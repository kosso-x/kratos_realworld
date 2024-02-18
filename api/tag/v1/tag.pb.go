// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: tag/v1/tag.proto

package v1

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

type GetTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTagRequest) Reset() {
	*x = GetTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_v1_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagRequest) ProtoMessage() {}

func (x *GetTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tag_v1_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagRequest.ProtoReflect.Descriptor instead.
func (*GetTagRequest) Descriptor() ([]byte, []int) {
	return file_tag_v1_tag_proto_rawDescGZIP(), []int{0}
}

type GetTagReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags []string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *GetTagReply) Reset() {
	*x = GetTagReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_v1_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagReply) ProtoMessage() {}

func (x *GetTagReply) ProtoReflect() protoreflect.Message {
	mi := &file_tag_v1_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagReply.ProtoReflect.Descriptor instead.
func (*GetTagReply) Descriptor() ([]byte, []int) {
	return file_tag_v1_tag_proto_rawDescGZIP(), []int{1}
}

func (x *GetTagReply) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_tag_v1_tag_proto protoreflect.FileDescriptor

var file_tag_v1_tag_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x61, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0f, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x21, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x32, 0x57, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x50, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x61,
	0x67, 0x73, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61,
	0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x67, 0x73, 0x42, 0x27, 0x0a, 0x0a, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x61, 0x67, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x17, 0x72, 0x65, 0x61, 0x6c, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tag_v1_tag_proto_rawDescOnce sync.Once
	file_tag_v1_tag_proto_rawDescData = file_tag_v1_tag_proto_rawDesc
)

func file_tag_v1_tag_proto_rawDescGZIP() []byte {
	file_tag_v1_tag_proto_rawDescOnce.Do(func() {
		file_tag_v1_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_tag_v1_tag_proto_rawDescData)
	})
	return file_tag_v1_tag_proto_rawDescData
}

var file_tag_v1_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tag_v1_tag_proto_goTypes = []interface{}{
	(*GetTagRequest)(nil), // 0: api.tag.v1.GetTagRequest
	(*GetTagReply)(nil),   // 1: api.tag.v1.GetTagReply
}
var file_tag_v1_tag_proto_depIdxs = []int32{
	0, // 0: api.tag.v1.Tag.GetTags:input_type -> api.tag.v1.GetTagRequest
	1, // 1: api.tag.v1.Tag.GetTags:output_type -> api.tag.v1.GetTagReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tag_v1_tag_proto_init() }
func file_tag_v1_tag_proto_init() {
	if File_tag_v1_tag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tag_v1_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagRequest); i {
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
		file_tag_v1_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagReply); i {
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
			RawDescriptor: file_tag_v1_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tag_v1_tag_proto_goTypes,
		DependencyIndexes: file_tag_v1_tag_proto_depIdxs,
		MessageInfos:      file_tag_v1_tag_proto_msgTypes,
	}.Build()
	File_tag_v1_tag_proto = out.File
	file_tag_v1_tag_proto_rawDesc = nil
	file_tag_v1_tag_proto_goTypes = nil
	file_tag_v1_tag_proto_depIdxs = nil
}
