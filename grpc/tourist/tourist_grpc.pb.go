// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: cmd/odas/grpc/tourist/tourist.proto

package tourist

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	grpc1 "odas/cmd/odas/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TouristService_SummaryByTime_FullMethodName             = "/TouristService/SummaryByTime"
	TouristService_SummaryByDate_FullMethodName             = "/TouristService/SummaryByDate"
	TouristService_ForecastPassengerFlowList_FullMethodName = "/TouristService/ForecastPassengerFlowList"
	TouristService_TouristLocal_FullMethodName              = "/TouristService/TouristLocal"
	TouristService_PassengerFlow_FullMethodName             = "/TouristService/PassengerFlow"
	TouristService_GroupList_FullMethodName                 = "/TouristService/GroupList"
	TouristService_InOutStat_FullMethodName                 = "/TouristService/InOutStat"
	TouristService_CreateGroup_FullMethodName               = "/TouristService/CreateGroup"
	TouristService_UpdateGroup_FullMethodName               = "/TouristService/UpdateGroup"
	TouristService_GroupById_FullMethodName                 = "/TouristService/GroupById"
	TouristService_InoutByGroups_FullMethodName             = "/TouristService/InoutByGroups"
	TouristService_DeviceJournalStat_FullMethodName         = "/TouristService/DeviceJournalStat"
	TouristService_InOutCompareFlow_FullMethodName          = "/TouristService/InOutCompareFlow"
)

// TouristServiceClient is the client API for TouristService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TouristServiceClient interface {
	// 获取按时间分组的出入园数据
	SummaryByTime(ctx context.Context, in *TouristSummaryRequest, opts ...grpc.CallOption) (*TouristSummaryResponse, error)
	// 按天获取出入园数据
	SummaryByDate(ctx context.Context, in *TouristSummaryRequest, opts ...grpc.CallOption) (*TouristSummaryResponse, error)
	// 客流预测每日以及今，明，后天数据
	ForecastPassengerFlowList(ctx context.Context, in *grpc1.PassedTimeSpanByOrderTypeV4Request, opts ...grpc.CallOption) (*ForecastPassengerFlowListResponse, error)
	// 供应商维度客流来源TopN
	TouristLocal(ctx context.Context, in *TouristLocalRequest, opts ...grpc.CallOption) (*TouristLocalResponse, error)
	// 每日客流数据(供应商维度)
	PassengerFlow(ctx context.Context, in *TouristStartEndRequest, opts ...grpc.CallOption) (*PassengerFlowResponse, error)
	// 获取统计组
	GroupList(ctx context.Context, in *GroupListRequest, opts ...grpc.CallOption) (*GroupListResponse, error)
	// 实时客流统计
	InOutStat(ctx context.Context, in *InOutStatRequest, opts ...grpc.CallOption) (*InOutStatResponse, error)
	// 创建出入园统计组
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error)
	// 更新出入园统计组
	UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*grpc1.Response, error)
	// 根据id查询统计组信息
	GroupById(ctx context.Context, in *GroupByIdRequest, opts ...grpc.CallOption) (*GroupListData, error)
	// 批量统计组获取出入园数据
	InoutByGroups(ctx context.Context, in *InoutByGroupsRequest, opts ...grpc.CallOption) (*InoutByGroupsResponse, error)
	// 终端管理出入、验证数据查询
	DeviceJournalStat(ctx context.Context, in *DeviceJournalStatRequest, opts ...grpc.CallOption) (*DeviceJournalStatResponse, error)
	// 实时对比客流数据
	InOutCompareFlow(ctx context.Context, in *InOutCompareRequest, opts ...grpc.CallOption) (*InoutCompareResponse, error)
}

type touristServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTouristServiceClient(cc grpc.ClientConnInterface) TouristServiceClient {
	return &touristServiceClient{cc}
}

func (c *touristServiceClient) SummaryByTime(ctx context.Context, in *TouristSummaryRequest, opts ...grpc.CallOption) (*TouristSummaryResponse, error) {
	out := new(TouristSummaryResponse)
	err := c.cc.Invoke(ctx, TouristService_SummaryByTime_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *touristServiceClient) SummaryByDate(ctx context.Context, in *TouristSummaryRequest, opts ...grpc.CallOption) (*TouristSummaryResponse, error) {
	out := new(TouristSummaryResponse)
	err := c.cc.Invoke(ctx, TouristService_SummaryByDate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *touristServiceClient) ForecastPassengerFlowList(ctx context.Context, in *grpc1.PassedTimeSpanByOrderTypeV4Request, opts ...grpc.CallOption) (*ForecastPassengerFlowListResponse, error) {
	out := new(ForecastPassengerFlowListResponse)
	err := c.cc.Invoke(ctx, TouristService_ForecastPassengerFlowList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *touristServiceClient) TouristLocal(ctx context.Context, in *TouristLocalRequest, opts ...grpc.CallOption) (*TouristLocalResponse, error) {
	out := new(TouristLocalResponse)
	err := c.cc.Invoke(ctx, TouristService_TouristLocal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *touristServiceClient) PassengerFlow(ctx context.Context, in *TouristStartEndRequest, opts ...grpc.CallOption) (*PassengerFlowResponse, error) {
	out := new(PassengerFlowResponse)
	err := c.cc.Invoke(ctx, TouristService_PassengerFlow_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *touristServiceClient) GroupList(ctx context.Context, in *GroupListRequest, opts ...grpc.CallOption) (*GroupListResponse, error) {
	out := new(GroupListResponse)
	err := c.cc.Invoke(ctx, TouristService_GroupList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *touristServiceClient) InOutStat(ctx context.Context, in *InOutStatRequest, opts ...grpc.CallOption) (*InOutStatResponse, error) {
	out := new(InOutStatResponse)
	err := c.cc.Invoke(ctx, TouristService_InOutStat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *touristServiceClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error) {
	out := new(CreateGroupResponse)
	err := c.cc.Invoke(ctx, TouristService_CreateGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *touristServiceClient) UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*grpc1.Response, error) {
	out := new(grpc1.Response)
	err := c.cc.Invoke(ctx, TouristService_UpdateGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *touristServiceClient) GroupById(ctx context.Context, in *GroupByIdRequest, opts ...grpc.CallOption) (*GroupListData, error) {
	out := new(GroupListData)
	err := c.cc.Invoke(ctx, TouristService_GroupById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *touristServiceClient) InoutByGroups(ctx context.Context, in *InoutByGroupsRequest, opts ...grpc.CallOption) (*InoutByGroupsResponse, error) {
	out := new(InoutByGroupsResponse)
	err := c.cc.Invoke(ctx, TouristService_InoutByGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *touristServiceClient) DeviceJournalStat(ctx context.Context, in *DeviceJournalStatRequest, opts ...grpc.CallOption) (*DeviceJournalStatResponse, error) {
	out := new(DeviceJournalStatResponse)
	err := c.cc.Invoke(ctx, TouristService_DeviceJournalStat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *touristServiceClient) InOutCompareFlow(ctx context.Context, in *InOutCompareRequest, opts ...grpc.CallOption) (*InoutCompareResponse, error) {
	out := new(InoutCompareResponse)
	err := c.cc.Invoke(ctx, TouristService_InOutCompareFlow_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TouristServiceServer is the server API for TouristService service.
// All implementations must embed UnimplementedTouristServiceServer
// for forward compatibility
type TouristServiceServer interface {
	// 获取按时间分组的出入园数据
	SummaryByTime(context.Context, *TouristSummaryRequest) (*TouristSummaryResponse, error)
	// 按天获取出入园数据
	SummaryByDate(context.Context, *TouristSummaryRequest) (*TouristSummaryResponse, error)
	// 客流预测每日以及今，明，后天数据
	ForecastPassengerFlowList(context.Context, *grpc1.PassedTimeSpanByOrderTypeV4Request) (*ForecastPassengerFlowListResponse, error)
	// 供应商维度客流来源TopN
	TouristLocal(context.Context, *TouristLocalRequest) (*TouristLocalResponse, error)
	// 每日客流数据(供应商维度)
	PassengerFlow(context.Context, *TouristStartEndRequest) (*PassengerFlowResponse, error)
	// 获取统计组
	GroupList(context.Context, *GroupListRequest) (*GroupListResponse, error)
	// 实时客流统计
	InOutStat(context.Context, *InOutStatRequest) (*InOutStatResponse, error)
	// 创建出入园统计组
	CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error)
	// 更新出入园统计组
	UpdateGroup(context.Context, *UpdateGroupRequest) (*grpc1.Response, error)
	// 根据id查询统计组信息
	GroupById(context.Context, *GroupByIdRequest) (*GroupListData, error)
	// 批量统计组获取出入园数据
	InoutByGroups(context.Context, *InoutByGroupsRequest) (*InoutByGroupsResponse, error)
	// 终端管理出入、验证数据查询
	DeviceJournalStat(context.Context, *DeviceJournalStatRequest) (*DeviceJournalStatResponse, error)
	// 实时对比客流数据
	InOutCompareFlow(context.Context, *InOutCompareRequest) (*InoutCompareResponse, error)
	mustEmbedUnimplementedTouristServiceServer()
}

// UnimplementedTouristServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTouristServiceServer struct {
}

func (UnimplementedTouristServiceServer) SummaryByTime(context.Context, *TouristSummaryRequest) (*TouristSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SummaryByTime not implemented")
}
func (UnimplementedTouristServiceServer) SummaryByDate(context.Context, *TouristSummaryRequest) (*TouristSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SummaryByDate not implemented")
}
func (UnimplementedTouristServiceServer) ForecastPassengerFlowList(context.Context, *grpc1.PassedTimeSpanByOrderTypeV4Request) (*ForecastPassengerFlowListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForecastPassengerFlowList not implemented")
}
func (UnimplementedTouristServiceServer) TouristLocal(context.Context, *TouristLocalRequest) (*TouristLocalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TouristLocal not implemented")
}
func (UnimplementedTouristServiceServer) PassengerFlow(context.Context, *TouristStartEndRequest) (*PassengerFlowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PassengerFlow not implemented")
}
func (UnimplementedTouristServiceServer) GroupList(context.Context, *GroupListRequest) (*GroupListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupList not implemented")
}
func (UnimplementedTouristServiceServer) InOutStat(context.Context, *InOutStatRequest) (*InOutStatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InOutStat not implemented")
}
func (UnimplementedTouristServiceServer) CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedTouristServiceServer) UpdateGroup(context.Context, *UpdateGroupRequest) (*grpc1.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroup not implemented")
}
func (UnimplementedTouristServiceServer) GroupById(context.Context, *GroupByIdRequest) (*GroupListData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupById not implemented")
}
func (UnimplementedTouristServiceServer) InoutByGroups(context.Context, *InoutByGroupsRequest) (*InoutByGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InoutByGroups not implemented")
}
func (UnimplementedTouristServiceServer) DeviceJournalStat(context.Context, *DeviceJournalStatRequest) (*DeviceJournalStatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeviceJournalStat not implemented")
}
func (UnimplementedTouristServiceServer) InOutCompareFlow(context.Context, *InOutCompareRequest) (*InoutCompareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InOutCompareFlow not implemented")
}
func (UnimplementedTouristServiceServer) mustEmbedUnimplementedTouristServiceServer() {}

// UnsafeTouristServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TouristServiceServer will
// result in compilation errors.
type UnsafeTouristServiceServer interface {
	mustEmbedUnimplementedTouristServiceServer()
}

func RegisterTouristServiceServer(s grpc.ServiceRegistrar, srv TouristServiceServer) {
	s.RegisterService(&TouristService_ServiceDesc, srv)
}

func _TouristService_SummaryByTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TouristSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TouristServiceServer).SummaryByTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TouristService_SummaryByTime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TouristServiceServer).SummaryByTime(ctx, req.(*TouristSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TouristService_SummaryByDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TouristSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TouristServiceServer).SummaryByDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TouristService_SummaryByDate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TouristServiceServer).SummaryByDate(ctx, req.(*TouristSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TouristService_ForecastPassengerFlowList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(grpc1.PassedTimeSpanByOrderTypeV4Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TouristServiceServer).ForecastPassengerFlowList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TouristService_ForecastPassengerFlowList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TouristServiceServer).ForecastPassengerFlowList(ctx, req.(*grpc1.PassedTimeSpanByOrderTypeV4Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _TouristService_TouristLocal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TouristLocalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TouristServiceServer).TouristLocal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TouristService_TouristLocal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TouristServiceServer).TouristLocal(ctx, req.(*TouristLocalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TouristService_PassengerFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TouristStartEndRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TouristServiceServer).PassengerFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TouristService_PassengerFlow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TouristServiceServer).PassengerFlow(ctx, req.(*TouristStartEndRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TouristService_GroupList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TouristServiceServer).GroupList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TouristService_GroupList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TouristServiceServer).GroupList(ctx, req.(*GroupListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TouristService_InOutStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InOutStatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TouristServiceServer).InOutStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TouristService_InOutStat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TouristServiceServer).InOutStat(ctx, req.(*InOutStatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TouristService_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TouristServiceServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TouristService_CreateGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TouristServiceServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TouristService_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TouristServiceServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TouristService_UpdateGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TouristServiceServer).UpdateGroup(ctx, req.(*UpdateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TouristService_GroupById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TouristServiceServer).GroupById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TouristService_GroupById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TouristServiceServer).GroupById(ctx, req.(*GroupByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TouristService_InoutByGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InoutByGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TouristServiceServer).InoutByGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TouristService_InoutByGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TouristServiceServer).InoutByGroups(ctx, req.(*InoutByGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TouristService_DeviceJournalStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceJournalStatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TouristServiceServer).DeviceJournalStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TouristService_DeviceJournalStat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TouristServiceServer).DeviceJournalStat(ctx, req.(*DeviceJournalStatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TouristService_InOutCompareFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InOutCompareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TouristServiceServer).InOutCompareFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TouristService_InOutCompareFlow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TouristServiceServer).InOutCompareFlow(ctx, req.(*InOutCompareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TouristService_ServiceDesc is the grpc.ServiceDesc for TouristService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TouristService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TouristService",
	HandlerType: (*TouristServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SummaryByTime",
			Handler:    _TouristService_SummaryByTime_Handler,
		},
		{
			MethodName: "SummaryByDate",
			Handler:    _TouristService_SummaryByDate_Handler,
		},
		{
			MethodName: "ForecastPassengerFlowList",
			Handler:    _TouristService_ForecastPassengerFlowList_Handler,
		},
		{
			MethodName: "TouristLocal",
			Handler:    _TouristService_TouristLocal_Handler,
		},
		{
			MethodName: "PassengerFlow",
			Handler:    _TouristService_PassengerFlow_Handler,
		},
		{
			MethodName: "GroupList",
			Handler:    _TouristService_GroupList_Handler,
		},
		{
			MethodName: "InOutStat",
			Handler:    _TouristService_InOutStat_Handler,
		},
		{
			MethodName: "CreateGroup",
			Handler:    _TouristService_CreateGroup_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _TouristService_UpdateGroup_Handler,
		},
		{
			MethodName: "GroupById",
			Handler:    _TouristService_GroupById_Handler,
		},
		{
			MethodName: "InoutByGroups",
			Handler:    _TouristService_InoutByGroups_Handler,
		},
		{
			MethodName: "DeviceJournalStat",
			Handler:    _TouristService_DeviceJournalStat_Handler,
		},
		{
			MethodName: "InOutCompareFlow",
			Handler:    _TouristService_InOutCompareFlow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmd/odas/grpc/tourist/tourist.proto",
}
