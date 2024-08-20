// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd/protoc-gen-go/testdata/protoeditions/legacy_enum.proto

package protoeditions

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/gofeaturespb"
	reflect "reflect"
	sync "sync"
)

// EnumTypeWithLegacyUnmarshalJSON comment.
type EnumTypeWithLegacyUnmarshalJSON int32

const (
	// EnumTypeWithLegacyUnmarshalJSON_ONE comment.
	EnumTypeWithLegacyUnmarshalJSON_FIRST EnumTypeWithLegacyUnmarshalJSON = 1
	// EnumTypeWithLegacyUnmarshalJSON_TWO comment.
	EnumTypeWithLegacyUnmarshalJSON_SECOND EnumTypeWithLegacyUnmarshalJSON = 2
)

// Enum value maps for EnumTypeWithLegacyUnmarshalJSON.
var (
	EnumTypeWithLegacyUnmarshalJSON_name = map[int32]string{
		1: "FIRST",
		2: "SECOND",
	}
	EnumTypeWithLegacyUnmarshalJSON_value = map[string]int32{
		"FIRST":  1,
		"SECOND": 2,
	}
)

func (x EnumTypeWithLegacyUnmarshalJSON) Enum() *EnumTypeWithLegacyUnmarshalJSON {
	p := new(EnumTypeWithLegacyUnmarshalJSON)
	*p = x
	return p
}

func (x EnumTypeWithLegacyUnmarshalJSON) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumTypeWithLegacyUnmarshalJSON) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_enumTypes[0].Descriptor()
}

func (EnumTypeWithLegacyUnmarshalJSON) Type() protoreflect.EnumType {
	return &file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_enumTypes[0]
}

func (x EnumTypeWithLegacyUnmarshalJSON) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EnumTypeWithLegacyUnmarshalJSON) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EnumTypeWithLegacyUnmarshalJSON(num)
	return nil
}

// Deprecated: Use EnumTypeWithLegacyUnmarshalJSON.Descriptor instead.
func (EnumTypeWithLegacyUnmarshalJSON) EnumDescriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDescGZIP(), []int{0}
}

type EnumWithoutUnmarshalJSON int32

const (
	EnumWithoutUnmarshalJSON_WITHOUT_UNMARSHAL_JSON_FOO EnumWithoutUnmarshalJSON = 0
	EnumWithoutUnmarshalJSON_WITHOUT_UNMARSHAL_JSON_BAR EnumWithoutUnmarshalJSON = 1
	EnumWithoutUnmarshalJSON_WITHOUT_UNMARSHAL_JSON_BAZ EnumWithoutUnmarshalJSON = 2
)

// Enum value maps for EnumWithoutUnmarshalJSON.
var (
	EnumWithoutUnmarshalJSON_name = map[int32]string{
		0: "WITHOUT_UNMARSHAL_JSON_FOO",
		1: "WITHOUT_UNMARSHAL_JSON_BAR",
		2: "WITHOUT_UNMARSHAL_JSON_BAZ",
	}
	EnumWithoutUnmarshalJSON_value = map[string]int32{
		"WITHOUT_UNMARSHAL_JSON_FOO": 0,
		"WITHOUT_UNMARSHAL_JSON_BAR": 1,
		"WITHOUT_UNMARSHAL_JSON_BAZ": 2,
	}
)

func (x EnumWithoutUnmarshalJSON) Enum() *EnumWithoutUnmarshalJSON {
	p := new(EnumWithoutUnmarshalJSON)
	*p = x
	return p
}

func (x EnumWithoutUnmarshalJSON) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumWithoutUnmarshalJSON) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_enumTypes[1].Descriptor()
}

func (EnumWithoutUnmarshalJSON) Type() protoreflect.EnumType {
	return &file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_enumTypes[1]
}

func (x EnumWithoutUnmarshalJSON) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumWithoutUnmarshalJSON.Descriptor instead.
func (EnumWithoutUnmarshalJSON) EnumDescriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDescGZIP(), []int{1}
}

// NestedEnumType1A comment.
type ContainerForNestedEnum_NestedEnum int32

const (
	// NestedEnum_VALUE comment.
	ContainerForNestedEnum_VALUE ContainerForNestedEnum_NestedEnum = 0
)

// Enum value maps for ContainerForNestedEnum_NestedEnum.
var (
	ContainerForNestedEnum_NestedEnum_name = map[int32]string{
		0: "VALUE",
	}
	ContainerForNestedEnum_NestedEnum_value = map[string]int32{
		"VALUE": 0,
	}
)

func (x ContainerForNestedEnum_NestedEnum) Enum() *ContainerForNestedEnum_NestedEnum {
	p := new(ContainerForNestedEnum_NestedEnum)
	*p = x
	return p
}

func (x ContainerForNestedEnum_NestedEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContainerForNestedEnum_NestedEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_enumTypes[2].Descriptor()
}

func (ContainerForNestedEnum_NestedEnum) Type() protoreflect.EnumType {
	return &file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_enumTypes[2]
}

func (x ContainerForNestedEnum_NestedEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ContainerForNestedEnum_NestedEnum) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ContainerForNestedEnum_NestedEnum(num)
	return nil
}

// Deprecated: Use ContainerForNestedEnum_NestedEnum.Descriptor instead.
func (ContainerForNestedEnum_NestedEnum) EnumDescriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDescGZIP(), []int{0, 0}
}

type ContainerForNestedEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ContainerForNestedEnum) Reset() {
	*x = ContainerForNestedEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerForNestedEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerForNestedEnum) ProtoMessage() {}

func (x *ContainerForNestedEnum) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerForNestedEnum.ProtoReflect.Descriptor instead.
func (*ContainerForNestedEnum) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDescGZIP(), []int{0}
}

var File_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto protoreflect.FileDescriptor

var file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63,
	0x79, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x21, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x5f, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a,
	0x16, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x4e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x17, 0x0a, 0x0a, 0x4e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x00,
	0x2a, 0x3e, 0x0a, 0x1f, 0x45, 0x6e, 0x75, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x57, 0x69, 0x74, 0x68,
	0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x55, 0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x4a,
	0x53, 0x4f, 0x4e, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x49, 0x52, 0x53, 0x54, 0x10, 0x01, 0x12, 0x0a,
	0x0a, 0x06, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x10, 0x02, 0x1a, 0x04, 0x3a, 0x02, 0x10, 0x02,
	0x2a, 0x83, 0x01, 0x0a, 0x18, 0x45, 0x6e, 0x75, 0x6d, 0x57, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74,
	0x55, 0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x1e, 0x0a,
	0x1a, 0x57, 0x49, 0x54, 0x48, 0x4f, 0x55, 0x54, 0x5f, 0x55, 0x4e, 0x4d, 0x41, 0x52, 0x53, 0x48,
	0x41, 0x4c, 0x5f, 0x4a, 0x53, 0x4f, 0x4e, 0x5f, 0x46, 0x4f, 0x4f, 0x10, 0x00, 0x12, 0x1e, 0x0a,
	0x1a, 0x57, 0x49, 0x54, 0x48, 0x4f, 0x55, 0x54, 0x5f, 0x55, 0x4e, 0x4d, 0x41, 0x52, 0x53, 0x48,
	0x41, 0x4c, 0x5f, 0x4a, 0x53, 0x4f, 0x4e, 0x5f, 0x42, 0x41, 0x52, 0x10, 0x01, 0x12, 0x1e, 0x0a,
	0x1a, 0x57, 0x49, 0x54, 0x48, 0x4f, 0x55, 0x54, 0x5f, 0x55, 0x4e, 0x4d, 0x41, 0x52, 0x53, 0x48,
	0x41, 0x4c, 0x5f, 0x4a, 0x53, 0x4f, 0x4e, 0x5f, 0x42, 0x41, 0x5a, 0x10, 0x02, 0x1a, 0x07, 0x3a,
	0x05, 0xd2, 0x3e, 0x02, 0x08, 0x00, 0x42, 0x4d, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x92, 0x03, 0x05,
	0xd2, 0x3e, 0x02, 0x08, 0x01, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70,
	0xe8, 0x07,
}

var (
	file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDescData = file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDesc
)

func file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDescData)
	})
	return file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDescData
}

var file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_goTypes = []any{
	(EnumTypeWithLegacyUnmarshalJSON)(0),   // 0: goproto.protoc.protoeditions.EnumTypeWithLegacyUnmarshalJSON
	(EnumWithoutUnmarshalJSON)(0),          // 1: goproto.protoc.protoeditions.EnumWithoutUnmarshalJSON
	(ContainerForNestedEnum_NestedEnum)(0), // 2: goproto.protoc.protoeditions.ContainerForNestedEnum.NestedEnum
	(*ContainerForNestedEnum)(nil),         // 3: goproto.protoc.protoeditions.ContainerForNestedEnum
}
var file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_init() }
func file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_init() {
	if File_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ContainerForNestedEnum); i {
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
			RawDescriptor: file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_depIdxs,
		EnumInfos:         file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_enumTypes,
		MessageInfos:      file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_msgTypes,
	}.Build()
	File_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto = out.File
	file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDesc = nil
	file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_goTypes = nil
	file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_depIdxs = nil
}
