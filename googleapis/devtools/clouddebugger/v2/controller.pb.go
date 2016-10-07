// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/devtools/clouddebugger/v2/controller.proto
// DO NOT EDIT!

/*
Package google_devtools_clouddebugger_v2 is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/genproto/googleapis/devtools/clouddebugger/v2/controller.proto
	google.golang.org/genproto/googleapis/devtools/clouddebugger/v2/data.proto
	google.golang.org/genproto/googleapis/devtools/clouddebugger/v2/debugger.proto

It has these top-level messages:
	RegisterDebuggeeRequest
	RegisterDebuggeeResponse
	ListActiveBreakpointsRequest
	ListActiveBreakpointsResponse
	UpdateActiveBreakpointRequest
	UpdateActiveBreakpointResponse
	FormatMessage
	StatusMessage
	SourceLocation
	Variable
	StackFrame
	Breakpoint
	Debuggee
	SetBreakpointRequest
	SetBreakpointResponse
	GetBreakpointRequest
	GetBreakpointResponse
	DeleteBreakpointRequest
	ListBreakpointsRequest
	ListBreakpointsResponse
	ListDebuggeesRequest
	ListDebuggeesResponse
*/
package google_devtools_clouddebugger_v2 // import "google.golang.org/genproto/googleapis/devtools/clouddebugger/v2"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import _ "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request to register a debuggee.
type RegisterDebuggeeRequest struct {
	// Debuggee information to register.
	// The fields `project`, `uniquifier`, `description` and `agent_version`
	// of the debuggee must be set.
	Debuggee *Debuggee `protobuf:"bytes,1,opt,name=debuggee" json:"debuggee,omitempty"`
}

func (m *RegisterDebuggeeRequest) Reset()                    { *m = RegisterDebuggeeRequest{} }
func (m *RegisterDebuggeeRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterDebuggeeRequest) ProtoMessage()               {}
func (*RegisterDebuggeeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RegisterDebuggeeRequest) GetDebuggee() *Debuggee {
	if m != nil {
		return m.Debuggee
	}
	return nil
}

// Response for registering a debuggee.
type RegisterDebuggeeResponse struct {
	// Debuggee resource.
	// The field `id` is guranteed to be set (in addition to the echoed fields).
	Debuggee *Debuggee `protobuf:"bytes,1,opt,name=debuggee" json:"debuggee,omitempty"`
}

func (m *RegisterDebuggeeResponse) Reset()                    { *m = RegisterDebuggeeResponse{} }
func (m *RegisterDebuggeeResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterDebuggeeResponse) ProtoMessage()               {}
func (*RegisterDebuggeeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegisterDebuggeeResponse) GetDebuggee() *Debuggee {
	if m != nil {
		return m.Debuggee
	}
	return nil
}

// Request to list active breakpoints.
type ListActiveBreakpointsRequest struct {
	// Identifies the debuggee.
	DebuggeeId string `protobuf:"bytes,1,opt,name=debuggee_id,json=debuggeeId" json:"debuggee_id,omitempty"`
	// A wait token that, if specified, blocks the method call until the list
	// of active breakpoints has changed, or a server selected timeout has
	// expired.  The value should be set from the last returned response.
	WaitToken string `protobuf:"bytes,2,opt,name=wait_token,json=waitToken" json:"wait_token,omitempty"`
	// If set to `true`, returns `google.rpc.Code.OK` status and sets the
	// `wait_expired` response field to `true` when the server-selected timeout
	// has expired (recommended).
	//
	// If set to `false`, returns `google.rpc.Code.ABORTED` status when the
	// server-selected timeout has expired (deprecated).
	SuccessOnTimeout bool `protobuf:"varint,3,opt,name=success_on_timeout,json=successOnTimeout" json:"success_on_timeout,omitempty"`
}

func (m *ListActiveBreakpointsRequest) Reset()                    { *m = ListActiveBreakpointsRequest{} }
func (m *ListActiveBreakpointsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListActiveBreakpointsRequest) ProtoMessage()               {}
func (*ListActiveBreakpointsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Response for listing active breakpoints.
type ListActiveBreakpointsResponse struct {
	// List of all active breakpoints.
	// The fields `id` and `location` are guaranteed to be set on each breakpoint.
	Breakpoints []*Breakpoint `protobuf:"bytes,1,rep,name=breakpoints" json:"breakpoints,omitempty"`
	// A wait token that can be used in the next method call to block until
	// the list of breakpoints changes.
	NextWaitToken string `protobuf:"bytes,2,opt,name=next_wait_token,json=nextWaitToken" json:"next_wait_token,omitempty"`
	// The `wait_expired` field is set to true by the server when the
	// request times out and the field `success_on_timeout` is set to true.
	WaitExpired bool `protobuf:"varint,3,opt,name=wait_expired,json=waitExpired" json:"wait_expired,omitempty"`
}

func (m *ListActiveBreakpointsResponse) Reset()                    { *m = ListActiveBreakpointsResponse{} }
func (m *ListActiveBreakpointsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListActiveBreakpointsResponse) ProtoMessage()               {}
func (*ListActiveBreakpointsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListActiveBreakpointsResponse) GetBreakpoints() []*Breakpoint {
	if m != nil {
		return m.Breakpoints
	}
	return nil
}

// Request to update an active breakpoint.
type UpdateActiveBreakpointRequest struct {
	// Identifies the debuggee being debugged.
	DebuggeeId string `protobuf:"bytes,1,opt,name=debuggee_id,json=debuggeeId" json:"debuggee_id,omitempty"`
	// Updated breakpoint information.
	// The field 'id' must be set.
	Breakpoint *Breakpoint `protobuf:"bytes,2,opt,name=breakpoint" json:"breakpoint,omitempty"`
}

func (m *UpdateActiveBreakpointRequest) Reset()                    { *m = UpdateActiveBreakpointRequest{} }
func (m *UpdateActiveBreakpointRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateActiveBreakpointRequest) ProtoMessage()               {}
func (*UpdateActiveBreakpointRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateActiveBreakpointRequest) GetBreakpoint() *Breakpoint {
	if m != nil {
		return m.Breakpoint
	}
	return nil
}

// Response for updating an active breakpoint.
// The message is defined to allow future extensions.
type UpdateActiveBreakpointResponse struct {
}

func (m *UpdateActiveBreakpointResponse) Reset()                    { *m = UpdateActiveBreakpointResponse{} }
func (m *UpdateActiveBreakpointResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateActiveBreakpointResponse) ProtoMessage()               {}
func (*UpdateActiveBreakpointResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*RegisterDebuggeeRequest)(nil), "google.devtools.clouddebugger.v2.RegisterDebuggeeRequest")
	proto.RegisterType((*RegisterDebuggeeResponse)(nil), "google.devtools.clouddebugger.v2.RegisterDebuggeeResponse")
	proto.RegisterType((*ListActiveBreakpointsRequest)(nil), "google.devtools.clouddebugger.v2.ListActiveBreakpointsRequest")
	proto.RegisterType((*ListActiveBreakpointsResponse)(nil), "google.devtools.clouddebugger.v2.ListActiveBreakpointsResponse")
	proto.RegisterType((*UpdateActiveBreakpointRequest)(nil), "google.devtools.clouddebugger.v2.UpdateActiveBreakpointRequest")
	proto.RegisterType((*UpdateActiveBreakpointResponse)(nil), "google.devtools.clouddebugger.v2.UpdateActiveBreakpointResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Controller2 service

type Controller2Client interface {
	// Registers the debuggee with the controller service.
	//
	// All agents attached to the same application should call this method with
	// the same request content to get back the same stable `debuggee_id`. Agents
	// should call this method again whenever `google.rpc.Code.NOT_FOUND` is
	// returned from any controller method.
	//
	// This allows the controller service to disable the agent or recover from any
	// data loss. If the debuggee is disabled by the server, the response will
	// have `is_disabled` set to `true`.
	RegisterDebuggee(ctx context.Context, in *RegisterDebuggeeRequest, opts ...grpc.CallOption) (*RegisterDebuggeeResponse, error)
	// Returns the list of all active breakpoints for the debuggee.
	//
	// The breakpoint specification (location, condition, and expression
	// fields) is semantically immutable, although the field values may
	// change. For example, an agent may update the location line number
	// to reflect the actual line where the breakpoint was set, but this
	// doesn't change the breakpoint semantics.
	//
	// This means that an agent does not need to check if a breakpoint has changed
	// when it encounters the same breakpoint on a successive call.
	// Moreover, an agent should remember the breakpoints that are completed
	// until the controller removes them from the active list to avoid
	// setting those breakpoints again.
	ListActiveBreakpoints(ctx context.Context, in *ListActiveBreakpointsRequest, opts ...grpc.CallOption) (*ListActiveBreakpointsResponse, error)
	// Updates the breakpoint state or mutable fields.
	// The entire Breakpoint message must be sent back to the controller
	// service.
	//
	// Updates to active breakpoint fields are only allowed if the new value
	// does not change the breakpoint specification. Updates to the `location`,
	// `condition` and `expression` fields should not alter the breakpoint
	// semantics. These may only make changes such as canonicalizing a value
	// or snapping the location to the correct line of code.
	UpdateActiveBreakpoint(ctx context.Context, in *UpdateActiveBreakpointRequest, opts ...grpc.CallOption) (*UpdateActiveBreakpointResponse, error)
}

type controller2Client struct {
	cc *grpc.ClientConn
}

func NewController2Client(cc *grpc.ClientConn) Controller2Client {
	return &controller2Client{cc}
}

func (c *controller2Client) RegisterDebuggee(ctx context.Context, in *RegisterDebuggeeRequest, opts ...grpc.CallOption) (*RegisterDebuggeeResponse, error) {
	out := new(RegisterDebuggeeResponse)
	err := grpc.Invoke(ctx, "/google.devtools.clouddebugger.v2.Controller2/RegisterDebuggee", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controller2Client) ListActiveBreakpoints(ctx context.Context, in *ListActiveBreakpointsRequest, opts ...grpc.CallOption) (*ListActiveBreakpointsResponse, error) {
	out := new(ListActiveBreakpointsResponse)
	err := grpc.Invoke(ctx, "/google.devtools.clouddebugger.v2.Controller2/ListActiveBreakpoints", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controller2Client) UpdateActiveBreakpoint(ctx context.Context, in *UpdateActiveBreakpointRequest, opts ...grpc.CallOption) (*UpdateActiveBreakpointResponse, error) {
	out := new(UpdateActiveBreakpointResponse)
	err := grpc.Invoke(ctx, "/google.devtools.clouddebugger.v2.Controller2/UpdateActiveBreakpoint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Controller2 service

type Controller2Server interface {
	// Registers the debuggee with the controller service.
	//
	// All agents attached to the same application should call this method with
	// the same request content to get back the same stable `debuggee_id`. Agents
	// should call this method again whenever `google.rpc.Code.NOT_FOUND` is
	// returned from any controller method.
	//
	// This allows the controller service to disable the agent or recover from any
	// data loss. If the debuggee is disabled by the server, the response will
	// have `is_disabled` set to `true`.
	RegisterDebuggee(context.Context, *RegisterDebuggeeRequest) (*RegisterDebuggeeResponse, error)
	// Returns the list of all active breakpoints for the debuggee.
	//
	// The breakpoint specification (location, condition, and expression
	// fields) is semantically immutable, although the field values may
	// change. For example, an agent may update the location line number
	// to reflect the actual line where the breakpoint was set, but this
	// doesn't change the breakpoint semantics.
	//
	// This means that an agent does not need to check if a breakpoint has changed
	// when it encounters the same breakpoint on a successive call.
	// Moreover, an agent should remember the breakpoints that are completed
	// until the controller removes them from the active list to avoid
	// setting those breakpoints again.
	ListActiveBreakpoints(context.Context, *ListActiveBreakpointsRequest) (*ListActiveBreakpointsResponse, error)
	// Updates the breakpoint state or mutable fields.
	// The entire Breakpoint message must be sent back to the controller
	// service.
	//
	// Updates to active breakpoint fields are only allowed if the new value
	// does not change the breakpoint specification. Updates to the `location`,
	// `condition` and `expression` fields should not alter the breakpoint
	// semantics. These may only make changes such as canonicalizing a value
	// or snapping the location to the correct line of code.
	UpdateActiveBreakpoint(context.Context, *UpdateActiveBreakpointRequest) (*UpdateActiveBreakpointResponse, error)
}

func RegisterController2Server(s *grpc.Server, srv Controller2Server) {
	s.RegisterService(&_Controller2_serviceDesc, srv)
}

func _Controller2_RegisterDebuggee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDebuggeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Controller2Server).RegisterDebuggee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouddebugger.v2.Controller2/RegisterDebuggee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Controller2Server).RegisterDebuggee(ctx, req.(*RegisterDebuggeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Controller2_ListActiveBreakpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListActiveBreakpointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Controller2Server).ListActiveBreakpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouddebugger.v2.Controller2/ListActiveBreakpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Controller2Server).ListActiveBreakpoints(ctx, req.(*ListActiveBreakpointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Controller2_UpdateActiveBreakpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateActiveBreakpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Controller2Server).UpdateActiveBreakpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouddebugger.v2.Controller2/UpdateActiveBreakpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Controller2Server).UpdateActiveBreakpoint(ctx, req.(*UpdateActiveBreakpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Controller2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.devtools.clouddebugger.v2.Controller2",
	HandlerType: (*Controller2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterDebuggee",
			Handler:    _Controller2_RegisterDebuggee_Handler,
		},
		{
			MethodName: "ListActiveBreakpoints",
			Handler:    _Controller2_ListActiveBreakpoints_Handler,
		},
		{
			MethodName: "UpdateActiveBreakpoint",
			Handler:    _Controller2_UpdateActiveBreakpoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/devtools/clouddebugger/v2/controller.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0xcb, 0x6e, 0xd3, 0x4c,
	0x14, 0x96, 0x5b, 0xfd, 0xbf, 0xda, 0x13, 0x50, 0xab, 0x91, 0x80, 0xc8, 0xa2, 0x90, 0x5a, 0xa8,
	0x42, 0x55, 0xe4, 0x91, 0x5c, 0x36, 0xcd, 0x82, 0x4b, 0xb8, 0xa9, 0xa8, 0x40, 0x14, 0x15, 0xb1,
	0x8c, 0x7c, 0x39, 0x1d, 0x46, 0x75, 0x66, 0x8c, 0x67, 0x1c, 0x8a, 0x2a, 0x36, 0x6c, 0x41, 0x6c,
	0x78, 0x18, 0x24, 0x5e, 0x03, 0x1e, 0x81, 0x15, 0x4f, 0xc1, 0xc4, 0x76, 0x2e, 0xbd, 0xa4, 0x49,
	0x03, 0x1b, 0x2b, 0x39, 0x67, 0xbe, 0xf3, 0x7d, 0xdf, 0x99, 0x33, 0x07, 0x5a, 0x4c, 0x4a, 0x16,
	0xa3, 0xcb, 0x64, 0xec, 0x0b, 0xe6, 0xca, 0x94, 0x51, 0x86, 0x22, 0x49, 0xa5, 0x96, 0xb4, 0x48,
	0xf9, 0x09, 0x57, 0x34, 0xc2, 0x9e, 0x96, 0x32, 0x56, 0x34, 0x8c, 0x65, 0x16, 0x45, 0x18, 0x64,
	0x8c, 0x61, 0x4a, 0x7b, 0x1e, 0x0d, 0xa5, 0xd0, 0xa9, 0x8c, 0x63, 0x4c, 0xdd, 0x1c, 0x45, 0x6a,
	0x65, 0xc5, 0x01, 0xc4, 0x3d, 0x06, 0x71, 0x7b, 0x9e, 0xbd, 0x33, 0x1b, 0xa7, 0xf9, 0x50, 0x85,
	0x69, 0x8f, 0x87, 0x68, 0x48, 0xf6, 0x39, 0xa3, 0xbe, 0x10, 0x52, 0xfb, 0x9a, 0x4b, 0xa1, 0x0a,
	0x32, 0xfb, 0xd9, 0xdf, 0xca, 0x8f, 0x7c, 0xed, 0x97, 0xb5, 0xb6, 0x18, 0xd7, 0x6f, 0xb2, 0xc0,
	0x0d, 0x65, 0x97, 0x16, 0xf5, 0x68, 0x9e, 0x08, 0xb2, 0x7d, 0x9a, 0xe8, 0xf7, 0x09, 0x2a, 0x8a,
	0x5d, 0xf3, 0xa3, 0xf8, 0x16, 0x20, 0xc7, 0x87, 0x6b, 0x6d, 0x64, 0x5c, 0x69, 0x4c, 0x1f, 0x15,
	0x65, 0xb1, 0x8d, 0x6f, 0x33, 0x54, 0x9a, 0x3c, 0x81, 0xa5, 0x92, 0x09, 0xab, 0x56, 0xcd, 0xba,
	0x5d, 0xf1, 0x36, 0xdd, 0x69, 0xbd, 0x71, 0x87, 0x45, 0x86, 0x58, 0x27, 0x80, 0xea, 0x69, 0x0a,
	0x95, 0x98, 0x26, 0xe0, 0x3f, 0xe3, 0xf8, 0x6c, 0xc1, 0xf5, 0x5d, 0x43, 0xf1, 0x20, 0xd4, 0xbc,
	0x87, 0xcd, 0x14, 0xfd, 0x83, 0x44, 0x72, 0xa1, 0xd5, 0xc0, 0xcc, 0x4d, 0xa8, 0x0c, 0x0e, 0x77,
	0x78, 0x94, 0x73, 0x2d, 0xb7, 0x61, 0x10, 0xda, 0x89, 0xc8, 0x1a, 0xc0, 0x3b, 0x9f, 0xeb, 0x8e,
	0x96, 0x07, 0x28, 0xaa, 0x0b, 0x79, 0x7e, 0xb9, 0x1f, 0xd9, 0xeb, 0x07, 0x48, 0x1d, 0x88, 0xca,
	0xc2, 0x10, 0x95, 0xea, 0x48, 0xd1, 0xd1, 0xbc, 0x8b, 0x32, 0xd3, 0xd5, 0x45, 0x73, 0x6c, 0xa9,
	0xbd, 0x5a, 0x66, 0x5e, 0x8a, 0xbd, 0x22, 0xee, 0x7c, 0xb7, 0x60, 0x6d, 0x82, 0x9c, 0xd2, 0xf8,
	0x0b, 0xa8, 0x04, 0xa3, 0xb0, 0xd1, 0xb3, 0x68, 0xbc, 0xd7, 0xa7, 0x7b, 0x1f, 0xd5, 0x6a, 0x8f,
	0x17, 0x20, 0x1b, 0xb0, 0x22, 0xf0, 0x50, 0x77, 0x4e, 0x79, 0xb8, 0xdc, 0x0f, 0xbf, 0x1e, 0xfa,
	0x58, 0x87, 0x4b, 0xf9, 0x11, 0x3c, 0x4c, 0x78, 0x8a, 0x51, 0xe9, 0xa0, 0xd2, 0x8f, 0x3d, 0x2e,
	0x42, 0xce, 0x17, 0x23, 0xfe, 0x55, 0x62, 0x06, 0x0b, 0x4f, 0xca, 0x9f, 0xb9, 0x99, 0xbb, 0x00,
	0x23, 0x71, 0xb9, 0x90, 0x8b, 0x9a, 0x1b, 0xc3, 0x3b, 0x35, 0xb8, 0x31, 0x49, 0x4f, 0xd1, 0x4d,
	0xef, 0xd3, 0x7f, 0x50, 0x79, 0x38, 0x7c, 0xc8, 0x1e, 0xf9, 0x66, 0xc1, 0xea, 0xc9, 0x99, 0x23,
	0xdb, 0xd3, 0x05, 0x4c, 0x78, 0x0a, 0x76, 0x63, 0x1e, 0x68, 0xa1, 0xcd, 0xa9, 0x7f, 0xfc, 0xf1,
	0xeb, 0xeb, 0xc2, 0x86, 0xb3, 0x7e, 0x7c, 0xdb, 0xd0, 0x41, 0xbb, 0x14, 0x4d, 0x4b, 0x68, 0xc3,
	0xda, 0x24, 0x3f, 0x2d, 0xb8, 0x72, 0xe6, 0xe4, 0x90, 0xbb, 0xd3, 0x35, 0x9c, 0xf7, 0x02, 0xec,
	0x7b, 0x73, 0xe3, 0x4b, 0x23, 0x8d, 0xdc, 0xc8, 0x1d, 0xe2, 0x4d, 0x34, 0x72, 0x34, 0x36, 0x15,
	0x1f, 0xe8, 0xf8, 0x78, 0xfe, 0xb6, 0xe0, 0xea, 0xd9, 0x77, 0x48, 0x66, 0xd0, 0x75, 0xee, 0x34,
	0xda, 0xf7, 0xe7, 0x2f, 0x50, 0x3a, 0x7b, 0x9e, 0x3b, 0x7b, 0x6a, 0x37, 0x2f, 0xee, 0x8c, 0x1e,
	0x8d, 0xfe, 0xb8, 0x26, 0x67, 0xee, 0xb0, 0xb9, 0x0d, 0xb7, 0xcc, 0x0e, 0x9e, 0xaa, 0xaa, 0xb9,
	0x32, 0x1a, 0xd9, 0x56, 0x7f, 0x19, 0xb7, 0xac, 0xe0, 0xff, 0x7c, 0x2b, 0x6f, 0xfd, 0x09, 0x00,
	0x00, 0xff, 0xff, 0x89, 0xdc, 0x57, 0xe9, 0xd7, 0x06, 0x00, 0x00,
}
