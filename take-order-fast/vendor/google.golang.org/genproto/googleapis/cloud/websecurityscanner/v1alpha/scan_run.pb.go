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
// source: google/cloud/websecurityscanner/v1alpha/scan_run.proto

package websecurityscanner

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

// Types of ScanRun execution state.
type ScanRun_ExecutionState int32

const (
	// Represents an invalid state caused by internal server error. This value
	// should never be returned.
	ScanRun_EXECUTION_STATE_UNSPECIFIED ScanRun_ExecutionState = 0
	// The scan is waiting in the queue.
	ScanRun_QUEUED ScanRun_ExecutionState = 1
	// The scan is in progress.
	ScanRun_SCANNING ScanRun_ExecutionState = 2
	// The scan is either finished or stopped by user.
	ScanRun_FINISHED ScanRun_ExecutionState = 3
)

// Enum value maps for ScanRun_ExecutionState.
var (
	ScanRun_ExecutionState_name = map[int32]string{
		0: "EXECUTION_STATE_UNSPECIFIED",
		1: "QUEUED",
		2: "SCANNING",
		3: "FINISHED",
	}
	ScanRun_ExecutionState_value = map[string]int32{
		"EXECUTION_STATE_UNSPECIFIED": 0,
		"QUEUED":                      1,
		"SCANNING":                    2,
		"FINISHED":                    3,
	}
)

func (x ScanRun_ExecutionState) Enum() *ScanRun_ExecutionState {
	p := new(ScanRun_ExecutionState)
	*p = x
	return p
}

func (x ScanRun_ExecutionState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScanRun_ExecutionState) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_enumTypes[0].Descriptor()
}

func (ScanRun_ExecutionState) Type() protoreflect.EnumType {
	return &file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_enumTypes[0]
}

func (x ScanRun_ExecutionState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ScanRun_ExecutionState.Descriptor instead.
func (ScanRun_ExecutionState) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_rawDescGZIP(), []int{0, 0}
}

// Types of ScanRun result state.
type ScanRun_ResultState int32

const (
	// Default value. This value is returned when the ScanRun is not yet
	// finished.
	ScanRun_RESULT_STATE_UNSPECIFIED ScanRun_ResultState = 0
	// The scan finished without errors.
	ScanRun_SUCCESS ScanRun_ResultState = 1
	// The scan finished with errors.
	ScanRun_ERROR ScanRun_ResultState = 2
	// The scan was terminated by user.
	ScanRun_KILLED ScanRun_ResultState = 3
)

// Enum value maps for ScanRun_ResultState.
var (
	ScanRun_ResultState_name = map[int32]string{
		0: "RESULT_STATE_UNSPECIFIED",
		1: "SUCCESS",
		2: "ERROR",
		3: "KILLED",
	}
	ScanRun_ResultState_value = map[string]int32{
		"RESULT_STATE_UNSPECIFIED": 0,
		"SUCCESS":                  1,
		"ERROR":                    2,
		"KILLED":                   3,
	}
)

func (x ScanRun_ResultState) Enum() *ScanRun_ResultState {
	p := new(ScanRun_ResultState)
	*p = x
	return p
}

func (x ScanRun_ResultState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScanRun_ResultState) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_enumTypes[1].Descriptor()
}

func (ScanRun_ResultState) Type() protoreflect.EnumType {
	return &file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_enumTypes[1]
}

func (x ScanRun_ResultState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ScanRun_ResultState.Descriptor instead.
func (ScanRun_ResultState) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_rawDescGZIP(), []int{0, 1}
}

// A ScanRun is a output-only resource representing an actual run of the scan.
type ScanRun struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the ScanRun. The name follows the format of
	// 'projects/{projectId}/scanConfigs/{scanConfigId}/scanRuns/{scanRunId}'.
	// The ScanRun IDs are generated by the system.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The execution state of the ScanRun.
	ExecutionState ScanRun_ExecutionState `protobuf:"varint,2,opt,name=execution_state,json=executionState,proto3,enum=google.cloud.websecurityscanner.v1alpha.ScanRun_ExecutionState" json:"execution_state,omitempty"`
	// The result state of the ScanRun. This field is only available after the
	// execution state reaches "FINISHED".
	ResultState ScanRun_ResultState `protobuf:"varint,3,opt,name=result_state,json=resultState,proto3,enum=google.cloud.websecurityscanner.v1alpha.ScanRun_ResultState" json:"result_state,omitempty"`
	// The time at which the ScanRun started.
	StartTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// The time at which the ScanRun reached termination state - that the ScanRun
	// is either finished or stopped by user.
	EndTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// The number of URLs crawled during this ScanRun. If the scan is in progress,
	// the value represents the number of URLs crawled up to now.
	UrlsCrawledCount int64 `protobuf:"varint,6,opt,name=urls_crawled_count,json=urlsCrawledCount,proto3" json:"urls_crawled_count,omitempty"`
	// The number of URLs tested during this ScanRun. If the scan is in progress,
	// the value represents the number of URLs tested up to now. The number of
	// URLs tested is usually larger than the number URLS crawled because
	// typically a crawled URL is tested with multiple test payloads.
	UrlsTestedCount int64 `protobuf:"varint,7,opt,name=urls_tested_count,json=urlsTestedCount,proto3" json:"urls_tested_count,omitempty"`
	// Whether the scan run has found any vulnerabilities.
	HasVulnerabilities bool `protobuf:"varint,8,opt,name=has_vulnerabilities,json=hasVulnerabilities,proto3" json:"has_vulnerabilities,omitempty"`
	// The percentage of total completion ranging from 0 to 100.
	// If the scan is in queue, the value is 0.
	// If the scan is running, the value ranges from 0 to 100.
	// If the scan is finished, the value is 100.
	ProgressPercent int32 `protobuf:"varint,9,opt,name=progress_percent,json=progressPercent,proto3" json:"progress_percent,omitempty"`
}

func (x *ScanRun) Reset() {
	*x = ScanRun{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanRun) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanRun) ProtoMessage() {}

func (x *ScanRun) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanRun.ProtoReflect.Descriptor instead.
func (*ScanRun) Descriptor() ([]byte, []int) {
	return file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_rawDescGZIP(), []int{0}
}

func (x *ScanRun) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ScanRun) GetExecutionState() ScanRun_ExecutionState {
	if x != nil {
		return x.ExecutionState
	}
	return ScanRun_EXECUTION_STATE_UNSPECIFIED
}

func (x *ScanRun) GetResultState() ScanRun_ResultState {
	if x != nil {
		return x.ResultState
	}
	return ScanRun_RESULT_STATE_UNSPECIFIED
}

func (x *ScanRun) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *ScanRun) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *ScanRun) GetUrlsCrawledCount() int64 {
	if x != nil {
		return x.UrlsCrawledCount
	}
	return 0
}

func (x *ScanRun) GetUrlsTestedCount() int64 {
	if x != nil {
		return x.UrlsTestedCount
	}
	return 0
}

func (x *ScanRun) GetHasVulnerabilities() bool {
	if x != nil {
		return x.HasVulnerabilities
	}
	return false
}

func (x *ScanRun) GetProgressPercent() int32 {
	if x != nil {
		return x.ProgressPercent
	}
	return 0
}

var File_google_cloud_websecurityscanner_v1alpha_scan_run_proto protoreflect.FileDescriptor

var file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x77,
	0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x72,
	0x75, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x06,
	0x0a, 0x07, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x75, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x68, 0x0a,
	0x0f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x75, 0x6e, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x5f, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x77, 0x65, 0x62,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x75, 0x6e, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x75, 0x72,
	0x6c, 0x73, 0x5f, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x75, 0x72, 0x6c, 0x73, 0x43, 0x72, 0x61, 0x77,
	0x6c, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x75, 0x72, 0x6c, 0x73,
	0x5f, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0f, 0x75, 0x72, 0x6c, 0x73, 0x54, 0x65, 0x73, 0x74, 0x65, 0x64, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x68, 0x61, 0x73, 0x5f, 0x76, 0x75, 0x6c, 0x6e,
	0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x12, 0x68, 0x61, 0x73, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x22, 0x59, 0x0a, 0x0e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x51, 0x55, 0x45, 0x55, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x0c, 0x0a, 0x08, 0x53, 0x43, 0x41, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0c, 0x0a,
	0x08, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x03, 0x22, 0x4f, 0x0a, 0x0b, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x52, 0x45,
	0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43,
	0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02,
	0x12, 0x0a, 0x0a, 0x06, 0x4b, 0x49, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x3a, 0x70, 0xea, 0x41,
	0x6d, 0x0a, 0x29, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x75, 0x6e, 0x12, 0x40, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d,
	0x2f, 0x73, 0x63, 0x61, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f, 0x7b, 0x73, 0x63,
	0x61, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x7d, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x52,
	0x75, 0x6e, 0x73, 0x2f, 0x7b, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x72, 0x75, 0x6e, 0x7d, 0x42, 0x98,
	0x01, 0x0a, 0x2b, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73,
	0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0c,
	0x53, 0x63, 0x61, 0x6e, 0x52, 0x75, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x59,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x77, 0x65, 0x62, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_rawDescOnce sync.Once
	file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_rawDescData = file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_rawDesc
)

func file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_rawDescGZIP() []byte {
	file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_rawDescOnce.Do(func() {
		file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_rawDescData)
	})
	return file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_rawDescData
}

var file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_goTypes = []interface{}{
	(ScanRun_ExecutionState)(0), // 0: google.cloud.websecurityscanner.v1alpha.ScanRun.ExecutionState
	(ScanRun_ResultState)(0),    // 1: google.cloud.websecurityscanner.v1alpha.ScanRun.ResultState
	(*ScanRun)(nil),             // 2: google.cloud.websecurityscanner.v1alpha.ScanRun
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_depIdxs = []int32{
	0, // 0: google.cloud.websecurityscanner.v1alpha.ScanRun.execution_state:type_name -> google.cloud.websecurityscanner.v1alpha.ScanRun.ExecutionState
	1, // 1: google.cloud.websecurityscanner.v1alpha.ScanRun.result_state:type_name -> google.cloud.websecurityscanner.v1alpha.ScanRun.ResultState
	3, // 2: google.cloud.websecurityscanner.v1alpha.ScanRun.start_time:type_name -> google.protobuf.Timestamp
	3, // 3: google.cloud.websecurityscanner.v1alpha.ScanRun.end_time:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_init() }
func file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_init() {
	if File_google_cloud_websecurityscanner_v1alpha_scan_run_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanRun); i {
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
			RawDescriptor: file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_goTypes,
		DependencyIndexes: file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_depIdxs,
		EnumInfos:         file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_enumTypes,
		MessageInfos:      file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_msgTypes,
	}.Build()
	File_google_cloud_websecurityscanner_v1alpha_scan_run_proto = out.File
	file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_rawDesc = nil
	file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_goTypes = nil
	file_google_cloud_websecurityscanner_v1alpha_scan_run_proto_depIdxs = nil
}
