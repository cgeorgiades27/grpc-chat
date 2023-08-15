// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: pkg/api/v1/chat.proto

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

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatClient interface {
	SendMessage(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatResponse, error)
	StreamMessage(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (Chat_StreamMessageClient, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) SendMessage(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatResponse, error) {
	out := new(ChatResponse)
	err := c.cc.Invoke(ctx, "/chat.v1.Chat/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) StreamMessage(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (Chat_StreamMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chat_ServiceDesc.Streams[0], "/chat.v1.Chat/StreamMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatStreamMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chat_StreamMessageClient interface {
	Recv() (*ChatResponse, error)
	grpc.ClientStream
}

type chatStreamMessageClient struct {
	grpc.ClientStream
}

func (x *chatStreamMessageClient) Recv() (*ChatResponse, error) {
	m := new(ChatResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServer is the server API for Chat service.
// All implementations must embed UnimplementedChatServer
// for forward compatibility
type ChatServer interface {
	SendMessage(context.Context, *ChatRequest) (*ChatResponse, error)
	StreamMessage(*ChatRequest, Chat_StreamMessageServer) error
	mustEmbedUnimplementedChatServer()
}

// UnimplementedChatServer must be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (UnimplementedChatServer) SendMessage(context.Context, *ChatRequest) (*ChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatServer) StreamMessage(*ChatRequest, Chat_StreamMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamMessage not implemented")
}
func (UnimplementedChatServer) mustEmbedUnimplementedChatServer() {}

// UnsafeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServer will
// result in compilation errors.
type UnsafeChatServer interface {
	mustEmbedUnimplementedChatServer()
}

func RegisterChatServer(s grpc.ServiceRegistrar, srv ChatServer) {
	s.RegisterService(&Chat_ServiceDesc, srv)
}

func _Chat_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.v1.Chat/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).SendMessage(ctx, req.(*ChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_StreamMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ChatRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServer).StreamMessage(m, &chatStreamMessageServer{stream})
}

type Chat_StreamMessageServer interface {
	Send(*ChatResponse) error
	grpc.ServerStream
}

type chatStreamMessageServer struct {
	grpc.ServerStream
}

func (x *chatStreamMessageServer) Send(m *ChatResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Chat_ServiceDesc is the grpc.ServiceDesc for Chat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.v1.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _Chat_SendMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamMessage",
			Handler:       _Chat_StreamMessage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/api/v1/chat.proto",
}
