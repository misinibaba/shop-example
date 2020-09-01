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
// source: google/datastore/v1beta3/entity.proto

package datastore

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
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

// A partition ID identifies a grouping of entities. The grouping is always
// by project and namespace, however the namespace ID may be empty.
//
// A partition ID contains several dimensions:
// project ID and namespace ID.
//
// Partition dimensions:
//
// - May be `""`.
// - Must be valid UTF-8 bytes.
// - Must have values that match regex `[A-Za-z\d\.\-_]{1,100}`
// If the value of any dimension matches regex `__.*__`, the partition is
// reserved/read-only.
// A reserved/read-only partition ID is forbidden in certain documented
// contexts.
//
// Foreign partition IDs (in which the project ID does
// not match the context project ID ) are discouraged.
// Reads and writes of foreign partition IDs may fail if the project is not in
// an active state.
type PartitionId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the project to which the entities belong.
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// If not empty, the ID of the namespace to which the entities belong.
	NamespaceId string `protobuf:"bytes,4,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
}

func (x *PartitionId) Reset() {
	*x = PartitionId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_datastore_v1beta3_entity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartitionId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartitionId) ProtoMessage() {}

func (x *PartitionId) ProtoReflect() protoreflect.Message {
	mi := &file_google_datastore_v1beta3_entity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartitionId.ProtoReflect.Descriptor instead.
func (*PartitionId) Descriptor() ([]byte, []int) {
	return file_google_datastore_v1beta3_entity_proto_rawDescGZIP(), []int{0}
}

func (x *PartitionId) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *PartitionId) GetNamespaceId() string {
	if x != nil {
		return x.NamespaceId
	}
	return ""
}

// A unique identifier for an entity.
// If a key's partition ID or any of its path kinds or names are
// reserved/read-only, the key is reserved/read-only.
// A reserved/read-only key is forbidden in certain documented contexts.
type Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Entities are partitioned into subsets, currently identified by a project
	// ID and namespace ID.
	// Queries are scoped to a single partition.
	PartitionId *PartitionId `protobuf:"bytes,1,opt,name=partition_id,json=partitionId,proto3" json:"partition_id,omitempty"`
	// The entity path.
	// An entity path consists of one or more elements composed of a kind and a
	// string or numerical identifier, which identify entities. The first
	// element identifies a _root entity_, the second element identifies
	// a _child_ of the root entity, the third element identifies a child of the
	// second entity, and so forth. The entities identified by all prefixes of
	// the path are called the element's _ancestors_.
	//
	// An entity path is always fully complete: *all* of the entity's ancestors
	// are required to be in the path along with the entity identifier itself.
	// The only exception is that in some documented cases, the identifier in the
	// last path element (for the entity) itself may be omitted. For example,
	// the last path element of the key of `Mutation.insert` may have no
	// identifier.
	//
	// A path can never be empty, and a path can have at most 100 elements.
	Path []*Key_PathElement `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
}

func (x *Key) Reset() {
	*x = Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_datastore_v1beta3_entity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Key) ProtoMessage() {}

func (x *Key) ProtoReflect() protoreflect.Message {
	mi := &file_google_datastore_v1beta3_entity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Key.ProtoReflect.Descriptor instead.
func (*Key) Descriptor() ([]byte, []int) {
	return file_google_datastore_v1beta3_entity_proto_rawDescGZIP(), []int{1}
}

func (x *Key) GetPartitionId() *PartitionId {
	if x != nil {
		return x.PartitionId
	}
	return nil
}

func (x *Key) GetPath() []*Key_PathElement {
	if x != nil {
		return x.Path
	}
	return nil
}

// An array value.
type ArrayValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Values in the array.
	// The order of this array may not be preserved if it contains a mix of
	// indexed and unindexed values.
	Values []*Value `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *ArrayValue) Reset() {
	*x = ArrayValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_datastore_v1beta3_entity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArrayValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrayValue) ProtoMessage() {}

func (x *ArrayValue) ProtoReflect() protoreflect.Message {
	mi := &file_google_datastore_v1beta3_entity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrayValue.ProtoReflect.Descriptor instead.
func (*ArrayValue) Descriptor() ([]byte, []int) {
	return file_google_datastore_v1beta3_entity_proto_rawDescGZIP(), []int{2}
}

func (x *ArrayValue) GetValues() []*Value {
	if x != nil {
		return x.Values
	}
	return nil
}

// A message that can hold any of the supported value types and associated
// metadata.
type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Must have a value set.
	//
	// Types that are assignable to ValueType:
	//	*Value_NullValue
	//	*Value_BooleanValue
	//	*Value_IntegerValue
	//	*Value_DoubleValue
	//	*Value_TimestampValue
	//	*Value_KeyValue
	//	*Value_StringValue
	//	*Value_BlobValue
	//	*Value_GeoPointValue
	//	*Value_EntityValue
	//	*Value_ArrayValue
	ValueType isValue_ValueType `protobuf_oneof:"value_type"`
	// The `meaning` field should only be populated for backwards compatibility.
	Meaning int32 `protobuf:"varint,14,opt,name=meaning,proto3" json:"meaning,omitempty"`
	// If the value should be excluded from all indexes including those defined
	// explicitly.
	ExcludeFromIndexes bool `protobuf:"varint,19,opt,name=exclude_from_indexes,json=excludeFromIndexes,proto3" json:"exclude_from_indexes,omitempty"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_datastore_v1beta3_entity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_google_datastore_v1beta3_entity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_google_datastore_v1beta3_entity_proto_rawDescGZIP(), []int{3}
}

func (m *Value) GetValueType() isValue_ValueType {
	if m != nil {
		return m.ValueType
	}
	return nil
}

func (x *Value) GetNullValue() _struct.NullValue {
	if x, ok := x.GetValueType().(*Value_NullValue); ok {
		return x.NullValue
	}
	return _struct.NullValue_NULL_VALUE
}

func (x *Value) GetBooleanValue() bool {
	if x, ok := x.GetValueType().(*Value_BooleanValue); ok {
		return x.BooleanValue
	}
	return false
}

func (x *Value) GetIntegerValue() int64 {
	if x, ok := x.GetValueType().(*Value_IntegerValue); ok {
		return x.IntegerValue
	}
	return 0
}

func (x *Value) GetDoubleValue() float64 {
	if x, ok := x.GetValueType().(*Value_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (x *Value) GetTimestampValue() *timestamp.Timestamp {
	if x, ok := x.GetValueType().(*Value_TimestampValue); ok {
		return x.TimestampValue
	}
	return nil
}

func (x *Value) GetKeyValue() *Key {
	if x, ok := x.GetValueType().(*Value_KeyValue); ok {
		return x.KeyValue
	}
	return nil
}

func (x *Value) GetStringValue() string {
	if x, ok := x.GetValueType().(*Value_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *Value) GetBlobValue() []byte {
	if x, ok := x.GetValueType().(*Value_BlobValue); ok {
		return x.BlobValue
	}
	return nil
}

func (x *Value) GetGeoPointValue() *latlng.LatLng {
	if x, ok := x.GetValueType().(*Value_GeoPointValue); ok {
		return x.GeoPointValue
	}
	return nil
}

func (x *Value) GetEntityValue() *Entity {
	if x, ok := x.GetValueType().(*Value_EntityValue); ok {
		return x.EntityValue
	}
	return nil
}

func (x *Value) GetArrayValue() *ArrayValue {
	if x, ok := x.GetValueType().(*Value_ArrayValue); ok {
		return x.ArrayValue
	}
	return nil
}

func (x *Value) GetMeaning() int32 {
	if x != nil {
		return x.Meaning
	}
	return 0
}

func (x *Value) GetExcludeFromIndexes() bool {
	if x != nil {
		return x.ExcludeFromIndexes
	}
	return false
}

type isValue_ValueType interface {
	isValue_ValueType()
}

type Value_NullValue struct {
	// A null value.
	NullValue _struct.NullValue `protobuf:"varint,11,opt,name=null_value,json=nullValue,proto3,enum=google.protobuf.NullValue,oneof"`
}

type Value_BooleanValue struct {
	// A boolean value.
	BooleanValue bool `protobuf:"varint,1,opt,name=boolean_value,json=booleanValue,proto3,oneof"`
}

type Value_IntegerValue struct {
	// An integer value.
	IntegerValue int64 `protobuf:"varint,2,opt,name=integer_value,json=integerValue,proto3,oneof"`
}

type Value_DoubleValue struct {
	// A double value.
	DoubleValue float64 `protobuf:"fixed64,3,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type Value_TimestampValue struct {
	// A timestamp value.
	// When stored in the Datastore, precise only to microseconds;
	// any additional precision is rounded down.
	TimestampValue *timestamp.Timestamp `protobuf:"bytes,10,opt,name=timestamp_value,json=timestampValue,proto3,oneof"`
}

type Value_KeyValue struct {
	// A key value.
	KeyValue *Key `protobuf:"bytes,5,opt,name=key_value,json=keyValue,proto3,oneof"`
}

type Value_StringValue struct {
	// A UTF-8 encoded string value.
	// When `exclude_from_indexes` is false (it is indexed), may have at most
	// 1500 bytes. Otherwise, may be set to at most 1,000,000 bytes.
	StringValue string `protobuf:"bytes,17,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Value_BlobValue struct {
	// A blob value.
	// May have at most 1,000,000 bytes.
	// When `exclude_from_indexes` is false, may have at most 1500 bytes.
	// In JSON requests, must be base64-encoded.
	BlobValue []byte `protobuf:"bytes,18,opt,name=blob_value,json=blobValue,proto3,oneof"`
}

type Value_GeoPointValue struct {
	// A geo point value representing a point on the surface of Earth.
	GeoPointValue *latlng.LatLng `protobuf:"bytes,8,opt,name=geo_point_value,json=geoPointValue,proto3,oneof"`
}

type Value_EntityValue struct {
	// An entity value.
	//
	// - May have no key.
	// - May have a key with an incomplete key path.
	// - May have a reserved/read-only key.
	EntityValue *Entity `protobuf:"bytes,6,opt,name=entity_value,json=entityValue,proto3,oneof"`
}

type Value_ArrayValue struct {
	// An array value.
	// Cannot contain another array value.
	// A `Value` instance that sets field `array_value` must not set fields
	// `meaning` or `exclude_from_indexes`.
	ArrayValue *ArrayValue `protobuf:"bytes,9,opt,name=array_value,json=arrayValue,proto3,oneof"`
}

func (*Value_NullValue) isValue_ValueType() {}

func (*Value_BooleanValue) isValue_ValueType() {}

func (*Value_IntegerValue) isValue_ValueType() {}

func (*Value_DoubleValue) isValue_ValueType() {}

func (*Value_TimestampValue) isValue_ValueType() {}

func (*Value_KeyValue) isValue_ValueType() {}

func (*Value_StringValue) isValue_ValueType() {}

func (*Value_BlobValue) isValue_ValueType() {}

func (*Value_GeoPointValue) isValue_ValueType() {}

func (*Value_EntityValue) isValue_ValueType() {}

func (*Value_ArrayValue) isValue_ValueType() {}

// A Datastore data object.
//
// An entity is limited to 1 megabyte when stored. That _roughly_
// corresponds to a limit of 1 megabyte for the serialized form of this
// message.
type Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The entity's key.
	//
	// An entity must have a key, unless otherwise documented (for example,
	// an entity in `Value.entity_value` may have no key).
	// An entity's kind is its key path's last element's kind,
	// or null if it has no key.
	Key *Key `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The entity's properties.
	// The map's keys are property names.
	// A property name matching regex `__.*__` is reserved.
	// A reserved property name is forbidden in certain documented contexts.
	// The name must not contain more than 500 characters.
	// The name cannot be `""`.
	Properties map[string]*Value `protobuf:"bytes,3,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Entity) Reset() {
	*x = Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_datastore_v1beta3_entity_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_google_datastore_v1beta3_entity_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_google_datastore_v1beta3_entity_proto_rawDescGZIP(), []int{4}
}

func (x *Entity) GetKey() *Key {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Entity) GetProperties() map[string]*Value {
	if x != nil {
		return x.Properties
	}
	return nil
}

// A (kind, ID/name) pair used to construct a key path.
//
// If either name or ID is set, the element is complete.
// If neither is set, the element is incomplete.
type Key_PathElement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The kind of the entity.
	// A kind matching regex `__.*__` is reserved/read-only.
	// A kind must not contain more than 1500 bytes when UTF-8 encoded.
	// Cannot be `""`.
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// The type of ID.
	//
	// Types that are assignable to IdType:
	//	*Key_PathElement_Id
	//	*Key_PathElement_Name
	IdType isKey_PathElement_IdType `protobuf_oneof:"id_type"`
}

func (x *Key_PathElement) Reset() {
	*x = Key_PathElement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_datastore_v1beta3_entity_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Key_PathElement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Key_PathElement) ProtoMessage() {}

func (x *Key_PathElement) ProtoReflect() protoreflect.Message {
	mi := &file_google_datastore_v1beta3_entity_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Key_PathElement.ProtoReflect.Descriptor instead.
func (*Key_PathElement) Descriptor() ([]byte, []int) {
	return file_google_datastore_v1beta3_entity_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Key_PathElement) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (m *Key_PathElement) GetIdType() isKey_PathElement_IdType {
	if m != nil {
		return m.IdType
	}
	return nil
}

func (x *Key_PathElement) GetId() int64 {
	if x, ok := x.GetIdType().(*Key_PathElement_Id); ok {
		return x.Id
	}
	return 0
}

func (x *Key_PathElement) GetName() string {
	if x, ok := x.GetIdType().(*Key_PathElement_Name); ok {
		return x.Name
	}
	return ""
}

type isKey_PathElement_IdType interface {
	isKey_PathElement_IdType()
}

type Key_PathElement_Id struct {
	// The auto-allocated ID of the entity.
	// Never equal to zero. Values less than zero are discouraged and may not
	// be supported in the future.
	Id int64 `protobuf:"varint,2,opt,name=id,proto3,oneof"`
}

type Key_PathElement_Name struct {
	// The name of the entity.
	// A name matching regex `__.*__` is reserved/read-only.
	// A name must not be more than 1500 bytes when UTF-8 encoded.
	// Cannot be `""`.
	Name string `protobuf:"bytes,3,opt,name=name,proto3,oneof"`
}

func (*Key_PathElement_Id) isKey_PathElement_IdType() {}

func (*Key_PathElement_Name) isKey_PathElement_IdType() {}

var File_google_datastore_v1beta3_entity_proto protoreflect.FileDescriptor

var file_google_datastore_v1beta3_entity_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x33, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x33, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6c, 0x61, 0x74, 0x6c,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0xe4, 0x01, 0x0a, 0x03, 0x4b, 0x65,
	0x79, 0x12, 0x48, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x33, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x52, 0x0b,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x33, 0x2e, 0x4b, 0x65, 0x79, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x45, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x1a, 0x54, 0x0a, 0x0b, 0x50, 0x61,
	0x74, 0x68, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x10, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x45, 0x0a, 0x0a, 0x41, 0x72, 0x72, 0x61, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x37,
	0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x33, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0xab, 0x05, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x6e, 0x75, 0x6c, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x75, 0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x25,
	0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0c, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x25, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0c,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x45, 0x0a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x33, 0x2e, 0x4b, 0x65, 0x79, 0x48, 0x00, 0x52, 0x08, 0x6b, 0x65,
	0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x62,
	0x6c, 0x6f, 0x62, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0c, 0x48,
	0x00, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x62, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3d, 0x0a, 0x0f,
	0x67, 0x65, 0x6f, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x4c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x0d, 0x67, 0x65,
	0x6f, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x33, 0x2e, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x48, 0x00, 0x52, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x47, 0x0a, 0x0b, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x33, 0x2e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52,
	0x0a, 0x61, 0x72, 0x72, 0x61, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x61, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x65,
	0x61, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x30, 0x0a, 0x14, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x18, 0x13, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x12, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x46, 0x72, 0x6f, 0x6d,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x42, 0x0c, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0xeb, 0x01, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x2f, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x33, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x50, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x33,
	0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x1a, 0x5e, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x35, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x33, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x42, 0xb2, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x33, 0x42, 0x0b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x33, 0x3b, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x33, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_datastore_v1beta3_entity_proto_rawDescOnce sync.Once
	file_google_datastore_v1beta3_entity_proto_rawDescData = file_google_datastore_v1beta3_entity_proto_rawDesc
)

func file_google_datastore_v1beta3_entity_proto_rawDescGZIP() []byte {
	file_google_datastore_v1beta3_entity_proto_rawDescOnce.Do(func() {
		file_google_datastore_v1beta3_entity_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_datastore_v1beta3_entity_proto_rawDescData)
	})
	return file_google_datastore_v1beta3_entity_proto_rawDescData
}

var file_google_datastore_v1beta3_entity_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_google_datastore_v1beta3_entity_proto_goTypes = []interface{}{
	(*PartitionId)(nil),         // 0: google.datastore.v1beta3.PartitionId
	(*Key)(nil),                 // 1: google.datastore.v1beta3.Key
	(*ArrayValue)(nil),          // 2: google.datastore.v1beta3.ArrayValue
	(*Value)(nil),               // 3: google.datastore.v1beta3.Value
	(*Entity)(nil),              // 4: google.datastore.v1beta3.Entity
	(*Key_PathElement)(nil),     // 5: google.datastore.v1beta3.Key.PathElement
	nil,                         // 6: google.datastore.v1beta3.Entity.PropertiesEntry
	(_struct.NullValue)(0),      // 7: google.protobuf.NullValue
	(*timestamp.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*latlng.LatLng)(nil),       // 9: google.type.LatLng
}
var file_google_datastore_v1beta3_entity_proto_depIdxs = []int32{
	0,  // 0: google.datastore.v1beta3.Key.partition_id:type_name -> google.datastore.v1beta3.PartitionId
	5,  // 1: google.datastore.v1beta3.Key.path:type_name -> google.datastore.v1beta3.Key.PathElement
	3,  // 2: google.datastore.v1beta3.ArrayValue.values:type_name -> google.datastore.v1beta3.Value
	7,  // 3: google.datastore.v1beta3.Value.null_value:type_name -> google.protobuf.NullValue
	8,  // 4: google.datastore.v1beta3.Value.timestamp_value:type_name -> google.protobuf.Timestamp
	1,  // 5: google.datastore.v1beta3.Value.key_value:type_name -> google.datastore.v1beta3.Key
	9,  // 6: google.datastore.v1beta3.Value.geo_point_value:type_name -> google.type.LatLng
	4,  // 7: google.datastore.v1beta3.Value.entity_value:type_name -> google.datastore.v1beta3.Entity
	2,  // 8: google.datastore.v1beta3.Value.array_value:type_name -> google.datastore.v1beta3.ArrayValue
	1,  // 9: google.datastore.v1beta3.Entity.key:type_name -> google.datastore.v1beta3.Key
	6,  // 10: google.datastore.v1beta3.Entity.properties:type_name -> google.datastore.v1beta3.Entity.PropertiesEntry
	3,  // 11: google.datastore.v1beta3.Entity.PropertiesEntry.value:type_name -> google.datastore.v1beta3.Value
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_google_datastore_v1beta3_entity_proto_init() }
func file_google_datastore_v1beta3_entity_proto_init() {
	if File_google_datastore_v1beta3_entity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_datastore_v1beta3_entity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartitionId); i {
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
		file_google_datastore_v1beta3_entity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Key); i {
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
		file_google_datastore_v1beta3_entity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArrayValue); i {
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
		file_google_datastore_v1beta3_entity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_google_datastore_v1beta3_entity_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entity); i {
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
		file_google_datastore_v1beta3_entity_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Key_PathElement); i {
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
	file_google_datastore_v1beta3_entity_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Value_NullValue)(nil),
		(*Value_BooleanValue)(nil),
		(*Value_IntegerValue)(nil),
		(*Value_DoubleValue)(nil),
		(*Value_TimestampValue)(nil),
		(*Value_KeyValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_BlobValue)(nil),
		(*Value_GeoPointValue)(nil),
		(*Value_EntityValue)(nil),
		(*Value_ArrayValue)(nil),
	}
	file_google_datastore_v1beta3_entity_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*Key_PathElement_Id)(nil),
		(*Key_PathElement_Name)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_datastore_v1beta3_entity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_datastore_v1beta3_entity_proto_goTypes,
		DependencyIndexes: file_google_datastore_v1beta3_entity_proto_depIdxs,
		MessageInfos:      file_google_datastore_v1beta3_entity_proto_msgTypes,
	}.Build()
	File_google_datastore_v1beta3_entity_proto = out.File
	file_google_datastore_v1beta3_entity_proto_rawDesc = nil
	file_google_datastore_v1beta3_entity_proto_goTypes = nil
	file_google_datastore_v1beta3_entity_proto_depIdxs = nil
}
