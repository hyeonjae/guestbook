// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// GuestbookClient is the client API for Guestbook service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GuestbookClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type guestbookClient struct {
	cc grpc.ClientConnInterface
}

func NewGuestbookClient(cc grpc.ClientConnInterface) GuestbookClient {
	return &guestbookClient{cc}
}

func (c *guestbookClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/hyeonjae.guestbook.v1.Guestbook/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guestbookClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/hyeonjae.guestbook.v1.Guestbook/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guestbookClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/hyeonjae.guestbook.v1.Guestbook/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GuestbookServer is the server API for Guestbook service.
// All implementations must embed UnimplementedGuestbookServer
// for forward compatibility
type GuestbookServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedGuestbookServer()
}

// UnimplementedGuestbookServer must be embedded to have forward compatible implementations.
type UnimplementedGuestbookServer struct {
}

func (UnimplementedGuestbookServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedGuestbookServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedGuestbookServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedGuestbookServer) mustEmbedUnimplementedGuestbookServer() {}

// UnsafeGuestbookServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GuestbookServer will
// result in compilation errors.
type UnsafeGuestbookServer interface {
	mustEmbedUnimplementedGuestbookServer()
}

func RegisterGuestbookServer(s grpc.ServiceRegistrar, srv GuestbookServer) {
	s.RegisterService(&Guestbook_ServiceDesc, srv)
}

func _Guestbook_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestbookServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hyeonjae.guestbook.v1.Guestbook/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestbookServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Guestbook_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestbookServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hyeonjae.guestbook.v1.Guestbook/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestbookServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Guestbook_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestbookServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hyeonjae.guestbook.v1.Guestbook/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestbookServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Guestbook_ServiceDesc is the grpc.ServiceDesc for Guestbook service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Guestbook_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hyeonjae.guestbook.v1.Guestbook",
	HandlerType: (*GuestbookServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Guestbook_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Guestbook_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Guestbook_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "guestbook.proto",
}
