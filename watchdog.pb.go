// Code generated by protoc-gen-go. DO NOT EDIT.
// source: watchdog.proto

package main

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Watch_Frequency int32

const (
	Watch_DAILY  Watch_Frequency = 0
	Watch_WEEKLY Watch_Frequency = 1
)

var Watch_Frequency_name = map[int32]string{
	0: "DAILY",
	1: "WEEKLY",
}

var Watch_Frequency_value = map[string]int32{
	"DAILY":  0,
	"WEEKLY": 1,
}

func (x Watch_Frequency) String() string {
	return proto.EnumName(Watch_Frequency_name, int32(x))
}

func (Watch_Frequency) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_554efde3f2d36ab7, []int{0, 0}
}

type Watch struct {
	Name                 string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Frequency            Watch_Frequency `protobuf:"varint,2,opt,name=frequency,proto3,enum=watchdog.Watch_Frequency" json:"frequency,omitempty"`
	LastSeen             int64           `protobuf:"varint,3,opt,name=LastSeen,proto3" json:"LastSeen,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Watch) Reset()         { *m = Watch{} }
func (m *Watch) String() string { return proto.CompactTextString(m) }
func (*Watch) ProtoMessage()    {}
func (*Watch) Descriptor() ([]byte, []int) {
	return fileDescriptor_554efde3f2d36ab7, []int{0}
}

func (m *Watch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Watch.Unmarshal(m, b)
}
func (m *Watch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Watch.Marshal(b, m, deterministic)
}
func (m *Watch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Watch.Merge(m, src)
}
func (m *Watch) XXX_Size() int {
	return xxx_messageInfo_Watch.Size(m)
}
func (m *Watch) XXX_DiscardUnknown() {
	xxx_messageInfo_Watch.DiscardUnknown(m)
}

var xxx_messageInfo_Watch proto.InternalMessageInfo

func (m *Watch) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Watch) GetFrequency() Watch_Frequency {
	if m != nil {
		return m.Frequency
	}
	return Watch_DAILY
}

func (m *Watch) GetLastSeen() int64 {
	if m != nil {
		return m.LastSeen
	}
	return 0
}

func init() {
	proto.RegisterEnum("watchdog.Watch_Frequency", Watch_Frequency_name, Watch_Frequency_value)
	proto.RegisterType((*Watch)(nil), "watchdog.Watch")
}

func init() { proto.RegisterFile("watchdog.proto", fileDescriptor_554efde3f2d36ab7) }

var fileDescriptor_554efde3f2d36ab7 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4f, 0x2c, 0x49,
	0xce, 0x48, 0xc9, 0x4f, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0xa6,
	0x30, 0x72, 0xb1, 0x86, 0x83, 0x38, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x39, 0x17, 0x67, 0x5a, 0x51, 0x6a, 0x61, 0x69,
	0x6a, 0x5e, 0x72, 0xa5, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x9f, 0x91, 0xa4, 0x1e, 0xdc, 0x2c, 0xb0,
	0x3e, 0x3d, 0x37, 0x98, 0x82, 0x20, 0x84, 0x5a, 0x21, 0x29, 0x2e, 0x0e, 0x9f, 0xc4, 0xe2, 0x92,
	0xe0, 0xd4, 0xd4, 0x3c, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x38, 0x5f, 0x49, 0x89, 0x8b,
	0x13, 0xae, 0x47, 0x88, 0x93, 0x8b, 0xd5, 0xc5, 0xd1, 0xd3, 0x27, 0x52, 0x80, 0x41, 0x88, 0x8b,
	0x8b, 0x2d, 0xdc, 0xd5, 0xd5, 0xdb, 0x27, 0x52, 0x80, 0xd1, 0x89, 0x2d, 0x8a, 0x25, 0x37, 0x31,
	0x33, 0x2f, 0x89, 0x0d, 0xec, 0x5e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x1b, 0xb7,
	0xb3, 0xc1, 0x00, 0x00, 0x00,
}
