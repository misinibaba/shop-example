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
// source: google/devtools/resultstore/v2/common.proto

package resultstore

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// These correspond to the prefix of the rule name. Eg cc_test has language CC.
type Language int32

const (
	// Language unspecified or not listed here.
	Language_LANGUAGE_UNSPECIFIED Language = 0
	// Not related to any particular language
	Language_NONE Language = 1
	// Android
	Language_ANDROID Language = 2
	// ActionScript (Flash)
	Language_AS Language = 3
	// C++ or C
	Language_CC Language = 4
	// Cascading-Style-Sheets
	Language_CSS Language = 5
	// Dart
	Language_DART Language = 6
	// Go
	Language_GO Language = 7
	// Google-Web-Toolkit
	Language_GWT Language = 8
	// Haskell
	Language_HASKELL Language = 9
	// Java
	Language_JAVA Language = 10
	// Javascript
	Language_JS Language = 11
	// Lisp
	Language_LISP Language = 12
	// Objective-C
	Language_OBJC Language = 13
	// Python
	Language_PY Language = 14
	// Shell (Typically Bash)
	Language_SH Language = 15
	// Swift
	Language_SWIFT Language = 16
	// Typescript
	Language_TS Language = 18
	// Webtesting
	Language_WEB Language = 19
	// Scala
	Language_SCALA Language = 20
	// Protocol Buffer
	Language_PROTO Language = 21
	// Extensible Markup Language
	Language_XML Language = 22
)

// Enum value maps for Language.
var (
	Language_name = map[int32]string{
		0:  "LANGUAGE_UNSPECIFIED",
		1:  "NONE",
		2:  "ANDROID",
		3:  "AS",
		4:  "CC",
		5:  "CSS",
		6:  "DART",
		7:  "GO",
		8:  "GWT",
		9:  "HASKELL",
		10: "JAVA",
		11: "JS",
		12: "LISP",
		13: "OBJC",
		14: "PY",
		15: "SH",
		16: "SWIFT",
		18: "TS",
		19: "WEB",
		20: "SCALA",
		21: "PROTO",
		22: "XML",
	}
	Language_value = map[string]int32{
		"LANGUAGE_UNSPECIFIED": 0,
		"NONE":                 1,
		"ANDROID":              2,
		"AS":                   3,
		"CC":                   4,
		"CSS":                  5,
		"DART":                 6,
		"GO":                   7,
		"GWT":                  8,
		"HASKELL":              9,
		"JAVA":                 10,
		"JS":                   11,
		"LISP":                 12,
		"OBJC":                 13,
		"PY":                   14,
		"SH":                   15,
		"SWIFT":                16,
		"TS":                   18,
		"WEB":                  19,
		"SCALA":                20,
		"PROTO":                21,
		"XML":                  22,
	}
)

func (x Language) Enum() *Language {
	p := new(Language)
	*p = x
	return p
}

func (x Language) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Language) Descriptor() protoreflect.EnumDescriptor {
	return file_google_devtools_resultstore_v2_common_proto_enumTypes[0].Descriptor()
}

func (Language) Type() protoreflect.EnumType {
	return &file_google_devtools_resultstore_v2_common_proto_enumTypes[0]
}

func (x Language) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Language.Descriptor instead.
func (Language) EnumDescriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_common_proto_rawDescGZIP(), []int{0}
}

// Status of a resource.
type Status int32

const (
	// The implicit default enum value. Should never be set.
	Status_STATUS_UNSPECIFIED Status = 0
	// Displays as "Building". Means the target is compiling, linking, etc.
	Status_BUILDING Status = 1
	// Displays as "Built". Means the target was built successfully.
	// If testing was requested, it should never reach this status: it should go
	// straight from BUILDING to TESTING.
	Status_BUILT Status = 2
	// Displays as "Broken". Means build failure such as compile error.
	Status_FAILED_TO_BUILD Status = 3
	// Displays as "Testing". Means the test is running.
	Status_TESTING Status = 4
	// Displays as "Passed". Means the test was run and passed.
	Status_PASSED Status = 5
	// Displays as "Failed". Means the test was run and failed.
	Status_FAILED Status = 6
	// Displays as "Timed out". Means the test didn't finish in time.
	Status_TIMED_OUT Status = 7
	// Displays as "Cancelled". Means the build or test was cancelled.
	// E.g. User hit control-C.
	Status_CANCELLED Status = 8
	// Displays as "Tool Failed". Means the build or test had internal tool
	// failure.
	Status_TOOL_FAILED Status = 9
	// Displays as "Incomplete". Means the build or test did not complete.  This
	// might happen when a build breakage or test failure causes the tool to stop
	// trying to build anything more or run any more tests, with the default
	// bazel --nokeep_going option or the --notest_keep_going option.
	Status_INCOMPLETE Status = 10
	// Displays as "Flaky". Means the aggregate status contains some runs that
	// were successful, and some that were not.
	Status_FLAKY Status = 11
	// Displays as "Unknown". Means the tool uploading to the server died
	// mid-upload or does not know the state.
	Status_UNKNOWN Status = 12
	// Displays as "Skipped". Means building and testing were skipped.
	// (E.g. Restricted to a different configuration.)
	Status_SKIPPED Status = 13
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0:  "STATUS_UNSPECIFIED",
		1:  "BUILDING",
		2:  "BUILT",
		3:  "FAILED_TO_BUILD",
		4:  "TESTING",
		5:  "PASSED",
		6:  "FAILED",
		7:  "TIMED_OUT",
		8:  "CANCELLED",
		9:  "TOOL_FAILED",
		10: "INCOMPLETE",
		11: "FLAKY",
		12: "UNKNOWN",
		13: "SKIPPED",
	}
	Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"BUILDING":           1,
		"BUILT":              2,
		"FAILED_TO_BUILD":    3,
		"TESTING":            4,
		"PASSED":             5,
		"FAILED":             6,
		"TIMED_OUT":          7,
		"CANCELLED":          8,
		"TOOL_FAILED":        9,
		"INCOMPLETE":         10,
		"FLAKY":              11,
		"UNKNOWN":            12,
		"SKIPPED":            13,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_google_devtools_resultstore_v2_common_proto_enumTypes[1].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_google_devtools_resultstore_v2_common_proto_enumTypes[1]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_common_proto_rawDescGZIP(), []int{1}
}

// Indicates the upload status of the invocation, whether it is
// post-processing, or immutable, etc.
type UploadStatus int32

const (
	// The implicit default enum value. Should never be set.
	UploadStatus_UPLOAD_STATUS_UNSPECIFIED UploadStatus = 0
	// The invocation is still uploading to the ResultStore.
	UploadStatus_UPLOADING UploadStatus = 1
	// The invocation upload is complete. The ResultStore is still post-processing
	// the invocation.
	UploadStatus_POST_PROCESSING UploadStatus = 2
	// All post-processing is complete, and the invocation is now immutable.
	UploadStatus_IMMUTABLE UploadStatus = 3
)

// Enum value maps for UploadStatus.
var (
	UploadStatus_name = map[int32]string{
		0: "UPLOAD_STATUS_UNSPECIFIED",
		1: "UPLOADING",
		2: "POST_PROCESSING",
		3: "IMMUTABLE",
	}
	UploadStatus_value = map[string]int32{
		"UPLOAD_STATUS_UNSPECIFIED": 0,
		"UPLOADING":                 1,
		"POST_PROCESSING":           2,
		"IMMUTABLE":                 3,
	}
)

func (x UploadStatus) Enum() *UploadStatus {
	p := new(UploadStatus)
	*p = x
	return p
}

func (x UploadStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UploadStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_google_devtools_resultstore_v2_common_proto_enumTypes[2].Descriptor()
}

func (UploadStatus) Type() protoreflect.EnumType {
	return &file_google_devtools_resultstore_v2_common_proto_enumTypes[2]
}

func (x UploadStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UploadStatus.Descriptor instead.
func (UploadStatus) EnumDescriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_common_proto_rawDescGZIP(), []int{2}
}

// Describes the status of a resource in both enum and string form.
// Only use description when conveying additional info not captured in the enum
// name.
type StatusAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Enum representation of the status.
	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=google.devtools.resultstore.v2.Status" json:"status,omitempty"`
	// A longer description about the status.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *StatusAttributes) Reset() {
	*x = StatusAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusAttributes) ProtoMessage() {}

func (x *StatusAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusAttributes.ProtoReflect.Descriptor instead.
func (*StatusAttributes) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_common_proto_rawDescGZIP(), []int{0}
}

func (x *StatusAttributes) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_STATUS_UNSPECIFIED
}

func (x *StatusAttributes) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// A generic key-value property definition.
type Property struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The key.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The value.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Property) Reset() {
	*x = Property{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Property) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Property) ProtoMessage() {}

func (x *Property) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Property.ProtoReflect.Descriptor instead.
func (*Property) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_common_proto_rawDescGZIP(), []int{1}
}

func (x *Property) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Property) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// The timing of a particular Invocation, Action, etc. The start_time is
// specified, stop time can be calculated by adding duration to start_time.
type Timing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The time the resource started running. This is in UTC Epoch time.
	StartTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// The duration for which the resource ran.
	Duration *duration.Duration `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *Timing) Reset() {
	*x = Timing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timing) ProtoMessage() {}

func (x *Timing) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timing.ProtoReflect.Descriptor instead.
func (*Timing) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_common_proto_rawDescGZIP(), []int{2}
}

func (x *Timing) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Timing) GetDuration() *duration.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

// Represents a dependency of a resource on another resource. This can be used
// to define a graph or a workflow paradigm through resources.
type Dependency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource depended upon. It may be a Target, ConfiguredTarget, or
	// Action.
	//
	// Types that are assignable to Resource:
	//	*Dependency_Target
	//	*Dependency_ConfiguredTarget
	//	*Dependency_Action
	Resource isDependency_Resource `protobuf_oneof:"resource"`
	// A label describing this dependency.
	// The label "Root Cause" is handled specially. It is used to point to the
	// exact resource that caused a resource to fail.
	Label string `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *Dependency) Reset() {
	*x = Dependency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dependency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dependency) ProtoMessage() {}

func (x *Dependency) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dependency.ProtoReflect.Descriptor instead.
func (*Dependency) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_common_proto_rawDescGZIP(), []int{3}
}

func (m *Dependency) GetResource() isDependency_Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (x *Dependency) GetTarget() string {
	if x, ok := x.GetResource().(*Dependency_Target); ok {
		return x.Target
	}
	return ""
}

func (x *Dependency) GetConfiguredTarget() string {
	if x, ok := x.GetResource().(*Dependency_ConfiguredTarget); ok {
		return x.ConfiguredTarget
	}
	return ""
}

func (x *Dependency) GetAction() string {
	if x, ok := x.GetResource().(*Dependency_Action); ok {
		return x.Action
	}
	return ""
}

func (x *Dependency) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type isDependency_Resource interface {
	isDependency_Resource()
}

type Dependency_Target struct {
	// The name of a target.  Its format must be:
	// invocations/${INVOCATION_ID}/targets/${url_encode(TARGET_ID)}
	// This must point to an target under the same invocation.
	Target string `protobuf:"bytes,1,opt,name=target,proto3,oneof"`
}

type Dependency_ConfiguredTarget struct {
	// The name of a configured target.  Its format must be:
	// invocations/${INVOCATION_ID}/targets/${url_encode(TARGET_ID)}/configuredTargets/${url_encode(CONFIG_ID)}
	// This must point to an configured target under the same invocation.
	ConfiguredTarget string `protobuf:"bytes,2,opt,name=configured_target,json=configuredTarget,proto3,oneof"`
}

type Dependency_Action struct {
	// The name of an action.  Its format must be:
	// invocations/${INVOCATION_ID}/targets/${url_encode(TARGET_ID)}/configuredTargets/${url_encode(CONFIG_ID)}/actions/${url_encode(ACTION_ID)}
	// This must point to an action under the same invocation.
	Action string `protobuf:"bytes,3,opt,name=action,proto3,oneof"`
}

func (*Dependency_Target) isDependency_Resource() {}

func (*Dependency_ConfiguredTarget) isDependency_Resource() {}

func (*Dependency_Action) isDependency_Resource() {}

var File_google_devtools_resultstore_v2_common_proto protoreflect.FileDescriptor

var file_google_devtools_resultstore_v2_common_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74,
	0x0a, 0x10, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x12, 0x3e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74,
	0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x32, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x7a, 0x0a, 0x06, 0x54, 0x69, 0x6d, 0x69,
	0x6e, 0x67, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a,
	0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x91, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65,
	0x6e, 0x63, 0x79, 0x12, 0x18, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x2d, 0x0a,
	0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x42, 0x0a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2a, 0xed, 0x01, 0x0a, 0x08, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x4e, 0x44,
	0x52, 0x4f, 0x49, 0x44, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x41, 0x53, 0x10, 0x03, 0x12, 0x06,
	0x0a, 0x02, 0x43, 0x43, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x53, 0x53, 0x10, 0x05, 0x12,
	0x08, 0x0a, 0x04, 0x44, 0x41, 0x52, 0x54, 0x10, 0x06, 0x12, 0x06, 0x0a, 0x02, 0x47, 0x4f, 0x10,
	0x07, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x57, 0x54, 0x10, 0x08, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x41,
	0x53, 0x4b, 0x45, 0x4c, 0x4c, 0x10, 0x09, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x41, 0x56, 0x41, 0x10,
	0x0a, 0x12, 0x06, 0x0a, 0x02, 0x4a, 0x53, 0x10, 0x0b, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x49, 0x53,
	0x50, 0x10, 0x0c, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x42, 0x4a, 0x43, 0x10, 0x0d, 0x12, 0x06, 0x0a,
	0x02, 0x50, 0x59, 0x10, 0x0e, 0x12, 0x06, 0x0a, 0x02, 0x53, 0x48, 0x10, 0x0f, 0x12, 0x09, 0x0a,
	0x05, 0x53, 0x57, 0x49, 0x46, 0x54, 0x10, 0x10, 0x12, 0x06, 0x0a, 0x02, 0x54, 0x53, 0x10, 0x12,
	0x12, 0x07, 0x0a, 0x03, 0x57, 0x45, 0x42, 0x10, 0x13, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x43, 0x41,
	0x4c, 0x41, 0x10, 0x14, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x10, 0x15, 0x12,
	0x07, 0x0a, 0x03, 0x58, 0x4d, 0x4c, 0x10, 0x16, 0x2a, 0xd7, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x42,
	0x55, 0x49, 0x4c, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x55, 0x49,
	0x4c, 0x54, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x54,
	0x4f, 0x5f, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x45, 0x53,
	0x54, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x53, 0x53, 0x45, 0x44,
	0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x06, 0x12, 0x0d,
	0x0a, 0x09, 0x54, 0x49, 0x4d, 0x45, 0x44, 0x5f, 0x4f, 0x55, 0x54, 0x10, 0x07, 0x12, 0x0d, 0x0a,
	0x09, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x08, 0x12, 0x0f, 0x0a, 0x0b,
	0x54, 0x4f, 0x4f, 0x4c, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x09, 0x12, 0x0e, 0x0a,
	0x0a, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x0a, 0x12, 0x09, 0x0a,
	0x05, 0x46, 0x4c, 0x41, 0x4b, 0x59, 0x10, 0x0b, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x0c, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x4b, 0x49, 0x50, 0x50, 0x45, 0x44,
	0x10, 0x0d, 0x2a, 0x60, 0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1d, 0x0a, 0x19, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01,
	0x12, 0x13, 0x0a, 0x0f, 0x50, 0x4f, 0x53, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53,
	0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4d, 0x4d, 0x55, 0x54, 0x41, 0x42,
	0x4c, 0x45, 0x10, 0x03, 0x42, 0x71, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x3b, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_devtools_resultstore_v2_common_proto_rawDescOnce sync.Once
	file_google_devtools_resultstore_v2_common_proto_rawDescData = file_google_devtools_resultstore_v2_common_proto_rawDesc
)

func file_google_devtools_resultstore_v2_common_proto_rawDescGZIP() []byte {
	file_google_devtools_resultstore_v2_common_proto_rawDescOnce.Do(func() {
		file_google_devtools_resultstore_v2_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_devtools_resultstore_v2_common_proto_rawDescData)
	})
	return file_google_devtools_resultstore_v2_common_proto_rawDescData
}

var file_google_devtools_resultstore_v2_common_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_google_devtools_resultstore_v2_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_devtools_resultstore_v2_common_proto_goTypes = []interface{}{
	(Language)(0),               // 0: google.devtools.resultstore.v2.Language
	(Status)(0),                 // 1: google.devtools.resultstore.v2.Status
	(UploadStatus)(0),           // 2: google.devtools.resultstore.v2.UploadStatus
	(*StatusAttributes)(nil),    // 3: google.devtools.resultstore.v2.StatusAttributes
	(*Property)(nil),            // 4: google.devtools.resultstore.v2.Property
	(*Timing)(nil),              // 5: google.devtools.resultstore.v2.Timing
	(*Dependency)(nil),          // 6: google.devtools.resultstore.v2.Dependency
	(*timestamp.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*duration.Duration)(nil),   // 8: google.protobuf.Duration
}
var file_google_devtools_resultstore_v2_common_proto_depIdxs = []int32{
	1, // 0: google.devtools.resultstore.v2.StatusAttributes.status:type_name -> google.devtools.resultstore.v2.Status
	7, // 1: google.devtools.resultstore.v2.Timing.start_time:type_name -> google.protobuf.Timestamp
	8, // 2: google.devtools.resultstore.v2.Timing.duration:type_name -> google.protobuf.Duration
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_devtools_resultstore_v2_common_proto_init() }
func file_google_devtools_resultstore_v2_common_proto_init() {
	if File_google_devtools_resultstore_v2_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_devtools_resultstore_v2_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusAttributes); i {
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
		file_google_devtools_resultstore_v2_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Property); i {
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
		file_google_devtools_resultstore_v2_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timing); i {
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
		file_google_devtools_resultstore_v2_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dependency); i {
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
	file_google_devtools_resultstore_v2_common_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Dependency_Target)(nil),
		(*Dependency_ConfiguredTarget)(nil),
		(*Dependency_Action)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_devtools_resultstore_v2_common_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_devtools_resultstore_v2_common_proto_goTypes,
		DependencyIndexes: file_google_devtools_resultstore_v2_common_proto_depIdxs,
		EnumInfos:         file_google_devtools_resultstore_v2_common_proto_enumTypes,
		MessageInfos:      file_google_devtools_resultstore_v2_common_proto_msgTypes,
	}.Build()
	File_google_devtools_resultstore_v2_common_proto = out.File
	file_google_devtools_resultstore_v2_common_proto_rawDesc = nil
	file_google_devtools_resultstore_v2_common_proto_goTypes = nil
	file_google_devtools_resultstore_v2_common_proto_depIdxs = nil
}
