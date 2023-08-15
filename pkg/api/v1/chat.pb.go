// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.17.3
// source: pkg/api/v1/chat.proto

package v1

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

type ChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInput string `protobuf:"bytes,1,opt,name=user_input,json=userInput,proto3" json:"user_input,omitempty"`
}

func (x *ChatRequest) Reset() {
	*x = ChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_v1_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRequest) ProtoMessage() {}

func (x *ChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_v1_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatRequest.ProtoReflect.Descriptor instead.
func (*ChatRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_v1_chat_proto_rawDescGZIP(), []int{0}
}

func (x *ChatRequest) GetUserInput() string {
	if x != nil {
		return x.UserInput
	}
	return ""
}

type ChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssistantReply string `protobuf:"bytes,1,opt,name=assistant_reply,json=assistantReply,proto3" json:"assistant_reply,omitempty"`
}

func (x *ChatResponse) Reset() {
	*x = ChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_v1_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatResponse) ProtoMessage() {}

func (x *ChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_v1_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatResponse.ProtoReflect.Descriptor instead.
func (*ChatResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_v1_chat_proto_rawDescGZIP(), []int{1}
}

func (x *ChatResponse) GetAssistantReply() string {
	if x != nil {
		return x.AssistantReply
	}
	return ""
}

var File_pkg_api_v1_chat_proto protoreflect.FileDescriptor

var file_pkg_api_v1_chat_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31,
	0x22, 0x2c, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x37,
	0x0a, 0x0c, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27,
	0x0a, 0x0f, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6c,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x82, 0x01, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74,
	0x12, 0x3a, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0d,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68,
	0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x2b, 0x5a, 0x29,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x67, 0x65, 0x6f, 0x72,
	0x67, 0x69, 0x61, 0x64, 0x65, 0x73, 0x32, 0x37, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x68,
	0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_api_v1_chat_proto_rawDescOnce sync.Once
	file_pkg_api_v1_chat_proto_rawDescData = file_pkg_api_v1_chat_proto_rawDesc
)

func file_pkg_api_v1_chat_proto_rawDescGZIP() []byte {
	file_pkg_api_v1_chat_proto_rawDescOnce.Do(func() {
		file_pkg_api_v1_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_api_v1_chat_proto_rawDescData)
	})
	return file_pkg_api_v1_chat_proto_rawDescData
}

var file_pkg_api_v1_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_api_v1_chat_proto_goTypes = []interface{}{
	(*ChatRequest)(nil),  // 0: chat.v1.ChatRequest
	(*ChatResponse)(nil), // 1: chat.v1.ChatResponse
}
var file_pkg_api_v1_chat_proto_depIdxs = []int32{
	0, // 0: chat.v1.Chat.SendMessage:input_type -> chat.v1.ChatRequest
	0, // 1: chat.v1.Chat.StreamMessage:input_type -> chat.v1.ChatRequest
	1, // 2: chat.v1.Chat.SendMessage:output_type -> chat.v1.ChatResponse
	1, // 3: chat.v1.Chat.StreamMessage:output_type -> chat.v1.ChatResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_api_v1_chat_proto_init() }
func file_pkg_api_v1_chat_proto_init() {
	if File_pkg_api_v1_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_api_v1_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatRequest); i {
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
		file_pkg_api_v1_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatResponse); i {
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
			RawDescriptor: file_pkg_api_v1_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_api_v1_chat_proto_goTypes,
		DependencyIndexes: file_pkg_api_v1_chat_proto_depIdxs,
		MessageInfos:      file_pkg_api_v1_chat_proto_msgTypes,
	}.Build()
	File_pkg_api_v1_chat_proto = out.File
	file_pkg_api_v1_chat_proto_rawDesc = nil
	file_pkg_api_v1_chat_proto_goTypes = nil
	file_pkg_api_v1_chat_proto_depIdxs = nil
}