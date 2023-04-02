// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: model/proto/route/routetypes.proto

package route

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

type RouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdAccount string `protobuf:"bytes,1,opt,name=id_account,json=idAccount,proto3" json:"id_account,omitempty"`
}

func (x *RouteRequest) Reset() {
	*x = RouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_route_routetypes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteRequest) ProtoMessage() {}

func (x *RouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_route_routetypes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteRequest.ProtoReflect.Descriptor instead.
func (*RouteRequest) Descriptor() ([]byte, []int) {
	return file_model_proto_route_routetypes_proto_rawDescGZIP(), []int{0}
}

func (x *RouteRequest) GetIdAccount() string {
	if x != nil {
		return x.IdAccount
	}
	return ""
}

type LocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdAccount string `protobuf:"bytes,1,opt,name=id_account,json=idAccount,proto3" json:"id_account,omitempty"`
	Point     *Point `protobuf:"bytes,2,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *LocationRequest) Reset() {
	*x = LocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_route_routetypes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationRequest) ProtoMessage() {}

func (x *LocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_route_routetypes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationRequest.ProtoReflect.Descriptor instead.
func (*LocationRequest) Descriptor() ([]byte, []int) {
	return file_model_proto_route_routetypes_proto_rawDescGZIP(), []int{1}
}

func (x *LocationRequest) GetIdAccount() string {
	if x != nil {
		return x.IdAccount
	}
	return ""
}

func (x *LocationRequest) GetPoint() *Point {
	if x != nil {
		return x.Point
	}
	return nil
}

type DestintationAndPolylineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdAccount   string         `protobuf:"bytes,1,opt,name=id_account,json=idAccount,proto3" json:"id_account,omitempty"`
	Destination *Point         `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	Polyline    *RoutePolyline `protobuf:"bytes,3,opt,name=polyline,proto3" json:"polyline,omitempty"`
}

func (x *DestintationAndPolylineRequest) Reset() {
	*x = DestintationAndPolylineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_route_routetypes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestintationAndPolylineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestintationAndPolylineRequest) ProtoMessage() {}

func (x *DestintationAndPolylineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_route_routetypes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestintationAndPolylineRequest.ProtoReflect.Descriptor instead.
func (*DestintationAndPolylineRequest) Descriptor() ([]byte, []int) {
	return file_model_proto_route_routetypes_proto_rawDescGZIP(), []int{2}
}

func (x *DestintationAndPolylineRequest) GetIdAccount() string {
	if x != nil {
		return x.IdAccount
	}
	return ""
}

func (x *DestintationAndPolylineRequest) GetDestination() *Point {
	if x != nil {
		return x.Destination
	}
	return nil
}

func (x *DestintationAndPolylineRequest) GetPolyline() *RoutePolyline {
	if x != nil {
		return x.Polyline
	}
	return nil
}

type LocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Point   *Point `protobuf:"bytes,3,opt,name=point,proto3,oneof" json:"point,omitempty"`
}

func (x *LocationResponse) Reset() {
	*x = LocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_route_routetypes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationResponse) ProtoMessage() {}

func (x *LocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_route_routetypes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationResponse.ProtoReflect.Descriptor instead.
func (*LocationResponse) Descriptor() ([]byte, []int) {
	return file_model_proto_route_routetypes_proto_rawDescGZIP(), []int{3}
}

func (x *LocationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *LocationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LocationResponse) GetPoint() *Point {
	if x != nil {
		return x.Point
	}
	return nil
}

type DestintationAndPolylineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success     bool           `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message     string         `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data        *RoutePolyline `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data,omitempty"`
	Destination *Point         `protobuf:"bytes,4,opt,name=destination,proto3,oneof" json:"destination,omitempty"`
}

func (x *DestintationAndPolylineResponse) Reset() {
	*x = DestintationAndPolylineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_route_routetypes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestintationAndPolylineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestintationAndPolylineResponse) ProtoMessage() {}

func (x *DestintationAndPolylineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_route_routetypes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestintationAndPolylineResponse.ProtoReflect.Descriptor instead.
func (*DestintationAndPolylineResponse) Descriptor() ([]byte, []int) {
	return file_model_proto_route_routetypes_proto_rawDescGZIP(), []int{4}
}

func (x *DestintationAndPolylineResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DestintationAndPolylineResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DestintationAndPolylineResponse) GetData() *RoutePolyline {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DestintationAndPolylineResponse) GetDestination() *Point {
	if x != nil {
		return x.Destination
	}
	return nil
}

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_route_routetypes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_route_routetypes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_model_proto_route_routetypes_proto_rawDescGZIP(), []int{5}
}

func (x *Point) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Point) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type RoutePolyline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Points []*Point `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
}

func (x *RoutePolyline) Reset() {
	*x = RoutePolyline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_route_routetypes_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoutePolyline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutePolyline) ProtoMessage() {}

func (x *RoutePolyline) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_route_routetypes_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutePolyline.ProtoReflect.Descriptor instead.
func (*RoutePolyline) Descriptor() ([]byte, []int) {
	return file_model_proto_route_routetypes_proto_rawDescGZIP(), []int{6}
}

func (x *RoutePolyline) GetPoints() []*Point {
	if x != nil {
		return x.Points
	}
	return nil
}

var File_model_proto_route_routetypes_proto protoreflect.FileDescriptor

var file_model_proto_route_routetypes_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x74,
	0x72, 0x69, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x22, 0x2d, 0x0a, 0x0c, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x64, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69,
	0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x62, 0x0a, 0x0f, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69,
	0x64, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x69, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0xbd, 0x01, 0x0a,
	0x1e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x64,
	0x50, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x69, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3c,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x74,
	0x72, 0x69, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x08,
	0x70, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x79, 0x6c, 0x69,
	0x6e, 0x65, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x87, 0x01, 0x0a,
	0x10, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x35, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x74, 0x72, 0x69, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x48, 0x00, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0xee, 0x01, 0x0a, 0x1f, 0x44, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x64, 0x50, 0x6f, 0x6c, 0x79, 0x6c, 0x69,
	0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3b,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65,
	0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x12, 0x41, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x74, 0x72, 0x69, 0x70,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x48, 0x01, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x41, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x43, 0x0a, 0x0d, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x42,
	0x41, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4a, 0x6f, 0x63, 0x65, 0x72, 0x64, 0x69, 0x6b, 0x69, 0x61, 0x77, 0x61, 0x6e, 0x6e, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x74, 0x72, 0x69, 0x70,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_proto_route_routetypes_proto_rawDescOnce sync.Once
	file_model_proto_route_routetypes_proto_rawDescData = file_model_proto_route_routetypes_proto_rawDesc
)

func file_model_proto_route_routetypes_proto_rawDescGZIP() []byte {
	file_model_proto_route_routetypes_proto_rawDescOnce.Do(func() {
		file_model_proto_route_routetypes_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_proto_route_routetypes_proto_rawDescData)
	})
	return file_model_proto_route_routetypes_proto_rawDescData
}

var file_model_proto_route_routetypes_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_model_proto_route_routetypes_proto_goTypes = []interface{}{
	(*RouteRequest)(nil),                    // 0: app.sharetrip.route.RouteRequest
	(*LocationRequest)(nil),                 // 1: app.sharetrip.route.LocationRequest
	(*DestintationAndPolylineRequest)(nil),  // 2: app.sharetrip.route.DestintationAndPolylineRequest
	(*LocationResponse)(nil),                // 3: app.sharetrip.route.LocationResponse
	(*DestintationAndPolylineResponse)(nil), // 4: app.sharetrip.route.DestintationAndPolylineResponse
	(*Point)(nil),                           // 5: app.sharetrip.route.Point
	(*RoutePolyline)(nil),                   // 6: app.sharetrip.route.RoutePolyline
}
var file_model_proto_route_routetypes_proto_depIdxs = []int32{
	5, // 0: app.sharetrip.route.LocationRequest.point:type_name -> app.sharetrip.route.Point
	5, // 1: app.sharetrip.route.DestintationAndPolylineRequest.destination:type_name -> app.sharetrip.route.Point
	6, // 2: app.sharetrip.route.DestintationAndPolylineRequest.polyline:type_name -> app.sharetrip.route.RoutePolyline
	5, // 3: app.sharetrip.route.LocationResponse.point:type_name -> app.sharetrip.route.Point
	6, // 4: app.sharetrip.route.DestintationAndPolylineResponse.data:type_name -> app.sharetrip.route.RoutePolyline
	5, // 5: app.sharetrip.route.DestintationAndPolylineResponse.destination:type_name -> app.sharetrip.route.Point
	5, // 6: app.sharetrip.route.RoutePolyline.points:type_name -> app.sharetrip.route.Point
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_model_proto_route_routetypes_proto_init() }
func file_model_proto_route_routetypes_proto_init() {
	if File_model_proto_route_routetypes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_proto_route_routetypes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteRequest); i {
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
		file_model_proto_route_routetypes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationRequest); i {
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
		file_model_proto_route_routetypes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestintationAndPolylineRequest); i {
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
		file_model_proto_route_routetypes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationResponse); i {
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
		file_model_proto_route_routetypes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestintationAndPolylineResponse); i {
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
		file_model_proto_route_routetypes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_model_proto_route_routetypes_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoutePolyline); i {
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
	file_model_proto_route_routetypes_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_model_proto_route_routetypes_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_model_proto_route_routetypes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_proto_route_routetypes_proto_goTypes,
		DependencyIndexes: file_model_proto_route_routetypes_proto_depIdxs,
		MessageInfos:      file_model_proto_route_routetypes_proto_msgTypes,
	}.Build()
	File_model_proto_route_routetypes_proto = out.File
	file_model_proto_route_routetypes_proto_rawDesc = nil
	file_model_proto_route_routetypes_proto_goTypes = nil
	file_model_proto_route_routetypes_proto_depIdxs = nil
}