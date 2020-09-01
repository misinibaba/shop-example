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
// source: google/identity/accesscontextmanager/v1/service_perimeter.proto

package accesscontextmanager

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

// Specifies the type of the Perimeter. There are two types: regular and
// bridge. Regular Service Perimeter contains resources, access levels, and
// restricted services. Every resource can be in at most ONE
// regular Service Perimeter.
//
// In addition to being in a regular service perimeter, a resource can also
// be in zero or more perimeter bridges.  A perimeter bridge only contains
// resources.  Cross project operations are permitted if all effected
// resources share some perimeter (whether bridge or regular). Perimeter
// Bridge does not contain access levels or services: those are governed
// entirely by the regular perimeter that resource is in.
//
// Perimeter Bridges are typically useful when building more complex toplogies
// with many independent perimeters that need to share some data with a common
// perimeter, but should not be able to share data among themselves.
type ServicePerimeter_PerimeterType int32

const (
	// Regular Perimeter.
	ServicePerimeter_PERIMETER_TYPE_REGULAR ServicePerimeter_PerimeterType = 0
	// Perimeter Bridge.
	ServicePerimeter_PERIMETER_TYPE_BRIDGE ServicePerimeter_PerimeterType = 1
)

// Enum value maps for ServicePerimeter_PerimeterType.
var (
	ServicePerimeter_PerimeterType_name = map[int32]string{
		0: "PERIMETER_TYPE_REGULAR",
		1: "PERIMETER_TYPE_BRIDGE",
	}
	ServicePerimeter_PerimeterType_value = map[string]int32{
		"PERIMETER_TYPE_REGULAR": 0,
		"PERIMETER_TYPE_BRIDGE":  1,
	}
)

func (x ServicePerimeter_PerimeterType) Enum() *ServicePerimeter_PerimeterType {
	p := new(ServicePerimeter_PerimeterType)
	*p = x
	return p
}

func (x ServicePerimeter_PerimeterType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServicePerimeter_PerimeterType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_identity_accesscontextmanager_v1_service_perimeter_proto_enumTypes[0].Descriptor()
}

func (ServicePerimeter_PerimeterType) Type() protoreflect.EnumType {
	return &file_google_identity_accesscontextmanager_v1_service_perimeter_proto_enumTypes[0]
}

func (x ServicePerimeter_PerimeterType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServicePerimeter_PerimeterType.Descriptor instead.
func (ServicePerimeter_PerimeterType) EnumDescriptor() ([]byte, []int) {
	return file_google_identity_accesscontextmanager_v1_service_perimeter_proto_rawDescGZIP(), []int{0, 0}
}

// `ServicePerimeter` describes a set of Google Cloud resources which can freely
// import and export data amongst themselves, but not export outside of the
// `ServicePerimeter`. If a request with a source within this `ServicePerimeter`
// has a target outside of the `ServicePerimeter`, the request will be blocked.
// Otherwise the request is allowed. There are two types of Service Perimeter -
// Regular and Bridge. Regular Service Perimeters cannot overlap, a single
// Google Cloud project can only belong to a single regular Service Perimeter.
// Service Perimeter Bridges can contain only Google Cloud projects as members,
// a single Google Cloud project may belong to multiple Service Perimeter
// Bridges.
type ServicePerimeter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Resource name for the ServicePerimeter.  The `short_name`
	// component must begin with a letter and only include alphanumeric and '_'.
	// Format: `accessPolicies/{policy_id}/servicePerimeters/{short_name}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Human readable title. Must be unique within the Policy.
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// Description of the `ServicePerimeter` and its use. Does not affect
	// behavior.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. Time the `ServicePerimeter` was created in UTC.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Time the `ServicePerimeter` was updated in UTC.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Perimeter type indicator. A single project is
	// allowed to be a member of single regular perimeter, but multiple service
	// perimeter bridges. A project cannot be a included in a perimeter bridge
	// without being included in regular perimeter. For perimeter bridges,
	// the restricted service list as well as access level lists must be
	// empty.
	PerimeterType ServicePerimeter_PerimeterType `protobuf:"varint,6,opt,name=perimeter_type,json=perimeterType,proto3,enum=google.identity.accesscontextmanager.v1.ServicePerimeter_PerimeterType" json:"perimeter_type,omitempty"`
	// Current ServicePerimeter configuration. Specifies sets of resources,
	// restricted services and access levels that determine perimeter
	// content and boundaries.
	Status *ServicePerimeterConfig `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	// Proposed (or dry run) ServicePerimeter configuration. This configuration
	// allows to specify and test ServicePerimeter configuration without enforcing
	// actual access restrictions. Only allowed to be set when the
	// "use_explicit_dry_run_spec" flag is set.
	Spec *ServicePerimeterConfig `protobuf:"bytes,8,opt,name=spec,proto3" json:"spec,omitempty"`
	// Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly
	// exists  for all Service Perimeters, and that spec is identical to the
	// status for those Service Perimeters. When this flag is set, it inhibits the
	// generation of the implicit spec, thereby allowing the user to explicitly
	// provide a configuration ("spec") to use in a dry-run version of the Service
	// Perimeter. This allows the user to test changes to the enforced config
	// ("status") without actually enforcing them. This testing is done through
	// analyzing the differences between currently enforced and suggested
	// restrictions. use_explicit_dry_run_spec must bet set to True if any of the
	// fields in the spec are set to non-default values.
	UseExplicitDryRunSpec bool `protobuf:"varint,9,opt,name=use_explicit_dry_run_spec,json=useExplicitDryRunSpec,proto3" json:"use_explicit_dry_run_spec,omitempty"`
}

func (x *ServicePerimeter) Reset() {
	*x = ServicePerimeter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_identity_accesscontextmanager_v1_service_perimeter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServicePerimeter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServicePerimeter) ProtoMessage() {}

func (x *ServicePerimeter) ProtoReflect() protoreflect.Message {
	mi := &file_google_identity_accesscontextmanager_v1_service_perimeter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServicePerimeter.ProtoReflect.Descriptor instead.
func (*ServicePerimeter) Descriptor() ([]byte, []int) {
	return file_google_identity_accesscontextmanager_v1_service_perimeter_proto_rawDescGZIP(), []int{0}
}

func (x *ServicePerimeter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServicePerimeter) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ServicePerimeter) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ServicePerimeter) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ServicePerimeter) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *ServicePerimeter) GetPerimeterType() ServicePerimeter_PerimeterType {
	if x != nil {
		return x.PerimeterType
	}
	return ServicePerimeter_PERIMETER_TYPE_REGULAR
}

func (x *ServicePerimeter) GetStatus() *ServicePerimeterConfig {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ServicePerimeter) GetSpec() *ServicePerimeterConfig {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *ServicePerimeter) GetUseExplicitDryRunSpec() bool {
	if x != nil {
		return x.UseExplicitDryRunSpec
	}
	return false
}

// `ServicePerimeterConfig` specifies a set of Google Cloud resources that
// describe specific Service Perimeter configuration.
type ServicePerimeterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of Google Cloud resources that are inside of the service perimeter.
	// Currently only projects are allowed. Format: `projects/{project_number}`
	Resources []string `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
	// A list of `AccessLevel` resource names that allow resources within the
	// `ServicePerimeter` to be accessed from the internet. `AccessLevels` listed
	// must be in the same policy as this `ServicePerimeter`. Referencing a
	// nonexistent `AccessLevel` is a syntax error. If no `AccessLevel` names are
	// listed, resources within the perimeter can only be accessed via Google
	// Cloud calls with request origins within the perimeter. Example:
	// `"accessPolicies/MY_POLICY/accessLevels/MY_LEVEL"`.
	// For Service Perimeter Bridge, must be empty.
	AccessLevels []string `protobuf:"bytes,2,rep,name=access_levels,json=accessLevels,proto3" json:"access_levels,omitempty"`
	// Google Cloud services that are subject to the Service Perimeter
	// restrictions. For example, if `storage.googleapis.com` is specified, access
	// to the storage buckets inside the perimeter must meet the perimeter's
	// access restrictions.
	RestrictedServices []string `protobuf:"bytes,4,rep,name=restricted_services,json=restrictedServices,proto3" json:"restricted_services,omitempty"`
	// Configuration for APIs allowed within Perimeter.
	VpcAccessibleServices *ServicePerimeterConfig_VpcAccessibleServices `protobuf:"bytes,10,opt,name=vpc_accessible_services,json=vpcAccessibleServices,proto3" json:"vpc_accessible_services,omitempty"`
}

func (x *ServicePerimeterConfig) Reset() {
	*x = ServicePerimeterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_identity_accesscontextmanager_v1_service_perimeter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServicePerimeterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServicePerimeterConfig) ProtoMessage() {}

func (x *ServicePerimeterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_identity_accesscontextmanager_v1_service_perimeter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServicePerimeterConfig.ProtoReflect.Descriptor instead.
func (*ServicePerimeterConfig) Descriptor() ([]byte, []int) {
	return file_google_identity_accesscontextmanager_v1_service_perimeter_proto_rawDescGZIP(), []int{1}
}

func (x *ServicePerimeterConfig) GetResources() []string {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *ServicePerimeterConfig) GetAccessLevels() []string {
	if x != nil {
		return x.AccessLevels
	}
	return nil
}

func (x *ServicePerimeterConfig) GetRestrictedServices() []string {
	if x != nil {
		return x.RestrictedServices
	}
	return nil
}

func (x *ServicePerimeterConfig) GetVpcAccessibleServices() *ServicePerimeterConfig_VpcAccessibleServices {
	if x != nil {
		return x.VpcAccessibleServices
	}
	return nil
}

// Specifies how APIs are allowed to communicate within the Service
// Perimeter.
type ServicePerimeterConfig_VpcAccessibleServices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether to restrict API calls within the Service Perimeter to the list of
	// APIs specified in 'allowed_services'.
	EnableRestriction bool `protobuf:"varint,1,opt,name=enable_restriction,json=enableRestriction,proto3" json:"enable_restriction,omitempty"`
	// The list of APIs usable within the Service Perimeter. Must be empty
	// unless 'enable_restriction' is True.
	AllowedServices []string `protobuf:"bytes,2,rep,name=allowed_services,json=allowedServices,proto3" json:"allowed_services,omitempty"`
}

func (x *ServicePerimeterConfig_VpcAccessibleServices) Reset() {
	*x = ServicePerimeterConfig_VpcAccessibleServices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_identity_accesscontextmanager_v1_service_perimeter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServicePerimeterConfig_VpcAccessibleServices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServicePerimeterConfig_VpcAccessibleServices) ProtoMessage() {}

func (x *ServicePerimeterConfig_VpcAccessibleServices) ProtoReflect() protoreflect.Message {
	mi := &file_google_identity_accesscontextmanager_v1_service_perimeter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServicePerimeterConfig_VpcAccessibleServices.ProtoReflect.Descriptor instead.
func (*ServicePerimeterConfig_VpcAccessibleServices) Descriptor() ([]byte, []int) {
	return file_google_identity_accesscontextmanager_v1_service_perimeter_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ServicePerimeterConfig_VpcAccessibleServices) GetEnableRestriction() bool {
	if x != nil {
		return x.EnableRestriction
	}
	return false
}

func (x *ServicePerimeterConfig_VpcAccessibleServices) GetAllowedServices() []string {
	if x != nil {
		return x.AllowedServices
	}
	return nil
}

var File_google_identity_accesscontextmanager_v1_service_perimeter_proto protoreflect.FileDescriptor

var file_google_identity_accesscontextmanager_v1_service_perimeter_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x04, 0x0a, 0x10, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x6e, 0x0a, 0x0e, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x47, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x57, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x53, 0x0a,
	0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x12, 0x38, 0x0a, 0x19, 0x75, 0x73, 0x65, 0x5f, 0x65, 0x78, 0x70, 0x6c, 0x69, 0x63,
	0x69, 0x74, 0x5f, 0x64, 0x72, 0x79, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x75, 0x73, 0x65, 0x45, 0x78, 0x70, 0x6c, 0x69, 0x63,
	0x69, 0x74, 0x44, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x22, 0x46, 0x0a, 0x0d,
	0x50, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x16, 0x50, 0x45, 0x52, 0x49, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x52, 0x45, 0x47, 0x55, 0x4c, 0x41, 0x52, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x45, 0x52,
	0x49, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x52, 0x49, 0x44,
	0x47, 0x45, 0x10, 0x01, 0x22, 0x8f, 0x03, 0x0a, 0x16, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x1c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x23, 0x0a,
	0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65, 0x64,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x12, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x8d, 0x01, 0x0a, 0x17, 0x76, 0x70, 0x63, 0x5f, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x55, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x56, 0x70, 0x63, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x69, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x15, 0x76, 0x70,
	0x63, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x71, 0x0a, 0x15, 0x56, 0x70, 0x63, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x69, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x12,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0xab, 0x02, 0x0a, 0x2b, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x15, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x5b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0xa2, 0x02, 0x04, 0x47,
	0x41, 0x43, 0x4d, 0xaa, 0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x27,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5c,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x2a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_identity_accesscontextmanager_v1_service_perimeter_proto_rawDescOnce sync.Once
	file_google_identity_accesscontextmanager_v1_service_perimeter_proto_rawDescData = file_google_identity_accesscontextmanager_v1_service_perimeter_proto_rawDesc
)

func file_google_identity_accesscontextmanager_v1_service_perimeter_proto_rawDescGZIP() []byte {
	file_google_identity_accesscontextmanager_v1_service_perimeter_proto_rawDescOnce.Do(func() {
		file_google_identity_accesscontextmanager_v1_service_perimeter_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_identity_accesscontextmanager_v1_service_perimeter_proto_rawDescData)
	})
	return file_google_identity_accesscontextmanager_v1_service_perimeter_proto_rawDescData
}

var file_google_identity_accesscontextmanager_v1_service_perimeter_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_identity_accesscontextmanager_v1_service_perimeter_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_identity_accesscontextmanager_v1_service_perimeter_proto_goTypes = []interface{}{
	(ServicePerimeter_PerimeterType)(0),                  // 0: google.identity.accesscontextmanager.v1.ServicePerimeter.PerimeterType
	(*ServicePerimeter)(nil),                             // 1: google.identity.accesscontextmanager.v1.ServicePerimeter
	(*ServicePerimeterConfig)(nil),                       // 2: google.identity.accesscontextmanager.v1.ServicePerimeterConfig
	(*ServicePerimeterConfig_VpcAccessibleServices)(nil), // 3: google.identity.accesscontextmanager.v1.ServicePerimeterConfig.VpcAccessibleServices
	(*timestamp.Timestamp)(nil),                          // 4: google.protobuf.Timestamp
}
var file_google_identity_accesscontextmanager_v1_service_perimeter_proto_depIdxs = []int32{
	4, // 0: google.identity.accesscontextmanager.v1.ServicePerimeter.create_time:type_name -> google.protobuf.Timestamp
	4, // 1: google.identity.accesscontextmanager.v1.ServicePerimeter.update_time:type_name -> google.protobuf.Timestamp
	0, // 2: google.identity.accesscontextmanager.v1.ServicePerimeter.perimeter_type:type_name -> google.identity.accesscontextmanager.v1.ServicePerimeter.PerimeterType
	2, // 3: google.identity.accesscontextmanager.v1.ServicePerimeter.status:type_name -> google.identity.accesscontextmanager.v1.ServicePerimeterConfig
	2, // 4: google.identity.accesscontextmanager.v1.ServicePerimeter.spec:type_name -> google.identity.accesscontextmanager.v1.ServicePerimeterConfig
	3, // 5: google.identity.accesscontextmanager.v1.ServicePerimeterConfig.vpc_accessible_services:type_name -> google.identity.accesscontextmanager.v1.ServicePerimeterConfig.VpcAccessibleServices
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_identity_accesscontextmanager_v1_service_perimeter_proto_init() }
func file_google_identity_accesscontextmanager_v1_service_perimeter_proto_init() {
	if File_google_identity_accesscontextmanager_v1_service_perimeter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_identity_accesscontextmanager_v1_service_perimeter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServicePerimeter); i {
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
		file_google_identity_accesscontextmanager_v1_service_perimeter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServicePerimeterConfig); i {
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
		file_google_identity_accesscontextmanager_v1_service_perimeter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServicePerimeterConfig_VpcAccessibleServices); i {
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
			RawDescriptor: file_google_identity_accesscontextmanager_v1_service_perimeter_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_identity_accesscontextmanager_v1_service_perimeter_proto_goTypes,
		DependencyIndexes: file_google_identity_accesscontextmanager_v1_service_perimeter_proto_depIdxs,
		EnumInfos:         file_google_identity_accesscontextmanager_v1_service_perimeter_proto_enumTypes,
		MessageInfos:      file_google_identity_accesscontextmanager_v1_service_perimeter_proto_msgTypes,
	}.Build()
	File_google_identity_accesscontextmanager_v1_service_perimeter_proto = out.File
	file_google_identity_accesscontextmanager_v1_service_perimeter_proto_rawDesc = nil
	file_google_identity_accesscontextmanager_v1_service_perimeter_proto_goTypes = nil
	file_google_identity_accesscontextmanager_v1_service_perimeter_proto_depIdxs = nil
}
