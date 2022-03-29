// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: providers/gcp/gcp.proto

package gcp

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

type ArtifactRegistryConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnableLocation []string `protobuf:"bytes,1,rep,name=enable_location,json=enableLocation,proto3" json:"enable_location,omitempty"`
}

func (x *ArtifactRegistryConf) Reset() {
	*x = ArtifactRegistryConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_providers_gcp_gcp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtifactRegistryConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtifactRegistryConf) ProtoMessage() {}

func (x *ArtifactRegistryConf) ProtoReflect() protoreflect.Message {
	mi := &file_providers_gcp_gcp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArtifactRegistryConf.ProtoReflect.Descriptor instead.
func (*ArtifactRegistryConf) Descriptor() ([]byte, []int) {
	return file_providers_gcp_gcp_proto_rawDescGZIP(), []int{0}
}

func (x *ArtifactRegistryConf) GetEnableLocation() []string {
	if x != nil {
		return x.EnableLocation
	}
	return nil
}

var File_providers_gcp_gcp_proto protoreflect.FileDescriptor

var file_providers_gcp_gcp_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x63, 0x70, 0x2f,
	0x67, 0x63, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x66, 0x6f, 0x75, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2e,
	0x67, 0x63, 0x70, 0x22, 0x3f, 0x0a, 0x14, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x27, 0x0a, 0x0f, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x2c, 0x5a, 0x2a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x67,
	0x63, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_providers_gcp_gcp_proto_rawDescOnce sync.Once
	file_providers_gcp_gcp_proto_rawDescData = file_providers_gcp_gcp_proto_rawDesc
)

func file_providers_gcp_gcp_proto_rawDescGZIP() []byte {
	file_providers_gcp_gcp_proto_rawDescOnce.Do(func() {
		file_providers_gcp_gcp_proto_rawDescData = protoimpl.X.CompressGZIP(file_providers_gcp_gcp_proto_rawDescData)
	})
	return file_providers_gcp_gcp_proto_rawDescData
}

var file_providers_gcp_gcp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_providers_gcp_gcp_proto_goTypes = []interface{}{
	(*ArtifactRegistryConf)(nil), // 0: foundation.providers.gcp.ArtifactRegistryConf
}
var file_providers_gcp_gcp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_providers_gcp_gcp_proto_init() }
func file_providers_gcp_gcp_proto_init() {
	if File_providers_gcp_gcp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_providers_gcp_gcp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArtifactRegistryConf); i {
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
			RawDescriptor: file_providers_gcp_gcp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_providers_gcp_gcp_proto_goTypes,
		DependencyIndexes: file_providers_gcp_gcp_proto_depIdxs,
		MessageInfos:      file_providers_gcp_gcp_proto_msgTypes,
	}.Build()
	File_providers_gcp_gcp_proto = out.File
	file_providers_gcp_gcp_proto_rawDesc = nil
	file_providers_gcp_gcp_proto_goTypes = nil
	file_providers_gcp_gcp_proto_depIdxs = nil
}