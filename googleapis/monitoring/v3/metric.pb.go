// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/monitoring/v3/metric.proto
// DO NOT EDIT!

package google_monitoring_v3 // import "google.golang.org/genproto/googleapis/monitoring/v3"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_api5 "google.golang.org/genproto/googleapis/api/metric"
import google_api4 "google.golang.org/genproto/googleapis/api/monitoredres"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A single data point in a time series.
type Point struct {
	// The time interval to which the data point applies.  For GAUGE metrics, only
	// the end time of the interval is used.  For DELTA metrics, the start and end
	// time should specify a non-zero interval, with subsequent points specifying
	// contiguous and non-overlapping intervals.  For CUMULATIVE metrics, the
	// start and end time should specify a non-zero interval, with subsequent
	// points specifying the same start time and increasing end times, until an
	// event resets the cumulative value to zero and sets a new start time for the
	// following points.
	Interval *TimeInterval `protobuf:"bytes,1,opt,name=interval" json:"interval,omitempty"`
	// The value of the data point.
	Value *TypedValue `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *Point) GetInterval() *TimeInterval {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *Point) GetValue() *TypedValue {
	if m != nil {
		return m.Value
	}
	return nil
}

// A collection of data points that describes the time-varying values
// of a metric. A time series is identified by a combination of a
// fully-specified monitored resource and a fully-specified metric.
// This type is used for both listing and creating time series.
type TimeSeries struct {
	// The associated metric. A fully-specified metric used to identify the time
	// series.
	Metric *google_api5.Metric `protobuf:"bytes,1,opt,name=metric" json:"metric,omitempty"`
	// The associated resource. A fully-specified monitored resource used to
	// identify the time series.
	Resource *google_api4.MonitoredResource `protobuf:"bytes,2,opt,name=resource" json:"resource,omitempty"`
	// The metric kind of the time series. When listing time series, this metric
	// kind might be different from the metric kind of the associated metric if
	// this time series is an alignment or reduction of other time series.
	//
	// When creating a time series, this field is optional. If present, it must be
	// the same as the metric kind of the associated metric. If the associated
	// metric's descriptor must be auto-created, then this field specifies the
	// metric kind of the new descriptor and must be either `GAUGE` (the default)
	// or `CUMULATIVE`.
	MetricKind google_api5.MetricDescriptor_MetricKind `protobuf:"varint,3,opt,name=metric_kind,json=metricKind,enum=google.api.MetricDescriptor_MetricKind" json:"metric_kind,omitempty"`
	// The value type of the time series. When listing time series, this value
	// type might be different from the value type of the associated metric if
	// this time series is an alignment or reduction of other time series.
	//
	// When creating a time series, this field is optional. If present, it must be
	// the same as the type of the data in the `points` field.
	ValueType google_api5.MetricDescriptor_ValueType `protobuf:"varint,4,opt,name=value_type,json=valueType,enum=google.api.MetricDescriptor_ValueType" json:"value_type,omitempty"`
	// The data points of this time series. When listing time series, the order of
	// the points is specified by the list method.
	//
	// When creating a time series, this field must contain exactly one point and
	// the point's type must be the same as the value type of the associated
	// metric. If the associated metric's descriptor must be auto-created, then
	// the value type of the descriptor is determined by the point's type, which
	// must be `BOOL`, `INT64`, `DOUBLE`, or `DISTRIBUTION`.
	Points []*Point `protobuf:"bytes,5,rep,name=points" json:"points,omitempty"`
}

func (m *TimeSeries) Reset()                    { *m = TimeSeries{} }
func (m *TimeSeries) String() string            { return proto.CompactTextString(m) }
func (*TimeSeries) ProtoMessage()               {}
func (*TimeSeries) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *TimeSeries) GetMetric() *google_api5.Metric {
	if m != nil {
		return m.Metric
	}
	return nil
}

func (m *TimeSeries) GetResource() *google_api4.MonitoredResource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *TimeSeries) GetPoints() []*Point {
	if m != nil {
		return m.Points
	}
	return nil
}

func init() {
	proto.RegisterType((*Point)(nil), "google.monitoring.v3.Point")
	proto.RegisterType((*TimeSeries)(nil), "google.monitoring.v3.TimeSeries")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/monitoring/v3/metric.proto", fileDescriptor5)
}

var fileDescriptor5 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x91, 0x4f, 0x4b, 0xc3, 0x30,
	0x14, 0xc0, 0xd9, 0xe6, 0xc6, 0x7c, 0x03, 0x0f, 0xc1, 0x43, 0x99, 0x08, 0x63, 0x07, 0xff, 0x1d,
	0x52, 0x58, 0x41, 0xf0, 0xa0, 0xc8, 0x50, 0x50, 0x44, 0x1c, 0x51, 0xbc, 0x8e, 0xd9, 0x3d, 0x42,
	0x70, 0x4d, 0x4a, 0xda, 0x15, 0x3c, 0xf9, 0xb9, 0xfc, 0x76, 0xa6, 0x49, 0xba, 0x31, 0xac, 0xa2,
	0x5e, 0x9a, 0x7f, 0xbf, 0xf7, 0x7b, 0x7d, 0xef, 0xc1, 0x25, 0x57, 0x8a, 0x2f, 0x90, 0x72, 0xb5,
	0x98, 0x49, 0x4e, 0x95, 0xe6, 0x21, 0x47, 0x99, 0x6a, 0x95, 0xab, 0xd0, 0x3d, 0xcd, 0x52, 0x91,
	0x85, 0x89, 0x92, 0x22, 0x57, 0x5a, 0x48, 0x1e, 0x16, 0x51, 0x98, 0x60, 0xae, 0x45, 0x4c, 0x2d,
	0x45, 0x76, 0xbd, 0x61, 0x8d, 0xd0, 0x22, 0xea, 0x9f, 0xff, 0xce, 0x6b, 0x3e, 0xde, 0xb6, 0x21,
	0xed, 0x3f, 0xfc, 0x21, 0xdc, 0xe5, 0xc5, 0xb9, 0xc6, 0x6c, 0x7d, 0x98, 0x9a, 0x93, 0x5a, 0xea,
	0x18, 0xbd, 0xf0, 0x5f, 0x75, 0xc6, 0x2a, 0x31, 0x17, 0xce, 0x30, 0x7c, 0x87, 0xf6, 0x44, 0x09,
	0x99, 0x93, 0x0b, 0xe8, 0x9a, 0x05, 0x75, 0x31, 0x5b, 0x04, 0x8d, 0x41, 0xe3, 0xa8, 0x37, 0x1a,
	0xd2, 0xba, 0x1e, 0xd0, 0x27, 0x91, 0xe0, 0xad, 0x27, 0xd9, 0x2a, 0x86, 0x9c, 0x42, 0xdb, 0x2c,
	0x4b, 0x0c, 0x9a, 0x36, 0x78, 0xf0, 0x4d, 0xf0, 0x5b, 0x8a, 0xf3, 0xe7, 0x92, 0x63, 0x0e, 0x1f,
	0x7e, 0x34, 0x01, 0x4a, 0xe5, 0x23, 0x6a, 0x81, 0x19, 0x39, 0x81, 0x8e, 0x6b, 0x99, 0xff, 0x09,
	0x52, 0x79, 0x4c, 0x21, 0xf4, 0xde, 0xbe, 0x30, 0x4f, 0x90, 0x33, 0xe8, 0x56, 0xfd, 0xf0, 0x59,
	0xf7, 0x37, 0xe8, 0xaa, 0x6b, 0xcc, 0x43, 0x6c, 0x85, 0x93, 0x1b, 0xe8, 0x39, 0xc9, 0xf4, 0x55,
	0xc8, 0x79, 0xd0, 0x32, 0xd1, 0x3b, 0xa3, 0xc3, 0xaf, 0xb9, 0xae, 0x30, 0x8b, 0xb5, 0x48, 0x8d,
	0xc6, 0x5f, 0xdc, 0x19, 0x9c, 0x41, 0xb2, 0xda, 0x93, 0x6b, 0x00, 0x5b, 0xc8, 0x34, 0x37, 0xa5,
	0x05, 0x5b, 0x56, 0x74, 0xf0, 0xa3, 0xc8, 0x96, 0x5f, 0x36, 0x82, 0x6d, 0x17, 0xd5, 0x96, 0x44,
	0xd0, 0x49, 0xcb, 0x39, 0x64, 0x41, 0x7b, 0xd0, 0x32, 0x95, 0xec, 0xd5, 0xf7, 0xcf, 0xce, 0x8a,
	0x79, 0x74, 0x7c, 0x0c, 0x81, 0x19, 0x66, 0x2d, 0x39, 0xee, 0xb9, 0xbc, 0x93, 0x72, 0xca, 0x93,
	0xc6, 0x4b, 0xc7, 0x8e, 0x3b, 0xfa, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x42, 0x89, 0xb9, 0x1f, 0x1a,
	0x03, 0x00, 0x00,
}
