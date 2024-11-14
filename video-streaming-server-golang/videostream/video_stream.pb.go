// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: video_stream.proto

package videostream

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

type VideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *VideoRequest) Reset() {
	*x = VideoRequest{}
	mi := &file_video_stream_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoRequest) ProtoMessage() {}

func (x *VideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_stream_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoRequest.ProtoReflect.Descriptor instead.
func (*VideoRequest) Descriptor() ([]byte, []int) {
	return file_video_stream_proto_rawDescGZIP(), []int{0}
}

func (x *VideoRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type VideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *VideoResponse) Reset() {
	*x = VideoResponse{}
	mi := &file_video_stream_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoResponse) ProtoMessage() {}

func (x *VideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_video_stream_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoResponse.ProtoReflect.Descriptor instead.
func (*VideoResponse) Descriptor() ([]byte, []int) {
	return file_video_stream_proto_rawDescGZIP(), []int{1}
}

func (x *VideoResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_video_stream_proto protoreflect.FileDescriptor

var file_video_stream_proto_rawDesc = []byte{
	0x0a, 0x12, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x22, 0x2a, 0x0a, 0x0c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x23, 0x0a,
	0x0d, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x32, 0x55, 0x0a, 0x0b, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x46, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x12, 0x19, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x2b, 0x5a, 0x29, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2d, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_video_stream_proto_rawDescOnce sync.Once
	file_video_stream_proto_rawDescData = file_video_stream_proto_rawDesc
)

func file_video_stream_proto_rawDescGZIP() []byte {
	file_video_stream_proto_rawDescOnce.Do(func() {
		file_video_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_video_stream_proto_rawDescData)
	})
	return file_video_stream_proto_rawDescData
}

var file_video_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_video_stream_proto_goTypes = []any{
	(*VideoRequest)(nil),  // 0: videostream.VideoRequest
	(*VideoResponse)(nil), // 1: videostream.VideoResponse
}
var file_video_stream_proto_depIdxs = []int32{
	0, // 0: videostream.VideoStream.StreamVideo:input_type -> videostream.VideoRequest
	1, // 1: videostream.VideoStream.StreamVideo:output_type -> videostream.VideoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_video_stream_proto_init() }
func file_video_stream_proto_init() {
	if File_video_stream_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_video_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_video_stream_proto_goTypes,
		DependencyIndexes: file_video_stream_proto_depIdxs,
		MessageInfos:      file_video_stream_proto_msgTypes,
	}.Build()
	File_video_stream_proto = out.File
	file_video_stream_proto_rawDesc = nil
	file_video_stream_proto_goTypes = nil
	file_video_stream_proto_depIdxs = nil
}
