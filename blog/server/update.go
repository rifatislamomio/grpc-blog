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

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	log.Println("Update blog was invoked with", in)
	data := BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Could not parse to ID")
	}

	filter := bson.M{"_id": oid}
	res, mErr := collection.UpdateOne(ctx, filter, bson.M{"$set": data})
	if mErr != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update blog post")
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Blog post not found with that id to update",
		)
	}

	log.Println("Blog post updated ")

	return &emptypb.Empty{}, nil
}
