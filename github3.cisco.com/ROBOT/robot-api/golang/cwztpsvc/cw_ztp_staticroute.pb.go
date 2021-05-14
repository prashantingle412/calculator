// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.6.1
// source: cw_ztp_staticroute/cw_ztp_staticroute.proto

package cwztpsvc

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

//database schema
type StaticRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unique uuid for subnet mask
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// subnet
	Subnet string `protobuf:"bytes,2,opt,name=subnet,proto3" json:"subnet,omitempty"`
	// mask
	Mask string `protobuf:"bytes,3,opt,name=mask,proto3" json:"mask,omitempty"`
	// status of the subnet and mask
	Status string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	// operation e.g add or delete
	Operation string `protobuf:"bytes,5,opt,name=operation,proto3" json:"operation,omitempty"`
	Message   string `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StaticRoute) Reset() {
	*x = StaticRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaticRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticRoute) ProtoMessage() {}

func (x *StaticRoute) ProtoReflect() protoreflect.Message {
	mi := &file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticRoute.ProtoReflect.Descriptor instead.
func (*StaticRoute) Descriptor() ([]byte, []int) {
	return file_cw_ztp_staticroute_cw_ztp_staticroute_proto_rawDescGZIP(), []int{0}
}

func (x *StaticRoute) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *StaticRoute) GetSubnet() string {
	if x != nil {
		return x.Subnet
	}
	return ""
}

func (x *StaticRoute) GetMask() string {
	if x != nil {
		return x.Mask
	}
	return ""
}

func (x *StaticRoute) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *StaticRoute) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *StaticRoute) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

//single object
type AddStaticRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subnet string `protobuf:"bytes,1,opt,name=subnet,proto3" json:"subnet,omitempty"`
	Mask   string `protobuf:"bytes,2,opt,name=mask,proto3" json:"mask,omitempty"`
}

func (x *AddStaticRoute) Reset() {
	*x = AddStaticRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddStaticRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddStaticRoute) ProtoMessage() {}

func (x *AddStaticRoute) ProtoReflect() protoreflect.Message {
	mi := &file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddStaticRoute.ProtoReflect.Descriptor instead.
func (*AddStaticRoute) Descriptor() ([]byte, []int) {
	return file_cw_ztp_staticroute_cw_ztp_staticroute_proto_rawDescGZIP(), []int{1}
}

func (x *AddStaticRoute) GetSubnet() string {
	if x != nil {
		return x.Subnet
	}
	return ""
}

func (x *AddStaticRoute) GetMask() string {
	if x != nil {
		return x.Mask
	}
	return ""
}

//request body for add static route []AddStaticRoute
type AddStaticRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddStaticRoutList []*AddStaticRoute `protobuf:"bytes,1,rep,name=addStaticRoutList,proto3" json:"addStaticRoutList,omitempty"`
}

func (x *AddStaticRouteRequest) Reset() {
	*x = AddStaticRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddStaticRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddStaticRouteRequest) ProtoMessage() {}

func (x *AddStaticRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddStaticRouteRequest.ProtoReflect.Descriptor instead.
func (*AddStaticRouteRequest) Descriptor() ([]byte, []int) {
	return file_cw_ztp_staticroute_cw_ztp_staticroute_proto_rawDescGZIP(), []int{2}
}

func (x *AddStaticRouteRequest) GetAddStaticRoutList() []*AddStaticRoute {
	if x != nil {
		return x.AddStaticRoutList
	}
	return nil
}

// response for AddStaticRoute
type AddStaticRouteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AddStaticRouteResponse) Reset() {
	*x = AddStaticRouteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddStaticRouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddStaticRouteResponse) ProtoMessage() {}

func (x *AddStaticRouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddStaticRouteResponse.ProtoReflect.Descriptor instead.
func (*AddStaticRouteResponse) Descriptor() ([]byte, []int) {
	return file_cw_ztp_staticroute_cw_ztp_staticroute_proto_rawDescGZIP(), []int{3}
}

func (x *AddStaticRouteResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AddStaticRouteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Fetch query request
type StaticRouteGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *Filter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *StaticRouteGetRequest) Reset() {
	*x = StaticRouteGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaticRouteGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticRouteGetRequest) ProtoMessage() {}

func (x *StaticRouteGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticRouteGetRequest.ProtoReflect.Descriptor instead.
func (*StaticRouteGetRequest) Descriptor() ([]byte, []int) {
	return file_cw_ztp_staticroute_cw_ztp_staticroute_proto_rawDescGZIP(), []int{4}
}

func (x *StaticRouteGetRequest) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize int32 `protobuf:"varint,1,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_cw_ztp_staticroute_cw_ztp_staticroute_proto_rawDescGZIP(), []int{5}
}

func (x *Filter) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// Fetch query response
type StaticRouteGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaticRouteEntityList []*StaticRoute `protobuf:"bytes,1,rep,name=StaticRouteEntityList,proto3" json:"StaticRouteEntityList,omitempty"`
	Status                string         `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *StaticRouteGetResponse) Reset() {
	*x = StaticRouteGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaticRouteGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticRouteGetResponse) ProtoMessage() {}

func (x *StaticRouteGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticRouteGetResponse.ProtoReflect.Descriptor instead.
func (*StaticRouteGetResponse) Descriptor() ([]byte, []int) {
	return file_cw_ztp_staticroute_cw_ztp_staticroute_proto_rawDescGZIP(), []int{6}
}

func (x *StaticRouteGetResponse) GetStaticRouteEntityList() []*StaticRoute {
	if x != nil {
		return x.StaticRouteEntityList
	}
	return nil
}

func (x *StaticRouteGetResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// for delete
type DeleteStaticRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *DeleteStaticRoute) Reset() {
	*x = DeleteStaticRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStaticRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStaticRoute) ProtoMessage() {}

func (x *DeleteStaticRoute) ProtoReflect() protoreflect.Message {
	mi := &file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStaticRoute.ProtoReflect.Descriptor instead.
func (*DeleteStaticRoute) Descriptor() ([]byte, []int) {
	return file_cw_ztp_staticroute_cw_ztp_staticroute_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteStaticRoute) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

// for delete request
type DeleteStaticRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeleteStaticRouteList []*DeleteStaticRoute `protobuf:"bytes,1,rep,name=deleteStaticRouteList,proto3" json:"deleteStaticRouteList,omitempty"`
}

func (x *DeleteStaticRouteRequest) Reset() {
	*x = DeleteStaticRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStaticRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStaticRouteRequest) ProtoMessage() {}

func (x *DeleteStaticRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStaticRouteRequest.ProtoReflect.Descriptor instead.
func (*DeleteStaticRouteRequest) Descriptor() ([]byte, []int) {
	return file_cw_ztp_staticroute_cw_ztp_staticroute_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteStaticRouteRequest) GetDeleteStaticRouteList() []*DeleteStaticRoute {
	if x != nil {
		return x.DeleteStaticRouteList
	}
	return nil
}

// for delete response
type DeleteStaticRouteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteStaticRouteResponse) Reset() {
	*x = DeleteStaticRouteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStaticRouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStaticRouteResponse) ProtoMessage() {}

func (x *DeleteStaticRouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStaticRouteResponse.ProtoReflect.Descriptor instead.
func (*DeleteStaticRouteResponse) Descriptor() ([]byte, []int) {
	return file_cw_ztp_staticroute_cw_ztp_staticroute_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteStaticRouteResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DeleteStaticRouteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_cw_ztp_staticroute_cw_ztp_staticroute_proto protoreflect.FileDescriptor

var file_cw_ztp_staticroute_cw_ztp_staticroute_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x63, 0x77, 0x5f, 0x7a, 0x74, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x2f, 0x63, 0x77, 0x5f, 0x7a, 0x74, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63,
	0x77, 0x7a, 0x74, 0x70, 0x73, 0x76, 0x63, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x62,
	0x6e, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6d, 0x61, 0x73, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3c, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d,
	0x61, 0x73, 0x6b, 0x22, 0x5f, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x11,
	0x61, 0x64, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x6f, 0x75, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x77, 0x7a, 0x74, 0x70, 0x73,
	0x76, 0x63, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x52, 0x11, 0x61, 0x64, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x6f, 0x75, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x46, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x41, 0x0a, 0x15,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x77, 0x7a, 0x74, 0x70, 0x73, 0x76, 0x63,
	0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22,
	0x24, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x7d, 0x0a, 0x16, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4b, 0x0a, 0x15, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x63, 0x77, 0x7a, 0x74, 0x70, 0x73, 0x76, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x15, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x27, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x6d, 0x0a,
	0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x51, 0x0a, 0x15, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x77, 0x7a, 0x74, 0x70,
	0x73, 0x76, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x15, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x49, 0x0a, 0x19,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xf6, 0x02, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6f,
	0x0a, 0x0e, 0x41, 0x64, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x12, 0x1f, 0x2e, 0x63, 0x77, 0x7a, 0x74, 0x70, 0x73, 0x76, 0x63, 0x2e, 0x41, 0x64, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x63, 0x77, 0x7a, 0x74, 0x70, 0x73, 0x76, 0x63, 0x2e, 0x41, 0x64, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12,
	0x75, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x12, 0x1f, 0x2e, 0x63, 0x77, 0x7a, 0x74, 0x70, 0x73, 0x76, 0x63, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x77, 0x7a, 0x74, 0x70, 0x73, 0x76, 0x63, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2f, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x78, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x63, 0x77,
	0x7a, 0x74, 0x70, 0x73, 0x76, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x63, 0x77, 0x7a, 0x74, 0x70, 0x73, 0x76, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x2a, 0x0f, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x3a, 0x01, 0x2a,
	0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x33, 0x2e, 0x63, 0x69, 0x73, 0x63,
	0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x4f, 0x42, 0x4f, 0x54, 0x2f, 0x72, 0x6f, 0x62, 0x6f,
	0x74, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x63, 0x77, 0x7a,
	0x74, 0x70, 0x73, 0x76, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cw_ztp_staticroute_cw_ztp_staticroute_proto_rawDescOnce sync.Once
	file_cw_ztp_staticroute_cw_ztp_staticroute_proto_rawDescData = file_cw_ztp_staticroute_cw_ztp_staticroute_proto_rawDesc
)

func file_cw_ztp_staticroute_cw_ztp_staticroute_proto_rawDescGZIP() []byte {
	file_cw_ztp_staticroute_cw_ztp_staticroute_proto_rawDescOnce.Do(func() {
		file_cw_ztp_staticroute_cw_ztp_staticroute_proto_rawDescData = protoimpl.X.CompressGZIP(file_cw_ztp_staticroute_cw_ztp_staticroute_proto_rawDescData)
	})
	return file_cw_ztp_staticroute_cw_ztp_staticroute_proto_rawDescData
}

var file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_cw_ztp_staticroute_cw_ztp_staticroute_proto_goTypes = []interface{}{
	(*StaticRoute)(nil),               // 0: cwztpsvc.StaticRoute
	(*AddStaticRoute)(nil),            // 1: cwztpsvc.AddStaticRoute
	(*AddStaticRouteRequest)(nil),     // 2: cwztpsvc.AddStaticRouteRequest
	(*AddStaticRouteResponse)(nil),    // 3: cwztpsvc.AddStaticRouteResponse
	(*StaticRouteGetRequest)(nil),     // 4: cwztpsvc.StaticRouteGetRequest
	(*Filter)(nil),                    // 5: cwztpsvc.Filter
	(*StaticRouteGetResponse)(nil),    // 6: cwztpsvc.StaticRouteGetResponse
	(*DeleteStaticRoute)(nil),         // 7: cwztpsvc.DeleteStaticRoute
	(*DeleteStaticRouteRequest)(nil),  // 8: cwztpsvc.DeleteStaticRouteRequest
	(*DeleteStaticRouteResponse)(nil), // 9: cwztpsvc.DeleteStaticRouteResponse
}
var file_cw_ztp_staticroute_cw_ztp_staticroute_proto_depIdxs = []int32{
	1, // 0: cwztpsvc.AddStaticRouteRequest.addStaticRoutList:type_name -> cwztpsvc.AddStaticRoute
	5, // 1: cwztpsvc.StaticRouteGetRequest.filter:type_name -> cwztpsvc.Filter
	0, // 2: cwztpsvc.StaticRouteGetResponse.StaticRouteEntityList:type_name -> cwztpsvc.StaticRoute
	7, // 3: cwztpsvc.DeleteStaticRouteRequest.deleteStaticRouteList:type_name -> cwztpsvc.DeleteStaticRoute
	2, // 4: cwztpsvc.StaticRouteService.AddStaticRoute:input_type -> cwztpsvc.AddStaticRouteRequest
	4, // 5: cwztpsvc.StaticRouteService.GetStaticRoute:input_type -> cwztpsvc.StaticRouteGetRequest
	8, // 6: cwztpsvc.StaticRouteService.DeleteStaticRoute:input_type -> cwztpsvc.DeleteStaticRouteRequest
	3, // 7: cwztpsvc.StaticRouteService.AddStaticRoute:output_type -> cwztpsvc.AddStaticRouteResponse
	6, // 8: cwztpsvc.StaticRouteService.GetStaticRoute:output_type -> cwztpsvc.StaticRouteGetResponse
	9, // 9: cwztpsvc.StaticRouteService.DeleteStaticRoute:output_type -> cwztpsvc.DeleteStaticRouteResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_cw_ztp_staticroute_cw_ztp_staticroute_proto_init() }
func file_cw_ztp_staticroute_cw_ztp_staticroute_proto_init() {
	if File_cw_ztp_staticroute_cw_ztp_staticroute_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaticRoute); i {
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
		file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddStaticRoute); i {
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
		file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddStaticRouteRequest); i {
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
		file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddStaticRouteResponse); i {
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
		file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaticRouteGetRequest); i {
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
		file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaticRouteGetResponse); i {
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
		file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStaticRoute); i {
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
		file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStaticRouteRequest); i {
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
		file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStaticRouteResponse); i {
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
			RawDescriptor: file_cw_ztp_staticroute_cw_ztp_staticroute_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cw_ztp_staticroute_cw_ztp_staticroute_proto_goTypes,
		DependencyIndexes: file_cw_ztp_staticroute_cw_ztp_staticroute_proto_depIdxs,
		MessageInfos:      file_cw_ztp_staticroute_cw_ztp_staticroute_proto_msgTypes,
	}.Build()
	File_cw_ztp_staticroute_cw_ztp_staticroute_proto = out.File
	file_cw_ztp_staticroute_cw_ztp_staticroute_proto_rawDesc = nil
	file_cw_ztp_staticroute_cw_ztp_staticroute_proto_goTypes = nil
	file_cw_ztp_staticroute_cw_ztp_staticroute_proto_depIdxs = nil
}