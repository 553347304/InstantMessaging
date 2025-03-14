// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.9.0
// source: rpc/user.proto

package user_rpc

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

const (
	User_UserCreate_FullMethodName     = "/user_rpc.User/UserCreate"
	User_UserInfo_FullMethodName       = "/user_rpc.User/UserInfo"
	User_UserOnlineList_FullMethodName = "/user_rpc.User/UserOnlineList"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error)
	UserInfo(ctx context.Context, in *IdList, opts ...grpc.CallOption) (*UserInfoResponse, error)
	UserOnlineList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserOnlineListResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error) {
	out := new(UserCreateResponse)
	err := c.cc.Invoke(ctx, User_UserCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserInfo(ctx context.Context, in *IdList, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, User_UserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserOnlineList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserOnlineListResponse, error) {
	out := new(UserOnlineListResponse)
	err := c.cc.Invoke(ctx, User_UserOnlineList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	UserCreate(context.Context, *UserCreateRequest) (*UserCreateResponse, error)
	UserInfo(context.Context, *IdList) (*UserInfoResponse, error)
	UserOnlineList(context.Context, *Empty) (*UserOnlineListResponse, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) UserCreate(context.Context, *UserCreateRequest) (*UserCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCreate not implemented")
}
func (UnimplementedUserServer) UserInfo(context.Context, *IdList) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedUserServer) UserOnlineList(context.Context, *Empty) (*UserOnlineListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserOnlineList not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_UserCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserCreate(ctx, req.(*UserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserInfo(ctx, req.(*IdList))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserOnlineList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserOnlineList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserOnlineList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserOnlineList(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_rpc.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserCreate",
			Handler:    _User_UserCreate_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _User_UserInfo_Handler,
		},
		{
			MethodName: "UserOnlineList",
			Handler:    _User_UserOnlineList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/user.proto",
}

const (
	Friend_IsFriend_FullMethodName   = "/user_rpc.Friend/IsFriend"
	Friend_FriendList_FullMethodName = "/user_rpc.Friend/FriendList"
)

// FriendClient is the client API for Friend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FriendClient interface {
	IsFriend(ctx context.Context, in *IsFriendRequest, opts ...grpc.CallOption) (*Empty, error)
	FriendList(ctx context.Context, in *ID, opts ...grpc.CallOption) (*FriendListResponse, error)
}

type friendClient struct {
	cc grpc.ClientConnInterface
}

func NewFriendClient(cc grpc.ClientConnInterface) FriendClient {
	return &friendClient{cc}
}

func (c *friendClient) IsFriend(ctx context.Context, in *IsFriendRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Friend_IsFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) FriendList(ctx context.Context, in *ID, opts ...grpc.CallOption) (*FriendListResponse, error) {
	out := new(FriendListResponse)
	err := c.cc.Invoke(ctx, Friend_FriendList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FriendServer is the server API for Friend service.
// All implementations must embed UnimplementedFriendServer
// for forward compatibility
type FriendServer interface {
	IsFriend(context.Context, *IsFriendRequest) (*Empty, error)
	FriendList(context.Context, *ID) (*FriendListResponse, error)
	mustEmbedUnimplementedFriendServer()
}

// UnimplementedFriendServer must be embedded to have forward compatible implementations.
type UnimplementedFriendServer struct {
}

func (UnimplementedFriendServer) IsFriend(context.Context, *IsFriendRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsFriend not implemented")
}
func (UnimplementedFriendServer) FriendList(context.Context, *ID) (*FriendListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FriendList not implemented")
}
func (UnimplementedFriendServer) mustEmbedUnimplementedFriendServer() {}

// UnsafeFriendServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FriendServer will
// result in compilation errors.
type UnsafeFriendServer interface {
	mustEmbedUnimplementedFriendServer()
}

func RegisterFriendServer(s grpc.ServiceRegistrar, srv FriendServer) {
	s.RegisterService(&Friend_ServiceDesc, srv)
}

func _Friend_IsFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).IsFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_IsFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).IsFriend(ctx, req.(*IsFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_FriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).FriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_FriendList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).FriendList(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

// Friend_ServiceDesc is the grpc.ServiceDesc for Friend service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Friend_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_rpc.Friend",
	HandlerType: (*FriendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsFriend",
			Handler:    _Friend_IsFriend_Handler,
		},
		{
			MethodName: "FriendList",
			Handler:    _Friend_FriendList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/user.proto",
}

const (
	UserCurtail_IsCurtail_FullMethodName = "/user_rpc.UserCurtail/IsCurtail"
)

// UserCurtailClient is the client API for UserCurtail service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserCurtailClient interface {
	IsCurtail(ctx context.Context, in *ID, opts ...grpc.CallOption) (*CurtailResponse, error)
}

type userCurtailClient struct {
	cc grpc.ClientConnInterface
}

func NewUserCurtailClient(cc grpc.ClientConnInterface) UserCurtailClient {
	return &userCurtailClient{cc}
}

func (c *userCurtailClient) IsCurtail(ctx context.Context, in *ID, opts ...grpc.CallOption) (*CurtailResponse, error) {
	out := new(CurtailResponse)
	err := c.cc.Invoke(ctx, UserCurtail_IsCurtail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserCurtailServer is the server API for UserCurtail service.
// All implementations must embed UnimplementedUserCurtailServer
// for forward compatibility
type UserCurtailServer interface {
	IsCurtail(context.Context, *ID) (*CurtailResponse, error)
	mustEmbedUnimplementedUserCurtailServer()
}

// UnimplementedUserCurtailServer must be embedded to have forward compatible implementations.
type UnimplementedUserCurtailServer struct {
}

func (UnimplementedUserCurtailServer) IsCurtail(context.Context, *ID) (*CurtailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsCurtail not implemented")
}
func (UnimplementedUserCurtailServer) mustEmbedUnimplementedUserCurtailServer() {}

// UnsafeUserCurtailServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserCurtailServer will
// result in compilation errors.
type UnsafeUserCurtailServer interface {
	mustEmbedUnimplementedUserCurtailServer()
}

func RegisterUserCurtailServer(s grpc.ServiceRegistrar, srv UserCurtailServer) {
	s.RegisterService(&UserCurtail_ServiceDesc, srv)
}

func _UserCurtail_IsCurtail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCurtailServer).IsCurtail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCurtail_IsCurtail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCurtailServer).IsCurtail(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

// UserCurtail_ServiceDesc is the grpc.ServiceDesc for UserCurtail service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserCurtail_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_rpc.UserCurtail",
	HandlerType: (*UserCurtailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsCurtail",
			Handler:    _UserCurtail_IsCurtail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/user.proto",
}
