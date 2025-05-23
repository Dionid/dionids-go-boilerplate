// Copyright 2020-2025 Buf Technologies, Inc.
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
// source: buf/alpha/registry/v1alpha1/verification_status.proto

package registryv1alpha1

import (
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

// VerificationStatus is the verification status of an owner on if we recognize them,
// an owner can be either user or organization.
type VerificationStatus int32

const (
	VerificationStatus_VERIFICATION_STATUS_UNSPECIFIED VerificationStatus = 0
	// OFFICIAL indicates that the owner is maintained by Buf.
	VerificationStatus_VERIFICATION_STATUS_OFFICIAL VerificationStatus = 1
	// VERIFIED_PUBLISHER indicates that the owner is a third-party that has been
	// verified by Buf.
	VerificationStatus_VERIFICATION_STATUS_VERIFIED_PUBLISHER VerificationStatus = 2
)

// Enum value maps for VerificationStatus.
var (
	VerificationStatus_name = map[int32]string{
		0: "VERIFICATION_STATUS_UNSPECIFIED",
		1: "VERIFICATION_STATUS_OFFICIAL",
		2: "VERIFICATION_STATUS_VERIFIED_PUBLISHER",
	}
	VerificationStatus_value = map[string]int32{
		"VERIFICATION_STATUS_UNSPECIFIED":        0,
		"VERIFICATION_STATUS_OFFICIAL":           1,
		"VERIFICATION_STATUS_VERIFIED_PUBLISHER": 2,
	}
)

func (x VerificationStatus) Enum() *VerificationStatus {
	p := new(VerificationStatus)
	*p = x
	return p
}

func (x VerificationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VerificationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_buf_alpha_registry_v1alpha1_verification_status_proto_enumTypes[0].Descriptor()
}

func (VerificationStatus) Type() protoreflect.EnumType {
	return &file_buf_alpha_registry_v1alpha1_verification_status_proto_enumTypes[0]
}

func (x VerificationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

var File_buf_alpha_registry_v1alpha1_verification_status_proto protoreflect.FileDescriptor

const file_buf_alpha_registry_v1alpha1_verification_status_proto_rawDesc = "" +
	"\n" +
	"5buf/alpha/registry/v1alpha1/verification_status.proto\x12\x1bbuf.alpha.registry.v1alpha1*\x87\x01\n" +
	"\x12VerificationStatus\x12#\n" +
	"\x1fVERIFICATION_STATUS_UNSPECIFIED\x10\x00\x12 \n" +
	"\x1cVERIFICATION_STATUS_OFFICIAL\x10\x01\x12*\n" +
	"&VERIFICATION_STATUS_VERIFIED_PUBLISHER\x10\x02B\xa4\x02\n" +
	"\x1fcom.buf.alpha.registry.v1alpha1B\x17VerificationStatusProtoP\x01ZYgithub.com/bufbuild/buf/private/gen/proto/go/buf/alpha/registry/v1alpha1;registryv1alpha1\xa2\x02\x03BAR\xaa\x02\x1bBuf.Alpha.Registry.V1alpha1\xca\x02\x1bBuf\\Alpha\\Registry\\V1alpha1\xe2\x02'Buf\\Alpha\\Registry\\V1alpha1\\GPBMetadata\xea\x02\x1eBuf::Alpha::Registry::V1alpha1b\x06proto3"

var file_buf_alpha_registry_v1alpha1_verification_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_buf_alpha_registry_v1alpha1_verification_status_proto_goTypes = []any{
	(VerificationStatus)(0), // 0: buf.alpha.registry.v1alpha1.VerificationStatus
}
var file_buf_alpha_registry_v1alpha1_verification_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_buf_alpha_registry_v1alpha1_verification_status_proto_init() }
func file_buf_alpha_registry_v1alpha1_verification_status_proto_init() {
	if File_buf_alpha_registry_v1alpha1_verification_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_buf_alpha_registry_v1alpha1_verification_status_proto_rawDesc), len(file_buf_alpha_registry_v1alpha1_verification_status_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_alpha_registry_v1alpha1_verification_status_proto_goTypes,
		DependencyIndexes: file_buf_alpha_registry_v1alpha1_verification_status_proto_depIdxs,
		EnumInfos:         file_buf_alpha_registry_v1alpha1_verification_status_proto_enumTypes,
	}.Build()
	File_buf_alpha_registry_v1alpha1_verification_status_proto = out.File
	file_buf_alpha_registry_v1alpha1_verification_status_proto_goTypes = nil
	file_buf_alpha_registry_v1alpha1_verification_status_proto_depIdxs = nil
}
