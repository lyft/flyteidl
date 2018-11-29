// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/core/execution.proto

package core // import "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Indicates various phases of Workflow Execution
type WorkflowExecutionPhase int32

const (
	WorkflowExecutionPhase_WORKFLOW_PHASE_UNDEFINED  WorkflowExecutionPhase = 0
	WorkflowExecutionPhase_WORKFLOW_PHASE_RUNNING    WorkflowExecutionPhase = 1
	WorkflowExecutionPhase_WORKFLOW_PHASE_SUCCEEDING WorkflowExecutionPhase = 2
	WorkflowExecutionPhase_WORKFLOW_PHASE_SUCCEEDED  WorkflowExecutionPhase = 3
	WorkflowExecutionPhase_WORKFLOW_PHASE_FAILING    WorkflowExecutionPhase = 4
	WorkflowExecutionPhase_WORKFLOW_PHASE_FAILED     WorkflowExecutionPhase = 5
	WorkflowExecutionPhase_WORKFLOW_PHASE_TIMED_OUT  WorkflowExecutionPhase = 6
	WorkflowExecutionPhase_WORKFLOW_PHASE_ABORTED    WorkflowExecutionPhase = 7
	WorkflowExecutionPhase_WORKFLOW_PHASE_QUEUED     WorkflowExecutionPhase = 8
)

var WorkflowExecutionPhase_name = map[int32]string{
	0: "WORKFLOW_PHASE_UNDEFINED",
	1: "WORKFLOW_PHASE_RUNNING",
	2: "WORKFLOW_PHASE_SUCCEEDING",
	3: "WORKFLOW_PHASE_SUCCEEDED",
	4: "WORKFLOW_PHASE_FAILING",
	5: "WORKFLOW_PHASE_FAILED",
	6: "WORKFLOW_PHASE_TIMED_OUT",
	7: "WORKFLOW_PHASE_ABORTED",
	8: "WORKFLOW_PHASE_QUEUED",
}
var WorkflowExecutionPhase_value = map[string]int32{
	"WORKFLOW_PHASE_UNDEFINED":  0,
	"WORKFLOW_PHASE_RUNNING":    1,
	"WORKFLOW_PHASE_SUCCEEDING": 2,
	"WORKFLOW_PHASE_SUCCEEDED":  3,
	"WORKFLOW_PHASE_FAILING":    4,
	"WORKFLOW_PHASE_FAILED":     5,
	"WORKFLOW_PHASE_TIMED_OUT":  6,
	"WORKFLOW_PHASE_ABORTED":    7,
	"WORKFLOW_PHASE_QUEUED":     8,
}

func (x WorkflowExecutionPhase) String() string {
	return proto.EnumName(WorkflowExecutionPhase_name, int32(x))
}
func (WorkflowExecutionPhase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_execution_c316b2f86bb68dc7, []int{0}
}

// Indicates various phases of Node Execution
type NodeExecutionPhase int32

const (
	NodeExecutionPhase_NODE_PHASE_UNDEFINED NodeExecutionPhase = 0
	NodeExecutionPhase_NODE_PHASE_RUNNING   NodeExecutionPhase = 1
	NodeExecutionPhase_NODE_PHASE_SUCCEEDED NodeExecutionPhase = 2
	NodeExecutionPhase_NODE_PHASE_FAILING   NodeExecutionPhase = 3
	NodeExecutionPhase_NODE_PHASE_FAILED    NodeExecutionPhase = 4
	NodeExecutionPhase_NODE_PHASE_TIMED_OUT NodeExecutionPhase = 5
	NodeExecutionPhase_NODE_PHASE_SKIPPED   NodeExecutionPhase = 6
	NodeExecutionPhase_NODE_PHASE_ABORTED   NodeExecutionPhase = 7
	NodeExecutionPhase_NODE_PHASE_QUEUED    NodeExecutionPhase = 8
)

var NodeExecutionPhase_name = map[int32]string{
	0: "NODE_PHASE_UNDEFINED",
	1: "NODE_PHASE_RUNNING",
	2: "NODE_PHASE_SUCCEEDED",
	3: "NODE_PHASE_FAILING",
	4: "NODE_PHASE_FAILED",
	5: "NODE_PHASE_TIMED_OUT",
	6: "NODE_PHASE_SKIPPED",
	7: "NODE_PHASE_ABORTED",
	8: "NODE_PHASE_QUEUED",
}
var NodeExecutionPhase_value = map[string]int32{
	"NODE_PHASE_UNDEFINED": 0,
	"NODE_PHASE_RUNNING":   1,
	"NODE_PHASE_SUCCEEDED": 2,
	"NODE_PHASE_FAILING":   3,
	"NODE_PHASE_FAILED":    4,
	"NODE_PHASE_TIMED_OUT": 5,
	"NODE_PHASE_SKIPPED":   6,
	"NODE_PHASE_ABORTED":   7,
	"NODE_PHASE_QUEUED":    8,
}

func (x NodeExecutionPhase) String() string {
	return proto.EnumName(NodeExecutionPhase_name, int32(x))
}
func (NodeExecutionPhase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_execution_c316b2f86bb68dc7, []int{1}
}

// Phases that task plugins can go through. Not all phases may be applicable to a specific plugin task,
// but this is the cumulative list that customers may want to know about for their task.
type TaskExecutionPhase int32

const (
	TaskExecutionPhase_TASK_PHASE_UNDEFINED TaskExecutionPhase = 0
	TaskExecutionPhase_TASK_PHASE_QUEUED    TaskExecutionPhase = 1
	TaskExecutionPhase_TASK_PHASE_RUNNABLE  TaskExecutionPhase = 2
	TaskExecutionPhase_TASK_PHASE_RUNNING   TaskExecutionPhase = 3
	TaskExecutionPhase_TASK_PHASE_SUCCEEDED TaskExecutionPhase = 4
	TaskExecutionPhase_TASK_PHASE_FAILED    TaskExecutionPhase = 5
	TaskExecutionPhase_TASK_PHASE_ABORTED   TaskExecutionPhase = 6
)

var TaskExecutionPhase_name = map[int32]string{
	0: "TASK_PHASE_UNDEFINED",
	1: "TASK_PHASE_QUEUED",
	2: "TASK_PHASE_RUNNABLE",
	3: "TASK_PHASE_RUNNING",
	4: "TASK_PHASE_SUCCEEDED",
	5: "TASK_PHASE_FAILED",
	6: "TASK_PHASE_ABORTED",
}
var TaskExecutionPhase_value = map[string]int32{
	"TASK_PHASE_UNDEFINED": 0,
	"TASK_PHASE_QUEUED":    1,
	"TASK_PHASE_RUNNABLE":  2,
	"TASK_PHASE_RUNNING":   3,
	"TASK_PHASE_SUCCEEDED": 4,
	"TASK_PHASE_FAILED":    5,
	"TASK_PHASE_ABORTED":   6,
}

func (x TaskExecutionPhase) String() string {
	return proto.EnumName(TaskExecutionPhase_name, int32(x))
}
func (TaskExecutionPhase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_execution_c316b2f86bb68dc7, []int{2}
}

// Represents the error message from the execution.
type ExecutionError struct {
	// Error code indicates a grouping of a type of error.
	// More Info: <Link>
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	// Detailed description of the error - including stack trace.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// Full error contents accessible via a URI
	ErrorUri             string   `protobuf:"bytes,3,opt,name=error_uri,json=errorUri,proto3" json:"error_uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecutionError) Reset()         { *m = ExecutionError{} }
func (m *ExecutionError) String() string { return proto.CompactTextString(m) }
func (*ExecutionError) ProtoMessage()    {}
func (*ExecutionError) Descriptor() ([]byte, []int) {
	return fileDescriptor_execution_c316b2f86bb68dc7, []int{0}
}
func (m *ExecutionError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionError.Unmarshal(m, b)
}
func (m *ExecutionError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionError.Marshal(b, m, deterministic)
}
func (dst *ExecutionError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionError.Merge(dst, src)
}
func (m *ExecutionError) XXX_Size() int {
	return xxx_messageInfo_ExecutionError.Size(m)
}
func (m *ExecutionError) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionError.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionError proto.InternalMessageInfo

func (m *ExecutionError) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *ExecutionError) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ExecutionError) GetErrorUri() string {
	if m != nil {
		return m.ErrorUri
	}
	return ""
}

// Log information for the task that is specific to a log sink
// When our log story is flushed out, we may have more metadata here like log link expiry
type TaskLog struct {
	Uri                  string   `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskLog) Reset()         { *m = TaskLog{} }
func (m *TaskLog) String() string { return proto.CompactTextString(m) }
func (*TaskLog) ProtoMessage()    {}
func (*TaskLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_execution_c316b2f86bb68dc7, []int{1}
}
func (m *TaskLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskLog.Unmarshal(m, b)
}
func (m *TaskLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskLog.Marshal(b, m, deterministic)
}
func (dst *TaskLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskLog.Merge(dst, src)
}
func (m *TaskLog) XXX_Size() int {
	return xxx_messageInfo_TaskLog.Size(m)
}
func (m *TaskLog) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskLog.DiscardUnknown(m)
}

var xxx_messageInfo_TaskLog proto.InternalMessageInfo

func (m *TaskLog) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func init() {
	proto.RegisterType((*ExecutionError)(nil), "flyteidl.core.ExecutionError")
	proto.RegisterType((*TaskLog)(nil), "flyteidl.core.TaskLog")
	proto.RegisterEnum("flyteidl.core.WorkflowExecutionPhase", WorkflowExecutionPhase_name, WorkflowExecutionPhase_value)
	proto.RegisterEnum("flyteidl.core.NodeExecutionPhase", NodeExecutionPhase_name, NodeExecutionPhase_value)
	proto.RegisterEnum("flyteidl.core.TaskExecutionPhase", TaskExecutionPhase_name, TaskExecutionPhase_value)
}

func init() {
	proto.RegisterFile("flyteidl/core/execution.proto", fileDescriptor_execution_c316b2f86bb68dc7)
}

var fileDescriptor_execution_c316b2f86bb68dc7 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x49, 0xdb, 0xb5, 0xdb, 0x91, 0x40, 0x9e, 0xa1, 0x25, 0x63, 0x4c, 0x42, 0xbb, 0x42,
	0x93, 0x68, 0x10, 0x3c, 0x41, 0x3a, 0x9f, 0x42, 0xd4, 0x92, 0x94, 0x36, 0x51, 0x25, 0xb8, 0xa8,
	0xfa, 0xc7, 0xcd, 0xa2, 0x75, 0xf3, 0xe4, 0xa6, 0x82, 0x3d, 0x0b, 0x0f, 0xc3, 0x2b, 0xf1, 0x08,
	0xc8, 0x69, 0xd3, 0x34, 0x9e, 0xef, 0x9c, 0xf3, 0xf3, 0xf9, 0x4e, 0xbe, 0x2f, 0x39, 0x70, 0xb1,
	0x5c, 0x3d, 0xa6, 0x3c, 0x59, 0xac, 0x9c, 0xb9, 0x90, 0xdc, 0xe1, 0xbf, 0xf9, 0x7c, 0x93, 0x26,
	0xe2, 0xbe, 0xfd, 0x20, 0x45, 0x2a, 0xe8, 0xf3, 0x1c, 0xb7, 0x15, 0xbe, 0xfc, 0x09, 0x2f, 0x30,
	0xbf, 0x81, 0x52, 0x0a, 0x49, 0x29, 0xd4, 0xe6, 0x62, 0xc1, 0x6d, 0xeb, 0x9d, 0xf5, 0xfe, 0x64,
	0x98, 0x9d, 0xa9, 0x0d, 0x8d, 0x3b, 0xbe, 0x5e, 0x4f, 0x63, 0x6e, 0x57, 0xb2, 0x72, 0xfe, 0x48,
	0xcf, 0xe1, 0x84, 0xab, 0xb6, 0xc9, 0x46, 0x26, 0x76, 0x35, 0x63, 0xc7, 0x59, 0x21, 0x92, 0xc9,
	0xe5, 0x39, 0x34, 0xc2, 0xe9, 0xfa, 0xb6, 0x2f, 0x62, 0x4a, 0xa0, 0xaa, 0x6e, 0x6c, 0x45, 0xd5,
	0xf1, 0xea, 0x4f, 0x05, 0x5a, 0x63, 0x21, 0x6f, 0x97, 0x2b, 0xf1, 0x6b, 0xff, 0x0a, 0x83, 0x9b,
	0xe9, 0x9a, 0xd3, 0xb7, 0x60, 0x8f, 0x83, 0x61, 0xaf, 0xdb, 0x0f, 0xc6, 0x93, 0xc1, 0x57, 0x77,
	0x84, 0x93, 0xc8, 0x67, 0xd8, 0xf5, 0x7c, 0x64, 0xe4, 0x19, 0x7d, 0x03, 0x2d, 0x8d, 0x0e, 0x23,
	0xdf, 0xf7, 0xfc, 0x2f, 0xc4, 0xa2, 0x17, 0x70, 0xa6, 0xb1, 0x51, 0x74, 0x7d, 0x8d, 0xc8, 0x14,
	0xae, 0x18, 0x84, 0x77, 0x18, 0x19, 0xa9, 0x1a, 0x84, 0xbb, 0xae, 0xd7, 0x57, 0x9d, 0x35, 0x7a,
	0x06, 0x4d, 0x03, 0x43, 0x46, 0x8e, 0x0c, 0xa2, 0xa1, 0xf7, 0x0d, 0xd9, 0x24, 0x88, 0x42, 0x52,
	0x37, 0x88, 0xba, 0x9d, 0x60, 0x18, 0x22, 0x23, 0x0d, 0x83, 0xe8, 0xf7, 0x08, 0x23, 0x64, 0xe4,
	0xf8, 0xea, 0x9f, 0x05, 0xd4, 0x17, 0x0b, 0xae, 0x25, 0x63, 0xc3, 0x2b, 0x3f, 0x60, 0x68, 0x48,
	0xa5, 0x05, 0xf4, 0x80, 0x14, 0x89, 0x94, 0x3b, 0x0a, 0xbb, 0x15, 0xad, 0x23, 0xb7, 0x5a, 0xa5,
	0x4d, 0x38, 0xd5, 0xea, 0xc8, 0x48, 0x4d, 0x13, 0x2a, 0x2c, 0x1e, 0x69, 0x42, 0xa3, 0x9e, 0x37,
	0x18, 0x20, 0x23, 0x75, 0xad, 0x5e, 0xd8, 0x2e, 0x0f, 0xd8, 0x5b, 0xfe, 0x6b, 0x01, 0x55, 0xbf,
	0xcb, 0x53, 0xcb, 0xa1, 0x3b, 0xea, 0x19, 0x2c, 0x37, 0xe1, 0xf4, 0x80, 0xec, 0x74, 0x2c, 0xfa,
	0x1a, 0x5e, 0x1e, 0x94, 0x55, 0x12, 0x6e, 0xa7, 0x8f, 0x5b, 0xc3, 0x1a, 0xd8, 0x1a, 0x2e, 0x4f,
	0x28, 0x22, 0xaa, 0x69, 0x13, 0xf6, 0x5f, 0xbc, 0x2c, 0x94, 0x1b, 0xab, 0x77, 0x3e, 0xfd, 0xf8,
	0x18, 0x27, 0xe9, 0xcd, 0x66, 0xd6, 0x9e, 0x8b, 0x3b, 0x67, 0xf5, 0xb8, 0x4c, 0x9d, 0xfd, 0x32,
	0xc6, 0xfc, 0xde, 0x79, 0x98, 0x7d, 0x88, 0x85, 0x53, 0xda, 0xcf, 0x59, 0x3d, 0x5b, 0xcb, 0xcf,
	0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x10, 0xa2, 0xff, 0x5a, 0xb7, 0x03, 0x00, 0x00,
}
