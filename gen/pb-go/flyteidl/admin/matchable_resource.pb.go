// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/matchable_resource.proto

package admin

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Defines a resource that can be configured by customizable Project-, ProjectDomain- or WorkflowAttributes
// based on matching tags.
type MatchableResource int32

const (
	// Applies to customizable task resource requests and limits.
	MatchableResource_TASK_RESOURCE MatchableResource = 0
	// Applies to configuring templated kubernetes cluster resources.
	MatchableResource_CLUSTER_RESOURCE MatchableResource = 1
	// Configures task and dynamic task execution queue assignment.
	MatchableResource_EXECUTION_QUEUE MatchableResource = 2
	// Configures the K8s cluster for the execution to be run
	MatchableResource_CLUSTER_LABEL MatchableResource = 3
)

var MatchableResource_name = map[int32]string{
	0: "TASK_RESOURCE",
	1: "CLUSTER_RESOURCE",
	2: "EXECUTION_QUEUE",
	3: "CLUSTER_LABEL",
}

var MatchableResource_value = map[string]int32{
	"TASK_RESOURCE":    0,
	"CLUSTER_RESOURCE": 1,
	"EXECUTION_QUEUE":  2,
	"CLUSTER_LABEL":    3,
}

func (x MatchableResource) String() string {
	return proto.EnumName(MatchableResource_name, int32(x))
}

func (MatchableResource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{0}
}

type TaskResourceSpec struct {
	Cpu                  string   `protobuf:"bytes,1,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Gpu                  string   `protobuf:"bytes,2,opt,name=gpu,proto3" json:"gpu,omitempty"`
	Memory               string   `protobuf:"bytes,3,opt,name=memory,proto3" json:"memory,omitempty"`
	Storage              string   `protobuf:"bytes,4,opt,name=storage,proto3" json:"storage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskResourceSpec) Reset()         { *m = TaskResourceSpec{} }
func (m *TaskResourceSpec) String() string { return proto.CompactTextString(m) }
func (*TaskResourceSpec) ProtoMessage()    {}
func (*TaskResourceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{0}
}

func (m *TaskResourceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskResourceSpec.Unmarshal(m, b)
}
func (m *TaskResourceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskResourceSpec.Marshal(b, m, deterministic)
}
func (m *TaskResourceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskResourceSpec.Merge(m, src)
}
func (m *TaskResourceSpec) XXX_Size() int {
	return xxx_messageInfo_TaskResourceSpec.Size(m)
}
func (m *TaskResourceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskResourceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_TaskResourceSpec proto.InternalMessageInfo

func (m *TaskResourceSpec) GetCpu() string {
	if m != nil {
		return m.Cpu
	}
	return ""
}

func (m *TaskResourceSpec) GetGpu() string {
	if m != nil {
		return m.Gpu
	}
	return ""
}

func (m *TaskResourceSpec) GetMemory() string {
	if m != nil {
		return m.Memory
	}
	return ""
}

func (m *TaskResourceSpec) GetStorage() string {
	if m != nil {
		return m.Storage
	}
	return ""
}

type TaskResourceAttributes struct {
	Defaults             *TaskResourceSpec `protobuf:"bytes,1,opt,name=defaults,proto3" json:"defaults,omitempty"`
	Limits               *TaskResourceSpec `protobuf:"bytes,2,opt,name=limits,proto3" json:"limits,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TaskResourceAttributes) Reset()         { *m = TaskResourceAttributes{} }
func (m *TaskResourceAttributes) String() string { return proto.CompactTextString(m) }
func (*TaskResourceAttributes) ProtoMessage()    {}
func (*TaskResourceAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{1}
}

func (m *TaskResourceAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskResourceAttributes.Unmarshal(m, b)
}
func (m *TaskResourceAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskResourceAttributes.Marshal(b, m, deterministic)
}
func (m *TaskResourceAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskResourceAttributes.Merge(m, src)
}
func (m *TaskResourceAttributes) XXX_Size() int {
	return xxx_messageInfo_TaskResourceAttributes.Size(m)
}
func (m *TaskResourceAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskResourceAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_TaskResourceAttributes proto.InternalMessageInfo

func (m *TaskResourceAttributes) GetDefaults() *TaskResourceSpec {
	if m != nil {
		return m.Defaults
	}
	return nil
}

func (m *TaskResourceAttributes) GetLimits() *TaskResourceSpec {
	if m != nil {
		return m.Limits
	}
	return nil
}

type ClusterResourceAttributes struct {
	// Custom resource attributes which will be applied in cluster resource creation (e.g. quotas).
	// Map keys are the *case-sensitive* names of variables in templatized resource files.
	// Map values should be the custom values which get substituted during resource creation.
	Attributes           map[string]string `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ClusterResourceAttributes) Reset()         { *m = ClusterResourceAttributes{} }
func (m *ClusterResourceAttributes) String() string { return proto.CompactTextString(m) }
func (*ClusterResourceAttributes) ProtoMessage()    {}
func (*ClusterResourceAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{2}
}

func (m *ClusterResourceAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterResourceAttributes.Unmarshal(m, b)
}
func (m *ClusterResourceAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterResourceAttributes.Marshal(b, m, deterministic)
}
func (m *ClusterResourceAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterResourceAttributes.Merge(m, src)
}
func (m *ClusterResourceAttributes) XXX_Size() int {
	return xxx_messageInfo_ClusterResourceAttributes.Size(m)
}
func (m *ClusterResourceAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterResourceAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterResourceAttributes proto.InternalMessageInfo

func (m *ClusterResourceAttributes) GetAttributes() map[string]string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type ExecutionQueueAttributes struct {
	// Tags used for assigning execution queues for tasks defined within this project.
	Tags                 []string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecutionQueueAttributes) Reset()         { *m = ExecutionQueueAttributes{} }
func (m *ExecutionQueueAttributes) String() string { return proto.CompactTextString(m) }
func (*ExecutionQueueAttributes) ProtoMessage()    {}
func (*ExecutionQueueAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{3}
}

func (m *ExecutionQueueAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionQueueAttributes.Unmarshal(m, b)
}
func (m *ExecutionQueueAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionQueueAttributes.Marshal(b, m, deterministic)
}
func (m *ExecutionQueueAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionQueueAttributes.Merge(m, src)
}
func (m *ExecutionQueueAttributes) XXX_Size() int {
	return xxx_messageInfo_ExecutionQueueAttributes.Size(m)
}
func (m *ExecutionQueueAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionQueueAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionQueueAttributes proto.InternalMessageInfo

func (m *ExecutionQueueAttributes) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// Generic container for encapsulating all types of the above attributes messages.
type MatchingAttributes struct {
	// Types that are valid to be assigned to Target:
	//	*MatchingAttributes_TaskResourceAttributes
	//	*MatchingAttributes_ClusterResourceAttributes
	//	*MatchingAttributes_ExecutionQueueAttributes
	//	*MatchingAttributes_ClusterLabel
	Target               isMatchingAttributes_Target `protobuf_oneof:"target"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *MatchingAttributes) Reset()         { *m = MatchingAttributes{} }
func (m *MatchingAttributes) String() string { return proto.CompactTextString(m) }
func (*MatchingAttributes) ProtoMessage()    {}
func (*MatchingAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{4}
}

func (m *MatchingAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchingAttributes.Unmarshal(m, b)
}
func (m *MatchingAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchingAttributes.Marshal(b, m, deterministic)
}
func (m *MatchingAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchingAttributes.Merge(m, src)
}
func (m *MatchingAttributes) XXX_Size() int {
	return xxx_messageInfo_MatchingAttributes.Size(m)
}
func (m *MatchingAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchingAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_MatchingAttributes proto.InternalMessageInfo

type isMatchingAttributes_Target interface {
	isMatchingAttributes_Target()
}

type MatchingAttributes_TaskResourceAttributes struct {
	TaskResourceAttributes *TaskResourceAttributes `protobuf:"bytes,1,opt,name=task_resource_attributes,json=taskResourceAttributes,proto3,oneof"`
}

type MatchingAttributes_ClusterResourceAttributes struct {
	ClusterResourceAttributes *ClusterResourceAttributes `protobuf:"bytes,2,opt,name=cluster_resource_attributes,json=clusterResourceAttributes,proto3,oneof"`
}

type MatchingAttributes_ExecutionQueueAttributes struct {
	ExecutionQueueAttributes *ExecutionQueueAttributes `protobuf:"bytes,3,opt,name=execution_queue_attributes,json=executionQueueAttributes,proto3,oneof"`
}

type MatchingAttributes_ClusterLabel struct {
	ClusterLabel string `protobuf:"bytes,4,opt,name=cluster_label,json=clusterLabel,proto3,oneof"`
}

func (*MatchingAttributes_TaskResourceAttributes) isMatchingAttributes_Target() {}

func (*MatchingAttributes_ClusterResourceAttributes) isMatchingAttributes_Target() {}

func (*MatchingAttributes_ExecutionQueueAttributes) isMatchingAttributes_Target() {}

func (*MatchingAttributes_ClusterLabel) isMatchingAttributes_Target() {}

func (m *MatchingAttributes) GetTarget() isMatchingAttributes_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *MatchingAttributes) GetTaskResourceAttributes() *TaskResourceAttributes {
	if x, ok := m.GetTarget().(*MatchingAttributes_TaskResourceAttributes); ok {
		return x.TaskResourceAttributes
	}
	return nil
}

func (m *MatchingAttributes) GetClusterResourceAttributes() *ClusterResourceAttributes {
	if x, ok := m.GetTarget().(*MatchingAttributes_ClusterResourceAttributes); ok {
		return x.ClusterResourceAttributes
	}
	return nil
}

func (m *MatchingAttributes) GetExecutionQueueAttributes() *ExecutionQueueAttributes {
	if x, ok := m.GetTarget().(*MatchingAttributes_ExecutionQueueAttributes); ok {
		return x.ExecutionQueueAttributes
	}
	return nil
}

func (m *MatchingAttributes) GetClusterLabel() string {
	if x, ok := m.GetTarget().(*MatchingAttributes_ClusterLabel); ok {
		return x.ClusterLabel
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MatchingAttributes) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MatchingAttributes_TaskResourceAttributes)(nil),
		(*MatchingAttributes_ClusterResourceAttributes)(nil),
		(*MatchingAttributes_ExecutionQueueAttributes)(nil),
		(*MatchingAttributes_ClusterLabel)(nil),
	}
}

func init() {
	proto.RegisterEnum("flyteidl.admin.MatchableResource", MatchableResource_name, MatchableResource_value)
	proto.RegisterType((*TaskResourceSpec)(nil), "flyteidl.admin.TaskResourceSpec")
	proto.RegisterType((*TaskResourceAttributes)(nil), "flyteidl.admin.TaskResourceAttributes")
	proto.RegisterType((*ClusterResourceAttributes)(nil), "flyteidl.admin.ClusterResourceAttributes")
	proto.RegisterMapType((map[string]string)(nil), "flyteidl.admin.ClusterResourceAttributes.AttributesEntry")
	proto.RegisterType((*ExecutionQueueAttributes)(nil), "flyteidl.admin.ExecutionQueueAttributes")
	proto.RegisterType((*MatchingAttributes)(nil), "flyteidl.admin.MatchingAttributes")
}

func init() {
	proto.RegisterFile("flyteidl/admin/matchable_resource.proto", fileDescriptor_1d15bcabb02640f4)
}

var fileDescriptor_1d15bcabb02640f4 = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x51, 0x8f, 0xd2, 0x4c,
	0x14, 0xa5, 0x74, 0x3f, 0xbe, 0xdd, 0xbb, 0xae, 0xdb, 0x1d, 0x37, 0xa4, 0xbb, 0xbe, 0x10, 0x12,
	0x15, 0x4d, 0x6c, 0xe3, 0xee, 0xcb, 0x6a, 0xf4, 0x01, 0x48, 0x13, 0x8c, 0xac, 0x9b, 0x1d, 0x20,
	0x51, 0x5f, 0xc8, 0xb4, 0x5c, 0x4a, 0xc3, 0x94, 0xd6, 0x76, 0xc6, 0xd8, 0x37, 0x7f, 0x82, 0x7f,
	0xc6, 0xff, 0x67, 0xda, 0x52, 0x28, 0x04, 0xcc, 0xbe, 0xcd, 0x3d, 0x3d, 0x77, 0xce, 0x99, 0x33,
	0xd3, 0x0b, 0x2f, 0xa6, 0x3c, 0x11, 0xe8, 0x4d, 0xb8, 0xc9, 0x26, 0xbe, 0xb7, 0x30, 0x7d, 0x26,
	0x9c, 0x19, 0xb3, 0x39, 0x8e, 0x23, 0x8c, 0x03, 0x19, 0x39, 0x68, 0x84, 0x51, 0x20, 0x02, 0xf2,
	0xb8, 0x20, 0x1a, 0x19, 0xb1, 0x39, 0x03, 0x6d, 0xc8, 0xe2, 0x39, 0x5d, 0xb2, 0x06, 0x21, 0x3a,
	0x44, 0x03, 0xd5, 0x09, 0xa5, 0xae, 0x34, 0x94, 0xd6, 0x11, 0x4d, 0x97, 0x29, 0xe2, 0x86, 0x52,
	0xaf, 0xe6, 0x88, 0x1b, 0x4a, 0x52, 0x87, 0x9a, 0x8f, 0x7e, 0x10, 0x25, 0xba, 0x9a, 0x81, 0xcb,
	0x8a, 0xe8, 0xf0, 0x7f, 0x2c, 0x82, 0x88, 0xb9, 0xa8, 0x1f, 0x64, 0x1f, 0x8a, 0xb2, 0xf9, 0x5b,
	0x81, 0x7a, 0x59, 0xaa, 0x2d, 0x44, 0xe4, 0xd9, 0x52, 0x60, 0x4c, 0xde, 0xc3, 0xe1, 0x04, 0xa7,
	0x4c, 0x72, 0x11, 0x67, 0xaa, 0xc7, 0x57, 0x0d, 0x63, 0xd3, 0xa7, 0xb1, 0x6d, 0x92, 0xae, 0x3a,
	0xc8, 0x0d, 0xd4, 0xb8, 0xe7, 0x7b, 0x22, 0xce, 0xfc, 0x3d, 0xa4, 0x77, 0xc9, 0x6f, 0xfe, 0x51,
	0xe0, 0xa2, 0xcb, 0x65, 0x2c, 0x30, 0xda, 0xe1, 0xea, 0x2b, 0x00, 0x5b, 0x55, 0xba, 0xd2, 0x50,
	0x5b, 0xc7, 0x57, 0x6f, 0xb7, 0xf7, 0xde, 0xdb, 0x6e, 0xac, 0x97, 0xd6, 0x42, 0x44, 0x09, 0x2d,
	0x6d, 0x76, 0xf9, 0x01, 0x4e, 0xb7, 0x3e, 0xa7, 0x11, 0xcf, 0x31, 0x29, 0x42, 0x9f, 0x63, 0x42,
	0xce, 0xe1, 0xbf, 0x1f, 0x8c, 0x4b, 0x5c, 0xc6, 0x9e, 0x17, 0xef, 0xaa, 0x37, 0x4a, 0xd3, 0x00,
	0xdd, 0xfa, 0x89, 0x8e, 0x14, 0x5e, 0xb0, 0xb8, 0x97, 0x28, 0xcb, 0xae, 0x09, 0x1c, 0x08, 0xe6,
	0xe6, 0x7e, 0x8f, 0x68, 0xb6, 0x6e, 0xfe, 0x52, 0x81, 0xdc, 0xa6, 0x2f, 0xc2, 0x5b, 0xb8, 0x25,
	0xaa, 0x0d, 0xba, 0x60, 0xf1, 0x7c, 0xf5, 0x44, 0xc6, 0x1b, 0xc7, 0x4d, 0xa3, 0x7c, 0xfe, 0xaf,
	0x28, 0xd7, 0x3b, 0xf5, 0x2a, 0xb4, 0x2e, 0x76, 0x5f, 0xed, 0x1c, 0x9e, 0x3a, 0x79, 0x44, 0x3b,
	0x65, 0xf2, 0x1b, 0x7b, 0xf9, 0xe0, 0x54, 0x7b, 0x15, 0x7a, 0xe1, 0xec, 0xbd, 0xb1, 0x19, 0x5c,
	0x62, 0x91, 0xcb, 0xf8, 0x7b, 0x1a, 0x4c, 0x59, 0x4b, 0xcd, 0xb4, 0x5a, 0xdb, 0x5a, 0xfb, 0x92,
	0xec, 0x55, 0xa8, 0x8e, 0xfb, 0x52, 0x7e, 0x06, 0x27, 0xc5, 0xb1, 0x38, 0xb3, 0x91, 0xe7, 0x8f,
	0xbd, 0x57, 0xa1, 0x8f, 0x96, 0x70, 0x3f, 0x45, 0x3b, 0x87, 0x50, 0x13, 0x2c, 0x72, 0x51, 0xbc,
	0x9a, 0xc0, 0xd9, 0x6d, 0xf1, 0x4f, 0x16, 0xce, 0xc9, 0x19, 0x9c, 0x0c, 0xdb, 0x83, 0x4f, 0x63,
	0x6a, 0x0d, 0xee, 0x46, 0xb4, 0x6b, 0x69, 0x15, 0x72, 0x0e, 0x5a, 0xb7, 0x3f, 0x1a, 0x0c, 0x2d,
	0xba, 0x46, 0x15, 0xf2, 0x04, 0x4e, 0xad, 0x2f, 0x56, 0x77, 0x34, 0xfc, 0x78, 0xf7, 0x79, 0x7c,
	0x3f, 0xb2, 0x46, 0x96, 0x56, 0x4d, 0xbb, 0x0b, 0x6a, 0xbf, 0xdd, 0xb1, 0xfa, 0x9a, 0xda, 0xb9,
	0xfe, 0xf6, 0xc6, 0xf5, 0xc4, 0x4c, 0xda, 0x86, 0x13, 0xf8, 0x26, 0x4f, 0xa6, 0xc2, 0x5c, 0x0d,
	0x06, 0x17, 0x17, 0x66, 0x68, 0xbf, 0x76, 0x03, 0x73, 0x73, 0x56, 0xd8, 0xb5, 0x6c, 0x32, 0x5c,
	0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x0c, 0xea, 0xcf, 0x44, 0x04, 0x00, 0x00,
}
