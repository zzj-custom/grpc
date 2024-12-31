// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.20.3
// source: cmd/odas/grpc/inout_authorization/inout_authorization.proto

package inout_authorization

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 定义创建请求参数和响应
type InoutAuthorizationCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Granter int64 `protobuf:"varint,1,opt,name=granter,proto3" json:"granter,omitempty"`
	Grantee int64 `protobuf:"varint,2,opt,name=grantee,proto3" json:"grantee,omitempty"`
	GroupId int64 `protobuf:"varint,3,opt,name=groupId,proto3" json:"groupId,omitempty"`
}

func (x *InoutAuthorizationCreateRequest) Reset() {
	*x = InoutAuthorizationCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InoutAuthorizationCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InoutAuthorizationCreateRequest) ProtoMessage() {}

func (x *InoutAuthorizationCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InoutAuthorizationCreateRequest.ProtoReflect.Descriptor instead.
func (*InoutAuthorizationCreateRequest) Descriptor() ([]byte, []int) {
	return file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_rawDescGZIP(), []int{0}
}

func (x *InoutAuthorizationCreateRequest) GetGranter() int64 {
	if x != nil {
		return x.Granter
	}
	return 0
}

func (x *InoutAuthorizationCreateRequest) GetGrantee() int64 {
	if x != nil {
		return x.Grantee
	}
	return 0
}

func (x *InoutAuthorizationCreateRequest) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type InoutAuthorizationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *InoutAuthorizationResponse) Reset() {
	*x = InoutAuthorizationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InoutAuthorizationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InoutAuthorizationResponse) ProtoMessage() {}

func (x *InoutAuthorizationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InoutAuthorizationResponse.ProtoReflect.Descriptor instead.
func (*InoutAuthorizationResponse) Descriptor() ([]byte, []int) {
	return file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_rawDescGZIP(), []int{1}
}

func (x *InoutAuthorizationResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *InoutAuthorizationResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// 定义编辑请求参数和响应
type InoutAuthorizationUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	GroupId int64 `protobuf:"varint,4,opt,name=groupId,proto3" json:"groupId,omitempty"`
}

func (x *InoutAuthorizationUpdateRequest) Reset() {
	*x = InoutAuthorizationUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InoutAuthorizationUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InoutAuthorizationUpdateRequest) ProtoMessage() {}

func (x *InoutAuthorizationUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InoutAuthorizationUpdateRequest.ProtoReflect.Descriptor instead.
func (*InoutAuthorizationUpdateRequest) Descriptor() ([]byte, []int) {
	return file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_rawDescGZIP(), []int{2}
}

func (x *InoutAuthorizationUpdateRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *InoutAuthorizationUpdateRequest) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

// 定义删除请求参数和响应
type InoutAuthorizationDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *InoutAuthorizationDeleteRequest) Reset() {
	*x = InoutAuthorizationDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InoutAuthorizationDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InoutAuthorizationDeleteRequest) ProtoMessage() {}

func (x *InoutAuthorizationDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InoutAuthorizationDeleteRequest.ProtoReflect.Descriptor instead.
func (*InoutAuthorizationDeleteRequest) Descriptor() ([]byte, []int) {
	return file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_rawDescGZIP(), []int{3}
}

func (x *InoutAuthorizationDeleteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 定义查询请求参数和响应
type InoutAuthorizationListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Granter  int64 `protobuf:"varint,3,opt,name=granter,proto3" json:"granter,omitempty"`
	Grantee  int64 `protobuf:"varint,4,opt,name=grantee,proto3" json:"grantee,omitempty"`
	GroupId  int64 `protobuf:"varint,5,opt,name=groupId,proto3" json:"groupId,omitempty"`
}

func (x *InoutAuthorizationListRequest) Reset() {
	*x = InoutAuthorizationListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InoutAuthorizationListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InoutAuthorizationListRequest) ProtoMessage() {}

func (x *InoutAuthorizationListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InoutAuthorizationListRequest.ProtoReflect.Descriptor instead.
func (*InoutAuthorizationListRequest) Descriptor() ([]byte, []int) {
	return file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_rawDescGZIP(), []int{4}
}

func (x *InoutAuthorizationListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *InoutAuthorizationListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *InoutAuthorizationListRequest) GetGranter() int64 {
	if x != nil {
		return x.Granter
	}
	return 0
}

func (x *InoutAuthorizationListRequest) GetGrantee() int64 {
	if x != nil {
		return x.Grantee
	}
	return 0
}

func (x *InoutAuthorizationListRequest) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type ListData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Granter   int64  `protobuf:"varint,2,opt,name=granter,proto3" json:"granter,omitempty"`
	Grantee   int64  `protobuf:"varint,3,opt,name=grantee,proto3" json:"grantee,omitempty"`
	GroupId   int64  `protobuf:"varint,4,opt,name=groupId,proto3" json:"groupId,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt string `protobuf:"bytes,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *ListData) Reset() {
	*x = ListData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListData) ProtoMessage() {}

func (x *ListData) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListData.ProtoReflect.Descriptor instead.
func (*ListData) Descriptor() ([]byte, []int) {
	return file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_rawDescGZIP(), []int{5}
}

func (x *ListData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListData) GetGranter() int64 {
	if x != nil {
		return x.Granter
	}
	return 0
}

func (x *ListData) GetGrantee() int64 {
	if x != nil {
		return x.Grantee
	}
	return 0
}

func (x *ListData) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *ListData) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ListData) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type InoutAuthorizationPagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Total    int64 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	Pages    int64 `protobuf:"varint,4,opt,name=pages,proto3" json:"pages,omitempty"`
}

func (x *InoutAuthorizationPagination) Reset() {
	*x = InoutAuthorizationPagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InoutAuthorizationPagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InoutAuthorizationPagination) ProtoMessage() {}

func (x *InoutAuthorizationPagination) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InoutAuthorizationPagination.ProtoReflect.Descriptor instead.
func (*InoutAuthorizationPagination) Descriptor() ([]byte, []int) {
	return file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_rawDescGZIP(), []int{6}
}

func (x *InoutAuthorizationPagination) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *InoutAuthorizationPagination) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *InoutAuthorizationPagination) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *InoutAuthorizationPagination) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

type InoutAuthorizationListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List       []*ListData                   `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Pagination *InoutAuthorizationPagination `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *InoutAuthorizationListResponse) Reset() {
	*x = InoutAuthorizationListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InoutAuthorizationListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InoutAuthorizationListResponse) ProtoMessage() {}

func (x *InoutAuthorizationListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InoutAuthorizationListResponse.ProtoReflect.Descriptor instead.
func (*InoutAuthorizationListResponse) Descriptor() ([]byte, []int) {
	return file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_rawDescGZIP(), []int{7}
}

func (x *InoutAuthorizationListResponse) GetList() []*ListData {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *InoutAuthorizationListResponse) GetPagination() *InoutAuthorizationPagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_cmd_odas_grpc_inout_authorization_inout_authorization_proto protoreflect.FileDescriptor

var file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x63, 0x6d, 0x64, 0x2f, 0x6f, 0x64, 0x61, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x69, 0x6e, 0x6f, 0x75, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x69, 0x6e, 0x6f, 0x75, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a,
	0x1f, 0x49, 0x6e, 0x6f, 0x75, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72,
	0x61, 0x6e, 0x74, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x61,
	0x6e, 0x74, 0x65, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x42,
	0x0a, 0x1a, 0x49, 0x6e, 0x6f, 0x75, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x22, 0x4b, 0x0a, 0x1f, 0x49, 0x6e, 0x6f, 0x75, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22,
	0x31, 0x0a, 0x1f, 0x49, 0x6e, 0x6f, 0x75, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x9d, 0x01, 0x0a, 0x1d, 0x49, 0x6e, 0x6f, 0x75, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x22, 0xa4, 0x01, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x61,
	0x6e, 0x74, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x61, 0x6e,
	0x74, 0x65, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x7a, 0x0a, 0x1c, 0x49, 0x6e, 0x6f,
	0x75, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x73, 0x22, 0x7e, 0x0a, 0x1e, 0x49, 0x6e, 0x6f, 0x75, 0x74, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x49, 0x6e, 0x6f,
	0x75, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xc7, 0x02, 0x0a, 0x19, 0x49, 0x6e, 0x6f, 0x75, 0x74, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e,
	0x49, 0x6e, 0x6f, 0x75, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x49, 0x6e, 0x6f, 0x75, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x49, 0x6e, 0x6f, 0x75, 0x74,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x49, 0x6e, 0x6f,
	0x75, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x49, 0x6e, 0x6f, 0x75, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x49, 0x6e, 0x6f, 0x75, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x49,
	0x6e, 0x6f, 0x75, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x49,
	0x6e, 0x6f, 0x75, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x3c, 0x5a, 0x3a, 0x6f, 0x64, 0x61, 0x73, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x6f, 0x64, 0x61, 0x73,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x6e, 0x6f, 0x75, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x69, 0x6e, 0x6f, 0x75, 0x74, 0x5f,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_rawDescOnce sync.Once
	file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_rawDescData = file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_rawDesc
)

func file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_rawDescGZIP() []byte {
	file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_rawDescOnce.Do(func() {
		file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_rawDescData)
	})
	return file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_rawDescData
}

var file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_goTypes = []interface{}{
	(*InoutAuthorizationCreateRequest)(nil), // 0: InoutAuthorizationCreateRequest
	(*InoutAuthorizationResponse)(nil),      // 1: InoutAuthorizationResponse
	(*InoutAuthorizationUpdateRequest)(nil), // 2: InoutAuthorizationUpdateRequest
	(*InoutAuthorizationDeleteRequest)(nil), // 3: InoutAuthorizationDeleteRequest
	(*InoutAuthorizationListRequest)(nil),   // 4: InoutAuthorizationListRequest
	(*ListData)(nil),                        // 5: ListData
	(*InoutAuthorizationPagination)(nil),    // 6: InoutAuthorizationPagination
	(*InoutAuthorizationListResponse)(nil),  // 7: InoutAuthorizationListResponse
}
var file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_depIdxs = []int32{
	5, // 0: InoutAuthorizationListResponse.list:type_name -> ListData
	6, // 1: InoutAuthorizationListResponse.pagination:type_name -> InoutAuthorizationPagination
	0, // 2: InoutAuthorizationService.Create:input_type -> InoutAuthorizationCreateRequest
	2, // 3: InoutAuthorizationService.Update:input_type -> InoutAuthorizationUpdateRequest
	3, // 4: InoutAuthorizationService.Delete:input_type -> InoutAuthorizationDeleteRequest
	4, // 5: InoutAuthorizationService.List:input_type -> InoutAuthorizationListRequest
	1, // 6: InoutAuthorizationService.Create:output_type -> InoutAuthorizationResponse
	1, // 7: InoutAuthorizationService.Update:output_type -> InoutAuthorizationResponse
	1, // 8: InoutAuthorizationService.Delete:output_type -> InoutAuthorizationResponse
	7, // 9: InoutAuthorizationService.List:output_type -> InoutAuthorizationListResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_init() }
func file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_init() {
	if File_cmd_odas_grpc_inout_authorization_inout_authorization_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InoutAuthorizationCreateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InoutAuthorizationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InoutAuthorizationUpdateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InoutAuthorizationDeleteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InoutAuthorizationListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InoutAuthorizationPagination); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InoutAuthorizationListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_goTypes,
		DependencyIndexes: file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_depIdxs,
		MessageInfos:      file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_msgTypes,
	}.Build()
	File_cmd_odas_grpc_inout_authorization_inout_authorization_proto = out.File
	file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_rawDesc = nil
	file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_goTypes = nil
	file_cmd_odas_grpc_inout_authorization_inout_authorization_proto_depIdxs = nil
}