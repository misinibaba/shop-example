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
// source: google/cloud/automl/v1beta1/text.proto

package automl

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

// Dataset metadata for classification.
type TextClassificationDatasetMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Type of the classification problem.
	ClassificationType ClassificationType `protobuf:"varint,1,opt,name=classification_type,json=classificationType,proto3,enum=google.cloud.automl.v1beta1.ClassificationType" json:"classification_type,omitempty"`
}

func (x *TextClassificationDatasetMetadata) Reset() {
	*x = TextClassificationDatasetMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_text_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextClassificationDatasetMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextClassificationDatasetMetadata) ProtoMessage() {}

func (x *TextClassificationDatasetMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_text_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextClassificationDatasetMetadata.ProtoReflect.Descriptor instead.
func (*TextClassificationDatasetMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_text_proto_rawDescGZIP(), []int{0}
}

func (x *TextClassificationDatasetMetadata) GetClassificationType() ClassificationType {
	if x != nil {
		return x.ClassificationType
	}
	return ClassificationType_CLASSIFICATION_TYPE_UNSPECIFIED
}

// Model metadata that is specific to text classification.
type TextClassificationModelMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Classification type of the dataset used to train this model.
	ClassificationType ClassificationType `protobuf:"varint,3,opt,name=classification_type,json=classificationType,proto3,enum=google.cloud.automl.v1beta1.ClassificationType" json:"classification_type,omitempty"`
}

func (x *TextClassificationModelMetadata) Reset() {
	*x = TextClassificationModelMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_text_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextClassificationModelMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextClassificationModelMetadata) ProtoMessage() {}

func (x *TextClassificationModelMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_text_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextClassificationModelMetadata.ProtoReflect.Descriptor instead.
func (*TextClassificationModelMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_text_proto_rawDescGZIP(), []int{1}
}

func (x *TextClassificationModelMetadata) GetClassificationType() ClassificationType {
	if x != nil {
		return x.ClassificationType
	}
	return ClassificationType_CLASSIFICATION_TYPE_UNSPECIFIED
}

// Dataset metadata that is specific to text extraction
type TextExtractionDatasetMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TextExtractionDatasetMetadata) Reset() {
	*x = TextExtractionDatasetMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_text_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextExtractionDatasetMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextExtractionDatasetMetadata) ProtoMessage() {}

func (x *TextExtractionDatasetMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_text_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextExtractionDatasetMetadata.ProtoReflect.Descriptor instead.
func (*TextExtractionDatasetMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_text_proto_rawDescGZIP(), []int{2}
}

// Model metadata that is specific to text extraction.
type TextExtractionModelMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TextExtractionModelMetadata) Reset() {
	*x = TextExtractionModelMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_text_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextExtractionModelMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextExtractionModelMetadata) ProtoMessage() {}

func (x *TextExtractionModelMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_text_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextExtractionModelMetadata.ProtoReflect.Descriptor instead.
func (*TextExtractionModelMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_text_proto_rawDescGZIP(), []int{3}
}

// Dataset metadata for text sentiment.
type TextSentimentDatasetMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. A sentiment is expressed as an integer ordinal, where higher value
	// means a more positive sentiment. The range of sentiments that will be used
	// is between 0 and sentiment_max (inclusive on both ends), and all the values
	// in the range must be represented in the dataset before a model can be
	// created.
	// sentiment_max value must be between 1 and 10 (inclusive).
	SentimentMax int32 `protobuf:"varint,1,opt,name=sentiment_max,json=sentimentMax,proto3" json:"sentiment_max,omitempty"`
}

func (x *TextSentimentDatasetMetadata) Reset() {
	*x = TextSentimentDatasetMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_text_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextSentimentDatasetMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextSentimentDatasetMetadata) ProtoMessage() {}

func (x *TextSentimentDatasetMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_text_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextSentimentDatasetMetadata.ProtoReflect.Descriptor instead.
func (*TextSentimentDatasetMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_text_proto_rawDescGZIP(), []int{4}
}

func (x *TextSentimentDatasetMetadata) GetSentimentMax() int32 {
	if x != nil {
		return x.SentimentMax
	}
	return 0
}

// Model metadata that is specific to text sentiment.
type TextSentimentModelMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TextSentimentModelMetadata) Reset() {
	*x = TextSentimentModelMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_text_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextSentimentModelMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextSentimentModelMetadata) ProtoMessage() {}

func (x *TextSentimentModelMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_text_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextSentimentModelMetadata.ProtoReflect.Descriptor instead.
func (*TextSentimentModelMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_text_proto_rawDescGZIP(), []int{5}
}

var File_google_cloud_automl_v1beta1_text_proto protoreflect.FileDescriptor

var file_google_cloud_automl_v1beta1_text_proto_rawDesc = []byte{
	0x0a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x65,
	0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x21, 0x54, 0x65, 0x78, 0x74, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x60, 0x0a, 0x13, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x12, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x83, 0x01,
	0x0a, 0x1f, 0x54, 0x65, 0x78, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x60, 0x0a, 0x13, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x12, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x1f, 0x0a, 0x1d, 0x54, 0x65, 0x78, 0x74, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x1d, 0x0a, 0x1b, 0x54, 0x65, 0x78, 0x74, 0x45, 0x78, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x43, 0x0a, 0x1c, 0x54, 0x65, 0x78, 0x74, 0x53, 0x65, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x6d, 0x61, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x65, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x78, 0x22, 0x1c, 0x0a, 0x1a, 0x54, 0x65, 0x78, 0x74,
	0x53, 0x65, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0xb0, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x09, 0x54, 0x65, 0x78, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0xca, 0x02, 0x1b, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x6c,
	0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x4c,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_cloud_automl_v1beta1_text_proto_rawDescOnce sync.Once
	file_google_cloud_automl_v1beta1_text_proto_rawDescData = file_google_cloud_automl_v1beta1_text_proto_rawDesc
)

func file_google_cloud_automl_v1beta1_text_proto_rawDescGZIP() []byte {
	file_google_cloud_automl_v1beta1_text_proto_rawDescOnce.Do(func() {
		file_google_cloud_automl_v1beta1_text_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_automl_v1beta1_text_proto_rawDescData)
	})
	return file_google_cloud_automl_v1beta1_text_proto_rawDescData
}

var file_google_cloud_automl_v1beta1_text_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_automl_v1beta1_text_proto_goTypes = []interface{}{
	(*TextClassificationDatasetMetadata)(nil), // 0: google.cloud.automl.v1beta1.TextClassificationDatasetMetadata
	(*TextClassificationModelMetadata)(nil),   // 1: google.cloud.automl.v1beta1.TextClassificationModelMetadata
	(*TextExtractionDatasetMetadata)(nil),     // 2: google.cloud.automl.v1beta1.TextExtractionDatasetMetadata
	(*TextExtractionModelMetadata)(nil),       // 3: google.cloud.automl.v1beta1.TextExtractionModelMetadata
	(*TextSentimentDatasetMetadata)(nil),      // 4: google.cloud.automl.v1beta1.TextSentimentDatasetMetadata
	(*TextSentimentModelMetadata)(nil),        // 5: google.cloud.automl.v1beta1.TextSentimentModelMetadata
	(ClassificationType)(0),                   // 6: google.cloud.automl.v1beta1.ClassificationType
}
var file_google_cloud_automl_v1beta1_text_proto_depIdxs = []int32{
	6, // 0: google.cloud.automl.v1beta1.TextClassificationDatasetMetadata.classification_type:type_name -> google.cloud.automl.v1beta1.ClassificationType
	6, // 1: google.cloud.automl.v1beta1.TextClassificationModelMetadata.classification_type:type_name -> google.cloud.automl.v1beta1.ClassificationType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_automl_v1beta1_text_proto_init() }
func file_google_cloud_automl_v1beta1_text_proto_init() {
	if File_google_cloud_automl_v1beta1_text_proto != nil {
		return
	}
	file_google_cloud_automl_v1beta1_classification_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_automl_v1beta1_text_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextClassificationDatasetMetadata); i {
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
		file_google_cloud_automl_v1beta1_text_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextClassificationModelMetadata); i {
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
		file_google_cloud_automl_v1beta1_text_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextExtractionDatasetMetadata); i {
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
		file_google_cloud_automl_v1beta1_text_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextExtractionModelMetadata); i {
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
		file_google_cloud_automl_v1beta1_text_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextSentimentDatasetMetadata); i {
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
		file_google_cloud_automl_v1beta1_text_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextSentimentModelMetadata); i {
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
			RawDescriptor: file_google_cloud_automl_v1beta1_text_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_automl_v1beta1_text_proto_goTypes,
		DependencyIndexes: file_google_cloud_automl_v1beta1_text_proto_depIdxs,
		MessageInfos:      file_google_cloud_automl_v1beta1_text_proto_msgTypes,
	}.Build()
	File_google_cloud_automl_v1beta1_text_proto = out.File
	file_google_cloud_automl_v1beta1_text_proto_rawDesc = nil
	file_google_cloud_automl_v1beta1_text_proto_goTypes = nil
	file_google_cloud_automl_v1beta1_text_proto_depIdxs = nil
}