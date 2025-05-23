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
// source: buf/plugin/check/v1/category.proto

//go:build protoopaque

package checkv1

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

// A category that a CheckService implements.
//
// Buf uses categories to include or exclude sets of rules via configuration.
type Category struct {
	state                     protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Id             string                 `protobuf:"bytes,1,opt,name=id,proto3"`
	xxx_hidden_Purpose        string                 `protobuf:"bytes,2,opt,name=purpose,proto3"`
	xxx_hidden_Deprecated     bool                   `protobuf:"varint,3,opt,name=deprecated,proto3"`
	xxx_hidden_ReplacementIds []string               `protobuf:"bytes,4,rep,name=replacement_ids,json=replacementIds,proto3"`
	unknownFields             protoimpl.UnknownFields
	sizeCache                 protoimpl.SizeCache
}

func (x *Category) Reset() {
	*x = Category{}
	mi := &file_buf_plugin_check_v1_category_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_buf_plugin_check_v1_category_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Category) GetId() string {
	if x != nil {
		return x.xxx_hidden_Id
	}
	return ""
}

func (x *Category) GetPurpose() string {
	if x != nil {
		return x.xxx_hidden_Purpose
	}
	return ""
}

func (x *Category) GetDeprecated() bool {
	if x != nil {
		return x.xxx_hidden_Deprecated
	}
	return false
}

func (x *Category) GetReplacementIds() []string {
	if x != nil {
		return x.xxx_hidden_ReplacementIds
	}
	return nil
}

func (x *Category) SetId(v string) {
	x.xxx_hidden_Id = v
}

func (x *Category) SetPurpose(v string) {
	x.xxx_hidden_Purpose = v
}

func (x *Category) SetDeprecated(v bool) {
	x.xxx_hidden_Deprecated = v
}

func (x *Category) SetReplacementIds(v []string) {
	x.xxx_hidden_ReplacementIds = v
}

type Category_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The ID of the category.
	//
	// Required.
	//
	// This uniquely identifies the Category.
	//
	// Category IDs must also be unique relative to Rule IDs.
	//
	// Rule and Category IDs must be unique across all plugins used at the same time with
	// Buf. That is, no two plugins can both publish the same Rule or Category ID.
	//
	// This must have at least three characters.
	// This must start and end with a capital letter from A-Z or digits from 0-9, and only
	// consist of capital letters from A-Z, digits from 0-0, and underscores.
	Id string
	// A user-displayable purpose of the category.
	//
	// Required.
	//
	// This should be a proper sentence that starts with a capital letter and ends in a period.
	Purpose string
	// Whether or not this Category is deprecated.
	//
	// If the Category is deprecated, it may be replaced by 0 or more Categories. These will be
	// denoted by replacement_ids.
	Deprecated bool
	// The IDs of the Categories that replace this Category, if this Category is deprecated.
	//
	// This means that the combination of the Categories specified by replacement_ids replace this
	// Category entirely, and this Category is considered equivalent to the AND of the categories
	// specified by replacement_ids.
	//
	// This will only be non-empty if deprecated is true.
	// This may be empty even if deprecated is true.
	//
	// It is not valid for a deprecated Category to specfiy another deprecated Category as a
	// replacement.
	ReplacementIds []string
}

func (b0 Category_builder) Build() *Category {
	m0 := &Category{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Id = b.Id
	x.xxx_hidden_Purpose = b.Purpose
	x.xxx_hidden_Deprecated = b.Deprecated
	x.xxx_hidden_ReplacementIds = b.ReplacementIds
	return m0
}

var File_buf_plugin_check_v1_category_proto protoreflect.FileDescriptor

const file_buf_plugin_check_v1_category_proto_rawDesc = "" +
	"\n" +
	"\"buf/plugin/check/v1/category.proto\x12\x13buf.plugin.check.v1\x1a\x1bbuf/validate/validate.proto\"\xb0\x03\n" +
	"\bCategory\x12:\n" +
	"\x02id\x18\x01 \x01(\tB*\xbaH'\xc8\x01\x01r\"\x10\x03\x18@2\x1c^[A-Z0-9][A-Z0-9_]*[A-Z0-9]$R\x02id\x125\n" +
	"\apurpose\x18\x02 \x01(\tB\x1b\xbaH\x18\xc8\x01\x01r\x13\x10\x02\x18\x80\x022\f^[A-Z].*[.]$R\apurpose\x12\x1e\n" +
	"\n" +
	"deprecated\x18\x03 \x01(\bR\n" +
	"deprecated\x12U\n" +
	"\x0freplacement_ids\x18\x04 \x03(\tB,\xbaH)\x92\x01&\"$r\"\x10\x03\x18@2\x1c^[A-Z0-9][A-Z0-9_]*[A-Z0-9]$R\x0ereplacementIds:\xb9\x01\xbaH\xb5\x01\x1a\xb2\x01\n" +
	"*deprecated_true_if_replacement_ids_present\x126deprecated must be true if replacement_ids are present\x1aL!has(this.replacement_ids) || (has(this.replacement_ids) && this.deprecated)BTZRbuf.build/gen/go/bufbuild/bufplugin/protocolbuffers/go/buf/plugin/check/v1;checkv1b\x06proto3"

var file_buf_plugin_check_v1_category_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_buf_plugin_check_v1_category_proto_goTypes = []any{
	(*Category)(nil), // 0: buf.plugin.check.v1.Category
}
var file_buf_plugin_check_v1_category_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_buf_plugin_check_v1_category_proto_init() }
func file_buf_plugin_check_v1_category_proto_init() {
	if File_buf_plugin_check_v1_category_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_buf_plugin_check_v1_category_proto_rawDesc), len(file_buf_plugin_check_v1_category_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_plugin_check_v1_category_proto_goTypes,
		DependencyIndexes: file_buf_plugin_check_v1_category_proto_depIdxs,
		MessageInfos:      file_buf_plugin_check_v1_category_proto_msgTypes,
	}.Build()
	File_buf_plugin_check_v1_category_proto = out.File
	file_buf_plugin_check_v1_category_proto_goTypes = nil
	file_buf_plugin_check_v1_category_proto_depIdxs = nil
}
