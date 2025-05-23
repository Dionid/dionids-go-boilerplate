// Copyright 2023-2025 Buf Technologies, Inc.
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
// source: buf/registry/plugin/v1beta1/download_service.proto

//go:build protoopaque

package pluginv1beta1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DownloadRequest struct {
	state             protoimpl.MessageState    `protogen:"opaque.v1"`
	xxx_hidden_Values *[]*DownloadRequest_Value `protobuf:"bytes,1,rep,name=values,proto3"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *DownloadRequest) Reset() {
	*x = DownloadRequest{}
	mi := &file_buf_registry_plugin_v1beta1_download_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadRequest) ProtoMessage() {}

func (x *DownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_plugin_v1beta1_download_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DownloadRequest) GetValues() []*DownloadRequest_Value {
	if x != nil {
		if x.xxx_hidden_Values != nil {
			return *x.xxx_hidden_Values
		}
	}
	return nil
}

func (x *DownloadRequest) SetValues(v []*DownloadRequest_Value) {
	x.xxx_hidden_Values = &v
}

type DownloadRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The references to get contents for.
	Values []*DownloadRequest_Value
}

func (b0 DownloadRequest_builder) Build() *DownloadRequest {
	m0 := &DownloadRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Values = &b.Values
	return m0
}

type DownloadResponse struct {
	state               protoimpl.MessageState       `protogen:"opaque.v1"`
	xxx_hidden_Contents *[]*DownloadResponse_Content `protobuf:"bytes,1,rep,name=contents,proto3"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *DownloadResponse) Reset() {
	*x = DownloadResponse{}
	mi := &file_buf_registry_plugin_v1beta1_download_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadResponse) ProtoMessage() {}

func (x *DownloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_plugin_v1beta1_download_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DownloadResponse) GetContents() []*DownloadResponse_Content {
	if x != nil {
		if x.xxx_hidden_Contents != nil {
			return *x.xxx_hidden_Contents
		}
	}
	return nil
}

func (x *DownloadResponse) SetContents(v []*DownloadResponse_Content) {
	x.xxx_hidden_Contents = &v
}

type DownloadResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Contents []*DownloadResponse_Content
}

func (b0 DownloadResponse_builder) Build() *DownloadResponse {
	m0 := &DownloadResponse{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Contents = &b.Contents
	return m0
}

// A request for content for a single version of a Plugin.
type DownloadRequest_Value struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_ResourceRef *ResourceRef           `protobuf:"bytes,1,opt,name=resource_ref,json=resourceRef,proto3"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *DownloadRequest_Value) Reset() {
	*x = DownloadRequest_Value{}
	mi := &file_buf_registry_plugin_v1beta1_download_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadRequest_Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadRequest_Value) ProtoMessage() {}

func (x *DownloadRequest_Value) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_plugin_v1beta1_download_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DownloadRequest_Value) GetResourceRef() *ResourceRef {
	if x != nil {
		return x.xxx_hidden_ResourceRef
	}
	return nil
}

func (x *DownloadRequest_Value) SetResourceRef(v *ResourceRef) {
	x.xxx_hidden_ResourceRef = v
}

func (x *DownloadRequest_Value) HasResourceRef() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_ResourceRef != nil
}

func (x *DownloadRequest_Value) ClearResourceRef() {
	x.xxx_hidden_ResourceRef = nil
}

type DownloadRequest_Value_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The reference to get content for.
	//
	// See the documentation on Reference for reference resolution details.
	//
	// Once the resource is resolved, the following content is returned:
	//   - If a Plugin is referenced, the content of the latest commit of the default label is
	//     returned.
	//   - If a Label is referenced, the content of the Commit of this Label is returned.
	//   - If a Commit is referenced, the content for this Commit is returned.
	ResourceRef *ResourceRef
}

func (b0 DownloadRequest_Value_builder) Build() *DownloadRequest_Value {
	m0 := &DownloadRequest_Value{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_ResourceRef = b.ResourceRef
	return m0
}

// Content for a single version of a Plugin.
type DownloadResponse_Content struct {
	state                      protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Commit          *Commit                `protobuf:"bytes,1,opt,name=commit,proto3"`
	xxx_hidden_CompressionType CompressionType        `protobuf:"varint,2,opt,name=compression_type,json=compressionType,proto3,enum=buf.registry.plugin.v1beta1.CompressionType"`
	xxx_hidden_Content         []byte                 `protobuf:"bytes,3,opt,name=content,proto3"`
	unknownFields              protoimpl.UnknownFields
	sizeCache                  protoimpl.SizeCache
}

func (x *DownloadResponse_Content) Reset() {
	*x = DownloadResponse_Content{}
	mi := &file_buf_registry_plugin_v1beta1_download_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadResponse_Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadResponse_Content) ProtoMessage() {}

func (x *DownloadResponse_Content) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_plugin_v1beta1_download_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DownloadResponse_Content) GetCommit() *Commit {
	if x != nil {
		return x.xxx_hidden_Commit
	}
	return nil
}

func (x *DownloadResponse_Content) GetCompressionType() CompressionType {
	if x != nil {
		return x.xxx_hidden_CompressionType
	}
	return CompressionType_COMPRESSION_TYPE_UNSPECIFIED
}

func (x *DownloadResponse_Content) GetContent() []byte {
	if x != nil {
		return x.xxx_hidden_Content
	}
	return nil
}

func (x *DownloadResponse_Content) SetCommit(v *Commit) {
	x.xxx_hidden_Commit = v
}

func (x *DownloadResponse_Content) SetCompressionType(v CompressionType) {
	x.xxx_hidden_CompressionType = v
}

func (x *DownloadResponse_Content) SetContent(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Content = v
}

func (x *DownloadResponse_Content) HasCommit() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Commit != nil
}

func (x *DownloadResponse_Content) ClearCommit() {
	x.xxx_hidden_Commit = nil
}

type DownloadResponse_Content_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The Commit associated with the Content.
	Commit *Commit
	// The compression type.
	CompressionType CompressionType
	// The content.
	Content []byte
}

func (b0 DownloadResponse_Content_builder) Build() *DownloadResponse_Content {
	m0 := &DownloadResponse_Content{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Commit = b.Commit
	x.xxx_hidden_CompressionType = b.CompressionType
	x.xxx_hidden_Content = b.Content
	return m0
}

var File_buf_registry_plugin_v1beta1_download_service_proto protoreflect.FileDescriptor

const file_buf_registry_plugin_v1beta1_download_service_proto_rawDesc = "" +
	"\n" +
	"2buf/registry/plugin/v1beta1/download_service.proto\x12\x1bbuf.registry.plugin.v1beta1\x1a(buf/registry/plugin/v1beta1/commit.proto\x1a-buf/registry/plugin/v1beta1/compression.proto\x1a*buf/registry/plugin/v1beta1/resource.proto\x1a\x1bbuf/validate/validate.proto\"\xc8\x01\n" +
	"\x0fDownloadRequest\x12W\n" +
	"\x06values\x18\x01 \x03(\v22.buf.registry.plugin.v1beta1.DownloadRequest.ValueB\v\xbaH\b\x92\x01\x05\b\x01\x10\xfa\x01R\x06values\x1a\\\n" +
	"\x05Value\x12S\n" +
	"\fresource_ref\x18\x01 \x01(\v2(.buf.registry.plugin.v1beta1.ResourceRefB\x06\xbaH\x03\xc8\x01\x01R\vresourceRef\"\xc8\x02\n" +
	"\x10DownloadResponse\x12[\n" +
	"\bcontents\x18\x01 \x03(\v25.buf.registry.plugin.v1beta1.DownloadResponse.ContentB\b\xbaH\x05\x92\x01\x02\b\x01R\bcontents\x1a\xd6\x01\n" +
	"\aContent\x12C\n" +
	"\x06commit\x18\x01 \x01(\v2#.buf.registry.plugin.v1beta1.CommitB\x06\xbaH\x03\xc8\x01\x01R\x06commit\x12d\n" +
	"\x10compression_type\x18\x02 \x01(\x0e2,.buf.registry.plugin.v1beta1.CompressionTypeB\v\xbaH\b\xc8\x01\x01\x82\x01\x02\x10\x01R\x0fcompressionType\x12 \n" +
	"\acontent\x18\x03 \x01(\fB\x06\xbaH\x03\xc8\x01\x01R\acontent2\x7f\n" +
	"\x0fDownloadService\x12l\n" +
	"\bDownload\x12,.buf.registry.plugin.v1beta1.DownloadRequest\x1a-.buf.registry.plugin.v1beta1.DownloadResponse\"\x03\x90\x02\x01BaZ_buf.build/gen/go/bufbuild/registry/protocolbuffers/go/buf/registry/plugin/v1beta1;pluginv1beta1b\x06proto3"

var file_buf_registry_plugin_v1beta1_download_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_buf_registry_plugin_v1beta1_download_service_proto_goTypes = []any{
	(*DownloadRequest)(nil),          // 0: buf.registry.plugin.v1beta1.DownloadRequest
	(*DownloadResponse)(nil),         // 1: buf.registry.plugin.v1beta1.DownloadResponse
	(*DownloadRequest_Value)(nil),    // 2: buf.registry.plugin.v1beta1.DownloadRequest.Value
	(*DownloadResponse_Content)(nil), // 3: buf.registry.plugin.v1beta1.DownloadResponse.Content
	(*ResourceRef)(nil),              // 4: buf.registry.plugin.v1beta1.ResourceRef
	(*Commit)(nil),                   // 5: buf.registry.plugin.v1beta1.Commit
	(CompressionType)(0),             // 6: buf.registry.plugin.v1beta1.CompressionType
}
var file_buf_registry_plugin_v1beta1_download_service_proto_depIdxs = []int32{
	2, // 0: buf.registry.plugin.v1beta1.DownloadRequest.values:type_name -> buf.registry.plugin.v1beta1.DownloadRequest.Value
	3, // 1: buf.registry.plugin.v1beta1.DownloadResponse.contents:type_name -> buf.registry.plugin.v1beta1.DownloadResponse.Content
	4, // 2: buf.registry.plugin.v1beta1.DownloadRequest.Value.resource_ref:type_name -> buf.registry.plugin.v1beta1.ResourceRef
	5, // 3: buf.registry.plugin.v1beta1.DownloadResponse.Content.commit:type_name -> buf.registry.plugin.v1beta1.Commit
	6, // 4: buf.registry.plugin.v1beta1.DownloadResponse.Content.compression_type:type_name -> buf.registry.plugin.v1beta1.CompressionType
	0, // 5: buf.registry.plugin.v1beta1.DownloadService.Download:input_type -> buf.registry.plugin.v1beta1.DownloadRequest
	1, // 6: buf.registry.plugin.v1beta1.DownloadService.Download:output_type -> buf.registry.plugin.v1beta1.DownloadResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_buf_registry_plugin_v1beta1_download_service_proto_init() }
func file_buf_registry_plugin_v1beta1_download_service_proto_init() {
	if File_buf_registry_plugin_v1beta1_download_service_proto != nil {
		return
	}
	file_buf_registry_plugin_v1beta1_commit_proto_init()
	file_buf_registry_plugin_v1beta1_compression_proto_init()
	file_buf_registry_plugin_v1beta1_resource_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_buf_registry_plugin_v1beta1_download_service_proto_rawDesc), len(file_buf_registry_plugin_v1beta1_download_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_buf_registry_plugin_v1beta1_download_service_proto_goTypes,
		DependencyIndexes: file_buf_registry_plugin_v1beta1_download_service_proto_depIdxs,
		MessageInfos:      file_buf_registry_plugin_v1beta1_download_service_proto_msgTypes,
	}.Build()
	File_buf_registry_plugin_v1beta1_download_service_proto = out.File
	file_buf_registry_plugin_v1beta1_download_service_proto_goTypes = nil
	file_buf_registry_plugin_v1beta1_download_service_proto_depIdxs = nil
}
