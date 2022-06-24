// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: schema/test_storage.proto

package schema

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TestResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success                bool                    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Plan                   *DeployPlan             `protobuf:"bytes,2,opt,name=plan,proto3" json:"plan,omitempty"`
	ComputedConfigurations *ComputedConfigurations `protobuf:"bytes,3,opt,name=computed_configurations,json=computedConfigurations,proto3" json:"computed_configurations,omitempty"`
}

func (x *TestResult) Reset() {
	*x = TestResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_test_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResult) ProtoMessage() {}

func (x *TestResult) ProtoReflect() protoreflect.Message {
	mi := &file_schema_test_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResult.ProtoReflect.Descriptor instead.
func (*TestResult) Descriptor() ([]byte, []int) {
	return file_schema_test_storage_proto_rawDescGZIP(), []int{0}
}

func (x *TestResult) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *TestResult) GetPlan() *DeployPlan {
	if x != nil {
		return x.Plan
	}
	return nil
}

func (x *TestResult) GetComputedConfigurations() *ComputedConfigurations {
	if x != nil {
		return x.ComputedConfigurations
	}
	return nil
}

type TestBundle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Created   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created,proto3" json:"created,omitempty"`
	Completed *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=completed,proto3" json:"completed,omitempty"` // Regardless of success or failure.
	Result    *TestResult            `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	TestLog   *LogRef                `protobuf:"bytes,2,opt,name=test_log,json=testLog,proto3" json:"test_log,omitempty"`
	ServerLog []*LogRef              `protobuf:"bytes,3,rep,name=server_log,json=serverLog,proto3" json:"server_log,omitempty"`
}

func (x *TestBundle) Reset() {
	*x = TestBundle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_test_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestBundle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestBundle) ProtoMessage() {}

func (x *TestBundle) ProtoReflect() protoreflect.Message {
	mi := &file_schema_test_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestBundle.ProtoReflect.Descriptor instead.
func (*TestBundle) Descriptor() ([]byte, []int) {
	return file_schema_test_storage_proto_rawDescGZIP(), []int{1}
}

func (x *TestBundle) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *TestBundle) GetCompleted() *timestamppb.Timestamp {
	if x != nil {
		return x.Completed
	}
	return nil
}

func (x *TestBundle) GetResult() *TestResult {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *TestBundle) GetTestLog() *LogRef {
	if x != nil {
		return x.TestLog
	}
	return nil
}

func (x *TestBundle) GetServerLog() []*LogRef {
	if x != nil {
		return x.ServerLog
	}
	return nil
}

type LogRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageName   string        `protobuf:"bytes,1,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`
	ContainerName string        `protobuf:"bytes,2,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty"`
	ContainerKind ContainerKind `protobuf:"varint,4,opt,name=container_kind,json=containerKind,proto3,enum=foundation.schema.ContainerKind" json:"container_kind,omitempty"`
	LogFile       string        `protobuf:"bytes,3,opt,name=log_file,json=logFile,proto3" json:"log_file,omitempty"`
	LogSize       uint64        `protobuf:"varint,5,opt,name=log_size,json=logSize,proto3" json:"log_size,omitempty"`
}

func (x *LogRef) Reset() {
	*x = LogRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_test_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRef) ProtoMessage() {}

func (x *LogRef) ProtoReflect() protoreflect.Message {
	mi := &file_schema_test_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRef.ProtoReflect.Descriptor instead.
func (*LogRef) Descriptor() ([]byte, []int) {
	return file_schema_test_storage_proto_rawDescGZIP(), []int{2}
}

func (x *LogRef) GetPackageName() string {
	if x != nil {
		return x.PackageName
	}
	return ""
}

func (x *LogRef) GetContainerName() string {
	if x != nil {
		return x.ContainerName
	}
	return ""
}

func (x *LogRef) GetContainerKind() ContainerKind {
	if x != nil {
		return x.ContainerKind
	}
	return ContainerKind_CONTAINER_KIND_UNSPECIFIED
}

func (x *LogRef) GetLogFile() string {
	if x != nil {
		return x.LogFile
	}
	return ""
}

func (x *LogRef) GetLogSize() uint64 {
	if x != nil {
		return x.LogSize
	}
	return 0
}

var File_schema_test_storage_proto protoreflect.FileDescriptor

var file_schema_test_storage_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a, 0x17,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x14, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x31, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x04, 0x70,
	0x6c, 0x61, 0x6e, 0x12, 0x62, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x16, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xa3, 0x02, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74,
	0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x35, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x34, 0x0a,
	0x08, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x66, 0x52, 0x07, 0x74, 0x65, 0x73, 0x74,
	0x4c, 0x6f, 0x67, 0x12, 0x38, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6c, 0x6f,
	0x67, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4c, 0x6f, 0x67, 0x52,
	0x65, 0x66, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x22, 0xd1, 0x01,
	0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x66, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f,
	0x6b, 0x69, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x0d, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c,
	0x6f, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c,
	0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x53, 0x69, 0x7a,
	0x65, 0x42, 0x25, 0x5a, 0x23, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61,
	0x62, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_test_storage_proto_rawDescOnce sync.Once
	file_schema_test_storage_proto_rawDescData = file_schema_test_storage_proto_rawDesc
)

func file_schema_test_storage_proto_rawDescGZIP() []byte {
	file_schema_test_storage_proto_rawDescOnce.Do(func() {
		file_schema_test_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_test_storage_proto_rawDescData)
	})
	return file_schema_test_storage_proto_rawDescData
}

var file_schema_test_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_schema_test_storage_proto_goTypes = []interface{}{
	(*TestResult)(nil),             // 0: foundation.schema.TestResult
	(*TestBundle)(nil),             // 1: foundation.schema.TestBundle
	(*LogRef)(nil),                 // 2: foundation.schema.LogRef
	(*DeployPlan)(nil),             // 3: foundation.schema.DeployPlan
	(*ComputedConfigurations)(nil), // 4: foundation.schema.ComputedConfigurations
	(*timestamppb.Timestamp)(nil),  // 5: google.protobuf.Timestamp
	(ContainerKind)(0),             // 6: foundation.schema.ContainerKind
}
var file_schema_test_storage_proto_depIdxs = []int32{
	3, // 0: foundation.schema.TestResult.plan:type_name -> foundation.schema.DeployPlan
	4, // 1: foundation.schema.TestResult.computed_configurations:type_name -> foundation.schema.ComputedConfigurations
	5, // 2: foundation.schema.TestBundle.created:type_name -> google.protobuf.Timestamp
	5, // 3: foundation.schema.TestBundle.completed:type_name -> google.protobuf.Timestamp
	0, // 4: foundation.schema.TestBundle.result:type_name -> foundation.schema.TestResult
	2, // 5: foundation.schema.TestBundle.test_log:type_name -> foundation.schema.LogRef
	2, // 6: foundation.schema.TestBundle.server_log:type_name -> foundation.schema.LogRef
	6, // 7: foundation.schema.LogRef.container_kind:type_name -> foundation.schema.ContainerKind
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_schema_test_storage_proto_init() }
func file_schema_test_storage_proto_init() {
	if File_schema_test_storage_proto != nil {
		return
	}
	file_schema_definition_proto_init()
	file_schema_serialized_proto_init()
	file_schema_runtime_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_schema_test_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResult); i {
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
		file_schema_test_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestBundle); i {
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
		file_schema_test_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogRef); i {
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
			RawDescriptor: file_schema_test_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_test_storage_proto_goTypes,
		DependencyIndexes: file_schema_test_storage_proto_depIdxs,
		MessageInfos:      file_schema_test_storage_proto_msgTypes,
	}.Build()
	File_schema_test_storage_proto = out.File
	file_schema_test_storage_proto_rawDesc = nil
	file_schema_test_storage_proto_goTypes = nil
	file_schema_test_storage_proto_depIdxs = nil
}