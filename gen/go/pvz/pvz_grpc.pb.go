// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.6
// source: pvz/pvz.proto

package pvz_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	PVZService_GetPVZList_FullMethodName = "/pvz.v1.PVZService/GetPVZList"
)

// PVZServiceClient is the client API for PVZService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PVZServiceClient interface {
	GetPVZList(ctx context.Context, in *GetPVZListRequest, opts ...grpc.CallOption) (*GetPVZListResponse, error)
}

type pVZServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPVZServiceClient(cc grpc.ClientConnInterface) PVZServiceClient {
	return &pVZServiceClient{cc}
}

func (c *pVZServiceClient) GetPVZList(ctx context.Context, in *GetPVZListRequest, opts ...grpc.CallOption) (*GetPVZListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPVZListResponse)
	err := c.cc.Invoke(ctx, PVZService_GetPVZList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PVZServiceServer is the server API for PVZService service.
// All implementations must embed UnimplementedPVZServiceServer
// for forward compatibility
type PVZServiceServer interface {
	GetPVZList(context.Context, *GetPVZListRequest) (*GetPVZListResponse, error)
	mustEmbedUnimplementedPVZServiceServer()
}

// UnimplementedPVZServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPVZServiceServer struct {
}

func (UnimplementedPVZServiceServer) GetPVZList(context.Context, *GetPVZListRequest) (*GetPVZListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPVZList not implemented")
}
func (UnimplementedPVZServiceServer) mustEmbedUnimplementedPVZServiceServer() {}

// UnsafePVZServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PVZServiceServer will
// result in compilation errors.
type UnsafePVZServiceServer interface {
	mustEmbedUnimplementedPVZServiceServer()
}

func RegisterPVZServiceServer(s grpc.ServiceRegistrar, srv PVZServiceServer) {
	s.RegisterService(&PVZService_ServiceDesc, srv)
}

func _PVZService_GetPVZList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPVZListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PVZServiceServer).GetPVZList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PVZService_GetPVZList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PVZServiceServer).GetPVZList(ctx, req.(*GetPVZListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PVZService_ServiceDesc is the grpc.ServiceDesc for PVZService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PVZService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pvz.v1.PVZService",
	HandlerType: (*PVZServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPVZList",
			Handler:    _PVZService_GetPVZList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pvz/pvz.proto",
}
