//
// Copyright 2019, Offchain Labs, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.10.1
// source: server.proto

package validatorserver

import (
	proto "github.com/golang/protobuf/proto"
	evm "github.com/offchainlabs/arbitrum/packages/arb-validator-core/evm"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type FindLogsArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromHeight string   `protobuf:"bytes,1,opt,name=fromHeight,proto3" json:"fromHeight,omitempty"`
	ToHeight   string   `protobuf:"bytes,2,opt,name=toHeight,proto3" json:"toHeight,omitempty"`
	Address    string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Topics     []string `protobuf:"bytes,4,rep,name=topics,proto3" json:"topics,omitempty"`
}

func (x *FindLogsArgs) Reset() {
	*x = FindLogsArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindLogsArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindLogsArgs) ProtoMessage() {}

func (x *FindLogsArgs) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindLogsArgs.ProtoReflect.Descriptor instead.
func (*FindLogsArgs) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0}
}

func (x *FindLogsArgs) GetFromHeight() string {
	if x != nil {
		return x.FromHeight
	}
	return ""
}

func (x *FindLogsArgs) GetToHeight() string {
	if x != nil {
		return x.ToHeight
	}
	return ""
}

func (x *FindLogsArgs) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *FindLogsArgs) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

type FindLogsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logs []*evm.FullLogBuf `protobuf:"bytes,4,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (x *FindLogsReply) Reset() {
	*x = FindLogsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindLogsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindLogsReply) ProtoMessage() {}

func (x *FindLogsReply) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindLogsReply.ProtoReflect.Descriptor instead.
func (*FindLogsReply) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{1}
}

func (x *FindLogsReply) GetLogs() []*evm.FullLogBuf {
	if x != nil {
		return x.Logs
	}
	return nil
}

type GetOutputMessageArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssertionNodeHash string `protobuf:"bytes,1,opt,name=AssertionNodeHash,proto3" json:"AssertionNodeHash,omitempty"`
	MsgIndex          string `protobuf:"bytes,2,opt,name=MsgIndex,proto3" json:"MsgIndex,omitempty"`
}

func (x *GetOutputMessageArgs) Reset() {
	*x = GetOutputMessageArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOutputMessageArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOutputMessageArgs) ProtoMessage() {}

func (x *GetOutputMessageArgs) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOutputMessageArgs.ProtoReflect.Descriptor instead.
func (*GetOutputMessageArgs) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{2}
}

func (x *GetOutputMessageArgs) GetAssertionNodeHash() string {
	if x != nil {
		return x.AssertionNodeHash
	}
	return ""
}

func (x *GetOutputMessageArgs) GetMsgIndex() string {
	if x != nil {
		return x.MsgIndex
	}
	return ""
}

type GetOutputMessageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Found  bool   `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
	RawVal string `protobuf:"bytes,2,opt,name=rawVal,proto3" json:"rawVal,omitempty"`
}

func (x *GetOutputMessageReply) Reset() {
	*x = GetOutputMessageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOutputMessageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOutputMessageReply) ProtoMessage() {}

func (x *GetOutputMessageReply) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOutputMessageReply.ProtoReflect.Descriptor instead.
func (*GetOutputMessageReply) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{3}
}

func (x *GetOutputMessageReply) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

func (x *GetOutputMessageReply) GetRawVal() string {
	if x != nil {
		return x.RawVal
	}
	return ""
}

type GetMessageResultArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxHash string `protobuf:"bytes,1,opt,name=txHash,proto3" json:"txHash,omitempty"`
}

func (x *GetMessageResultArgs) Reset() {
	*x = GetMessageResultArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageResultArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageResultArgs) ProtoMessage() {}

func (x *GetMessageResultArgs) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageResultArgs.ProtoReflect.Descriptor instead.
func (*GetMessageResultArgs) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{4}
}

func (x *GetMessageResultArgs) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

type GetMessageResultReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tx *evm.TxInfoBuf `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (x *GetMessageResultReply) Reset() {
	*x = GetMessageResultReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageResultReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageResultReply) ProtoMessage() {}

func (x *GetMessageResultReply) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageResultReply.ProtoReflect.Descriptor instead.
func (*GetMessageResultReply) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{5}
}

func (x *GetMessageResultReply) GetTx() *evm.TxInfoBuf {
	if x != nil {
		return x.Tx
	}
	return nil
}

type GetAssertionCountArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAssertionCountArgs) Reset() {
	*x = GetAssertionCountArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssertionCountArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssertionCountArgs) ProtoMessage() {}

func (x *GetAssertionCountArgs) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssertionCountArgs.ProtoReflect.Descriptor instead.
func (*GetAssertionCountArgs) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{6}
}

type GetAssertionCountReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssertionCount int32 `protobuf:"varint,1,opt,name=assertionCount,proto3" json:"assertionCount,omitempty"`
}

func (x *GetAssertionCountReply) Reset() {
	*x = GetAssertionCountReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssertionCountReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssertionCountReply) ProtoMessage() {}

func (x *GetAssertionCountReply) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssertionCountReply.ProtoReflect.Descriptor instead.
func (*GetAssertionCountReply) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{7}
}

func (x *GetAssertionCountReply) GetAssertionCount() int32 {
	if x != nil {
		return x.AssertionCount
	}
	return 0
}

type GetVMInfoArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetVMInfoArgs) Reset() {
	*x = GetVMInfoArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVMInfoArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVMInfoArgs) ProtoMessage() {}

func (x *GetVMInfoArgs) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVMInfoArgs.ProtoReflect.Descriptor instead.
func (*GetVMInfoArgs) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{8}
}

type GetVMInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VmID string `protobuf:"bytes,1,opt,name=vmID,proto3" json:"vmID,omitempty"`
}

func (x *GetVMInfoReply) Reset() {
	*x = GetVMInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVMInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVMInfoReply) ProtoMessage() {}

func (x *GetVMInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVMInfoReply.ProtoReflect.Descriptor instead.
func (*GetVMInfoReply) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{9}
}

func (x *GetVMInfoReply) GetVmID() string {
	if x != nil {
		return x.VmID
	}
	return ""
}

type CallMessageArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContractAddress string `protobuf:"bytes,1,opt,name=contractAddress,proto3" json:"contractAddress,omitempty"`
	Sender          string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Data            string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CallMessageArgs) Reset() {
	*x = CallMessageArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallMessageArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallMessageArgs) ProtoMessage() {}

func (x *CallMessageArgs) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallMessageArgs.ProtoReflect.Descriptor instead.
func (*CallMessageArgs) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{10}
}

func (x *CallMessageArgs) GetContractAddress() string {
	if x != nil {
		return x.ContractAddress
	}
	return ""
}

func (x *CallMessageArgs) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *CallMessageArgs) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type CallMessageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawVal string `protobuf:"bytes,1,opt,name=rawVal,proto3" json:"rawVal,omitempty"`
}

func (x *CallMessageReply) Reset() {
	*x = CallMessageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallMessageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallMessageReply) ProtoMessage() {}

func (x *CallMessageReply) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallMessageReply.ProtoReflect.Descriptor instead.
func (*CallMessageReply) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{11}
}

func (x *CallMessageReply) GetRawVal() string {
	if x != nil {
		return x.RawVal
	}
	return ""
}

var File_server_proto protoreflect.FileDescriptor

var file_server_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a,
	0x0d, 0x65, 0x76, 0x6d, 0x2f, 0x65, 0x76, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c,
	0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x73, 0x41, 0x72, 0x67, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x74, 0x6f, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x6f, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x22, 0x34, 0x0a, 0x0d,
	0x46, 0x69, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x23, 0x0a,
	0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x65, 0x76,
	0x6d, 0x2e, 0x46, 0x75, 0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x42, 0x75, 0x66, 0x52, 0x04, 0x6c, 0x6f,
	0x67, 0x73, 0x22, 0x60, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x72, 0x67, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x41, 0x73,
	0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x41, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x6f, 0x64, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x73, 0x67, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4d, 0x73, 0x67, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x22, 0x45, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x77, 0x56, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x61, 0x77, 0x56, 0x61, 0x6c, 0x22, 0x2e, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x41,
	0x72, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x22, 0x37, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x1e, 0x0a, 0x02, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x65, 0x76, 0x6d, 0x2e, 0x54, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x75, 0x66,
	0x52, 0x02, 0x74, 0x78, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x72,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x72, 0x67, 0x73, 0x22, 0x40, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x73, 0x73, 0x65, 0x72,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0e, 0x61, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x56, 0x4d, 0x49, 0x6e, 0x66, 0x6f, 0x41, 0x72, 0x67, 0x73,
	0x22, 0x24, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x56, 0x4d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x76, 0x6d, 0x49, 0x44, 0x22, 0x67, 0x0a, 0x0f, 0x43, 0x61, 0x6c, 0x6c, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x72, 0x67, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x2a, 0x0a, 0x10, 0x43, 0x61, 0x6c, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x77, 0x56, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x61, 0x77, 0x56, 0x61, 0x6c, 0x32, 0xaa, 0x04, 0x0a, 0x0f,
	0x52, 0x6f, 0x6c, 0x6c, 0x75, 0x70, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x61, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x25, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x26, 0x2e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x61, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x25, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x26, 0x2e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x52, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x21, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x49, 0x0a, 0x08, 0x46, 0x69, 0x6e,
	0x64, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x1d, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x73,
	0x41, 0x72, 0x67, 0x73, 0x1a, 0x1e, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x64, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x72,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x2e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x72, 0x67,
	0x73, 0x1a, 0x27, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4c, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x56, 0x4d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x4d, 0x49,
	0x6e, 0x66, 0x6f, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x1f, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x4d, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x66, 0x66, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c,
	0x61, 0x62, 0x73, 0x2f, 0x61, 0x72, 0x62, 0x69, 0x74, 0x72, 0x75, 0x6d, 0x2f, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x61, 0x72, 0x62, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_proto_rawDescOnce sync.Once
	file_server_proto_rawDescData = file_server_proto_rawDesc
)

func file_server_proto_rawDescGZIP() []byte {
	file_server_proto_rawDescOnce.Do(func() {
		file_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_proto_rawDescData)
	})
	return file_server_proto_rawDescData
}

var file_server_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_server_proto_goTypes = []interface{}{
	(*FindLogsArgs)(nil),           // 0: validatorserver.FindLogsArgs
	(*FindLogsReply)(nil),          // 1: validatorserver.FindLogsReply
	(*GetOutputMessageArgs)(nil),   // 2: validatorserver.GetOutputMessageArgs
	(*GetOutputMessageReply)(nil),  // 3: validatorserver.GetOutputMessageReply
	(*GetMessageResultArgs)(nil),   // 4: validatorserver.GetMessageResultArgs
	(*GetMessageResultReply)(nil),  // 5: validatorserver.GetMessageResultReply
	(*GetAssertionCountArgs)(nil),  // 6: validatorserver.GetAssertionCountArgs
	(*GetAssertionCountReply)(nil), // 7: validatorserver.GetAssertionCountReply
	(*GetVMInfoArgs)(nil),          // 8: validatorserver.GetVMInfoArgs
	(*GetVMInfoReply)(nil),         // 9: validatorserver.GetVMInfoReply
	(*CallMessageArgs)(nil),        // 10: validatorserver.CallMessageArgs
	(*CallMessageReply)(nil),       // 11: validatorserver.CallMessageReply
	(*evm.FullLogBuf)(nil),         // 12: evm.FullLogBuf
	(*evm.TxInfoBuf)(nil),          // 13: evm.TxInfoBuf
}
var file_server_proto_depIdxs = []int32{
	12, // 0: validatorserver.FindLogsReply.logs:type_name -> evm.FullLogBuf
	13, // 1: validatorserver.GetMessageResultReply.tx:type_name -> evm.TxInfoBuf
	2,  // 2: validatorserver.RollupValidator.GetOutputMessage:input_type -> validatorserver.GetOutputMessageArgs
	4,  // 3: validatorserver.RollupValidator.GetMessageResult:input_type -> validatorserver.GetMessageResultArgs
	10, // 4: validatorserver.RollupValidator.CallMessage:input_type -> validatorserver.CallMessageArgs
	0,  // 5: validatorserver.RollupValidator.FindLogs:input_type -> validatorserver.FindLogsArgs
	6,  // 6: validatorserver.RollupValidator.GetAssertionCount:input_type -> validatorserver.GetAssertionCountArgs
	8,  // 7: validatorserver.RollupValidator.GetVMInfo:input_type -> validatorserver.GetVMInfoArgs
	3,  // 8: validatorserver.RollupValidator.GetOutputMessage:output_type -> validatorserver.GetOutputMessageReply
	5,  // 9: validatorserver.RollupValidator.GetMessageResult:output_type -> validatorserver.GetMessageResultReply
	11, // 10: validatorserver.RollupValidator.CallMessage:output_type -> validatorserver.CallMessageReply
	1,  // 11: validatorserver.RollupValidator.FindLogs:output_type -> validatorserver.FindLogsReply
	7,  // 12: validatorserver.RollupValidator.GetAssertionCount:output_type -> validatorserver.GetAssertionCountReply
	9,  // 13: validatorserver.RollupValidator.GetVMInfo:output_type -> validatorserver.GetVMInfoReply
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_server_proto_init() }
func file_server_proto_init() {
	if File_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindLogsArgs); i {
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
		file_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindLogsReply); i {
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
		file_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOutputMessageArgs); i {
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
		file_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOutputMessageReply); i {
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
		file_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageResultArgs); i {
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
		file_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageResultReply); i {
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
		file_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAssertionCountArgs); i {
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
		file_server_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAssertionCountReply); i {
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
		file_server_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVMInfoArgs); i {
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
		file_server_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVMInfoReply); i {
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
		file_server_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallMessageArgs); i {
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
		file_server_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallMessageReply); i {
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
			RawDescriptor: file_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_proto_goTypes,
		DependencyIndexes: file_server_proto_depIdxs,
		MessageInfos:      file_server_proto_msgTypes,
	}.Build()
	File_server_proto = out.File
	file_server_proto_rawDesc = nil
	file_server_proto_goTypes = nil
	file_server_proto_depIdxs = nil
}
