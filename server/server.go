package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/saskaradit/grpc-blog/blog"
	"github.com/saskaradit/grpc-blog/blogpb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// if we crash the code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Hello im the blog server")

	mgo := connectToMongo()

	// Connect to gRPC
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln("Failed to connect", err)
	}
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	blogpb.RegisterBlogServiceServer(s, &blog.Server{})

	// Register reflection service on gRPC server
	reflection.Register(s)

	go func() {
		fmt.Println("Starting Server")
		if err := s.Serve(lis); err != nil {
			log.Fatalln("Failed to serve", err)
		}
	}()
	// Wait for control C
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch
	fmt.Println("\nStopping the server")
	s.Stop()
	fmt.Println("Closing the listener")
	lis.Close()
	fmt.Println("Closing MongoDB Connection")
	mgo.Disconnect(context.TODO())
}

func connectToMongo() *mongo.Client {
	fmt.Println("Connecting to MongoDB")
	// Connect to mongo
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalln("failed to connect", err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatalln("Failed to connect client", err)
	}

	blog.Collection = client.Database("raddb").Collection("blog")

	return client
}
