package main

import (
	pb "gRPC/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const address = "127.0.0.1:50051"

func main() {
	connection, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer connection.Close()

	client := pb.NewBlogServiceClient(connection)

	postId := CreateBlogPost(client) //create a blog post
	ReadBlogPost(client, postId)     //read the newly created blog post

	UpdateBlogPost(client, postId) //update the newly create blog post
	ReadBlogPost(client, postId)   //read the updated blog post

	DeletePost(client, postId) //delete the blog post

	//create two more blog posts
	CreateBlogPost(client)
	CreateBlogPost(client)

	//read blog list
	BlogList(client)
}
