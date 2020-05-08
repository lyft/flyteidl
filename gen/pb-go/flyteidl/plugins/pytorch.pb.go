// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/plugins/pytorch.proto

package plugins

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

type PyTorchOperatorTask struct {
	Workers              int32    `protobuf:"varint,1,opt,name=workers,proto3" json:"workers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PyTorchOperatorTask) Reset()         { *m = PyTorchOperatorTask{} }
func (m *PyTorchOperatorTask) String() string { return proto.CompactTextString(m) }
func (*PyTorchOperatorTask) ProtoMessage()    {}
func (*PyTorchOperatorTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_4df8a9374b28b766, []int{0}
}

func (m *PyTorchOperatorTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PyTorchOperatorTask.Unmarshal(m, b)
}
func (m *PyTorchOperatorTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PyTorchOperatorTask.Marshal(b, m, deterministic)
}
func (m *PyTorchOperatorTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PyTorchOperatorTask.Merge(m, src)
}
func (m *PyTorchOperatorTask) XXX_Size() int {
	return xxx_messageInfo_PyTorchOperatorTask.Size(m)
}
func (m *PyTorchOperatorTask) XXX_DiscardUnknown() {
	xxx_messageInfo_PyTorchOperatorTask.DiscardUnknown(m)
}

var xxx_messageInfo_PyTorchOperatorTask proto.InternalMessageInfo

func (m *PyTorchOperatorTask) GetWorkers() int32 {
	if m != nil {
		return m.Workers
	}
	return 0
}

func init() {
	proto.RegisterType((*PyTorchOperatorTask)(nil), "flyteidl.plugins.PyTorchOperatorTask")
}

func init() { proto.RegisterFile("flyteidl/plugins/pytorch.proto", fileDescriptor_4df8a9374b28b766) }

var fileDescriptor_4df8a9374b28b766 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcb, 0xa9, 0x2c,
	0x49, 0xcd, 0x4c, 0xc9, 0xd1, 0x2f, 0xc8, 0x29, 0x4d, 0xcf, 0xcc, 0x2b, 0xd6, 0x2f, 0xa8, 0x2c,
	0xc9, 0x2f, 0x4a, 0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x80, 0xc9, 0xeb, 0x41,
	0xe5, 0x95, 0xf4, 0xb9, 0x84, 0x03, 0x2a, 0x43, 0x40, 0x4a, 0xfc, 0x0b, 0x52, 0x8b, 0x12, 0x4b,
	0xf2, 0x8b, 0x42, 0x12, 0x8b, 0xb3, 0x85, 0x24, 0xb8, 0xd8, 0xcb, 0xf3, 0x8b, 0xb2, 0x53, 0x8b,
	0x8a, 0x25, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x60, 0x5c, 0x27, 0xd3, 0x28, 0xe3, 0xf4, 0xcc,
	0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0x9c, 0xca, 0xb4, 0x12, 0x7d, 0xb8, 0xa5,
	0xe9, 0xa9, 0x79, 0xfa, 0x05, 0x49, 0xba, 0xe9, 0xf9, 0xfa, 0xe8, 0xee, 0x48, 0x62, 0x03, 0x3b,
	0xc0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x46, 0x28, 0x35, 0xa2, 0x00, 0x00, 0x00,
}
