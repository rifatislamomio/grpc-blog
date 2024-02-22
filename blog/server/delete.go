package main

import (
	"context"
	pb "gRPC/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {
	log.Println("Delete blog was invoked with", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Cannot converted to OID")
	}

	filter := bson.M{"_id": oid}

	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil,
			status.Errorf(codes.Internal, "Unexpected error occurred")
	}

	if res.DeletedCount != 1 {
		return nil,
			status.Errorf(codes.NotFound, "Failed to delete blog post as it was not found with that id")
	}

	log.Println("Blog post deleted ")

	return &emptypb.Empty{}, nil
}
