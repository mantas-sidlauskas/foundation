// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: std/runtime/config.proto

package runtime

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

// The runtime configuration is generated from the deployment plan. It is
// injected into the server at runtime. Because RuntimeConfig is serialized
// as JSON, the field names have great significance and should be changed
// with care.
type RuntimeConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Environment *ServerEnvironment `protobuf:"bytes,1,opt,name=environment,proto3" json:"environment,omitempty"`
	Current     *Server            `protobuf:"bytes,2,opt,name=current,proto3" json:"current,omitempty"`
	// References to other servers that this server depends on.
	StackEntry []*Server `protobuf:"bytes,3,rep,name=stack_entry,json=stackEntry,proto3" json:"stack_entry,omitempty"`
}

func (x *RuntimeConfig) Reset() {
	*x = RuntimeConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_std_runtime_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuntimeConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuntimeConfig) ProtoMessage() {}

func (x *RuntimeConfig) ProtoReflect() protoreflect.Message {
	mi := &file_std_runtime_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuntimeConfig.ProtoReflect.Descriptor instead.
func (*RuntimeConfig) Descriptor() ([]byte, []int) {
	return file_std_runtime_config_proto_rawDescGZIP(), []int{0}
}

func (x *RuntimeConfig) GetEnvironment() *ServerEnvironment {
	if x != nil {
		return x.Environment
	}
	return nil
}

func (x *RuntimeConfig) GetCurrent() *Server {
	if x != nil {
		return x.Current
	}
	return nil
}

func (x *RuntimeConfig) GetStackEntry() []*Server {
	if x != nil {
		return x.StackEntry
	}
	return nil
}

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageName string         `protobuf:"bytes,1,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`
	ModuleName  string         `protobuf:"bytes,2,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty"`
	Port        []*Server_Port `protobuf:"bytes,3,rep,name=port,proto3" json:"port,omitempty"`
	Service     []*Service     `protobuf:"bytes,4,rep,name=service,proto3" json:"service,omitempty"`
	ImageRef    string         `protobuf:"bytes,5,opt,name=image_ref,json=imageRef,proto3" json:"image_ref,omitempty"` // Only set for current.
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_std_runtime_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_std_runtime_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_std_runtime_config_proto_rawDescGZIP(), []int{1}
}

func (x *Server) GetPackageName() string {
	if x != nil {
		return x.PackageName
	}
	return ""
}

func (x *Server) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

func (x *Server) GetPort() []*Server_Port {
	if x != nil {
		return x.Port
	}
	return nil
}

func (x *Server) GetService() []*Service {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *Server) GetImageRef() string {
	if x != nil {
		return x.ImageRef
	}
	return ""
}

type ServerEnvironment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // Empty if ephemeral is true.
	Purpose string `protobuf:"bytes,2,opt,name=purpose,proto3" json:"purpose,omitempty"`
	// Typically only set for tests. Signals that this environment is single-use and not meant to be user serviceable.
	Ephemeral bool `protobuf:"varint,3,opt,name=ephemeral,proto3" json:"ephemeral,omitempty"`
}

func (x *ServerEnvironment) Reset() {
	*x = ServerEnvironment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_std_runtime_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerEnvironment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerEnvironment) ProtoMessage() {}

func (x *ServerEnvironment) ProtoReflect() protoreflect.Message {
	mi := &file_std_runtime_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerEnvironment.ProtoReflect.Descriptor instead.
func (*ServerEnvironment) Descriptor() ([]byte, []int) {
	return file_std_runtime_config_proto_rawDescGZIP(), []int{2}
}

func (x *ServerEnvironment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServerEnvironment) GetPurpose() string {
	if x != nil {
		return x.Purpose
	}
	return ""
}

func (x *ServerEnvironment) GetEphemeral() bool {
	if x != nil {
		return x.Ephemeral
	}
	return false
}

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner    string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`       // Package name.
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`         // Scoped to package name.
	Endpoint string `protobuf:"bytes,3,opt,name=endpoint,proto3" json:"endpoint,omitempty"` // E.g. hostname:port.
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_std_runtime_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_std_runtime_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_std_runtime_config_proto_rawDescGZIP(), []int{3}
}

func (x *Service) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Service) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type Server_Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *Server_Port) Reset() {
	*x = Server_Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_std_runtime_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_Port) ProtoMessage() {}

func (x *Server_Port) ProtoReflect() protoreflect.Message {
	mi := &file_std_runtime_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_Port.ProtoReflect.Descriptor instead.
func (*Server_Port) Descriptor() ([]byte, []int) {
	return file_std_runtime_config_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Server_Port) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Server_Port) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_std_runtime_config_proto protoreflect.FileDescriptor

var file_std_runtime_config_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x74, 0x64, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x74, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x22, 0xd7, 0x01, 0x0a, 0x0d, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x4b, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x74, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x38, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x73, 0x74, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x3f, 0x0a, 0x0b, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x74,
	0x64, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x8d, 0x02, 0x0a,
	0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x74, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x39, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x73, 0x74, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x66, 0x1a, 0x2e, 0x0a, 0x04,
	0x50, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x5f, 0x0a, 0x11,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x65, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x65, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x22, 0x4f, 0x0a,
	0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x2a,
	0x5a, 0x28, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x2e,
	0x64, 0x65, 0x76, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73,
	0x74, 0x64, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_std_runtime_config_proto_rawDescOnce sync.Once
	file_std_runtime_config_proto_rawDescData = file_std_runtime_config_proto_rawDesc
)

func file_std_runtime_config_proto_rawDescGZIP() []byte {
	file_std_runtime_config_proto_rawDescOnce.Do(func() {
		file_std_runtime_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_std_runtime_config_proto_rawDescData)
	})
	return file_std_runtime_config_proto_rawDescData
}

var file_std_runtime_config_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_std_runtime_config_proto_goTypes = []interface{}{
	(*RuntimeConfig)(nil),     // 0: foundation.std.runtime.RuntimeConfig
	(*Server)(nil),            // 1: foundation.std.runtime.Server
	(*ServerEnvironment)(nil), // 2: foundation.std.runtime.ServerEnvironment
	(*Service)(nil),           // 3: foundation.std.runtime.Service
	(*Server_Port)(nil),       // 4: foundation.std.runtime.Server.Port
}
var file_std_runtime_config_proto_depIdxs = []int32{
	2, // 0: foundation.std.runtime.RuntimeConfig.environment:type_name -> foundation.std.runtime.ServerEnvironment
	1, // 1: foundation.std.runtime.RuntimeConfig.current:type_name -> foundation.std.runtime.Server
	1, // 2: foundation.std.runtime.RuntimeConfig.stack_entry:type_name -> foundation.std.runtime.Server
	4, // 3: foundation.std.runtime.Server.port:type_name -> foundation.std.runtime.Server.Port
	3, // 4: foundation.std.runtime.Server.service:type_name -> foundation.std.runtime.Service
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_std_runtime_config_proto_init() }
func file_std_runtime_config_proto_init() {
	if File_std_runtime_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_std_runtime_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuntimeConfig); i {
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
		file_std_runtime_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
		file_std_runtime_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerEnvironment); i {
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
		file_std_runtime_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_std_runtime_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_Port); i {
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
			RawDescriptor: file_std_runtime_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_std_runtime_config_proto_goTypes,
		DependencyIndexes: file_std_runtime_config_proto_depIdxs,
		MessageInfos:      file_std_runtime_config_proto_msgTypes,
	}.Build()
	File_std_runtime_config_proto = out.File
	file_std_runtime_config_proto_rawDesc = nil
	file_std_runtime_config_proto_goTypes = nil
	file_std_runtime_config_proto_depIdxs = nil
}
