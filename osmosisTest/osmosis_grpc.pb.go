// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: osmosisTest/osmosis.proto

package lavaPB

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
	Service_GetBlockByHeight_FullMethodName        = "/cosmos.base.tendermint.v1beta1.Service/GetBlockByHeight"
	Service_GetLatestBlock_FullMethodName          = "/cosmos.base.tendermint.v1beta1.Service/GetLatestBlock"
	Service_GetLatestValidatorSet_FullMethodName   = "/cosmos.base.tendermint.v1beta1.Service/GetLatestValidatorSet"
	Service_GetNodeInfo_FullMethodName             = "/cosmos.base.tendermint.v1beta1.Service/GetNodeInfo"
	Service_GetSyncing_FullMethodName              = "/cosmos.base.tendermint.v1beta1.Service/GetSyncing"
	Service_GetValidatorSetByHeight_FullMethodName = "/cosmos.base.tendermint.v1beta1.Service/GetValidatorSetByHeight"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	GetBlockByHeight(ctx context.Context, in *GetHeightRequest, opts ...grpc.CallOption) (*GetBlockByHeightResponse, error)
	GetLatestBlock(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetLatestBlockResponse, error)
	GetLatestValidatorSet(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetLatestValidatorSetResponse, error)
	GetNodeInfo(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetNodeInfoResponse, error)
	GetSyncing(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetSyncingResponse, error)
	GetValidatorSetByHeight(ctx context.Context, in *GetHeightRequest, opts ...grpc.CallOption) (*GetValidatorSetByHeightResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) GetBlockByHeight(ctx context.Context, in *GetHeightRequest, opts ...grpc.CallOption) (*GetBlockByHeightResponse, error) {
	out := new(GetBlockByHeightResponse)
	err := c.cc.Invoke(ctx, Service_GetBlockByHeight_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetLatestBlock(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetLatestBlockResponse, error) {
	out := new(GetLatestBlockResponse)
	err := c.cc.Invoke(ctx, Service_GetLatestBlock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetLatestValidatorSet(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetLatestValidatorSetResponse, error) {
	out := new(GetLatestValidatorSetResponse)
	err := c.cc.Invoke(ctx, Service_GetLatestValidatorSet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetNodeInfo(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetNodeInfoResponse, error) {
	out := new(GetNodeInfoResponse)
	err := c.cc.Invoke(ctx, Service_GetNodeInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetSyncing(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetSyncingResponse, error) {
	out := new(GetSyncingResponse)
	err := c.cc.Invoke(ctx, Service_GetSyncing_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetValidatorSetByHeight(ctx context.Context, in *GetHeightRequest, opts ...grpc.CallOption) (*GetValidatorSetByHeightResponse, error) {
	out := new(GetValidatorSetByHeightResponse)
	err := c.cc.Invoke(ctx, Service_GetValidatorSetByHeight_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	GetBlockByHeight(context.Context, *GetHeightRequest) (*GetBlockByHeightResponse, error)
	GetLatestBlock(context.Context, *GetRequest) (*GetLatestBlockResponse, error)
	GetLatestValidatorSet(context.Context, *GetRequest) (*GetLatestValidatorSetResponse, error)
	GetNodeInfo(context.Context, *GetRequest) (*GetNodeInfoResponse, error)
	GetSyncing(context.Context, *GetRequest) (*GetSyncingResponse, error)
	GetValidatorSetByHeight(context.Context, *GetHeightRequest) (*GetValidatorSetByHeightResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) GetBlockByHeight(context.Context, *GetHeightRequest) (*GetBlockByHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockByHeight not implemented")
}
func (UnimplementedServiceServer) GetLatestBlock(context.Context, *GetRequest) (*GetLatestBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestBlock not implemented")
}
func (UnimplementedServiceServer) GetLatestValidatorSet(context.Context, *GetRequest) (*GetLatestValidatorSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestValidatorSet not implemented")
}
func (UnimplementedServiceServer) GetNodeInfo(context.Context, *GetRequest) (*GetNodeInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeInfo not implemented")
}
func (UnimplementedServiceServer) GetSyncing(context.Context, *GetRequest) (*GetSyncingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSyncing not implemented")
}
func (UnimplementedServiceServer) GetValidatorSetByHeight(context.Context, *GetHeightRequest) (*GetValidatorSetByHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValidatorSetByHeight not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_GetBlockByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetBlockByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetBlockByHeight_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetBlockByHeight(ctx, req.(*GetHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetLatestBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetLatestBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetLatestBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetLatestBlock(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetLatestValidatorSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetLatestValidatorSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetLatestValidatorSet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetLatestValidatorSet(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetNodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetNodeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetNodeInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetNodeInfo(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetSyncing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetSyncing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetSyncing_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetSyncing(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetValidatorSetByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetValidatorSetByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetValidatorSetByHeight_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetValidatorSetByHeight(ctx, req.(*GetHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.base.tendermint.v1beta1.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlockByHeight",
			Handler:    _Service_GetBlockByHeight_Handler,
		},
		{
			MethodName: "GetLatestBlock",
			Handler:    _Service_GetLatestBlock_Handler,
		},
		{
			MethodName: "GetLatestValidatorSet",
			Handler:    _Service_GetLatestValidatorSet_Handler,
		},
		{
			MethodName: "GetNodeInfo",
			Handler:    _Service_GetNodeInfo_Handler,
		},
		{
			MethodName: "GetSyncing",
			Handler:    _Service_GetSyncing_Handler,
		},
		{
			MethodName: "GetValidatorSetByHeight",
			Handler:    _Service_GetValidatorSetByHeight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "osmosisTest/osmosis.proto",
}
