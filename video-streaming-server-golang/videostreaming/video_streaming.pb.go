// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: video_streaming.proto

package videostreaming

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

	VideoName string `protobuf:"bytes,1,opt,name=video_name,json=videoName,proto3" json:"video_name,omitempty"`
}

func (x *VideoRequest) Reset() {
	*x = VideoRequest{}
	mi := &file_video_streaming_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoRequest) ProtoMessage() {}

func (x *VideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_streaming_proto_msgTypes[0]
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
	return file_video_streaming_proto_rawDescGZIP(), []int{0}
}

func (x *VideoRequest) GetVideoName() string {
	if x != nil {
		return x.VideoName
	}
	return ""
}

type VideoChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *VideoChunk) Reset() {
	*x = VideoChunk{}
	mi := &file_video_streaming_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VideoChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoChunk) ProtoMessage() {}

func (x *VideoChunk) ProtoReflect() protoreflect.Message {
	mi := &file_video_streaming_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoChunk.ProtoReflect.Descriptor instead.
func (*VideoChunk) Descriptor() ([]byte, []int) {
	return file_video_streaming_proto_rawDescGZIP(), []int{1}
}

func (x *VideoChunk) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_video_streaming_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_video_streaming_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_video_streaming_proto_rawDescGZIP(), []int{2}
}

type VideoList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoNames []string `protobuf:"bytes,1,rep,name=video_names,json=videoNames,proto3" json:"video_names,omitempty"`
}

func (x *VideoList) Reset() {
	*x = VideoList{}
	mi := &file_video_streaming_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VideoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoList) ProtoMessage() {}

func (x *VideoList) ProtoReflect() protoreflect.Message {
	mi := &file_video_streaming_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoList.ProtoReflect.Descriptor instead.
func (*VideoList) Descriptor() ([]byte, []int) {
	return file_video_streaming_proto_rawDescGZIP(), []int{3}
}

func (x *VideoList) GetVideoNames() []string {
	if x != nil {
		return x.VideoNames
	}
	return nil
}

var File_video_streaming_proto protoreflect.FileDescriptor

var file_video_streaming_proto_rawDesc = []byte{
	0x0a, 0x15, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x22, 0x2d, 0x0a, 0x0c, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x07,
	0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2c, 0x0a, 0x09, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x32, 0xa6, 0x01, 0x0a, 0x15, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4b, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x1c,
	0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x00, 0x30, 0x01, 0x12, 0x40, 0x0a, 0x0a,
	0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x12, 0x15, 0x2e, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x19, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x2e,
	0x5a, 0x2c, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2d, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_video_streaming_proto_rawDescOnce sync.Once
	file_video_streaming_proto_rawDescData = file_video_streaming_proto_rawDesc
)

func file_video_streaming_proto_rawDescGZIP() []byte {
	file_video_streaming_proto_rawDescOnce.Do(func() {
		file_video_streaming_proto_rawDescData = protoimpl.X.CompressGZIP(file_video_streaming_proto_rawDescData)
	})
	return file_video_streaming_proto_rawDescData
}

var file_video_streaming_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_video_streaming_proto_goTypes = []any{
	(*VideoRequest)(nil), // 0: videostreaming.VideoRequest
	(*VideoChunk)(nil),   // 1: videostreaming.VideoChunk
	(*Empty)(nil),        // 2: videostreaming.Empty
	(*VideoList)(nil),    // 3: videostreaming.VideoList
}
var file_video_streaming_proto_depIdxs = []int32{
	0, // 0: videostreaming.VideoStreamingService.StreamVideo:input_type -> videostreaming.VideoRequest
	2, // 1: videostreaming.VideoStreamingService.ListVideos:input_type -> videostreaming.Empty
	1, // 2: videostreaming.VideoStreamingService.StreamVideo:output_type -> videostreaming.VideoChunk
	3, // 3: videostreaming.VideoStreamingService.ListVideos:output_type -> videostreaming.VideoList
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_video_streaming_proto_init() }
func file_video_streaming_proto_init() {
	if File_video_streaming_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_video_streaming_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_video_streaming_proto_goTypes,
		DependencyIndexes: file_video_streaming_proto_depIdxs,
		MessageInfos:      file_video_streaming_proto_msgTypes,
	}.Build()
	File_video_streaming_proto = out.File
	file_video_streaming_proto_rawDesc = nil
	file_video_streaming_proto_goTypes = nil
	file_video_streaming_proto_depIdxs = nil
}
