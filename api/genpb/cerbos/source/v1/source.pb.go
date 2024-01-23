// Copyright 2021-2024 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: cerbos/source/v1/source.proto

package sourcev1

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

type Error_Kind int32

const (
	Error_KIND_UNSPECIFIED      Error_Kind = 0
	Error_KIND_PARSE_ERROR      Error_Kind = 1
	Error_KIND_VALIDATION_ERROR Error_Kind = 2
)

// Enum value maps for Error_Kind.
var (
	Error_Kind_name = map[int32]string{
		0: "KIND_UNSPECIFIED",
		1: "KIND_PARSE_ERROR",
		2: "KIND_VALIDATION_ERROR",
	}
	Error_Kind_value = map[string]int32{
		"KIND_UNSPECIFIED":      0,
		"KIND_PARSE_ERROR":      1,
		"KIND_VALIDATION_ERROR": 2,
	}
)

func (x Error_Kind) Enum() *Error_Kind {
	p := new(Error_Kind)
	*p = x
	return p
}

func (x Error_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Error_Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_cerbos_source_v1_source_proto_enumTypes[0].Descriptor()
}

func (Error_Kind) Type() protoreflect.EnumType {
	return &file_cerbos_source_v1_source_proto_enumTypes[0]
}

func (x Error_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Error_Kind.Descriptor instead.
func (Error_Kind) EnumDescriptor() ([]byte, []int) {
	return file_cerbos_source_v1_source_proto_rawDescGZIP(), []int{1, 0}
}

type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Line   uint32 `protobuf:"varint,1,opt,name=line,proto3" json:"line,omitempty"`
	Column uint32 `protobuf:"varint,2,opt,name=column,proto3" json:"column,omitempty"`
	Path   string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cerbos_source_v1_source_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_source_v1_source_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_cerbos_source_v1_source_proto_rawDescGZIP(), []int{0}
}

func (x *Position) GetLine() uint32 {
	if x != nil {
		return x.Line
	}
	return 0
}

func (x *Position) GetColumn() uint32 {
	if x != nil {
		return x.Column
	}
	return 0
}

func (x *Position) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind     Error_Kind `protobuf:"varint,1,opt,name=kind,proto3,enum=cerbos.source.v1.Error_Kind" json:"kind,omitempty"`
	Position *Position  `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	Message  string     `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Context  string     `protobuf:"bytes,4,opt,name=context,proto3" json:"context,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cerbos_source_v1_source_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_source_v1_source_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_cerbos_source_v1_source_proto_rawDescGZIP(), []int{1}
}

func (x *Error) GetKind() Error_Kind {
	if x != nil {
		return x.Kind
	}
	return Error_KIND_UNSPECIFIED
}

func (x *Error) GetPosition() *Position {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetContext() string {
	if x != nil {
		return x.Context
	}
	return ""
}

type SourceContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldPositions map[string]*Position `protobuf:"bytes,1,rep,name=field_positions,json=fieldPositions,proto3" json:"field_positions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Errors         []*Error             `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *SourceContext) Reset() {
	*x = SourceContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cerbos_source_v1_source_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceContext) ProtoMessage() {}

func (x *SourceContext) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_source_v1_source_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceContext.ProtoReflect.Descriptor instead.
func (*SourceContext) Descriptor() ([]byte, []int) {
	return file_cerbos_source_v1_source_proto_rawDescGZIP(), []int{2}
}

func (x *SourceContext) GetFieldPositions() map[string]*Position {
	if x != nil {
		return x.FieldPositions
	}
	return nil
}

func (x *SourceContext) GetErrors() []*Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_cerbos_source_v1_source_proto protoreflect.FileDescriptor

var file_cerbos_source_v1_source_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x22, 0x4a, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6c, 0x69, 0x6e,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0xf4, 0x01,
	0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x30, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x4b,
	0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x36, 0x0a, 0x08, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x65,
	0x72, 0x62, 0x6f, 0x73, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x4d, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x14, 0x0a,
	0x10, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x53,
	0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x4b, 0x49, 0x4e,
	0x44, 0x5f, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x02, 0x22, 0xfd, 0x01, 0x0a, 0x0d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x5c, 0x0a, 0x0f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x33, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x5d, 0x0a, 0x13, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x62, 0x2f, 0x63, 0x65, 0x72, 0x62, 0x6f,
	0x73, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cerbos_source_v1_source_proto_rawDescOnce sync.Once
	file_cerbos_source_v1_source_proto_rawDescData = file_cerbos_source_v1_source_proto_rawDesc
)

func file_cerbos_source_v1_source_proto_rawDescGZIP() []byte {
	file_cerbos_source_v1_source_proto_rawDescOnce.Do(func() {
		file_cerbos_source_v1_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_cerbos_source_v1_source_proto_rawDescData)
	})
	return file_cerbos_source_v1_source_proto_rawDescData
}

var file_cerbos_source_v1_source_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cerbos_source_v1_source_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cerbos_source_v1_source_proto_goTypes = []interface{}{
	(Error_Kind)(0),       // 0: cerbos.source.v1.Error.Kind
	(*Position)(nil),      // 1: cerbos.source.v1.Position
	(*Error)(nil),         // 2: cerbos.source.v1.Error
	(*SourceContext)(nil), // 3: cerbos.source.v1.SourceContext
	nil,                   // 4: cerbos.source.v1.SourceContext.FieldPositionsEntry
}
var file_cerbos_source_v1_source_proto_depIdxs = []int32{
	0, // 0: cerbos.source.v1.Error.kind:type_name -> cerbos.source.v1.Error.Kind
	1, // 1: cerbos.source.v1.Error.position:type_name -> cerbos.source.v1.Position
	4, // 2: cerbos.source.v1.SourceContext.field_positions:type_name -> cerbos.source.v1.SourceContext.FieldPositionsEntry
	2, // 3: cerbos.source.v1.SourceContext.errors:type_name -> cerbos.source.v1.Error
	1, // 4: cerbos.source.v1.SourceContext.FieldPositionsEntry.value:type_name -> cerbos.source.v1.Position
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_cerbos_source_v1_source_proto_init() }
func file_cerbos_source_v1_source_proto_init() {
	if File_cerbos_source_v1_source_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cerbos_source_v1_source_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
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
		file_cerbos_source_v1_source_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_cerbos_source_v1_source_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceContext); i {
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
			RawDescriptor: file_cerbos_source_v1_source_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cerbos_source_v1_source_proto_goTypes,
		DependencyIndexes: file_cerbos_source_v1_source_proto_depIdxs,
		EnumInfos:         file_cerbos_source_v1_source_proto_enumTypes,
		MessageInfos:      file_cerbos_source_v1_source_proto_msgTypes,
	}.Build()
	File_cerbos_source_v1_source_proto = out.File
	file_cerbos_source_v1_source_proto_rawDesc = nil
	file_cerbos_source_v1_source_proto_goTypes = nil
	file_cerbos_source_v1_source_proto_depIdxs = nil
}