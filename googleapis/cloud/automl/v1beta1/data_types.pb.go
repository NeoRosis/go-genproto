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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/cloud/automl/v1beta1/data_types.proto

package automl

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// `TypeCode` is used as a part of
// [DataType][google.cloud.automl.v1beta1.DataType].
type TypeCode int32

const (
	// Not specified. Should not be used.
	TypeCode_TYPE_CODE_UNSPECIFIED TypeCode = 0
	// Encoded as `number`, or the strings `"NaN"`, `"Infinity"`, or
	// `"-Infinity"`.
	TypeCode_FLOAT64 TypeCode = 3
	// Must be between 0AD and 9999AD. Encoded as `string` according to
	// [time_format][google.cloud.automl.v1beta1.DataType.time_format], or, if
	// that format is not set, then in RFC 3339 `date-time` format, where
	// `time-offset` = `"Z"` (e.g. 1985-04-12T23:20:50.52Z).
	TypeCode_TIMESTAMP TypeCode = 4
	// Encoded as `string`.
	TypeCode_STRING TypeCode = 6
	// Encoded as `list`, where the list elements are represented according to
	//
	// [list_element_type][google.cloud.automl.v1beta1.DataType.list_element_type].
	TypeCode_ARRAY TypeCode = 8
	// Encoded as `struct`, where field values are represented according to
	// [struct_type][google.cloud.automl.v1beta1.DataType.struct_type].
	TypeCode_STRUCT TypeCode = 9
	// Values of this type are not further understood by AutoML,
	// e.g. AutoML is unable to tell the order of values (as it could with
	// FLOAT64), or is unable to say if one value contains another (as it
	// could with STRING).
	// Encoded as `string` (bytes should be base64-encoded, as described in RFC
	// 4648, section 4).
	TypeCode_CATEGORY TypeCode = 10
)

// Enum value maps for TypeCode.
var (
	TypeCode_name = map[int32]string{
		0:  "TYPE_CODE_UNSPECIFIED",
		3:  "FLOAT64",
		4:  "TIMESTAMP",
		6:  "STRING",
		8:  "ARRAY",
		9:  "STRUCT",
		10: "CATEGORY",
	}
	TypeCode_value = map[string]int32{
		"TYPE_CODE_UNSPECIFIED": 0,
		"FLOAT64":               3,
		"TIMESTAMP":             4,
		"STRING":                6,
		"ARRAY":                 8,
		"STRUCT":                9,
		"CATEGORY":              10,
	}
)

func (x TypeCode) Enum() *TypeCode {
	p := new(TypeCode)
	*p = x
	return p
}

func (x TypeCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TypeCode) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_automl_v1beta1_data_types_proto_enumTypes[0].Descriptor()
}

func (TypeCode) Type() protoreflect.EnumType {
	return &file_google_cloud_automl_v1beta1_data_types_proto_enumTypes[0]
}

func (x TypeCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TypeCode.Descriptor instead.
func (TypeCode) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_data_types_proto_rawDescGZIP(), []int{0}
}

// Indicated the type of data that can be stored in a structured data entity
// (e.g. a table).
type DataType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Details of DataType-s that need additional specification.
	//
	// Types that are assignable to Details:
	//
	//	*DataType_ListElementType
	//	*DataType_StructType
	//	*DataType_TimeFormat
	Details isDataType_Details `protobuf_oneof:"details"`
	// Required. The [TypeCode][google.cloud.automl.v1beta1.TypeCode] for this type.
	TypeCode TypeCode `protobuf:"varint,1,opt,name=type_code,json=typeCode,proto3,enum=google.cloud.automl.v1beta1.TypeCode" json:"type_code,omitempty"`
	// If true, this DataType can also be `NULL`. In .CSV files `NULL` value is
	// expressed as an empty string.
	Nullable bool `protobuf:"varint,4,opt,name=nullable,proto3" json:"nullable,omitempty"`
}

func (x *DataType) Reset() {
	*x = DataType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_data_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType) ProtoMessage() {}

func (x *DataType) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_data_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType.ProtoReflect.Descriptor instead.
func (*DataType) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_data_types_proto_rawDescGZIP(), []int{0}
}

func (m *DataType) GetDetails() isDataType_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (x *DataType) GetListElementType() *DataType {
	if x, ok := x.GetDetails().(*DataType_ListElementType); ok {
		return x.ListElementType
	}
	return nil
}

func (x *DataType) GetStructType() *StructType {
	if x, ok := x.GetDetails().(*DataType_StructType); ok {
		return x.StructType
	}
	return nil
}

func (x *DataType) GetTimeFormat() string {
	if x, ok := x.GetDetails().(*DataType_TimeFormat); ok {
		return x.TimeFormat
	}
	return ""
}

func (x *DataType) GetTypeCode() TypeCode {
	if x != nil {
		return x.TypeCode
	}
	return TypeCode_TYPE_CODE_UNSPECIFIED
}

func (x *DataType) GetNullable() bool {
	if x != nil {
		return x.Nullable
	}
	return false
}

type isDataType_Details interface {
	isDataType_Details()
}

type DataType_ListElementType struct {
	// If [type_code][google.cloud.automl.v1beta1.DataType.type_code] == [ARRAY][google.cloud.automl.v1beta1.TypeCode.ARRAY],
	// then `list_element_type` is the type of the elements.
	ListElementType *DataType `protobuf:"bytes,2,opt,name=list_element_type,json=listElementType,proto3,oneof"`
}

type DataType_StructType struct {
	// If [type_code][google.cloud.automl.v1beta1.DataType.type_code] == [STRUCT][google.cloud.automl.v1beta1.TypeCode.STRUCT], then `struct_type`
	// provides type information for the struct's fields.
	StructType *StructType `protobuf:"bytes,3,opt,name=struct_type,json=structType,proto3,oneof"`
}

type DataType_TimeFormat struct {
	// If [type_code][google.cloud.automl.v1beta1.DataType.type_code] == [TIMESTAMP][google.cloud.automl.v1beta1.TypeCode.TIMESTAMP]
	// then `time_format` provides the format in which that time field is
	// expressed. The time_format must either be one of:
	// * `UNIX_SECONDS`
	// * `UNIX_MILLISECONDS`
	// * `UNIX_MICROSECONDS`
	// * `UNIX_NANOSECONDS`
	// (for respectively number of seconds, milliseconds, microseconds and
	// nanoseconds since start of the Unix epoch);
	// or be written in `strftime` syntax. If time_format is not set, then the
	// default format as described on the type_code is used.
	TimeFormat string `protobuf:"bytes,5,opt,name=time_format,json=timeFormat,proto3,oneof"`
}

func (*DataType_ListElementType) isDataType_Details() {}

func (*DataType_StructType) isDataType_Details() {}

func (*DataType_TimeFormat) isDataType_Details() {}

// `StructType` defines the DataType-s of a [STRUCT][google.cloud.automl.v1beta1.TypeCode.STRUCT] type.
type StructType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unordered map of struct field names to their data types.
	// Fields cannot be added or removed via Update. Their names and
	// data types are still mutable.
	Fields map[string]*DataType `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *StructType) Reset() {
	*x = StructType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_data_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StructType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StructType) ProtoMessage() {}

func (x *StructType) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_data_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StructType.ProtoReflect.Descriptor instead.
func (*StructType) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_data_types_proto_rawDescGZIP(), []int{1}
}

func (x *StructType) GetFields() map[string]*DataType {
	if x != nil {
		return x.Fields
	}
	return nil
}

var File_google_cloud_automl_v1beta1_data_types_proto protoreflect.FileDescriptor

var file_google_cloud_automl_v1beta1_data_types_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x22, 0xb9, 0x02, 0x0a, 0x08,
	0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x53, 0x0a, 0x11, 0x6c, 0x69, 0x73, 0x74,
	0x5f, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x6c, 0x69,
	0x73, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4a, 0x0a,
	0x0b, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x42, 0x0a, 0x09,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x75, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x6e, 0x75, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x09, 0x0a, 0x07,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0xbb, 0x01, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x1a, 0x60, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x3b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x72, 0x0a, 0x08, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x46, 0x4c, 0x4f, 0x41, 0x54, 0x36, 0x34, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x49, 0x4d,
	0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49,
	0x4e, 0x47, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x52, 0x52, 0x41, 0x59, 0x10, 0x08, 0x12,
	0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x55, 0x43, 0x54, 0x10, 0x09, 0x12, 0x0c, 0x0a, 0x08, 0x43,
	0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x10, 0x0a, 0x42, 0xa5, 0x01, 0x0a, 0x1f, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x50, 0x01, 0x5a,
	0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x6c, 0xca, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x6c, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xea, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x4c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_automl_v1beta1_data_types_proto_rawDescOnce sync.Once
	file_google_cloud_automl_v1beta1_data_types_proto_rawDescData = file_google_cloud_automl_v1beta1_data_types_proto_rawDesc
)

func file_google_cloud_automl_v1beta1_data_types_proto_rawDescGZIP() []byte {
	file_google_cloud_automl_v1beta1_data_types_proto_rawDescOnce.Do(func() {
		file_google_cloud_automl_v1beta1_data_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_automl_v1beta1_data_types_proto_rawDescData)
	})
	return file_google_cloud_automl_v1beta1_data_types_proto_rawDescData
}

var file_google_cloud_automl_v1beta1_data_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_automl_v1beta1_data_types_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_automl_v1beta1_data_types_proto_goTypes = []interface{}{
	(TypeCode)(0),      // 0: google.cloud.automl.v1beta1.TypeCode
	(*DataType)(nil),   // 1: google.cloud.automl.v1beta1.DataType
	(*StructType)(nil), // 2: google.cloud.automl.v1beta1.StructType
	nil,                // 3: google.cloud.automl.v1beta1.StructType.FieldsEntry
}
var file_google_cloud_automl_v1beta1_data_types_proto_depIdxs = []int32{
	1, // 0: google.cloud.automl.v1beta1.DataType.list_element_type:type_name -> google.cloud.automl.v1beta1.DataType
	2, // 1: google.cloud.automl.v1beta1.DataType.struct_type:type_name -> google.cloud.automl.v1beta1.StructType
	0, // 2: google.cloud.automl.v1beta1.DataType.type_code:type_name -> google.cloud.automl.v1beta1.TypeCode
	3, // 3: google.cloud.automl.v1beta1.StructType.fields:type_name -> google.cloud.automl.v1beta1.StructType.FieldsEntry
	1, // 4: google.cloud.automl.v1beta1.StructType.FieldsEntry.value:type_name -> google.cloud.automl.v1beta1.DataType
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_automl_v1beta1_data_types_proto_init() }
func file_google_cloud_automl_v1beta1_data_types_proto_init() {
	if File_google_cloud_automl_v1beta1_data_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_automl_v1beta1_data_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType); i {
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
		file_google_cloud_automl_v1beta1_data_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StructType); i {
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
	file_google_cloud_automl_v1beta1_data_types_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*DataType_ListElementType)(nil),
		(*DataType_StructType)(nil),
		(*DataType_TimeFormat)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_automl_v1beta1_data_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_automl_v1beta1_data_types_proto_goTypes,
		DependencyIndexes: file_google_cloud_automl_v1beta1_data_types_proto_depIdxs,
		EnumInfos:         file_google_cloud_automl_v1beta1_data_types_proto_enumTypes,
		MessageInfos:      file_google_cloud_automl_v1beta1_data_types_proto_msgTypes,
	}.Build()
	File_google_cloud_automl_v1beta1_data_types_proto = out.File
	file_google_cloud_automl_v1beta1_data_types_proto_rawDesc = nil
	file_google_cloud_automl_v1beta1_data_types_proto_goTypes = nil
	file_google_cloud_automl_v1beta1_data_types_proto_depIdxs = nil
}
