// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/plugins/spark.proto

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

type SparkApplication_Type int32

const (
	SparkApplication_PYTHON SparkApplication_Type = 0
	SparkApplication_JAVA   SparkApplication_Type = 1
	SparkApplication_SCALA  SparkApplication_Type = 2
	SparkApplication_R      SparkApplication_Type = 3
)

var SparkApplication_Type_name = map[int32]string{
	0: "PYTHON",
	1: "JAVA",
	2: "SCALA",
	3: "R",
}

var SparkApplication_Type_value = map[string]int32{
	"PYTHON": 0,
	"JAVA":   1,
	"SCALA":  2,
	"R":      3,
}

func (x SparkApplication_Type) String() string {
	return proto.EnumName(SparkApplication_Type_name, int32(x))
}

func (SparkApplication_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ca8a069b9820144a, []int{0, 0}
}

type SparkApplication struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SparkApplication) Reset()         { *m = SparkApplication{} }
func (m *SparkApplication) String() string { return proto.CompactTextString(m) }
func (*SparkApplication) ProtoMessage()    {}
func (*SparkApplication) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca8a069b9820144a, []int{0}
}

func (m *SparkApplication) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SparkApplication.Unmarshal(m, b)
}
func (m *SparkApplication) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SparkApplication.Marshal(b, m, deterministic)
}
func (m *SparkApplication) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SparkApplication.Merge(m, src)
}
func (m *SparkApplication) XXX_Size() int {
	return xxx_messageInfo_SparkApplication.Size(m)
}
func (m *SparkApplication) XXX_DiscardUnknown() {
	xxx_messageInfo_SparkApplication.DiscardUnknown(m)
}

var xxx_messageInfo_SparkApplication proto.InternalMessageInfo

// Custom Proto for Spark Plugin.
type SparkJob struct {
	ApplicationType      SparkApplication_Type `protobuf:"varint,1,opt,name=applicationType,proto3,enum=flyteidl.plugins.SparkApplication_Type" json:"applicationType,omitempty"`
	MainApplicationFile  string                `protobuf:"bytes,2,opt,name=mainApplicationFile,proto3" json:"mainApplicationFile,omitempty"`
	MainClass            string                `protobuf:"bytes,3,opt,name=mainClass,proto3" json:"mainClass,omitempty"`
	SparkConf            map[string]string     `protobuf:"bytes,4,rep,name=sparkConf,proto3" json:"sparkConf,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	HadoopConf           map[string]string     `protobuf:"bytes,5,rep,name=hadoopConf,proto3" json:"hadoopConf,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ExecutorPath         string                `protobuf:"bytes,6,opt,name=executorPath,proto3" json:"executorPath,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SparkJob) Reset()         { *m = SparkJob{} }
func (m *SparkJob) String() string { return proto.CompactTextString(m) }
func (*SparkJob) ProtoMessage()    {}
func (*SparkJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca8a069b9820144a, []int{1}
}

func (m *SparkJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SparkJob.Unmarshal(m, b)
}
func (m *SparkJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SparkJob.Marshal(b, m, deterministic)
}
func (m *SparkJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SparkJob.Merge(m, src)
}
func (m *SparkJob) XXX_Size() int {
	return xxx_messageInfo_SparkJob.Size(m)
}
func (m *SparkJob) XXX_DiscardUnknown() {
	xxx_messageInfo_SparkJob.DiscardUnknown(m)
}

var xxx_messageInfo_SparkJob proto.InternalMessageInfo

func (m *SparkJob) GetApplicationType() SparkApplication_Type {
	if m != nil {
		return m.ApplicationType
	}
	return SparkApplication_PYTHON
}

func (m *SparkJob) GetMainApplicationFile() string {
	if m != nil {
		return m.MainApplicationFile
	}
	return ""
}

func (m *SparkJob) GetMainClass() string {
	if m != nil {
		return m.MainClass
	}
	return ""
}

func (m *SparkJob) GetSparkConf() map[string]string {
	if m != nil {
		return m.SparkConf
	}
	return nil
}

func (m *SparkJob) GetHadoopConf() map[string]string {
	if m != nil {
		return m.HadoopConf
	}
	return nil
}

func (m *SparkJob) GetExecutorPath() string {
	if m != nil {
		return m.ExecutorPath
	}
	return ""
}

func init() {
	proto.RegisterEnum("flyteidl.plugins.SparkApplication_Type", SparkApplication_Type_name, SparkApplication_Type_value)
	proto.RegisterType((*SparkApplication)(nil), "flyteidl.plugins.SparkApplication")
	proto.RegisterType((*SparkJob)(nil), "flyteidl.plugins.SparkJob")
	proto.RegisterMapType((map[string]string)(nil), "flyteidl.plugins.SparkJob.HadoopConfEntry")
	proto.RegisterMapType((map[string]string)(nil), "flyteidl.plugins.SparkJob.SparkConfEntry")
}

func init() { proto.RegisterFile("flyteidl/plugins/spark.proto", fileDescriptor_ca8a069b9820144a) }

var fileDescriptor_ca8a069b9820144a = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x4f, 0xea, 0x40,
	0x14, 0xc5, 0x5f, 0x29, 0x25, 0xf4, 0xbe, 0x17, 0x68, 0xe6, 0xb9, 0x68, 0x08, 0x0b, 0xd2, 0x8d,
	0x68, 0x62, 0x6b, 0x70, 0xe1, 0x9f, 0xe8, 0xa2, 0x10, 0x95, 0x34, 0x46, 0xb1, 0x10, 0x13, 0xdd,
	0x4d, 0xa1, 0x94, 0x86, 0xa1, 0x33, 0x69, 0xa7, 0xc6, 0x7e, 0x5e, 0xbf, 0x88, 0xe9, 0x20, 0xff,
	0x1a, 0x35, 0x71, 0x37, 0xbd, 0xf7, 0x9c, 0xdf, 0x3d, 0xbd, 0xb9, 0xd0, 0x9c, 0x92, 0x8c, 0xfb,
	0xe1, 0x84, 0x58, 0x8c, 0xa4, 0x41, 0x18, 0x25, 0x56, 0xc2, 0x70, 0x3c, 0x37, 0x59, 0x4c, 0x39,
	0x45, 0xda, 0xaa, 0x6b, 0x7e, 0x76, 0x8d, 0x2e, 0x68, 0xc3, 0x5c, 0x60, 0x33, 0x46, 0xc2, 0x31,
	0xe6, 0x21, 0x8d, 0x0c, 0x13, 0xca, 0xa3, 0x8c, 0xf9, 0x08, 0xa0, 0x32, 0x78, 0x1e, 0xf5, 0x1f,
	0xee, 0xb5, 0x3f, 0xa8, 0x0a, 0x65, 0xc7, 0x7e, 0xb2, 0x35, 0x09, 0xa9, 0xa0, 0x0c, 0x7b, 0xf6,
	0x9d, 0xad, 0x95, 0x90, 0x02, 0x92, 0xab, 0xc9, 0xc6, 0xbb, 0x0c, 0x55, 0x01, 0x71, 0xa8, 0x87,
	0x1e, 0xa1, 0x8e, 0x37, 0xac, 0x9c, 0xa3, 0x4b, 0x2d, 0xa9, 0x5d, 0xeb, 0xec, 0x9b, 0xc5, 0xe1,
	0x66, 0x71, 0xb2, 0x99, 0xcb, 0xdd, 0xa2, 0x1f, 0x1d, 0xc3, 0xff, 0x05, 0x0e, 0xa3, 0x2d, 0xe1,
	0x4d, 0x48, 0x7c, 0xbd, 0xd4, 0x92, 0xda, 0xaa, 0xfb, 0x55, 0x0b, 0x35, 0x41, 0xcd, 0xcb, 0x3d,
	0x82, 0x93, 0x44, 0x97, 0x85, 0x6e, 0x53, 0x40, 0xb7, 0xa0, 0x8a, 0xa5, 0xf4, 0x68, 0x34, 0xd5,
	0xcb, 0x2d, 0xb9, 0xfd, 0xb7, 0x73, 0xf0, 0x4d, 0x38, 0x87, 0x7a, 0xcb, 0x47, 0xae, 0xbd, 0x8e,
	0x78, 0x9c, 0xb9, 0x1b, 0x2f, 0x72, 0x00, 0x66, 0x78, 0x42, 0x29, 0x13, 0x24, 0x45, 0x90, 0x0e,
	0x7f, 0x20, 0xf5, 0xd7, 0xe2, 0x25, 0x6a, 0xcb, 0x8d, 0x0c, 0xf8, 0xe7, 0xbf, 0xf9, 0xe3, 0x94,
	0xd3, 0x78, 0x80, 0xf9, 0x4c, 0xaf, 0x88, 0xd4, 0x3b, 0xb5, 0xc6, 0x25, 0xd4, 0x76, 0xc3, 0x20,
	0x0d, 0xe4, 0xb9, 0x9f, 0x89, 0x0d, 0xab, 0x6e, 0xfe, 0x44, 0x7b, 0xa0, 0xbc, 0x62, 0x92, 0xae,
	0xd6, 0xb3, 0xfc, 0xb8, 0x28, 0x9d, 0x49, 0x8d, 0x2b, 0xa8, 0x17, 0x02, 0xfc, 0xc6, 0xde, 0x3d,
	0x7f, 0x39, 0x0d, 0x42, 0x3e, 0x4b, 0x3d, 0x73, 0x4c, 0x17, 0x96, 0xf8, 0x49, 0x1a, 0x07, 0xd6,
	0xfa, 0xde, 0x02, 0x3f, 0xb2, 0x98, 0x77, 0x14, 0x50, 0xab, 0x78, 0x82, 0x5e, 0x45, 0x5c, 0xdf,
	0xc9, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x94, 0x6f, 0xaa, 0x2d, 0x9d, 0x02, 0x00, 0x00,
}
