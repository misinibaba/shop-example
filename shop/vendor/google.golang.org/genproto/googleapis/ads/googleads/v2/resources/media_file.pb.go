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
// source: google/ads/googleads/v2/resources/media_file.proto

package resources

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
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

// A media file.
type MediaFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the media file.
	// Media file resource names have the form:
	//
	// `customers/{customer_id}/mediaFiles/{media_file_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the media file.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Immutable. Type of the media file.
	Type enums.MediaTypeEnum_MediaType `protobuf:"varint,5,opt,name=type,proto3,enum=google.ads.googleads.v2.enums.MediaTypeEnum_MediaType" json:"type,omitempty"`
	// Output only. The mime type of the media file.
	MimeType enums.MimeTypeEnum_MimeType `protobuf:"varint,6,opt,name=mime_type,json=mimeType,proto3,enum=google.ads.googleads.v2.enums.MimeTypeEnum_MimeType" json:"mime_type,omitempty"`
	// Immutable. The URL of where the original media file was downloaded from (or a file
	// name). Only used for media of type AUDIO and IMAGE.
	SourceUrl *wrappers.StringValue `protobuf:"bytes,7,opt,name=source_url,json=sourceUrl,proto3" json:"source_url,omitempty"`
	// Immutable. The name of the media file. The name can be used by clients to help
	// identify previously uploaded media.
	Name *wrappers.StringValue `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The size of the media file in bytes.
	FileSize *wrappers.Int64Value `protobuf:"bytes,9,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	// The specific type of the media file.
	//
	// Types that are assignable to Mediatype:
	//	*MediaFile_Image
	//	*MediaFile_MediaBundle
	//	*MediaFile_Audio
	//	*MediaFile_Video
	Mediatype isMediaFile_Mediatype `protobuf_oneof:"mediatype"`
}

func (x *MediaFile) Reset() {
	*x = MediaFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_resources_media_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaFile) ProtoMessage() {}

func (x *MediaFile) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_resources_media_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaFile.ProtoReflect.Descriptor instead.
func (*MediaFile) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_resources_media_file_proto_rawDescGZIP(), []int{0}
}

func (x *MediaFile) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *MediaFile) GetId() *wrappers.Int64Value {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *MediaFile) GetType() enums.MediaTypeEnum_MediaType {
	if x != nil {
		return x.Type
	}
	return enums.MediaTypeEnum_UNSPECIFIED
}

func (x *MediaFile) GetMimeType() enums.MimeTypeEnum_MimeType {
	if x != nil {
		return x.MimeType
	}
	return enums.MimeTypeEnum_UNSPECIFIED
}

func (x *MediaFile) GetSourceUrl() *wrappers.StringValue {
	if x != nil {
		return x.SourceUrl
	}
	return nil
}

func (x *MediaFile) GetName() *wrappers.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *MediaFile) GetFileSize() *wrappers.Int64Value {
	if x != nil {
		return x.FileSize
	}
	return nil
}

func (m *MediaFile) GetMediatype() isMediaFile_Mediatype {
	if m != nil {
		return m.Mediatype
	}
	return nil
}

func (x *MediaFile) GetImage() *MediaImage {
	if x, ok := x.GetMediatype().(*MediaFile_Image); ok {
		return x.Image
	}
	return nil
}

func (x *MediaFile) GetMediaBundle() *MediaBundle {
	if x, ok := x.GetMediatype().(*MediaFile_MediaBundle); ok {
		return x.MediaBundle
	}
	return nil
}

func (x *MediaFile) GetAudio() *MediaAudio {
	if x, ok := x.GetMediatype().(*MediaFile_Audio); ok {
		return x.Audio
	}
	return nil
}

func (x *MediaFile) GetVideo() *MediaVideo {
	if x, ok := x.GetMediatype().(*MediaFile_Video); ok {
		return x.Video
	}
	return nil
}

type isMediaFile_Mediatype interface {
	isMediaFile_Mediatype()
}

type MediaFile_Image struct {
	// Immutable. Encapsulates an Image.
	Image *MediaImage `protobuf:"bytes,3,opt,name=image,proto3,oneof"`
}

type MediaFile_MediaBundle struct {
	// Immutable. A ZIP archive media the content of which contains HTML5 assets.
	MediaBundle *MediaBundle `protobuf:"bytes,4,opt,name=media_bundle,json=mediaBundle,proto3,oneof"`
}

type MediaFile_Audio struct {
	// Output only. Encapsulates an Audio.
	Audio *MediaAudio `protobuf:"bytes,10,opt,name=audio,proto3,oneof"`
}

type MediaFile_Video struct {
	// Immutable. Encapsulates a Video.
	Video *MediaVideo `protobuf:"bytes,11,opt,name=video,proto3,oneof"`
}

func (*MediaFile_Image) isMediaFile_Mediatype() {}

func (*MediaFile_MediaBundle) isMediaFile_Mediatype() {}

func (*MediaFile_Audio) isMediaFile_Mediatype() {}

func (*MediaFile_Video) isMediaFile_Mediatype() {}

// Encapsulates an Image.
type MediaImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. Raw image data.
	Data *wrappers.BytesValue `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *MediaImage) Reset() {
	*x = MediaImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_resources_media_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaImage) ProtoMessage() {}

func (x *MediaImage) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_resources_media_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaImage.ProtoReflect.Descriptor instead.
func (*MediaImage) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_resources_media_file_proto_rawDescGZIP(), []int{1}
}

func (x *MediaImage) GetData() *wrappers.BytesValue {
	if x != nil {
		return x.Data
	}
	return nil
}

// Represents a ZIP archive media the content of which contains HTML5 assets.
type MediaBundle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. Raw zipped data.
	Data *wrappers.BytesValue `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *MediaBundle) Reset() {
	*x = MediaBundle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_resources_media_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaBundle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaBundle) ProtoMessage() {}

func (x *MediaBundle) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_resources_media_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaBundle.ProtoReflect.Descriptor instead.
func (*MediaBundle) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_resources_media_file_proto_rawDescGZIP(), []int{2}
}

func (x *MediaBundle) GetData() *wrappers.BytesValue {
	if x != nil {
		return x.Data
	}
	return nil
}

// Encapsulates an Audio.
type MediaAudio struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The duration of the Audio in milliseconds.
	AdDurationMillis *wrappers.Int64Value `protobuf:"bytes,1,opt,name=ad_duration_millis,json=adDurationMillis,proto3" json:"ad_duration_millis,omitempty"`
}

func (x *MediaAudio) Reset() {
	*x = MediaAudio{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_resources_media_file_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaAudio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaAudio) ProtoMessage() {}

func (x *MediaAudio) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_resources_media_file_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaAudio.ProtoReflect.Descriptor instead.
func (*MediaAudio) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_resources_media_file_proto_rawDescGZIP(), []int{3}
}

func (x *MediaAudio) GetAdDurationMillis() *wrappers.Int64Value {
	if x != nil {
		return x.AdDurationMillis
	}
	return nil
}

// Encapsulates a Video.
type MediaVideo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The duration of the Video in milliseconds.
	AdDurationMillis *wrappers.Int64Value `protobuf:"bytes,1,opt,name=ad_duration_millis,json=adDurationMillis,proto3" json:"ad_duration_millis,omitempty"`
	// Immutable. The YouTube video ID (as seen in YouTube URLs).
	YoutubeVideoId *wrappers.StringValue `protobuf:"bytes,2,opt,name=youtube_video_id,json=youtubeVideoId,proto3" json:"youtube_video_id,omitempty"`
	// Output only. The Advertising Digital Identification code for this video, as defined by
	// the American Association of Advertising Agencies, used mainly for
	// television commercials.
	AdvertisingIdCode *wrappers.StringValue `protobuf:"bytes,3,opt,name=advertising_id_code,json=advertisingIdCode,proto3" json:"advertising_id_code,omitempty"`
	// Output only. The Industry Standard Commercial Identifier code for this video, used
	// mainly for television commercials.
	IsciCode *wrappers.StringValue `protobuf:"bytes,4,opt,name=isci_code,json=isciCode,proto3" json:"isci_code,omitempty"`
}

func (x *MediaVideo) Reset() {
	*x = MediaVideo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_resources_media_file_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaVideo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaVideo) ProtoMessage() {}

func (x *MediaVideo) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_resources_media_file_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaVideo.ProtoReflect.Descriptor instead.
func (*MediaVideo) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_resources_media_file_proto_rawDescGZIP(), []int{4}
}

func (x *MediaVideo) GetAdDurationMillis() *wrappers.Int64Value {
	if x != nil {
		return x.AdDurationMillis
	}
	return nil
}

func (x *MediaVideo) GetYoutubeVideoId() *wrappers.StringValue {
	if x != nil {
		return x.YoutubeVideoId
	}
	return nil
}

func (x *MediaVideo) GetAdvertisingIdCode() *wrappers.StringValue {
	if x != nil {
		return x.AdvertisingIdCode
	}
	return nil
}

func (x *MediaVideo) GetIsciCode() *wrappers.StringValue {
	if x != nil {
		return x.IsciCode
	}
	return nil
}

var File_google_ads_googleads_v2_resources_media_file_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v2_resources_media_file_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x91, 0x07, 0x0a, 0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x4f,
	0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x24, 0x0a, 0x22, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x30, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x4f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x56, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x4d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e,
	0x75, 0x6d, 0x2e, 0x4d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41,
	0x05, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x35, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x4a, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x58,
	0x0a, 0x0c, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x42, 0x75,
	0x6e, 0x64, 0x6c, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x0b, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x4a, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69,
	0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x05, 0x61,
	0x75, 0x64, 0x69, 0x6f, 0x12, 0x4a, 0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x3a, 0x55, 0xea, 0x41, 0x52, 0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2c, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x7d, 0x2f,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x7d, 0x42, 0x0b, 0x0a, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x42, 0x0a, 0x0a, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0,
	0x41, 0x05, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x43, 0x0a, 0x0b, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5c, 0x0a,
	0x0a, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x4e, 0x0a, 0x12, 0x61,
	0x64, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x10, 0x61, 0x64, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x22, 0xbc, 0x02, 0x0a, 0x0a,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x4e, 0x0a, 0x12, 0x61, 0x64,
	0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x10, 0x61, 0x64, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x12, 0x4b, 0x0a, 0x10, 0x79, 0x6f,
	0x75, 0x74, 0x75, 0x62, 0x65, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x0e, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x51, 0x0a, 0x13, 0x61, 0x64, 0x76, 0x65, 0x72,
	0x74, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x11, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69,
	0x73, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x69, 0x73,
	0x63, 0x69, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x08, 0x69, 0x73, 0x63, 0x69, 0x43, 0x6f, 0x64, 0x65, 0x42, 0xfb, 0x01, 0x0a, 0x25, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x42, 0x0e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e,
	0x56, 0x32, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x5c, 0x56, 0x32, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x32, 0x3a, 0x3a, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v2_resources_media_file_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v2_resources_media_file_proto_rawDescData = file_google_ads_googleads_v2_resources_media_file_proto_rawDesc
)

func file_google_ads_googleads_v2_resources_media_file_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v2_resources_media_file_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v2_resources_media_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v2_resources_media_file_proto_rawDescData)
	})
	return file_google_ads_googleads_v2_resources_media_file_proto_rawDescData
}

var file_google_ads_googleads_v2_resources_media_file_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_ads_googleads_v2_resources_media_file_proto_goTypes = []interface{}{
	(*MediaFile)(nil),                  // 0: google.ads.googleads.v2.resources.MediaFile
	(*MediaImage)(nil),                 // 1: google.ads.googleads.v2.resources.MediaImage
	(*MediaBundle)(nil),                // 2: google.ads.googleads.v2.resources.MediaBundle
	(*MediaAudio)(nil),                 // 3: google.ads.googleads.v2.resources.MediaAudio
	(*MediaVideo)(nil),                 // 4: google.ads.googleads.v2.resources.MediaVideo
	(*wrappers.Int64Value)(nil),        // 5: google.protobuf.Int64Value
	(enums.MediaTypeEnum_MediaType)(0), // 6: google.ads.googleads.v2.enums.MediaTypeEnum.MediaType
	(enums.MimeTypeEnum_MimeType)(0),   // 7: google.ads.googleads.v2.enums.MimeTypeEnum.MimeType
	(*wrappers.StringValue)(nil),       // 8: google.protobuf.StringValue
	(*wrappers.BytesValue)(nil),        // 9: google.protobuf.BytesValue
}
var file_google_ads_googleads_v2_resources_media_file_proto_depIdxs = []int32{
	5,  // 0: google.ads.googleads.v2.resources.MediaFile.id:type_name -> google.protobuf.Int64Value
	6,  // 1: google.ads.googleads.v2.resources.MediaFile.type:type_name -> google.ads.googleads.v2.enums.MediaTypeEnum.MediaType
	7,  // 2: google.ads.googleads.v2.resources.MediaFile.mime_type:type_name -> google.ads.googleads.v2.enums.MimeTypeEnum.MimeType
	8,  // 3: google.ads.googleads.v2.resources.MediaFile.source_url:type_name -> google.protobuf.StringValue
	8,  // 4: google.ads.googleads.v2.resources.MediaFile.name:type_name -> google.protobuf.StringValue
	5,  // 5: google.ads.googleads.v2.resources.MediaFile.file_size:type_name -> google.protobuf.Int64Value
	1,  // 6: google.ads.googleads.v2.resources.MediaFile.image:type_name -> google.ads.googleads.v2.resources.MediaImage
	2,  // 7: google.ads.googleads.v2.resources.MediaFile.media_bundle:type_name -> google.ads.googleads.v2.resources.MediaBundle
	3,  // 8: google.ads.googleads.v2.resources.MediaFile.audio:type_name -> google.ads.googleads.v2.resources.MediaAudio
	4,  // 9: google.ads.googleads.v2.resources.MediaFile.video:type_name -> google.ads.googleads.v2.resources.MediaVideo
	9,  // 10: google.ads.googleads.v2.resources.MediaImage.data:type_name -> google.protobuf.BytesValue
	9,  // 11: google.ads.googleads.v2.resources.MediaBundle.data:type_name -> google.protobuf.BytesValue
	5,  // 12: google.ads.googleads.v2.resources.MediaAudio.ad_duration_millis:type_name -> google.protobuf.Int64Value
	5,  // 13: google.ads.googleads.v2.resources.MediaVideo.ad_duration_millis:type_name -> google.protobuf.Int64Value
	8,  // 14: google.ads.googleads.v2.resources.MediaVideo.youtube_video_id:type_name -> google.protobuf.StringValue
	8,  // 15: google.ads.googleads.v2.resources.MediaVideo.advertising_id_code:type_name -> google.protobuf.StringValue
	8,  // 16: google.ads.googleads.v2.resources.MediaVideo.isci_code:type_name -> google.protobuf.StringValue
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v2_resources_media_file_proto_init() }
func file_google_ads_googleads_v2_resources_media_file_proto_init() {
	if File_google_ads_googleads_v2_resources_media_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v2_resources_media_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaFile); i {
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
		file_google_ads_googleads_v2_resources_media_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaImage); i {
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
		file_google_ads_googleads_v2_resources_media_file_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaBundle); i {
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
		file_google_ads_googleads_v2_resources_media_file_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaAudio); i {
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
		file_google_ads_googleads_v2_resources_media_file_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaVideo); i {
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
	file_google_ads_googleads_v2_resources_media_file_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*MediaFile_Image)(nil),
		(*MediaFile_MediaBundle)(nil),
		(*MediaFile_Audio)(nil),
		(*MediaFile_Video)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v2_resources_media_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v2_resources_media_file_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v2_resources_media_file_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v2_resources_media_file_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v2_resources_media_file_proto = out.File
	file_google_ads_googleads_v2_resources_media_file_proto_rawDesc = nil
	file_google_ads_googleads_v2_resources_media_file_proto_goTypes = nil
	file_google_ads_googleads_v2_resources_media_file_proto_depIdxs = nil
}
