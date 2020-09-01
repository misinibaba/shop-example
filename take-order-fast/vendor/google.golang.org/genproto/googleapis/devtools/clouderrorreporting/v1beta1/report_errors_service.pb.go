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
// source: google/devtools/clouderrorreporting/v1beta1/report_errors_service.proto

package clouderrorreporting

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// A request for reporting an individual error event.
type ReportErrorEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the Google Cloud Platform project. Written
	// as `projects/` plus the
	// [Google Cloud Platform project
	// ID](https://support.google.com/cloud/answer/6158840). Example:
	// `projects/my-project-123`.
	ProjectName string `protobuf:"bytes,1,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	// Required. The error event to be reported.
	Event *ReportedErrorEvent `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *ReportErrorEventRequest) Reset() {
	*x = ReportErrorEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportErrorEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportErrorEventRequest) ProtoMessage() {}

func (x *ReportErrorEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportErrorEventRequest.ProtoReflect.Descriptor instead.
func (*ReportErrorEventRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_rawDescGZIP(), []int{0}
}

func (x *ReportErrorEventRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *ReportErrorEventRequest) GetEvent() *ReportedErrorEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

// Response for reporting an individual error event.
// Data may be added to this message in the future.
type ReportErrorEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReportErrorEventResponse) Reset() {
	*x = ReportErrorEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportErrorEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportErrorEventResponse) ProtoMessage() {}

func (x *ReportErrorEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportErrorEventResponse.ProtoReflect.Descriptor instead.
func (*ReportErrorEventResponse) Descriptor() ([]byte, []int) {
	return file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_rawDescGZIP(), []int{1}
}

// An error event which is reported to the Error Reporting system.
type ReportedErrorEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. Time when the event occurred.
	// If not provided, the time when the event was received by the
	// Error Reporting system will be used.
	EventTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	// Required. The service context in which this error has occurred.
	ServiceContext *ServiceContext `protobuf:"bytes,2,opt,name=service_context,json=serviceContext,proto3" json:"service_context,omitempty"`
	// Required. The error message.
	// If no `context.reportLocation` is provided, the message must contain a
	// header (typically consisting of the exception type name and an error
	// message) and an exception stack trace in one of the supported programming
	// languages and formats.
	// Supported languages are Java, Python, JavaScript, Ruby, C#, PHP, and Go.
	// Supported stack trace formats are:
	//
	// * **Java**: Must be the return value of
	// [`Throwable.printStackTrace()`](https://docs.oracle.com/javase/7/docs/api/java/lang/Throwable.html#printStackTrace%28%29).
	// * **Python**: Must be the return value of
	// [`traceback.format_exc()`](https://docs.python.org/2/library/traceback.html#traceback.format_exc).
	// * **JavaScript**: Must be the value of
	// [`error.stack`](https://github.com/v8/v8/wiki/Stack-Trace-API) as returned
	// by V8.
	// * **Ruby**: Must contain frames returned by
	// [`Exception.backtrace`](https://ruby-doc.org/core-2.2.0/Exception.html#method-i-backtrace).
	// * **C#**: Must be the return value of
	// [`Exception.ToString()`](https://msdn.microsoft.com/en-us/library/system.exception.tostring.aspx).
	// * **PHP**: Must start with `PHP (Notice|Parse error|Fatal error|Warning)`
	// and contain the result of
	// [`(string)$exception`](http://php.net/manual/en/exception.tostring.php).
	// * **Go**: Must be the return value of
	// [`runtime.Stack()`](https://golang.org/pkg/runtime/debug/#Stack).
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// Optional. A description of the context in which the error occurred.
	Context *ErrorContext `protobuf:"bytes,4,opt,name=context,proto3" json:"context,omitempty"`
}

func (x *ReportedErrorEvent) Reset() {
	*x = ReportedErrorEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportedErrorEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportedErrorEvent) ProtoMessage() {}

func (x *ReportedErrorEvent) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportedErrorEvent.ProtoReflect.Descriptor instead.
func (*ReportedErrorEvent) Descriptor() ([]byte, []int) {
	return file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_rawDescGZIP(), []int{2}
}

func (x *ReportedErrorEvent) GetEventTime() *timestamp.Timestamp {
	if x != nil {
		return x.EventTime
	}
	return nil
}

func (x *ReportedErrorEvent) GetServiceContext() *ServiceContext {
	if x != nil {
		return x.ServiceContext
	}
	return nil
}

func (x *ReportedErrorEvent) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ReportedErrorEvent) GetContext() *ErrorContext {
	if x != nil {
		return x.Context
	}
	return nil
}

var File_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto protoreflect.FileDescriptor

var file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_rawDesc = []byte{
	0x0a, 0x47, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x17, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x56, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2d, 0x0a, 0x2b,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x0b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x5a, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x22, 0x1a, 0x0a, 0x18, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0xb8, 0x02, 0x0a, 0x12, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x09, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x69, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f,
	0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x58, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74,
	0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x32, 0xe5, 0x02, 0x0a, 0x13,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0xf5, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x44, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x45,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x54, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x39, 0x22, 0x30, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a,
	0x7d, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x3a,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0xda, 0x41, 0x12, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x2c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x56, 0xca, 0x41, 0x22,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x42, 0xfc, 0x01, 0x0a, 0x2f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x18, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x5e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f,
	0x6f, 0x6c, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0xf8, 0x01, 0x01, 0xaa, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x23, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_rawDescOnce sync.Once
	file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_rawDescData = file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_rawDesc
)

func file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_rawDescGZIP() []byte {
	file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_rawDescOnce.Do(func() {
		file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_rawDescData)
	})
	return file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_rawDescData
}

var file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_goTypes = []interface{}{
	(*ReportErrorEventRequest)(nil),  // 0: google.devtools.clouderrorreporting.v1beta1.ReportErrorEventRequest
	(*ReportErrorEventResponse)(nil), // 1: google.devtools.clouderrorreporting.v1beta1.ReportErrorEventResponse
	(*ReportedErrorEvent)(nil),       // 2: google.devtools.clouderrorreporting.v1beta1.ReportedErrorEvent
	(*timestamp.Timestamp)(nil),      // 3: google.protobuf.Timestamp
	(*ServiceContext)(nil),           // 4: google.devtools.clouderrorreporting.v1beta1.ServiceContext
	(*ErrorContext)(nil),             // 5: google.devtools.clouderrorreporting.v1beta1.ErrorContext
}
var file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_depIdxs = []int32{
	2, // 0: google.devtools.clouderrorreporting.v1beta1.ReportErrorEventRequest.event:type_name -> google.devtools.clouderrorreporting.v1beta1.ReportedErrorEvent
	3, // 1: google.devtools.clouderrorreporting.v1beta1.ReportedErrorEvent.event_time:type_name -> google.protobuf.Timestamp
	4, // 2: google.devtools.clouderrorreporting.v1beta1.ReportedErrorEvent.service_context:type_name -> google.devtools.clouderrorreporting.v1beta1.ServiceContext
	5, // 3: google.devtools.clouderrorreporting.v1beta1.ReportedErrorEvent.context:type_name -> google.devtools.clouderrorreporting.v1beta1.ErrorContext
	0, // 4: google.devtools.clouderrorreporting.v1beta1.ReportErrorsService.ReportErrorEvent:input_type -> google.devtools.clouderrorreporting.v1beta1.ReportErrorEventRequest
	1, // 5: google.devtools.clouderrorreporting.v1beta1.ReportErrorsService.ReportErrorEvent:output_type -> google.devtools.clouderrorreporting.v1beta1.ReportErrorEventResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_init() }
func file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_init() {
	if File_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto != nil {
		return
	}
	file_google_devtools_clouderrorreporting_v1beta1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportErrorEventRequest); i {
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
		file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportErrorEventResponse); i {
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
		file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportedErrorEvent); i {
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
			RawDescriptor: file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_goTypes,
		DependencyIndexes: file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_depIdxs,
		MessageInfos:      file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_msgTypes,
	}.Build()
	File_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto = out.File
	file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_rawDesc = nil
	file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_goTypes = nil
	file_google_devtools_clouderrorreporting_v1beta1_report_errors_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReportErrorsServiceClient is the client API for ReportErrorsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReportErrorsServiceClient interface {
	// Report an individual error event.
	//
	// This endpoint accepts **either** an OAuth token,
	// **or** an [API key](https://support.google.com/cloud/answer/6158862)
	// for authentication. To use an API key, append it to the URL as the value of
	// a `key` parameter. For example:
	//
	// `POST
	// https://clouderrorreporting.googleapis.com/v1beta1/projects/example-project/events:report?key=123ABC456`
	ReportErrorEvent(ctx context.Context, in *ReportErrorEventRequest, opts ...grpc.CallOption) (*ReportErrorEventResponse, error)
}

type reportErrorsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReportErrorsServiceClient(cc grpc.ClientConnInterface) ReportErrorsServiceClient {
	return &reportErrorsServiceClient{cc}
}

func (c *reportErrorsServiceClient) ReportErrorEvent(ctx context.Context, in *ReportErrorEventRequest, opts ...grpc.CallOption) (*ReportErrorEventResponse, error) {
	out := new(ReportErrorEventResponse)
	err := c.cc.Invoke(ctx, "/google.devtools.clouderrorreporting.v1beta1.ReportErrorsService/ReportErrorEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportErrorsServiceServer is the server API for ReportErrorsService service.
type ReportErrorsServiceServer interface {
	// Report an individual error event.
	//
	// This endpoint accepts **either** an OAuth token,
	// **or** an [API key](https://support.google.com/cloud/answer/6158862)
	// for authentication. To use an API key, append it to the URL as the value of
	// a `key` parameter. For example:
	//
	// `POST
	// https://clouderrorreporting.googleapis.com/v1beta1/projects/example-project/events:report?key=123ABC456`
	ReportErrorEvent(context.Context, *ReportErrorEventRequest) (*ReportErrorEventResponse, error)
}

// UnimplementedReportErrorsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedReportErrorsServiceServer struct {
}

func (*UnimplementedReportErrorsServiceServer) ReportErrorEvent(context.Context, *ReportErrorEventRequest) (*ReportErrorEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportErrorEvent not implemented")
}

func RegisterReportErrorsServiceServer(s *grpc.Server, srv ReportErrorsServiceServer) {
	s.RegisterService(&_ReportErrorsService_serviceDesc, srv)
}

func _ReportErrorsService_ReportErrorEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportErrorEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportErrorsServiceServer).ReportErrorEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouderrorreporting.v1beta1.ReportErrorsService/ReportErrorEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportErrorsServiceServer).ReportErrorEvent(ctx, req.(*ReportErrorEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReportErrorsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.devtools.clouderrorreporting.v1beta1.ReportErrorsService",
	HandlerType: (*ReportErrorsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportErrorEvent",
			Handler:    _ReportErrorsService_ReportErrorEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/devtools/clouderrorreporting/v1beta1/report_errors_service.proto",
}
