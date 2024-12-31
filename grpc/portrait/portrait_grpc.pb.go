// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: cmd/odas/grpc/portrait/portrait.proto

package portrait

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
	PortraitService_AgeSummary_FullMethodName        = "/PortraitService/AgeSummary"
	PortraitService_Fellow_FullMethodName            = "/PortraitService/Fellow"
	PortraitService_PaymentMethod_FullMethodName     = "/PortraitService/PaymentMethod"
	PortraitService_CityRank_FullMethodName          = "/PortraitService/CityRank"
	PortraitService_ProvinceRank_FullMethodName      = "/PortraitService/ProvinceRank"
	PortraitService_AgeCompareSummary_FullMethodName = "/PortraitService/AgeCompareSummary"
)

// PortraitServiceClient is the client API for PortraitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortraitServiceClient interface {
	// 定义方法 AgeSummary 接受 AgeRequest 消息，返回 QueryAgeSummaryBySidResponse 消息
	AgeSummary(ctx context.Context, in *AgeRequest, opts ...grpc.CallOption) (*AgeSummaryResponse, error)
	// 定义方法 Fellow 接受 Request 消息，返回 FellowResponse 消息
	Fellow(ctx context.Context, in *FellowRequest, opts ...grpc.CallOption) (*FellowResponse, error)
	// 定义方法 PaymentMethod 接受 Request 消息，返回 PaymentMethod 消息
	PaymentMethod(ctx context.Context, in *FellowRequest, opts ...grpc.CallOption) (*PaymentMethodResponse, error)
	// 定义方法 CityRankBySid 接受 CityRankRequest 消息，返回 CityRankResponse 消息
	CityRank(ctx context.Context, in *CityRankRequest, opts ...grpc.CallOption) (*CityRankResponse, error)
	// 定义方法 ProvinceRank 接受 ProvinceRankRequest 消息，返回 ProvinceRankResponse 消息
	ProvinceRank(ctx context.Context, in *ProvinceRankRequest, opts ...grpc.CallOption) (*ProvinceRankResponse, error)
	// 定义方法 AgeCompareSummary 接受 AgeRequest 消息，返回 AgeCompareSummaryBySidResponse 消息
	AgeCompareSummary(ctx context.Context, in *AgeCompareRequest, opts ...grpc.CallOption) (*AgeCompareSummaryResponse, error)
}

type portraitServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPortraitServiceClient(cc grpc.ClientConnInterface) PortraitServiceClient {
	return &portraitServiceClient{cc}
}

func (c *portraitServiceClient) AgeSummary(ctx context.Context, in *AgeRequest, opts ...grpc.CallOption) (*AgeSummaryResponse, error) {
	out := new(AgeSummaryResponse)
	err := c.cc.Invoke(ctx, PortraitService_AgeSummary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portraitServiceClient) Fellow(ctx context.Context, in *FellowRequest, opts ...grpc.CallOption) (*FellowResponse, error) {
	out := new(FellowResponse)
	err := c.cc.Invoke(ctx, PortraitService_Fellow_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portraitServiceClient) PaymentMethod(ctx context.Context, in *FellowRequest, opts ...grpc.CallOption) (*PaymentMethodResponse, error) {
	out := new(PaymentMethodResponse)
	err := c.cc.Invoke(ctx, PortraitService_PaymentMethod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portraitServiceClient) CityRank(ctx context.Context, in *CityRankRequest, opts ...grpc.CallOption) (*CityRankResponse, error) {
	out := new(CityRankResponse)
	err := c.cc.Invoke(ctx, PortraitService_CityRank_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portraitServiceClient) ProvinceRank(ctx context.Context, in *ProvinceRankRequest, opts ...grpc.CallOption) (*ProvinceRankResponse, error) {
	out := new(ProvinceRankResponse)
	err := c.cc.Invoke(ctx, PortraitService_ProvinceRank_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portraitServiceClient) AgeCompareSummary(ctx context.Context, in *AgeCompareRequest, opts ...grpc.CallOption) (*AgeCompareSummaryResponse, error) {
	out := new(AgeCompareSummaryResponse)
	err := c.cc.Invoke(ctx, PortraitService_AgeCompareSummary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortraitServiceServer is the server API for PortraitService service.
// All implementations must embed UnimplementedPortraitServiceServer
// for forward compatibility
type PortraitServiceServer interface {
	// 定义方法 AgeSummary 接受 AgeRequest 消息，返回 QueryAgeSummaryBySidResponse 消息
	AgeSummary(context.Context, *AgeRequest) (*AgeSummaryResponse, error)
	// 定义方法 Fellow 接受 Request 消息，返回 FellowResponse 消息
	Fellow(context.Context, *FellowRequest) (*FellowResponse, error)
	// 定义方法 PaymentMethod 接受 Request 消息，返回 PaymentMethod 消息
	PaymentMethod(context.Context, *FellowRequest) (*PaymentMethodResponse, error)
	// 定义方法 CityRankBySid 接受 CityRankRequest 消息，返回 CityRankResponse 消息
	CityRank(context.Context, *CityRankRequest) (*CityRankResponse, error)
	// 定义方法 ProvinceRank 接受 ProvinceRankRequest 消息，返回 ProvinceRankResponse 消息
	ProvinceRank(context.Context, *ProvinceRankRequest) (*ProvinceRankResponse, error)
	// 定义方法 AgeCompareSummary 接受 AgeRequest 消息，返回 AgeCompareSummaryBySidResponse 消息
	AgeCompareSummary(context.Context, *AgeCompareRequest) (*AgeCompareSummaryResponse, error)
	mustEmbedUnimplementedPortraitServiceServer()
}

// UnimplementedPortraitServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPortraitServiceServer struct {
}

func (UnimplementedPortraitServiceServer) AgeSummary(context.Context, *AgeRequest) (*AgeSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AgeSummary not implemented")
}
func (UnimplementedPortraitServiceServer) Fellow(context.Context, *FellowRequest) (*FellowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fellow not implemented")
}
func (UnimplementedPortraitServiceServer) PaymentMethod(context.Context, *FellowRequest) (*PaymentMethodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PaymentMethod not implemented")
}
func (UnimplementedPortraitServiceServer) CityRank(context.Context, *CityRankRequest) (*CityRankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CityRank not implemented")
}
func (UnimplementedPortraitServiceServer) ProvinceRank(context.Context, *ProvinceRankRequest) (*ProvinceRankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProvinceRank not implemented")
}
func (UnimplementedPortraitServiceServer) AgeCompareSummary(context.Context, *AgeCompareRequest) (*AgeCompareSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AgeCompareSummary not implemented")
}
func (UnimplementedPortraitServiceServer) mustEmbedUnimplementedPortraitServiceServer() {}

// UnsafePortraitServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortraitServiceServer will
// result in compilation errors.
type UnsafePortraitServiceServer interface {
	mustEmbedUnimplementedPortraitServiceServer()
}

func RegisterPortraitServiceServer(s grpc.ServiceRegistrar, srv PortraitServiceServer) {
	s.RegisterService(&PortraitService_ServiceDesc, srv)
}

func _PortraitService_AgeSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortraitServiceServer).AgeSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortraitService_AgeSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortraitServiceServer).AgeSummary(ctx, req.(*AgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortraitService_Fellow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FellowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortraitServiceServer).Fellow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortraitService_Fellow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortraitServiceServer).Fellow(ctx, req.(*FellowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortraitService_PaymentMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FellowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortraitServiceServer).PaymentMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortraitService_PaymentMethod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortraitServiceServer).PaymentMethod(ctx, req.(*FellowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortraitService_CityRank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CityRankRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortraitServiceServer).CityRank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortraitService_CityRank_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortraitServiceServer).CityRank(ctx, req.(*CityRankRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortraitService_ProvinceRank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProvinceRankRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortraitServiceServer).ProvinceRank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortraitService_ProvinceRank_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortraitServiceServer).ProvinceRank(ctx, req.(*ProvinceRankRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortraitService_AgeCompareSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgeCompareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortraitServiceServer).AgeCompareSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortraitService_AgeCompareSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortraitServiceServer).AgeCompareSummary(ctx, req.(*AgeCompareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PortraitService_ServiceDesc is the grpc.ServiceDesc for PortraitService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PortraitService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PortraitService",
	HandlerType: (*PortraitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AgeSummary",
			Handler:    _PortraitService_AgeSummary_Handler,
		},
		{
			MethodName: "Fellow",
			Handler:    _PortraitService_Fellow_Handler,
		},
		{
			MethodName: "PaymentMethod",
			Handler:    _PortraitService_PaymentMethod_Handler,
		},
		{
			MethodName: "CityRank",
			Handler:    _PortraitService_CityRank_Handler,
		},
		{
			MethodName: "ProvinceRank",
			Handler:    _PortraitService_ProvinceRank_Handler,
		},
		{
			MethodName: "AgeCompareSummary",
			Handler:    _PortraitService_AgeCompareSummary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmd/odas/grpc/portrait/portrait.proto",
}