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
// source: google/ads/googleads/v2/enums/custom_placeholder_field.proto

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

// Possible values for Custom placeholder fields.
type CustomPlaceholderFieldEnum_CustomPlaceholderField int32

const (
	// Not specified.
	CustomPlaceholderFieldEnum_UNSPECIFIED CustomPlaceholderFieldEnum_CustomPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	CustomPlaceholderFieldEnum_UNKNOWN CustomPlaceholderFieldEnum_CustomPlaceholderField = 1
	// Data Type: STRING. Required. Combination ID and ID2 must be unique per
	// offer.
	CustomPlaceholderFieldEnum_ID CustomPlaceholderFieldEnum_CustomPlaceholderField = 2
	// Data Type: STRING. Combination ID and ID2 must be unique per offer.
	CustomPlaceholderFieldEnum_ID2 CustomPlaceholderFieldEnum_CustomPlaceholderField = 3
	// Data Type: STRING. Required. Main headline with product name to be shown
	// in dynamic ad.
	CustomPlaceholderFieldEnum_ITEM_TITLE CustomPlaceholderFieldEnum_CustomPlaceholderField = 4
	// Data Type: STRING. Optional text to be shown in the image ad.
	CustomPlaceholderFieldEnum_ITEM_SUBTITLE CustomPlaceholderFieldEnum_CustomPlaceholderField = 5
	// Data Type: STRING. Optional description of the product to be shown in the
	// ad.
	CustomPlaceholderFieldEnum_ITEM_DESCRIPTION CustomPlaceholderFieldEnum_CustomPlaceholderField = 6
	// Data Type: STRING. Full address of your offer or service, including
	// postal code. This will be used to identify the closest product to the
	// user when there are multiple offers in the feed that are relevant to the
	// user.
	CustomPlaceholderFieldEnum_ITEM_ADDRESS CustomPlaceholderFieldEnum_CustomPlaceholderField = 7
	// Data Type: STRING. Price to be shown in the ad.
	// Example: "100.00 USD"
	CustomPlaceholderFieldEnum_PRICE CustomPlaceholderFieldEnum_CustomPlaceholderField = 8
	// Data Type: STRING. Formatted price to be shown in the ad.
	// Example: "Starting at $100.00 USD", "$80 - $100"
	CustomPlaceholderFieldEnum_FORMATTED_PRICE CustomPlaceholderFieldEnum_CustomPlaceholderField = 9
	// Data Type: STRING. Sale price to be shown in the ad.
	// Example: "80.00 USD"
	CustomPlaceholderFieldEnum_SALE_PRICE CustomPlaceholderFieldEnum_CustomPlaceholderField = 10
	// Data Type: STRING. Formatted sale price to be shown in the ad.
	// Example: "On sale for $80.00", "$60 - $80"
	CustomPlaceholderFieldEnum_FORMATTED_SALE_PRICE CustomPlaceholderFieldEnum_CustomPlaceholderField = 11
	// Data Type: URL. Image to be displayed in the ad. Highly recommended for
	// image ads.
	CustomPlaceholderFieldEnum_IMAGE_URL CustomPlaceholderFieldEnum_CustomPlaceholderField = 12
	// Data Type: STRING. Used as a recommendation engine signal to serve items
	// in the same category.
	CustomPlaceholderFieldEnum_ITEM_CATEGORY CustomPlaceholderFieldEnum_CustomPlaceholderField = 13
	// Data Type: URL_LIST. Final URLs for the ad when using Upgraded
	// URLs. User will be redirected to these URLs when they click on an ad, or
	// when they click on a specific product for ads that have multiple
	// products.
	CustomPlaceholderFieldEnum_FINAL_URLS CustomPlaceholderFieldEnum_CustomPlaceholderField = 14
	// Data Type: URL_LIST. Final mobile URLs for the ad when using Upgraded
	// URLs.
	CustomPlaceholderFieldEnum_FINAL_MOBILE_URLS CustomPlaceholderFieldEnum_CustomPlaceholderField = 15
	// Data Type: URL. Tracking template for the ad when using Upgraded URLs.
	CustomPlaceholderFieldEnum_TRACKING_URL CustomPlaceholderFieldEnum_CustomPlaceholderField = 16
	// Data Type: STRING_LIST. Keywords used for product retrieval.
	CustomPlaceholderFieldEnum_CONTEXTUAL_KEYWORDS CustomPlaceholderFieldEnum_CustomPlaceholderField = 17
	// Data Type: STRING. Android app link. Must be formatted as:
	// android-app://{package_id}/{scheme}/{host_path}.
	// The components are defined as follows:
	// package_id: app ID as specified in Google Play.
	// scheme: the scheme to pass to the application. Can be HTTP, or a custom
	//   scheme.
	// host_path: identifies the specific content within your application.
	CustomPlaceholderFieldEnum_ANDROID_APP_LINK CustomPlaceholderFieldEnum_CustomPlaceholderField = 18
	// Data Type: STRING_LIST. List of recommended IDs to show together with
	// this item.
	CustomPlaceholderFieldEnum_SIMILAR_IDS CustomPlaceholderFieldEnum_CustomPlaceholderField = 19
	// Data Type: STRING. iOS app link.
	CustomPlaceholderFieldEnum_IOS_APP_LINK CustomPlaceholderFieldEnum_CustomPlaceholderField = 20
	// Data Type: INT64. iOS app store ID.
	CustomPlaceholderFieldEnum_IOS_APP_STORE_ID CustomPlaceholderFieldEnum_CustomPlaceholderField = 21
)

// Enum value maps for CustomPlaceholderFieldEnum_CustomPlaceholderField.
var (
	CustomPlaceholderFieldEnum_CustomPlaceholderField_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "ID",
		3:  "ID2",
		4:  "ITEM_TITLE",
		5:  "ITEM_SUBTITLE",
		6:  "ITEM_DESCRIPTION",
		7:  "ITEM_ADDRESS",
		8:  "PRICE",
		9:  "FORMATTED_PRICE",
		10: "SALE_PRICE",
		11: "FORMATTED_SALE_PRICE",
		12: "IMAGE_URL",
		13: "ITEM_CATEGORY",
		14: "FINAL_URLS",
		15: "FINAL_MOBILE_URLS",
		16: "TRACKING_URL",
		17: "CONTEXTUAL_KEYWORDS",
		18: "ANDROID_APP_LINK",
		19: "SIMILAR_IDS",
		20: "IOS_APP_LINK",
		21: "IOS_APP_STORE_ID",
	}
	CustomPlaceholderFieldEnum_CustomPlaceholderField_value = map[string]int32{
		"UNSPECIFIED":          0,
		"UNKNOWN":              1,
		"ID":                   2,
		"ID2":                  3,
		"ITEM_TITLE":           4,
		"ITEM_SUBTITLE":        5,
		"ITEM_DESCRIPTION":     6,
		"ITEM_ADDRESS":         7,
		"PRICE":                8,
		"FORMATTED_PRICE":      9,
		"SALE_PRICE":           10,
		"FORMATTED_SALE_PRICE": 11,
		"IMAGE_URL":            12,
		"ITEM_CATEGORY":        13,
		"FINAL_URLS":           14,
		"FINAL_MOBILE_URLS":    15,
		"TRACKING_URL":         16,
		"CONTEXTUAL_KEYWORDS":  17,
		"ANDROID_APP_LINK":     18,
		"SIMILAR_IDS":          19,
		"IOS_APP_LINK":         20,
		"IOS_APP_STORE_ID":     21,
	}
)

func (x CustomPlaceholderFieldEnum_CustomPlaceholderField) Enum() *CustomPlaceholderFieldEnum_CustomPlaceholderField {
	p := new(CustomPlaceholderFieldEnum_CustomPlaceholderField)
	*p = x
	return p
}

func (x CustomPlaceholderFieldEnum_CustomPlaceholderField) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CustomPlaceholderFieldEnum_CustomPlaceholderField) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_enumTypes[0].Descriptor()
}

func (CustomPlaceholderFieldEnum_CustomPlaceholderField) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_enumTypes[0]
}

func (x CustomPlaceholderFieldEnum_CustomPlaceholderField) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CustomPlaceholderFieldEnum_CustomPlaceholderField.Descriptor instead.
func (CustomPlaceholderFieldEnum_CustomPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_rawDescGZIP(), []int{0, 0}
}

// Values for Custom placeholder fields.
// For more information about dynamic remarketing feeds, see
// https://support.google.com/google-ads/answer/6053288.
type CustomPlaceholderFieldEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CustomPlaceholderFieldEnum) Reset() {
	*x = CustomPlaceholderFieldEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomPlaceholderFieldEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomPlaceholderFieldEnum) ProtoMessage() {}

func (x *CustomPlaceholderFieldEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomPlaceholderFieldEnum.ProtoReflect.Descriptor instead.
func (*CustomPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v2_enums_custom_placeholder_field_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x03, 0x0a, 0x1a,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x9f, 0x03, 0x0a, 0x16, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x44, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x49,
	0x44, 0x32, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x54, 0x49, 0x54,
	0x4c, 0x45, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x53, 0x55, 0x42,
	0x54, 0x49, 0x54, 0x4c, 0x45, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x54, 0x45, 0x4d, 0x5f,
	0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x06, 0x12, 0x10, 0x0a,
	0x0c, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x10, 0x07, 0x12,
	0x09, 0x0a, 0x05, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x08, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x4f,
	0x52, 0x4d, 0x41, 0x54, 0x54, 0x45, 0x44, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x09, 0x12,
	0x0e, 0x0a, 0x0a, 0x53, 0x41, 0x4c, 0x45, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x0a, 0x12,
	0x18, 0x0a, 0x14, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x54, 0x45, 0x44, 0x5f, 0x53, 0x41, 0x4c,
	0x45, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x0b, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4d, 0x41,
	0x47, 0x45, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x0c, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x54, 0x45, 0x4d,
	0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x10, 0x0d, 0x12, 0x0e, 0x0a, 0x0a, 0x46,
	0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x53, 0x10, 0x0e, 0x12, 0x15, 0x0a, 0x11, 0x46,
	0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x55, 0x52, 0x4c, 0x53,
	0x10, 0x0f, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x55,
	0x52, 0x4c, 0x10, 0x10, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x58, 0x54, 0x55,
	0x41, 0x4c, 0x5f, 0x4b, 0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x53, 0x10, 0x11, 0x12, 0x14, 0x0a,
	0x10, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x4c, 0x49, 0x4e,
	0x4b, 0x10, 0x12, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x49, 0x4d, 0x49, 0x4c, 0x41, 0x52, 0x5f, 0x49,
	0x44, 0x53, 0x10, 0x13, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4f, 0x53, 0x5f, 0x41, 0x50, 0x50, 0x5f,
	0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x14, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x4f, 0x53, 0x5f, 0x41, 0x50,
	0x50, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x49, 0x44, 0x10, 0x15, 0x42, 0xf0, 0x01, 0x0a,
	0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x42, 0x1b, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1d, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x2e, 0x56, 0x32, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1d, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x5c, 0x56, 0x32, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x21, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x32, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_rawDescData = file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_rawDesc
)

func file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_rawDescData)
	})
	return file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_rawDescData
}

var file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_goTypes = []interface{}{
	(CustomPlaceholderFieldEnum_CustomPlaceholderField)(0), // 0: google.ads.googleads.v2.enums.CustomPlaceholderFieldEnum.CustomPlaceholderField
	(*CustomPlaceholderFieldEnum)(nil),                     // 1: google.ads.googleads.v2.enums.CustomPlaceholderFieldEnum
}
var file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_init() }
func file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_init() {
	if File_google_ads_googleads_v2_enums_custom_placeholder_field_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomPlaceholderFieldEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v2_enums_custom_placeholder_field_proto = out.File
	file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_rawDesc = nil
	file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_goTypes = nil
	file_google_ads_googleads_v2_enums_custom_placeholder_field_proto_depIdxs = nil
}
