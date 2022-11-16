//
// BeerProto
//
// Another beer format, written in protocol buffer.
//
// Copyright (c) 2020 Ross Merrigan

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: beerproto/v1/equipment.proto

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

type EquipmentItemType_EquipmentBaseForm int32

const (
	EquipmentItemType_EQUIPMENT_BASE_FORM_UNSPECIFIED EquipmentItemType_EquipmentBaseForm = 0
	// HLT
	EquipmentItemType_EQUIPMENT_BASE_FORM_HLT EquipmentItemType_EquipmentBaseForm = 1
	// Mash Tun
	EquipmentItemType_EQUIPMENT_BASE_FORM_MASH_TUN EquipmentItemType_EquipmentBaseForm = 2
	// Lauter Tun
	EquipmentItemType_EQUIPMENT_BASE_FORM_LAUTER_TUN EquipmentItemType_EquipmentBaseForm = 3
	// Brew Kettle
	EquipmentItemType_EQUIPMENT_BASE_FORM_BREW_KETTLE EquipmentItemType_EquipmentBaseForm = 4
	// Fermenter
	EquipmentItemType_EQUIPMENT_BASE_FORM_FERMENTER EquipmentItemType_EquipmentBaseForm = 5
	// Aging Vessel
	EquipmentItemType_EQUIPMENT_BASE_FORM_AGING_VESSEL EquipmentItemType_EquipmentBaseForm = 6
	// Packaging Vessel
	EquipmentItemType_EQUIPMENT_BASE_FORM_PACKAGING_VESSEL EquipmentItemType_EquipmentBaseForm = 7
)

// Enum value maps for EquipmentItemType_EquipmentBaseForm.
var (
	EquipmentItemType_EquipmentBaseForm_name = map[int32]string{
		0: "EQUIPMENT_BASE_FORM_UNSPECIFIED",
		1: "EQUIPMENT_BASE_FORM_HLT",
		2: "EQUIPMENT_BASE_FORM_MASH_TUN",
		3: "EQUIPMENT_BASE_FORM_LAUTER_TUN",
		4: "EQUIPMENT_BASE_FORM_BREW_KETTLE",
		5: "EQUIPMENT_BASE_FORM_FERMENTER",
		6: "EQUIPMENT_BASE_FORM_AGING_VESSEL",
		7: "EQUIPMENT_BASE_FORM_PACKAGING_VESSEL",
	}
	EquipmentItemType_EquipmentBaseForm_value = map[string]int32{
		"EQUIPMENT_BASE_FORM_UNSPECIFIED":      0,
		"EQUIPMENT_BASE_FORM_HLT":              1,
		"EQUIPMENT_BASE_FORM_MASH_TUN":         2,
		"EQUIPMENT_BASE_FORM_LAUTER_TUN":       3,
		"EQUIPMENT_BASE_FORM_BREW_KETTLE":      4,
		"EQUIPMENT_BASE_FORM_FERMENTER":        5,
		"EQUIPMENT_BASE_FORM_AGING_VESSEL":     6,
		"EQUIPMENT_BASE_FORM_PACKAGING_VESSEL": 7,
	}
)

func (x EquipmentItemType_EquipmentBaseForm) Enum() *EquipmentItemType_EquipmentBaseForm {
	p := new(EquipmentItemType_EquipmentBaseForm)
	*p = x
	return p
}

func (x EquipmentItemType_EquipmentBaseForm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EquipmentItemType_EquipmentBaseForm) Descriptor() protoreflect.EnumDescriptor {
	return file_beerproto_v1_equipment_proto_enumTypes[0].Descriptor()
}

func (EquipmentItemType_EquipmentBaseForm) Type() protoreflect.EnumType {
	return &file_beerproto_v1_equipment_proto_enumTypes[0]
}

func (x EquipmentItemType_EquipmentBaseForm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EquipmentItemType_EquipmentBaseForm.Descriptor instead.
func (EquipmentItemType_EquipmentBaseForm) EnumDescriptor() ([]byte, []int) {
	return file_beerproto_v1_equipment_proto_rawDescGZIP(), []int{0, 0}
}

// EquipmentType provides necessary information for individual brewing equipment
type EquipmentItemType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Notes string `protobuf:"bytes,2,opt,name=notes,proto3" json:"notes,omitempty"`
	// The volume boiled off during 1 hour, measured before and after at room temperature.
	BoilRatePerHour *VolumeType                         `protobuf:"bytes,3,opt,name=boil_rate_per_hour,json=boilRatePerHour,proto3" json:"boil_rate_per_hour,omitempty"`
	Type            string                              `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Form            EquipmentItemType_EquipmentBaseForm `protobuf:"varint,5,opt,name=form,proto3,enum=beerproto.v1.EquipmentItemType_EquipmentBaseForm" json:"form,omitempty"`
	// The volume that leaves the kettle, especially important for non-immersion chillers that cool the wort as it leaves the kettle.
	DrainRatePerMinute *VolumeType `protobuf:"bytes,6,opt,name=drain_rate_per_minute,json=drainRatePerMinute,proto3" json:"drain_rate_per_minute,omitempty"`
	// The specific heat of the piece of equipment, expressed in Cal/(g*C), especially important for when the mashtun is not preheated.
	SpecificHeat *SpecificHeatType `protobuf:"bytes,7,opt,name=specific_heat,json=specificHeat,proto3" json:"specific_heat,omitempty"`
	// The apparent volume absorbed by grain, typical values are 0.125 qt/lb (1.04 L/kg) for a mashtun, 0.08 gal/lb (0.66 L/kg) for BIAB.
	GrainAbsorptionRate *SpecificVolumeType `protobuf:"bytes,8,opt,name=grain_absorption_rate,json=grainAbsorptionRate,proto3" json:"grain_absorption_rate,omitempty"`
	Name                string              `protobuf:"bytes,9,opt,name=name,proto3" json:"name,omitempty"`
	MaximumVolume       *VolumeType         `protobuf:"bytes,10,opt,name=maximum_volume,json=maximumVolume,proto3" json:"maximum_volume,omitempty"`
	// The weight of the piece of equipment, especially important for when the mashtun is not preheated.
	Weight     *MassType    `protobuf:"bytes,11,opt,name=weight,proto3" json:"weight,omitempty"`
	Loss       *VolumeType  `protobuf:"bytes,12,opt,name=loss,proto3" json:"loss,omitempty"`
	Efficiency *PercentType `protobuf:"bytes,13,opt,name=efficiency,proto3" json:"efficiency,omitempty"`
}

func (x *EquipmentItemType) Reset() {
	*x = EquipmentItemType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beerproto_v1_equipment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EquipmentItemType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EquipmentItemType) ProtoMessage() {}

func (x *EquipmentItemType) ProtoReflect() protoreflect.Message {
	mi := &file_beerproto_v1_equipment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EquipmentItemType.ProtoReflect.Descriptor instead.
func (*EquipmentItemType) Descriptor() ([]byte, []int) {
	return file_beerproto_v1_equipment_proto_rawDescGZIP(), []int{0}
}

func (x *EquipmentItemType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EquipmentItemType) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *EquipmentItemType) GetBoilRatePerHour() *VolumeType {
	if x != nil {
		return x.BoilRatePerHour
	}
	return nil
}

func (x *EquipmentItemType) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *EquipmentItemType) GetForm() EquipmentItemType_EquipmentBaseForm {
	if x != nil {
		return x.Form
	}
	return EquipmentItemType_EQUIPMENT_BASE_FORM_UNSPECIFIED
}

func (x *EquipmentItemType) GetDrainRatePerMinute() *VolumeType {
	if x != nil {
		return x.DrainRatePerMinute
	}
	return nil
}

func (x *EquipmentItemType) GetSpecificHeat() *SpecificHeatType {
	if x != nil {
		return x.SpecificHeat
	}
	return nil
}

func (x *EquipmentItemType) GetGrainAbsorptionRate() *SpecificVolumeType {
	if x != nil {
		return x.GrainAbsorptionRate
	}
	return nil
}

func (x *EquipmentItemType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EquipmentItemType) GetMaximumVolume() *VolumeType {
	if x != nil {
		return x.MaximumVolume
	}
	return nil
}

func (x *EquipmentItemType) GetWeight() *MassType {
	if x != nil {
		return x.Weight
	}
	return nil
}

func (x *EquipmentItemType) GetLoss() *VolumeType {
	if x != nil {
		return x.Loss
	}
	return nil
}

func (x *EquipmentItemType) GetEfficiency() *PercentType {
	if x != nil {
		return x.Efficiency
	}
	return nil
}

// Provides necessary information for brewing equipment set
type EquipmentType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	EquipmentItems      []*EquipmentItemType `protobuf:"bytes,3,rep,name=equipment_items,json=equipmentItems,proto3" json:"equipment_items,omitempty"`
	BrewhouseEfficiency *PercentType         `protobuf:"bytes,4,opt,name=brewhouse_efficiency,json=brewhouseEfficiency,proto3" json:"brewhouse_efficiency,omitempty"`
}

func (x *EquipmentType) Reset() {
	*x = EquipmentType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beerproto_v1_equipment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EquipmentType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EquipmentType) ProtoMessage() {}

func (x *EquipmentType) ProtoReflect() protoreflect.Message {
	mi := &file_beerproto_v1_equipment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EquipmentType.ProtoReflect.Descriptor instead.
func (*EquipmentType) Descriptor() ([]byte, []int) {
	return file_beerproto_v1_equipment_proto_rawDescGZIP(), []int{1}
}

func (x *EquipmentType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EquipmentType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EquipmentType) GetEquipmentItems() []*EquipmentItemType {
	if x != nil {
		return x.EquipmentItems
	}
	return nil
}

func (x *EquipmentType) GetBrewhouseEfficiency() *PercentType {
	if x != nil {
		return x.BrewhouseEfficiency
	}
	return nil
}

var File_beerproto_v1_equipment_proto protoreflect.FileDescriptor

var file_beerproto_v1_equipment_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x62, 0x65,
	0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xe7, 0x07, 0x0a, 0x11, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x45,
	0x0a, 0x12, 0x62, 0x6f, 0x69, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f,
	0x68, 0x6f, 0x75, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x65, 0x65,
	0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x62, 0x6f, 0x69, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x72, 0x48, 0x6f, 0x75, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x45, 0x0a, 0x04, 0x66, 0x6f, 0x72,
	0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x42, 0x61, 0x73, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x04, 0x66, 0x6f, 0x72, 0x6d,
	0x12, 0x4b, 0x0a, 0x15, 0x64, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x70,
	0x65, 0x72, 0x5f, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x12, 0x64, 0x72, 0x61, 0x69, 0x6e,
	0x52, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x12, 0x43, 0x0a,
	0x0d, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x5f, 0x68, 0x65, 0x61, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x48, 0x65, 0x61, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x48, 0x65,
	0x61, 0x74, 0x12, 0x54, 0x0a, 0x15, 0x67, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x61, 0x62, 0x73, 0x6f,
	0x72, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x13, 0x67, 0x72, 0x61, 0x69, 0x6e, 0x41, 0x62, 0x73, 0x6f, 0x72, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x0e,
	0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d,
	0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x2e, 0x0a,
	0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x73,
	0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x2c, 0x0a,
	0x04, 0x6c, 0x6f, 0x73, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x65,
	0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x6c, 0x6f, 0x73, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x65,
	0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x65, 0x66, 0x66, 0x69,
	0x63, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x22, 0xb3, 0x02, 0x0a, 0x11, 0x45, 0x71, 0x75, 0x69, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x73, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x23, 0x0a, 0x1f,
	0x45, 0x51, 0x55, 0x49, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x46,
	0x4f, 0x52, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x51, 0x55, 0x49, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x42,
	0x41, 0x53, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x48, 0x4c, 0x54, 0x10, 0x01, 0x12, 0x20,
	0x0a, 0x1c, 0x45, 0x51, 0x55, 0x49, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x42, 0x41, 0x53, 0x45,
	0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x4d, 0x41, 0x53, 0x48, 0x5f, 0x54, 0x55, 0x4e, 0x10, 0x02,
	0x12, 0x22, 0x0a, 0x1e, 0x45, 0x51, 0x55, 0x49, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x42, 0x41,
	0x53, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x4c, 0x41, 0x55, 0x54, 0x45, 0x52, 0x5f, 0x54,
	0x55, 0x4e, 0x10, 0x03, 0x12, 0x23, 0x0a, 0x1f, 0x45, 0x51, 0x55, 0x49, 0x50, 0x4d, 0x45, 0x4e,
	0x54, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x42, 0x52, 0x45, 0x57,
	0x5f, 0x4b, 0x45, 0x54, 0x54, 0x4c, 0x45, 0x10, 0x04, 0x12, 0x21, 0x0a, 0x1d, 0x45, 0x51, 0x55,
	0x49, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x4d,
	0x5f, 0x46, 0x45, 0x52, 0x4d, 0x45, 0x4e, 0x54, 0x45, 0x52, 0x10, 0x05, 0x12, 0x24, 0x0a, 0x20,
	0x45, 0x51, 0x55, 0x49, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x46,
	0x4f, 0x52, 0x4d, 0x5f, 0x41, 0x47, 0x49, 0x4e, 0x47, 0x5f, 0x56, 0x45, 0x53, 0x53, 0x45, 0x4c,
	0x10, 0x06, 0x12, 0x28, 0x0a, 0x24, 0x45, 0x51, 0x55, 0x49, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x42, 0x41, 0x53, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x50, 0x41, 0x43, 0x4b, 0x41, 0x47,
	0x49, 0x4e, 0x47, 0x5f, 0x56, 0x45, 0x53, 0x53, 0x45, 0x4c, 0x10, 0x07, 0x22, 0xcb, 0x01, 0x0a,
	0x0d, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x48, 0x0a, 0x0f, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x65,
	0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x65, 0x71,
	0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x4c, 0x0a, 0x14,
	0x62, 0x72, 0x65, 0x77, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x65, 0x66, 0x66, 0x69, 0x63, 0x69,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x65, 0x65,
	0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x13, 0x62, 0x72, 0x65, 0x77, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x45, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x42, 0xaf, 0x01, 0x0a, 0x10, 0x63,
	0x6f, 0x6d, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x42,
	0x0e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65,
	0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x5f, 0x67, 0x6f, 0x2f, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76,
	0x31, 0x3b, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x42, 0x58, 0x58, 0xaa, 0x02, 0x0c, 0x42, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x0c, 0x42, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x18, 0x42, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x42,
	0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_beerproto_v1_equipment_proto_rawDescOnce sync.Once
	file_beerproto_v1_equipment_proto_rawDescData = file_beerproto_v1_equipment_proto_rawDesc
)

func file_beerproto_v1_equipment_proto_rawDescGZIP() []byte {
	file_beerproto_v1_equipment_proto_rawDescOnce.Do(func() {
		file_beerproto_v1_equipment_proto_rawDescData = protoimpl.X.CompressGZIP(file_beerproto_v1_equipment_proto_rawDescData)
	})
	return file_beerproto_v1_equipment_proto_rawDescData
}

var file_beerproto_v1_equipment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_beerproto_v1_equipment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_beerproto_v1_equipment_proto_goTypes = []interface{}{
	(EquipmentItemType_EquipmentBaseForm)(0), // 0: beerproto.v1.EquipmentItemType.EquipmentBaseForm
	(*EquipmentItemType)(nil),                // 1: beerproto.v1.EquipmentItemType
	(*EquipmentType)(nil),                    // 2: beerproto.v1.EquipmentType
	(*VolumeType)(nil),                       // 3: beerproto.v1.VolumeType
	(*SpecificHeatType)(nil),                 // 4: beerproto.v1.SpecificHeatType
	(*SpecificVolumeType)(nil),               // 5: beerproto.v1.SpecificVolumeType
	(*MassType)(nil),                         // 6: beerproto.v1.MassType
	(*PercentType)(nil),                      // 7: beerproto.v1.PercentType
}
var file_beerproto_v1_equipment_proto_depIdxs = []int32{
	3,  // 0: beerproto.v1.EquipmentItemType.boil_rate_per_hour:type_name -> beerproto.v1.VolumeType
	0,  // 1: beerproto.v1.EquipmentItemType.form:type_name -> beerproto.v1.EquipmentItemType.EquipmentBaseForm
	3,  // 2: beerproto.v1.EquipmentItemType.drain_rate_per_minute:type_name -> beerproto.v1.VolumeType
	4,  // 3: beerproto.v1.EquipmentItemType.specific_heat:type_name -> beerproto.v1.SpecificHeatType
	5,  // 4: beerproto.v1.EquipmentItemType.grain_absorption_rate:type_name -> beerproto.v1.SpecificVolumeType
	3,  // 5: beerproto.v1.EquipmentItemType.maximum_volume:type_name -> beerproto.v1.VolumeType
	6,  // 6: beerproto.v1.EquipmentItemType.weight:type_name -> beerproto.v1.MassType
	3,  // 7: beerproto.v1.EquipmentItemType.loss:type_name -> beerproto.v1.VolumeType
	7,  // 8: beerproto.v1.EquipmentItemType.efficiency:type_name -> beerproto.v1.PercentType
	1,  // 9: beerproto.v1.EquipmentType.equipment_items:type_name -> beerproto.v1.EquipmentItemType
	7,  // 10: beerproto.v1.EquipmentType.brewhouse_efficiency:type_name -> beerproto.v1.PercentType
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_beerproto_v1_equipment_proto_init() }
func file_beerproto_v1_equipment_proto_init() {
	if File_beerproto_v1_equipment_proto != nil {
		return
	}
	file_beerproto_v1_measureable_units_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_beerproto_v1_equipment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EquipmentItemType); i {
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
		file_beerproto_v1_equipment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EquipmentType); i {
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
			RawDescriptor: file_beerproto_v1_equipment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_beerproto_v1_equipment_proto_goTypes,
		DependencyIndexes: file_beerproto_v1_equipment_proto_depIdxs,
		EnumInfos:         file_beerproto_v1_equipment_proto_enumTypes,
		MessageInfos:      file_beerproto_v1_equipment_proto_msgTypes,
	}.Build()
	File_beerproto_v1_equipment_proto = out.File
	file_beerproto_v1_equipment_proto_rawDesc = nil
	file_beerproto_v1_equipment_proto_goTypes = nil
	file_beerproto_v1_equipment_proto_depIdxs = nil
}
