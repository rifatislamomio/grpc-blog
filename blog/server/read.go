package main

import (
	"context"
	"fmt"
	pb "gRPC/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Println("Read blog was invoked with", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Could not parse to ID")
	}

	data := &BlogItem{}
	filter := bson.M{"_id": oid}
	res := collection.FindOne(ctx, filter)

	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not find blog post")
	}

	if err != nil {
		return nil, status.Errorf(
			codes.Internal, fmt.Sprintf("Internal Error: %v", err))
	}

	return documentToBlog(data), nil
}
