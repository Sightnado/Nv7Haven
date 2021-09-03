// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: proto/anarchy.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AnarchyElement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color      string   `protobuf:"bytes,1,opt,name=Color,proto3" json:"Color,omitempty"`
	Comment    string   `protobuf:"bytes,2,opt,name=Comment,proto3" json:"Comment,omitempty"`
	CreatedOn  int64    `protobuf:"varint,3,opt,name=CreatedOn,proto3" json:"CreatedOn,omitempty"`
	Creator    string   `protobuf:"bytes,4,opt,name=Creator,proto3" json:"Creator,omitempty"`
	Name       string   `protobuf:"bytes,5,opt,name=Name,proto3" json:"Name,omitempty"`
	Parents    []string `protobuf:"bytes,6,rep,name=Parents,proto3" json:"Parents,omitempty"`
	Uses       int64    `protobuf:"varint,8,opt,name=Uses,proto3" json:"Uses,omitempty"`
	FoundBy    int64    `protobuf:"varint,9,opt,name=FoundBy,proto3" json:"FoundBy,omitempty"`
	Complexity int64    `protobuf:"varint,10,opt,name=Complexity,proto3" json:"Complexity,omitempty"`
}

func (x *AnarchyElement) Reset() {
	*x = AnarchyElement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_anarchy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnarchyElement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnarchyElement) ProtoMessage() {}

func (x *AnarchyElement) ProtoReflect() protoreflect.Message {
	mi := &file_proto_anarchy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnarchyElement.ProtoReflect.Descriptor instead.
func (*AnarchyElement) Descriptor() ([]byte, []int) {
	return file_proto_anarchy_proto_rawDescGZIP(), []int{0}
}

func (x *AnarchyElement) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *AnarchyElement) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *AnarchyElement) GetCreatedOn() int64 {
	if x != nil {
		return x.CreatedOn
	}
	return 0
}

func (x *AnarchyElement) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *AnarchyElement) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AnarchyElement) GetParents() []string {
	if x != nil {
		return x.Parents
	}
	return nil
}

func (x *AnarchyElement) GetUses() int64 {
	if x != nil {
		return x.Uses
	}
	return 0
}

func (x *AnarchyElement) GetFoundBy() int64 {
	if x != nil {
		return x.FoundBy
	}
	return 0
}

func (x *AnarchyElement) GetComplexity() int64 {
	if x != nil {
		return x.Complexity
	}
	return 0
}

type AnarchyCombination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Elem1 string `protobuf:"bytes,1,opt,name=elem1,proto3" json:"elem1,omitempty"`
	Elem2 string `protobuf:"bytes,2,opt,name=elem2,proto3" json:"elem2,omitempty"`
}

func (x *AnarchyCombination) Reset() {
	*x = AnarchyCombination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_anarchy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnarchyCombination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnarchyCombination) ProtoMessage() {}

func (x *AnarchyCombination) ProtoReflect() protoreflect.Message {
	mi := &file_proto_anarchy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnarchyCombination.ProtoReflect.Descriptor instead.
func (*AnarchyCombination) Descriptor() ([]byte, []int) {
	return file_proto_anarchy_proto_rawDescGZIP(), []int{1}
}

func (x *AnarchyCombination) GetElem1() string {
	if x != nil {
		return x.Elem1
	}
	return ""
}

func (x *AnarchyCombination) GetElem2() string {
	if x != nil {
		return x.Elem2
	}
	return ""
}

type AnarchyCombinationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data   string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Exists bool   `protobuf:"varint,2,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *AnarchyCombinationResult) Reset() {
	*x = AnarchyCombinationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_anarchy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnarchyCombinationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnarchyCombinationResult) ProtoMessage() {}

func (x *AnarchyCombinationResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_anarchy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnarchyCombinationResult.ProtoReflect.Descriptor instead.
func (*AnarchyCombinationResult) Descriptor() ([]byte, []int) {
	return file_proto_anarchy_proto_rawDescGZIP(), []int{2}
}

func (x *AnarchyCombinationResult) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *AnarchyCombinationResult) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

type AnarchyInventory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Found []string `protobuf:"bytes,1,rep,name=Found,proto3" json:"Found,omitempty"`
}

func (x *AnarchyInventory) Reset() {
	*x = AnarchyInventory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_anarchy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnarchyInventory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnarchyInventory) ProtoMessage() {}

func (x *AnarchyInventory) ProtoReflect() protoreflect.Message {
	mi := &file_proto_anarchy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnarchyInventory.ProtoReflect.Descriptor instead.
func (*AnarchyInventory) Descriptor() ([]byte, []int) {
	return file_proto_anarchy_proto_rawDescGZIP(), []int{3}
}

func (x *AnarchyInventory) GetFound() []string {
	if x != nil {
		return x.Found
	}
	return nil
}

type AnarchyUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid     string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Element string `protobuf:"bytes,2,opt,name=element,proto3" json:"element,omitempty"`
}

func (x *AnarchyUserRequest) Reset() {
	*x = AnarchyUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_anarchy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnarchyUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnarchyUserRequest) ProtoMessage() {}

func (x *AnarchyUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_anarchy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnarchyUserRequest.ProtoReflect.Descriptor instead.
func (*AnarchyUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_anarchy_proto_rawDescGZIP(), []int{4}
}

func (x *AnarchyUserRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *AnarchyUserRequest) GetElement() string {
	if x != nil {
		return x.Element
	}
	return ""
}

var File_proto_anarchy_proto protoreflect.FileDescriptor

var file_proto_anarchy_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x6e, 0x61, 0x72, 0x63, 0x68, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x6e, 0x61, 0x72, 0x63, 0x68, 0x79, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x01, 0x0a, 0x0e,
	0x41, 0x6e, 0x61, 0x72, 0x63, 0x68, 0x79, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43,
	0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x50, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x55, 0x73, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x6f, 0x75, 0x6e,
	0x64, 0x42, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x46, 0x6f, 0x75, 0x6e, 0x64,
	0x42, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x69, 0x74, 0x79,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x69,
	0x74, 0x79, 0x22, 0x40, 0x0a, 0x12, 0x41, 0x6e, 0x61, 0x72, 0x63, 0x68, 0x79, 0x43, 0x6f, 0x6d,
	0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6c, 0x65, 0x6d,
	0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6c, 0x65, 0x6d, 0x31, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6c, 0x65, 0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6c, 0x65, 0x6d, 0x32, 0x22, 0x46, 0x0a, 0x18, 0x41, 0x6e, 0x61, 0x72, 0x63, 0x68, 0x79, 0x43,
	0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x22, 0x28, 0x0a, 0x10,
	0x41, 0x6e, 0x61, 0x72, 0x63, 0x68, 0x79, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x40, 0x0a, 0x12, 0x41, 0x6e, 0x61, 0x72, 0x63, 0x68,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0xa9, 0x02, 0x0a, 0x07, 0x41, 0x6e, 0x61,
	0x72, 0x63, 0x68, 0x79, 0x12, 0x42, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x12,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x17, 0x2e,
	0x61, 0x6e, 0x61, 0x72, 0x63, 0x68, 0x79, 0x2e, 0x41, 0x6e, 0x61, 0x72, 0x63, 0x68, 0x79, 0x45,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x61, 0x6e, 0x61,
	0x72, 0x63, 0x68, 0x79, 0x2e, 0x41, 0x6e, 0x61, 0x72, 0x63, 0x68, 0x79, 0x43, 0x6f, 0x6d, 0x62,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x21, 0x2e, 0x61, 0x6e, 0x61, 0x72, 0x63, 0x68,
	0x79, 0x2e, 0x41, 0x6e, 0x61, 0x72, 0x63, 0x68, 0x79, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x06,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x1a, 0x19, 0x2e, 0x61, 0x6e, 0x61, 0x72, 0x63, 0x68, 0x79, 0x2e, 0x41,
	0x6e, 0x61, 0x72, 0x63, 0x68, 0x79, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x22,
	0x00, 0x12, 0x41, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1b, 0x2e,
	0x61, 0x6e, 0x61, 0x72, 0x63, 0x68, 0x79, 0x2e, 0x41, 0x6e, 0x61, 0x72, 0x63, 0x68, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x4e, 0x76, 0x37, 0x68, 0x61,
	0x76, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_anarchy_proto_rawDescOnce sync.Once
	file_proto_anarchy_proto_rawDescData = file_proto_anarchy_proto_rawDesc
)

func file_proto_anarchy_proto_rawDescGZIP() []byte {
	file_proto_anarchy_proto_rawDescOnce.Do(func() {
		file_proto_anarchy_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_anarchy_proto_rawDescData)
	})
	return file_proto_anarchy_proto_rawDescData
}

var file_proto_anarchy_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_anarchy_proto_goTypes = []interface{}{
	(*AnarchyElement)(nil),           // 0: anarchy.AnarchyElement
	(*AnarchyCombination)(nil),       // 1: anarchy.AnarchyCombination
	(*AnarchyCombinationResult)(nil), // 2: anarchy.AnarchyCombinationResult
	(*AnarchyInventory)(nil),         // 3: anarchy.AnarchyInventory
	(*AnarchyUserRequest)(nil),       // 4: anarchy.AnarchyUserRequest
	(*wrapperspb.StringValue)(nil),   // 5: google.protobuf.StringValue
	(*emptypb.Empty)(nil),            // 6: google.protobuf.Empty
}
var file_proto_anarchy_proto_depIdxs = []int32{
	5, // 0: anarchy.Anarchy.GetElem:input_type -> google.protobuf.StringValue
	1, // 1: anarchy.Anarchy.GetCombination:input_type -> anarchy.AnarchyCombination
	5, // 2: anarchy.Anarchy.GetInv:input_type -> google.protobuf.StringValue
	4, // 3: anarchy.Anarchy.AddFound:input_type -> anarchy.AnarchyUserRequest
	0, // 4: anarchy.Anarchy.GetElem:output_type -> anarchy.AnarchyElement
	2, // 5: anarchy.Anarchy.GetCombination:output_type -> anarchy.AnarchyCombinationResult
	3, // 6: anarchy.Anarchy.GetInv:output_type -> anarchy.AnarchyInventory
	6, // 7: anarchy.Anarchy.AddFound:output_type -> google.protobuf.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_anarchy_proto_init() }
func file_proto_anarchy_proto_init() {
	if File_proto_anarchy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_anarchy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnarchyElement); i {
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
		file_proto_anarchy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnarchyCombination); i {
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
		file_proto_anarchy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnarchyCombinationResult); i {
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
		file_proto_anarchy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnarchyInventory); i {
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
		file_proto_anarchy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnarchyUserRequest); i {
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
			RawDescriptor: file_proto_anarchy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_anarchy_proto_goTypes,
		DependencyIndexes: file_proto_anarchy_proto_depIdxs,
		MessageInfos:      file_proto_anarchy_proto_msgTypes,
	}.Build()
	File_proto_anarchy_proto = out.File
	file_proto_anarchy_proto_rawDesc = nil
	file_proto_anarchy_proto_goTypes = nil
	file_proto_anarchy_proto_depIdxs = nil
}
