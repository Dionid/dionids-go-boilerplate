// Copyright 2024-2025 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: buf/plugin/descriptor/v1/file_descriptor.proto

//go:build protoopaque

package descriptorv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A file descriptor.
//
// A FileDescriptor is represented as a FileDescriptorProto, with the additional property of whether
// or not the File is an import.
type FileDescriptor struct {
	state                          protoimpl.MessageState            `protogen:"opaque.v1"`
	xxx_hidden_FileDescriptorProto *descriptorpb.FileDescriptorProto `protobuf:"bytes,1,opt,name=file_descriptor_proto,json=fileDescriptorProto,proto3"`
	xxx_hidden_IsImport            bool                              `protobuf:"varint,2,opt,name=is_import,json=isImport,proto3"`
	xxx_hidden_IsSyntaxUnspecified bool                              `protobuf:"varint,3,opt,name=is_syntax_unspecified,json=isSyntaxUnspecified,proto3"`
	xxx_hidden_UnusedDependency    []int32                           `protobuf:"varint,4,rep,packed,name=unused_dependency,json=unusedDependency,proto3"`
	unknownFields                  protoimpl.UnknownFields
	sizeCache                      protoimpl.SizeCache
}

func (x *FileDescriptor) Reset() {
	*x = FileDescriptor{}
	mi := &file_buf_plugin_descriptor_v1_file_descriptor_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileDescriptor) ProtoMessage() {}

func (x *FileDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_buf_plugin_descriptor_v1_file_descriptor_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *FileDescriptor) GetFileDescriptorProto() *descriptorpb.FileDescriptorProto {
	if x != nil {
		return x.xxx_hidden_FileDescriptorProto
	}
	return nil
}

func (x *FileDescriptor) GetIsImport() bool {
	if x != nil {
		return x.xxx_hidden_IsImport
	}
	return false
}

func (x *FileDescriptor) GetIsSyntaxUnspecified() bool {
	if x != nil {
		return x.xxx_hidden_IsSyntaxUnspecified
	}
	return false
}

func (x *FileDescriptor) GetUnusedDependency() []int32 {
	if x != nil {
		return x.xxx_hidden_UnusedDependency
	}
	return nil
}

func (x *FileDescriptor) SetFileDescriptorProto(v *descriptorpb.FileDescriptorProto) {
	x.xxx_hidden_FileDescriptorProto = v
}

func (x *FileDescriptor) SetIsImport(v bool) {
	x.xxx_hidden_IsImport = v
}

func (x *FileDescriptor) SetIsSyntaxUnspecified(v bool) {
	x.xxx_hidden_IsSyntaxUnspecified = v
}

func (x *FileDescriptor) SetUnusedDependency(v []int32) {
	x.xxx_hidden_UnusedDependency = v
}

func (x *FileDescriptor) HasFileDescriptorProto() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_FileDescriptorProto != nil
}

func (x *FileDescriptor) ClearFileDescriptorProto() {
	x.xxx_hidden_FileDescriptorProto = nil
}

type FileDescriptor_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The FileDescriptorProto that represents the file.
	//
	// Required.
	FileDescriptorProto *descriptorpb.FileDescriptorProto
	// Whether or not the file is considered an "import".
	//
	// An import is a file that is either:
	//
	//   - A Well-Known Type included from the compiler and imported by a targeted file.
	//   - A file that was included from a Buf module dependency and imported by a targeted file.
	//   - A file that was not targeted, but was imported by a targeted file.
	//
	// We use "import" as this matches with the protoc concept of --include_imports, however
	// import is a bit of an overloaded term.
	IsImport bool
	// Whether the file did not have a syntax explicitly specified.
	//
	// Per the FileDescriptorProto spec, it would be fine in this case to just leave the syntax field
	// unset to denote this and to set the syntax field to "proto2" if it is specified. However,
	// protoc does not set the syntax field if it was "proto2". Plugins may want to differentiate
	// between "proto2" and unset, and this field allows them to.
	IsSyntaxUnspecified bool
	// The indexes within the dependency field on FileDescriptorProto for those dependencies that
	// are not used.
	//
	// This matches the shape of the public_dependency and weak_dependency fields.
	UnusedDependency []int32
}

func (b0 FileDescriptor_builder) Build() *FileDescriptor {
	m0 := &FileDescriptor{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_FileDescriptorProto = b.FileDescriptorProto
	x.xxx_hidden_IsImport = b.IsImport
	x.xxx_hidden_IsSyntaxUnspecified = b.IsSyntaxUnspecified
	x.xxx_hidden_UnusedDependency = b.UnusedDependency
	return m0
}

var File_buf_plugin_descriptor_v1_file_descriptor_proto protoreflect.FileDescriptor

const file_buf_plugin_descriptor_v1_file_descriptor_proto_rawDesc = "" +
	"\n" +
	".buf/plugin/descriptor/v1/file_descriptor.proto\x12\x18buf.plugin.descriptor.v1\x1a\x1bbuf/validate/validate.proto\x1a google/protobuf/descriptor.proto\"\xa8\x02\n" +
	"\x0eFileDescriptor\x12\x97\x01\n" +
	"\x15file_descriptor_proto\x18\x01 \x01(\v2$.google.protobuf.FileDescriptorProtoB=\xbaH:\xba\x014\n" +
	"\fname_present\x12\x14name must be present\x1a\x0ehas(this.name)\xc8\x01\x01R\x13fileDescriptorProto\x12\x1b\n" +
	"\tis_import\x18\x02 \x01(\bR\bisImport\x122\n" +
	"\x15is_syntax_unspecified\x18\x03 \x01(\bR\x13isSyntaxUnspecified\x12+\n" +
	"\x11unused_dependency\x18\x04 \x03(\x05R\x10unusedDependencyB^Z\\buf.build/gen/go/bufbuild/bufplugin/protocolbuffers/go/buf/plugin/descriptor/v1;descriptorv1b\x06proto3"

var file_buf_plugin_descriptor_v1_file_descriptor_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_buf_plugin_descriptor_v1_file_descriptor_proto_goTypes = []any{
	(*FileDescriptor)(nil),                   // 0: buf.plugin.descriptor.v1.FileDescriptor
	(*descriptorpb.FileDescriptorProto)(nil), // 1: google.protobuf.FileDescriptorProto
}
var file_buf_plugin_descriptor_v1_file_descriptor_proto_depIdxs = []int32{
	1, // 0: buf.plugin.descriptor.v1.FileDescriptor.file_descriptor_proto:type_name -> google.protobuf.FileDescriptorProto
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_buf_plugin_descriptor_v1_file_descriptor_proto_init() }
func file_buf_plugin_descriptor_v1_file_descriptor_proto_init() {
	if File_buf_plugin_descriptor_v1_file_descriptor_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_buf_plugin_descriptor_v1_file_descriptor_proto_rawDesc), len(file_buf_plugin_descriptor_v1_file_descriptor_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_plugin_descriptor_v1_file_descriptor_proto_goTypes,
		DependencyIndexes: file_buf_plugin_descriptor_v1_file_descriptor_proto_depIdxs,
		MessageInfos:      file_buf_plugin_descriptor_v1_file_descriptor_proto_msgTypes,
	}.Build()
	File_buf_plugin_descriptor_v1_file_descriptor_proto = out.File
	file_buf_plugin_descriptor_v1_file_descriptor_proto_goTypes = nil
	file_buf_plugin_descriptor_v1_file_descriptor_proto_depIdxs = nil
}
