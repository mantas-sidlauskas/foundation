// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: internal/fnfs/workspace/wsremote/protocol.proto

package wsremote

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wscontents "namespacelabs.dev/foundation/internal/wscontents"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModuleName string `protobuf:"bytes,1,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty"`
	Rel        string `protobuf:"bytes,2,opt,name=rel,proto3" json:"rel,omitempty"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_fnfs_workspace_wsremote_protocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_internal_fnfs_workspace_wsremote_protocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_internal_fnfs_workspace_wsremote_protocol_proto_rawDescGZIP(), []int{0}
}

func (x *Signature) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

func (x *Signature) GetRel() string {
	if x != nil {
		return x.Rel
	}
	return ""
}

type PushRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature *Signature              `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	FileEvent []*wscontents.FileEvent `protobuf:"bytes,2,rep,name=file_event,json=fileEvent,proto3" json:"file_event,omitempty"`
}

func (x *PushRequest) Reset() {
	*x = PushRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_fnfs_workspace_wsremote_protocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushRequest) ProtoMessage() {}

func (x *PushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_fnfs_workspace_wsremote_protocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushRequest.ProtoReflect.Descriptor instead.
func (*PushRequest) Descriptor() ([]byte, []int) {
	return file_internal_fnfs_workspace_wsremote_protocol_proto_rawDescGZIP(), []int{1}
}

func (x *PushRequest) GetSignature() *Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *PushRequest) GetFileEvent() []*wscontents.FileEvent {
	if x != nil {
		return x.FileEvent
	}
	return nil
}

type PushResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PushResponse) Reset() {
	*x = PushResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_fnfs_workspace_wsremote_protocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushResponse) ProtoMessage() {}

func (x *PushResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_fnfs_workspace_wsremote_protocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushResponse.ProtoReflect.Descriptor instead.
func (*PushResponse) Descriptor() ([]byte, []int) {
	return file_internal_fnfs_workspace_wsremote_protocol_proto_rawDescGZIP(), []int{2}
}

var File_internal_fnfs_workspace_wsremote_protocol_proto protoreflect.FileDescriptor

var file_internal_fnfs_workspace_wsremote_protocol_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x66, 0x6e, 0x66, 0x73, 0x2f,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2f, 0x77, 0x73, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1c, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x77, 0x73, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x1a,
	0x1f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x77, 0x73, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3e, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x72, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x6c,
	0x22, 0x9e, 0x01, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x45, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x77, 0x73, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x48, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x77, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0x70, 0x0a, 0x0f, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x04, 0x50, 0x75, 0x73, 0x68, 0x12, 0x29, 0x2e, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x77, 0x73, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x50, 0x75, 0x73, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x77, 0x73,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x3f, 0x5a, 0x3d, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6c, 0x61, 0x62, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x66, 0x6e, 0x66,
	0x73, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2f, 0x77, 0x73, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_fnfs_workspace_wsremote_protocol_proto_rawDescOnce sync.Once
	file_internal_fnfs_workspace_wsremote_protocol_proto_rawDescData = file_internal_fnfs_workspace_wsremote_protocol_proto_rawDesc
)

func file_internal_fnfs_workspace_wsremote_protocol_proto_rawDescGZIP() []byte {
	file_internal_fnfs_workspace_wsremote_protocol_proto_rawDescOnce.Do(func() {
		file_internal_fnfs_workspace_wsremote_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_fnfs_workspace_wsremote_protocol_proto_rawDescData)
	})
	return file_internal_fnfs_workspace_wsremote_protocol_proto_rawDescData
}

var file_internal_fnfs_workspace_wsremote_protocol_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_fnfs_workspace_wsremote_protocol_proto_goTypes = []interface{}{
	(*Signature)(nil),            // 0: foundation.internal.wsremote.Signature
	(*PushRequest)(nil),          // 1: foundation.internal.wsremote.PushRequest
	(*PushResponse)(nil),         // 2: foundation.internal.wsremote.PushResponse
	(*wscontents.FileEvent)(nil), // 3: foundation.internal.wscontents.FileEvent
}
var file_internal_fnfs_workspace_wsremote_protocol_proto_depIdxs = []int32{
	0, // 0: foundation.internal.wsremote.PushRequest.signature:type_name -> foundation.internal.wsremote.Signature
	3, // 1: foundation.internal.wsremote.PushRequest.file_event:type_name -> foundation.internal.wscontents.FileEvent
	1, // 2: foundation.internal.wsremote.FileSyncService.Push:input_type -> foundation.internal.wsremote.PushRequest
	2, // 3: foundation.internal.wsremote.FileSyncService.Push:output_type -> foundation.internal.wsremote.PushResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_fnfs_workspace_wsremote_protocol_proto_init() }
func file_internal_fnfs_workspace_wsremote_protocol_proto_init() {
	if File_internal_fnfs_workspace_wsremote_protocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_fnfs_workspace_wsremote_protocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signature); i {
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
		file_internal_fnfs_workspace_wsremote_protocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushRequest); i {
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
		file_internal_fnfs_workspace_wsremote_protocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushResponse); i {
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
			RawDescriptor: file_internal_fnfs_workspace_wsremote_protocol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_fnfs_workspace_wsremote_protocol_proto_goTypes,
		DependencyIndexes: file_internal_fnfs_workspace_wsremote_protocol_proto_depIdxs,
		MessageInfos:      file_internal_fnfs_workspace_wsremote_protocol_proto_msgTypes,
	}.Build()
	File_internal_fnfs_workspace_wsremote_protocol_proto = out.File
	file_internal_fnfs_workspace_wsremote_protocol_proto_rawDesc = nil
	file_internal_fnfs_workspace_wsremote_protocol_proto_goTypes = nil
	file_internal_fnfs_workspace_wsremote_protocol_proto_depIdxs = nil
}