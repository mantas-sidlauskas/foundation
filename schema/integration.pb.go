// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: schema/integration.proto

package schema

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NodejsIntegration_NodePkgMgr int32

const (
	NodejsIntegration_PKG_MGR_UNKNOWN NodejsIntegration_NodePkgMgr = 0
	NodejsIntegration_NPM             NodejsIntegration_NodePkgMgr = 1
	NodejsIntegration_YARN            NodejsIntegration_NodePkgMgr = 2
	NodejsIntegration_PNPM            NodejsIntegration_NodePkgMgr = 3
)

// Enum value maps for NodejsIntegration_NodePkgMgr.
var (
	NodejsIntegration_NodePkgMgr_name = map[int32]string{
		0: "PKG_MGR_UNKNOWN",
		1: "NPM",
		2: "YARN",
		3: "PNPM",
	}
	NodejsIntegration_NodePkgMgr_value = map[string]int32{
		"PKG_MGR_UNKNOWN": 0,
		"NPM":             1,
		"YARN":            2,
		"PNPM":            3,
	}
)

func (x NodejsIntegration_NodePkgMgr) Enum() *NodejsIntegration_NodePkgMgr {
	p := new(NodejsIntegration_NodePkgMgr)
	*p = x
	return p
}

func (x NodejsIntegration_NodePkgMgr) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodejsIntegration_NodePkgMgr) Descriptor() protoreflect.EnumDescriptor {
	return file_schema_integration_proto_enumTypes[0].Descriptor()
}

func (NodejsIntegration_NodePkgMgr) Type() protoreflect.EnumType {
	return &file_schema_integration_proto_enumTypes[0]
}

func (x NodejsIntegration_NodePkgMgr) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodejsIntegration_NodePkgMgr.Descriptor instead.
func (NodejsIntegration_NodePkgMgr) EnumDescriptor() ([]byte, []int) {
	return file_schema_integration_proto_rawDescGZIP(), []int{3, 0}
}

type Integration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Contains the integration-specific configuration, see below.
	Data *anypb.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Integration) Reset() {
	*x = Integration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_integration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Integration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Integration) ProtoMessage() {}

func (x *Integration) ProtoReflect() protoreflect.Message {
	mi := &file_schema_integration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Integration.ProtoReflect.Descriptor instead.
func (*Integration) Descriptor() ([]byte, []int) {
	return file_schema_integration_proto_rawDescGZIP(), []int{0}
}

func (x *Integration) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type DockerIntegration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dockerfile string `protobuf:"bytes,1,opt,name=dockerfile,proto3" json:"dockerfile,omitempty"`
}

func (x *DockerIntegration) Reset() {
	*x = DockerIntegration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_integration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DockerIntegration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DockerIntegration) ProtoMessage() {}

func (x *DockerIntegration) ProtoReflect() protoreflect.Message {
	mi := &file_schema_integration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DockerIntegration.ProtoReflect.Descriptor instead.
func (*DockerIntegration) Descriptor() ([]byte, []int) {
	return file_schema_integration_proto_rawDescGZIP(), []int{1}
}

func (x *DockerIntegration) GetDockerfile() string {
	if x != nil {
		return x.Dockerfile
	}
	return ""
}

type GoIntegration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pkg string `protobuf:"bytes,1,opt,name=pkg,proto3" json:"pkg,omitempty"`
}

func (x *GoIntegration) Reset() {
	*x = GoIntegration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_integration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoIntegration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoIntegration) ProtoMessage() {}

func (x *GoIntegration) ProtoReflect() protoreflect.Message {
	mi := &file_schema_integration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoIntegration.ProtoReflect.Descriptor instead.
func (*GoIntegration) Descriptor() ([]byte, []int) {
	return file_schema_integration_proto_rawDescGZIP(), []int{2}
}

func (x *GoIntegration) GetPkg() string {
	if x != nil {
		return x.Pkg
	}
	return ""
}

type NodejsIntegration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path to `package.json`, relative to the Namespace package. Default is "."
	Pkg string `protobuf:"bytes,1,opt,name=pkg,proto3" json:"pkg,omitempty"`
	// Detected Node.js package manager.
	NodePkgMgr NodejsIntegration_NodePkgMgr `protobuf:"varint,2,opt,name=node_pkg_mgr,json=nodePkgMgr,proto3,enum=foundation.schema.NodejsIntegration_NodePkgMgr" json:"node_pkg_mgr,omitempty"`
}

func (x *NodejsIntegration) Reset() {
	*x = NodejsIntegration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_integration_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodejsIntegration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodejsIntegration) ProtoMessage() {}

func (x *NodejsIntegration) ProtoReflect() protoreflect.Message {
	mi := &file_schema_integration_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodejsIntegration.ProtoReflect.Descriptor instead.
func (*NodejsIntegration) Descriptor() ([]byte, []int) {
	return file_schema_integration_proto_rawDescGZIP(), []int{3}
}

func (x *NodejsIntegration) GetPkg() string {
	if x != nil {
		return x.Pkg
	}
	return ""
}

func (x *NodejsIntegration) GetNodePkgMgr() NodejsIntegration_NodePkgMgr {
	if x != nil {
		return x.NodePkgMgr
	}
	return NodejsIntegration_PKG_MGR_UNKNOWN
}

var File_schema_integration_proto protoreflect.FileDescriptor

var file_schema_integration_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x0b, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x33, 0x0a, 0x11, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x21, 0x0a, 0x0d, 0x47, 0x6f, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6b, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x6b, 0x67, 0x22, 0xb8, 0x01, 0x0a, 0x11, 0x4e, 0x6f,
	0x64, 0x65, 0x6a, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x70, 0x6b, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x6b,
	0x67, 0x12, 0x51, 0x0a, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x70, 0x6b, 0x67, 0x5f, 0x6d, 0x67,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x6a, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x50, 0x6b, 0x67, 0x4d, 0x67, 0x72, 0x52, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x50, 0x6b,
	0x67, 0x4d, 0x67, 0x72, 0x22, 0x3e, 0x0a, 0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x6b, 0x67, 0x4d,
	0x67, 0x72, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x4b, 0x47, 0x5f, 0x4d, 0x47, 0x52, 0x5f, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x50, 0x4d, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x59, 0x41, 0x52, 0x4e, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4e,
	0x50, 0x4d, 0x10, 0x03, 0x42, 0x25, 0x5a, 0x23, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_schema_integration_proto_rawDescOnce sync.Once
	file_schema_integration_proto_rawDescData = file_schema_integration_proto_rawDesc
)

func file_schema_integration_proto_rawDescGZIP() []byte {
	file_schema_integration_proto_rawDescOnce.Do(func() {
		file_schema_integration_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_integration_proto_rawDescData)
	})
	return file_schema_integration_proto_rawDescData
}

var file_schema_integration_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_schema_integration_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_schema_integration_proto_goTypes = []interface{}{
	(NodejsIntegration_NodePkgMgr)(0), // 0: foundation.schema.NodejsIntegration.NodePkgMgr
	(*Integration)(nil),               // 1: foundation.schema.Integration
	(*DockerIntegration)(nil),         // 2: foundation.schema.DockerIntegration
	(*GoIntegration)(nil),             // 3: foundation.schema.GoIntegration
	(*NodejsIntegration)(nil),         // 4: foundation.schema.NodejsIntegration
	(*anypb.Any)(nil),                 // 5: google.protobuf.Any
}
var file_schema_integration_proto_depIdxs = []int32{
	5, // 0: foundation.schema.Integration.data:type_name -> google.protobuf.Any
	0, // 1: foundation.schema.NodejsIntegration.node_pkg_mgr:type_name -> foundation.schema.NodejsIntegration.NodePkgMgr
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_schema_integration_proto_init() }
func file_schema_integration_proto_init() {
	if File_schema_integration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_integration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Integration); i {
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
		file_schema_integration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DockerIntegration); i {
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
		file_schema_integration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoIntegration); i {
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
		file_schema_integration_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodejsIntegration); i {
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
			RawDescriptor: file_schema_integration_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_integration_proto_goTypes,
		DependencyIndexes: file_schema_integration_proto_depIdxs,
		EnumInfos:         file_schema_integration_proto_enumTypes,
		MessageInfos:      file_schema_integration_proto_msgTypes,
	}.Build()
	File_schema_integration_proto = out.File
	file_schema_integration_proto_rawDesc = nil
	file_schema_integration_proto_goTypes = nil
	file_schema_integration_proto_depIdxs = nil
}
