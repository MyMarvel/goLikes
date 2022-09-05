package gRPCServer

import (
	"context"
	"errors"

	"likes_handler/internal/controllers"
	"likes_handler/tools/proto"
)

var controllersFactory controllers.IFactory

func InitControllersFactory(initFactory controllers.IFactory) {
	controllersFactory = initFactory
}

type GRPCServer struct {
	proto.UnimplementedLikesServer
}

func (s *GRPCServer) DoLike(ctx context.Context, in *proto.DoLikeParams) (*proto.GetLikeCountReply, error) {
	var acc controllers.IAccount = controllersFactory.NewIAccount(in.CurrentAccount)
	isAccountAlreadyLikes, err := acc.IsAccountLikes(in.TargetAccount)
	if err == nil {
		if !isAccountAlreadyLikes {
			likesCount, err := acc.IncrementLikeCounter(in.TargetAccount)
			return &proto.GetLikeCountReply{Likes: int64(likesCount)}, err
		} else {
			return &proto.GetLikeCountReply{}, errors.New("Already liked")
		}
	}
	return &proto.GetLikeCountReply{}, err
}

func (s *GRPCServer) GetLikeCount(ctx context.Context, in *proto.GetLikeCountParams) (*proto.GetLikeCountReply, error) {
	var acc controllers.IAccount = controllersFactory.NewIAccount(in.TargetAccount)
	likesCount, err := acc.GetLikesCount(in.TargetAccount)
	return &proto.GetLikeCountReply{Likes: int64(likesCount)}, err
}
