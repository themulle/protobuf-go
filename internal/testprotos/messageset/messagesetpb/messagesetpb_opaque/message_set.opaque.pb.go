// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/messageset/messagesetpb/messagesetpb_opaque/message_set.opaque.proto

package messagesetpb_opaque

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/gofeaturespb"
	reflect "reflect"
)

type MessageSet struct {
	state           protoimpl.MessageState `protogen:"opaque.v1"`
	extensionFields protoimpl.ExtensionFields
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *MessageSet) Reset() {
	*x = MessageSet{}
	mi := &file_internal_testprotos_messageset_messagesetpb_messagesetpb_opaque_message_set_opaque_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageSet) ProtoMessage() {}

func (x *MessageSet) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_messageset_messagesetpb_messagesetpb_opaque_message_set_opaque_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

type MessageSet_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

}

func (b0 MessageSet_builder) Build() *MessageSet {
	m0 := &MessageSet{}
	b, x := &b0, m0
	_, _ = b, x
	return m0
}

type MessageSetContainer struct {
	state                 protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_MessageSet *MessageSet            `protobuf:"bytes,1,opt,name=message_set,json=messageSet"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *MessageSetContainer) Reset() {
	*x = MessageSetContainer{}
	mi := &file_internal_testprotos_messageset_messagesetpb_messagesetpb_opaque_message_set_opaque_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageSetContainer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageSetContainer) ProtoMessage() {}

func (x *MessageSetContainer) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_messageset_messagesetpb_messagesetpb_opaque_message_set_opaque_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MessageSetContainer) GetMessageSet() *MessageSet {
	if x != nil {
		return x.xxx_hidden_MessageSet
	}
	return nil
}

func (x *MessageSetContainer) SetMessageSet(v *MessageSet) {
	x.xxx_hidden_MessageSet = v
}

func (x *MessageSetContainer) HasMessageSet() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_MessageSet != nil
}

func (x *MessageSetContainer) ClearMessageSet() {
	x.xxx_hidden_MessageSet = nil
}

type MessageSetContainer_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	MessageSet *MessageSet
}

func (b0 MessageSetContainer_builder) Build() *MessageSetContainer {
	m0 := &MessageSetContainer{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_MessageSet = b.MessageSet
	return m0
}

var File_internal_testprotos_messageset_messagesetpb_messagesetpb_opaque_message_set_opaque_proto protoreflect.FileDescriptor

const file_internal_testprotos_messageset_messagesetpb_messagesetpb_opaque_message_set_opaque_proto_rawDesc = "\nXinternal/testprotos/messageset/messagesetpb/messagesetpb_opaque/message_set.opaque.protoopaque.goproto.proto.messageset!google/protobuf/go_features.proto\"\n\nMessageSet*\xff\xff\xff\xff:\"c\nMessageSetContainerL\nmessage_set (2+.opaque.goproto.proto.messageset.MessageSetR\nmessageSetBdZZgoogle.golang.org/protobuf/internal/testprotos/messageset/messagesetpb/messagesetpb_opaque\x92\xd2>beditionsp\xe8"

var file_internal_testprotos_messageset_messagesetpb_messagesetpb_opaque_message_set_opaque_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_testprotos_messageset_messagesetpb_messagesetpb_opaque_message_set_opaque_proto_goTypes = []any{
	(*MessageSet)(nil),          // 0: opaque.goproto.proto.messageset.MessageSet
	(*MessageSetContainer)(nil), // 1: opaque.goproto.proto.messageset.MessageSetContainer
}
var file_internal_testprotos_messageset_messagesetpb_messagesetpb_opaque_message_set_opaque_proto_depIdxs = []int32{
	0, // 0: opaque.goproto.proto.messageset.MessageSetContainer.message_set:type_name -> opaque.goproto.proto.messageset.MessageSet
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_internal_testprotos_messageset_messagesetpb_messagesetpb_opaque_message_set_opaque_proto_init()
}
func file_internal_testprotos_messageset_messagesetpb_messagesetpb_opaque_message_set_opaque_proto_init() {
	if File_internal_testprotos_messageset_messagesetpb_messagesetpb_opaque_message_set_opaque_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: []byte(file_internal_testprotos_messageset_messagesetpb_messagesetpb_opaque_message_set_opaque_proto_rawDesc),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_messageset_messagesetpb_messagesetpb_opaque_message_set_opaque_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_messageset_messagesetpb_messagesetpb_opaque_message_set_opaque_proto_depIdxs,
		MessageInfos:      file_internal_testprotos_messageset_messagesetpb_messagesetpb_opaque_message_set_opaque_proto_msgTypes,
	}.Build()
	File_internal_testprotos_messageset_messagesetpb_messagesetpb_opaque_message_set_opaque_proto = out.File
	file_internal_testprotos_messageset_messagesetpb_messagesetpb_opaque_message_set_opaque_proto_goTypes = nil
	file_internal_testprotos_messageset_messagesetpb_messagesetpb_opaque_message_set_opaque_proto_depIdxs = nil
}
