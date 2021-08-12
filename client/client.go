package main

import (
	"fmt"
	"log"

	"github.com/saskaradit/grpc-blog/blog"
	"github.com/saskaradit/grpc-blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello im the blog client")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalln("Could not connect", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)
	blog.Create(c)
	// readBlog(c, "6114c3ea66cdc314f5dfbcdf")
	// updateBlog(c, "6114c3ea66cdc314f5dfbcdf")
	// deleteBlog(c, "6114c3ea66cdc314f5dfbcdf")
	blog.List(c)
}
