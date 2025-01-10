// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test Protobuf definitions with proto3 syntax.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/textpb3/test.proto

package textpb3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type Enum int32

const (
	Enum_ZERO Enum = 0
	Enum_ONE  Enum = 1
	Enum_TWO  Enum = 2
	Enum_TEN  Enum = 10
)

// Enum value maps for Enum.
var (
	Enum_name = map[int32]string{
		0:  "ZERO",
		1:  "ONE",
		2:  "TWO",
		10: "TEN",
	}
	Enum_value = map[string]int32{
		"ZERO": 0,
		"ONE":  1,
		"TWO":  2,
		"TEN":  10,
	}
)

func (x Enum) Enum() *Enum {
	p := new(Enum)
	*p = x
	return p
}

func (x Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_testprotos_textpb3_test_proto_enumTypes[0].Descriptor()
}

func (Enum) Type() protoreflect.EnumType {
	return &file_internal_testprotos_textpb3_test_proto_enumTypes[0]
}

func (x Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Enum.Descriptor instead.
func (Enum) EnumDescriptor() ([]byte, []int) {
	return file_internal_testprotos_textpb3_test_proto_rawDescGZIP(), []int{0}
}

type OptionalEnums_NestedEnum int32

const (
	OptionalEnums_CERO OptionalEnums_NestedEnum = 0
	OptionalEnums_UNO  OptionalEnums_NestedEnum = 1
	OptionalEnums_DOS  OptionalEnums_NestedEnum = 2
	OptionalEnums_DIEZ OptionalEnums_NestedEnum = 10
)

// Enum value maps for OptionalEnums_NestedEnum.
var (
	OptionalEnums_NestedEnum_name = map[int32]string{
		0:  "CERO",
		1:  "UNO",
		2:  "DOS",
		10: "DIEZ",
	}
	OptionalEnums_NestedEnum_value = map[string]int32{
		"CERO": 0,
		"UNO":  1,
		"DOS":  2,
		"DIEZ": 10,
	}
)

func (x OptionalEnums_NestedEnum) Enum() *OptionalEnums_NestedEnum {
	p := new(OptionalEnums_NestedEnum)
	*p = x
	return p
}

func (x OptionalEnums_NestedEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OptionalEnums_NestedEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_testprotos_textpb3_test_proto_enumTypes[1].Descriptor()
}

func (OptionalEnums_NestedEnum) Type() protoreflect.EnumType {
	return &file_internal_testprotos_textpb3_test_proto_enumTypes[1]
}

func (x OptionalEnums_NestedEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OptionalEnums_NestedEnum.Descriptor instead.
func (OptionalEnums_NestedEnum) EnumDescriptor() ([]byte, []int) {
	return file_internal_testprotos_textpb3_test_proto_rawDescGZIP(), []int{3, 0}
}

type Enums_NestedEnum int32

const (
	Enums_CERO Enums_NestedEnum = 0
	Enums_UNO  Enums_NestedEnum = 1
	Enums_DOS  Enums_NestedEnum = 2
	Enums_DIEZ Enums_NestedEnum = 10
)

// Enum value maps for Enums_NestedEnum.
var (
	Enums_NestedEnum_name = map[int32]string{
		0:  "CERO",
		1:  "UNO",
		2:  "DOS",
		10: "DIEZ",
	}
	Enums_NestedEnum_value = map[string]int32{
		"CERO": 0,
		"UNO":  1,
		"DOS":  2,
		"DIEZ": 10,
	}
)

func (x Enums_NestedEnum) Enum() *Enums_NestedEnum {
	p := new(Enums_NestedEnum)
	*p = x
	return p
}

func (x Enums_NestedEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enums_NestedEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_testprotos_textpb3_test_proto_enumTypes[2].Descriptor()
}

func (Enums_NestedEnum) Type() protoreflect.EnumType {
	return &file_internal_testprotos_textpb3_test_proto_enumTypes[2]
}

func (x Enums_NestedEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Enums_NestedEnum.Descriptor instead.
func (Enums_NestedEnum) EnumDescriptor() ([]byte, []int) {
	return file_internal_testprotos_textpb3_test_proto_rawDescGZIP(), []int{4, 0}
}

// Scalars contains scalar field types.
type Scalars struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SBool         bool                   `protobuf:"varint,1,opt,name=s_bool,json=sBool,proto3" json:"s_bool,omitempty"`
	SInt32        int32                  `protobuf:"varint,2,opt,name=s_int32,json=sInt32,proto3" json:"s_int32,omitempty"`
	SInt64        int64                  `protobuf:"varint,3,opt,name=s_int64,json=sInt64,proto3" json:"s_int64,omitempty"`
	SUint32       uint32                 `protobuf:"varint,4,opt,name=s_uint32,json=sUint32,proto3" json:"s_uint32,omitempty"`
	SUint64       uint64                 `protobuf:"varint,5,opt,name=s_uint64,json=sUint64,proto3" json:"s_uint64,omitempty"`
	SSint32       int32                  `protobuf:"zigzag32,6,opt,name=s_sint32,json=sSint32,proto3" json:"s_sint32,omitempty"`
	SSint64       int64                  `protobuf:"zigzag64,7,opt,name=s_sint64,json=sSint64,proto3" json:"s_sint64,omitempty"`
	SFixed32      uint32                 `protobuf:"fixed32,8,opt,name=s_fixed32,json=sFixed32,proto3" json:"s_fixed32,omitempty"`
	SFixed64      uint64                 `protobuf:"fixed64,9,opt,name=s_fixed64,json=sFixed64,proto3" json:"s_fixed64,omitempty"`
	SSfixed32     int32                  `protobuf:"fixed32,10,opt,name=s_sfixed32,json=sSfixed32,proto3" json:"s_sfixed32,omitempty"`
	SSfixed64     int64                  `protobuf:"fixed64,11,opt,name=s_sfixed64,json=sSfixed64,proto3" json:"s_sfixed64,omitempty"`
	SFloat        float32                `protobuf:"fixed32,20,opt,name=s_float,json=sFloat,proto3" json:"s_float,omitempty"`
	SDouble       float64                `protobuf:"fixed64,21,opt,name=s_double,json=sDouble,proto3" json:"s_double,omitempty"`
	SBytes        []byte                 `protobuf:"bytes,14,opt,name=s_bytes,json=sBytes,proto3" json:"s_bytes,omitempty"`
	SString       string                 `protobuf:"bytes,13,opt,name=s_string,json=sString,proto3" json:"s_string,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Scalars) Reset() {
	*x = Scalars{}
	mi := &file_internal_testprotos_textpb3_test_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Scalars) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scalars) ProtoMessage() {}

func (x *Scalars) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_textpb3_test_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scalars.ProtoReflect.Descriptor instead.
func (*Scalars) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_textpb3_test_proto_rawDescGZIP(), []int{0}
}

func (x *Scalars) GetSBool() bool {
	if x != nil {
		return x.SBool
	}
	return false
}

func (x *Scalars) GetSInt32() int32 {
	if x != nil {
		return x.SInt32
	}
	return 0
}

func (x *Scalars) GetSInt64() int64 {
	if x != nil {
		return x.SInt64
	}
	return 0
}

func (x *Scalars) GetSUint32() uint32 {
	if x != nil {
		return x.SUint32
	}
	return 0
}

func (x *Scalars) GetSUint64() uint64 {
	if x != nil {
		return x.SUint64
	}
	return 0
}

func (x *Scalars) GetSSint32() int32 {
	if x != nil {
		return x.SSint32
	}
	return 0
}

func (x *Scalars) GetSSint64() int64 {
	if x != nil {
		return x.SSint64
	}
	return 0
}

func (x *Scalars) GetSFixed32() uint32 {
	if x != nil {
		return x.SFixed32
	}
	return 0
}

func (x *Scalars) GetSFixed64() uint64 {
	if x != nil {
		return x.SFixed64
	}
	return 0
}

func (x *Scalars) GetSSfixed32() int32 {
	if x != nil {
		return x.SSfixed32
	}
	return 0
}

func (x *Scalars) GetSSfixed64() int64 {
	if x != nil {
		return x.SSfixed64
	}
	return 0
}

func (x *Scalars) GetSFloat() float32 {
	if x != nil {
		return x.SFloat
	}
	return 0
}

func (x *Scalars) GetSDouble() float64 {
	if x != nil {
		return x.SDouble
	}
	return 0
}

func (x *Scalars) GetSBytes() []byte {
	if x != nil {
		return x.SBytes
	}
	return nil
}

func (x *Scalars) GetSString() string {
	if x != nil {
		return x.SString
	}
	return ""
}

// Message contains repeated fields.
type Repeats struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RptBool       []bool                 `protobuf:"varint,1,rep,packed,name=rpt_bool,json=rptBool,proto3" json:"rpt_bool,omitempty"`
	RptInt32      []int32                `protobuf:"varint,2,rep,packed,name=rpt_int32,json=rptInt32,proto3" json:"rpt_int32,omitempty"`
	RptInt64      []int64                `protobuf:"varint,3,rep,packed,name=rpt_int64,json=rptInt64,proto3" json:"rpt_int64,omitempty"`
	RptUint32     []uint32               `protobuf:"varint,4,rep,packed,name=rpt_uint32,json=rptUint32,proto3" json:"rpt_uint32,omitempty"`
	RptUint64     []uint64               `protobuf:"varint,5,rep,packed,name=rpt_uint64,json=rptUint64,proto3" json:"rpt_uint64,omitempty"`
	RptFloat      []float32              `protobuf:"fixed32,6,rep,packed,name=rpt_float,json=rptFloat,proto3" json:"rpt_float,omitempty"`
	RptDouble     []float64              `protobuf:"fixed64,7,rep,packed,name=rpt_double,json=rptDouble,proto3" json:"rpt_double,omitempty"`
	RptString     []string               `protobuf:"bytes,8,rep,name=rpt_string,json=rptString,proto3" json:"rpt_string,omitempty"`
	RptBytes      [][]byte               `protobuf:"bytes,9,rep,name=rpt_bytes,json=rptBytes,proto3" json:"rpt_bytes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Repeats) Reset() {
	*x = Repeats{}
	mi := &file_internal_testprotos_textpb3_test_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Repeats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Repeats) ProtoMessage() {}

func (x *Repeats) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_textpb3_test_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Repeats.ProtoReflect.Descriptor instead.
func (*Repeats) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_textpb3_test_proto_rawDescGZIP(), []int{1}
}

func (x *Repeats) GetRptBool() []bool {
	if x != nil {
		return x.RptBool
	}
	return nil
}

func (x *Repeats) GetRptInt32() []int32 {
	if x != nil {
		return x.RptInt32
	}
	return nil
}

func (x *Repeats) GetRptInt64() []int64 {
	if x != nil {
		return x.RptInt64
	}
	return nil
}

func (x *Repeats) GetRptUint32() []uint32 {
	if x != nil {
		return x.RptUint32
	}
	return nil
}

func (x *Repeats) GetRptUint64() []uint64 {
	if x != nil {
		return x.RptUint64
	}
	return nil
}

func (x *Repeats) GetRptFloat() []float32 {
	if x != nil {
		return x.RptFloat
	}
	return nil
}

func (x *Repeats) GetRptDouble() []float64 {
	if x != nil {
		return x.RptDouble
	}
	return nil
}

func (x *Repeats) GetRptString() []string {
	if x != nil {
		return x.RptString
	}
	return nil
}

func (x *Repeats) GetRptBytes() [][]byte {
	if x != nil {
		return x.RptBytes
	}
	return nil
}

type Proto3Optional struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OptBool       *bool                  `protobuf:"varint,1,opt,name=opt_bool,json=optBool,proto3,oneof" json:"opt_bool,omitempty"`
	OptInt32      *int32                 `protobuf:"varint,2,opt,name=opt_int32,json=optInt32,proto3,oneof" json:"opt_int32,omitempty"`
	OptInt64      *int64                 `protobuf:"varint,3,opt,name=opt_int64,json=optInt64,proto3,oneof" json:"opt_int64,omitempty"`
	OptUint32     *uint32                `protobuf:"varint,4,opt,name=opt_uint32,json=optUint32,proto3,oneof" json:"opt_uint32,omitempty"`
	OptUint64     *uint64                `protobuf:"varint,5,opt,name=opt_uint64,json=optUint64,proto3,oneof" json:"opt_uint64,omitempty"`
	OptSint32     *int32                 `protobuf:"zigzag32,12,opt,name=opt_sint32,json=optSint32,proto3,oneof" json:"opt_sint32,omitempty"`
	OptSint64     *int64                 `protobuf:"zigzag64,13,opt,name=opt_sint64,json=optSint64,proto3,oneof" json:"opt_sint64,omitempty"`
	OptFixed32    *uint32                `protobuf:"fixed32,14,opt,name=opt_fixed32,json=optFixed32,proto3,oneof" json:"opt_fixed32,omitempty"`
	OptFixed64    *uint64                `protobuf:"fixed64,15,opt,name=opt_fixed64,json=optFixed64,proto3,oneof" json:"opt_fixed64,omitempty"`
	OptSfixed32   *int32                 `protobuf:"fixed32,16,opt,name=opt_sfixed32,json=optSfixed32,proto3,oneof" json:"opt_sfixed32,omitempty"`
	OptSfixed64   *int64                 `protobuf:"fixed64,17,opt,name=opt_sfixed64,json=optSfixed64,proto3,oneof" json:"opt_sfixed64,omitempty"`
	OptFloat      *float32               `protobuf:"fixed32,20,opt,name=opt_float,json=optFloat,proto3,oneof" json:"opt_float,omitempty"`
	OptDouble     *float64               `protobuf:"fixed64,21,opt,name=opt_double,json=optDouble,proto3,oneof" json:"opt_double,omitempty"`
	OptBytes      []byte                 `protobuf:"bytes,8,opt,name=opt_bytes,json=optBytes,proto3,oneof" json:"opt_bytes,omitempty"`
	OptString     *string                `protobuf:"bytes,9,opt,name=opt_string,json=optString,proto3,oneof" json:"opt_string,omitempty"`
	OptEnum       *Enum                  `protobuf:"varint,10,opt,name=opt_enum,json=optEnum,proto3,enum=pb3.Enum,oneof" json:"opt_enum,omitempty"`
	OptMessage    *Nested                `protobuf:"bytes,11,opt,name=opt_message,json=optMessage,proto3,oneof" json:"opt_message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Proto3Optional) Reset() {
	*x = Proto3Optional{}
	mi := &file_internal_testprotos_textpb3_test_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Proto3Optional) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proto3Optional) ProtoMessage() {}

func (x *Proto3Optional) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_textpb3_test_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proto3Optional.ProtoReflect.Descriptor instead.
func (*Proto3Optional) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_textpb3_test_proto_rawDescGZIP(), []int{2}
}

func (x *Proto3Optional) GetOptBool() bool {
	if x != nil && x.OptBool != nil {
		return *x.OptBool
	}
	return false
}

func (x *Proto3Optional) GetOptInt32() int32 {
	if x != nil && x.OptInt32 != nil {
		return *x.OptInt32
	}
	return 0
}

func (x *Proto3Optional) GetOptInt64() int64 {
	if x != nil && x.OptInt64 != nil {
		return *x.OptInt64
	}
	return 0
}

func (x *Proto3Optional) GetOptUint32() uint32 {
	if x != nil && x.OptUint32 != nil {
		return *x.OptUint32
	}
	return 0
}

func (x *Proto3Optional) GetOptUint64() uint64 {
	if x != nil && x.OptUint64 != nil {
		return *x.OptUint64
	}
	return 0
}

func (x *Proto3Optional) GetOptSint32() int32 {
	if x != nil && x.OptSint32 != nil {
		return *x.OptSint32
	}
	return 0
}

func (x *Proto3Optional) GetOptSint64() int64 {
	if x != nil && x.OptSint64 != nil {
		return *x.OptSint64
	}
	return 0
}

func (x *Proto3Optional) GetOptFixed32() uint32 {
	if x != nil && x.OptFixed32 != nil {
		return *x.OptFixed32
	}
	return 0
}

func (x *Proto3Optional) GetOptFixed64() uint64 {
	if x != nil && x.OptFixed64 != nil {
		return *x.OptFixed64
	}
	return 0
}

func (x *Proto3Optional) GetOptSfixed32() int32 {
	if x != nil && x.OptSfixed32 != nil {
		return *x.OptSfixed32
	}
	return 0
}

func (x *Proto3Optional) GetOptSfixed64() int64 {
	if x != nil && x.OptSfixed64 != nil {
		return *x.OptSfixed64
	}
	return 0
}

func (x *Proto3Optional) GetOptFloat() float32 {
	if x != nil && x.OptFloat != nil {
		return *x.OptFloat
	}
	return 0
}

func (x *Proto3Optional) GetOptDouble() float64 {
	if x != nil && x.OptDouble != nil {
		return *x.OptDouble
	}
	return 0
}

func (x *Proto3Optional) GetOptBytes() []byte {
	if x != nil {
		return x.OptBytes
	}
	return nil
}

func (x *Proto3Optional) GetOptString() string {
	if x != nil && x.OptString != nil {
		return *x.OptString
	}
	return ""
}

func (x *Proto3Optional) GetOptEnum() Enum {
	if x != nil && x.OptEnum != nil {
		return *x.OptEnum
	}
	return Enum_ZERO
}

func (x *Proto3Optional) GetOptMessage() *Nested {
	if x != nil {
		return x.OptMessage
	}
	return nil
}

type OptionalEnums struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	OptEnum       *Enum                    `protobuf:"varint,1,opt,name=opt_enum,json=optEnum,proto3,enum=pb3.Enum,oneof" json:"opt_enum,omitempty"`
	OptNestedEnum OptionalEnums_NestedEnum `protobuf:"varint,3,opt,name=opt_nested_enum,json=optNestedEnum,proto3,enum=pb3.OptionalEnums_NestedEnum" json:"opt_nested_enum,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OptionalEnums) Reset() {
	*x = OptionalEnums{}
	mi := &file_internal_testprotos_textpb3_test_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OptionalEnums) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionalEnums) ProtoMessage() {}

func (x *OptionalEnums) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_textpb3_test_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionalEnums.ProtoReflect.Descriptor instead.
func (*OptionalEnums) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_textpb3_test_proto_rawDescGZIP(), []int{3}
}

func (x *OptionalEnums) GetOptEnum() Enum {
	if x != nil && x.OptEnum != nil {
		return *x.OptEnum
	}
	return Enum_ZERO
}

func (x *OptionalEnums) GetOptNestedEnum() OptionalEnums_NestedEnum {
	if x != nil {
		return x.OptNestedEnum
	}
	return OptionalEnums_CERO
}

// Message contains enum fields.
type Enums struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SEnum         Enum                   `protobuf:"varint,1,opt,name=s_enum,json=sEnum,proto3,enum=pb3.Enum" json:"s_enum,omitempty"`
	SNestedEnum   Enums_NestedEnum       `protobuf:"varint,3,opt,name=s_nested_enum,json=sNestedEnum,proto3,enum=pb3.Enums_NestedEnum" json:"s_nested_enum,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Enums) Reset() {
	*x = Enums{}
	mi := &file_internal_testprotos_textpb3_test_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Enums) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Enums) ProtoMessage() {}

func (x *Enums) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_textpb3_test_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Enums.ProtoReflect.Descriptor instead.
func (*Enums) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_textpb3_test_proto_rawDescGZIP(), []int{4}
}

func (x *Enums) GetSEnum() Enum {
	if x != nil {
		return x.SEnum
	}
	return Enum_ZERO
}

func (x *Enums) GetSNestedEnum() Enums_NestedEnum {
	if x != nil {
		return x.SNestedEnum
	}
	return Enums_CERO
}

// Message contains nested message field.
type Nests struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SNested       *Nested                `protobuf:"bytes,2,opt,name=s_nested,json=sNested,proto3" json:"s_nested,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Nests) Reset() {
	*x = Nests{}
	mi := &file_internal_testprotos_textpb3_test_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Nests) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nests) ProtoMessage() {}

func (x *Nests) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_textpb3_test_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nests.ProtoReflect.Descriptor instead.
func (*Nests) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_textpb3_test_proto_rawDescGZIP(), []int{5}
}

func (x *Nests) GetSNested() *Nested {
	if x != nil {
		return x.SNested
	}
	return nil
}

// Message type used as submessage.
type Nested struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SString       string                 `protobuf:"bytes,1,opt,name=s_string,json=sString,proto3" json:"s_string,omitempty"`
	SNested       *Nested                `protobuf:"bytes,2,opt,name=s_nested,json=sNested,proto3" json:"s_nested,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Nested) Reset() {
	*x = Nested{}
	mi := &file_internal_testprotos_textpb3_test_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Nested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nested) ProtoMessage() {}

func (x *Nested) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_textpb3_test_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nested.ProtoReflect.Descriptor instead.
func (*Nested) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_textpb3_test_proto_rawDescGZIP(), []int{6}
}

func (x *Nested) GetSString() string {
	if x != nil {
		return x.SString
	}
	return ""
}

func (x *Nested) GetSNested() *Nested {
	if x != nil {
		return x.SNested
	}
	return nil
}

// Message contains oneof field.
type Oneofs struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Union:
	//
	//	*Oneofs_OneofEnum
	//	*Oneofs_OneofString
	//	*Oneofs_OneofNested
	Union         isOneofs_Union `protobuf_oneof:"union"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Oneofs) Reset() {
	*x = Oneofs{}
	mi := &file_internal_testprotos_textpb3_test_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Oneofs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Oneofs) ProtoMessage() {}

func (x *Oneofs) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_textpb3_test_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Oneofs.ProtoReflect.Descriptor instead.
func (*Oneofs) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_textpb3_test_proto_rawDescGZIP(), []int{7}
}

func (x *Oneofs) GetUnion() isOneofs_Union {
	if x != nil {
		return x.Union
	}
	return nil
}

func (x *Oneofs) GetOneofEnum() Enum {
	if x != nil {
		if x, ok := x.Union.(*Oneofs_OneofEnum); ok {
			return x.OneofEnum
		}
	}
	return Enum_ZERO
}

func (x *Oneofs) GetOneofString() string {
	if x != nil {
		if x, ok := x.Union.(*Oneofs_OneofString); ok {
			return x.OneofString
		}
	}
	return ""
}

func (x *Oneofs) GetOneofNested() *Nested {
	if x != nil {
		if x, ok := x.Union.(*Oneofs_OneofNested); ok {
			return x.OneofNested
		}
	}
	return nil
}

type isOneofs_Union interface {
	isOneofs_Union()
}

type Oneofs_OneofEnum struct {
	OneofEnum Enum `protobuf:"varint,1,opt,name=oneof_enum,json=oneofEnum,proto3,enum=pb3.Enum,oneof"`
}

type Oneofs_OneofString struct {
	OneofString string `protobuf:"bytes,2,opt,name=oneof_string,json=oneofString,proto3,oneof"`
}

type Oneofs_OneofNested struct {
	OneofNested *Nested `protobuf:"bytes,3,opt,name=oneof_nested,json=oneofNested,proto3,oneof"`
}

func (*Oneofs_OneofEnum) isOneofs_Union() {}

func (*Oneofs_OneofString) isOneofs_Union() {}

func (*Oneofs_OneofNested) isOneofs_Union() {}

// Message contains map fields.
type Maps struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Int32ToStr    map[int32]string       `protobuf:"bytes,1,rep,name=int32_to_str,json=int32ToStr,proto3" json:"int32_to_str,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	BoolToUint32  map[bool]uint32        `protobuf:"bytes,2,rep,name=bool_to_uint32,json=boolToUint32,proto3" json:"bool_to_uint32,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	Uint64ToEnum  map[uint64]Enum        `protobuf:"bytes,3,rep,name=uint64_to_enum,json=uint64ToEnum,proto3" json:"uint64_to_enum,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value,enum=pb3.Enum"`
	StrToNested   map[string]*Nested     `protobuf:"bytes,4,rep,name=str_to_nested,json=strToNested,proto3" json:"str_to_nested,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	StrToOneofs   map[string]*Oneofs     `protobuf:"bytes,5,rep,name=str_to_oneofs,json=strToOneofs,proto3" json:"str_to_oneofs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Maps) Reset() {
	*x = Maps{}
	mi := &file_internal_testprotos_textpb3_test_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Maps) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Maps) ProtoMessage() {}

func (x *Maps) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_textpb3_test_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Maps.ProtoReflect.Descriptor instead.
func (*Maps) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_textpb3_test_proto_rawDescGZIP(), []int{8}
}

func (x *Maps) GetInt32ToStr() map[int32]string {
	if x != nil {
		return x.Int32ToStr
	}
	return nil
}

func (x *Maps) GetBoolToUint32() map[bool]uint32 {
	if x != nil {
		return x.BoolToUint32
	}
	return nil
}

func (x *Maps) GetUint64ToEnum() map[uint64]Enum {
	if x != nil {
		return x.Uint64ToEnum
	}
	return nil
}

func (x *Maps) GetStrToNested() map[string]*Nested {
	if x != nil {
		return x.StrToNested
	}
	return nil
}

func (x *Maps) GetStrToOneofs() map[string]*Oneofs {
	if x != nil {
		return x.StrToOneofs
	}
	return nil
}

// Message for testing json_name option.
type JSONNames struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SString       string                 `protobuf:"bytes,1,opt,name=s_string,json=foo_bar,proto3" json:"s_string,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JSONNames) Reset() {
	*x = JSONNames{}
	mi := &file_internal_testprotos_textpb3_test_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JSONNames) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JSONNames) ProtoMessage() {}

func (x *JSONNames) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_textpb3_test_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JSONNames.ProtoReflect.Descriptor instead.
func (*JSONNames) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_textpb3_test_proto_rawDescGZIP(), []int{9}
}

func (x *JSONNames) GetSString() string {
	if x != nil {
		return x.SString
	}
	return ""
}

// Message contains reserved field name.
type ReservedFieldNames struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OptInt32      int32                  `protobuf:"varint,1,opt,name=opt_int32,json=optInt32,proto3" json:"opt_int32,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReservedFieldNames) Reset() {
	*x = ReservedFieldNames{}
	mi := &file_internal_testprotos_textpb3_test_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReservedFieldNames) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservedFieldNames) ProtoMessage() {}

func (x *ReservedFieldNames) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_textpb3_test_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservedFieldNames.ProtoReflect.Descriptor instead.
func (*ReservedFieldNames) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_textpb3_test_proto_rawDescGZIP(), []int{10}
}

func (x *ReservedFieldNames) GetOptInt32() int32 {
	if x != nil {
		return x.OptInt32
	}
	return 0
}

var File_internal_testprotos_textpb3_test_proto protoreflect.FileDescriptor

const file_internal_testprotos_textpb3_test_proto_rawDesc = "\n&internal/testprotos/textpb3/test.protopb3\"\x9e\nScalars\ns_bool (RsBool\ns_int32 (RsInt32\ns_int64 (RsInt64\ns_uint32 (\rRsUint32\ns_uint64 (RsUint64\ns_sint32 (RsSint32\ns_sint64 (RsSint64\n	s_fixed32 (RsFixed32\n	s_fixed64	 (RsFixed64\n\ns_sfixed32\n (R	sSfixed32\n\ns_sfixed64 (R	sSfixed64\ns_float (RsFloat\ns_double (RsDouble\ns_bytes (RsBytes\ns_string\r (	RsString\"\x94\nRepeats\nrpt_bool (RrptBool\n	rpt_int32 (RrptInt32\n	rpt_int64 (RrptInt64\n\nrpt_uint32 (\rR	rptUint32\n\nrpt_uint64 (R	rptUint64\n	rpt_float (RrptFloat\n\nrpt_double (R	rptDouble\n\nrpt_string (	R	rptString\n	rpt_bytes	 (RrptBytes\"\x88\nProto3Optional\nopt_bool (H\x00RoptBool\x88 \n	opt_int32 (HRoptInt32\x88 \n	opt_int64 (HRoptInt64\x88\"\n\nopt_uint32 (\rHR	optUint32\x88\"\n\nopt_uint64 (HR	optUint64\x88\"\n\nopt_sint32 (HR	optSint32\x88\"\n\nopt_sint64\r (HR	optSint64\x88$\nopt_fixed32 (HR\noptFixed32\x88$\nopt_fixed64 (HR\noptFixed64\x88&\nopt_sfixed32 (H	RoptSfixed32\x88&\nopt_sfixed64 (H\nRoptSfixed64\x88 \n	opt_float (HRoptFloat\x88\"\n\nopt_double (HR	optDouble\x88 \n	opt_bytes (H\rRoptBytes\x88\"\n\nopt_string	 (	HR	optString\x88)\nopt_enum\n (2	.pb3.EnumHRoptEnum\x881\nopt_message (2.pb3.NestedHR\noptMessage\x88B\n	_opt_boolB\n\n_opt_int32B\n\n_opt_int64B\r\n_opt_uint32B\r\n_opt_uint64B\r\n_opt_sint32B\r\n_opt_sint64B\n_opt_fixed32B\n_opt_fixed64B\n\r_opt_sfixed32B\n\r_opt_sfixed64B\n\n_opt_floatB\r\n_opt_doubleB\n\n_opt_bytesB\r\n_opt_stringB\n	_opt_enumB\n_opt_message\"\xc2\n\rOptionalEnums)\nopt_enum (2	.pb3.EnumH\x00RoptEnum\x88E\nopt_nested_enum (2.pb3.OptionalEnums.NestedEnumR\roptNestedEnum\"2\n\nNestedEnum\nCERO\x00\nUNO\nDOS\nDIEZ\nB\n	_opt_enum\"\x98\nEnums \ns_enum (2	.pb3.EnumRsEnum9\n\rs_nested_enum (2.pb3.Enums.NestedEnumRsNestedEnum\"2\n\nNestedEnum\nCERO\x00\nUNO\nDOS\nDIEZ\n\"/\nNests&\ns_nested (2.pb3.NestedRsNested\"K\nNested\ns_string (	RsString&\ns_nested (2.pb3.NestedRsNested\"\x94\nOneofs*\n\noneof_enum (2	.pb3.EnumH\x00R	oneofEnum#\noneof_string (	H\x00RoneofString0\noneof_nested (2.pb3.NestedH\x00RoneofNestedB\nunion\"\xaf\nMaps;\nint32_to_str (2.pb3.Maps.Int32ToStrEntryR\nint32ToStrA\nbool_to_uint32 (2.pb3.Maps.BoolToUint32EntryRboolToUint32A\nuint64_to_enum (2.pb3.Maps.Uint64ToEnumEntryRuint64ToEnum>\n\rstr_to_nested (2.pb3.Maps.StrToNestedEntryRstrToNested>\n\rstr_to_oneofs (2.pb3.Maps.StrToOneofsEntryRstrToOneofs=\nInt32ToStrEntry\nkey (Rkey\nvalue (	Rvalue:8?\nBoolToUint32Entry\nkey (Rkey\nvalue (\rRvalue:8J\nUint64ToEnumEntry\nkey (Rkey\nvalue (2	.pb3.EnumRvalue:8K\nStrToNestedEntry\nkey (	Rkey!\nvalue (2.pb3.NestedRvalue:8K\nStrToOneofsEntry\nkey (	Rkey!\nvalue (2.pb3.OneofsRvalue:8\"&\n	JSONNames\ns_string (	Rfoo_bar\"A\nReservedFieldNames\n	opt_int32 (RoptInt32Rreserved_field*+\nEnum\nZERO\x00\nONE\nTWO\nTEN\nB8Z6google.golang.org/protobuf/internal/testprotos/textpb3bproto3"

var (
	file_internal_testprotos_textpb3_test_proto_rawDescOnce sync.Once
	file_internal_testprotos_textpb3_test_proto_rawDescData = file_internal_testprotos_textpb3_test_proto_rawDesc
)

func file_internal_testprotos_textpb3_test_proto_rawDescGZIP() []byte {
	file_internal_testprotos_textpb3_test_proto_rawDescOnce.Do(func() {
		file_internal_testprotos_textpb3_test_proto_rawDescData = string(protoimpl.X.CompressGZIP([]byte(file_internal_testprotos_textpb3_test_proto_rawDescData)))
	})
	return []byte(file_internal_testprotos_textpb3_test_proto_rawDescData)
}

var file_internal_testprotos_textpb3_test_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_internal_testprotos_textpb3_test_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_internal_testprotos_textpb3_test_proto_goTypes = []any{
	(Enum)(0),                     // 0: pb3.Enum
	(OptionalEnums_NestedEnum)(0), // 1: pb3.OptionalEnums.NestedEnum
	(Enums_NestedEnum)(0),         // 2: pb3.Enums.NestedEnum
	(*Scalars)(nil),               // 3: pb3.Scalars
	(*Repeats)(nil),               // 4: pb3.Repeats
	(*Proto3Optional)(nil),        // 5: pb3.Proto3Optional
	(*OptionalEnums)(nil),         // 6: pb3.OptionalEnums
	(*Enums)(nil),                 // 7: pb3.Enums
	(*Nests)(nil),                 // 8: pb3.Nests
	(*Nested)(nil),                // 9: pb3.Nested
	(*Oneofs)(nil),                // 10: pb3.Oneofs
	(*Maps)(nil),                  // 11: pb3.Maps
	(*JSONNames)(nil),             // 12: pb3.JSONNames
	(*ReservedFieldNames)(nil),    // 13: pb3.ReservedFieldNames
	nil,                           // 14: pb3.Maps.Int32ToStrEntry
	nil,                           // 15: pb3.Maps.BoolToUint32Entry
	nil,                           // 16: pb3.Maps.Uint64ToEnumEntry
	nil,                           // 17: pb3.Maps.StrToNestedEntry
	nil,                           // 18: pb3.Maps.StrToOneofsEntry
}
var file_internal_testprotos_textpb3_test_proto_depIdxs = []int32{
	0,  // 0: pb3.Proto3Optional.opt_enum:type_name -> pb3.Enum
	9,  // 1: pb3.Proto3Optional.opt_message:type_name -> pb3.Nested
	0,  // 2: pb3.OptionalEnums.opt_enum:type_name -> pb3.Enum
	1,  // 3: pb3.OptionalEnums.opt_nested_enum:type_name -> pb3.OptionalEnums.NestedEnum
	0,  // 4: pb3.Enums.s_enum:type_name -> pb3.Enum
	2,  // 5: pb3.Enums.s_nested_enum:type_name -> pb3.Enums.NestedEnum
	9,  // 6: pb3.Nests.s_nested:type_name -> pb3.Nested
	9,  // 7: pb3.Nested.s_nested:type_name -> pb3.Nested
	0,  // 8: pb3.Oneofs.oneof_enum:type_name -> pb3.Enum
	9,  // 9: pb3.Oneofs.oneof_nested:type_name -> pb3.Nested
	14, // 10: pb3.Maps.int32_to_str:type_name -> pb3.Maps.Int32ToStrEntry
	15, // 11: pb3.Maps.bool_to_uint32:type_name -> pb3.Maps.BoolToUint32Entry
	16, // 12: pb3.Maps.uint64_to_enum:type_name -> pb3.Maps.Uint64ToEnumEntry
	17, // 13: pb3.Maps.str_to_nested:type_name -> pb3.Maps.StrToNestedEntry
	18, // 14: pb3.Maps.str_to_oneofs:type_name -> pb3.Maps.StrToOneofsEntry
	0,  // 15: pb3.Maps.Uint64ToEnumEntry.value:type_name -> pb3.Enum
	9,  // 16: pb3.Maps.StrToNestedEntry.value:type_name -> pb3.Nested
	10, // 17: pb3.Maps.StrToOneofsEntry.value:type_name -> pb3.Oneofs
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_internal_testprotos_textpb3_test_proto_init() }
func file_internal_testprotos_textpb3_test_proto_init() {
	if File_internal_testprotos_textpb3_test_proto != nil {
		return
	}
	file_internal_testprotos_textpb3_test_proto_msgTypes[2].OneofWrappers = []any{}
	file_internal_testprotos_textpb3_test_proto_msgTypes[3].OneofWrappers = []any{}
	file_internal_testprotos_textpb3_test_proto_msgTypes[7].OneofWrappers = []any{
		(*Oneofs_OneofEnum)(nil),
		(*Oneofs_OneofString)(nil),
		(*Oneofs_OneofNested)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: []byte(file_internal_testprotos_textpb3_test_proto_rawDesc),
			NumEnums:      3,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_textpb3_test_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_textpb3_test_proto_depIdxs,
		EnumInfos:         file_internal_testprotos_textpb3_test_proto_enumTypes,
		MessageInfos:      file_internal_testprotos_textpb3_test_proto_msgTypes,
	}.Build()
	File_internal_testprotos_textpb3_test_proto = out.File
	file_internal_testprotos_textpb3_test_proto_goTypes = nil
	file_internal_testprotos_textpb3_test_proto_depIdxs = nil
}
