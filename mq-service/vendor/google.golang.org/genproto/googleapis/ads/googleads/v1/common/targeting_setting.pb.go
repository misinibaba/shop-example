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
// source: google/ads/googleads/v1/common/targeting_setting.proto

package common

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
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

// Settings for the targeting-related features, at the campaign and ad group
// levels. For more details about the targeting setting, visit
// https://support.google.com/google-ads/answer/7365594
type TargetingSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The per-targeting-dimension setting to restrict the reach of your campaign
	// or ad group.
	TargetRestrictions []*TargetRestriction `protobuf:"bytes,1,rep,name=target_restrictions,json=targetRestrictions,proto3" json:"target_restrictions,omitempty"`
}

func (x *TargetingSetting) Reset() {
	*x = TargetingSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v1_common_targeting_setting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetingSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetingSetting) ProtoMessage() {}

func (x *TargetingSetting) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v1_common_targeting_setting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetingSetting.ProtoReflect.Descriptor instead.
func (*TargetingSetting) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v1_common_targeting_setting_proto_rawDescGZIP(), []int{0}
}

func (x *TargetingSetting) GetTargetRestrictions() []*TargetRestriction {
	if x != nil {
		return x.TargetRestrictions
	}
	return nil
}

// The list of per-targeting-dimension targeting settings.
type TargetRestriction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The targeting dimension that these settings apply to.
	TargetingDimension enums.TargetingDimensionEnum_TargetingDimension `protobuf:"varint,1,opt,name=targeting_dimension,json=targetingDimension,proto3,enum=google.ads.googleads.v1.enums.TargetingDimensionEnum_TargetingDimension" json:"targeting_dimension,omitempty"`
	// Indicates whether to restrict your ads to show only for the criteria you
	// have selected for this targeting_dimension, or to target all values for
	// this targeting_dimension and show ads based on your targeting in other
	// TargetingDimensions. A value of `true` means that these criteria will only
	// apply bid modifiers, and not affect targeting. A value of `false` means
	// that these criteria will restrict targeting as well as applying bid
	// modifiers.
	BidOnly *wrappers.BoolValue `protobuf:"bytes,2,opt,name=bid_only,json=bidOnly,proto3" json:"bid_only,omitempty"`
}

func (x *TargetRestriction) Reset() {
	*x = TargetRestriction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v1_common_targeting_setting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetRestriction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetRestriction) ProtoMessage() {}

func (x *TargetRestriction) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v1_common_targeting_setting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetRestriction.ProtoReflect.Descriptor instead.
func (*TargetRestriction) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v1_common_targeting_setting_proto_rawDescGZIP(), []int{1}
}

func (x *TargetRestriction) GetTargetingDimension() enums.TargetingDimensionEnum_TargetingDimension {
	if x != nil {
		return x.TargetingDimension
	}
	return enums.TargetingDimensionEnum_UNSPECIFIED
}

func (x *TargetRestriction) GetBidOnly() *wrappers.BoolValue {
	if x != nil {
		return x.BidOnly
	}
	return nil
}

var File_google_ads_googleads_v1_common_targeting_setting_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v1_common_targeting_setting_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x5f, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x76, 0x0a, 0x10, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x62, 0x0a, 0x13, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x65,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x12, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x74, 0x72,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xc5, 0x01, 0x0a, 0x11, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x79, 0x0a,
	0x13, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x69, 0x6d, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x48, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x75,
	0x6d, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x6d, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x44,
	0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x08, 0x62, 0x69, 0x64, 0x5f,
	0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x62, 0x69, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x42,
	0xf0, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x15, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x44, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xca, 0x02, 0x1e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xea, 0x02, 0x22,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v1_common_targeting_setting_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v1_common_targeting_setting_proto_rawDescData = file_google_ads_googleads_v1_common_targeting_setting_proto_rawDesc
)

func file_google_ads_googleads_v1_common_targeting_setting_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v1_common_targeting_setting_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v1_common_targeting_setting_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v1_common_targeting_setting_proto_rawDescData)
	})
	return file_google_ads_googleads_v1_common_targeting_setting_proto_rawDescData
}

var file_google_ads_googleads_v1_common_targeting_setting_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_ads_googleads_v1_common_targeting_setting_proto_goTypes = []interface{}{
	(*TargetingSetting)(nil),                             // 0: google.ads.googleads.v1.common.TargetingSetting
	(*TargetRestriction)(nil),                            // 1: google.ads.googleads.v1.common.TargetRestriction
	(enums.TargetingDimensionEnum_TargetingDimension)(0), // 2: google.ads.googleads.v1.enums.TargetingDimensionEnum.TargetingDimension
	(*wrappers.BoolValue)(nil),                           // 3: google.protobuf.BoolValue
}
var file_google_ads_googleads_v1_common_targeting_setting_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v1.common.TargetingSetting.target_restrictions:type_name -> google.ads.googleads.v1.common.TargetRestriction
	2, // 1: google.ads.googleads.v1.common.TargetRestriction.targeting_dimension:type_name -> google.ads.googleads.v1.enums.TargetingDimensionEnum.TargetingDimension
	3, // 2: google.ads.googleads.v1.common.TargetRestriction.bid_only:type_name -> google.protobuf.BoolValue
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v1_common_targeting_setting_proto_init() }
func file_google_ads_googleads_v1_common_targeting_setting_proto_init() {
	if File_google_ads_googleads_v1_common_targeting_setting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v1_common_targeting_setting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetingSetting); i {
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
		file_google_ads_googleads_v1_common_targeting_setting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetRestriction); i {
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
			RawDescriptor: file_google_ads_googleads_v1_common_targeting_setting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v1_common_targeting_setting_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v1_common_targeting_setting_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v1_common_targeting_setting_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v1_common_targeting_setting_proto = out.File
	file_google_ads_googleads_v1_common_targeting_setting_proto_rawDesc = nil
	file_google_ads_googleads_v1_common_targeting_setting_proto_goTypes = nil
	file_google_ads_googleads_v1_common_targeting_setting_proto_depIdxs = nil
}
