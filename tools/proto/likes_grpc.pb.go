// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: likes.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LikesClient is the client API for Likes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LikesClient interface {
	DoLike(ctx context.Context, in *DoLikeParams, opts ...grpc.CallOption) (*GetLikeCountReply, error)
	GetLikeCount(ctx context.Context, in *GetLikeCountParams, opts ...grpc.CallOption) (*GetLikeCountReply, error)
}

type likesClient struct {
	cc grpc.ClientConnInterface
}

func NewLikesClient(cc grpc.ClientConnInterface) LikesClient {
	return &likesClient{cc}
}

func (c *likesClient) DoLike(ctx context.Context, in *DoLikeParams, opts ...grpc.CallOption) (*GetLikeCountReply, error) {
	out := new(GetLikeCountReply)
	err := c.cc.Invoke(ctx, "/likes.Likes/doLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likesClient) GetLikeCount(ctx context.Context, in *GetLikeCountParams, opts ...grpc.CallOption) (*GetLikeCountReply, error) {
	out := new(GetLikeCountReply)
	err := c.cc.Invoke(ctx, "/likes.Likes/getLikeCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LikesServer is the server API for Likes service.
// All implementations must embed UnimplementedLikesServer
// for forward compatibility
type LikesServer interface {
	DoLike(context.Context, *DoLikeParams) (*GetLikeCountReply, error)
	GetLikeCount(context.Context, *GetLikeCountParams) (*GetLikeCountReply, error)
	mustEmbedUnimplementedLikesServer()
}

// UnimplementedLikesServer must be embedded to have forward compatible implementations.
type UnimplementedLikesServer struct {
}

func (UnimplementedLikesServer) DoLike(context.Context, *DoLikeParams) (*GetLikeCountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoLike not implemented")
}
func (UnimplementedLikesServer) GetLikeCount(context.Context, *GetLikeCountParams) (*GetLikeCountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLikeCount not implemented")
}
func (UnimplementedLikesServer) mustEmbedUnimplementedLikesServer() {}

// UnsafeLikesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LikesServer will
// result in compilation errors.
type UnsafeLikesServer interface {
	mustEmbedUnimplementedLikesServer()
}

func RegisterLikesServer(s grpc.ServiceRegistrar, srv LikesServer) {
	s.RegisterService(&Likes_ServiceDesc, srv)
}

func _Likes_DoLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoLikeParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikesServer).DoLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/likes.Likes/doLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikesServer).DoLike(ctx, req.(*DoLikeParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Likes_GetLikeCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLikeCountParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikesServer).GetLikeCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/likes.Likes/getLikeCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikesServer).GetLikeCount(ctx, req.(*GetLikeCountParams))
	}
	return interceptor(ctx, in, info, handler)
}

// Likes_ServiceDesc is the grpc.ServiceDesc for Likes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Likes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "likes.Likes",
	HandlerType: (*LikesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "doLike",
			Handler:    _Likes_DoLike_Handler,
		},
		{
			MethodName: "getLikeCount",
			Handler:    _Likes_GetLikeCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "likes.proto",
}
