package blog

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/saskaradit/grpc-blog/blogpb"
)

// Create function creates a blog to the mongo
func Create(c blogpb.BlogServiceClient) {
	blog := &blogpb.Blog{
		AuthorId: "Rad",
		Title:    "Rad First Blog",
		Content:  "Content of the first blog",
	}
	res, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalln("Unexpected error:", err)
	}
	fmt.Println("Blog has been created", res)
}

// Read function retrieves a single blog
func Read(c blogpb.BlogServiceClient, blogID string) {
	_, err := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: blogID})
	if err != nil {
		fmt.Println("Error happenned while reading", err)
	}
	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogID}
	res, err := c.ReadBlog(context.Background(), readBlogReq)
	if err != nil {
		fmt.Println("Error happened while reading", err)
	}
	fmt.Println("Blog was read", res)
}

// Delete function deletes a single blog
func Delete(c blogpb.BlogServiceClient, blogID string) {
	res, err := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: blogID})
	if err != nil {
		fmt.Println("Something happenned whle deleting", err)
	}
	fmt.Println("Successfully deleted", res)
}

// Update function updates a single blog
func Update(c blogpb.BlogServiceClient, blogID string) {
	newBlog := &blogpb.Blog{
		Id:       blogID,
		AuthorId: "Change Author",
		Title:    "Next Blog",
		Content:  "Content of the first blog",
	}
	res, err := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: newBlog})
	if err != nil {
		fmt.Println("Error while updating", err)
	}
	fmt.Println("Blog was updated", res)
}

// List function lists all the blogs on the database
func List(c blogpb.BlogServiceClient) {
	stream, err := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
	if err != nil {
		log.Fatalln("error while calling ListBlog RPC", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Something happenned", err)
		}
		fmt.Println(res.GetBlog())
	}
}
