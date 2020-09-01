// Copyright 2018 Google Inc.
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
// source: google/bigtable/v1/bigtable_service.proto

package bigtable

import (
	context "context"
	reflect "reflect"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

var File_google_bigtable_v1_bigtable_service_proto protoreflect.FileDescriptor

var file_google_bigtable_v1_bigtable_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x62, 0x69,
	0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xdd, 0x08, 0x0a, 0x0f, 0x42, 0x69, 0x67, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa5, 0x01, 0x0a, 0x08, 0x52,
	0x65, 0x61, 0x64, 0x52, 0x6f, 0x77, 0x73, 0x12, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61,
	0x64, 0x52, 0x6f, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x6f, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x4c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x46, 0x22, 0x41, 0x2f, 0x76, 0x31, 0x2f,
	0x7b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x2f, 0x2a, 0x7d, 0x2f, 0x72, 0x6f, 0x77, 0x73, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x3a, 0x01, 0x2a,
	0x30, 0x01, 0x12, 0xb7, 0x01, 0x0a, 0x0d, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x6f, 0x77,
	0x4b, 0x65, 0x79, 0x73, 0x12, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69,
	0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x52, 0x6f, 0x77, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x6f, 0x77, 0x4b, 0x65, 0x79,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x49, 0x12, 0x47, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x7a, 0x6f,
	0x6e, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x2a,
	0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x72, 0x6f, 0x77, 0x73, 0x3a,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x30, 0x01, 0x12, 0xa3, 0x01, 0x0a,
	0x09, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x77, 0x12, 0x24, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x58, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x52,
	0x22, 0x4d, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x7a, 0x6f, 0x6e,
	0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x72, 0x6f, 0x77, 0x73, 0x2f, 0x7b,
	0x72, 0x6f, 0x77, 0x5f, 0x6b, 0x65, 0x79, 0x7d, 0x3a, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x3a,
	0x01, 0x2a, 0x12, 0xaa, 0x01, 0x0a, 0x0a, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x77,
	0x73, 0x12, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x77,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x75,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x4d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x47, 0x22, 0x42, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2f, 0x2a,
	0x7d, 0x3a, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x77, 0x73, 0x3a, 0x01, 0x2a, 0x12,
	0xd2, 0x01, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x6e, 0x64, 0x4d, 0x75, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x77, 0x12, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62,
	0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x41, 0x6e, 0x64, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x6e,
	0x64, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x60, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x5a, 0x22, 0x55, 0x2f, 0x76, 0x31, 0x2f,
	0x7b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x2f, 0x2a, 0x7d, 0x2f, 0x72, 0x6f, 0x77, 0x73, 0x2f, 0x7b, 0x72, 0x6f, 0x77, 0x5f, 0x6b, 0x65,
	0x79, 0x7d, 0x3a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x6e, 0x64, 0x4d, 0x75, 0x74, 0x61, 0x74,
	0x65, 0x3a, 0x01, 0x2a, 0x12, 0xbf, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x61, 0x64, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x6f, 0x77, 0x12, 0x2d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x61, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x52, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x6f, 0x77, 0x22, 0x61, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x5b, 0x22, 0x56, 0x2f, 0x76, 0x31,
	0x2f, 0x7b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x2f, 0x2a, 0x2f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x72, 0x6f, 0x77, 0x73, 0x2f, 0x7b, 0x72, 0x6f, 0x77, 0x5f, 0x6b,
	0x65, 0x79, 0x7d, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x70, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x42, 0x15, 0x42, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x69, 0x67,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_google_bigtable_v1_bigtable_service_proto_goTypes = []interface{}{
	(*ReadRowsRequest)(nil),           // 0: google.bigtable.v1.ReadRowsRequest
	(*SampleRowKeysRequest)(nil),      // 1: google.bigtable.v1.SampleRowKeysRequest
	(*MutateRowRequest)(nil),          // 2: google.bigtable.v1.MutateRowRequest
	(*MutateRowsRequest)(nil),         // 3: google.bigtable.v1.MutateRowsRequest
	(*CheckAndMutateRowRequest)(nil),  // 4: google.bigtable.v1.CheckAndMutateRowRequest
	(*ReadModifyWriteRowRequest)(nil), // 5: google.bigtable.v1.ReadModifyWriteRowRequest
	(*ReadRowsResponse)(nil),          // 6: google.bigtable.v1.ReadRowsResponse
	(*SampleRowKeysResponse)(nil),     // 7: google.bigtable.v1.SampleRowKeysResponse
	(*empty.Empty)(nil),               // 8: google.protobuf.Empty
	(*MutateRowsResponse)(nil),        // 9: google.bigtable.v1.MutateRowsResponse
	(*CheckAndMutateRowResponse)(nil), // 10: google.bigtable.v1.CheckAndMutateRowResponse
	(*Row)(nil),                       // 11: google.bigtable.v1.Row
}
var file_google_bigtable_v1_bigtable_service_proto_depIdxs = []int32{
	0,  // 0: google.bigtable.v1.BigtableService.ReadRows:input_type -> google.bigtable.v1.ReadRowsRequest
	1,  // 1: google.bigtable.v1.BigtableService.SampleRowKeys:input_type -> google.bigtable.v1.SampleRowKeysRequest
	2,  // 2: google.bigtable.v1.BigtableService.MutateRow:input_type -> google.bigtable.v1.MutateRowRequest
	3,  // 3: google.bigtable.v1.BigtableService.MutateRows:input_type -> google.bigtable.v1.MutateRowsRequest
	4,  // 4: google.bigtable.v1.BigtableService.CheckAndMutateRow:input_type -> google.bigtable.v1.CheckAndMutateRowRequest
	5,  // 5: google.bigtable.v1.BigtableService.ReadModifyWriteRow:input_type -> google.bigtable.v1.ReadModifyWriteRowRequest
	6,  // 6: google.bigtable.v1.BigtableService.ReadRows:output_type -> google.bigtable.v1.ReadRowsResponse
	7,  // 7: google.bigtable.v1.BigtableService.SampleRowKeys:output_type -> google.bigtable.v1.SampleRowKeysResponse
	8,  // 8: google.bigtable.v1.BigtableService.MutateRow:output_type -> google.protobuf.Empty
	9,  // 9: google.bigtable.v1.BigtableService.MutateRows:output_type -> google.bigtable.v1.MutateRowsResponse
	10, // 10: google.bigtable.v1.BigtableService.CheckAndMutateRow:output_type -> google.bigtable.v1.CheckAndMutateRowResponse
	11, // 11: google.bigtable.v1.BigtableService.ReadModifyWriteRow:output_type -> google.bigtable.v1.Row
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_google_bigtable_v1_bigtable_service_proto_init() }
func file_google_bigtable_v1_bigtable_service_proto_init() {
	if File_google_bigtable_v1_bigtable_service_proto != nil {
		return
	}
	file_google_bigtable_v1_bigtable_data_proto_init()
	file_google_bigtable_v1_bigtable_service_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_bigtable_v1_bigtable_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_bigtable_v1_bigtable_service_proto_goTypes,
		DependencyIndexes: file_google_bigtable_v1_bigtable_service_proto_depIdxs,
	}.Build()
	File_google_bigtable_v1_bigtable_service_proto = out.File
	file_google_bigtable_v1_bigtable_service_proto_rawDesc = nil
	file_google_bigtable_v1_bigtable_service_proto_goTypes = nil
	file_google_bigtable_v1_bigtable_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BigtableServiceClient is the client API for BigtableService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BigtableServiceClient interface {
	// Streams back the contents of all requested rows, optionally applying
	// the same Reader filter to each. Depending on their size, rows may be
	// broken up across multiple responses, but atomicity of each row will still
	// be preserved.
	ReadRows(ctx context.Context, in *ReadRowsRequest, opts ...grpc.CallOption) (BigtableService_ReadRowsClient, error)
	// Returns a sample of row keys in the table. The returned row keys will
	// delimit contiguous sections of the table of approximately equal size,
	// which can be used to break up the data for distributed tasks like
	// mapreduces.
	SampleRowKeys(ctx context.Context, in *SampleRowKeysRequest, opts ...grpc.CallOption) (BigtableService_SampleRowKeysClient, error)
	// Mutates a row atomically. Cells already present in the row are left
	// unchanged unless explicitly changed by 'mutation'.
	MutateRow(ctx context.Context, in *MutateRowRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Mutates multiple rows in a batch. Each individual row is mutated
	// atomically as in MutateRow, but the entire batch is not executed
	// atomically.
	MutateRows(ctx context.Context, in *MutateRowsRequest, opts ...grpc.CallOption) (*MutateRowsResponse, error)
	// Mutates a row atomically based on the output of a predicate Reader filter.
	CheckAndMutateRow(ctx context.Context, in *CheckAndMutateRowRequest, opts ...grpc.CallOption) (*CheckAndMutateRowResponse, error)
	// Modifies a row atomically, reading the latest existing timestamp/value from
	// the specified columns and writing a new value at
	// max(existing timestamp, current server time) based on pre-defined
	// read/modify/write rules. Returns the new contents of all modified cells.
	ReadModifyWriteRow(ctx context.Context, in *ReadModifyWriteRowRequest, opts ...grpc.CallOption) (*Row, error)
}

type bigtableServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBigtableServiceClient(cc grpc.ClientConnInterface) BigtableServiceClient {
	return &bigtableServiceClient{cc}
}

func (c *bigtableServiceClient) ReadRows(ctx context.Context, in *ReadRowsRequest, opts ...grpc.CallOption) (BigtableService_ReadRowsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BigtableService_serviceDesc.Streams[0], "/google.bigtable.v1.BigtableService/ReadRows", opts...)
	if err != nil {
		return nil, err
	}
	x := &bigtableServiceReadRowsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BigtableService_ReadRowsClient interface {
	Recv() (*ReadRowsResponse, error)
	grpc.ClientStream
}

type bigtableServiceReadRowsClient struct {
	grpc.ClientStream
}

func (x *bigtableServiceReadRowsClient) Recv() (*ReadRowsResponse, error) {
	m := new(ReadRowsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bigtableServiceClient) SampleRowKeys(ctx context.Context, in *SampleRowKeysRequest, opts ...grpc.CallOption) (BigtableService_SampleRowKeysClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BigtableService_serviceDesc.Streams[1], "/google.bigtable.v1.BigtableService/SampleRowKeys", opts...)
	if err != nil {
		return nil, err
	}
	x := &bigtableServiceSampleRowKeysClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BigtableService_SampleRowKeysClient interface {
	Recv() (*SampleRowKeysResponse, error)
	grpc.ClientStream
}

type bigtableServiceSampleRowKeysClient struct {
	grpc.ClientStream
}

func (x *bigtableServiceSampleRowKeysClient) Recv() (*SampleRowKeysResponse, error) {
	m := new(SampleRowKeysResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bigtableServiceClient) MutateRow(ctx context.Context, in *MutateRowRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.bigtable.v1.BigtableService/MutateRow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableServiceClient) MutateRows(ctx context.Context, in *MutateRowsRequest, opts ...grpc.CallOption) (*MutateRowsResponse, error) {
	out := new(MutateRowsResponse)
	err := c.cc.Invoke(ctx, "/google.bigtable.v1.BigtableService/MutateRows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableServiceClient) CheckAndMutateRow(ctx context.Context, in *CheckAndMutateRowRequest, opts ...grpc.CallOption) (*CheckAndMutateRowResponse, error) {
	out := new(CheckAndMutateRowResponse)
	err := c.cc.Invoke(ctx, "/google.bigtable.v1.BigtableService/CheckAndMutateRow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableServiceClient) ReadModifyWriteRow(ctx context.Context, in *ReadModifyWriteRowRequest, opts ...grpc.CallOption) (*Row, error) {
	out := new(Row)
	err := c.cc.Invoke(ctx, "/google.bigtable.v1.BigtableService/ReadModifyWriteRow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BigtableServiceServer is the server API for BigtableService service.
type BigtableServiceServer interface {
	// Streams back the contents of all requested rows, optionally applying
	// the same Reader filter to each. Depending on their size, rows may be
	// broken up across multiple responses, but atomicity of each row will still
	// be preserved.
	ReadRows(*ReadRowsRequest, BigtableService_ReadRowsServer) error
	// Returns a sample of row keys in the table. The returned row keys will
	// delimit contiguous sections of the table of approximately equal size,
	// which can be used to break up the data for distributed tasks like
	// mapreduces.
	SampleRowKeys(*SampleRowKeysRequest, BigtableService_SampleRowKeysServer) error
	// Mutates a row atomically. Cells already present in the row are left
	// unchanged unless explicitly changed by 'mutation'.
	MutateRow(context.Context, *MutateRowRequest) (*empty.Empty, error)
	// Mutates multiple rows in a batch. Each individual row is mutated
	// atomically as in MutateRow, but the entire batch is not executed
	// atomically.
	MutateRows(context.Context, *MutateRowsRequest) (*MutateRowsResponse, error)
	// Mutates a row atomically based on the output of a predicate Reader filter.
	CheckAndMutateRow(context.Context, *CheckAndMutateRowRequest) (*CheckAndMutateRowResponse, error)
	// Modifies a row atomically, reading the latest existing timestamp/value from
	// the specified columns and writing a new value at
	// max(existing timestamp, current server time) based on pre-defined
	// read/modify/write rules. Returns the new contents of all modified cells.
	ReadModifyWriteRow(context.Context, *ReadModifyWriteRowRequest) (*Row, error)
}

// UnimplementedBigtableServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBigtableServiceServer struct {
}

func (*UnimplementedBigtableServiceServer) ReadRows(*ReadRowsRequest, BigtableService_ReadRowsServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadRows not implemented")
}
func (*UnimplementedBigtableServiceServer) SampleRowKeys(*SampleRowKeysRequest, BigtableService_SampleRowKeysServer) error {
	return status.Errorf(codes.Unimplemented, "method SampleRowKeys not implemented")
}
func (*UnimplementedBigtableServiceServer) MutateRow(context.Context, *MutateRowRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateRow not implemented")
}
func (*UnimplementedBigtableServiceServer) MutateRows(context.Context, *MutateRowsRequest) (*MutateRowsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateRows not implemented")
}
func (*UnimplementedBigtableServiceServer) CheckAndMutateRow(context.Context, *CheckAndMutateRowRequest) (*CheckAndMutateRowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAndMutateRow not implemented")
}
func (*UnimplementedBigtableServiceServer) ReadModifyWriteRow(context.Context, *ReadModifyWriteRowRequest) (*Row, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadModifyWriteRow not implemented")
}

func RegisterBigtableServiceServer(s *grpc.Server, srv BigtableServiceServer) {
	s.RegisterService(&_BigtableService_serviceDesc, srv)
}

func _BigtableService_ReadRows_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadRowsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BigtableServiceServer).ReadRows(m, &bigtableServiceReadRowsServer{stream})
}

type BigtableService_ReadRowsServer interface {
	Send(*ReadRowsResponse) error
	grpc.ServerStream
}

type bigtableServiceReadRowsServer struct {
	grpc.ServerStream
}

func (x *bigtableServiceReadRowsServer) Send(m *ReadRowsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BigtableService_SampleRowKeys_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SampleRowKeysRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BigtableServiceServer).SampleRowKeys(m, &bigtableServiceSampleRowKeysServer{stream})
}

type BigtableService_SampleRowKeysServer interface {
	Send(*SampleRowKeysResponse) error
	grpc.ServerStream
}

type bigtableServiceSampleRowKeysServer struct {
	grpc.ServerStream
}

func (x *bigtableServiceSampleRowKeysServer) Send(m *SampleRowKeysResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BigtableService_MutateRow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateRowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableServiceServer).MutateRow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.v1.BigtableService/MutateRow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableServiceServer).MutateRow(ctx, req.(*MutateRowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableService_MutateRows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateRowsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableServiceServer).MutateRows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.v1.BigtableService/MutateRows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableServiceServer).MutateRows(ctx, req.(*MutateRowsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableService_CheckAndMutateRow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAndMutateRowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableServiceServer).CheckAndMutateRow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.v1.BigtableService/CheckAndMutateRow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableServiceServer).CheckAndMutateRow(ctx, req.(*CheckAndMutateRowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableService_ReadModifyWriteRow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadModifyWriteRowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableServiceServer).ReadModifyWriteRow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.v1.BigtableService/ReadModifyWriteRow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableServiceServer).ReadModifyWriteRow(ctx, req.(*ReadModifyWriteRowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BigtableService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.bigtable.v1.BigtableService",
	HandlerType: (*BigtableServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateRow",
			Handler:    _BigtableService_MutateRow_Handler,
		},
		{
			MethodName: "MutateRows",
			Handler:    _BigtableService_MutateRows_Handler,
		},
		{
			MethodName: "CheckAndMutateRow",
			Handler:    _BigtableService_CheckAndMutateRow_Handler,
		},
		{
			MethodName: "ReadModifyWriteRow",
			Handler:    _BigtableService_ReadModifyWriteRow_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReadRows",
			Handler:       _BigtableService_ReadRows_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SampleRowKeys",
			Handler:       _BigtableService_SampleRowKeys_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "google/bigtable/v1/bigtable_service.proto",
}
