// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/core/execution.proto

package core

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type WorkflowExecution_Phase int32

const (
	WorkflowExecution_UNDEFINED  WorkflowExecution_Phase = 0
	WorkflowExecution_QUEUED     WorkflowExecution_Phase = 1
	WorkflowExecution_RUNNING    WorkflowExecution_Phase = 2
	WorkflowExecution_SUCCEEDING WorkflowExecution_Phase = 3
	WorkflowExecution_SUCCEEDED  WorkflowExecution_Phase = 4
	WorkflowExecution_FAILING    WorkflowExecution_Phase = 5
	WorkflowExecution_FAILED     WorkflowExecution_Phase = 6
	WorkflowExecution_ABORTED    WorkflowExecution_Phase = 7
	WorkflowExecution_TIMED_OUT  WorkflowExecution_Phase = 8
)

var WorkflowExecution_Phase_name = map[int32]string{
	0: "UNDEFINED",
	1: "QUEUED",
	2: "RUNNING",
	3: "SUCCEEDING",
	4: "SUCCEEDED",
	5: "FAILING",
	6: "FAILED",
	7: "ABORTED",
	8: "TIMED_OUT",
}

var WorkflowExecution_Phase_value = map[string]int32{
	"UNDEFINED":  0,
	"QUEUED":     1,
	"RUNNING":    2,
	"SUCCEEDING": 3,
	"SUCCEEDED":  4,
	"FAILING":    5,
	"FAILED":     6,
	"ABORTED":    7,
	"TIMED_OUT":  8,
}

func (x WorkflowExecution_Phase) String() string {
	return proto.EnumName(WorkflowExecution_Phase_name, int32(x))
}

func (WorkflowExecution_Phase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1523842fd9084ee4, []int{0, 0}
}

type NodeExecution_Phase int32

const (
	NodeExecution_UNDEFINED       NodeExecution_Phase = 0
	NodeExecution_QUEUED          NodeExecution_Phase = 1
	NodeExecution_RUNNING         NodeExecution_Phase = 2
	NodeExecution_SUCCEEDED       NodeExecution_Phase = 3
	NodeExecution_FAILING         NodeExecution_Phase = 4
	NodeExecution_FAILED          NodeExecution_Phase = 5
	NodeExecution_ABORTED         NodeExecution_Phase = 6
	NodeExecution_SKIPPED         NodeExecution_Phase = 7
	NodeExecution_TIMED_OUT       NodeExecution_Phase = 8
	NodeExecution_DYNAMIC_RUNNING NodeExecution_Phase = 9
	NodeExecution_RECOVERED       NodeExecution_Phase = 10
)

var NodeExecution_Phase_name = map[int32]string{
	0:  "UNDEFINED",
	1:  "QUEUED",
	2:  "RUNNING",
	3:  "SUCCEEDED",
	4:  "FAILING",
	5:  "FAILED",
	6:  "ABORTED",
	7:  "SKIPPED",
	8:  "TIMED_OUT",
	9:  "DYNAMIC_RUNNING",
	10: "RECOVERED",
}

var NodeExecution_Phase_value = map[string]int32{
	"UNDEFINED":       0,
	"QUEUED":          1,
	"RUNNING":         2,
	"SUCCEEDED":       3,
	"FAILING":         4,
	"FAILED":          5,
	"ABORTED":         6,
	"SKIPPED":         7,
	"TIMED_OUT":       8,
	"DYNAMIC_RUNNING": 9,
	"RECOVERED":       10,
}

func (x NodeExecution_Phase) String() string {
	return proto.EnumName(NodeExecution_Phase_name, int32(x))
}

func (NodeExecution_Phase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1523842fd9084ee4, []int{1, 0}
}

type TaskExecution_Phase int32

const (
	TaskExecution_UNDEFINED TaskExecution_Phase = 0
	TaskExecution_QUEUED    TaskExecution_Phase = 1
	TaskExecution_RUNNING   TaskExecution_Phase = 2
	TaskExecution_SUCCEEDED TaskExecution_Phase = 3
	TaskExecution_ABORTED   TaskExecution_Phase = 4
	TaskExecution_FAILED    TaskExecution_Phase = 5
	// To indicate cases where task is initializing, like: ErrImagePull, ContainerCreating, PodInitializing
	TaskExecution_INITIALIZING TaskExecution_Phase = 6
	// To address cases, where underlying resource is not available: Backoff error, Resource quota exceeded
	TaskExecution_WAITING_FOR_RESOURCES TaskExecution_Phase = 7
)

var TaskExecution_Phase_name = map[int32]string{
	0: "UNDEFINED",
	1: "QUEUED",
	2: "RUNNING",
	3: "SUCCEEDED",
	4: "ABORTED",
	5: "FAILED",
	6: "INITIALIZING",
	7: "WAITING_FOR_RESOURCES",
}

var TaskExecution_Phase_value = map[string]int32{
	"UNDEFINED":             0,
	"QUEUED":                1,
	"RUNNING":               2,
	"SUCCEEDED":             3,
	"ABORTED":               4,
	"FAILED":                5,
	"INITIALIZING":          6,
	"WAITING_FOR_RESOURCES": 7,
}

func (x TaskExecution_Phase) String() string {
	return proto.EnumName(TaskExecution_Phase_name, int32(x))
}

func (TaskExecution_Phase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1523842fd9084ee4, []int{2, 0}
}

// Error type: System or User
type ExecutionError_ErrorKind int32

const (
	ExecutionError_UNKNOWN ExecutionError_ErrorKind = 0
	ExecutionError_USER    ExecutionError_ErrorKind = 1
	ExecutionError_SYSTEM  ExecutionError_ErrorKind = 2
)

var ExecutionError_ErrorKind_name = map[int32]string{
	0: "UNKNOWN",
	1: "USER",
	2: "SYSTEM",
}

var ExecutionError_ErrorKind_value = map[string]int32{
	"UNKNOWN": 0,
	"USER":    1,
	"SYSTEM":  2,
}

func (x ExecutionError_ErrorKind) String() string {
	return proto.EnumName(ExecutionError_ErrorKind_name, int32(x))
}

func (ExecutionError_ErrorKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1523842fd9084ee4, []int{3, 0}
}

type TaskLog_MessageFormat int32

const (
	TaskLog_UNKNOWN TaskLog_MessageFormat = 0
	TaskLog_CSV     TaskLog_MessageFormat = 1
	TaskLog_JSON    TaskLog_MessageFormat = 2
)

var TaskLog_MessageFormat_name = map[int32]string{
	0: "UNKNOWN",
	1: "CSV",
	2: "JSON",
}

var TaskLog_MessageFormat_value = map[string]int32{
	"UNKNOWN": 0,
	"CSV":     1,
	"JSON":    2,
}

func (x TaskLog_MessageFormat) String() string {
	return proto.EnumName(TaskLog_MessageFormat_name, int32(x))
}

func (TaskLog_MessageFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1523842fd9084ee4, []int{4, 0}
}

type QualityOfService_Tier int32

const (
	// Default: no quality of service specified.
	QualityOfService_UNDEFINED QualityOfService_Tier = 0
	QualityOfService_HIGH      QualityOfService_Tier = 1
	QualityOfService_MEDIUM    QualityOfService_Tier = 2
	QualityOfService_LOW       QualityOfService_Tier = 3
)

var QualityOfService_Tier_name = map[int32]string{
	0: "UNDEFINED",
	1: "HIGH",
	2: "MEDIUM",
	3: "LOW",
}

var QualityOfService_Tier_value = map[string]int32{
	"UNDEFINED": 0,
	"HIGH":      1,
	"MEDIUM":    2,
	"LOW":       3,
}

func (x QualityOfService_Tier) String() string {
	return proto.EnumName(QualityOfService_Tier_name, int32(x))
}

func (QualityOfService_Tier) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1523842fd9084ee4, []int{6, 0}
}

// Indicates various phases of Workflow Execution
type WorkflowExecution struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowExecution) Reset()         { *m = WorkflowExecution{} }
func (m *WorkflowExecution) String() string { return proto.CompactTextString(m) }
func (*WorkflowExecution) ProtoMessage()    {}
func (*WorkflowExecution) Descriptor() ([]byte, []int) {
	return fileDescriptor_1523842fd9084ee4, []int{0}
}

func (m *WorkflowExecution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowExecution.Unmarshal(m, b)
}
func (m *WorkflowExecution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowExecution.Marshal(b, m, deterministic)
}
func (m *WorkflowExecution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowExecution.Merge(m, src)
}
func (m *WorkflowExecution) XXX_Size() int {
	return xxx_messageInfo_WorkflowExecution.Size(m)
}
func (m *WorkflowExecution) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowExecution.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowExecution proto.InternalMessageInfo

// Indicates various phases of Node Execution
type NodeExecution struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeExecution) Reset()         { *m = NodeExecution{} }
func (m *NodeExecution) String() string { return proto.CompactTextString(m) }
func (*NodeExecution) ProtoMessage()    {}
func (*NodeExecution) Descriptor() ([]byte, []int) {
	return fileDescriptor_1523842fd9084ee4, []int{1}
}

func (m *NodeExecution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecution.Unmarshal(m, b)
}
func (m *NodeExecution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecution.Marshal(b, m, deterministic)
}
func (m *NodeExecution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecution.Merge(m, src)
}
func (m *NodeExecution) XXX_Size() int {
	return xxx_messageInfo_NodeExecution.Size(m)
}
func (m *NodeExecution) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecution.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecution proto.InternalMessageInfo

// Phases that task plugins can go through. Not all phases may be applicable to a specific plugin task,
// but this is the cumulative list that customers may want to know about for their task.
type TaskExecution struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskExecution) Reset()         { *m = TaskExecution{} }
func (m *TaskExecution) String() string { return proto.CompactTextString(m) }
func (*TaskExecution) ProtoMessage()    {}
func (*TaskExecution) Descriptor() ([]byte, []int) {
	return fileDescriptor_1523842fd9084ee4, []int{2}
}

func (m *TaskExecution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecution.Unmarshal(m, b)
}
func (m *TaskExecution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecution.Marshal(b, m, deterministic)
}
func (m *TaskExecution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecution.Merge(m, src)
}
func (m *TaskExecution) XXX_Size() int {
	return xxx_messageInfo_TaskExecution.Size(m)
}
func (m *TaskExecution) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecution.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecution proto.InternalMessageInfo

// Represents the error message from the execution.
type ExecutionError struct {
	// Error code indicates a grouping of a type of error.
	// More Info: <Link>
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	// Detailed description of the error - including stack trace.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// Full error contents accessible via a URI
	ErrorUri             string                   `protobuf:"bytes,3,opt,name=error_uri,json=errorUri,proto3" json:"error_uri,omitempty"`
	Kind                 ExecutionError_ErrorKind `protobuf:"varint,4,opt,name=kind,proto3,enum=flyteidl.core.ExecutionError_ErrorKind" json:"kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ExecutionError) Reset()         { *m = ExecutionError{} }
func (m *ExecutionError) String() string { return proto.CompactTextString(m) }
func (*ExecutionError) ProtoMessage()    {}
func (*ExecutionError) Descriptor() ([]byte, []int) {
	return fileDescriptor_1523842fd9084ee4, []int{3}
}

func (m *ExecutionError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionError.Unmarshal(m, b)
}
func (m *ExecutionError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionError.Marshal(b, m, deterministic)
}
func (m *ExecutionError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionError.Merge(m, src)
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

func (m *ExecutionError) GetKind() ExecutionError_ErrorKind {
	if m != nil {
		return m.Kind
	}
	return ExecutionError_UNKNOWN
}

// Log information for the task that is specific to a log sink
// When our log story is flushed out, we may have more metadata here like log link expiry
type TaskLog struct {
	Uri                  string                `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	Name                 string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MessageFormat        TaskLog_MessageFormat `protobuf:"varint,3,opt,name=message_format,json=messageFormat,proto3,enum=flyteidl.core.TaskLog_MessageFormat" json:"message_format,omitempty"`
	Ttl                  *duration.Duration    `protobuf:"bytes,4,opt,name=ttl,proto3" json:"ttl,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TaskLog) Reset()         { *m = TaskLog{} }
func (m *TaskLog) String() string { return proto.CompactTextString(m) }
func (*TaskLog) ProtoMessage()    {}
func (*TaskLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_1523842fd9084ee4, []int{4}
}

func (m *TaskLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskLog.Unmarshal(m, b)
}
func (m *TaskLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskLog.Marshal(b, m, deterministic)
}
func (m *TaskLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskLog.Merge(m, src)
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

func (m *TaskLog) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TaskLog) GetMessageFormat() TaskLog_MessageFormat {
	if m != nil {
		return m.MessageFormat
	}
	return TaskLog_UNKNOWN
}

func (m *TaskLog) GetTtl() *duration.Duration {
	if m != nil {
		return m.Ttl
	}
	return nil
}

// Represents customized execution run-time attributes.
type QualityOfServiceSpec struct {
	// Indicates how much queueing delay an execution can tolerate.
	QueueingBudget       *duration.Duration `protobuf:"bytes,1,opt,name=queueing_budget,json=queueingBudget,proto3" json:"queueing_budget,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *QualityOfServiceSpec) Reset()         { *m = QualityOfServiceSpec{} }
func (m *QualityOfServiceSpec) String() string { return proto.CompactTextString(m) }
func (*QualityOfServiceSpec) ProtoMessage()    {}
func (*QualityOfServiceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_1523842fd9084ee4, []int{5}
}

func (m *QualityOfServiceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QualityOfServiceSpec.Unmarshal(m, b)
}
func (m *QualityOfServiceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QualityOfServiceSpec.Marshal(b, m, deterministic)
}
func (m *QualityOfServiceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QualityOfServiceSpec.Merge(m, src)
}
func (m *QualityOfServiceSpec) XXX_Size() int {
	return xxx_messageInfo_QualityOfServiceSpec.Size(m)
}
func (m *QualityOfServiceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_QualityOfServiceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_QualityOfServiceSpec proto.InternalMessageInfo

func (m *QualityOfServiceSpec) GetQueueingBudget() *duration.Duration {
	if m != nil {
		return m.QueueingBudget
	}
	return nil
}

// Indicates the priority of an execution.
type QualityOfService struct {
	// Types that are valid to be assigned to Designation:
	//	*QualityOfService_Tier_
	//	*QualityOfService_Spec
	Designation          isQualityOfService_Designation `protobuf_oneof:"designation"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *QualityOfService) Reset()         { *m = QualityOfService{} }
func (m *QualityOfService) String() string { return proto.CompactTextString(m) }
func (*QualityOfService) ProtoMessage()    {}
func (*QualityOfService) Descriptor() ([]byte, []int) {
	return fileDescriptor_1523842fd9084ee4, []int{6}
}

func (m *QualityOfService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QualityOfService.Unmarshal(m, b)
}
func (m *QualityOfService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QualityOfService.Marshal(b, m, deterministic)
}
func (m *QualityOfService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QualityOfService.Merge(m, src)
}
func (m *QualityOfService) XXX_Size() int {
	return xxx_messageInfo_QualityOfService.Size(m)
}
func (m *QualityOfService) XXX_DiscardUnknown() {
	xxx_messageInfo_QualityOfService.DiscardUnknown(m)
}

var xxx_messageInfo_QualityOfService proto.InternalMessageInfo

type isQualityOfService_Designation interface {
	isQualityOfService_Designation()
}

type QualityOfService_Tier_ struct {
	Tier QualityOfService_Tier `protobuf:"varint,1,opt,name=tier,proto3,enum=flyteidl.core.QualityOfService_Tier,oneof"`
}

type QualityOfService_Spec struct {
	Spec *QualityOfServiceSpec `protobuf:"bytes,2,opt,name=spec,proto3,oneof"`
}

func (*QualityOfService_Tier_) isQualityOfService_Designation() {}

func (*QualityOfService_Spec) isQualityOfService_Designation() {}

func (m *QualityOfService) GetDesignation() isQualityOfService_Designation {
	if m != nil {
		return m.Designation
	}
	return nil
}

func (m *QualityOfService) GetTier() QualityOfService_Tier {
	if x, ok := m.GetDesignation().(*QualityOfService_Tier_); ok {
		return x.Tier
	}
	return QualityOfService_UNDEFINED
}

func (m *QualityOfService) GetSpec() *QualityOfServiceSpec {
	if x, ok := m.GetDesignation().(*QualityOfService_Spec); ok {
		return x.Spec
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*QualityOfService) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*QualityOfService_Tier_)(nil),
		(*QualityOfService_Spec)(nil),
	}
}

func init() {
	proto.RegisterEnum("flyteidl.core.WorkflowExecution_Phase", WorkflowExecution_Phase_name, WorkflowExecution_Phase_value)
	proto.RegisterEnum("flyteidl.core.NodeExecution_Phase", NodeExecution_Phase_name, NodeExecution_Phase_value)
	proto.RegisterEnum("flyteidl.core.TaskExecution_Phase", TaskExecution_Phase_name, TaskExecution_Phase_value)
	proto.RegisterEnum("flyteidl.core.ExecutionError_ErrorKind", ExecutionError_ErrorKind_name, ExecutionError_ErrorKind_value)
	proto.RegisterEnum("flyteidl.core.TaskLog_MessageFormat", TaskLog_MessageFormat_name, TaskLog_MessageFormat_value)
	proto.RegisterEnum("flyteidl.core.QualityOfService_Tier", QualityOfService_Tier_name, QualityOfService_Tier_value)
	proto.RegisterType((*WorkflowExecution)(nil), "flyteidl.core.WorkflowExecution")
	proto.RegisterType((*NodeExecution)(nil), "flyteidl.core.NodeExecution")
	proto.RegisterType((*TaskExecution)(nil), "flyteidl.core.TaskExecution")
	proto.RegisterType((*ExecutionError)(nil), "flyteidl.core.ExecutionError")
	proto.RegisterType((*TaskLog)(nil), "flyteidl.core.TaskLog")
	proto.RegisterType((*QualityOfServiceSpec)(nil), "flyteidl.core.QualityOfServiceSpec")
	proto.RegisterType((*QualityOfService)(nil), "flyteidl.core.QualityOfService")
}

func init() { proto.RegisterFile("flyteidl/core/execution.proto", fileDescriptor_1523842fd9084ee4) }

var fileDescriptor_1523842fd9084ee4 = []byte{
	// 724 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xdd, 0x6e, 0xda, 0x48,
	0x14, 0xc7, 0x63, 0x70, 0xf8, 0x38, 0x2c, 0x64, 0x76, 0x76, 0x57, 0x22, 0xbb, 0xda, 0x55, 0xe4,
	0x5d, 0x69, 0x23, 0x55, 0xb5, 0x25, 0x1a, 0x55, 0x6a, 0x7b, 0x05, 0x78, 0x48, 0x5c, 0xc0, 0x4e,
	0xfc, 0x11, 0x94, 0xdc, 0x20, 0xc0, 0x83, 0x63, 0x05, 0x18, 0x6a, 0xec, 0xb6, 0xb9, 0x6e, 0x9f,
	0xa1, 0x52, 0xef, 0xfb, 0x0c, 0x7d, 0x86, 0x5e, 0xf5, 0x01, 0xfa, 0x34, 0xd5, 0x0c, 0x38, 0x01,
	0x5a, 0xb5, 0xaa, 0xd4, 0x1b, 0x34, 0x67, 0xce, 0xf9, 0x9f, 0xf9, 0xfd, 0x0f, 0x9e, 0x81, 0xbf,
	0xc7, 0x93, 0x9b, 0x98, 0x86, 0xfe, 0x44, 0x1b, 0xb1, 0x88, 0x6a, 0xf4, 0x25, 0x1d, 0x25, 0x71,
	0xc8, 0x66, 0xea, 0x3c, 0x62, 0x31, 0xc3, 0xe5, 0x34, 0xad, 0xf2, 0xf4, 0x9f, 0xff, 0x04, 0x8c,
	0x05, 0x13, 0xaa, 0x89, 0xe4, 0x30, 0x19, 0x6b, 0x7e, 0x12, 0x0d, 0xee, 0xca, 0x95, 0xb7, 0x12,
	0xfc, 0xda, 0x63, 0xd1, 0xf5, 0x78, 0xc2, 0x5e, 0x90, 0xb4, 0x95, 0xf2, 0x4a, 0x82, 0xdd, 0xd3,
	0xab, 0xc1, 0x82, 0xe2, 0x32, 0x14, 0x3d, 0x53, 0x27, 0x2d, 0xc3, 0x24, 0x3a, 0xda, 0xc1, 0x00,
	0xb9, 0x33, 0x8f, 0x78, 0x44, 0x47, 0x12, 0x2e, 0x41, 0xde, 0xf6, 0x4c, 0xd3, 0x30, 0x8f, 0x51,
	0x06, 0x57, 0x00, 0x1c, 0xaf, 0xd9, 0x24, 0x44, 0xe7, 0x71, 0x96, 0xeb, 0x56, 0x31, 0xd1, 0x91,
	0xcc, 0x6b, 0x5b, 0x75, 0xa3, 0xc3, 0x73, 0xbb, 0xbc, 0x09, 0x0f, 0x88, 0x8e, 0x72, 0x3c, 0x51,
	0x6f, 0x58, 0xb6, 0x4b, 0x74, 0x94, 0xe7, 0x22, 0xd7, 0xe8, 0x12, 0xbd, 0x6f, 0x79, 0x2e, 0x2a,
	0x28, 0xef, 0x25, 0x28, 0x9b, 0xcc, 0xa7, 0x77, 0x5c, 0xef, 0x7e, 0x98, 0x6b, 0x83, 0x23, 0xbb,
	0xce, 0x21, 0xaf, 0x71, 0xec, 0xae, 0x73, 0x08, 0x28, 0xa7, 0x6d, 0x9c, 0x9e, 0x7e, 0x05, 0x0a,
	0xff, 0x06, 0x7b, 0xfa, 0x85, 0x59, 0xef, 0x1a, 0xcd, 0x7e, 0x7a, 0x4a, 0x91, 0xd7, 0xd8, 0xa4,
	0x69, 0x9d, 0x13, 0x9b, 0xe8, 0x08, 0x94, 0x37, 0x12, 0x94, 0xdd, 0xc1, 0xe2, 0xfa, 0x0e, 0xfc,
	0xf5, 0x4f, 0x00, 0x4f, 0xf9, 0x36, 0xc1, 0x11, 0xfc, 0x62, 0x98, 0x86, 0x6b, 0xd4, 0x3b, 0xc6,
	0x25, 0x57, 0xe6, 0xf0, 0x3e, 0xfc, 0xd1, 0xab, 0x1b, 0xae, 0x61, 0x1e, 0xf7, 0x5b, 0x96, 0xdd,
	0xb7, 0x89, 0x63, 0x79, 0x76, 0x93, 0x38, 0x28, 0xaf, 0x7c, 0x90, 0xa0, 0x72, 0x0b, 0x45, 0xa2,
	0x88, 0x45, 0x18, 0x83, 0x3c, 0x62, 0x3e, 0xad, 0x4a, 0x07, 0xd2, 0x61, 0xd1, 0x16, 0x6b, 0x5c,
	0x85, 0xfc, 0x94, 0x2e, 0x16, 0x83, 0x80, 0x56, 0x33, 0x62, 0x3b, 0x0d, 0xf1, 0x5f, 0x50, 0xa4,
	0x5c, 0xd6, 0x4f, 0xa2, 0xb0, 0x9a, 0x15, 0xb9, 0x82, 0xd8, 0xf0, 0xa2, 0x10, 0x3f, 0x01, 0xf9,
	0x3a, 0x9c, 0xf9, 0x55, 0xf9, 0x40, 0x3a, 0xac, 0xd4, 0xfe, 0x57, 0x37, 0xbe, 0x44, 0x75, 0xf3,
	0x5c, 0x55, 0xfc, 0xb6, 0xc3, 0x99, 0x6f, 0x0b, 0x91, 0xa2, 0x42, 0xf1, 0x76, 0x8b, 0xbb, 0xf5,
	0xcc, 0xb6, 0x69, 0xf5, 0x4c, 0xb4, 0x83, 0x0b, 0x20, 0x7b, 0x0e, 0xb1, 0x91, 0xc4, 0x7d, 0x3b,
	0x17, 0x8e, 0x4b, 0xba, 0x28, 0xa3, 0x7c, 0x92, 0x20, 0xcf, 0x67, 0xdc, 0x61, 0x01, 0x46, 0x90,
	0xe5, 0x3c, 0x4b, 0x0b, 0x7c, 0xc9, 0x5d, 0xcd, 0x06, 0xd3, 0x14, 0x5f, 0xac, 0x71, 0x1b, 0x2a,
	0x2b, 0x1b, 0xfd, 0x31, 0x8b, 0xa6, 0x83, 0x58, 0x18, 0xa8, 0xd4, 0xfe, 0xdb, 0x02, 0x5d, 0x75,
	0x55, 0xbb, 0xcb, 0xe2, 0x96, 0xa8, 0xb5, 0xcb, 0xd3, 0xf5, 0x10, 0xdf, 0x83, 0x6c, 0x1c, 0x4f,
	0x84, 0xd5, 0x52, 0x6d, 0x5f, 0x5d, 0xde, 0x32, 0x35, 0xbd, 0x65, 0xaa, 0xbe, 0xba, 0x65, 0x36,
	0xaf, 0x52, 0x34, 0x28, 0x6f, 0x34, 0xdb, 0xf4, 0x97, 0x87, 0x6c, 0xd3, 0x39, 0x47, 0x12, 0x37,
	0xfa, 0xd4, 0xb1, 0x4c, 0x94, 0x51, 0x2e, 0xe1, 0xf7, 0xb3, 0x64, 0x30, 0x09, 0xe3, 0x1b, 0x6b,
	0xec, 0xd0, 0xe8, 0x79, 0x38, 0xa2, 0xce, 0x9c, 0x8e, 0x70, 0x03, 0xf6, 0x9e, 0x25, 0x34, 0xa1,
	0xe1, 0x2c, 0xe8, 0x0f, 0x13, 0x3f, 0xa0, 0xb1, 0x30, 0xfd, 0x4d, 0x82, 0x4a, 0xaa, 0x68, 0x08,
	0x81, 0xf2, 0x51, 0x02, 0xb4, 0xdd, 0x1c, 0x3f, 0x06, 0x39, 0x0e, 0x69, 0x24, 0xba, 0x7d, 0x39,
	0x91, 0xed, 0x72, 0xd5, 0x0d, 0x69, 0x74, 0xb2, 0x63, 0x0b, 0x0d, 0x7e, 0x04, 0xf2, 0x62, 0x4e,
	0x47, 0x62, 0xd6, 0xa5, 0xda, 0xbf, 0xdf, 0xd1, 0x72, 0x1f, 0x5c, 0xca, 0x25, 0xca, 0x11, 0xc8,
	0xbc, 0xd5, 0xf6, 0xa5, 0x28, 0x80, 0x7c, 0x62, 0x1c, 0x9f, 0x2c, 0xff, 0xf1, 0x2e, 0xd1, 0x0d,
	0xaf, 0x8b, 0x32, 0x7c, 0x4e, 0x1d, 0xab, 0x87, 0xb2, 0x8d, 0x32, 0x94, 0x7c, 0xba, 0x08, 0x83,
	0x99, 0x30, 0xd8, 0x78, 0x78, 0x79, 0x14, 0x84, 0xf1, 0x55, 0x32, 0x54, 0x47, 0x6c, 0xaa, 0x89,
	0xd3, 0x59, 0x14, 0x68, 0xb7, 0xcf, 0x64, 0x40, 0x67, 0xda, 0x7c, 0x78, 0x3f, 0x60, 0xda, 0xc6,
	0xcb, 0x39, 0xcc, 0x89, 0x59, 0x3d, 0xf8, 0x1c, 0x00, 0x00, 0xff, 0xff, 0x26, 0xb3, 0xbb, 0x44,
	0x51, 0x05, 0x00, 0x00,
}
