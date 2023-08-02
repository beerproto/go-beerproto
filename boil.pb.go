//*
// BeerProto
//
// Another beer format, written in protocol buffer.
//
// Copyright (c) 2020 Ross Merrigan

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: beerproto/v1/boil.proto

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

// BoilProcedureType defines the procedure for performing a boil. A boil procedure with no steps is the same as a standard single step boil
type BoilProcedureType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PreBoilSize *VolumeType     `protobuf:"bytes,2,opt,name=pre_boil_size,json=preBoilSize,proto3" json:"pre_boil_size,omitempty"`
	BoilTime    *TimeType       `protobuf:"bytes,3,opt,name=boil_time,json=boilTime,proto3" json:"boil_time,omitempty"`
	BoilSteps   []*BoilStepType `protobuf:"bytes,4,rep,name=boil_steps,json=boilSteps,proto3" json:"boil_steps,omitempty"`
	Name        string          `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Description string          `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Notes       string          `protobuf:"bytes,7,opt,name=notes,proto3" json:"notes,omitempty"`
}

func (x *BoilProcedureType) Reset() {
	*x = BoilProcedureType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beerproto_v1_boil_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoilProcedureType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoilProcedureType) ProtoMessage() {}

func (x *BoilProcedureType) ProtoReflect() protoreflect.Message {
	mi := &file_beerproto_v1_boil_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoilProcedureType.ProtoReflect.Descriptor instead.
func (*BoilProcedureType) Descriptor() ([]byte, []int) {
	return file_beerproto_v1_boil_proto_rawDescGZIP(), []int{0}
}

func (x *BoilProcedureType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BoilProcedureType) GetPreBoilSize() *VolumeType {
	if x != nil {
		return x.PreBoilSize
	}
	return nil
}

func (x *BoilProcedureType) GetBoilTime() *TimeType {
	if x != nil {
		return x.BoilTime
	}
	return nil
}

func (x *BoilProcedureType) GetBoilSteps() []*BoilStepType {
	if x != nil {
		return x.BoilSteps
	}
	return nil
}

func (x *BoilProcedureType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BoilProcedureType) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BoilProcedureType) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

var File_beerproto_v1_boil_proto protoreflect.FileDescriptor

var file_beerproto_v1_boil_proto_rawDesc = []byte{
	0x0a, 0x17, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x62,
	0x6f, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x65, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x69, 0x6c, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x75, 0x6e, 0x69, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x02, 0x0a, 0x11,
	0x42, 0x6f, 0x69, 0x6c, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x3c, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x5f, 0x62, 0x6f, 0x69, 0x6c, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x42, 0x6f, 0x69, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x33, 0x0a, 0x09, 0x62, 0x6f, 0x69, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x62, 0x6f, 0x69, 0x6c,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x62, 0x6f, 0x69, 0x6c, 0x5f, 0x73, 0x74, 0x65,
	0x70, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x69, 0x6c, 0x53, 0x74, 0x65, 0x70,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x62, 0x6f, 0x69, 0x6c, 0x53, 0x74, 0x65, 0x70, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x42, 0xaa, 0x01, 0x0a, 0x10,
	0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31,
	0x42, 0x09, 0x42, 0x6f, 0x69, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x6f,
	0x2f, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x65,
	0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa,
	0x02, 0x0c, 0x42, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x0c, 0x42, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18,
	0x42, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x42, 0x65, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_beerproto_v1_boil_proto_rawDescOnce sync.Once
	file_beerproto_v1_boil_proto_rawDescData = file_beerproto_v1_boil_proto_rawDesc
)

func file_beerproto_v1_boil_proto_rawDescGZIP() []byte {
	file_beerproto_v1_boil_proto_rawDescOnce.Do(func() {
		file_beerproto_v1_boil_proto_rawDescData = protoimpl.X.CompressGZIP(file_beerproto_v1_boil_proto_rawDescData)
	})
	return file_beerproto_v1_boil_proto_rawDescData
}

var file_beerproto_v1_boil_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_beerproto_v1_boil_proto_goTypes = []interface{}{
	(*BoilProcedureType)(nil), // 0: beerproto.v1.BoilProcedureType
	(*VolumeType)(nil),        // 1: beerproto.v1.VolumeType
	(*TimeType)(nil),          // 2: beerproto.v1.TimeType
	(*BoilStepType)(nil),      // 3: beerproto.v1.BoilStepType
}
var file_beerproto_v1_boil_proto_depIdxs = []int32{
	1, // 0: beerproto.v1.BoilProcedureType.pre_boil_size:type_name -> beerproto.v1.VolumeType
	2, // 1: beerproto.v1.BoilProcedureType.boil_time:type_name -> beerproto.v1.TimeType
	3, // 2: beerproto.v1.BoilProcedureType.boil_steps:type_name -> beerproto.v1.BoilStepType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_beerproto_v1_boil_proto_init() }
func file_beerproto_v1_boil_proto_init() {
	if File_beerproto_v1_boil_proto != nil {
		return
	}
	file_beerproto_v1_boil_step_proto_init()
	file_beerproto_v1_measureable_units_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_beerproto_v1_boil_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoilProcedureType); i {
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
			RawDescriptor: file_beerproto_v1_boil_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_beerproto_v1_boil_proto_goTypes,
		DependencyIndexes: file_beerproto_v1_boil_proto_depIdxs,
		MessageInfos:      file_beerproto_v1_boil_proto_msgTypes,
	}.Build()
	File_beerproto_v1_boil_proto = out.File
	file_beerproto_v1_boil_proto_rawDesc = nil
	file_beerproto_v1_boil_proto_goTypes = nil
	file_beerproto_v1_boil_proto_depIdxs = nil
}
