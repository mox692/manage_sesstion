// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package session

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// SessionClient is the client API for Session service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SessionClient interface {
	SetSession(ctx context.Context, in *SessionRequest, opts ...grpc.CallOption) (*SetStatus, error)
	GetSession(ctx context.Context, in *SessionRequest, opts ...grpc.CallOption) (*GetStatus, error)
}

type sessionClient struct {
	cc grpc.ClientConnInterface
}

func NewSessionClient(cc grpc.ClientConnInterface) SessionClient {
	return &sessionClient{cc}
}

func (c *sessionClient) SetSession(ctx context.Context, in *SessionRequest, opts ...grpc.CallOption) (*SetStatus, error) {
	out := new(SetStatus)
	err := c.cc.Invoke(ctx, "/session.Session/SetSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionClient) GetSession(ctx context.Context, in *SessionRequest, opts ...grpc.CallOption) (*GetStatus, error) {
	out := new(GetStatus)
	err := c.cc.Invoke(ctx, "/session.Session/GetSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionServer is the server API for Session service.
// All implementations must embed UnimplementedSessionServer
// for forward compatibility
type SessionServer interface {
	SetSession(context.Context, *SessionRequest) (*SetStatus, error)
	GetSession(context.Context, *SessionRequest) (*GetStatus, error)
	mustEmbedUnimplementedSessionServer()
}

// UnimplementedSessionServer must be embedded to have forward compatible implementations.
type UnimplementedSessionServer struct {
}

func (UnimplementedSessionServer) SetSession(context.Context, *SessionRequest) (*SetStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSession not implemented")
}
func (UnimplementedSessionServer) GetSession(context.Context, *SessionRequest) (*GetStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSession not implemented")
}
func (UnimplementedSessionServer) mustEmbedUnimplementedSessionServer() {}

// UnsafeSessionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SessionServer will
// result in compilation errors.
type UnsafeSessionServer interface {
	mustEmbedUnimplementedSessionServer()
}

func RegisterSessionServer(s grpc.ServiceRegistrar, srv SessionServer) {
	s.RegisterService(&_Session_serviceDesc, srv)
}

func _Session_SetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServer).SetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.Session/SetSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServer).SetSession(ctx, req.(*SessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Session_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServer).GetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.Session/GetSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServer).GetSession(ctx, req.(*SessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Session_serviceDesc = grpc.ServiceDesc{
	ServiceName: "session.Session",
	HandlerType: (*SessionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetSession",
			Handler:    _Session_SetSession_Handler,
		},
		{
			MethodName: "GetSession",
			Handler:    _Session_GetSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "session.proto",
}
