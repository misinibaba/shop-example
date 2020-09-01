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
// source: google/cloud/automl/v1beta1/model.proto

package automl

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Deployment state of the model.
type Model_DeploymentState int32

const (
	// Should not be used, an un-set enum has this value by default.
	Model_DEPLOYMENT_STATE_UNSPECIFIED Model_DeploymentState = 0
	// Model is deployed.
	Model_DEPLOYED Model_DeploymentState = 1
	// Model is not deployed.
	Model_UNDEPLOYED Model_DeploymentState = 2
)

// Enum value maps for Model_DeploymentState.
var (
	Model_DeploymentState_name = map[int32]string{
		0: "DEPLOYMENT_STATE_UNSPECIFIED",
		1: "DEPLOYED",
		2: "UNDEPLOYED",
	}
	Model_DeploymentState_value = map[string]int32{
		"DEPLOYMENT_STATE_UNSPECIFIED": 0,
		"DEPLOYED":                     1,
		"UNDEPLOYED":                   2,
	}
)

func (x Model_DeploymentState) Enum() *Model_DeploymentState {
	p := new(Model_DeploymentState)
	*p = x
	return p
}

func (x Model_DeploymentState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Model_DeploymentState) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_automl_v1beta1_model_proto_enumTypes[0].Descriptor()
}

func (Model_DeploymentState) Type() protoreflect.EnumType {
	return &file_google_cloud_automl_v1beta1_model_proto_enumTypes[0]
}

func (x Model_DeploymentState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Model_DeploymentState.Descriptor instead.
func (Model_DeploymentState) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_model_proto_rawDescGZIP(), []int{0, 0}
}

// API proto representing a trained machine learning model.
type Model struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required.
	// The model metadata that is specific to the problem type.
	// Must match the metadata type of the dataset used to train the model.
	//
	// Types that are assignable to ModelMetadata:
	//	*Model_TranslationModelMetadata
	//	*Model_ImageClassificationModelMetadata
	//	*Model_TextClassificationModelMetadata
	//	*Model_ImageObjectDetectionModelMetadata
	//	*Model_VideoClassificationModelMetadata
	//	*Model_VideoObjectTrackingModelMetadata
	//	*Model_TextExtractionModelMetadata
	//	*Model_TablesModelMetadata
	//	*Model_TextSentimentModelMetadata
	ModelMetadata isModel_ModelMetadata `protobuf_oneof:"model_metadata"`
	// Output only. Resource name of the model.
	// Format: `projects/{project_id}/locations/{location_id}/models/{model_id}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The name of the model to show in the interface. The name can be
	// up to 32 characters long and can consist only of ASCII Latin letters A-Z
	// and a-z, underscores
	// (_), and ASCII digits 0-9. It must start with a letter.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Required. The resource ID of the dataset used to create the model. The dataset must
	// come from the same ancestor project and location.
	DatasetId string `protobuf:"bytes,3,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	// Output only. Timestamp when the model training finished  and can be used for prediction.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Timestamp when this model was last updated.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,11,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. Deployment state of the model. A model can only serve
	// prediction requests after it gets deployed.
	DeploymentState Model_DeploymentState `protobuf:"varint,8,opt,name=deployment_state,json=deploymentState,proto3,enum=google.cloud.automl.v1beta1.Model_DeploymentState" json:"deployment_state,omitempty"`
}

func (x *Model) Reset() {
	*x = Model{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Model) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Model) ProtoMessage() {}

func (x *Model) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Model.ProtoReflect.Descriptor instead.
func (*Model) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_model_proto_rawDescGZIP(), []int{0}
}

func (m *Model) GetModelMetadata() isModel_ModelMetadata {
	if m != nil {
		return m.ModelMetadata
	}
	return nil
}

func (x *Model) GetTranslationModelMetadata() *TranslationModelMetadata {
	if x, ok := x.GetModelMetadata().(*Model_TranslationModelMetadata); ok {
		return x.TranslationModelMetadata
	}
	return nil
}

func (x *Model) GetImageClassificationModelMetadata() *ImageClassificationModelMetadata {
	if x, ok := x.GetModelMetadata().(*Model_ImageClassificationModelMetadata); ok {
		return x.ImageClassificationModelMetadata
	}
	return nil
}

func (x *Model) GetTextClassificationModelMetadata() *TextClassificationModelMetadata {
	if x, ok := x.GetModelMetadata().(*Model_TextClassificationModelMetadata); ok {
		return x.TextClassificationModelMetadata
	}
	return nil
}

func (x *Model) GetImageObjectDetectionModelMetadata() *ImageObjectDetectionModelMetadata {
	if x, ok := x.GetModelMetadata().(*Model_ImageObjectDetectionModelMetadata); ok {
		return x.ImageObjectDetectionModelMetadata
	}
	return nil
}

func (x *Model) GetVideoClassificationModelMetadata() *VideoClassificationModelMetadata {
	if x, ok := x.GetModelMetadata().(*Model_VideoClassificationModelMetadata); ok {
		return x.VideoClassificationModelMetadata
	}
	return nil
}

func (x *Model) GetVideoObjectTrackingModelMetadata() *VideoObjectTrackingModelMetadata {
	if x, ok := x.GetModelMetadata().(*Model_VideoObjectTrackingModelMetadata); ok {
		return x.VideoObjectTrackingModelMetadata
	}
	return nil
}

func (x *Model) GetTextExtractionModelMetadata() *TextExtractionModelMetadata {
	if x, ok := x.GetModelMetadata().(*Model_TextExtractionModelMetadata); ok {
		return x.TextExtractionModelMetadata
	}
	return nil
}

func (x *Model) GetTablesModelMetadata() *TablesModelMetadata {
	if x, ok := x.GetModelMetadata().(*Model_TablesModelMetadata); ok {
		return x.TablesModelMetadata
	}
	return nil
}

func (x *Model) GetTextSentimentModelMetadata() *TextSentimentModelMetadata {
	if x, ok := x.GetModelMetadata().(*Model_TextSentimentModelMetadata); ok {
		return x.TextSentimentModelMetadata
	}
	return nil
}

func (x *Model) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Model) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Model) GetDatasetId() string {
	if x != nil {
		return x.DatasetId
	}
	return ""
}

func (x *Model) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Model) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Model) GetDeploymentState() Model_DeploymentState {
	if x != nil {
		return x.DeploymentState
	}
	return Model_DEPLOYMENT_STATE_UNSPECIFIED
}

type isModel_ModelMetadata interface {
	isModel_ModelMetadata()
}

type Model_TranslationModelMetadata struct {
	// Metadata for translation models.
	TranslationModelMetadata *TranslationModelMetadata `protobuf:"bytes,15,opt,name=translation_model_metadata,json=translationModelMetadata,proto3,oneof"`
}

type Model_ImageClassificationModelMetadata struct {
	// Metadata for image classification models.
	ImageClassificationModelMetadata *ImageClassificationModelMetadata `protobuf:"bytes,13,opt,name=image_classification_model_metadata,json=imageClassificationModelMetadata,proto3,oneof"`
}

type Model_TextClassificationModelMetadata struct {
	// Metadata for text classification models.
	TextClassificationModelMetadata *TextClassificationModelMetadata `protobuf:"bytes,14,opt,name=text_classification_model_metadata,json=textClassificationModelMetadata,proto3,oneof"`
}

type Model_ImageObjectDetectionModelMetadata struct {
	// Metadata for image object detection models.
	ImageObjectDetectionModelMetadata *ImageObjectDetectionModelMetadata `protobuf:"bytes,20,opt,name=image_object_detection_model_metadata,json=imageObjectDetectionModelMetadata,proto3,oneof"`
}

type Model_VideoClassificationModelMetadata struct {
	// Metadata for video classification models.
	VideoClassificationModelMetadata *VideoClassificationModelMetadata `protobuf:"bytes,23,opt,name=video_classification_model_metadata,json=videoClassificationModelMetadata,proto3,oneof"`
}

type Model_VideoObjectTrackingModelMetadata struct {
	// Metadata for video object tracking models.
	VideoObjectTrackingModelMetadata *VideoObjectTrackingModelMetadata `protobuf:"bytes,21,opt,name=video_object_tracking_model_metadata,json=videoObjectTrackingModelMetadata,proto3,oneof"`
}

type Model_TextExtractionModelMetadata struct {
	// Metadata for text extraction models.
	TextExtractionModelMetadata *TextExtractionModelMetadata `protobuf:"bytes,19,opt,name=text_extraction_model_metadata,json=textExtractionModelMetadata,proto3,oneof"`
}

type Model_TablesModelMetadata struct {
	// Metadata for Tables models.
	TablesModelMetadata *TablesModelMetadata `protobuf:"bytes,24,opt,name=tables_model_metadata,json=tablesModelMetadata,proto3,oneof"`
}

type Model_TextSentimentModelMetadata struct {
	// Metadata for text sentiment models.
	TextSentimentModelMetadata *TextSentimentModelMetadata `protobuf:"bytes,22,opt,name=text_sentiment_model_metadata,json=textSentimentModelMetadata,proto3,oneof"`
}

func (*Model_TranslationModelMetadata) isModel_ModelMetadata() {}

func (*Model_ImageClassificationModelMetadata) isModel_ModelMetadata() {}

func (*Model_TextClassificationModelMetadata) isModel_ModelMetadata() {}

func (*Model_ImageObjectDetectionModelMetadata) isModel_ModelMetadata() {}

func (*Model_VideoClassificationModelMetadata) isModel_ModelMetadata() {}

func (*Model_VideoObjectTrackingModelMetadata) isModel_ModelMetadata() {}

func (*Model_TextExtractionModelMetadata) isModel_ModelMetadata() {}

func (*Model_TablesModelMetadata) isModel_ModelMetadata() {}

func (*Model_TextSentimentModelMetadata) isModel_ModelMetadata() {}

var File_google_cloud_automl_v1beta1_model_proto protoreflect.FileDescriptor

var file_google_cloud_automl_v1beta1_model_proto_rawDesc = []byte{
	0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x0d, 0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x75, 0x0a,
	0x1a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x18, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x8e, 0x01, 0x0a, 0x23, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x48, 0x00, 0x52, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x8b, 0x01, 0x0a, 0x22, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x54, 0x65, 0x78, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x48, 0x00, 0x52, 0x1f, 0x74, 0x65, 0x78, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x92, 0x01, 0x0a, 0x25, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x21, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x8e, 0x01, 0x0a, 0x23, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x20, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x8f, 0x01, 0x0a, 0x24, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x20, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x7f, 0x0a, 0x1e, 0x74,
	0x65, 0x78, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x13, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52,
	0x1b, 0x74, 0x65, 0x78, 0x74, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x66, 0x0a, 0x15,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52,
	0x13, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x7c, 0x0a, 0x1d, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x65, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x53, 0x65,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x1a, 0x74, 0x65, 0x78, 0x74, 0x53, 0x65, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x5d, 0x0a, 0x10, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x0f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x22, 0x51, 0x0a, 0x0f, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x1c, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45,
	0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59,
	0x45, 0x44, 0x10, 0x02, 0x3a, 0x58, 0xea, 0x41, 0x55, 0x0a, 0x1b, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x36, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x7b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x7d, 0x42, 0x10,
	0x0a, 0x0e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x42, 0xa5, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0xca, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x6c, 0x5c,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x4c, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_automl_v1beta1_model_proto_rawDescOnce sync.Once
	file_google_cloud_automl_v1beta1_model_proto_rawDescData = file_google_cloud_automl_v1beta1_model_proto_rawDesc
)

func file_google_cloud_automl_v1beta1_model_proto_rawDescGZIP() []byte {
	file_google_cloud_automl_v1beta1_model_proto_rawDescOnce.Do(func() {
		file_google_cloud_automl_v1beta1_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_automl_v1beta1_model_proto_rawDescData)
	})
	return file_google_cloud_automl_v1beta1_model_proto_rawDescData
}

var file_google_cloud_automl_v1beta1_model_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_automl_v1beta1_model_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_automl_v1beta1_model_proto_goTypes = []interface{}{
	(Model_DeploymentState)(0),                // 0: google.cloud.automl.v1beta1.Model.DeploymentState
	(*Model)(nil),                             // 1: google.cloud.automl.v1beta1.Model
	(*TranslationModelMetadata)(nil),          // 2: google.cloud.automl.v1beta1.TranslationModelMetadata
	(*ImageClassificationModelMetadata)(nil),  // 3: google.cloud.automl.v1beta1.ImageClassificationModelMetadata
	(*TextClassificationModelMetadata)(nil),   // 4: google.cloud.automl.v1beta1.TextClassificationModelMetadata
	(*ImageObjectDetectionModelMetadata)(nil), // 5: google.cloud.automl.v1beta1.ImageObjectDetectionModelMetadata
	(*VideoClassificationModelMetadata)(nil),  // 6: google.cloud.automl.v1beta1.VideoClassificationModelMetadata
	(*VideoObjectTrackingModelMetadata)(nil),  // 7: google.cloud.automl.v1beta1.VideoObjectTrackingModelMetadata
	(*TextExtractionModelMetadata)(nil),       // 8: google.cloud.automl.v1beta1.TextExtractionModelMetadata
	(*TablesModelMetadata)(nil),               // 9: google.cloud.automl.v1beta1.TablesModelMetadata
	(*TextSentimentModelMetadata)(nil),        // 10: google.cloud.automl.v1beta1.TextSentimentModelMetadata
	(*timestamp.Timestamp)(nil),               // 11: google.protobuf.Timestamp
}
var file_google_cloud_automl_v1beta1_model_proto_depIdxs = []int32{
	2,  // 0: google.cloud.automl.v1beta1.Model.translation_model_metadata:type_name -> google.cloud.automl.v1beta1.TranslationModelMetadata
	3,  // 1: google.cloud.automl.v1beta1.Model.image_classification_model_metadata:type_name -> google.cloud.automl.v1beta1.ImageClassificationModelMetadata
	4,  // 2: google.cloud.automl.v1beta1.Model.text_classification_model_metadata:type_name -> google.cloud.automl.v1beta1.TextClassificationModelMetadata
	5,  // 3: google.cloud.automl.v1beta1.Model.image_object_detection_model_metadata:type_name -> google.cloud.automl.v1beta1.ImageObjectDetectionModelMetadata
	6,  // 4: google.cloud.automl.v1beta1.Model.video_classification_model_metadata:type_name -> google.cloud.automl.v1beta1.VideoClassificationModelMetadata
	7,  // 5: google.cloud.automl.v1beta1.Model.video_object_tracking_model_metadata:type_name -> google.cloud.automl.v1beta1.VideoObjectTrackingModelMetadata
	8,  // 6: google.cloud.automl.v1beta1.Model.text_extraction_model_metadata:type_name -> google.cloud.automl.v1beta1.TextExtractionModelMetadata
	9,  // 7: google.cloud.automl.v1beta1.Model.tables_model_metadata:type_name -> google.cloud.automl.v1beta1.TablesModelMetadata
	10, // 8: google.cloud.automl.v1beta1.Model.text_sentiment_model_metadata:type_name -> google.cloud.automl.v1beta1.TextSentimentModelMetadata
	11, // 9: google.cloud.automl.v1beta1.Model.create_time:type_name -> google.protobuf.Timestamp
	11, // 10: google.cloud.automl.v1beta1.Model.update_time:type_name -> google.protobuf.Timestamp
	0,  // 11: google.cloud.automl.v1beta1.Model.deployment_state:type_name -> google.cloud.automl.v1beta1.Model.DeploymentState
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_google_cloud_automl_v1beta1_model_proto_init() }
func file_google_cloud_automl_v1beta1_model_proto_init() {
	if File_google_cloud_automl_v1beta1_model_proto != nil {
		return
	}
	file_google_cloud_automl_v1beta1_image_proto_init()
	file_google_cloud_automl_v1beta1_tables_proto_init()
	file_google_cloud_automl_v1beta1_text_proto_init()
	file_google_cloud_automl_v1beta1_translation_proto_init()
	file_google_cloud_automl_v1beta1_video_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_automl_v1beta1_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Model); i {
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
	file_google_cloud_automl_v1beta1_model_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Model_TranslationModelMetadata)(nil),
		(*Model_ImageClassificationModelMetadata)(nil),
		(*Model_TextClassificationModelMetadata)(nil),
		(*Model_ImageObjectDetectionModelMetadata)(nil),
		(*Model_VideoClassificationModelMetadata)(nil),
		(*Model_VideoObjectTrackingModelMetadata)(nil),
		(*Model_TextExtractionModelMetadata)(nil),
		(*Model_TablesModelMetadata)(nil),
		(*Model_TextSentimentModelMetadata)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_automl_v1beta1_model_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_automl_v1beta1_model_proto_goTypes,
		DependencyIndexes: file_google_cloud_automl_v1beta1_model_proto_depIdxs,
		EnumInfos:         file_google_cloud_automl_v1beta1_model_proto_enumTypes,
		MessageInfos:      file_google_cloud_automl_v1beta1_model_proto_msgTypes,
	}.Build()
	File_google_cloud_automl_v1beta1_model_proto = out.File
	file_google_cloud_automl_v1beta1_model_proto_rawDesc = nil
	file_google_cloud_automl_v1beta1_model_proto_goTypes = nil
	file_google_cloud_automl_v1beta1_model_proto_depIdxs = nil
}
