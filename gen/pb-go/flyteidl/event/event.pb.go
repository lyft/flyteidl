// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/event/event.proto

package event // import "github.com/lyft/flyteidl/gen/pb-go/flyteidl/event"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/golang/protobuf/ptypes/struct"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import core "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WorkflowExecutionEvent struct {
	// Workflow execution id
	ExecutionId *core.WorkflowExecutionIdentifier `protobuf:"bytes,1,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	// the id of the originator (Propeller) of the event
	ProducerId string                      `protobuf:"bytes,2,opt,name=producer_id,json=producerId,proto3" json:"producer_id,omitempty"`
	Phase      core.WorkflowExecutionPhase `protobuf:"varint,3,opt,name=phase,proto3,enum=flyteidl.core.WorkflowExecutionPhase" json:"phase,omitempty"`
	// This timestamp represents when the original event occurred, it is generated
	// by the executor of the workflow.
	OccurredAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	// Types that are valid to be assigned to OutputResult:
	//	*WorkflowExecutionEvent_OutputUri
	//	*WorkflowExecutionEvent_Error
	OutputResult         isWorkflowExecutionEvent_OutputResult `protobuf_oneof:"output_result"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *WorkflowExecutionEvent) Reset()         { *m = WorkflowExecutionEvent{} }
func (m *WorkflowExecutionEvent) String() string { return proto.CompactTextString(m) }
func (*WorkflowExecutionEvent) ProtoMessage()    {}
func (*WorkflowExecutionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_7f0ab3e135fe4708, []int{0}
}
func (m *WorkflowExecutionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowExecutionEvent.Unmarshal(m, b)
}
func (m *WorkflowExecutionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowExecutionEvent.Marshal(b, m, deterministic)
}
func (dst *WorkflowExecutionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowExecutionEvent.Merge(dst, src)
}
func (m *WorkflowExecutionEvent) XXX_Size() int {
	return xxx_messageInfo_WorkflowExecutionEvent.Size(m)
}
func (m *WorkflowExecutionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowExecutionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowExecutionEvent proto.InternalMessageInfo

func (m *WorkflowExecutionEvent) GetExecutionId() *core.WorkflowExecutionIdentifier {
	if m != nil {
		return m.ExecutionId
	}
	return nil
}

func (m *WorkflowExecutionEvent) GetProducerId() string {
	if m != nil {
		return m.ProducerId
	}
	return ""
}

func (m *WorkflowExecutionEvent) GetPhase() core.WorkflowExecutionPhase {
	if m != nil {
		return m.Phase
	}
	return core.WorkflowExecutionPhase_WORKFLOW_PHASE_UNDEFINED
}

func (m *WorkflowExecutionEvent) GetOccurredAt() *timestamp.Timestamp {
	if m != nil {
		return m.OccurredAt
	}
	return nil
}

type isWorkflowExecutionEvent_OutputResult interface {
	isWorkflowExecutionEvent_OutputResult()
}

type WorkflowExecutionEvent_OutputUri struct {
	OutputUri string `protobuf:"bytes,5,opt,name=output_uri,json=outputUri,proto3,oneof"`
}

type WorkflowExecutionEvent_Error struct {
	Error *core.ExecutionError `protobuf:"bytes,6,opt,name=error,proto3,oneof"`
}

func (*WorkflowExecutionEvent_OutputUri) isWorkflowExecutionEvent_OutputResult() {}

func (*WorkflowExecutionEvent_Error) isWorkflowExecutionEvent_OutputResult() {}

func (m *WorkflowExecutionEvent) GetOutputResult() isWorkflowExecutionEvent_OutputResult {
	if m != nil {
		return m.OutputResult
	}
	return nil
}

func (m *WorkflowExecutionEvent) GetOutputUri() string {
	if x, ok := m.GetOutputResult().(*WorkflowExecutionEvent_OutputUri); ok {
		return x.OutputUri
	}
	return ""
}

func (m *WorkflowExecutionEvent) GetError() *core.ExecutionError {
	if x, ok := m.GetOutputResult().(*WorkflowExecutionEvent_Error); ok {
		return x.Error
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*WorkflowExecutionEvent) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _WorkflowExecutionEvent_OneofMarshaler, _WorkflowExecutionEvent_OneofUnmarshaler, _WorkflowExecutionEvent_OneofSizer, []interface{}{
		(*WorkflowExecutionEvent_OutputUri)(nil),
		(*WorkflowExecutionEvent_Error)(nil),
	}
}

func _WorkflowExecutionEvent_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*WorkflowExecutionEvent)
	// output_result
	switch x := m.OutputResult.(type) {
	case *WorkflowExecutionEvent_OutputUri:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.OutputUri)
	case *WorkflowExecutionEvent_Error:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("WorkflowExecutionEvent.OutputResult has unexpected type %T", x)
	}
	return nil
}

func _WorkflowExecutionEvent_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*WorkflowExecutionEvent)
	switch tag {
	case 5: // output_result.output_uri
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.OutputResult = &WorkflowExecutionEvent_OutputUri{x}
		return true, err
	case 6: // output_result.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(core.ExecutionError)
		err := b.DecodeMessage(msg)
		m.OutputResult = &WorkflowExecutionEvent_Error{msg}
		return true, err
	default:
		return false, nil
	}
}

func _WorkflowExecutionEvent_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*WorkflowExecutionEvent)
	// output_result
	switch x := m.OutputResult.(type) {
	case *WorkflowExecutionEvent_OutputUri:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.OutputUri)))
		n += len(x.OutputUri)
	case *WorkflowExecutionEvent_Error:
		s := proto.Size(x.Error)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type NodeExecutionEvent struct {
	// Unique identifier for this node execution
	Id *core.NodeExecutionIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// the id of the originator (Propeller) of the event
	ProducerId string                  `protobuf:"bytes,2,opt,name=producer_id,json=producerId,proto3" json:"producer_id,omitempty"`
	Phase      core.NodeExecutionPhase `protobuf:"varint,3,opt,name=phase,proto3,enum=flyteidl.core.NodeExecutionPhase" json:"phase,omitempty"`
	// This timestamp represents when the original event occurred, it is generated
	// by the executor of the node.
	OccurredAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	InputUri   string               `protobuf:"bytes,5,opt,name=input_uri,json=inputUri,proto3" json:"input_uri,omitempty"`
	// Types that are valid to be assigned to OutputResult:
	//	*NodeExecutionEvent_OutputUri
	//	*NodeExecutionEvent_Error
	OutputResult         isNodeExecutionEvent_OutputResult `protobuf_oneof:"output_result"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *NodeExecutionEvent) Reset()         { *m = NodeExecutionEvent{} }
func (m *NodeExecutionEvent) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionEvent) ProtoMessage()    {}
func (*NodeExecutionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_7f0ab3e135fe4708, []int{1}
}
func (m *NodeExecutionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionEvent.Unmarshal(m, b)
}
func (m *NodeExecutionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionEvent.Marshal(b, m, deterministic)
}
func (dst *NodeExecutionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionEvent.Merge(dst, src)
}
func (m *NodeExecutionEvent) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionEvent.Size(m)
}
func (m *NodeExecutionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionEvent proto.InternalMessageInfo

func (m *NodeExecutionEvent) GetId() *core.NodeExecutionIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *NodeExecutionEvent) GetProducerId() string {
	if m != nil {
		return m.ProducerId
	}
	return ""
}

func (m *NodeExecutionEvent) GetPhase() core.NodeExecutionPhase {
	if m != nil {
		return m.Phase
	}
	return core.NodeExecutionPhase_NODE_PHASE_UNDEFINED
}

func (m *NodeExecutionEvent) GetOccurredAt() *timestamp.Timestamp {
	if m != nil {
		return m.OccurredAt
	}
	return nil
}

func (m *NodeExecutionEvent) GetInputUri() string {
	if m != nil {
		return m.InputUri
	}
	return ""
}

type isNodeExecutionEvent_OutputResult interface {
	isNodeExecutionEvent_OutputResult()
}

type NodeExecutionEvent_OutputUri struct {
	OutputUri string `protobuf:"bytes,6,opt,name=output_uri,json=outputUri,proto3,oneof"`
}

type NodeExecutionEvent_Error struct {
	Error *core.ExecutionError `protobuf:"bytes,7,opt,name=error,proto3,oneof"`
}

func (*NodeExecutionEvent_OutputUri) isNodeExecutionEvent_OutputResult() {}

func (*NodeExecutionEvent_Error) isNodeExecutionEvent_OutputResult() {}

func (m *NodeExecutionEvent) GetOutputResult() isNodeExecutionEvent_OutputResult {
	if m != nil {
		return m.OutputResult
	}
	return nil
}

func (m *NodeExecutionEvent) GetOutputUri() string {
	if x, ok := m.GetOutputResult().(*NodeExecutionEvent_OutputUri); ok {
		return x.OutputUri
	}
	return ""
}

func (m *NodeExecutionEvent) GetError() *core.ExecutionError {
	if x, ok := m.GetOutputResult().(*NodeExecutionEvent_Error); ok {
		return x.Error
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*NodeExecutionEvent) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _NodeExecutionEvent_OneofMarshaler, _NodeExecutionEvent_OneofUnmarshaler, _NodeExecutionEvent_OneofSizer, []interface{}{
		(*NodeExecutionEvent_OutputUri)(nil),
		(*NodeExecutionEvent_Error)(nil),
	}
}

func _NodeExecutionEvent_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*NodeExecutionEvent)
	// output_result
	switch x := m.OutputResult.(type) {
	case *NodeExecutionEvent_OutputUri:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.OutputUri)
	case *NodeExecutionEvent_Error:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("NodeExecutionEvent.OutputResult has unexpected type %T", x)
	}
	return nil
}

func _NodeExecutionEvent_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*NodeExecutionEvent)
	switch tag {
	case 6: // output_result.output_uri
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.OutputResult = &NodeExecutionEvent_OutputUri{x}
		return true, err
	case 7: // output_result.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(core.ExecutionError)
		err := b.DecodeMessage(msg)
		m.OutputResult = &NodeExecutionEvent_Error{msg}
		return true, err
	default:
		return false, nil
	}
}

func _NodeExecutionEvent_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*NodeExecutionEvent)
	// output_result
	switch x := m.OutputResult.(type) {
	case *NodeExecutionEvent_OutputUri:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.OutputUri)))
		n += len(x.OutputUri)
	case *NodeExecutionEvent_Error:
		s := proto.Size(x.Error)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Plugin specific execution event information. For tasks like Python, Hive, Spark, DynamicJob.
type TaskExecutionEvent struct {
	// ID of the task. In combination with the retryAttempt this will indicate
	// the task execution uniquely for a given parent node execution.
	TaskId *core.Identifier `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// A task execution is always kicked off by a node execution, the event consumer
	// will use the parent_id to relate the task to it's parent node execution
	ParentNodeExecutionId *core.NodeExecutionIdentifier `protobuf:"bytes,2,opt,name=parent_node_execution_id,json=parentNodeExecutionId,proto3" json:"parent_node_execution_id,omitempty"`
	// retry attempt number for this task, ie., 2 for the second attempt
	RetryAttempt uint32 `protobuf:"varint,3,opt,name=retry_attempt,json=retryAttempt,proto3" json:"retry_attempt,omitempty"`
	// Phase associated with the event
	Phase core.TaskExecutionPhase `protobuf:"varint,4,opt,name=phase,proto3,enum=flyteidl.core.TaskExecutionPhase" json:"phase,omitempty"`
	// id of the process that sent this event, mainly for trace debugging
	ProducerId string `protobuf:"bytes,5,opt,name=producer_id,json=producerId,proto3" json:"producer_id,omitempty"`
	// log information for the task execution
	Logs []*core.TaskLog `protobuf:"bytes,6,rep,name=logs,proto3" json:"logs,omitempty"`
	// This timestamp represents when the original event occurred, it is generated
	// by the executor of the task.
	OccurredAt *timestamp.Timestamp `protobuf:"bytes,7,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	// URI of the input file, it encodes all the information
	// including Cloud source provider. ie., s3://...
	InputUri string `protobuf:"bytes,8,opt,name=input_uri,json=inputUri,proto3" json:"input_uri,omitempty"`
	// Types that are valid to be assigned to OutputResult:
	//	*TaskExecutionEvent_OutputUri
	//	*TaskExecutionEvent_Error
	OutputResult isTaskExecutionEvent_OutputResult `protobuf_oneof:"output_result"`
	// Custom data that the task plugin sends back. This is extensible to allow various plugins in the system.
	CustomInfo           *_struct.Struct `protobuf:"bytes,11,opt,name=custom_info,json=customInfo,proto3" json:"custom_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TaskExecutionEvent) Reset()         { *m = TaskExecutionEvent{} }
func (m *TaskExecutionEvent) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionEvent) ProtoMessage()    {}
func (*TaskExecutionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_7f0ab3e135fe4708, []int{2}
}
func (m *TaskExecutionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionEvent.Unmarshal(m, b)
}
func (m *TaskExecutionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionEvent.Marshal(b, m, deterministic)
}
func (dst *TaskExecutionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionEvent.Merge(dst, src)
}
func (m *TaskExecutionEvent) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionEvent.Size(m)
}
func (m *TaskExecutionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionEvent proto.InternalMessageInfo

func (m *TaskExecutionEvent) GetTaskId() *core.Identifier {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *TaskExecutionEvent) GetParentNodeExecutionId() *core.NodeExecutionIdentifier {
	if m != nil {
		return m.ParentNodeExecutionId
	}
	return nil
}

func (m *TaskExecutionEvent) GetRetryAttempt() uint32 {
	if m != nil {
		return m.RetryAttempt
	}
	return 0
}

func (m *TaskExecutionEvent) GetPhase() core.TaskExecutionPhase {
	if m != nil {
		return m.Phase
	}
	return core.TaskExecutionPhase_TASK_PHASE_UNDEFINED
}

func (m *TaskExecutionEvent) GetProducerId() string {
	if m != nil {
		return m.ProducerId
	}
	return ""
}

func (m *TaskExecutionEvent) GetLogs() []*core.TaskLog {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *TaskExecutionEvent) GetOccurredAt() *timestamp.Timestamp {
	if m != nil {
		return m.OccurredAt
	}
	return nil
}

func (m *TaskExecutionEvent) GetInputUri() string {
	if m != nil {
		return m.InputUri
	}
	return ""
}

type isTaskExecutionEvent_OutputResult interface {
	isTaskExecutionEvent_OutputResult()
}

type TaskExecutionEvent_OutputUri struct {
	OutputUri string `protobuf:"bytes,9,opt,name=output_uri,json=outputUri,proto3,oneof"`
}

type TaskExecutionEvent_Error struct {
	Error *core.ExecutionError `protobuf:"bytes,10,opt,name=error,proto3,oneof"`
}

func (*TaskExecutionEvent_OutputUri) isTaskExecutionEvent_OutputResult() {}

func (*TaskExecutionEvent_Error) isTaskExecutionEvent_OutputResult() {}

func (m *TaskExecutionEvent) GetOutputResult() isTaskExecutionEvent_OutputResult {
	if m != nil {
		return m.OutputResult
	}
	return nil
}

func (m *TaskExecutionEvent) GetOutputUri() string {
	if x, ok := m.GetOutputResult().(*TaskExecutionEvent_OutputUri); ok {
		return x.OutputUri
	}
	return ""
}

func (m *TaskExecutionEvent) GetError() *core.ExecutionError {
	if x, ok := m.GetOutputResult().(*TaskExecutionEvent_Error); ok {
		return x.Error
	}
	return nil
}

func (m *TaskExecutionEvent) GetCustomInfo() *_struct.Struct {
	if m != nil {
		return m.CustomInfo
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TaskExecutionEvent) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TaskExecutionEvent_OneofMarshaler, _TaskExecutionEvent_OneofUnmarshaler, _TaskExecutionEvent_OneofSizer, []interface{}{
		(*TaskExecutionEvent_OutputUri)(nil),
		(*TaskExecutionEvent_Error)(nil),
	}
}

func _TaskExecutionEvent_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TaskExecutionEvent)
	// output_result
	switch x := m.OutputResult.(type) {
	case *TaskExecutionEvent_OutputUri:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.OutputUri)
	case *TaskExecutionEvent_Error:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TaskExecutionEvent.OutputResult has unexpected type %T", x)
	}
	return nil
}

func _TaskExecutionEvent_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TaskExecutionEvent)
	switch tag {
	case 9: // output_result.output_uri
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.OutputResult = &TaskExecutionEvent_OutputUri{x}
		return true, err
	case 10: // output_result.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(core.ExecutionError)
		err := b.DecodeMessage(msg)
		m.OutputResult = &TaskExecutionEvent_Error{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TaskExecutionEvent_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TaskExecutionEvent)
	// output_result
	switch x := m.OutputResult.(type) {
	case *TaskExecutionEvent_OutputUri:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.OutputUri)))
		n += len(x.OutputUri)
	case *TaskExecutionEvent_Error:
		s := proto.Size(x.Error)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*WorkflowExecutionEvent)(nil), "flyteidl.event.WorkflowExecutionEvent")
	proto.RegisterType((*NodeExecutionEvent)(nil), "flyteidl.event.NodeExecutionEvent")
	proto.RegisterType((*TaskExecutionEvent)(nil), "flyteidl.event.TaskExecutionEvent")
}

func init() { proto.RegisterFile("flyteidl/event/event.proto", fileDescriptor_event_7f0ab3e135fe4708) }

var fileDescriptor_event_7f0ab3e135fe4708 = []byte{
	// 573 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x6f, 0x6b, 0xd3, 0x5e,
	0x14, 0xc7, 0xd7, 0xac, 0x7f, 0xd6, 0x93, 0x75, 0x3f, 0xb8, 0xf0, 0x9b, 0xb1, 0x3a, 0x5b, 0x27,
	0x4a, 0x19, 0x98, 0x60, 0x87, 0x7f, 0x60, 0x8f, 0x36, 0x18, 0xac, 0xa0, 0x22, 0x71, 0x22, 0xf8,
	0x24, 0xa4, 0xc9, 0x49, 0x76, 0x69, 0x9a, 0x1b, 0x6e, 0x4e, 0xd4, 0xbe, 0x4c, 0x5f, 0x86, 0xf8,
	0x26, 0x24, 0x49, 0x93, 0x2e, 0xe9, 0xa8, 0x4e, 0x7d, 0x52, 0xe8, 0xb9, 0xdf, 0x7b, 0xee, 0xf7,
	0x7e, 0xce, 0x37, 0x17, 0xfa, 0x5e, 0xb0, 0x20, 0xe4, 0x6e, 0x60, 0xe0, 0x67, 0x0c, 0x29, 0xff,
	0xd5, 0x23, 0x29, 0x48, 0xb0, 0xbd, 0x62, 0x4d, 0xcf, 0xaa, 0xfd, 0x83, 0x52, 0xeb, 0x08, 0x89,
	0x06, 0x7e, 0x45, 0x27, 0x21, 0x2e, 0xc2, 0x5c, 0xde, 0x7f, 0x50, 0x5d, 0xe6, 0x2e, 0x86, 0xc4,
	0x3d, 0x8e, 0x72, 0xb9, 0x3e, 0xf0, 0x85, 0xf0, 0x03, 0x34, 0xb2, 0x7f, 0xd3, 0xc4, 0x33, 0x88,
	0xcf, 0x31, 0x26, 0x7b, 0x1e, 0x2d, 0x05, 0xf7, 0xeb, 0x82, 0x98, 0x64, 0xe2, 0x2c, 0xdd, 0x1c,
	0xfe, 0x50, 0x60, 0xff, 0xa3, 0x90, 0x33, 0x2f, 0x10, 0x5f, 0xce, 0x8b, 0xa3, 0xcf, 0x53, 0x63,
	0xec, 0x0d, 0xec, 0x96, 0x66, 0x2c, 0xee, 0x6a, 0x8d, 0x61, 0x63, 0xa4, 0x8e, 0x8f, 0xf4, 0xd2,
	0x7f, 0x6a, 0x48, 0x5f, 0xdb, 0x3c, 0x29, 0x1d, 0x9a, 0x2a, 0xae, 0x8a, 0x6c, 0x00, 0x6a, 0x24,
	0x85, 0x9b, 0x38, 0x28, 0xd3, 0x6e, 0xca, 0xb0, 0x31, 0xea, 0x9a, 0x50, 0x94, 0x26, 0x2e, 0x3b,
	0x81, 0x56, 0x74, 0x65, 0xc7, 0xa8, 0x6d, 0x0f, 0x1b, 0xa3, 0xbd, 0xf1, 0xe3, 0x5f, 0x1d, 0xf4,
	0x2e, 0x15, 0x9b, 0xf9, 0x1e, 0x76, 0x02, 0xaa, 0x70, 0x9c, 0x44, 0x4a, 0x74, 0x2d, 0x9b, 0xb4,
	0x66, 0xe6, 0xb5, 0xaf, 0xe7, 0x77, 0xd7, 0x8b, 0xbb, 0xeb, 0x97, 0x05, 0x1c, 0x13, 0x0a, 0xf9,
	0x29, 0xb1, 0x01, 0x80, 0x48, 0x28, 0x4a, 0xc8, 0x4a, 0x24, 0xd7, 0x5a, 0xa9, 0xb3, 0x8b, 0x2d,
	0xb3, 0x9b, 0xd7, 0x3e, 0x48, 0xce, 0x9e, 0x43, 0x0b, 0xa5, 0x14, 0x52, 0x6b, 0x67, 0x7d, 0x0f,
	0x6a, 0xd6, 0x56, 0xe0, 0x52, 0xd1, 0xc5, 0x96, 0x99, 0xab, 0xcf, 0xfe, 0x83, 0xde, 0xb2, 0xaf,
	0xc4, 0x38, 0x09, 0xe8, 0xf0, 0xbb, 0x02, 0xec, 0xad, 0x70, 0xb1, 0x46, 0xfa, 0x05, 0x28, 0x25,
	0xdf, 0x27, 0xb5, 0xde, 0x15, 0xf9, 0x35, 0xb6, 0x0a, 0xff, 0x0d, 0xa4, 0x2f, 0xab, 0x48, 0x1f,
	0x6e, 0xea, 0xfd, 0xef, 0x70, 0xde, 0x83, 0x2e, 0x0f, 0x2b, 0x34, 0xcd, 0x9d, 0xac, 0x90, 0xa2,
	0xac, 0xb2, 0x6e, 0x6f, 0x60, 0xdd, 0xf9, 0x3b, 0xd6, 0xdf, 0x9a, 0xc0, 0x2e, 0xed, 0x78, 0x56,
	0x63, 0x3d, 0x86, 0x0e, 0xd9, 0xf1, 0x6c, 0x15, 0xe8, 0xbb, 0xb5, 0x03, 0xae, 0x31, 0x6e, 0xa7,
	0xca, 0x89, 0xcb, 0x2c, 0xd0, 0x22, 0x5b, 0x62, 0x48, 0x56, 0x28, 0x5c, 0xb4, 0x2a, 0x5f, 0x85,
	0x72, 0xab, 0xa9, 0xfd, 0x9f, 0xf7, 0xa9, 0x2d, 0xb3, 0x47, 0xd0, 0x93, 0x48, 0x72, 0x61, 0xd9,
	0x44, 0x38, 0x8f, 0x28, 0x9b, 0x57, 0xcf, 0xdc, 0xcd, 0x8a, 0xa7, 0x79, 0x6d, 0x35, 0xcc, 0xe6,
	0x8d, 0xc3, 0xac, 0xdc, 0xb5, 0x32, 0xcc, 0x5a, 0x4c, 0x5a, 0x6b, 0x31, 0x39, 0x82, 0x66, 0x20,
	0xfc, 0x58, 0x6b, 0x0f, 0xb7, 0x47, 0xea, 0x78, 0xff, 0x86, 0xc6, 0xaf, 0x85, 0x6f, 0x66, 0x9a,
	0x7a, 0x32, 0x3a, 0x7f, 0x9e, 0x8c, 0x9d, 0x8d, 0xc9, 0xe8, 0x6e, 0x48, 0x06, 0xdc, 0x26, 0x19,
	0xec, 0x15, 0xa8, 0x4e, 0x12, 0x93, 0x98, 0x5b, 0x3c, 0xf4, 0x84, 0xa6, 0x66, 0x9b, 0xef, 0xac,
	0x39, 0x7e, 0x9f, 0x3d, 0x8b, 0x26, 0xe4, 0xda, 0x49, 0xe8, 0x89, 0xb5, 0x4c, 0x9d, 0x1d, 0x7f,
	0x7a, 0xe6, 0x73, 0xba, 0x4a, 0xa6, 0xba, 0x23, 0xe6, 0x46, 0xb0, 0xf0, 0xc8, 0x28, 0x9f, 0x67,
	0x1f, 0x43, 0x23, 0x9a, 0x3e, 0xf5, 0x85, 0x51, 0x7d, 0xfc, 0xa7, 0xed, 0xec, 0x88, 0xe3, 0x9f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x9c, 0xcf, 0x79, 0x15, 0x06, 0x00, 0x00,
}
