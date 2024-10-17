// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.11.4
// source: github.com/moby/buildkit/session/upload/upload.proto

package upload

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
	Upload_Pull_FullMethodName = "/moby.upload.v1.Upload/Pull"
)

// UploadClient is the client API for Upload service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UploadClient interface {
	Pull(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[BytesMessage, BytesMessage], error)
}

type uploadClient struct {
	cc grpc.ClientConnInterface
}

func NewUploadClient(cc grpc.ClientConnInterface) UploadClient {
	return &uploadClient{cc}
}

func (c *uploadClient) Pull(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[BytesMessage, BytesMessage], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Upload_ServiceDesc.Streams[0], Upload_Pull_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[BytesMessage, BytesMessage]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Upload_PullClient = grpc.BidiStreamingClient[BytesMessage, BytesMessage]

// UploadServer is the server API for Upload service.
// All implementations should embed UnimplementedUploadServer
// for forward compatibility.
type UploadServer interface {
	Pull(grpc.BidiStreamingServer[BytesMessage, BytesMessage]) error
}

// UnimplementedUploadServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUploadServer struct{}

func (UnimplementedUploadServer) Pull(grpc.BidiStreamingServer[BytesMessage, BytesMessage]) error {
	return status.Errorf(codes.Unimplemented, "method Pull not implemented")
}
func (UnimplementedUploadServer) testEmbeddedByValue() {}

// UnsafeUploadServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UploadServer will
// result in compilation errors.
type UnsafeUploadServer interface {
	mustEmbedUnimplementedUploadServer()
}

func RegisterUploadServer(s grpc.ServiceRegistrar, srv UploadServer) {
	// If the following call pancis, it indicates UnimplementedUploadServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Upload_ServiceDesc, srv)
}

func _Upload_Pull_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UploadServer).Pull(&grpc.GenericServerStream[BytesMessage, BytesMessage]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Upload_PullServer = grpc.BidiStreamingServer[BytesMessage, BytesMessage]

// Upload_ServiceDesc is the grpc.ServiceDesc for Upload service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Upload_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "moby.upload.v1.Upload",
	HandlerType: (*UploadServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Pull",
			Handler:       _Upload_Pull_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/moby/buildkit/session/upload/upload.proto",
}