// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: api/proto/file.proto

package proto

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

type OpenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Caller *Caller `protobuf:"bytes,1,opt,name=caller,proto3" json:"caller,omitempty"`
	Path   string  `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Flags  uint32  `protobuf:"varint,3,opt,name=flags,proto3" json:"flags,omitempty"`
}

func (x *OpenRequest) Reset() {
	*x = OpenRequest{}
	mi := &file_api_proto_file_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OpenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenRequest) ProtoMessage() {}

func (x *OpenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_file_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenRequest.ProtoReflect.Descriptor instead.
func (*OpenRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_file_proto_rawDescGZIP(), []int{0}
}

func (x *OpenRequest) GetCaller() *Caller {
	if x != nil {
		return x.Caller
	}
	return nil
}

func (x *OpenRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *OpenRequest) GetFlags() uint32 {
	if x != nil {
		return x.Flags
	}
	return 0
}

type OpenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *OpenReply) Reset() {
	*x = OpenReply{}
	mi := &file_api_proto_file_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OpenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenReply) ProtoMessage() {}

func (x *OpenReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_file_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenReply.ProtoReflect.Descriptor instead.
func (*OpenReply) Descriptor() ([]byte, []int) {
	return file_api_proto_file_proto_rawDescGZIP(), []int{1}
}

func (x *OpenReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Caller *Caller `protobuf:"bytes,1,opt,name=caller,proto3" json:"caller,omitempty"`
	Path   string  `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Flags  uint32  `protobuf:"varint,3,opt,name=flags,proto3" json:"flags,omitempty"`
	Mode   uint32  `protobuf:"varint,4,opt,name=mode,proto3" json:"mode,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	mi := &file_api_proto_file_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_file_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_file_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRequest) GetCaller() *Caller {
	if x != nil {
		return x.Caller
	}
	return nil
}

func (x *CreateRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *CreateRequest) GetFlags() uint32 {
	if x != nil {
		return x.Flags
	}
	return 0
}

func (x *CreateRequest) GetMode() uint32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

type CreateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateReply) Reset() {
	*x = CreateReply{}
	mi := &file_api_proto_file_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReply) ProtoMessage() {}

func (x *CreateReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_file_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReply.ProtoReflect.Descriptor instead.
func (*CreateReply) Descriptor() ([]byte, []int) {
	return file_api_proto_file_proto_rawDescGZIP(), []int{3}
}

func (x *CreateReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path   string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Offset int64  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Size   uint32 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	mi := &file_api_proto_file_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_file_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_file_proto_rawDescGZIP(), []int{4}
}

func (x *ReadRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ReadRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ReadRequest) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ReadReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bytes  []byte `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Size   int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Status int32  `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ReadReply) Reset() {
	*x = ReadReply{}
	mi := &file_api_proto_file_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadReply) ProtoMessage() {}

func (x *ReadReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_file_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadReply.ProtoReflect.Descriptor instead.
func (*ReadReply) Descriptor() ([]byte, []int) {
	return file_api_proto_file_proto_rawDescGZIP(), []int{5}
}

func (x *ReadReply) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *ReadReply) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ReadReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type WriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path   string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Bytes  []byte `protobuf:"bytes,2,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Offset int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	mi := &file_api_proto_file_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_file_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_file_proto_rawDescGZIP(), []int{6}
}

func (x *WriteRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *WriteRequest) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *WriteRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type WriteReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Written uint32 `protobuf:"varint,1,opt,name=written,proto3" json:"written,omitempty"`
	Status  int32  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *WriteReply) Reset() {
	*x = WriteReply{}
	mi := &file_api_proto_file_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WriteReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteReply) ProtoMessage() {}

func (x *WriteReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_file_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteReply.ProtoReflect.Descriptor instead.
func (*WriteReply) Descriptor() ([]byte, []int) {
	return file_api_proto_file_proto_rawDescGZIP(), []int{7}
}

func (x *WriteReply) GetWritten() uint32 {
	if x != nil {
		return x.Written
	}
	return 0
}

func (x *WriteReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_api_proto_file_proto protoreflect.FileDescriptor

var file_api_proto_file_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x65,
	0x1a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x0b, 0x4f, 0x70, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x69, 0x65, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x22, 0x23, 0x0a, 0x09, 0x4f,
	0x70, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x77, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x28, 0x0a, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x65, 0x2e, 0x43, 0x61, 0x6c,
	0x6c, 0x65, 0x72, 0x52, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x25, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x4d, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22,
	0x4d, 0x0a, 0x09, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x50,
	0x0a, 0x0c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x22, 0x3e, 0x0a, 0x0a, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x77, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x77, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x32, 0xe2, 0x01, 0x0a, 0x07, 0x52, 0x70, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x32, 0x0a, 0x04,
	0x4f, 0x70, 0x65, 0x6e, 0x12, 0x15, 0x2e, 0x67, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x65, 0x2e,
	0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x69, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x38, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x67, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x69, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x32, 0x0a, 0x04, 0x52, 0x65,
	0x61, 0x64, 0x12, 0x15, 0x2e, 0x67, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x65, 0x2e, 0x52, 0x65,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x69, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x35,
	0x0a, 0x05, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x69, 0x65, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x67, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x65, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_file_proto_rawDescOnce sync.Once
	file_api_proto_file_proto_rawDescData = file_api_proto_file_proto_rawDesc
)

func file_api_proto_file_proto_rawDescGZIP() []byte {
	file_api_proto_file_proto_rawDescOnce.Do(func() {
		file_api_proto_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_file_proto_rawDescData)
	})
	return file_api_proto_file_proto_rawDescData
}

var file_api_proto_file_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_proto_file_proto_goTypes = []any{
	(*OpenRequest)(nil),   // 0: gmountie.OpenRequest
	(*OpenReply)(nil),     // 1: gmountie.OpenReply
	(*CreateRequest)(nil), // 2: gmountie.CreateRequest
	(*CreateReply)(nil),   // 3: gmountie.CreateReply
	(*ReadRequest)(nil),   // 4: gmountie.ReadRequest
	(*ReadReply)(nil),     // 5: gmountie.ReadReply
	(*WriteRequest)(nil),  // 6: gmountie.WriteRequest
	(*WriteReply)(nil),    // 7: gmountie.WriteReply
	(*Caller)(nil),        // 8: gmountie.Caller
}
var file_api_proto_file_proto_depIdxs = []int32{
	8, // 0: gmountie.OpenRequest.caller:type_name -> gmountie.Caller
	8, // 1: gmountie.CreateRequest.caller:type_name -> gmountie.Caller
	0, // 2: gmountie.RpcFile.Open:input_type -> gmountie.OpenRequest
	2, // 3: gmountie.RpcFile.Create:input_type -> gmountie.CreateRequest
	4, // 4: gmountie.RpcFile.Read:input_type -> gmountie.ReadRequest
	6, // 5: gmountie.RpcFile.Write:input_type -> gmountie.WriteRequest
	1, // 6: gmountie.RpcFile.Open:output_type -> gmountie.OpenReply
	3, // 7: gmountie.RpcFile.Create:output_type -> gmountie.CreateReply
	5, // 8: gmountie.RpcFile.Read:output_type -> gmountie.ReadReply
	7, // 9: gmountie.RpcFile.Write:output_type -> gmountie.WriteReply
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_file_proto_init() }
func file_api_proto_file_proto_init() {
	if File_api_proto_file_proto != nil {
		return
	}
	file_api_proto_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_file_proto_goTypes,
		DependencyIndexes: file_api_proto_file_proto_depIdxs,
		MessageInfos:      file_api_proto_file_proto_msgTypes,
	}.Build()
	File_api_proto_file_proto = out.File
	file_api_proto_file_proto_rawDesc = nil
	file_api_proto_file_proto_goTypes = nil
	file_api_proto_file_proto_depIdxs = nil
}
