// Copyright 2016 Google Inc.
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
// source: google/appengine/v1/application.proto

package appengine

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

// An Application resource contains the top-level configuration of an App
// Engine application.
type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Full path to the Application resource in the API.
	// Example: `apps/myapp`.
	//
	// @OutputOnly
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Identifier of the Application resource. This identifier is equivalent
	// to the project ID of the Google Cloud Platform project where you want to
	// deploy your application.
	// Example: `myapp`.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// HTTP path dispatch rules for requests to the application that do not
	// explicitly target a service or version. Rules are order-dependent.
	//
	// @OutputOnly
	DispatchRules []*UrlDispatchRule `protobuf:"bytes,3,rep,name=dispatch_rules,json=dispatchRules,proto3" json:"dispatch_rules,omitempty"`
	// Google Apps authentication domain that controls which users can access
	// this application.
	//
	// Defaults to open access for any Google Account.
	AuthDomain string `protobuf:"bytes,6,opt,name=auth_domain,json=authDomain,proto3" json:"auth_domain,omitempty"`
	// Location from which this application will be run. Application instances
	// will run out of data centers in the chosen location, which is also where
	// all of the application's end user content is stored.
	//
	// Defaults to `us-central`.
	//
	// Options are:
	//
	// `us-central` - Central US
	//
	// `europe-west` - Western Europe
	//
	// `us-east1` - Eastern US
	LocationId string `protobuf:"bytes,7,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	// Google Cloud Storage bucket that can be used for storing files
	// associated with this application. This bucket is associated with the
	// application and can be used by the gcloud deployment commands.
	//
	// @OutputOnly
	CodeBucket string `protobuf:"bytes,8,opt,name=code_bucket,json=codeBucket,proto3" json:"code_bucket,omitempty"`
	// Cookie expiration policy for this application.
	//
	// @OutputOnly
	DefaultCookieExpiration *duration.Duration `protobuf:"bytes,9,opt,name=default_cookie_expiration,json=defaultCookieExpiration,proto3" json:"default_cookie_expiration,omitempty"`
	// Hostname used to reach this application, as resolved by App Engine.
	//
	// @OutputOnly
	DefaultHostname string `protobuf:"bytes,11,opt,name=default_hostname,json=defaultHostname,proto3" json:"default_hostname,omitempty"`
	// Google Cloud Storage bucket that can be used by this application to store
	// content.
	//
	// @OutputOnly
	DefaultBucket string `protobuf:"bytes,12,opt,name=default_bucket,json=defaultBucket,proto3" json:"default_bucket,omitempty"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_v1_application_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_v1_application_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_google_appengine_v1_application_proto_rawDescGZIP(), []int{0}
}

func (x *Application) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Application) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Application) GetDispatchRules() []*UrlDispatchRule {
	if x != nil {
		return x.DispatchRules
	}
	return nil
}

func (x *Application) GetAuthDomain() string {
	if x != nil {
		return x.AuthDomain
	}
	return ""
}

func (x *Application) GetLocationId() string {
	if x != nil {
		return x.LocationId
	}
	return ""
}

func (x *Application) GetCodeBucket() string {
	if x != nil {
		return x.CodeBucket
	}
	return ""
}

func (x *Application) GetDefaultCookieExpiration() *duration.Duration {
	if x != nil {
		return x.DefaultCookieExpiration
	}
	return nil
}

func (x *Application) GetDefaultHostname() string {
	if x != nil {
		return x.DefaultHostname
	}
	return ""
}

func (x *Application) GetDefaultBucket() string {
	if x != nil {
		return x.DefaultBucket
	}
	return ""
}

// Rules to match an HTTP request and dispatch that request to a service.
type UrlDispatchRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Domain name to match against. The wildcard "`*`" is supported if
	// specified before a period: "`*.`".
	//
	// Defaults to matching all domains: "`*`".
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// Pathname within the host. Must start with a "`/`". A
	// single "`*`" can be included at the end of the path. The sum
	// of the lengths of the domain and path may not exceed 100
	// characters.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// Resource ID of a service in this application that should
	// serve the matched request. The service must already
	// exist. Example: `default`.
	Service string `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *UrlDispatchRule) Reset() {
	*x = UrlDispatchRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_v1_application_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlDispatchRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlDispatchRule) ProtoMessage() {}

func (x *UrlDispatchRule) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_v1_application_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlDispatchRule.ProtoReflect.Descriptor instead.
func (*UrlDispatchRule) Descriptor() ([]byte, []int) {
	return file_google_appengine_v1_application_proto_rawDescGZIP(), []int{1}
}

func (x *UrlDispatchRule) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *UrlDispatchRule) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *UrlDispatchRule) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

var File_google_appengine_v1_application_proto protoreflect.FileDescriptor

var file_google_appengine_v1_application_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x03, 0x0a, 0x0b, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x4b,
	0x0a, 0x0e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x72, 0x6c,
	0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x0d, 0x64, 0x69,
	0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x0b,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x55,
	0x0a, 0x19, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65,
	0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x17, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x5f, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x57, 0x0a, 0x0f, 0x55, 0x72, 0x6c, 0x44, 0x69,
	0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x42, 0x6b, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_appengine_v1_application_proto_rawDescOnce sync.Once
	file_google_appengine_v1_application_proto_rawDescData = file_google_appengine_v1_application_proto_rawDesc
)

func file_google_appengine_v1_application_proto_rawDescGZIP() []byte {
	file_google_appengine_v1_application_proto_rawDescOnce.Do(func() {
		file_google_appengine_v1_application_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_appengine_v1_application_proto_rawDescData)
	})
	return file_google_appengine_v1_application_proto_rawDescData
}

var file_google_appengine_v1_application_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_appengine_v1_application_proto_goTypes = []interface{}{
	(*Application)(nil),       // 0: google.appengine.v1.Application
	(*UrlDispatchRule)(nil),   // 1: google.appengine.v1.UrlDispatchRule
	(*duration.Duration)(nil), // 2: google.protobuf.Duration
}
var file_google_appengine_v1_application_proto_depIdxs = []int32{
	1, // 0: google.appengine.v1.Application.dispatch_rules:type_name -> google.appengine.v1.UrlDispatchRule
	2, // 1: google.appengine.v1.Application.default_cookie_expiration:type_name -> google.protobuf.Duration
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_appengine_v1_application_proto_init() }
func file_google_appengine_v1_application_proto_init() {
	if File_google_appengine_v1_application_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_appengine_v1_application_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
		file_google_appengine_v1_application_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UrlDispatchRule); i {
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
			RawDescriptor: file_google_appengine_v1_application_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_appengine_v1_application_proto_goTypes,
		DependencyIndexes: file_google_appengine_v1_application_proto_depIdxs,
		MessageInfos:      file_google_appengine_v1_application_proto_msgTypes,
	}.Build()
	File_google_appengine_v1_application_proto = out.File
	file_google_appengine_v1_application_proto_rawDesc = nil
	file_google_appengine_v1_application_proto_goTypes = nil
	file_google_appengine_v1_application_proto_depIdxs = nil
}
