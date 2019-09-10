// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/p2p/v1/message.proto

package v1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type Message struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Payload              []byte   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	From                 []byte   `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b86fd3d35f51f8e, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Message) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Message) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

type TestPing struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestPing) Reset()         { *m = TestPing{} }
func (m *TestPing) String() string { return proto.CompactTextString(m) }
func (*TestPing) ProtoMessage()    {}
func (*TestPing) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b86fd3d35f51f8e, []int{1}
}

func (m *TestPing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestPing.Unmarshal(m, b)
}
func (m *TestPing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestPing.Marshal(b, m, deterministic)
}
func (m *TestPing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestPing.Merge(m, src)
}
func (m *TestPing) XXX_Size() int {
	return xxx_messageInfo_TestPing.Size(m)
}
func (m *TestPing) XXX_DiscardUnknown() {
	xxx_messageInfo_TestPing.DiscardUnknown(m)
}

var xxx_messageInfo_TestPing proto.InternalMessageInfo

func (m *TestPing) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type TestPong struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestPong) Reset()         { *m = TestPong{} }
func (m *TestPong) String() string { return proto.CompactTextString(m) }
func (*TestPong) ProtoMessage()    {}
func (*TestPong) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b86fd3d35f51f8e, []int{2}
}

func (m *TestPong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestPong.Unmarshal(m, b)
}
func (m *TestPong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestPong.Marshal(b, m, deterministic)
}
func (m *TestPong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestPong.Merge(m, src)
}
func (m *TestPong) XXX_Size() int {
	return xxx_messageInfo_TestPong.Size(m)
}
func (m *TestPong) XXX_DiscardUnknown() {
	xxx_messageInfo_TestPong.DiscardUnknown(m)
}

var xxx_messageInfo_TestPong proto.InternalMessageInfo

func (m *TestPong) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "airbloc.p2p.v1.Message")
	proto.RegisterType((*TestPing)(nil), "airbloc.p2p.v1.TestPing")
	proto.RegisterType((*TestPong)(nil), "airbloc.p2p.v1.TestPong")
}

func init() { proto.RegisterFile("proto/p2p/v1/message.proto", fileDescriptor_0b86fd3d35f51f8e) }

var fileDescriptor_0b86fd3d35f51f8e = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x30, 0x2a, 0xd0, 0x2f, 0x33, 0xd4, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f,
	0xd5, 0x03, 0x0b, 0x0a, 0xf1, 0x25, 0x66, 0x16, 0x25, 0xe5, 0xe4, 0x27, 0xeb, 0x15, 0x18, 0x15,
	0xe8, 0x95, 0x19, 0x2a, 0xf9, 0x72, 0xb1, 0xfb, 0x42, 0x14, 0x08, 0x89, 0x70, 0xb1, 0x96, 0xe4,
	0x17, 0x64, 0x26, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x42, 0x12, 0x5c, 0xec,
	0x05, 0x89, 0x95, 0x39, 0xf9, 0x89, 0x29, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x30, 0xae,
	0x90, 0x10, 0x17, 0x4b, 0x5a, 0x51, 0x7e, 0xae, 0x04, 0x33, 0x58, 0x18, 0xcc, 0x56, 0x52, 0xe1,
	0xe2, 0x08, 0x49, 0x2d, 0x2e, 0x09, 0xc8, 0xcc, 0x4b, 0x07, 0xe9, 0x84, 0xda, 0x0d, 0x35, 0x11,
	0xc6, 0x85, 0xab, 0xca, 0xc7, 0xa7, 0xca, 0x49, 0x27, 0x4a, 0x2b, 0x3d, 0xb3, 0x24, 0xa3, 0x34,
	0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0xea, 0x6e, 0x18, 0xad, 0x9b, 0x9e, 0xaf, 0x8f, 0xec, 0xcd,
	0x24, 0x36, 0x30, 0xcf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x27, 0x4d, 0xc5, 0xe6, 0xfd, 0x00,
	0x00, 0x00,
}
