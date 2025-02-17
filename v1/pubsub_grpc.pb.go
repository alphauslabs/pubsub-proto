// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/v1/pubsub.proto

package pubsubproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PubSubService_Publish_FullMethodName                 = "/pubsubproto.PubSubService/Publish"
	PubSubService_Subscribe_FullMethodName               = "/pubsubproto.PubSubService/Subscribe"
	PubSubService_Acknowledge_FullMethodName             = "/pubsubproto.PubSubService/Acknowledge"
	PubSubService_ModifyVisibilityTimeout_FullMethodName = "/pubsubproto.PubSubService/ModifyVisibilityTimeout"
)

// PubSubServiceClient is the client API for PubSubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PubSubServiceClient interface {
	// Publish a message to a topic
	Publish(ctx context.Context, in *Message, opts ...grpc.CallOption) (*PublishResponse, error)
	// Subscribe to a topic and receive messages
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Message], error)
	// Acknowledge receipt of a message
	Acknowledge(ctx context.Context, in *AcknowledgeRequest, opts ...grpc.CallOption) (*AcknowledgeResponse, error)
	// Modify the visibility timeout of a message
	ModifyVisibilityTimeout(ctx context.Context, in *ModifyVisibilityTimeoutRequest, opts ...grpc.CallOption) (*ModifyVisibilityTimeoutResponse, error)
}

type pubSubServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPubSubServiceClient(cc grpc.ClientConnInterface) PubSubServiceClient {
	return &pubSubServiceClient{cc}
}

func (c *pubSubServiceClient) Publish(ctx context.Context, in *Message, opts ...grpc.CallOption) (*PublishResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublishResponse)
	err := c.cc.Invoke(ctx, PubSubService_Publish_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubSubServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Message], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &PubSubService_ServiceDesc.Streams[0], PubSubService_Subscribe_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SubscribeRequest, Message]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type PubSubService_SubscribeClient = grpc.ServerStreamingClient[Message]

func (c *pubSubServiceClient) Acknowledge(ctx context.Context, in *AcknowledgeRequest, opts ...grpc.CallOption) (*AcknowledgeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AcknowledgeResponse)
	err := c.cc.Invoke(ctx, PubSubService_Acknowledge_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubSubServiceClient) ModifyVisibilityTimeout(ctx context.Context, in *ModifyVisibilityTimeoutRequest, opts ...grpc.CallOption) (*ModifyVisibilityTimeoutResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ModifyVisibilityTimeoutResponse)
	err := c.cc.Invoke(ctx, PubSubService_ModifyVisibilityTimeout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PubSubServiceServer is the server API for PubSubService service.
// All implementations must embed UnimplementedPubSubServiceServer
// for forward compatibility.
type PubSubServiceServer interface {
	// Publish a message to a topic
	Publish(context.Context, *Message) (*PublishResponse, error)
	// Subscribe to a topic and receive messages
	Subscribe(*SubscribeRequest, grpc.ServerStreamingServer[Message]) error
	// Acknowledge receipt of a message
	Acknowledge(context.Context, *AcknowledgeRequest) (*AcknowledgeResponse, error)
	// Modify the visibility timeout of a message
	ModifyVisibilityTimeout(context.Context, *ModifyVisibilityTimeoutRequest) (*ModifyVisibilityTimeoutResponse, error)
	mustEmbedUnimplementedPubSubServiceServer()
}

// UnimplementedPubSubServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPubSubServiceServer struct{}

func (UnimplementedPubSubServiceServer) Publish(context.Context, *Message) (*PublishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedPubSubServiceServer) Subscribe(*SubscribeRequest, grpc.ServerStreamingServer[Message]) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedPubSubServiceServer) Acknowledge(context.Context, *AcknowledgeRequest) (*AcknowledgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Acknowledge not implemented")
}
func (UnimplementedPubSubServiceServer) ModifyVisibilityTimeout(context.Context, *ModifyVisibilityTimeoutRequest) (*ModifyVisibilityTimeoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyVisibilityTimeout not implemented")
}
func (UnimplementedPubSubServiceServer) mustEmbedUnimplementedPubSubServiceServer() {}
func (UnimplementedPubSubServiceServer) testEmbeddedByValue()                       {}

// UnsafePubSubServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PubSubServiceServer will
// result in compilation errors.
type UnsafePubSubServiceServer interface {
	mustEmbedUnimplementedPubSubServiceServer()
}

func RegisterPubSubServiceServer(s grpc.ServiceRegistrar, srv PubSubServiceServer) {
	// If the following call pancis, it indicates UnimplementedPubSubServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PubSubService_ServiceDesc, srv)
}

func _PubSubService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSubServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubSubService_Publish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSubServiceServer).Publish(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubSubService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PubSubServiceServer).Subscribe(m, &grpc.GenericServerStream[SubscribeRequest, Message]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type PubSubService_SubscribeServer = grpc.ServerStreamingServer[Message]

func _PubSubService_Acknowledge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcknowledgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSubServiceServer).Acknowledge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubSubService_Acknowledge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSubServiceServer).Acknowledge(ctx, req.(*AcknowledgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubSubService_ModifyVisibilityTimeout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyVisibilityTimeoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSubServiceServer).ModifyVisibilityTimeout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubSubService_ModifyVisibilityTimeout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSubServiceServer).ModifyVisibilityTimeout(ctx, req.(*ModifyVisibilityTimeoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PubSubService_ServiceDesc is the grpc.ServiceDesc for PubSubService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PubSubService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pubsubproto.PubSubService",
	HandlerType: (*PubSubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _PubSubService_Publish_Handler,
		},
		{
			MethodName: "Acknowledge",
			Handler:    _PubSubService_Acknowledge_Handler,
		},
		{
			MethodName: "ModifyVisibilityTimeout",
			Handler:    _PubSubService_ModifyVisibilityTimeout_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _PubSubService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/v1/pubsub.proto",
}
