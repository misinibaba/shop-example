// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.2
// source: google/ads/googleads/v3/enums/income_range_type.proto

package enums

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The type of demographic income ranges (e.g. between 0% to 50%).
type IncomeRangeTypeEnum_IncomeRangeType int32

const (
	// Not specified.
	IncomeRangeTypeEnum_UNSPECIFIED IncomeRangeTypeEnum_IncomeRangeType = 0
	// Used for return value only. Represents value unknown in this version.
	IncomeRangeTypeEnum_UNKNOWN IncomeRangeTypeEnum_IncomeRangeType = 1
	// 0%-50%.
	IncomeRangeTypeEnum_INCOME_RANGE_0_50 IncomeRangeTypeEnum_IncomeRangeType = 510001
	// 50% to 60%.
	IncomeRangeTypeEnum_INCOME_RANGE_50_60 IncomeRangeTypeEnum_IncomeRangeType = 510002
	// 60% to 70%.
	IncomeRangeTypeEnum_INCOME_RANGE_60_70 IncomeRangeTypeEnum_IncomeRangeType = 510003
	// 70% to 80%.
	IncomeRangeTypeEnum_INCOME_RANGE_70_80 IncomeRangeTypeEnum_IncomeRangeType = 510004
	// 80% to 90%.
	IncomeRangeTypeEnum_INCOME_RANGE_80_90 IncomeRangeTypeEnum_IncomeRangeType = 510005
	// Greater than 90%.
	IncomeRangeTypeEnum_INCOME_RANGE_90_UP IncomeRangeTypeEnum_IncomeRangeType = 510006
	// Undetermined income range.
	IncomeRangeTypeEnum_INCOME_RANGE_UNDETERMINED IncomeRangeTypeEnum_IncomeRangeType = 510000
)

// Enum value maps for IncomeRangeTypeEnum_IncomeRangeType.
var (
	IncomeRangeTypeEnum_IncomeRangeType_name = map[int32]string{
		0:      "UNSPECIFIED",
		1:      "UNKNOWN",
		510001: "INCOME_RANGE_0_50",
		510002: "INCOME_RANGE_50_60",
		510003: "INCOME_RANGE_60_70",
		510004: "INCOME_RANGE_70_80",
		510005: "INCOME_RANGE_80_90",
		510006: "INCOME_RANGE_90_UP",
		510000: "INCOME_RANGE_UNDETERMINED",
	}
	IncomeRangeTypeEnum_IncomeRangeType_value = map[string]int32{
		"UNSPECIFIED":               0,
		"UNKNOWN":                   1,
		"INCOME_RANGE_0_50":         510001,
		"INCOME_RANGE_50_60":        510002,
		"INCOME_RANGE_60_70":        510003,
		"INCOME_RANGE_70_80":        510004,
		"INCOME_RANGE_80_90":        510005,
		"INCOME_RANGE_90_UP":        510006,
		"INCOME_RANGE_UNDETERMINED": 510000,
	}
)

func (x IncomeRangeTypeEnum_IncomeRangeType) Enum() *IncomeRangeTypeEnum_IncomeRangeType {
	p := new(IncomeRangeTypeEnum_IncomeRangeType)
	*p = x
	return p
}

func (x IncomeRangeTypeEnum_IncomeRangeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IncomeRangeTypeEnum_IncomeRangeType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v3_enums_income_range_type_proto_enumTypes[0].Descriptor()
}

func (IncomeRangeTypeEnum_IncomeRangeType) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v3_enums_income_range_type_proto_enumTypes[0]
}

func (x IncomeRangeTypeEnum_IncomeRangeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IncomeRangeTypeEnum_IncomeRangeType.Descriptor instead.
func (IncomeRangeTypeEnum_IncomeRangeType) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v3_enums_income_range_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing the type of demographic income ranges.
type IncomeRangeTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IncomeRangeTypeEnum) Reset() {
	*x = IncomeRangeTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v3_enums_income_range_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncomeRangeTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncomeRangeTypeEnum) ProtoMessage() {}

func (x *IncomeRangeTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v3_enums_income_range_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncomeRangeTypeEnum.ProtoReflect.Descriptor instead.
func (*IncomeRangeTypeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v3_enums_income_range_type_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v3_enums_income_range_type_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v3_enums_income_range_type_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x33,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x02, 0x0a, 0x13, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xeb, 0x01, 0x0a,
	0x0f, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x17,
	0x0a, 0x11, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x30,
	0x5f, 0x35, 0x30, 0x10, 0xb1, 0x90, 0x1f, 0x12, 0x18, 0x0a, 0x12, 0x49, 0x4e, 0x43, 0x4f, 0x4d,
	0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x35, 0x30, 0x5f, 0x36, 0x30, 0x10, 0xb2, 0x90,
	0x1f, 0x12, 0x18, 0x0a, 0x12, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47,
	0x45, 0x5f, 0x36, 0x30, 0x5f, 0x37, 0x30, 0x10, 0xb3, 0x90, 0x1f, 0x12, 0x18, 0x0a, 0x12, 0x49,
	0x4e, 0x43, 0x4f, 0x4d, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x37, 0x30, 0x5f, 0x38,
	0x30, 0x10, 0xb4, 0x90, 0x1f, 0x12, 0x18, 0x0a, 0x12, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x45, 0x5f,
	0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x38, 0x30, 0x5f, 0x39, 0x30, 0x10, 0xb5, 0x90, 0x1f, 0x12,
	0x18, 0x0a, 0x12, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f,
	0x39, 0x30, 0x5f, 0x55, 0x50, 0x10, 0xb6, 0x90, 0x1f, 0x12, 0x1f, 0x0a, 0x19, 0x49, 0x4e, 0x43,
	0x4f, 0x4d, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x54, 0x45,
	0x52, 0x4d, 0x49, 0x4e, 0x45, 0x44, 0x10, 0xb0, 0x90, 0x1f, 0x42, 0xe9, 0x01, 0x0a, 0x21, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x42, 0x14, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x33,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47,
	0x41, 0x41, 0xaa, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x33, 0x2e, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0xca, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x33, 0x5c, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x33, 0x3a,
	0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v3_enums_income_range_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v3_enums_income_range_type_proto_rawDescData = file_google_ads_googleads_v3_enums_income_range_type_proto_rawDesc
)

func file_google_ads_googleads_v3_enums_income_range_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v3_enums_income_range_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v3_enums_income_range_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v3_enums_income_range_type_proto_rawDescData)
	})
	return file_google_ads_googleads_v3_enums_income_range_type_proto_rawDescData
}

var file_google_ads_googleads_v3_enums_income_range_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v3_enums_income_range_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v3_enums_income_range_type_proto_goTypes = []interface{}{
	(IncomeRangeTypeEnum_IncomeRangeType)(0), // 0: google.ads.googleads.v3.enums.IncomeRangeTypeEnum.IncomeRangeType
	(*IncomeRangeTypeEnum)(nil),              // 1: google.ads.googleads.v3.enums.IncomeRangeTypeEnum
}
var file_google_ads_googleads_v3_enums_income_range_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v3_enums_income_range_type_proto_init() }
func file_google_ads_googleads_v3_enums_income_range_type_proto_init() {
	if File_google_ads_googleads_v3_enums_income_range_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v3_enums_income_range_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncomeRangeTypeEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v3_enums_income_range_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v3_enums_income_range_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v3_enums_income_range_type_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v3_enums_income_range_type_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v3_enums_income_range_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v3_enums_income_range_type_proto = out.File
	file_google_ads_googleads_v3_enums_income_range_type_proto_rawDesc = nil
	file_google_ads_googleads_v3_enums_income_range_type_proto_goTypes = nil
	file_google_ads_googleads_v3_enums_income_range_type_proto_depIdxs = nil
}
