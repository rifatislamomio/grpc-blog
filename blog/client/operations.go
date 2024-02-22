package main

import (
	"context"
	pb "gRPC/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
)

func CreateBlogPost(c pb.BlogServiceClient) string {
	log.Println("Create blog post was invoked")

	post := &pb.Blog{
		AuthorId: "Rft",
		Title:    "First blog",
		Content:  "This is my first blog post",
	}

	res, err := c.CreateBlog(context.Background(), post)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Blog post has been created ", res.Id)

	return res.Id
}

func ReadBlogPost(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("Read blog post was invoked")

	req := &pb.BlogId{
		Id: id,
	}

	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Fatal("Blog post not found ", err)
	}

	log.Println("Fetched blog post ", res)

	return res
}

func UpdateBlogPost(c pb.BlogServiceClient, id string) {
	log.Println("Update blog post was invoked")

	req := &pb.Blog{
		Id:       id,
		AuthorId: "Omio",
		Title:    "*Updated title*",
		Content:  "Blog post with updated content",
	}

	_, err := c.UpdateBlog(context.Background(), req)
	if err != nil {
		log.Fatal("Failed to update blog post ", err)
	}

	log.Println("Blog post updated")
}

func DeletePost(c pb.BlogServiceClient, id string) {
	log.Println("Delete blog post was invoked")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatal("Failed to delete blog post ", err)
	}

	log.Println("Blog post deleted")
}

func BlogList(c pb.BlogServiceClient) {
	log.Println("List blog was invoked")

	stream, err := c.ListBlog(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatal("Failed to stream list blogs")
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("Failed to receive message", err)
		}

		log.Println("Blog post-> ", msg)
	}
}
