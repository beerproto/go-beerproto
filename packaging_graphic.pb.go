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
// source: beerproto/v1/packaging_graphic.proto

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

type PackagingGraphicType_PositionUnit int32

const (
	PackagingGraphicType_POSITION_UNIT_UNSPECIFIED      PackagingGraphicType_PositionUnit = 0
	PackagingGraphicType_POSITION_UNIT_BODY_FRONT       PackagingGraphicType_PositionUnit = 1
	PackagingGraphicType_POSITION_UNIT_BODY_BACK        PackagingGraphicType_PositionUnit = 2
	PackagingGraphicType_POSITION_UNIT_BODY_WRAP_AROUND PackagingGraphicType_PositionUnit = 3
	PackagingGraphicType_POSITION_UNIT_NECK_FRONT       PackagingGraphicType_PositionUnit = 4
	PackagingGraphicType_POSITION_UNIT_NECK_BACK        PackagingGraphicType_PositionUnit = 5
	PackagingGraphicType_POSITION_UNIT_NECK_WRAP_AROUND PackagingGraphicType_PositionUnit = 6
	PackagingGraphicType_POSITION_UNIT_CAP              PackagingGraphicType_PositionUnit = 7
	PackagingGraphicType_POSITION_UNIT_CARRIER          PackagingGraphicType_PositionUnit = 8
)

// Enum value maps for PackagingGraphicType_PositionUnit.
var (
	PackagingGraphicType_PositionUnit_name = map[int32]string{
		0: "POSITION_UNIT_UNSPECIFIED",
		1: "POSITION_UNIT_BODY_FRONT",
		2: "POSITION_UNIT_BODY_BACK",
		3: "POSITION_UNIT_BODY_WRAP_AROUND",
		4: "POSITION_UNIT_NECK_FRONT",
		5: "POSITION_UNIT_NECK_BACK",
		6: "POSITION_UNIT_NECK_WRAP_AROUND",
		7: "POSITION_UNIT_CAP",
		8: "POSITION_UNIT_CARRIER",
	}
	PackagingGraphicType_PositionUnit_value = map[string]int32{
		"POSITION_UNIT_UNSPECIFIED":      0,
		"POSITION_UNIT_BODY_FRONT":       1,
		"POSITION_UNIT_BODY_BACK":        2,
		"POSITION_UNIT_BODY_WRAP_AROUND": 3,
		"POSITION_UNIT_NECK_FRONT":       4,
		"POSITION_UNIT_NECK_BACK":        5,
		"POSITION_UNIT_NECK_WRAP_AROUND": 6,
		"POSITION_UNIT_CAP":              7,
		"POSITION_UNIT_CARRIER":          8,
	}
)

func (x PackagingGraphicType_PositionUnit) Enum() *PackagingGraphicType_PositionUnit {
	p := new(PackagingGraphicType_PositionUnit)
	*p = x
	return p
}

func (x PackagingGraphicType_PositionUnit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PackagingGraphicType_PositionUnit) Descriptor() protoreflect.EnumDescriptor {
	return file_beerproto_v1_packaging_graphic_proto_enumTypes[0].Descriptor()
}

func (PackagingGraphicType_PositionUnit) Type() protoreflect.EnumType {
	return &file_beerproto_v1_packaging_graphic_proto_enumTypes[0]
}

func (x PackagingGraphicType_PositionUnit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PackagingGraphicType_PositionUnit.Descriptor instead.
func (PackagingGraphicType_PositionUnit) EnumDescriptor() ([]byte, []int) {
	return file_beerproto_v1_packaging_graphic_proto_rawDescGZIP(), []int{0, 0}
}

type PackagingGraphicType_GraphicType int32

const (
	PackagingGraphicType_GRAPHIC_TYPE_UNSPECIFIED PackagingGraphicType_GraphicType = 0
	PackagingGraphicType_GRAPHIC_TYPE_SVG         PackagingGraphicType_GraphicType = 1
	PackagingGraphicType_GRAPHIC_TYPE_SVGZ        PackagingGraphicType_GraphicType = 2
	PackagingGraphicType_GRAPHIC_TYPE_AI          PackagingGraphicType_GraphicType = 3
	PackagingGraphicType_GRAPHIC_TYPE_CDR         PackagingGraphicType_GraphicType = 4
	PackagingGraphicType_GRAPHIC_TYPE_CDX         PackagingGraphicType_GraphicType = 5
	PackagingGraphicType_GRAPHIC_TYPE_ODG         PackagingGraphicType_GraphicType = 6
	PackagingGraphicType_GRAPHIC_TYPE_EPS         PackagingGraphicType_GraphicType = 7
	PackagingGraphicType_GRAPHIC_TYPE_PDF         PackagingGraphicType_GraphicType = 8
	PackagingGraphicType_GRAPHIC_TYPE_PNG         PackagingGraphicType_GraphicType = 9
	PackagingGraphicType_GRAPHIC_TYPE_JPG         PackagingGraphicType_GraphicType = 10
	PackagingGraphicType_GRAPHIC_TYPE_GIF         PackagingGraphicType_GraphicType = 11
)

// Enum value maps for PackagingGraphicType_GraphicType.
var (
	PackagingGraphicType_GraphicType_name = map[int32]string{
		0:  "GRAPHIC_TYPE_UNSPECIFIED",
		1:  "GRAPHIC_TYPE_SVG",
		2:  "GRAPHIC_TYPE_SVGZ",
		3:  "GRAPHIC_TYPE_AI",
		4:  "GRAPHIC_TYPE_CDR",
		5:  "GRAPHIC_TYPE_CDX",
		6:  "GRAPHIC_TYPE_ODG",
		7:  "GRAPHIC_TYPE_EPS",
		8:  "GRAPHIC_TYPE_PDF",
		9:  "GRAPHIC_TYPE_PNG",
		10: "GRAPHIC_TYPE_JPG",
		11: "GRAPHIC_TYPE_GIF",
	}
	PackagingGraphicType_GraphicType_value = map[string]int32{
		"GRAPHIC_TYPE_UNSPECIFIED": 0,
		"GRAPHIC_TYPE_SVG":         1,
		"GRAPHIC_TYPE_SVGZ":        2,
		"GRAPHIC_TYPE_AI":          3,
		"GRAPHIC_TYPE_CDR":         4,
		"GRAPHIC_TYPE_CDX":         5,
		"GRAPHIC_TYPE_ODG":         6,
		"GRAPHIC_TYPE_EPS":         7,
		"GRAPHIC_TYPE_PDF":         8,
		"GRAPHIC_TYPE_PNG":         9,
		"GRAPHIC_TYPE_JPG":         10,
		"GRAPHIC_TYPE_GIF":         11,
	}
)

func (x PackagingGraphicType_GraphicType) Enum() *PackagingGraphicType_GraphicType {
	p := new(PackagingGraphicType_GraphicType)
	*p = x
	return p
}

func (x PackagingGraphicType_GraphicType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PackagingGraphicType_GraphicType) Descriptor() protoreflect.EnumDescriptor {
	return file_beerproto_v1_packaging_graphic_proto_enumTypes[1].Descriptor()
}

func (PackagingGraphicType_GraphicType) Type() protoreflect.EnumType {
	return &file_beerproto_v1_packaging_graphic_proto_enumTypes[1]
}

func (x PackagingGraphicType_GraphicType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PackagingGraphicType_GraphicType.Descriptor instead.
func (PackagingGraphicType_GraphicType) EnumDescriptor() ([]byte, []int) {
	return file_beerproto_v1_packaging_graphic_proto_rawDescGZIP(), []int{0, 1}
}

type PackagingGraphicType_UnitsType int32

const (
	PackagingGraphicType_UNITS_TYPE_UNSPECIFIED PackagingGraphicType_UnitsType = 0
	PackagingGraphicType_UNITS_TYPE_MM          PackagingGraphicType_UnitsType = 1
	PackagingGraphicType_UNITS_TYPE_IN          PackagingGraphicType_UnitsType = 2
)

// Enum value maps for PackagingGraphicType_UnitsType.
var (
	PackagingGraphicType_UnitsType_name = map[int32]string{
		0: "UNITS_TYPE_UNSPECIFIED",
		1: "UNITS_TYPE_MM",
		2: "UNITS_TYPE_IN",
	}
	PackagingGraphicType_UnitsType_value = map[string]int32{
		"UNITS_TYPE_UNSPECIFIED": 0,
		"UNITS_TYPE_MM":          1,
		"UNITS_TYPE_IN":          2,
	}
)

func (x PackagingGraphicType_UnitsType) Enum() *PackagingGraphicType_UnitsType {
	p := new(PackagingGraphicType_UnitsType)
	*p = x
	return p
}

func (x PackagingGraphicType_UnitsType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PackagingGraphicType_UnitsType) Descriptor() protoreflect.EnumDescriptor {
	return file_beerproto_v1_packaging_graphic_proto_enumTypes[2].Descriptor()
}

func (PackagingGraphicType_UnitsType) Type() protoreflect.EnumType {
	return &file_beerproto_v1_packaging_graphic_proto_enumTypes[2]
}

func (x PackagingGraphicType_UnitsType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PackagingGraphicType_UnitsType.Descriptor instead.
func (PackagingGraphicType_UnitsType) EnumDescriptor() ([]byte, []int) {
	return file_beerproto_v1_packaging_graphic_proto_rawDescGZIP(), []int{0, 2}
}

// Representation of a graphic to be placed on a vessel.
type PackagingGraphicType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position PackagingGraphicType_PositionUnit `protobuf:"varint,1,opt,name=position,proto3,enum=beerproto.v1.PackagingGraphicType_PositionUnit" json:"position,omitempty"`
	// File type
	Type PackagingGraphicType_GraphicType `protobuf:"varint,2,opt,name=type,proto3,enum=beerproto.v1.PackagingGraphicType_GraphicType" json:"type,omitempty"`
	// Base64 encoded file.
	Base64Data string `protobuf:"bytes,3,opt,name=base64_data,json=base64Data,proto3" json:"base64_data,omitempty"`
	// URLS to hosted version of image.
	Urls []string `protobuf:"bytes,4,rep,name=urls,proto3" json:"urls,omitempty"`
	// Dots per inch.
	Dpi    int32 `protobuf:"varint,5,opt,name=dpi,proto3" json:"dpi,omitempty"`
	Width  int64 `protobuf:"varint,6,opt,name=width,proto3" json:"width,omitempty"`
	Height int64 `protobuf:"varint,7,opt,name=height,proto3" json:"height,omitempty"`
	// The unit type which are used for measurements.
	Units PackagingGraphicType_UnitsType `protobuf:"varint,8,opt,name=units,proto3,enum=beerproto.v1.PackagingGraphicType_UnitsType" json:"units,omitempty"`
}

func (x *PackagingGraphicType) Reset() {
	*x = PackagingGraphicType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beerproto_v1_packaging_graphic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackagingGraphicType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackagingGraphicType) ProtoMessage() {}

func (x *PackagingGraphicType) ProtoReflect() protoreflect.Message {
	mi := &file_beerproto_v1_packaging_graphic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackagingGraphicType.ProtoReflect.Descriptor instead.
func (*PackagingGraphicType) Descriptor() ([]byte, []int) {
	return file_beerproto_v1_packaging_graphic_proto_rawDescGZIP(), []int{0}
}

func (x *PackagingGraphicType) GetPosition() PackagingGraphicType_PositionUnit {
	if x != nil {
		return x.Position
	}
	return PackagingGraphicType_POSITION_UNIT_UNSPECIFIED
}

func (x *PackagingGraphicType) GetType() PackagingGraphicType_GraphicType {
	if x != nil {
		return x.Type
	}
	return PackagingGraphicType_GRAPHIC_TYPE_UNSPECIFIED
}

func (x *PackagingGraphicType) GetBase64Data() string {
	if x != nil {
		return x.Base64Data
	}
	return ""
}

func (x *PackagingGraphicType) GetUrls() []string {
	if x != nil {
		return x.Urls
	}
	return nil
}

func (x *PackagingGraphicType) GetDpi() int32 {
	if x != nil {
		return x.Dpi
	}
	return 0
}

func (x *PackagingGraphicType) GetWidth() int64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *PackagingGraphicType) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *PackagingGraphicType) GetUnits() PackagingGraphicType_UnitsType {
	if x != nil {
		return x.Units
	}
	return PackagingGraphicType_UNITS_TYPE_UNSPECIFIED
}

var File_beerproto_v1_packaging_graphic_proto protoreflect.FileDescriptor

var file_beerproto_v1_packaging_graphic_proto_rawDesc = []byte{
	0x0a, 0x24, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x76, 0x31, 0x22, 0xef, 0x07, 0x0a, 0x14, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4b, 0x0a,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2f, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x54,
	0x79, 0x70, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74,
	0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x72, 0x6c, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x70, 0x69, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x64, 0x70, 0x69, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x42, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x61, 0x70, 0x68,
	0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x22, 0x9d, 0x02, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x4f, 0x53, 0x49,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x4f, 0x53, 0x49, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x42, 0x4f, 0x44, 0x59, 0x5f, 0x46, 0x52,
	0x4f, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x42, 0x4f, 0x44, 0x59, 0x5f, 0x42, 0x41, 0x43, 0x4b,
	0x10, 0x02, 0x12, 0x22, 0x0a, 0x1e, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55,
	0x4e, 0x49, 0x54, 0x5f, 0x42, 0x4f, 0x44, 0x59, 0x5f, 0x57, 0x52, 0x41, 0x50, 0x5f, 0x41, 0x52,
	0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x4e, 0x45, 0x43, 0x4b, 0x5f, 0x46, 0x52, 0x4f,
	0x4e, 0x54, 0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x4e, 0x45, 0x43, 0x4b, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x10,
	0x05, 0x12, 0x22, 0x0a, 0x1e, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e,
	0x49, 0x54, 0x5f, 0x4e, 0x45, 0x43, 0x4b, 0x5f, 0x57, 0x52, 0x41, 0x50, 0x5f, 0x41, 0x52, 0x4f,
	0x55, 0x4e, 0x44, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x43, 0x41, 0x50, 0x10, 0x07, 0x12, 0x19, 0x0a, 0x15,
	0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x43, 0x41,
	0x52, 0x52, 0x49, 0x45, 0x52, 0x10, 0x08, 0x22, 0x9d, 0x02, 0x0a, 0x0b, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x47, 0x52, 0x41, 0x50, 0x48,
	0x49, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x47, 0x52, 0x41, 0x50, 0x48, 0x49, 0x43,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x56, 0x47, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x47,
	0x52, 0x41, 0x50, 0x48, 0x49, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x56, 0x47, 0x5a,
	0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x47, 0x52, 0x41, 0x50, 0x48, 0x49, 0x43, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x41, 0x49, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x47, 0x52, 0x41, 0x50, 0x48,
	0x49, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x44, 0x52, 0x10, 0x04, 0x12, 0x14, 0x0a,
	0x10, 0x47, 0x52, 0x41, 0x50, 0x48, 0x49, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x44,
	0x58, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x47, 0x52, 0x41, 0x50, 0x48, 0x49, 0x43, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4f, 0x44, 0x47, 0x10, 0x06, 0x12, 0x14, 0x0a, 0x10, 0x47, 0x52, 0x41,
	0x50, 0x48, 0x49, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x50, 0x53, 0x10, 0x07, 0x12,
	0x14, 0x0a, 0x10, 0x47, 0x52, 0x41, 0x50, 0x48, 0x49, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x50, 0x44, 0x46, 0x10, 0x08, 0x12, 0x14, 0x0a, 0x10, 0x47, 0x52, 0x41, 0x50, 0x48, 0x49, 0x43,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4e, 0x47, 0x10, 0x09, 0x12, 0x14, 0x0a, 0x10, 0x47,
	0x52, 0x41, 0x50, 0x48, 0x49, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4a, 0x50, 0x47, 0x10,
	0x0a, 0x12, 0x14, 0x0a, 0x10, 0x47, 0x52, 0x41, 0x50, 0x48, 0x49, 0x43, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x47, 0x49, 0x46, 0x10, 0x0b, 0x22, 0x4d, 0x0a, 0x09, 0x55, 0x6e, 0x69, 0x74, 0x73,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x55, 0x4e, 0x49, 0x54, 0x53, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x49, 0x54, 0x53, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d,
	0x4d, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x49, 0x54, 0x53, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x49, 0x4e, 0x10, 0x02, 0x42, 0xb6, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x62,
	0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x42, 0x15, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x65, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x6f, 0x2f, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa, 0x02, 0x0c, 0x42, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x42, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x42, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0d, 0x42, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_beerproto_v1_packaging_graphic_proto_rawDescOnce sync.Once
	file_beerproto_v1_packaging_graphic_proto_rawDescData = file_beerproto_v1_packaging_graphic_proto_rawDesc
)

func file_beerproto_v1_packaging_graphic_proto_rawDescGZIP() []byte {
	file_beerproto_v1_packaging_graphic_proto_rawDescOnce.Do(func() {
		file_beerproto_v1_packaging_graphic_proto_rawDescData = protoimpl.X.CompressGZIP(file_beerproto_v1_packaging_graphic_proto_rawDescData)
	})
	return file_beerproto_v1_packaging_graphic_proto_rawDescData
}

var file_beerproto_v1_packaging_graphic_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_beerproto_v1_packaging_graphic_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_beerproto_v1_packaging_graphic_proto_goTypes = []interface{}{
	(PackagingGraphicType_PositionUnit)(0), // 0: beerproto.v1.PackagingGraphicType.PositionUnit
	(PackagingGraphicType_GraphicType)(0),  // 1: beerproto.v1.PackagingGraphicType.GraphicType
	(PackagingGraphicType_UnitsType)(0),    // 2: beerproto.v1.PackagingGraphicType.UnitsType
	(*PackagingGraphicType)(nil),           // 3: beerproto.v1.PackagingGraphicType
}
var file_beerproto_v1_packaging_graphic_proto_depIdxs = []int32{
	0, // 0: beerproto.v1.PackagingGraphicType.position:type_name -> beerproto.v1.PackagingGraphicType.PositionUnit
	1, // 1: beerproto.v1.PackagingGraphicType.type:type_name -> beerproto.v1.PackagingGraphicType.GraphicType
	2, // 2: beerproto.v1.PackagingGraphicType.units:type_name -> beerproto.v1.PackagingGraphicType.UnitsType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_beerproto_v1_packaging_graphic_proto_init() }
func file_beerproto_v1_packaging_graphic_proto_init() {
	if File_beerproto_v1_packaging_graphic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_beerproto_v1_packaging_graphic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackagingGraphicType); i {
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
			RawDescriptor: file_beerproto_v1_packaging_graphic_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_beerproto_v1_packaging_graphic_proto_goTypes,
		DependencyIndexes: file_beerproto_v1_packaging_graphic_proto_depIdxs,
		EnumInfos:         file_beerproto_v1_packaging_graphic_proto_enumTypes,
		MessageInfos:      file_beerproto_v1_packaging_graphic_proto_msgTypes,
	}.Build()
	File_beerproto_v1_packaging_graphic_proto = out.File
	file_beerproto_v1_packaging_graphic_proto_rawDesc = nil
	file_beerproto_v1_packaging_graphic_proto_goTypes = nil
	file_beerproto_v1_packaging_graphic_proto_depIdxs = nil
}
