// Code generated by protoc-gen-go. DO NOT EDIT.
// source: admin/task.proto

package admin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TaskCreateRequest struct {
	Id                   *Identifier `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Version              string      `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Spec                 *TaskSpec   `protobuf:"bytes,3,opt,name=spec" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TaskCreateRequest) Reset()         { *m = TaskCreateRequest{} }
func (m *TaskCreateRequest) String() string { return proto.CompactTextString(m) }
func (*TaskCreateRequest) ProtoMessage()    {}
func (*TaskCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_0967cce0e4ca2be2, []int{0}
}
func (m *TaskCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskCreateRequest.Unmarshal(m, b)
}
func (m *TaskCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskCreateRequest.Marshal(b, m, deterministic)
}
func (dst *TaskCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskCreateRequest.Merge(dst, src)
}
func (m *TaskCreateRequest) XXX_Size() int {
	return xxx_messageInfo_TaskCreateRequest.Size(m)
}
func (m *TaskCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskCreateRequest proto.InternalMessageInfo

func (m *TaskCreateRequest) GetId() *Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *TaskCreateRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *TaskCreateRequest) GetSpec() *TaskSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type TaskCreateResponse struct {
	Urn                  string   `protobuf:"bytes,1,opt,name=urn" json:"urn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskCreateResponse) Reset()         { *m = TaskCreateResponse{} }
func (m *TaskCreateResponse) String() string { return proto.CompactTextString(m) }
func (*TaskCreateResponse) ProtoMessage()    {}
func (*TaskCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_0967cce0e4ca2be2, []int{1}
}
func (m *TaskCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskCreateResponse.Unmarshal(m, b)
}
func (m *TaskCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskCreateResponse.Marshal(b, m, deterministic)
}
func (dst *TaskCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskCreateResponse.Merge(dst, src)
}
func (m *TaskCreateResponse) XXX_Size() int {
	return xxx_messageInfo_TaskCreateResponse.Size(m)
}
func (m *TaskCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskCreateResponse proto.InternalMessageInfo

func (m *TaskCreateResponse) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

type Task struct {
	Id                   *Identifier `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Version              string      `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Urn                  string      `protobuf:"bytes,3,opt,name=urn" json:"urn,omitempty"`
	Spec                 *TaskSpec   `protobuf:"bytes,6,opt,name=spec" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_0967cce0e4ca2be2, []int{2}
}
func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (dst *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(dst, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetId() *Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Task) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Task) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *Task) GetSpec() *TaskSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type TaskList struct {
	Tasks                []*Task  `protobuf:"bytes,1,rep,name=tasks" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskList) Reset()         { *m = TaskList{} }
func (m *TaskList) String() string { return proto.CompactTextString(m) }
func (*TaskList) ProtoMessage()    {}
func (*TaskList) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_0967cce0e4ca2be2, []int{3}
}
func (m *TaskList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskList.Unmarshal(m, b)
}
func (m *TaskList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskList.Marshal(b, m, deterministic)
}
func (dst *TaskList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskList.Merge(dst, src)
}
func (m *TaskList) XXX_Size() int {
	return xxx_messageInfo_TaskList.Size(m)
}
func (m *TaskList) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskList.DiscardUnknown(m)
}

var xxx_messageInfo_TaskList proto.InternalMessageInfo

func (m *TaskList) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type TaskSpec struct {
	Task                 *core.TaskTemplate `protobuf:"bytes,1,opt,name=task" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TaskSpec) Reset()         { *m = TaskSpec{} }
func (m *TaskSpec) String() string { return proto.CompactTextString(m) }
func (*TaskSpec) ProtoMessage()    {}
func (*TaskSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_0967cce0e4ca2be2, []int{4}
}
func (m *TaskSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskSpec.Unmarshal(m, b)
}
func (m *TaskSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskSpec.Marshal(b, m, deterministic)
}
func (dst *TaskSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskSpec.Merge(dst, src)
}
func (m *TaskSpec) XXX_Size() int {
	return xxx_messageInfo_TaskSpec.Size(m)
}
func (m *TaskSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskSpec.DiscardUnknown(m)
}

var xxx_messageInfo_TaskSpec proto.InternalMessageInfo

func (m *TaskSpec) GetTask() *core.TaskTemplate {
	if m != nil {
		return m.Task
	}
	return nil
}

func init() {
	proto.RegisterType((*TaskCreateRequest)(nil), "admin.TaskCreateRequest")
	proto.RegisterType((*TaskCreateResponse)(nil), "admin.TaskCreateResponse")
	proto.RegisterType((*Task)(nil), "admin.Task")
	proto.RegisterType((*TaskList)(nil), "admin.TaskList")
	proto.RegisterType((*TaskSpec)(nil), "admin.TaskSpec")
}

func init() { proto.RegisterFile("admin/task.proto", fileDescriptor_task_0967cce0e4ca2be2) }

var fileDescriptor_task_0967cce0e4ca2be2 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x51, 0x3d, 0x4f, 0xf3, 0x30,
	0x10, 0x56, 0x3e, 0xda, 0xbe, 0xbd, 0x0c, 0x6f, 0xea, 0xc9, 0xea, 0x14, 0x82, 0x54, 0x65, 0x21,
	0x48, 0xe1, 0x1f, 0xc0, 0x84, 0xc4, 0x64, 0x3a, 0xb1, 0x85, 0xe4, 0x90, 0xac, 0x12, 0xdb, 0xd8,
	0x0e, 0x03, 0xbf, 0x1e, 0xf9, 0x1c, 0xaa, 0x4e, 0x2c, 0x6c, 0xc9, 0xf3, 0x71, 0xf7, 0xf8, 0x1e,
	0x28, 0xfb, 0x71, 0x92, 0xea, 0xd6, 0xf7, 0xee, 0xd4, 0x1a, 0xab, 0xbd, 0x66, 0x2b, 0x42, 0xf6,
	0xe5, 0xa0, 0x2d, 0x12, 0xee, 0x22, 0xb1, 0x67, 0x51, 0x3a, 0xe8, 0x69, 0xd2, 0x2a, 0x62, 0xf5,
	0x0c, 0xbb, 0x63, 0xef, 0x4e, 0x0f, 0x16, 0x7b, 0x8f, 0x02, 0x3f, 0x66, 0x74, 0x9e, 0x5d, 0x41,
	0x2a, 0x47, 0x9e, 0x54, 0x49, 0x53, 0x74, 0xbb, 0x96, 0x5c, 0xed, 0xe3, 0x88, 0xca, 0xcb, 0x37,
	0x89, 0x56, 0xa4, 0x72, 0x64, 0x1c, 0x36, 0x9f, 0x68, 0x9d, 0xd4, 0x8a, 0xa7, 0x55, 0xd2, 0x6c,
	0xc5, 0xcf, 0x2f, 0xbb, 0x86, 0xdc, 0x19, 0x1c, 0x78, 0x46, 0xf6, 0xff, 0x8b, 0x3d, 0x2c, 0x79,
	0x36, 0x38, 0x08, 0x22, 0xeb, 0x03, 0xb0, 0xcb, 0xb5, 0xce, 0x68, 0xe5, 0x90, 0x95, 0x90, 0xcd,
	0x56, 0xd1, 0xe2, 0xad, 0x08, 0x9f, 0xf5, 0x17, 0xe4, 0x41, 0xf7, 0xb7, 0x44, 0xcb, 0xd8, 0xec,
	0x3c, 0xf6, 0x9c, 0x71, 0xfd, 0x5b, 0xc6, 0x1b, 0xf8, 0x17, 0x90, 0x27, 0x49, 0x17, 0x59, 0xd1,
	0x25, 0x79, 0x52, 0x65, 0x4d, 0xd1, 0x15, 0x17, 0x0e, 0x11, 0x99, 0xba, 0x8b, 0xf2, 0x30, 0x80,
	0x1d, 0x20, 0x0f, 0xe0, 0x12, 0x98, 0xb5, 0xa1, 0x0a, 0x12, 0x1f, 0x71, 0x32, 0xef, 0xe1, 0xc9,
	0xc4, 0xdf, 0x6f, 0x5e, 0x62, 0x59, 0xaf, 0x6b, 0x6a, 0xe3, 0xee, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0x15, 0xc2, 0xa5, 0x24, 0xce, 0x01, 0x00, 0x00,
}