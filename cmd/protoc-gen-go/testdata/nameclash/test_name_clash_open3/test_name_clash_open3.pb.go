// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This proto verifies that we keep the name mangling algorithm (which is
// position dependent) intact in the protoc_gen_go generator. The field names
// and the getter names have to be kept intact over time, both in the OPEN and
// in the HYBRID API. How fields are "mangled" is described in a comment per
// field.

// The order of "evaluation" of fields is important. Fields are evaluated in
// order of appearance, except the oneof union names, that are evaluated after
// their first member. For each field, check if there is a previous field name
// or getter name that clashes with this field or it's getter. In case there is
// a clash, add an _ to the field name and repeat. In the case of oneof's, the
// union will be renamed if it clashes with it's first member, but not if it
// clashes with it's second.

// This scheme is here for backwards compatibility.
// The type of clashes that can be are the following:
// 1 - My field name clashes with their getter name
// 2 - My getter name clashes with their field name

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd/protoc-gen-go/testdata/nameclash/test_name_clash_open3.proto

package test_name_clash_open3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/gofeaturespb"
	reflect "reflect"
	sync "sync"
)

type M0 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	I1            int32                  `protobuf:"varint,1,opt,name=i1" json:"i1,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *M0) Reset() {
	*x = M0{}
	mi := &file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *M0) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M0) ProtoMessage() {}

func (x *M0) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M0.ProtoReflect.Descriptor instead.
func (*M0) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_rawDescGZIP(), []int{0}
}

func (x *M0) GetI1() int32 {
	if x != nil {
		return x.I1
	}
	return 0
}

type M1 struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// initial name in Go     | Clashes with field | type | final name
	// Foo                    | -                  | -    | Foo
	// GetFoo                 | foo                | 1    | GetFoo_
	// GetGetFoo              | -                  | -    | GetGetFoo
	Foo           *M0 `protobuf:"bytes,1,opt,name=foo" json:"foo,omitempty"`
	GetFoo_       *M0 `protobuf:"bytes,2,opt,name=get_foo,json=getFoo" json:"get_foo,omitempty"`
	GetGetFoo     *M0 `protobuf:"bytes,3,opt,name=get_get_foo,json=getGetFoo" json:"get_get_foo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *M1) Reset() {
	*x = M1{}
	mi := &file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *M1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M1) ProtoMessage() {}

func (x *M1) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M1.ProtoReflect.Descriptor instead.
func (*M1) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_rawDescGZIP(), []int{1}
}

func (x *M1) GetFoo() *M0 {
	if x != nil {
		return x.Foo
	}
	return nil
}

func (x *M1) GetGetFoo_() *M0 {
	if x != nil {
		return x.GetFoo_
	}
	return nil
}

func (x *M1) GetGetGetFoo() *M0 {
	if x != nil {
		return x.GetGetFoo
	}
	return nil
}

type M2 struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// initial name in Go     | Clashes with field | type | final name
	// GetGetFoo              | -                  | -    | GetGetFoo
	// GetFoo                 | get_get_foo        | 2    | GetFoo_
	// Foo                    | -                  | -    | Foo
	GetGetFoo     *M0 `protobuf:"bytes,3,opt,name=get_get_foo,json=getGetFoo" json:"get_get_foo,omitempty"`
	GetFoo_       *M0 `protobuf:"bytes,2,opt,name=get_foo,json=getFoo" json:"get_foo,omitempty"`
	Foo           *M0 `protobuf:"bytes,1,opt,name=foo" json:"foo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *M2) Reset() {
	*x = M2{}
	mi := &file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *M2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M2) ProtoMessage() {}

func (x *M2) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M2.ProtoReflect.Descriptor instead.
func (*M2) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_rawDescGZIP(), []int{2}
}

func (x *M2) GetGetGetFoo() *M0 {
	if x != nil {
		return x.GetGetFoo
	}
	return nil
}

func (x *M2) GetGetFoo_() *M0 {
	if x != nil {
		return x.GetFoo_
	}
	return nil
}

func (x *M2) GetFoo() *M0 {
	if x != nil {
		return x.Foo
	}
	return nil
}

type M3 struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// initial name in Go     | Clashes with field | type | final name
	// GetFoo                 | -                  | -    | GetFoo
	// GetGetFoo              | get_foo            | 1    | GetGetFoo_
	// Foo                    | get_foo            | 2    | Foo_
	GetFoo        *M0 `protobuf:"bytes,2,opt,name=get_foo,json=getFoo" json:"get_foo,omitempty"`
	GetGetFoo_    *M0 `protobuf:"bytes,3,opt,name=get_get_foo,json=getGetFoo" json:"get_get_foo,omitempty"`
	Foo_          *M0 `protobuf:"bytes,1,opt,name=foo" json:"foo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *M3) Reset() {
	*x = M3{}
	mi := &file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *M3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M3) ProtoMessage() {}

func (x *M3) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M3.ProtoReflect.Descriptor instead.
func (*M3) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_rawDescGZIP(), []int{3}
}

func (x *M3) GetGetFoo() *M0 {
	if x != nil {
		return x.GetFoo
	}
	return nil
}

func (x *M3) GetGetGetFoo_() *M0 {
	if x != nil {
		return x.GetGetFoo_
	}
	return nil
}

func (x *M3) GetFoo_() *M0 {
	if x != nil {
		return x.Foo_
	}
	return nil
}

type M4 struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// initial name in Go     | Clashes with field | type | final name
	// GetFoo                 | -                  | -    | GetFoo
	// GetGetFoo              | get_foo            | 1    | GetGetFoo_
	//
	//	GetGetGetFoo          | -                  | -    | GetGetGetFoo
	//	                      |                    |      |
	//
	// Foo                    | get_foo            | 2    | Foo_
	GetFoo *M0 `protobuf:"bytes,2,opt,name=get_foo,json=getFoo" json:"get_foo,omitempty"`
	// Types that are valid to be assigned to GetGetFoo_:
	//
	//	*M4_GetGetGetFoo
	GetGetFoo_    isM4_GetGetFoo_ `protobuf_oneof:"get_get_foo"`
	Foo_          *M0             `protobuf:"bytes,1,opt,name=foo" json:"foo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *M4) Reset() {
	*x = M4{}
	mi := &file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *M4) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M4) ProtoMessage() {}

func (x *M4) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M4.ProtoReflect.Descriptor instead.
func (*M4) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_rawDescGZIP(), []int{4}
}

func (x *M4) GetGetFoo() *M0 {
	if x != nil {
		return x.GetFoo
	}
	return nil
}

func (x *M4) GetGetGetFoo_() isM4_GetGetFoo_ {
	if x != nil {
		return x.GetGetFoo_
	}
	return nil
}

func (x *M4) GetGetGetGetFoo() int32 {
	if x != nil {
		if x, ok := x.GetGetFoo_.(*M4_GetGetGetFoo); ok {
			return x.GetGetGetFoo
		}
	}
	return 0
}

func (x *M4) GetFoo_() *M0 {
	if x != nil {
		return x.Foo_
	}
	return nil
}

type isM4_GetGetFoo_ interface {
	isM4_GetGetFoo_()
}

type M4_GetGetGetFoo struct {
	GetGetGetFoo int32 `protobuf:"varint,3,opt,name=get_get_get_foo,json=getGetGetFoo,oneof"`
}

func (*M4_GetGetGetFoo) isM4_GetGetFoo_() {}

type M5 struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Note evaluation order - get_get_foo before get_get_get_foo
	// initial name in Go     | Clashes with field | type | final name
	// GetFoo                 | -                  | -    | GetFoo
	// GetGetGetFoo           | -                  | -    | GetGetGetFoo
	//
	//	GetGetFoo             | get_foo            | 1    | GetGetFoo_
	//	                      |                    |      |
	//
	// Foo                    | get_foo            | 2    | Foo_
	GetFoo *M0 `protobuf:"bytes,2,opt,name=get_foo,json=getFoo" json:"get_foo,omitempty"`
	// Types that are valid to be assigned to GetGetGetFoo:
	//
	//	*M5_GetGetFoo_
	GetGetGetFoo  isM5_GetGetGetFoo `protobuf_oneof:"get_get_get_foo"`
	Foo_          *M0               `protobuf:"bytes,1,opt,name=foo" json:"foo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *M5) Reset() {
	*x = M5{}
	mi := &file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *M5) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M5) ProtoMessage() {}

func (x *M5) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M5.ProtoReflect.Descriptor instead.
func (*M5) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_rawDescGZIP(), []int{5}
}

func (x *M5) GetGetFoo() *M0 {
	if x != nil {
		return x.GetFoo
	}
	return nil
}

func (x *M5) GetGetGetGetFoo() isM5_GetGetGetFoo {
	if x != nil {
		return x.GetGetGetFoo
	}
	return nil
}

func (x *M5) GetGetGetFoo_() int32 {
	if x != nil {
		if x, ok := x.GetGetGetFoo.(*M5_GetGetFoo_); ok {
			return x.GetGetFoo_
		}
	}
	return 0
}

func (x *M5) GetFoo_() *M0 {
	if x != nil {
		return x.Foo_
	}
	return nil
}

type isM5_GetGetGetFoo interface {
	isM5_GetGetGetFoo()
}

type M5_GetGetFoo_ struct {
	GetGetFoo_ int32 `protobuf:"varint,3,opt,name=get_get_foo,json=getGetFoo,oneof"`
}

func (*M5_GetGetFoo_) isM5_GetGetGetFoo() {}

type M6 struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Note evaluation order - get_get_get_foo before get_get_foo
	// initial name in Go     | Clashes with field | type | final name
	// GetGetFoo              | -                  | -    | GetGetFoo
	//
	//	GetGetGetFoo          | -                  | -    | GetGetGetFoo
	//	                      |                    |      |
	//
	// GetFoo                 | get_get_foo        | 2    | GetFoo_
	// Foo                    | -                  | -    | Foo
	//
	// Types that are valid to be assigned to GetGetFoo:
	//
	//	*M6_GetGetGetFoo
	GetGetFoo     isM6_GetGetFoo `protobuf_oneof:"get_get_foo"`
	GetFoo_       *M0            `protobuf:"bytes,2,opt,name=get_foo,json=getFoo" json:"get_foo,omitempty"`
	Foo           *M0            `protobuf:"bytes,1,opt,name=foo" json:"foo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *M6) Reset() {
	*x = M6{}
	mi := &file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *M6) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M6) ProtoMessage() {}

func (x *M6) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M6.ProtoReflect.Descriptor instead.
func (*M6) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_rawDescGZIP(), []int{6}
}

func (x *M6) GetGetGetFoo() isM6_GetGetFoo {
	if x != nil {
		return x.GetGetFoo
	}
	return nil
}

func (x *M6) GetGetGetGetFoo() int32 {
	if x != nil {
		if x, ok := x.GetGetFoo.(*M6_GetGetGetFoo); ok {
			return x.GetGetGetFoo
		}
	}
	return 0
}

func (x *M6) GetGetFoo_() *M0 {
	if x != nil {
		return x.GetFoo_
	}
	return nil
}

func (x *M6) GetFoo() *M0 {
	if x != nil {
		return x.Foo
	}
	return nil
}

type isM6_GetGetFoo interface {
	isM6_GetGetFoo()
}

type M6_GetGetGetFoo struct {
	GetGetGetFoo int32 `protobuf:"varint,3,opt,name=get_get_get_foo,json=getGetGetFoo,oneof"`
}

func (*M6_GetGetGetFoo) isM6_GetGetFoo() {}

type M7 struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Note evaluation order - bar before get_get_foo, then get_get_get_foo
	// initial name in Go     | Clashes with field | type | final name
	// GetGetFoo              | -                  | -    | GetGetFoo
	//
	//	Bar                   | -                  | -    | Bar
	//	GetFoo                | foo                | 1    | GetFoo_
	//	                      |                    |      |
	//
	// Foo                    | -                  | -    | Foo
	//
	// Types that are valid to be assigned to GetGetFoo:
	//
	//	*M7_Bar
	//	*M7_GetFoo_
	GetGetFoo     isM7_GetGetFoo `protobuf_oneof:"get_get_foo"`
	Foo           *M0            `protobuf:"bytes,1,opt,name=foo" json:"foo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *M7) Reset() {
	*x = M7{}
	mi := &file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *M7) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M7) ProtoMessage() {}

func (x *M7) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M7.ProtoReflect.Descriptor instead.
func (*M7) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_rawDescGZIP(), []int{7}
}

func (x *M7) GetGetGetFoo() isM7_GetGetFoo {
	if x != nil {
		return x.GetGetFoo
	}
	return nil
}

func (x *M7) GetBar() bool {
	if x != nil {
		if x, ok := x.GetGetFoo.(*M7_Bar); ok {
			return x.Bar
		}
	}
	return false
}

func (x *M7) GetGetFoo_() int32 {
	if x != nil {
		if x, ok := x.GetGetFoo.(*M7_GetFoo_); ok {
			return x.GetFoo_
		}
	}
	return 0
}

func (x *M7) GetFoo() *M0 {
	if x != nil {
		return x.Foo
	}
	return nil
}

type isM7_GetGetFoo interface {
	isM7_GetGetFoo()
}

type M7_Bar struct {
	Bar bool `protobuf:"varint,4,opt,name=bar,oneof"`
}

type M7_GetFoo_ struct {
	GetFoo_ int32 `protobuf:"varint,3,opt,name=get_foo,json=getFoo,oneof"`
}

func (*M7_Bar) isM7_GetGetFoo() {}

func (*M7_GetFoo_) isM7_GetGetFoo() {}

type M8 struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Note evaluation order - get_get_foo before get_get_get_foo
	// initial name in Go     | Clashes with field | type | final name
	// GetGetGetFoo           | get_get_foo        | 1    | GetGetGetFoo_
	//
	//	GetGetFoo             | -                  | -    | GetGetFoo
	//	                      |                    |      |
	//
	// GetFoo                 | get_get_foo        | 2    | GetFoo_
	// Foo                    | -                  | -    | Foo
	//
	// Types that are valid to be assigned to GetGetGetFoo_:
	//
	//	*M8_GetGetFoo
	GetGetGetFoo_ isM8_GetGetGetFoo_ `protobuf_oneof:"get_get_get_foo"`
	GetFoo_       *M0                `protobuf:"bytes,2,opt,name=get_foo,json=getFoo" json:"get_foo,omitempty"`
	Foo           *M0                `protobuf:"bytes,1,opt,name=foo" json:"foo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *M8) Reset() {
	*x = M8{}
	mi := &file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *M8) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M8) ProtoMessage() {}

func (x *M8) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M8.ProtoReflect.Descriptor instead.
func (*M8) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_rawDescGZIP(), []int{8}
}

func (x *M8) GetGetGetGetFoo_() isM8_GetGetGetFoo_ {
	if x != nil {
		return x.GetGetGetFoo_
	}
	return nil
}

func (x *M8) GetGetGetFoo() int32 {
	if x != nil {
		if x, ok := x.GetGetGetFoo_.(*M8_GetGetFoo); ok {
			return x.GetGetFoo
		}
	}
	return 0
}

func (x *M8) GetGetFoo_() *M0 {
	if x != nil {
		return x.GetFoo_
	}
	return nil
}

func (x *M8) GetFoo() *M0 {
	if x != nil {
		return x.Foo
	}
	return nil
}

type isM8_GetGetGetFoo_ interface {
	isM8_GetGetGetFoo_()
}

type M8_GetGetFoo struct {
	GetGetFoo int32 `protobuf:"varint,3,opt,name=get_get_foo,json=getGetFoo,oneof"`
}

func (*M8_GetGetFoo) isM8_GetGetGetFoo_() {}

type M9 struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Note evaluation order - get_get_foo before get_get_get_foo, then get_foo
	// initial name in Go     | Clashes with field | type | final name
	// GetGetGetFoo           | get_get_foo        | 1    | GetGetGetFoo_
	//
	//	GetGetFoo             | -                  | -    | GetGetFoo
	//	GetFoo                | get_get_foo        | 2    | GetFoo_
	//	                      |                    |      |
	//
	// Foo                    | -                  | -    | Foo
	//
	// Types that are valid to be assigned to GetGetGetFoo_:
	//
	//	*M9_GetGetFoo
	//	*M9_GetFoo_
	GetGetGetFoo_ isM9_GetGetGetFoo_ `protobuf_oneof:"get_get_get_foo"`
	Foo           *M0                `protobuf:"bytes,1,opt,name=foo" json:"foo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *M9) Reset() {
	*x = M9{}
	mi := &file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *M9) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M9) ProtoMessage() {}

func (x *M9) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M9.ProtoReflect.Descriptor instead.
func (*M9) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_rawDescGZIP(), []int{9}
}

func (x *M9) GetGetGetGetFoo_() isM9_GetGetGetFoo_ {
	if x != nil {
		return x.GetGetGetFoo_
	}
	return nil
}

func (x *M9) GetGetGetFoo() int32 {
	if x != nil {
		if x, ok := x.GetGetGetFoo_.(*M9_GetGetFoo); ok {
			return x.GetGetFoo
		}
	}
	return 0
}

func (x *M9) GetGetFoo_() int32 {
	if x != nil {
		if x, ok := x.GetGetGetFoo_.(*M9_GetFoo_); ok {
			return x.GetFoo_
		}
	}
	return 0
}

func (x *M9) GetFoo() *M0 {
	if x != nil {
		return x.Foo
	}
	return nil
}

type isM9_GetGetGetFoo_ interface {
	isM9_GetGetGetFoo_()
}

type M9_GetGetFoo struct {
	GetGetFoo int32 `protobuf:"varint,3,opt,name=get_get_foo,json=getGetFoo,oneof"`
}

type M9_GetFoo_ struct {
	GetFoo_ int32 `protobuf:"varint,2,opt,name=get_foo,json=getFoo,oneof"`
}

func (*M9_GetGetFoo) isM9_GetGetGetFoo_() {}

func (*M9_GetFoo_) isM9_GetGetGetFoo_() {}

var File_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto protoreflect.FileDescriptor

const file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_rawDesc = "\n@cmd/protoc-gen-go/testdata/nameclash/test_name_clash_open3.proto%net.proto2.go.testdata.nameclashopen3!google/protobuf/go_features.proto\"\nM0\ni1 (Ri1\"\xd0\nM1;\nfoo (2).net.proto2.go.testdata.nameclashopen3.M0RfooB\nget_foo (2).net.proto2.go.testdata.nameclashopen3.M0RgetFooI\nget_get_foo (2).net.proto2.go.testdata.nameclashopen3.M0R	getGetFoo\"\xd0\nM2I\nget_get_foo (2).net.proto2.go.testdata.nameclashopen3.M0R	getGetFooB\nget_foo (2).net.proto2.go.testdata.nameclashopen3.M0RgetFoo;\nfoo (2).net.proto2.go.testdata.nameclashopen3.M0Rfoo\"\xd0\nM3B\nget_foo (2).net.proto2.go.testdata.nameclashopen3.M0RgetFooI\nget_get_foo (2).net.proto2.go.testdata.nameclashopen3.M0R	getGetFoo;\nfoo (2).net.proto2.go.testdata.nameclashopen3.M0Rfoo\"\xbd\nM4B\nget_foo (2).net.proto2.go.testdata.nameclashopen3.M0RgetFoo'\nget_get_get_foo (H\x00RgetGetGetFoo;\nfoo (2).net.proto2.go.testdata.nameclashopen3.M0RfooB\r\nget_get_foo\"\xba\nM5B\nget_foo (2).net.proto2.go.testdata.nameclashopen3.M0RgetFoo \nget_get_foo (H\x00R	getGetFoo;\nfoo (2).net.proto2.go.testdata.nameclashopen3.M0RfooB\nget_get_get_foo\"\xbd\nM6'\nget_get_get_foo (H\x00RgetGetGetFooB\nget_foo (2).net.proto2.go.testdata.nameclashopen3.M0RgetFoo;\nfoo (2).net.proto2.go.testdata.nameclashopen3.M0RfooB\r\nget_get_foo\"\nM7\nbar (H\x00Rbar\nget_foo (H\x00RgetFoo;\nfoo (2).net.proto2.go.testdata.nameclashopen3.M0RfooB\r\nget_get_foo\"\xba\nM8 \nget_get_foo (H\x00R	getGetFooB\nget_foo (2).net.proto2.go.testdata.nameclashopen3.M0RgetFoo;\nfoo (2).net.proto2.go.testdata.nameclashopen3.M0RfooB\nget_get_get_foo\"\x91\nM9 \nget_get_foo (H\x00R	getGetFoo\nget_foo (H\x00RgetFoo;\nfoo (2).net.proto2.go.testdata.nameclashopen3.M0RfooB\nget_get_get_fooBaZUgoogle.golang.org/protobuf/cmd/protoc-gen-go/testdata/nameclash/test_name_clash_open3\x92\xd2>beditionsp\xe8"

var (
	file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_rawDescData = file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_rawDesc
)

func file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_rawDescData = string(protoimpl.X.CompressGZIP([]byte(file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_rawDescData)))
	})
	return []byte(file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_rawDescData)
}

var file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_goTypes = []any{
	(*M0)(nil), // 0: net.proto2.go.testdata.nameclashopen3.M0
	(*M1)(nil), // 1: net.proto2.go.testdata.nameclashopen3.M1
	(*M2)(nil), // 2: net.proto2.go.testdata.nameclashopen3.M2
	(*M3)(nil), // 3: net.proto2.go.testdata.nameclashopen3.M3
	(*M4)(nil), // 4: net.proto2.go.testdata.nameclashopen3.M4
	(*M5)(nil), // 5: net.proto2.go.testdata.nameclashopen3.M5
	(*M6)(nil), // 6: net.proto2.go.testdata.nameclashopen3.M6
	(*M7)(nil), // 7: net.proto2.go.testdata.nameclashopen3.M7
	(*M8)(nil), // 8: net.proto2.go.testdata.nameclashopen3.M8
	(*M9)(nil), // 9: net.proto2.go.testdata.nameclashopen3.M9
}
var file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_depIdxs = []int32{
	0,  // 0: net.proto2.go.testdata.nameclashopen3.M1.foo:type_name -> net.proto2.go.testdata.nameclashopen3.M0
	0,  // 1: net.proto2.go.testdata.nameclashopen3.M1.get_foo:type_name -> net.proto2.go.testdata.nameclashopen3.M0
	0,  // 2: net.proto2.go.testdata.nameclashopen3.M1.get_get_foo:type_name -> net.proto2.go.testdata.nameclashopen3.M0
	0,  // 3: net.proto2.go.testdata.nameclashopen3.M2.get_get_foo:type_name -> net.proto2.go.testdata.nameclashopen3.M0
	0,  // 4: net.proto2.go.testdata.nameclashopen3.M2.get_foo:type_name -> net.proto2.go.testdata.nameclashopen3.M0
	0,  // 5: net.proto2.go.testdata.nameclashopen3.M2.foo:type_name -> net.proto2.go.testdata.nameclashopen3.M0
	0,  // 6: net.proto2.go.testdata.nameclashopen3.M3.get_foo:type_name -> net.proto2.go.testdata.nameclashopen3.M0
	0,  // 7: net.proto2.go.testdata.nameclashopen3.M3.get_get_foo:type_name -> net.proto2.go.testdata.nameclashopen3.M0
	0,  // 8: net.proto2.go.testdata.nameclashopen3.M3.foo:type_name -> net.proto2.go.testdata.nameclashopen3.M0
	0,  // 9: net.proto2.go.testdata.nameclashopen3.M4.get_foo:type_name -> net.proto2.go.testdata.nameclashopen3.M0
	0,  // 10: net.proto2.go.testdata.nameclashopen3.M4.foo:type_name -> net.proto2.go.testdata.nameclashopen3.M0
	0,  // 11: net.proto2.go.testdata.nameclashopen3.M5.get_foo:type_name -> net.proto2.go.testdata.nameclashopen3.M0
	0,  // 12: net.proto2.go.testdata.nameclashopen3.M5.foo:type_name -> net.proto2.go.testdata.nameclashopen3.M0
	0,  // 13: net.proto2.go.testdata.nameclashopen3.M6.get_foo:type_name -> net.proto2.go.testdata.nameclashopen3.M0
	0,  // 14: net.proto2.go.testdata.nameclashopen3.M6.foo:type_name -> net.proto2.go.testdata.nameclashopen3.M0
	0,  // 15: net.proto2.go.testdata.nameclashopen3.M7.foo:type_name -> net.proto2.go.testdata.nameclashopen3.M0
	0,  // 16: net.proto2.go.testdata.nameclashopen3.M8.get_foo:type_name -> net.proto2.go.testdata.nameclashopen3.M0
	0,  // 17: net.proto2.go.testdata.nameclashopen3.M8.foo:type_name -> net.proto2.go.testdata.nameclashopen3.M0
	0,  // 18: net.proto2.go.testdata.nameclashopen3.M9.foo:type_name -> net.proto2.go.testdata.nameclashopen3.M0
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_init() }
func file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_init() {
	if File_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto != nil {
		return
	}
	file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes[4].OneofWrappers = []any{
		(*M4_GetGetGetFoo)(nil),
	}
	file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes[5].OneofWrappers = []any{
		(*M5_GetGetFoo_)(nil),
	}
	file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes[6].OneofWrappers = []any{
		(*M6_GetGetGetFoo)(nil),
	}
	file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes[7].OneofWrappers = []any{
		(*M7_Bar)(nil),
		(*M7_GetFoo_)(nil),
	}
	file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes[8].OneofWrappers = []any{
		(*M8_GetGetFoo)(nil),
	}
	file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes[9].OneofWrappers = []any{
		(*M9_GetGetFoo)(nil),
		(*M9_GetFoo_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: []byte(file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_rawDesc),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_depIdxs,
		MessageInfos:      file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_msgTypes,
	}.Build()
	File_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto = out.File
	file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_goTypes = nil
	file_cmd_protoc_gen_go_testdata_nameclash_test_name_clash_open3_proto_depIdxs = nil
}
