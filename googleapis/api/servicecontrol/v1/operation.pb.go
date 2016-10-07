// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/api/servicecontrol/v1/operation.proto
// DO NOT EDIT!

package google_api_servicecontrol_v1 // import "google.golang.org/genproto/googleapis/api/servicecontrol/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import google_protobuf3 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Defines the importance of the data contained in the operation.
type Operation_Importance int32

const (
	// The API implementation may cache and aggregate the data.
	// The data may be lost when rare and unexpected system failures occur.
	Operation_LOW Operation_Importance = 0
	// The API implementation doesn't cache and aggregate the data.
	// If the method returns successfully, it's guaranteed that the data has
	// been persisted in durable storage.
	Operation_HIGH Operation_Importance = 1
)

var Operation_Importance_name = map[int32]string{
	0: "LOW",
	1: "HIGH",
}
var Operation_Importance_value = map[string]int32{
	"LOW":  0,
	"HIGH": 1,
}

func (x Operation_Importance) String() string {
	return proto.EnumName(Operation_Importance_name, int32(x))
}
func (Operation_Importance) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 0} }

// Represents information regarding an operation.
type Operation struct {
	// Identity of the operation. This must be unique within the scope of the
	// service that generated the operation. If the service calls
	// Check() and Report() on the same operation, the two calls should carry
	// the same id.
	//
	// UUID version 4 is recommended, though not required.
	// In scenarios where an operation is computed from existing information
	// and an idempotent id is desirable for deduplication purpose, UUID version 5
	// is recommended. See RFC 4122 for details.
	OperationId string `protobuf:"bytes,1,opt,name=operation_id,json=operationId" json:"operation_id,omitempty"`
	// Fully qualified name of the operation. Reserved for future use.
	OperationName string `protobuf:"bytes,2,opt,name=operation_name,json=operationName" json:"operation_name,omitempty"`
	// Identity of the consumer who is using the service.
	// This field should be filled in for the operations initiated by a
	// consumer, but not for service-initiated operations that are
	// not related to a specific consumer.
	//
	// This can be in one of the following formats:
	//   project:<project_id>,
	//   project_number:<project_number>,
	//   api_key:<api_key>.
	ConsumerId string `protobuf:"bytes,3,opt,name=consumer_id,json=consumerId" json:"consumer_id,omitempty"`
	// Required. Start time of the operation.
	StartTime *google_protobuf3.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	// End time of the operation.
	// Required when the operation is used in [ServiceController.Report][google.api.servicecontrol.v1.ServiceController.Report],
	// but optional when the operation is used in [ServiceController.Check][google.api.servicecontrol.v1.ServiceController.Check].
	EndTime *google_protobuf3.Timestamp `protobuf:"bytes,5,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	// Labels describing the operation. Only the following labels are allowed:
	//
	// - Labels describing monitored resources as defined in
	//   the service configuration.
	// - Default labels of metric values. When specified, labels defined in the
	//   metric value override these default.
	// - The following labels defined by Google Cloud Platform:
	//     - `cloud.googleapis.com/location` describing the location where the
	//        operation happened,
	//     - `servicecontrol.googleapis.com/user_agent` describing the user agent
	//        of the API request,
	//     - `servicecontrol.googleapis.com/service_agent` describing the service
	//        used to handle the API request (e.g. ESP),
	//     - `servicecontrol.googleapis.com/platform` describing the platform
	//        where the API is served (e.g. GAE, GCE, GKE).
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Represents information about this operation. Each MetricValueSet
	// corresponds to a metric defined in the service configuration.
	// The data type used in the MetricValueSet must agree with
	// the data type specified in the metric definition.
	//
	// Within a single operation, it is not allowed to have more than one
	// MetricValue instances that have the same metric names and identical
	// label value combinations. If a request has such duplicated MetricValue
	// instances, the entire request is rejected with
	// an invalid argument error.
	MetricValueSets []*MetricValueSet `protobuf:"bytes,7,rep,name=metric_value_sets,json=metricValueSets" json:"metric_value_sets,omitempty"`
	// Represents information to be logged.
	LogEntries []*LogEntry `protobuf:"bytes,8,rep,name=log_entries,json=logEntries" json:"log_entries,omitempty"`
	// DO NOT USE. This is an experimental field.
	Importance Operation_Importance `protobuf:"varint,11,opt,name=importance,enum=google.api.servicecontrol.v1.Operation_Importance" json:"importance,omitempty"`
}

func (m *Operation) Reset()                    { *m = Operation{} }
func (m *Operation) String() string            { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()               {}
func (*Operation) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Operation) GetStartTime() *google_protobuf3.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Operation) GetEndTime() *google_protobuf3.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *Operation) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Operation) GetMetricValueSets() []*MetricValueSet {
	if m != nil {
		return m.MetricValueSets
	}
	return nil
}

func (m *Operation) GetLogEntries() []*LogEntry {
	if m != nil {
		return m.LogEntries
	}
	return nil
}

func init() {
	proto.RegisterType((*Operation)(nil), "google.api.servicecontrol.v1.Operation")
	proto.RegisterEnum("google.api.servicecontrol.v1.Operation_Importance", Operation_Importance_name, Operation_Importance_value)
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/api/servicecontrol/v1/operation.proto", fileDescriptor4)
}

var fileDescriptor4 = []byte{
	// 490 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x92, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x86, 0x71, 0xd3, 0x7c, 0x8d, 0x21, 0x84, 0x15, 0x07, 0x2b, 0x42, 0x6a, 0xa8, 0x04, 0xe2,
	0x80, 0x76, 0xd5, 0x54, 0x08, 0x0a, 0xb7, 0x4a, 0xa8, 0x8d, 0x1a, 0xda, 0xca, 0x20, 0xe0, 0x66,
	0x39, 0xce, 0x74, 0x59, 0x61, 0x7b, 0x2d, 0xef, 0x26, 0x52, 0xfe, 0x01, 0x3f, 0x99, 0x23, 0xeb,
	0xf5, 0x47, 0xd2, 0x4b, 0x40, 0xe4, 0x62, 0xed, 0xcc, 0xce, 0xfb, 0xec, 0xeb, 0x99, 0x81, 0x2b,
	0x2e, 0x25, 0x8f, 0x91, 0x72, 0x19, 0x87, 0x29, 0xa7, 0x32, 0xe7, 0x8c, 0x63, 0x9a, 0xe5, 0x52,
	0x4b, 0x56, 0x5e, 0x85, 0x99, 0x50, 0xcc, 0x7c, 0x98, 0xc2, 0x7c, 0x25, 0x22, 0x8c, 0x64, 0xaa,
	0x73, 0x19, 0xb3, 0xd5, 0x09, 0x93, 0x19, 0xe6, 0xa1, 0x16, 0x32, 0xa5, 0x56, 0x40, 0x9e, 0x55,
	0x30, 0x53, 0x4d, 0xef, 0x57, 0xd3, 0xd5, 0xc9, 0x68, 0xfa, 0x3f, 0x4f, 0xdd, 0x09, 0xce, 0xc2,
	0x34, 0x95, 0xda, 0xbe, 0xa3, 0xca, 0x87, 0x46, 0x7b, 0xb9, 0x8e, 0x25, 0x0f, 0xd0, 0x44, 0xeb,
	0x0a, 0x76, 0xbd, 0x0f, 0x2c, 0x41, 0x9d, 0x8b, 0x28, 0x58, 0x85, 0xf1, 0x12, 0x2b, 0xde, 0x07,
	0x2e, 0xf4, 0x8f, 0xe5, 0x9c, 0x46, 0x32, 0x61, 0x25, 0x93, 0xd9, 0x8b, 0xf9, 0xf2, 0x8e, 0x65,
	0x7a, 0x9d, 0xa1, 0x62, 0x5a, 0x24, 0xa8, 0x74, 0x98, 0x64, 0x9b, 0x53, 0x29, 0x3e, 0xfe, 0xd5,
	0x86, 0xfe, 0x4d, 0xdd, 0x56, 0xf2, 0x1c, 0x1e, 0x36, 0x3d, 0x0e, 0xc4, 0xc2, 0x73, 0xc6, 0xce,
	0xab, 0xbe, 0xef, 0x36, 0xb9, 0xe9, 0x82, 0xbc, 0x80, 0xc1, 0xa6, 0x24, 0x0d, 0x13, 0xf4, 0x0e,
	0x6c, 0xd1, 0xa3, 0x26, 0x7b, 0x6d, 0x92, 0xe4, 0x08, 0x5c, 0xe3, 0x5a, 0x2d, 0x13, 0xcc, 0x0b,
	0x50, 0xcb, 0xd6, 0x40, 0x9d, 0x32, 0x9c, 0x33, 0x00, 0xe3, 0x23, 0xd7, 0x41, 0xe1, 0xc8, 0x3b,
	0x34, 0xf7, 0xee, 0x64, 0x44, 0xab, 0xd6, 0xd4, 0xfe, 0xe9, 0x97, 0xda, 0xae, 0xdf, 0xb7, 0xd5,
	0x45, 0x4c, 0xde, 0x40, 0x0f, 0xd3, 0x45, 0x29, 0x6c, 0xff, 0x55, 0xd8, 0x35, 0xb5, 0x56, 0x76,
	0x05, 0x9d, 0x38, 0x9c, 0x63, 0xac, 0xbc, 0xce, 0xb8, 0x65, 0x44, 0xa7, 0x74, 0xd7, 0xfa, 0xd0,
	0xa6, 0x2b, 0x74, 0x66, 0x55, 0x1f, 0x8b, 0x11, 0xfa, 0x15, 0x82, 0x7c, 0x87, 0x27, 0xdb, 0xa3,
	0x08, 0x14, 0x6a, 0xe5, 0x75, 0x2d, 0xf7, 0xf5, 0x6e, 0xee, 0x27, 0x2b, 0xfb, 0x5a, 0xa8, 0x3e,
	0xa3, 0xf6, 0x1f, 0x27, 0xf7, 0x62, 0x45, 0x2e, 0xc0, 0xad, 0x37, 0x46, 0xa0, 0xf2, 0x7a, 0x96,
	0xf9, 0x72, 0x37, 0x73, 0x26, 0x79, 0x69, 0x0f, 0xe2, 0xf2, 0x64, 0x94, 0xc4, 0x07, 0x10, 0x49,
	0x26, 0x73, 0x1d, 0xa6, 0x11, 0x7a, 0xae, 0x69, 0xd4, 0x60, 0x32, 0xf9, 0xd7, 0x7f, 0x9e, 0x36,
	0x4a, 0x7f, 0x8b, 0x32, 0x3a, 0x03, 0x77, 0xab, 0x1b, 0x64, 0x08, 0xad, 0x9f, 0xb8, 0xae, 0xd6,
	0xa4, 0x38, 0x92, 0xa7, 0xd0, 0xb6, 0x0d, 0xa9, 0xb6, 0xa2, 0x0c, 0xde, 0x1f, 0xbc, 0x73, 0x8e,
	0x8f, 0x00, 0x36, 0x50, 0xd2, 0x85, 0xd6, 0xec, 0xe6, 0xdb, 0xf0, 0x01, 0xe9, 0xc1, 0xe1, 0xe5,
	0xf4, 0xe2, 0x72, 0xe8, 0x9c, 0xbf, 0x85, 0xb1, 0x59, 0xe1, 0x9d, 0x06, 0xcf, 0x07, 0x8d, 0xc3,
	0xdb, 0x62, 0xd2, 0xb7, 0xce, 0x6f, 0xc7, 0x99, 0x77, 0xec, 0xd4, 0x4f, 0xff, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x89, 0xb4, 0x10, 0x3f, 0x5c, 0x04, 0x00, 0x00,
}
