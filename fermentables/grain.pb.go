//*
// BeerProto
//
// Another beer format, written in protocol buffer.
//
// Copyright (c) 2020 Ross Merrigan

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: grain.proto

package fermentables

import (
	beerproto_go "github.com/beerproto/beerproto_go"
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

type GrainType_StandardType int32

const (
	GrainType_NULL_StandardType GrainType_StandardType = 0
	// American Society of Brewing Chemists
	GrainType_ASBC GrainType_StandardType = 1
	// European Brewery Convention
	GrainType_EBC GrainType_StandardType = 2
	// ION
	GrainType_ION GrainType_StandardType = 3
)

// Enum value maps for GrainType_StandardType.
var (
	GrainType_StandardType_name = map[int32]string{
		0: "NULL_StandardType",
		1: "ASBC",
		2: "EBC",
		3: "ION",
	}
	GrainType_StandardType_value = map[string]int32{
		"NULL_StandardType": 0,
		"ASBC":              1,
		"EBC":               2,
		"ION":               3,
	}
)

func (x GrainType_StandardType) Enum() *GrainType_StandardType {
	p := new(GrainType_StandardType)
	*p = x
	return p
}

func (x GrainType_StandardType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GrainType_StandardType) Descriptor() protoreflect.EnumDescriptor {
	return file_grain_proto_enumTypes[0].Descriptor()
}

func (GrainType_StandardType) Type() protoreflect.EnumType {
	return &file_grain_proto_enumTypes[0]
}

func (x GrainType_StandardType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GrainType_StandardType.Descriptor instead.
func (GrainType_StandardType) EnumDescriptor() ([]byte, []int) {
	return file_grain_proto_rawDescGZIP(), []int{0, 0}
}

type GrainType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string                         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string                         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Producer string                         `protobuf:"bytes,3,opt,name=producer,proto3" json:"producer,omitempty"`
	Moisture *beerproto_go.PercentRangeType `protobuf:"bytes,4,opt,name=moisture,proto3" json:"moisture,omitempty"`
	// Percentage yield compared to succrose of a fine grind. eg 80%
	FineGrind *beerproto_go.PercentRangeType `protobuf:"bytes,5,opt,name=fine_grind,json=fineGrind,proto3" json:"fine_grind,omitempty"`
	// Percentage yield compared to succrose of a coarse grind. eg 78%
	CoarseGrind *beerproto_go.PercentRangeType `protobuf:"bytes,6,opt,name=coarse_grind,json=coarseGrind,proto3" json:"coarse_grind,omitempty"`
	// The difference between fine and coarse grind, a difference more than 2 percent can indicate a protein or step mash may be desirable. eg 2%
	FineCoarseDifference *beerproto_go.PercentRangeType `protobuf:"bytes,7,opt,name=fine_coarse_difference,json=fineCoarseDifference,proto3" json:"fine_coarse_difference,omitempty"`
	// The potential yield as a percent
	Yield *beerproto_go.PercentRangeType `protobuf:"bytes,8,opt,name=yield,proto3" json:"yield,omitempty"`
	// The potential yield of the fermentable ingredient for 1 lb of grain mashed in 1 gallon of water. eg 1.037
	Potential *beerproto_go.GravityRangeType `protobuf:"bytes,9,opt,name=potential,proto3" json:"potential,omitempty"`
	// Diastatic power is a measurement of malted grains enzymatic content. A value of 35 Lintner is needed to self convert, while a value of 100 or more is desirable.
	DiastaticPower  *beerproto_go.DiastaticPowerRangeType `protobuf:"bytes,10,opt,name=diastatic_power,json=diastaticPower,proto3" json:"diastatic_power,omitempty"`
	Protein         *beerproto_go.PercentType             `protobuf:"bytes,11,opt,name=protein,proto3" json:"protein,omitempty"`
	SolubleProtein  *beerproto_go.PercentType             `protobuf:"bytes,12,opt,name=soluble_protein,json=solubleProtein,proto3" json:"soluble_protein,omitempty"`
	TotalNitrogen   *beerproto_go.PercentType             `protobuf:"bytes,13,opt,name=total_nitrogen,json=totalNitrogen,proto3" json:"total_nitrogen,omitempty"`
	SolubleNitrogen *beerproto_go.PercentType             `protobuf:"bytes,14,opt,name=soluble_nitrogen,json=solubleNitrogen,proto3" json:"soluble_nitrogen,omitempty"`
	Maximum         *beerproto_go.PercentType             `protobuf:"bytes,15,opt,name=maximum,proto3" json:"maximum,omitempty"`
	// Friability is the measure of a malts ability to crumble during the crush, and is used as an indicator for easy gelatinization of the grain and starches, as well as modification of the malt. Value of 85% of higher indicates a well modified malt and is suitable for single step mashes. Lower values may require a step mash.
	Friability *beerproto_go.PercentRangeType `protobuf:"bytes,16,opt,name=friability,proto3" json:"friability,omitempty"`
	Color      *beerproto_go.ColorRangeType   `protobuf:"bytes,17,opt,name=color,proto3" json:"color,omitempty"`
	// Values of 180 or more may suggest a glucan rest and avoiding fly sparging.
	BetaGlucan *beerproto_go.ConcentrationRangeType `protobuf:"bytes,18,opt,name=beta_glucan,json=betaGlucan,proto3" json:"beta_glucan,omitempty"`
	// Free Amino Nitrogen (FAN) is a critical yeast nutrient. Typical values for base malt is 170.
	Fan *beerproto_go.PercentRangeType `protobuf:"bytes,19,opt,name=fan,proto3" json:"fan,omitempty"`
	// Where diastatic power gives the total amount of all enzymes, alpha amylase, also known as dextrinizing units, refers to only the total amount of alpa amylase in the malted grain. A value of 25-50 is desirable for base malt.
	AlphaAmylase     *beerproto_go.TimeType      `protobuf:"bytes,20,opt,name=alpha_amylase,json=alphaAmylase,proto3" json:"alpha_amylase,omitempty"`
	Saccharification *beerproto_go.TimeRangeType `protobuf:"bytes,21,opt,name=saccharification,proto3" json:"saccharification,omitempty"`
	// The measure of wort viscosity, typically associated with the breakdown of beta-glucans. The higher the viscosity, the greater the need for a glucan rest and the less suitable for a fly sparge.
	Viscosity *beerproto_go.ViscosityRangeType `protobuf:"bytes,22,opt,name=viscosity,proto3" json:"viscosity,omitempty"`
	// The amount of DMS precursors, namely S-methyl methionine (SMM) and dimethyl sulfoxide (DMSO) in the malt which convert to dimethyl sulfide (DMS).
	DmsP *beerproto_go.ConcentrationType `protobuf:"bytes,23,opt,name=dms_p,json=dmsP,proto3" json:"dms_p,omitempty"`
	// The Kolbach Index, also known as soluble to total ratio of nitrogen or protein, is used to indcate the degree of malt modification. A value above 35% is desired for simple single infusion mashing, undermodified malt may require multiple step mashes or decoction.
	KolbachIndex float64 `protobuf:"fixed64,24,opt,name=kolbach_index,json=kolbachIndex,proto3" json:"kolbach_index,omitempty"`
	// The pH of the resultant wort for 1 lb of grain mashed in 1 gallon of distilled water. Used in many water chemistry / mash pH prediction software.
	DiPh       *beerproto_go.AcidityType `protobuf:"bytes,25,opt,name=di_ph,json=diPh,proto3" json:"di_ph,omitempty"`
	GrainGroup beerproto_go.GrainGroup   `protobuf:"varint,26,opt,name=grain_group,json=grainGroup,proto3,enum=beerproto.GrainGroup" json:"grain_group,omitempty"`
	Country    string                    `protobuf:"bytes,27,opt,name=country,proto3" json:"country,omitempty"`
	Standard   GrainType_StandardType    `protobuf:"varint,28,opt,name=standard,proto3,enum=fermentables.GrainType_StandardType" json:"standard,omitempty"`
	Notes      string                    `protobuf:"bytes,29,opt,name=notes,proto3" json:"notes,omitempty"`
}

func (x *GrainType) Reset() {
	*x = GrainType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrainType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrainType) ProtoMessage() {}

func (x *GrainType) ProtoReflect() protoreflect.Message {
	mi := &file_grain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrainType.ProtoReflect.Descriptor instead.
func (*GrainType) Descriptor() ([]byte, []int) {
	return file_grain_proto_rawDescGZIP(), []int{0}
}

func (x *GrainType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GrainType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GrainType) GetProducer() string {
	if x != nil {
		return x.Producer
	}
	return ""
}

func (x *GrainType) GetMoisture() *beerproto_go.PercentRangeType {
	if x != nil {
		return x.Moisture
	}
	return nil
}

func (x *GrainType) GetFineGrind() *beerproto_go.PercentRangeType {
	if x != nil {
		return x.FineGrind
	}
	return nil
}

func (x *GrainType) GetCoarseGrind() *beerproto_go.PercentRangeType {
	if x != nil {
		return x.CoarseGrind
	}
	return nil
}

func (x *GrainType) GetFineCoarseDifference() *beerproto_go.PercentRangeType {
	if x != nil {
		return x.FineCoarseDifference
	}
	return nil
}

func (x *GrainType) GetYield() *beerproto_go.PercentRangeType {
	if x != nil {
		return x.Yield
	}
	return nil
}

func (x *GrainType) GetPotential() *beerproto_go.GravityRangeType {
	if x != nil {
		return x.Potential
	}
	return nil
}

func (x *GrainType) GetDiastaticPower() *beerproto_go.DiastaticPowerRangeType {
	if x != nil {
		return x.DiastaticPower
	}
	return nil
}

func (x *GrainType) GetProtein() *beerproto_go.PercentType {
	if x != nil {
		return x.Protein
	}
	return nil
}

func (x *GrainType) GetSolubleProtein() *beerproto_go.PercentType {
	if x != nil {
		return x.SolubleProtein
	}
	return nil
}

func (x *GrainType) GetTotalNitrogen() *beerproto_go.PercentType {
	if x != nil {
		return x.TotalNitrogen
	}
	return nil
}

func (x *GrainType) GetSolubleNitrogen() *beerproto_go.PercentType {
	if x != nil {
		return x.SolubleNitrogen
	}
	return nil
}

func (x *GrainType) GetMaximum() *beerproto_go.PercentType {
	if x != nil {
		return x.Maximum
	}
	return nil
}

func (x *GrainType) GetFriability() *beerproto_go.PercentRangeType {
	if x != nil {
		return x.Friability
	}
	return nil
}

func (x *GrainType) GetColor() *beerproto_go.ColorRangeType {
	if x != nil {
		return x.Color
	}
	return nil
}

func (x *GrainType) GetBetaGlucan() *beerproto_go.ConcentrationRangeType {
	if x != nil {
		return x.BetaGlucan
	}
	return nil
}

func (x *GrainType) GetFan() *beerproto_go.PercentRangeType {
	if x != nil {
		return x.Fan
	}
	return nil
}

func (x *GrainType) GetAlphaAmylase() *beerproto_go.TimeType {
	if x != nil {
		return x.AlphaAmylase
	}
	return nil
}

func (x *GrainType) GetSaccharification() *beerproto_go.TimeRangeType {
	if x != nil {
		return x.Saccharification
	}
	return nil
}

func (x *GrainType) GetViscosity() *beerproto_go.ViscosityRangeType {
	if x != nil {
		return x.Viscosity
	}
	return nil
}

func (x *GrainType) GetDmsP() *beerproto_go.ConcentrationType {
	if x != nil {
		return x.DmsP
	}
	return nil
}

func (x *GrainType) GetKolbachIndex() float64 {
	if x != nil {
		return x.KolbachIndex
	}
	return 0
}

func (x *GrainType) GetDiPh() *beerproto_go.AcidityType {
	if x != nil {
		return x.DiPh
	}
	return nil
}

func (x *GrainType) GetGrainGroup() beerproto_go.GrainGroup {
	if x != nil {
		return x.GrainGroup
	}
	return beerproto_go.GrainGroup(0)
}

func (x *GrainType) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *GrainType) GetStandard() GrainType_StandardType {
	if x != nil {
		return x.Standard
	}
	return GrainType_NULL_StandardType
}

func (x *GrainType) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

var File_grain_proto protoreflect.FileDescriptor

var file_grain_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x72, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x66,
	0x65, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x1a, 0x0a, 0x62, 0x65, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x0c, 0x0a, 0x09, 0x47, 0x72, 0x61, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x08, 0x6d, 0x6f, 0x69, 0x73, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x6d, 0x6f, 0x69, 0x73, 0x74, 0x75, 0x72, 0x65, 0x12, 0x3a,
	0x0a, 0x0a, 0x66, 0x69, 0x6e, 0x65, 0x5f, 0x67, 0x72, 0x69, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x09, 0x66, 0x69, 0x6e, 0x65, 0x47, 0x72, 0x69, 0x6e, 0x64, 0x12, 0x3e, 0x0a, 0x0c, 0x63, 0x6f,
	0x61, 0x72, 0x73, 0x65, 0x5f, 0x67, 0x72, 0x69, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x63,
	0x6f, 0x61, 0x72, 0x73, 0x65, 0x47, 0x72, 0x69, 0x6e, 0x64, 0x12, 0x51, 0x0a, 0x16, 0x66, 0x69,
	0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x61, 0x72, 0x73, 0x65, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x65, 0x65,
	0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x14, 0x66, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x61,
	0x72, 0x73, 0x65, 0x44, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x31, 0x0a,
	0x05, 0x79, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62,
	0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x79, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x39, 0x0a, 0x09, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x72, 0x61, 0x76, 0x69, 0x74, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x09, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x4b, 0x0a, 0x0f, 0x64,
	0x69, 0x61, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x44, 0x69, 0x61, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x64, 0x69, 0x61, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x74,
	0x65, 0x69, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x65, 0x65, 0x72,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x12, 0x3f, 0x0a, 0x0f, 0x73, 0x6f,
	0x6c, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x73, 0x6f, 0x6c,
	0x75, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x12, 0x3d, 0x0a, 0x0e, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6e, 0x69, 0x74, 0x72, 0x6f, 0x67, 0x65, 0x6e, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x4e, 0x69, 0x74, 0x72, 0x6f, 0x67, 0x65, 0x6e, 0x12, 0x41, 0x0a, 0x10, 0x73, 0x6f,
	0x6c, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x69, 0x74, 0x72, 0x6f, 0x67, 0x65, 0x6e, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x73, 0x6f,
	0x6c, 0x75, 0x62, 0x6c, 0x65, 0x4e, 0x69, 0x74, 0x72, 0x6f, 0x67, 0x65, 0x6e, 0x12, 0x30, 0x0a,
	0x07, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x12,
	0x3b, 0x0a, 0x0a, 0x66, 0x72, 0x69, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0a, 0x66, 0x72, 0x69, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x2f, 0x0a, 0x05,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x65,
	0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x42, 0x0a,
	0x0b, 0x62, 0x65, 0x74, 0x61, 0x5f, 0x67, 0x6c, 0x75, 0x63, 0x61, 0x6e, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6f, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x62, 0x65, 0x74, 0x61, 0x47, 0x6c, 0x75, 0x63, 0x61,
	0x6e, 0x12, 0x2d, 0x0a, 0x03, 0x66, 0x61, 0x6e, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x03, 0x66, 0x61, 0x6e,
	0x12, 0x38, 0x0a, 0x0d, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5f, 0x61, 0x6d, 0x79, 0x6c, 0x61, 0x73,
	0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x41, 0x6d, 0x79, 0x6c, 0x61, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x10, 0x73, 0x61,
	0x63, 0x63, 0x68, 0x61, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x10,
	0x73, 0x61, 0x63, 0x63, 0x68, 0x61, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x3b, 0x0a, 0x09, 0x76, 0x69, 0x73, 0x63, 0x6f, 0x73, 0x69, 0x74, 0x79, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x56, 0x69, 0x73, 0x63, 0x6f, 0x73, 0x69, 0x74, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x09, 0x76, 0x69, 0x73, 0x63, 0x6f, 0x73, 0x69, 0x74, 0x79, 0x12, 0x31, 0x0a,
	0x05, 0x64, 0x6d, 0x73, 0x5f, 0x70, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62,
	0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x6e, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x64, 0x6d, 0x73, 0x50,
	0x12, 0x23, 0x0a, 0x0d, 0x6b, 0x6f, 0x6c, 0x62, 0x61, 0x63, 0x68, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x18, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x6b, 0x6f, 0x6c, 0x62, 0x61, 0x63, 0x68,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2b, 0x0a, 0x05, 0x64, 0x69, 0x5f, 0x70, 0x68, 0x18, 0x19,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x63, 0x69, 0x64, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x64, 0x69,
	0x50, 0x68, 0x12, 0x36, 0x0a, 0x0b, 0x67, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x72, 0x61, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0a,
	0x67, 0x72, 0x61, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x40, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64,
	0x18, 0x1c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x66, 0x65, 0x72, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x2e,
	0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x73, 0x74,
	0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18,
	0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x41, 0x0a, 0x0c,
	0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11,
	0x4e, 0x55, 0x4c, 0x4c, 0x5f, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x41, 0x53, 0x42, 0x43, 0x10, 0x01, 0x12, 0x07, 0x0a,
	0x03, 0x45, 0x42, 0x43, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x42,
	0x83, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x66, 0x65, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x6f, 0x2f, 0x66, 0x65, 0x72,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2f, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03,
	0x42, 0x50, 0x46, 0xaa, 0x02, 0x1f, 0x42, 0x65, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x46, 0x65, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grain_proto_rawDescOnce sync.Once
	file_grain_proto_rawDescData = file_grain_proto_rawDesc
)

func file_grain_proto_rawDescGZIP() []byte {
	file_grain_proto_rawDescOnce.Do(func() {
		file_grain_proto_rawDescData = protoimpl.X.CompressGZIP(file_grain_proto_rawDescData)
	})
	return file_grain_proto_rawDescData
}

var file_grain_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_grain_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_grain_proto_goTypes = []interface{}{
	(GrainType_StandardType)(0),                  // 0: fermentables.GrainType.StandardType
	(*GrainType)(nil),                            // 1: fermentables.GrainType
	(*beerproto_go.PercentRangeType)(nil),        // 2: beerproto.PercentRangeType
	(*beerproto_go.GravityRangeType)(nil),        // 3: beerproto.GravityRangeType
	(*beerproto_go.DiastaticPowerRangeType)(nil), // 4: beerproto.DiastaticPowerRangeType
	(*beerproto_go.PercentType)(nil),             // 5: beerproto.PercentType
	(*beerproto_go.ColorRangeType)(nil),          // 6: beerproto.ColorRangeType
	(*beerproto_go.ConcentrationRangeType)(nil),  // 7: beerproto.ConcentrationRangeType
	(*beerproto_go.TimeType)(nil),                // 8: beerproto.TimeType
	(*beerproto_go.TimeRangeType)(nil),           // 9: beerproto.TimeRangeType
	(*beerproto_go.ViscosityRangeType)(nil),      // 10: beerproto.ViscosityRangeType
	(*beerproto_go.ConcentrationType)(nil),       // 11: beerproto.ConcentrationType
	(*beerproto_go.AcidityType)(nil),             // 12: beerproto.AcidityType
	(beerproto_go.GrainGroup)(0),                 // 13: beerproto.GrainGroup
}
var file_grain_proto_depIdxs = []int32{
	2,  // 0: fermentables.GrainType.moisture:type_name -> beerproto.PercentRangeType
	2,  // 1: fermentables.GrainType.fine_grind:type_name -> beerproto.PercentRangeType
	2,  // 2: fermentables.GrainType.coarse_grind:type_name -> beerproto.PercentRangeType
	2,  // 3: fermentables.GrainType.fine_coarse_difference:type_name -> beerproto.PercentRangeType
	2,  // 4: fermentables.GrainType.yield:type_name -> beerproto.PercentRangeType
	3,  // 5: fermentables.GrainType.potential:type_name -> beerproto.GravityRangeType
	4,  // 6: fermentables.GrainType.diastatic_power:type_name -> beerproto.DiastaticPowerRangeType
	5,  // 7: fermentables.GrainType.protein:type_name -> beerproto.PercentType
	5,  // 8: fermentables.GrainType.soluble_protein:type_name -> beerproto.PercentType
	5,  // 9: fermentables.GrainType.total_nitrogen:type_name -> beerproto.PercentType
	5,  // 10: fermentables.GrainType.soluble_nitrogen:type_name -> beerproto.PercentType
	5,  // 11: fermentables.GrainType.maximum:type_name -> beerproto.PercentType
	2,  // 12: fermentables.GrainType.friability:type_name -> beerproto.PercentRangeType
	6,  // 13: fermentables.GrainType.color:type_name -> beerproto.ColorRangeType
	7,  // 14: fermentables.GrainType.beta_glucan:type_name -> beerproto.ConcentrationRangeType
	2,  // 15: fermentables.GrainType.fan:type_name -> beerproto.PercentRangeType
	8,  // 16: fermentables.GrainType.alpha_amylase:type_name -> beerproto.TimeType
	9,  // 17: fermentables.GrainType.saccharification:type_name -> beerproto.TimeRangeType
	10, // 18: fermentables.GrainType.viscosity:type_name -> beerproto.ViscosityRangeType
	11, // 19: fermentables.GrainType.dms_p:type_name -> beerproto.ConcentrationType
	12, // 20: fermentables.GrainType.di_ph:type_name -> beerproto.AcidityType
	13, // 21: fermentables.GrainType.grain_group:type_name -> beerproto.GrainGroup
	0,  // 22: fermentables.GrainType.standard:type_name -> fermentables.GrainType.StandardType
	23, // [23:23] is the sub-list for method output_type
	23, // [23:23] is the sub-list for method input_type
	23, // [23:23] is the sub-list for extension type_name
	23, // [23:23] is the sub-list for extension extendee
	0,  // [0:23] is the sub-list for field type_name
}

func init() { file_grain_proto_init() }
func file_grain_proto_init() {
	if File_grain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrainType); i {
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
			RawDescriptor: file_grain_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grain_proto_goTypes,
		DependencyIndexes: file_grain_proto_depIdxs,
		EnumInfos:         file_grain_proto_enumTypes,
		MessageInfos:      file_grain_proto_msgTypes,
	}.Build()
	File_grain_proto = out.File
	file_grain_proto_rawDesc = nil
	file_grain_proto_goTypes = nil
	file_grain_proto_depIdxs = nil
}
