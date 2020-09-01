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
// source: google/ads/googleads/v3/enums/customer_match_upload_key_type.proto

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

// Enum describing possible customer match upload key types.
type CustomerMatchUploadKeyTypeEnum_CustomerMatchUploadKeyType int32

const (
	// Not specified.
	CustomerMatchUploadKeyTypeEnum_UNSPECIFIED CustomerMatchUploadKeyTypeEnum_CustomerMatchUploadKeyType = 0
	// Used for return value only. Represents value unknown in this version.
	CustomerMatchUploadKeyTypeEnum_UNKNOWN CustomerMatchUploadKeyTypeEnum_CustomerMatchUploadKeyType = 1
	// Members are matched from customer info such as email address, phone
	// number or physical address.
	CustomerMatchUploadKeyTypeEnum_CONTACT_INFO CustomerMatchUploadKeyTypeEnum_CustomerMatchUploadKeyType = 2
	// Members are matched from a user id generated and assigned by the
	// advertiser.
	CustomerMatchUploadKeyTypeEnum_CRM_ID CustomerMatchUploadKeyTypeEnum_CustomerMatchUploadKeyType = 3
	// Members are matched from mobile advertising ids.
	CustomerMatchUploadKeyTypeEnum_MOBILE_ADVERTISING_ID CustomerMatchUploadKeyTypeEnum_CustomerMatchUploadKeyType = 4
)

// Enum value maps for CustomerMatchUploadKeyTypeEnum_CustomerMatchUploadKeyType.
var (
	CustomerMatchUploadKeyTypeEnum_CustomerMatchUploadKeyType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "CONTACT_INFO",
		3: "CRM_ID",
		4: "MOBILE_ADVERTISING_ID",
	}
	CustomerMatchUploadKeyTypeEnum_CustomerMatchUploadKeyType_value = map[string]int32{
		"UNSPECIFIED":           0,
		"UNKNOWN":               1,
		"CONTACT_INFO":          2,
		"CRM_ID":                3,
		"MOBILE_ADVERTISING_ID": 4,
	}
)

func (x CustomerMatchUploadKeyTypeEnum_CustomerMatchUploadKeyType) Enum() *CustomerMatchUploadKeyTypeEnum_CustomerMatchUploadKeyType {
	p := new(CustomerMatchUploadKeyTypeEnum_CustomerMatchUploadKeyType)
	*p = x
	return p
}

func (x CustomerMatchUploadKeyTypeEnum_CustomerMatchUploadKeyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CustomerMatchUploadKeyTypeEnum_CustomerMatchUploadKeyType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_enumTypes[0].Descriptor()
}

func (CustomerMatchUploadKeyTypeEnum_CustomerMatchUploadKeyType) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_enumTypes[0]
}

func (x CustomerMatchUploadKeyTypeEnum_CustomerMatchUploadKeyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CustomerMatchUploadKeyTypeEnum_CustomerMatchUploadKeyType.Descriptor instead.
func (CustomerMatchUploadKeyTypeEnum_CustomerMatchUploadKeyType) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_rawDescGZIP(), []int{0, 0}
}

// Indicates what type of data are the user list's members matched from.
type CustomerMatchUploadKeyTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CustomerMatchUploadKeyTypeEnum) Reset() {
	*x = CustomerMatchUploadKeyTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerMatchUploadKeyTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerMatchUploadKeyTypeEnum) ProtoMessage() {}

func (x *CustomerMatchUploadKeyTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerMatchUploadKeyTypeEnum.ProtoReflect.Descriptor instead.
func (*CustomerMatchUploadKeyTypeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x95, 0x01, 0x0a, 0x1e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x45, 0x6e, 0x75, 0x6d, 0x22, 0x73, 0x0a, 0x1a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4b, 0x65, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01,
	0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x43, 0x54, 0x5f, 0x49, 0x4e, 0x46, 0x4f,
	0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x4d, 0x5f, 0x49, 0x44, 0x10, 0x03, 0x12, 0x19,
	0x0a, 0x15, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x41, 0x44, 0x56, 0x45, 0x52, 0x54, 0x49,
	0x53, 0x49, 0x4e, 0x47, 0x5f, 0x49, 0x44, 0x10, 0x04, 0x42, 0xf4, 0x01, 0x0a, 0x21, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42,
	0x1f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1d, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x2e, 0x56, 0x33, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1d, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x5c, 0x56, 0x33, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x21, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x33, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_rawDescData = file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_rawDesc
)

func file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_rawDescData)
	})
	return file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_rawDescData
}

var file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_goTypes = []interface{}{
	(CustomerMatchUploadKeyTypeEnum_CustomerMatchUploadKeyType)(0), // 0: google.ads.googleads.v3.enums.CustomerMatchUploadKeyTypeEnum.CustomerMatchUploadKeyType
	(*CustomerMatchUploadKeyTypeEnum)(nil),                         // 1: google.ads.googleads.v3.enums.CustomerMatchUploadKeyTypeEnum
}
var file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_init() }
func file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_init() {
	if File_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerMatchUploadKeyTypeEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto = out.File
	file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_rawDesc = nil
	file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_goTypes = nil
	file_google_ads_googleads_v3_enums_customer_match_upload_key_type_proto_depIdxs = nil
}
