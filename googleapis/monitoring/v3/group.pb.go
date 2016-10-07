// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/monitoring/v3/group.proto
// DO NOT EDIT!

package google_monitoring_v3 // import "google.golang.org/genproto/googleapis/monitoring/v3"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The description of a dynamic collection of monitored resources. Each group
// has a filter that is matched against monitored resources and their associated
// metadata. If a group's filter matches an available monitored resource, then
// that resource is a member of that group.  Groups can contain any number of
// monitored resources, and each monitored resource can be a member of any
// number of groups.
//
// Groups can be nested in parent-child hierarchies. The `parentName` field
// identifies an optional parent for each group.  If a group has a parent, then
// the only monitored resources available to be matched by the group's filter
// are the resources contained in the parent group.  In other words, a group
// contains the monitored resources that match its filter and the filters of all
// the group's ancestors.  A group without a parent can contain any monitored
// resource.
//
// For example, consider an infrastructure running a set of instances with two
// user-defined tags: `"environment"` and `"role"`. A parent group has a filter,
// `environment="production"`.  A child of that parent group has a filter,
// `role="transcoder"`.  The parent group contains all instances in the
// production environment, regardless of their roles.  The child group contains
// instances that have the transcoder role *and* are in the production
// environment.
//
// The monitored resources contained in a group can change at any moment,
// depending on what resources exist and what filters are associated with the
// group and its ancestors.
type Group struct {
	// Output only. The name of this group. The format is
	// `"projects/{project_id_or_number}/groups/{group_id}"`.
	// When creating a group, this field is ignored and a new name is created
	// consisting of the project specified in the call to `CreateGroup`
	// and a unique `{group_id}` that is generated automatically.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// A user-assigned name for this group, used only for display purposes.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	// The name of the group's parent, if it has one.
	// The format is `"projects/{project_id_or_number}/groups/{group_id}"`.
	// For groups with no parent, `parentName` is the empty string, `""`.
	ParentName string `protobuf:"bytes,3,opt,name=parent_name,json=parentName" json:"parent_name,omitempty"`
	// The filter used to determine which monitored resources belong to this group.
	Filter string `protobuf:"bytes,5,opt,name=filter" json:"filter,omitempty"`
	// If true, the members of this group are considered to be a cluster.
	// The system can perform additional analysis on groups that are clusters.
	IsCluster bool `protobuf:"varint,6,opt,name=is_cluster,json=isCluster" json:"is_cluster,omitempty"`
}

func (m *Group) Reset()                    { *m = Group{} }
func (m *Group) String() string            { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()               {}
func (*Group) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func init() {
	proto.RegisterType((*Group)(nil), "google.monitoring.v3.Group")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/monitoring/v3/group.proto", fileDescriptor3)
}

var fileDescriptor3 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x8f, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x89, 0xda, 0xc1, 0xde, 0xba, 0x0a, 0x22, 0xb3, 0x11, 0x7f, 0x56, 0x5d, 0x25, 0x8b,
	0x79, 0x00, 0xa1, 0x2e, 0xdc, 0x49, 0xe9, 0x0b, 0x94, 0x58, 0xe3, 0x25, 0x90, 0xc9, 0x0d, 0x49,
	0x66, 0xc0, 0x17, 0xf1, 0x79, 0x35, 0x37, 0x03, 0x6e, 0xba, 0x4b, 0xce, 0xf7, 0x71, 0x39, 0x07,
	0x5e, 0x90, 0x08, 0xbd, 0x55, 0x48, 0xde, 0x04, 0x54, 0x94, 0x50, 0xa3, 0x0d, 0x31, 0x51, 0x21,
	0xdd, 0x90, 0x89, 0x2e, 0xeb, 0x91, 0x82, 0x2b, 0x94, 0x5c, 0x40, 0x3d, 0x0f, 0x1a, 0x13, 0x4d,
	0x51, 0xb1, 0x24, 0x6f, 0x97, 0x03, 0xff, 0x86, 0x9a, 0x87, 0xe7, 0x1f, 0x01, 0xab, 0xb7, 0x6a,
	0x49, 0x09, 0x57, 0xc1, 0x8c, 0xb6, 0x17, 0x8f, 0x62, 0xbb, 0x3e, 0xf0, 0x5b, 0x3e, 0xc1, 0xcd,
	0xa7, 0xcb, 0xd1, 0x9b, 0xef, 0x23, 0xb3, 0x0b, 0x66, 0x9b, 0x25, 0x7b, 0xaf, 0xca, 0x03, 0x6c,
	0xa2, 0x49, 0x36, 0x94, 0x66, 0x5c, 0xb2, 0x01, 0x2d, 0x62, 0xe1, 0x0e, 0xba, 0x2f, 0xe7, 0x8b,
	0x4d, 0xfd, 0x8a, 0xd9, 0xf2, 0x93, 0xf7, 0x00, 0x2e, 0x1f, 0x4f, 0x7e, 0xca, 0x95, 0x75, 0x7f,
	0xec, 0xfa, 0xb0, 0x76, 0xf9, 0xb5, 0x05, 0xbb, 0x2d, 0xf4, 0x27, 0x1a, 0xd5, 0xb9, 0xd2, 0x3b,
	0xe0, 0xc6, 0xfb, 0x3a, 0x6b, 0x2f, 0x3e, 0x3a, 0xde, 0x37, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x76, 0x8a, 0xec, 0x63, 0x22, 0x01, 0x00, 0x00,
}
