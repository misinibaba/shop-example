// Copyright 2017 Google Inc.
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
// source: google/api/servicecontrol/v1/log_entry.proto

package servicecontrol

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_type "google.golang.org/genproto/googleapis/logging/type"
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

// An individual log entry.
type LogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The log to which this log entry belongs. Examples: `"syslog"`,
	// `"book_log"`.
	Name string `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	// The time the event described by the log entry occurred. If
	// omitted, defaults to operation start time.
	Timestamp *timestamp.Timestamp `protobuf:"bytes,11,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The severity of the log entry. The default value is
	// `LogSeverity.DEFAULT`.
	Severity _type.LogSeverity `protobuf:"varint,12,opt,name=severity,proto3,enum=google.logging.type.LogSeverity" json:"severity,omitempty"`
	// A unique ID for the log entry used for deduplication. If omitted,
	// the implementation will generate one based on operation_id.
	InsertId string `protobuf:"bytes,4,opt,name=insert_id,json=insertId,proto3" json:"insert_id,omitempty"`
	// A set of user-defined (key, value) data that provides additional
	// information about the log entry.
	Labels map[string]string `protobuf:"bytes,13,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The log entry payload, which can be one of multiple types.
	//
	// Types that are assignable to Payload:
	//	*LogEntry_ProtoPayload
	//	*LogEntry_TextPayload
	//	*LogEntry_StructPayload
	Payload isLogEntry_Payload `protobuf_oneof:"payload"`
}

func (x *LogEntry) Reset() {
	*x = LogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_servicecontrol_v1_log_entry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntry) ProtoMessage() {}

func (x *LogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_servicecontrol_v1_log_entry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntry.ProtoReflect.Descriptor instead.
func (*LogEntry) Descriptor() ([]byte, []int) {
	return file_google_api_servicecontrol_v1_log_entry_proto_rawDescGZIP(), []int{0}
}

func (x *LogEntry) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LogEntry) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *LogEntry) GetSeverity() _type.LogSeverity {
	if x != nil {
		return x.Severity
	}
	return _type.LogSeverity_DEFAULT
}

func (x *LogEntry) GetInsertId() string {
	if x != nil {
		return x.InsertId
	}
	return ""
}

func (x *LogEntry) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (m *LogEntry) GetPayload() isLogEntry_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *LogEntry) GetProtoPayload() *any.Any {
	if x, ok := x.GetPayload().(*LogEntry_ProtoPayload); ok {
		return x.ProtoPayload
	}
	return nil
}

func (x *LogEntry) GetTextPayload() string {
	if x, ok := x.GetPayload().(*LogEntry_TextPayload); ok {
		return x.TextPayload
	}
	return ""
}

func (x *LogEntry) GetStructPayload() *_struct.Struct {
	if x, ok := x.GetPayload().(*LogEntry_StructPayload); ok {
		return x.StructPayload
	}
	return nil
}

type isLogEntry_Payload interface {
	isLogEntry_Payload()
}

type LogEntry_ProtoPayload struct {
	// The log entry payload, represented as a protocol buffer that is
	// expressed as a JSON object. The only accepted type currently is
	// [AuditLog][google.cloud.audit.AuditLog].
	ProtoPayload *any.Any `protobuf:"bytes,2,opt,name=proto_payload,json=protoPayload,proto3,oneof"`
}

type LogEntry_TextPayload struct {
	// The log entry payload, represented as a Unicode string (UTF-8).
	TextPayload string `protobuf:"bytes,3,opt,name=text_payload,json=textPayload,proto3,oneof"`
}

type LogEntry_StructPayload struct {
	// The log entry payload, represented as a structure that
	// is expressed as a JSON object.
	StructPayload *_struct.Struct `protobuf:"bytes,6,opt,name=struct_payload,json=structPayload,proto3,oneof"`
}

func (*LogEntry_ProtoPayload) isLogEntry_Payload() {}

func (*LogEntry_TextPayload) isLogEntry_Payload() {}

func (*LogEntry_StructPayload) isLogEntry_Payload() {}

var File_google_api_servicecontrol_v1_log_entry_proto protoreflect.FileDescriptor

var file_google_api_servicecontrol_v1_log_entry_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6c,
	0x6f, 0x67, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x03, 0x0a,
	0x08, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3c, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x4c, 0x6f, 0x67, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76,
	0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74,
	0x49, 0x64, 0x12, 0x4a, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0d, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x3b,
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x0c, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x23, 0x0a, 0x0c, 0x74,
	0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x65, 0x78, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x40, 0x0a, 0x0e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x09, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x7f, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x4c, 0x6f,
	0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_api_servicecontrol_v1_log_entry_proto_rawDescOnce sync.Once
	file_google_api_servicecontrol_v1_log_entry_proto_rawDescData = file_google_api_servicecontrol_v1_log_entry_proto_rawDesc
)

func file_google_api_servicecontrol_v1_log_entry_proto_rawDescGZIP() []byte {
	file_google_api_servicecontrol_v1_log_entry_proto_rawDescOnce.Do(func() {
		file_google_api_servicecontrol_v1_log_entry_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_api_servicecontrol_v1_log_entry_proto_rawDescData)
	})
	return file_google_api_servicecontrol_v1_log_entry_proto_rawDescData
}

var file_google_api_servicecontrol_v1_log_entry_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_api_servicecontrol_v1_log_entry_proto_goTypes = []interface{}{
	(*LogEntry)(nil),            // 0: google.api.servicecontrol.v1.LogEntry
	nil,                         // 1: google.api.servicecontrol.v1.LogEntry.LabelsEntry
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(_type.LogSeverity)(0),      // 3: google.logging.type.LogSeverity
	(*any.Any)(nil),             // 4: google.protobuf.Any
	(*_struct.Struct)(nil),      // 5: google.protobuf.Struct
}
var file_google_api_servicecontrol_v1_log_entry_proto_depIdxs = []int32{
	2, // 0: google.api.servicecontrol.v1.LogEntry.timestamp:type_name -> google.protobuf.Timestamp
	3, // 1: google.api.servicecontrol.v1.LogEntry.severity:type_name -> google.logging.type.LogSeverity
	1, // 2: google.api.servicecontrol.v1.LogEntry.labels:type_name -> google.api.servicecontrol.v1.LogEntry.LabelsEntry
	4, // 3: google.api.servicecontrol.v1.LogEntry.proto_payload:type_name -> google.protobuf.Any
	5, // 4: google.api.servicecontrol.v1.LogEntry.struct_payload:type_name -> google.protobuf.Struct
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_api_servicecontrol_v1_log_entry_proto_init() }
func file_google_api_servicecontrol_v1_log_entry_proto_init() {
	if File_google_api_servicecontrol_v1_log_entry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_api_servicecontrol_v1_log_entry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntry); i {
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
	file_google_api_servicecontrol_v1_log_entry_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*LogEntry_ProtoPayload)(nil),
		(*LogEntry_TextPayload)(nil),
		(*LogEntry_StructPayload)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_api_servicecontrol_v1_log_entry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_api_servicecontrol_v1_log_entry_proto_goTypes,
		DependencyIndexes: file_google_api_servicecontrol_v1_log_entry_proto_depIdxs,
		MessageInfos:      file_google_api_servicecontrol_v1_log_entry_proto_msgTypes,
	}.Build()
	File_google_api_servicecontrol_v1_log_entry_proto = out.File
	file_google_api_servicecontrol_v1_log_entry_proto_rawDesc = nil
	file_google_api_servicecontrol_v1_log_entry_proto_goTypes = nil
	file_google_api_servicecontrol_v1_log_entry_proto_depIdxs = nil
}
