// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.0--rc2
// source: hostMark_service.proto

package hostMark

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hostMark_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hostMark_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_hostMark_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostmark *HostMark `protobuf:"bytes,1,opt,name=hostmark,proto3" json:"hostmark,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hostMark_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hostMark_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_hostMark_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetHostmark() *HostMark {
	if x != nil {
		return x.Hostmark
	}
	return nil
}

type GetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllRequest) Reset() {
	*x = GetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hostMark_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRequest) ProtoMessage() {}

func (x *GetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hostMark_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRequest.ProtoReflect.Descriptor instead.
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return file_hostMark_service_proto_rawDescGZIP(), []int{2}
}

type GetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostMark []*HostMark `protobuf:"bytes,1,rep,name=hostMark,proto3" json:"hostMark,omitempty"`
}

func (x *GetAllResponse) Reset() {
	*x = GetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hostMark_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllResponse) ProtoMessage() {}

func (x *GetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hostMark_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllResponse.ProtoReflect.Descriptor instead.
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return file_hostMark_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllResponse) GetHostMark() []*HostMark {
	if x != nil {
		return x.HostMark
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hostMark_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hostMark_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_hostMark_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hostMark_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hostMark_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_hostMark_service_proto_rawDescGZIP(), []int{5}
}

type CreateHostMarkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostMark *HostMark `protobuf:"bytes,1,opt,name=hostMark,proto3" json:"hostMark,omitempty"`
}

func (x *CreateHostMarkRequest) Reset() {
	*x = CreateHostMarkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hostMark_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHostMarkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHostMarkRequest) ProtoMessage() {}

func (x *CreateHostMarkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hostMark_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHostMarkRequest.ProtoReflect.Descriptor instead.
func (*CreateHostMarkRequest) Descriptor() ([]byte, []int) {
	return file_hostMark_service_proto_rawDescGZIP(), []int{6}
}

func (x *CreateHostMarkRequest) GetHostMark() *HostMark {
	if x != nil {
		return x.HostMark
	}
	return nil
}

type CreateHostMarkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostMark *HostMark `protobuf:"bytes,1,opt,name=hostMark,proto3" json:"hostMark,omitempty"`
}

func (x *CreateHostMarkResponse) Reset() {
	*x = CreateHostMarkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hostMark_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHostMarkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHostMarkResponse) ProtoMessage() {}

func (x *CreateHostMarkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hostMark_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHostMarkResponse.ProtoReflect.Descriptor instead.
func (*CreateHostMarkResponse) Descriptor() ([]byte, []int) {
	return file_hostMark_service_proto_rawDescGZIP(), []int{7}
}

func (x *CreateHostMarkResponse) GetHostMark() *HostMark {
	if x != nil {
		return x.HostMark
	}
	return nil
}

type EditHostMarkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostMark *HostMark `protobuf:"bytes,1,opt,name=hostMark,proto3" json:"hostMark,omitempty"`
}

func (x *EditHostMarkRequest) Reset() {
	*x = EditHostMarkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hostMark_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditHostMarkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditHostMarkRequest) ProtoMessage() {}

func (x *EditHostMarkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hostMark_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditHostMarkRequest.ProtoReflect.Descriptor instead.
func (*EditHostMarkRequest) Descriptor() ([]byte, []int) {
	return file_hostMark_service_proto_rawDescGZIP(), []int{8}
}

func (x *EditHostMarkRequest) GetHostMark() *HostMark {
	if x != nil {
		return x.HostMark
	}
	return nil
}

type EditHostMarkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostMark *HostMark `protobuf:"bytes,1,opt,name=hostMark,proto3" json:"hostMark,omitempty"`
}

func (x *EditHostMarkResponse) Reset() {
	*x = EditHostMarkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hostMark_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditHostMarkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditHostMarkResponse) ProtoMessage() {}

func (x *EditHostMarkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hostMark_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditHostMarkResponse.ProtoReflect.Descriptor instead.
func (*EditHostMarkResponse) Descriptor() ([]byte, []int) {
	return file_hostMark_service_proto_rawDescGZIP(), []int{9}
}

func (x *EditHostMarkResponse) GetHostMark() *HostMark {
	if x != nil {
		return x.HostMark
	}
	return nil
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartDay   string `protobuf:"bytes,1,opt,name=startDay,proto3" json:"startDay,omitempty"`
	StartMonth string `protobuf:"bytes,2,opt,name=startMonth,proto3" json:"startMonth,omitempty"`
	StartYear  string `protobuf:"bytes,3,opt,name=startYear,proto3" json:"startYear,omitempty"`
	EndDay     string `protobuf:"bytes,4,opt,name=endDay,proto3" json:"endDay,omitempty"`
	EndMonth   string `protobuf:"bytes,5,opt,name=endMonth,proto3" json:"endMonth,omitempty"`
	EndYear    string `protobuf:"bytes,6,opt,name=endYear,proto3" json:"endYear,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hostMark_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hostMark_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_hostMark_service_proto_rawDescGZIP(), []int{10}
}

func (x *SearchRequest) GetStartDay() string {
	if x != nil {
		return x.StartDay
	}
	return ""
}

func (x *SearchRequest) GetStartMonth() string {
	if x != nil {
		return x.StartMonth
	}
	return ""
}

func (x *SearchRequest) GetStartYear() string {
	if x != nil {
		return x.StartYear
	}
	return ""
}

func (x *SearchRequest) GetEndDay() string {
	if x != nil {
		return x.EndDay
	}
	return ""
}

func (x *SearchRequest) GetEndMonth() string {
	if x != nil {
		return x.EndMonth
	}
	return ""
}

func (x *SearchRequest) GetEndYear() string {
	if x != nil {
		return x.EndYear
	}
	return ""
}

type HostMark struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username     string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Grade        int32  `protobuf:"varint,3,opt,name=grade,proto3" json:"grade,omitempty"`
	HostUsername string `protobuf:"bytes,4,opt,name=hostUsername,proto3" json:"hostUsername,omitempty"`
}

func (x *HostMark) Reset() {
	*x = HostMark{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hostMark_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostMark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostMark) ProtoMessage() {}

func (x *HostMark) ProtoReflect() protoreflect.Message {
	mi := &file_hostMark_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostMark.ProtoReflect.Descriptor instead.
func (*HostMark) Descriptor() ([]byte, []int) {
	return file_hostMark_service_proto_rawDescGZIP(), []int{11}
}

func (x *HostMark) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *HostMark) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *HostMark) GetGrade() int32 {
	if x != nil {
		return x.Grade
	}
	return 0
}

func (x *HostMark) GetHostUsername() string {
	if x != nil {
		return x.HostUsername
	}
	return ""
}

var File_hostMark_service_proto protoreflect.FileDescriptor

var file_hostMark_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x68, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x4d, 0x61,
	0x72, 0x6b, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3d,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a,
	0x08, 0x68, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x4d,
	0x61, 0x72, 0x6b, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x0f, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x40,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2e, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x2e, 0x48, 0x6f,
	0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b,
	0x22, 0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x47, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73,
	0x74, 0x4d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x08,
	0x68, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x68, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x61,
	0x72, 0x6b, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x22, 0x48, 0x0a, 0x16,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x4d, 0x61,
	0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x4d,
	0x61, 0x72, 0x6b, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x52, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x22, 0x45, 0x0a, 0x13, 0x45, 0x64, 0x69, 0x74, 0x48, 0x6f,
	0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a,
	0x08, 0x68, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x4d,
	0x61, 0x72, 0x6b, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x22, 0x46, 0x0a,
	0x14, 0x45, 0x64, 0x69, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x4d, 0x61,
	0x72, 0x6b, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x52, 0x08, 0x68, 0x6f, 0x73,
	0x74, 0x4d, 0x61, 0x72, 0x6b, 0x22, 0xb7, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x44, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x44, 0x61, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x6f, 0x6e, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x6f,
	0x6e, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x59, 0x65, 0x61, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x59, 0x65, 0x61,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64,
	0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x59, 0x65, 0x61, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x59, 0x65, 0x61, 0x72, 0x22,
	0x70, 0x0a, 0x08, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x68, 0x6f, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x6f, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x32, 0xec, 0x03, 0x0a, 0x0f, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x68,
	0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x10, 0x12, 0x0e, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x12, 0x52, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x17, 0x2e, 0x68, 0x6f,
	0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x72,
	0x6b, 0x2f, 0x61, 0x6c, 0x6c, 0x12, 0x70, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48,
	0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x12, 0x1f, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x4d, 0x61,
	0x72, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x4d,
	0x61, 0x72, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x61,
	0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x3a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x22, 0x09, 0x2f, 0x68,
	0x6f, 0x73, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x6a, 0x0a, 0x0c, 0x45, 0x64, 0x69, 0x74, 0x48,
	0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x12, 0x1d, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x4d, 0x61,
	0x72, 0x6b, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72,
	0x6b, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x08,
	0x68, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x1a, 0x09, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x6d,
	0x61, 0x72, 0x6b, 0x12, 0x5b, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x73,
	0x74, 0x4d, 0x61, 0x72, 0x6b, 0x12, 0x17, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x68, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10,
	0x2a, 0x0e, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x6c, 0x61, 0x6f, 0x6c, 0x75, 0x64, 0x61, 0x73, 0x74, 0x72, 0x75, 0x6a, 0x61, 0x2f, 0x4a, 0x42,
	0x53, 0x50, 0x4c, 0x53, 0x2f, 0x53, 0x6b, 0x69, 0x74, 0x6e, 0x69, 0x63, 0x61, 0x2f, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x68, 0x6f, 0x73,
	0x74, 0x4d, 0x61, 0x72, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hostMark_service_proto_rawDescOnce sync.Once
	file_hostMark_service_proto_rawDescData = file_hostMark_service_proto_rawDesc
)

func file_hostMark_service_proto_rawDescGZIP() []byte {
	file_hostMark_service_proto_rawDescOnce.Do(func() {
		file_hostMark_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_hostMark_service_proto_rawDescData)
	})
	return file_hostMark_service_proto_rawDescData
}

var file_hostMark_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_hostMark_service_proto_goTypes = []interface{}{
	(*GetRequest)(nil),             // 0: hostMark.GetRequest
	(*GetResponse)(nil),            // 1: hostMark.GetResponse
	(*GetAllRequest)(nil),          // 2: hostMark.GetAllRequest
	(*GetAllResponse)(nil),         // 3: hostMark.GetAllResponse
	(*DeleteRequest)(nil),          // 4: hostMark.DeleteRequest
	(*DeleteResponse)(nil),         // 5: hostMark.DeleteResponse
	(*CreateHostMarkRequest)(nil),  // 6: hostMark.CreateHostMarkRequest
	(*CreateHostMarkResponse)(nil), // 7: hostMark.CreateHostMarkResponse
	(*EditHostMarkRequest)(nil),    // 8: hostMark.EditHostMarkRequest
	(*EditHostMarkResponse)(nil),   // 9: hostMark.EditHostMarkResponse
	(*SearchRequest)(nil),          // 10: hostMark.SearchRequest
	(*HostMark)(nil),               // 11: hostMark.HostMark
}
var file_hostMark_service_proto_depIdxs = []int32{
	11, // 0: hostMark.GetResponse.hostmark:type_name -> hostMark.HostMark
	11, // 1: hostMark.GetAllResponse.hostMark:type_name -> hostMark.HostMark
	11, // 2: hostMark.CreateHostMarkRequest.hostMark:type_name -> hostMark.HostMark
	11, // 3: hostMark.CreateHostMarkResponse.hostMark:type_name -> hostMark.HostMark
	11, // 4: hostMark.EditHostMarkRequest.hostMark:type_name -> hostMark.HostMark
	11, // 5: hostMark.EditHostMarkResponse.hostMark:type_name -> hostMark.HostMark
	0,  // 6: hostMark.HostMarkService.Get:input_type -> hostMark.GetRequest
	2,  // 7: hostMark.HostMarkService.GetAll:input_type -> hostMark.GetAllRequest
	6,  // 8: hostMark.HostMarkService.CreateHostMark:input_type -> hostMark.CreateHostMarkRequest
	8,  // 9: hostMark.HostMarkService.EditHostMark:input_type -> hostMark.EditHostMarkRequest
	4,  // 10: hostMark.HostMarkService.DeleteHostMark:input_type -> hostMark.DeleteRequest
	1,  // 11: hostMark.HostMarkService.Get:output_type -> hostMark.GetResponse
	3,  // 12: hostMark.HostMarkService.GetAll:output_type -> hostMark.GetAllResponse
	7,  // 13: hostMark.HostMarkService.CreateHostMark:output_type -> hostMark.CreateHostMarkResponse
	9,  // 14: hostMark.HostMarkService.EditHostMark:output_type -> hostMark.EditHostMarkResponse
	5,  // 15: hostMark.HostMarkService.DeleteHostMark:output_type -> hostMark.DeleteResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_hostMark_service_proto_init() }
func file_hostMark_service_proto_init() {
	if File_hostMark_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hostMark_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_hostMark_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_hostMark_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRequest); i {
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
		file_hostMark_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllResponse); i {
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
		file_hostMark_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_hostMark_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
		file_hostMark_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHostMarkRequest); i {
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
		file_hostMark_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHostMarkResponse); i {
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
		file_hostMark_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditHostMarkRequest); i {
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
		file_hostMark_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditHostMarkResponse); i {
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
		file_hostMark_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_hostMark_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostMark); i {
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
			RawDescriptor: file_hostMark_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hostMark_service_proto_goTypes,
		DependencyIndexes: file_hostMark_service_proto_depIdxs,
		MessageInfos:      file_hostMark_service_proto_msgTypes,
	}.Build()
	File_hostMark_service_proto = out.File
	file_hostMark_service_proto_rawDesc = nil
	file_hostMark_service_proto_goTypes = nil
	file_hostMark_service_proto_depIdxs = nil
}
