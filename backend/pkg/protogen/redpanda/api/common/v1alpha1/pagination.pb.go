// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: redpanda/api/common/v1alpha1/pagination.proto

package commonv1alpha1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// KeySetPageToken represents a pagination token for KeySet pagination.
// It marks the beginning of a page where records start from the key that
// satisfies the condition key >= value_greater_equal. Records are sorted
// alphabetically by key in ascending order.
type KeySetPageToken struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	Key               string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	ValueGreaterEqual string                 `protobuf:"bytes,2,opt,name=value_greater_equal,json=valueGreaterEqual,proto3" json:"value_greater_equal,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *KeySetPageToken) Reset() {
	*x = KeySetPageToken{}
	mi := &file_redpanda_api_common_v1alpha1_pagination_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KeySetPageToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeySetPageToken) ProtoMessage() {}

func (x *KeySetPageToken) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_api_common_v1alpha1_pagination_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeySetPageToken.ProtoReflect.Descriptor instead.
func (*KeySetPageToken) Descriptor() ([]byte, []int) {
	return file_redpanda_api_common_v1alpha1_pagination_proto_rawDescGZIP(), []int{0}
}

func (x *KeySetPageToken) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeySetPageToken) GetValueGreaterEqual() string {
	if x != nil {
		return x.ValueGreaterEqual
	}
	return ""
}

var File_redpanda_api_common_v1alpha1_pagination_proto protoreflect.FileDescriptor

var file_redpanda_api_common_v1alpha1_pagination_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1c, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x53, 0x0a,
	0x0f, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x2e, 0x0a, 0x13, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x67, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x72, 0x5f, 0x65, 0x71, 0x75, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x47, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x45, 0x71, 0x75,
	0x61, 0x6c, 0x42, 0xab, 0x02, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61,
	0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0f, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x63, 0x62, 0x75, 0x66, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65, 0x64,
	0x70, 0x61, 0x6e, 0x64, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73,
	0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2,
	0x02, 0x03, 0x52, 0x41, 0x43, 0xaa, 0x02, 0x1c, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61,
	0x2e, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x1c, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x5c,
	0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xe2, 0x02, 0x28, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x5c, 0x41,
	0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x1f, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_redpanda_api_common_v1alpha1_pagination_proto_rawDescOnce sync.Once
	file_redpanda_api_common_v1alpha1_pagination_proto_rawDescData = file_redpanda_api_common_v1alpha1_pagination_proto_rawDesc
)

func file_redpanda_api_common_v1alpha1_pagination_proto_rawDescGZIP() []byte {
	file_redpanda_api_common_v1alpha1_pagination_proto_rawDescOnce.Do(func() {
		file_redpanda_api_common_v1alpha1_pagination_proto_rawDescData = protoimpl.X.CompressGZIP(file_redpanda_api_common_v1alpha1_pagination_proto_rawDescData)
	})
	return file_redpanda_api_common_v1alpha1_pagination_proto_rawDescData
}

var file_redpanda_api_common_v1alpha1_pagination_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_redpanda_api_common_v1alpha1_pagination_proto_goTypes = []any{
	(*KeySetPageToken)(nil), // 0: redpanda.api.common.v1alpha1.KeySetPageToken
}
var file_redpanda_api_common_v1alpha1_pagination_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_redpanda_api_common_v1alpha1_pagination_proto_init() }
func file_redpanda_api_common_v1alpha1_pagination_proto_init() {
	if File_redpanda_api_common_v1alpha1_pagination_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_redpanda_api_common_v1alpha1_pagination_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_redpanda_api_common_v1alpha1_pagination_proto_goTypes,
		DependencyIndexes: file_redpanda_api_common_v1alpha1_pagination_proto_depIdxs,
		MessageInfos:      file_redpanda_api_common_v1alpha1_pagination_proto_msgTypes,
	}.Build()
	File_redpanda_api_common_v1alpha1_pagination_proto = out.File
	file_redpanda_api_common_v1alpha1_pagination_proto_rawDesc = nil
	file_redpanda_api_common_v1alpha1_pagination_proto_goTypes = nil
	file_redpanda_api_common_v1alpha1_pagination_proto_depIdxs = nil
}
