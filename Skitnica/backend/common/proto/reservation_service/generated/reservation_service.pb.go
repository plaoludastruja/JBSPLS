// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: reservation_service.proto

package reservation

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
		mi := &file_reservation_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_service_proto_msgTypes[0]
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
	return file_reservation_service_proto_rawDescGZIP(), []int{0}
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

	Reservation *Reservation `protobuf:"bytes,1,opt,name=reservation,proto3" json:"reservation,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_service_proto_msgTypes[1]
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
	return file_reservation_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetReservation() *Reservation {
	if x != nil {
		return x.Reservation
	}
	return nil
}

type GetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GetAllRequest) Reset() {
	*x = GetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRequest) ProtoMessage() {}

func (x *GetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_service_proto_msgTypes[2]
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
	return file_reservation_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type GetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reservations []*Reservation `protobuf:"bytes,1,rep,name=reservations,proto3" json:"reservations,omitempty"`
}

func (x *GetAllResponse) Reset() {
	*x = GetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllResponse) ProtoMessage() {}

func (x *GetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_service_proto_msgTypes[3]
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
	return file_reservation_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllResponse) GetReservations() []*Reservation {
	if x != nil {
		return x.Reservations
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
		mi := &file_reservation_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_service_proto_msgTypes[4]
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
	return file_reservation_service_proto_rawDescGZIP(), []int{4}
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
		mi := &file_reservation_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_service_proto_msgTypes[5]
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
	return file_reservation_service_proto_rawDescGZIP(), []int{5}
}

type CreateReservationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reservation *Reservation `protobuf:"bytes,1,opt,name=reservation,proto3" json:"reservation,omitempty"`
}

func (x *CreateReservationRequest) Reset() {
	*x = CreateReservationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReservationRequest) ProtoMessage() {}

func (x *CreateReservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReservationRequest.ProtoReflect.Descriptor instead.
func (*CreateReservationRequest) Descriptor() ([]byte, []int) {
	return file_reservation_service_proto_rawDescGZIP(), []int{6}
}

func (x *CreateReservationRequest) GetReservation() *Reservation {
	if x != nil {
		return x.Reservation
	}
	return nil
}

type CreateReservationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reservation *Reservation `protobuf:"bytes,1,opt,name=reservation,proto3" json:"reservation,omitempty"`
}

func (x *CreateReservationResponse) Reset() {
	*x = CreateReservationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReservationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReservationResponse) ProtoMessage() {}

func (x *CreateReservationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReservationResponse.ProtoReflect.Descriptor instead.
func (*CreateReservationResponse) Descriptor() ([]byte, []int) {
	return file_reservation_service_proto_rawDescGZIP(), []int{7}
}

func (x *CreateReservationResponse) GetReservation() *Reservation {
	if x != nil {
		return x.Reservation
	}
	return nil
}

type EditReservationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reservation *Reservation `protobuf:"bytes,1,opt,name=reservation,proto3" json:"reservation,omitempty"`
}

func (x *EditReservationRequest) Reset() {
	*x = EditReservationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditReservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditReservationRequest) ProtoMessage() {}

func (x *EditReservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditReservationRequest.ProtoReflect.Descriptor instead.
func (*EditReservationRequest) Descriptor() ([]byte, []int) {
	return file_reservation_service_proto_rawDescGZIP(), []int{8}
}

func (x *EditReservationRequest) GetReservation() *Reservation {
	if x != nil {
		return x.Reservation
	}
	return nil
}

type EditReservationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reservation *Reservation `protobuf:"bytes,1,opt,name=reservation,proto3" json:"reservation,omitempty"`
}

func (x *EditReservationResponse) Reset() {
	*x = EditReservationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditReservationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditReservationResponse) ProtoMessage() {}

func (x *EditReservationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditReservationResponse.ProtoReflect.Descriptor instead.
func (*EditReservationResponse) Descriptor() ([]byte, []int) {
	return file_reservation_service_proto_rawDescGZIP(), []int{9}
}

func (x *EditReservationResponse) GetReservation() *Reservation {
	if x != nil {
		return x.Reservation
	}
	return nil
}

type Reservation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AccomodationId string `protobuf:"bytes,2,opt,name=accomodationId,proto3" json:"accomodationId,omitempty"`
	Username       string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	StartDate      string `protobuf:"bytes,4,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate        string `protobuf:"bytes,5,opt,name=endDate,proto3" json:"endDate,omitempty"`
	GuestNumber    int32  `protobuf:"varint,6,opt,name=guestNumber,proto3" json:"guestNumber,omitempty"`
	Status         string `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Reservation) Reset() {
	*x = Reservation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reservation) ProtoMessage() {}

func (x *Reservation) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reservation.ProtoReflect.Descriptor instead.
func (*Reservation) Descriptor() ([]byte, []int) {
	return file_reservation_service_proto_rawDescGZIP(), []int{10}
}

func (x *Reservation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Reservation) GetAccomodationId() string {
	if x != nil {
		return x.AccomodationId
	}
	return ""
}

func (x *Reservation) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Reservation) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *Reservation) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *Reservation) GetGuestNumber() int32 {
	if x != nil {
		return x.GuestNumber
	}
	return 0
}

func (x *Reservation) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_reservation_service_proto protoreflect.FileDescriptor

var file_reservation_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x2b, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4e, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c,
	0x0a, 0x0c, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x1f, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x10, 0x0a,
	0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x56, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0b, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x57, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x54, 0x0a, 0x16, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0b, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x55, 0x0a, 0x17, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd3, 0x01,
	0x0a, 0x0b, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a,
	0x0e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x67, 0x75, 0x65,
	0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x67, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x32, 0xc3, 0x04, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x17, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0x66, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x6c, 0x6c, 0x2f, 0x7b, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x85, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e,
	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x0c, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x7f, 0x0a, 0x0f, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x0c, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x67, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x2a, 0x11, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6f, 0x6c, 0x75, 0x64, 0x61,
	0x73, 0x74, 0x72, 0x75, 0x6a, 0x61, 0x2f, 0x4a, 0x42, 0x53, 0x50, 0x4c, 0x53, 0x2f, 0x53, 0x6b,
	0x69, 0x74, 0x6e, 0x69, 0x63, 0x61, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reservation_service_proto_rawDescOnce sync.Once
	file_reservation_service_proto_rawDescData = file_reservation_service_proto_rawDesc
)

func file_reservation_service_proto_rawDescGZIP() []byte {
	file_reservation_service_proto_rawDescOnce.Do(func() {
		file_reservation_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_reservation_service_proto_rawDescData)
	})
	return file_reservation_service_proto_rawDescData
}

var file_reservation_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_reservation_service_proto_goTypes = []interface{}{
	(*GetRequest)(nil),                // 0: reservation.GetRequest
	(*GetResponse)(nil),               // 1: reservation.GetResponse
	(*GetAllRequest)(nil),             // 2: reservation.GetAllRequest
	(*GetAllResponse)(nil),            // 3: reservation.GetAllResponse
	(*DeleteRequest)(nil),             // 4: reservation.DeleteRequest
	(*DeleteResponse)(nil),            // 5: reservation.DeleteResponse
	(*CreateReservationRequest)(nil),  // 6: reservation.CreateReservationRequest
	(*CreateReservationResponse)(nil), // 7: reservation.CreateReservationResponse
	(*EditReservationRequest)(nil),    // 8: reservation.EditReservationRequest
	(*EditReservationResponse)(nil),   // 9: reservation.EditReservationResponse
	(*Reservation)(nil),               // 10: reservation.Reservation
}
var file_reservation_service_proto_depIdxs = []int32{
	10, // 0: reservation.GetResponse.reservation:type_name -> reservation.Reservation
	10, // 1: reservation.GetAllResponse.reservations:type_name -> reservation.Reservation
	10, // 2: reservation.CreateReservationRequest.reservation:type_name -> reservation.Reservation
	10, // 3: reservation.CreateReservationResponse.reservation:type_name -> reservation.Reservation
	10, // 4: reservation.EditReservationRequest.reservation:type_name -> reservation.Reservation
	10, // 5: reservation.EditReservationResponse.reservation:type_name -> reservation.Reservation
	0,  // 6: reservation.ReservationService.Get:input_type -> reservation.GetRequest
	2,  // 7: reservation.ReservationService.GetAll:input_type -> reservation.GetAllRequest
	6,  // 8: reservation.ReservationService.CreateReservation:input_type -> reservation.CreateReservationRequest
	8,  // 9: reservation.ReservationService.EditReservation:input_type -> reservation.EditReservationRequest
	4,  // 10: reservation.ReservationService.DeleteReservation:input_type -> reservation.DeleteRequest
	1,  // 11: reservation.ReservationService.Get:output_type -> reservation.GetResponse
	3,  // 12: reservation.ReservationService.GetAll:output_type -> reservation.GetAllResponse
	7,  // 13: reservation.ReservationService.CreateReservation:output_type -> reservation.CreateReservationResponse
	9,  // 14: reservation.ReservationService.EditReservation:output_type -> reservation.EditReservationResponse
	5,  // 15: reservation.ReservationService.DeleteReservation:output_type -> reservation.DeleteResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_reservation_service_proto_init() }
func file_reservation_service_proto_init() {
	if File_reservation_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reservation_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_reservation_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_reservation_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_reservation_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_reservation_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_reservation_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_reservation_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReservationRequest); i {
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
		file_reservation_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReservationResponse); i {
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
		file_reservation_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditReservationRequest); i {
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
		file_reservation_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditReservationResponse); i {
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
		file_reservation_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reservation); i {
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
			RawDescriptor: file_reservation_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reservation_service_proto_goTypes,
		DependencyIndexes: file_reservation_service_proto_depIdxs,
		MessageInfos:      file_reservation_service_proto_msgTypes,
	}.Build()
	File_reservation_service_proto = out.File
	file_reservation_service_proto_rawDesc = nil
	file_reservation_service_proto_goTypes = nil
	file_reservation_service_proto_depIdxs = nil
}
