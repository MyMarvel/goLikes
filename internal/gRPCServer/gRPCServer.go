package gRPCServer

import (
	"context"

	"likes_handler/tools/proto"
)

type GRPCServer struct {
	proto.UnimplementedLikesServer
}

func (s *GRPCServer) DoLike(ctx context.Context, in *proto.DoLikeParams) (*proto.GetLikeCountReply, error) {
	return &proto.GetLikeCountReply{Likes: "1"}, nil
}

func (s *GRPCServer) GetLikeCount(ctx context.Context, in *proto.GetLikeCountParams) (*proto.GetLikeCountReply, error) {
	return &proto.GetLikeCountReply{Likes: "1"}, nil
}
