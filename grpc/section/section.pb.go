// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.20.3
// source: cmd/odas/grpc/section/section.proto

package section

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

// 定义 VerifiedSummaryRequest 消息
type TodaySectionDeductionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 字段
	Lid []int64 `protobuf:"varint,1,rep,packed,name=lid,proto3" json:"lid,omitempty"`
}

func (x *TodaySectionDeductionRequest) Reset() {
	*x = TodaySectionDeductionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_odas_grpc_section_section_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodaySectionDeductionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodaySectionDeductionRequest) ProtoMessage() {}

func (x *TodaySectionDeductionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_odas_grpc_section_section_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodaySectionDeductionRequest.ProtoReflect.Descriptor instead.
func (*TodaySectionDeductionRequest) Descriptor() ([]byte, []int) {
	return file_cmd_odas_grpc_section_section_proto_rawDescGZIP(), []int{0}
}

func (x *TodaySectionDeductionRequest) GetLid() []int64 {
	if x != nil {
		return x.Lid
	}
	return nil
}

// 定义 TodaySectionDeductionResponse 消息返回格式
type TodaySectionDeductionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*SectionDeductionList `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *TodaySectionDeductionResponse) Reset() {
	*x = TodaySectionDeductionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_odas_grpc_section_section_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodaySectionDeductionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodaySectionDeductionResponse) ProtoMessage() {}

func (x *TodaySectionDeductionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_odas_grpc_section_section_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodaySectionDeductionResponse.ProtoReflect.Descriptor instead.
func (*TodaySectionDeductionResponse) Descriptor() ([]byte, []int) {
	return file_cmd_odas_grpc_section_section_proto_rawDescGZIP(), []int{1}
}

func (x *TodaySectionDeductionResponse) GetList() []*SectionDeductionList {
	if x != nil {
		return x.List
	}
	return nil
}

type SectionDeductionList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BeginDate               string   `protobuf:"bytes,1,opt,name=beginDate,proto3" json:"beginDate,omitempty"`
	BeginTime               string   `protobuf:"bytes,2,opt,name=beginTime,proto3" json:"beginTime,omitempty"`
	EndDate                 string   `protobuf:"bytes,3,opt,name=endDate,proto3" json:"endDate,omitempty"`
	EndTime                 string   `protobuf:"bytes,4,opt,name=endTime,proto3" json:"endTime,omitempty"`
	ReservedStorage         int64    `protobuf:"varint,5,opt,name=reservedStorage,proto3" json:"reservedStorage,omitempty"`
	RuleId                  int64    `protobuf:"varint,6,opt,name=ruleId,proto3" json:"ruleId,omitempty"`
	RuleName                string   `protobuf:"bytes,7,opt,name=ruleName,proto3" json:"ruleName,omitempty"`
	SectionAheadCheckInTime int64    `protobuf:"varint,8,opt,name=sectionAheadCheckInTime,proto3" json:"sectionAheadCheckInTime,omitempty"`
	SectionDelayCheckInTime int64    `protobuf:"varint,9,opt,name=sectionDelayCheckInTime,proto3" json:"sectionDelayCheckInTime,omitempty"`
	SectionTimeId           int64    `protobuf:"varint,10,opt,name=sectionTimeId,proto3" json:"sectionTimeId,omitempty"`
	Shop                    string   `protobuf:"bytes,11,opt,name=shop,proto3" json:"shop,omitempty"`
	SoldStorage             int64    `protobuf:"varint,12,opt,name=soldStorage,proto3" json:"soldStorage,omitempty"`
	TicketGroup             int64    `protobuf:"varint,13,opt,name=ticketGroup,proto3" json:"ticketGroup,omitempty"`
	TicketIds               []int64  `protobuf:"varint,14,rep,packed,name=ticketIds,proto3" json:"ticketIds,omitempty"`
	TicketNames             []string `protobuf:"bytes,15,rep,name=ticketNames,proto3" json:"ticketNames,omitempty"`
	TotalStorage            int64    `protobuf:"varint,16,opt,name=totalStorage,proto3" json:"totalStorage,omitempty"`
}

func (x *SectionDeductionList) Reset() {
	*x = SectionDeductionList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_odas_grpc_section_section_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SectionDeductionList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SectionDeductionList) ProtoMessage() {}

func (x *SectionDeductionList) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_odas_grpc_section_section_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SectionDeductionList.ProtoReflect.Descriptor instead.
func (*SectionDeductionList) Descriptor() ([]byte, []int) {
	return file_cmd_odas_grpc_section_section_proto_rawDescGZIP(), []int{2}
}

func (x *SectionDeductionList) GetBeginDate() string {
	if x != nil {
		return x.BeginDate
	}
	return ""
}

func (x *SectionDeductionList) GetBeginTime() string {
	if x != nil {
		return x.BeginTime
	}
	return ""
}

func (x *SectionDeductionList) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *SectionDeductionList) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *SectionDeductionList) GetReservedStorage() int64 {
	if x != nil {
		return x.ReservedStorage
	}
	return 0
}

func (x *SectionDeductionList) GetRuleId() int64 {
	if x != nil {
		return x.RuleId
	}
	return 0
}

func (x *SectionDeductionList) GetRuleName() string {
	if x != nil {
		return x.RuleName
	}
	return ""
}

func (x *SectionDeductionList) GetSectionAheadCheckInTime() int64 {
	if x != nil {
		return x.SectionAheadCheckInTime
	}
	return 0
}

func (x *SectionDeductionList) GetSectionDelayCheckInTime() int64 {
	if x != nil {
		return x.SectionDelayCheckInTime
	}
	return 0
}

func (x *SectionDeductionList) GetSectionTimeId() int64 {
	if x != nil {
		return x.SectionTimeId
	}
	return 0
}

func (x *SectionDeductionList) GetShop() string {
	if x != nil {
		return x.Shop
	}
	return ""
}

func (x *SectionDeductionList) GetSoldStorage() int64 {
	if x != nil {
		return x.SoldStorage
	}
	return 0
}

func (x *SectionDeductionList) GetTicketGroup() int64 {
	if x != nil {
		return x.TicketGroup
	}
	return 0
}

func (x *SectionDeductionList) GetTicketIds() []int64 {
	if x != nil {
		return x.TicketIds
	}
	return nil
}

func (x *SectionDeductionList) GetTicketNames() []string {
	if x != nil {
		return x.TicketNames
	}
	return nil
}

func (x *SectionDeductionList) GetTotalStorage() int64 {
	if x != nil {
		return x.TotalStorage
	}
	return 0
}

var File_cmd_odas_grpc_section_section_proto protoreflect.FileDescriptor

var file_cmd_odas_grpc_section_section_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6d, 0x64, 0x2f, 0x6f, 0x64, 0x61, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x1c, 0x54, 0x6f, 0x64, 0x61, 0x79, 0x53, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x03, 0x6c, 0x69, 0x64, 0x22, 0x4a, 0x0a, 0x1d, 0x54, 0x6f, 0x64, 0x61, 0x79,
	0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x22, 0xba, 0x04, 0x0a, 0x14, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x62, 0x65, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x65,
	0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62,
	0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f,
	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x75, 0x6c, 0x65, 0x49, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x17, 0x73, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x68, 0x65, 0x61, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x73, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x68, 0x65, 0x61, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x17, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x65, 0x6c, 0x61, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65,
	0x6c, 0x61, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24,
	0x0a, 0x0d, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x6f, 0x6c, 0x64,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73,
	0x6f, 0x6c, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x32, 0x79, 0x0a, 0x17, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x64, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x1b, 0x54,
	0x6f, 0x64, 0x61, 0x79, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x64, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x2e, 0x54, 0x6f, 0x64,
	0x61, 0x79, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x54, 0x6f, 0x64, 0x61,
	0x79, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22, 0x6f,
	0x64, 0x61, 0x73, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x6f, 0x64, 0x61, 0x73, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_odas_grpc_section_section_proto_rawDescOnce sync.Once
	file_cmd_odas_grpc_section_section_proto_rawDescData = file_cmd_odas_grpc_section_section_proto_rawDesc
)

func file_cmd_odas_grpc_section_section_proto_rawDescGZIP() []byte {
	file_cmd_odas_grpc_section_section_proto_rawDescOnce.Do(func() {
		file_cmd_odas_grpc_section_section_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_odas_grpc_section_section_proto_rawDescData)
	})
	return file_cmd_odas_grpc_section_section_proto_rawDescData
}

var file_cmd_odas_grpc_section_section_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cmd_odas_grpc_section_section_proto_goTypes = []interface{}{
	(*TodaySectionDeductionRequest)(nil),  // 0: TodaySectionDeductionRequest
	(*TodaySectionDeductionResponse)(nil), // 1: TodaySectionDeductionResponse
	(*SectionDeductionList)(nil),          // 2: SectionDeductionList
}
var file_cmd_odas_grpc_section_section_proto_depIdxs = []int32{
	2, // 0: TodaySectionDeductionResponse.list:type_name -> SectionDeductionList
	0, // 1: SectionDeductionService.TodaySectionDeductionDetail:input_type -> TodaySectionDeductionRequest
	1, // 2: SectionDeductionService.TodaySectionDeductionDetail:output_type -> TodaySectionDeductionResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cmd_odas_grpc_section_section_proto_init() }
func file_cmd_odas_grpc_section_section_proto_init() {
	if File_cmd_odas_grpc_section_section_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_odas_grpc_section_section_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodaySectionDeductionRequest); i {
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
		file_cmd_odas_grpc_section_section_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodaySectionDeductionResponse); i {
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
		file_cmd_odas_grpc_section_section_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SectionDeductionList); i {
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
			RawDescriptor: file_cmd_odas_grpc_section_section_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cmd_odas_grpc_section_section_proto_goTypes,
		DependencyIndexes: file_cmd_odas_grpc_section_section_proto_depIdxs,
		MessageInfos:      file_cmd_odas_grpc_section_section_proto_msgTypes,
	}.Build()
	File_cmd_odas_grpc_section_section_proto = out.File
	file_cmd_odas_grpc_section_section_proto_rawDesc = nil
	file_cmd_odas_grpc_section_section_proto_goTypes = nil
	file_cmd_odas_grpc_section_section_proto_depIdxs = nil
}
