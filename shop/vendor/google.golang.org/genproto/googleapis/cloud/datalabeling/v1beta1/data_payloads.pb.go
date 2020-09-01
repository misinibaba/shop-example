// Copyright 2019 Google LLC.
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
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.2
// source: google/cloud/datalabeling/v1beta1/data_payloads.proto

package datalabeling

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

// Container of information about an image.
type ImagePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Image format.
	MimeType string `protobuf:"bytes,1,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// A byte string of a thumbnail image.
	ImageThumbnail []byte `protobuf:"bytes,2,opt,name=image_thumbnail,json=imageThumbnail,proto3" json:"image_thumbnail,omitempty"`
	// Image uri from the user bucket.
	ImageUri string `protobuf:"bytes,3,opt,name=image_uri,json=imageUri,proto3" json:"image_uri,omitempty"`
	// Signed uri of the image file in the service bucket.
	SignedUri string `protobuf:"bytes,4,opt,name=signed_uri,json=signedUri,proto3" json:"signed_uri,omitempty"`
}

func (x *ImagePayload) Reset() {
	*x = ImagePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datalabeling_v1beta1_data_payloads_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImagePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImagePayload) ProtoMessage() {}

func (x *ImagePayload) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datalabeling_v1beta1_data_payloads_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImagePayload.ProtoReflect.Descriptor instead.
func (*ImagePayload) Descriptor() ([]byte, []int) {
	return file_google_cloud_datalabeling_v1beta1_data_payloads_proto_rawDescGZIP(), []int{0}
}

func (x *ImagePayload) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *ImagePayload) GetImageThumbnail() []byte {
	if x != nil {
		return x.ImageThumbnail
	}
	return nil
}

func (x *ImagePayload) GetImageUri() string {
	if x != nil {
		return x.ImageUri
	}
	return ""
}

func (x *ImagePayload) GetSignedUri() string {
	if x != nil {
		return x.SignedUri
	}
	return ""
}

// Container of information about a piece of text.
type TextPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Text content.
	TextContent string `protobuf:"bytes,1,opt,name=text_content,json=textContent,proto3" json:"text_content,omitempty"`
}

func (x *TextPayload) Reset() {
	*x = TextPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datalabeling_v1beta1_data_payloads_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextPayload) ProtoMessage() {}

func (x *TextPayload) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datalabeling_v1beta1_data_payloads_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextPayload.ProtoReflect.Descriptor instead.
func (*TextPayload) Descriptor() ([]byte, []int) {
	return file_google_cloud_datalabeling_v1beta1_data_payloads_proto_rawDescGZIP(), []int{1}
}

func (x *TextPayload) GetTextContent() string {
	if x != nil {
		return x.TextContent
	}
	return ""
}

// Container of information of a video thumbnail.
type VideoThumbnail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A byte string of the video frame.
	Thumbnail []byte `protobuf:"bytes,1,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	// Time offset relative to the beginning of the video, corresponding to the
	// video frame where the thumbnail has been extracted from.
	TimeOffset *duration.Duration `protobuf:"bytes,2,opt,name=time_offset,json=timeOffset,proto3" json:"time_offset,omitempty"`
}

func (x *VideoThumbnail) Reset() {
	*x = VideoThumbnail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datalabeling_v1beta1_data_payloads_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoThumbnail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoThumbnail) ProtoMessage() {}

func (x *VideoThumbnail) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datalabeling_v1beta1_data_payloads_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoThumbnail.ProtoReflect.Descriptor instead.
func (*VideoThumbnail) Descriptor() ([]byte, []int) {
	return file_google_cloud_datalabeling_v1beta1_data_payloads_proto_rawDescGZIP(), []int{2}
}

func (x *VideoThumbnail) GetThumbnail() []byte {
	if x != nil {
		return x.Thumbnail
	}
	return nil
}

func (x *VideoThumbnail) GetTimeOffset() *duration.Duration {
	if x != nil {
		return x.TimeOffset
	}
	return nil
}

// Container of information of a video.
type VideoPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Video format.
	MimeType string `protobuf:"bytes,1,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// Video uri from the user bucket.
	VideoUri string `protobuf:"bytes,2,opt,name=video_uri,json=videoUri,proto3" json:"video_uri,omitempty"`
	// The list of video thumbnails.
	VideoThumbnails []*VideoThumbnail `protobuf:"bytes,3,rep,name=video_thumbnails,json=videoThumbnails,proto3" json:"video_thumbnails,omitempty"`
	// FPS of the video.
	FrameRate float32 `protobuf:"fixed32,4,opt,name=frame_rate,json=frameRate,proto3" json:"frame_rate,omitempty"`
	// Signed uri of the video file in the service bucket.
	SignedUri string `protobuf:"bytes,5,opt,name=signed_uri,json=signedUri,proto3" json:"signed_uri,omitempty"`
}

func (x *VideoPayload) Reset() {
	*x = VideoPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datalabeling_v1beta1_data_payloads_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoPayload) ProtoMessage() {}

func (x *VideoPayload) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datalabeling_v1beta1_data_payloads_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoPayload.ProtoReflect.Descriptor instead.
func (*VideoPayload) Descriptor() ([]byte, []int) {
	return file_google_cloud_datalabeling_v1beta1_data_payloads_proto_rawDescGZIP(), []int{3}
}

func (x *VideoPayload) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *VideoPayload) GetVideoUri() string {
	if x != nil {
		return x.VideoUri
	}
	return ""
}

func (x *VideoPayload) GetVideoThumbnails() []*VideoThumbnail {
	if x != nil {
		return x.VideoThumbnails
	}
	return nil
}

func (x *VideoPayload) GetFrameRate() float32 {
	if x != nil {
		return x.FrameRate
	}
	return 0
}

func (x *VideoPayload) GetSignedUri() string {
	if x != nil {
		return x.SignedUri
	}
	return ""
}

var File_google_cloud_datalabeling_v1beta1_data_payloads_proto protoreflect.FileDescriptor

var file_google_cloud_datalabeling_v1beta1_data_payloads_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a, 0x0c, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69,
	0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d,
	0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x69, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x72, 0x69, 0x22, 0x30, 0x0a, 0x0b,
	0x54, 0x65, 0x78, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x74,
	0x65, 0x78, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x74, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x6a,
	0x0a, 0x0e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x3a,
	0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a,
	0x74, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0xe4, 0x01, 0x0a, 0x0c, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d,
	0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x55, 0x72, 0x69, 0x12, 0x5c, 0x0a, 0x10, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x74,
	0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61,
	0x69, 0x6c, 0x52, 0x0f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x61,
	0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x75, 0x72, 0x69,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x72,
	0x69, 0x42, 0x78, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x50, 0x01, 0x5a, 0x4d, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x64,
	0x61, 0x74, 0x61, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_datalabeling_v1beta1_data_payloads_proto_rawDescOnce sync.Once
	file_google_cloud_datalabeling_v1beta1_data_payloads_proto_rawDescData = file_google_cloud_datalabeling_v1beta1_data_payloads_proto_rawDesc
)

func file_google_cloud_datalabeling_v1beta1_data_payloads_proto_rawDescGZIP() []byte {
	file_google_cloud_datalabeling_v1beta1_data_payloads_proto_rawDescOnce.Do(func() {
		file_google_cloud_datalabeling_v1beta1_data_payloads_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_datalabeling_v1beta1_data_payloads_proto_rawDescData)
	})
	return file_google_cloud_datalabeling_v1beta1_data_payloads_proto_rawDescData
}

var file_google_cloud_datalabeling_v1beta1_data_payloads_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_datalabeling_v1beta1_data_payloads_proto_goTypes = []interface{}{
	(*ImagePayload)(nil),      // 0: google.cloud.datalabeling.v1beta1.ImagePayload
	(*TextPayload)(nil),       // 1: google.cloud.datalabeling.v1beta1.TextPayload
	(*VideoThumbnail)(nil),    // 2: google.cloud.datalabeling.v1beta1.VideoThumbnail
	(*VideoPayload)(nil),      // 3: google.cloud.datalabeling.v1beta1.VideoPayload
	(*duration.Duration)(nil), // 4: google.protobuf.Duration
}
var file_google_cloud_datalabeling_v1beta1_data_payloads_proto_depIdxs = []int32{
	4, // 0: google.cloud.datalabeling.v1beta1.VideoThumbnail.time_offset:type_name -> google.protobuf.Duration
	2, // 1: google.cloud.datalabeling.v1beta1.VideoPayload.video_thumbnails:type_name -> google.cloud.datalabeling.v1beta1.VideoThumbnail
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_datalabeling_v1beta1_data_payloads_proto_init() }
func file_google_cloud_datalabeling_v1beta1_data_payloads_proto_init() {
	if File_google_cloud_datalabeling_v1beta1_data_payloads_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_datalabeling_v1beta1_data_payloads_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImagePayload); i {
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
		file_google_cloud_datalabeling_v1beta1_data_payloads_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextPayload); i {
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
		file_google_cloud_datalabeling_v1beta1_data_payloads_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoThumbnail); i {
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
		file_google_cloud_datalabeling_v1beta1_data_payloads_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoPayload); i {
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
			RawDescriptor: file_google_cloud_datalabeling_v1beta1_data_payloads_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_datalabeling_v1beta1_data_payloads_proto_goTypes,
		DependencyIndexes: file_google_cloud_datalabeling_v1beta1_data_payloads_proto_depIdxs,
		MessageInfos:      file_google_cloud_datalabeling_v1beta1_data_payloads_proto_msgTypes,
	}.Build()
	File_google_cloud_datalabeling_v1beta1_data_payloads_proto = out.File
	file_google_cloud_datalabeling_v1beta1_data_payloads_proto_rawDesc = nil
	file_google_cloud_datalabeling_v1beta1_data_payloads_proto_goTypes = nil
	file_google_cloud_datalabeling_v1beta1_data_payloads_proto_depIdxs = nil
}
