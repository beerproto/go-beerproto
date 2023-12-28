//*
// BeerProto
//
// Another beer format, written in protocol buffer.
//
// Copyright (c) 2020 Ross Merrigan

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: beerproto/v1/fermentation.proto

package beerprotov1

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

// FermentationProcedureType defines the procedure for performing fermentation
type FermentationProcedureType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description       string                  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Notes             string                  `protobuf:"bytes,3,opt,name=notes,proto3" json:"notes,omitempty"`
	FermentationSteps []*FermentationStepType `protobuf:"bytes,4,rep,name=fermentation_steps,json=fermentationSteps,proto3" json:"fermentation_steps,omitempty"`
	Name              string                  `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FermentationProcedureType) Reset() {
	*x = FermentationProcedureType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beerproto_v1_fermentation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FermentationProcedureType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FermentationProcedureType) ProtoMessage() {}

func (x *FermentationProcedureType) ProtoReflect() protoreflect.Message {
	mi := &file_beerproto_v1_fermentation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FermentationProcedureType.ProtoReflect.Descriptor instead.
func (*FermentationProcedureType) Descriptor() ([]byte, []int) {
	return file_beerproto_v1_fermentation_proto_rawDescGZIP(), []int{0}
}

func (x *FermentationProcedureType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FermentationProcedureType) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *FermentationProcedureType) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *FermentationProcedureType) GetFermentationSteps() []*FermentationStepType {
	if x != nil {
		return x.FermentationSteps
	}
	return nil
}

func (x *FermentationProcedureType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_beerproto_v1_fermentation_proto protoreflect.FileDescriptor

var file_beerproto_v1_fermentation_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x66,
	0x65, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x1a,
	0x24, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65,
	0x72, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x01, 0x0a, 0x19, 0x46, 0x65, 0x72, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x51, 0x0a, 0x12, 0x66,
	0x65, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x65, 0x70,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x65, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x11, 0x66, 0x65, 0x72,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x65, 0x70, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x42, 0xb2, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x46, 0x65, 0x72, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x6f, 0x2f,
	0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x65, 0x65,
	0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa, 0x02,
	0x0c, 0x42, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c,
	0x42, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x42,
	0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x42, 0x65, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_beerproto_v1_fermentation_proto_rawDescOnce sync.Once
	file_beerproto_v1_fermentation_proto_rawDescData = file_beerproto_v1_fermentation_proto_rawDesc
)

func file_beerproto_v1_fermentation_proto_rawDescGZIP() []byte {
	file_beerproto_v1_fermentation_proto_rawDescOnce.Do(func() {
		file_beerproto_v1_fermentation_proto_rawDescData = protoimpl.X.CompressGZIP(file_beerproto_v1_fermentation_proto_rawDescData)
	})
	return file_beerproto_v1_fermentation_proto_rawDescData
}

var file_beerproto_v1_fermentation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_beerproto_v1_fermentation_proto_goTypes = []interface{}{
	(*FermentationProcedureType)(nil), // 0: beerproto.v1.FermentationProcedureType
	(*FermentationStepType)(nil),      // 1: beerproto.v1.FermentationStepType
}
var file_beerproto_v1_fermentation_proto_depIdxs = []int32{
	1, // 0: beerproto.v1.FermentationProcedureType.fermentation_steps:type_name -> beerproto.v1.FermentationStepType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_beerproto_v1_fermentation_proto_init() }
func file_beerproto_v1_fermentation_proto_init() {
	if File_beerproto_v1_fermentation_proto != nil {
		return
	}
	file_beerproto_v1_fermentation_step_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_beerproto_v1_fermentation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FermentationProcedureType); i {
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
			RawDescriptor: file_beerproto_v1_fermentation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_beerproto_v1_fermentation_proto_goTypes,
		DependencyIndexes: file_beerproto_v1_fermentation_proto_depIdxs,
		MessageInfos:      file_beerproto_v1_fermentation_proto_msgTypes,
	}.Build()
	File_beerproto_v1_fermentation_proto = out.File
	file_beerproto_v1_fermentation_proto_rawDesc = nil
	file_beerproto_v1_fermentation_proto_goTypes = nil
	file_beerproto_v1_fermentation_proto_depIdxs = nil
}
