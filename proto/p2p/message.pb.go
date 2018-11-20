// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/p2p/message.proto

package airbloc_p2p_v1

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

type CID int32

const (
	CID_AIRBLOC CID = 0
)

var CID_name = map[int32]string{
	0: "AIRBLOC",
}
var CID_value = map[string]int32{
	"AIRBLOC": 0,
}

func (x CID) String() string {
	return proto.EnumName(CID_name, int32(x))
}
func (CID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_message_7ea4b82cd1fe5a40, []int{0}
}

type Topic int32

const (
	Topic_UNKNOWN                  Topic = 0
	Topic_RENCRYPTION_KEY_REQUEST  Topic = 1
	Topic_RENCRYPTION_KEY_RESPONSE Topic = 2
	Topic_IDENTITY_REQUEST         Topic = 3
	Topic_IDENTITY_RESPONSE        Topic = 4
	Topic_DATASYNC_REQUEST         Topic = 5
	Topic_DATASYNC_RESPONSE        Topic = 6
	Topic_DAC_REQUEST              Topic = 7
	Topic_DAC_RESPONSE             Topic = 8
	Topic_TEST_PING                Topic = 9
	Topic_TEST_PONG                Topic = 10
)

var Topic_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "RENCRYPTION_KEY_REQUEST",
	2:  "RENCRYPTION_KEY_RESPONSE",
	3:  "IDENTITY_REQUEST",
	4:  "IDENTITY_RESPONSE",
	5:  "DATASYNC_REQUEST",
	6:  "DATASYNC_RESPONSE",
	7:  "DAC_REQUEST",
	8:  "DAC_RESPONSE",
	9:  "TEST_PING",
	10: "TEST_PONG",
}
var Topic_value = map[string]int32{
	"UNKNOWN":                  0,
	"RENCRYPTION_KEY_REQUEST":  1,
	"RENCRYPTION_KEY_RESPONSE": 2,
	"IDENTITY_REQUEST":         3,
	"IDENTITY_RESPONSE":        4,
	"DATASYNC_REQUEST":         5,
	"DATASYNC_RESPONSE":        6,
	"DAC_REQUEST":              7,
	"DAC_RESPONSE":             8,
	"TEST_PING":                9,
	"TEST_PONG":                10,
}

func (x Topic) String() string {
	return proto.EnumName(Topic_name, int32(x))
}
func (Topic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_message_7ea4b82cd1fe5a40, []int{1}
}

type Message struct {
	Topic                Topic    `protobuf:"varint,1,opt,name=topic,proto3,enum=airbloc.p2p.v1.Topic" json:"topic,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	From                 []byte   `protobuf:"bytes,3,opt,name=From,proto3" json:"From,omitempty"`
	Protocol             []byte   `protobuf:"bytes,4,opt,name=Protocol,proto3" json:"Protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_7ea4b82cd1fe5a40, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetTopic() Topic {
	if m != nil {
		return m.Topic
	}
	return Topic_UNKNOWN
}

func (m *Message) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Message) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Message) GetProtocol() []byte {
	if m != nil {
		return m.Protocol
	}
	return nil
}

type ReEncryptionKeyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReEncryptionKeyRequest) Reset()         { *m = ReEncryptionKeyRequest{} }
func (m *ReEncryptionKeyRequest) String() string { return proto.CompactTextString(m) }
func (*ReEncryptionKeyRequest) ProtoMessage()    {}
func (*ReEncryptionKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_7ea4b82cd1fe5a40, []int{1}
}
func (m *ReEncryptionKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReEncryptionKeyRequest.Unmarshal(m, b)
}
func (m *ReEncryptionKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReEncryptionKeyRequest.Marshal(b, m, deterministic)
}
func (dst *ReEncryptionKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReEncryptionKeyRequest.Merge(dst, src)
}
func (m *ReEncryptionKeyRequest) XXX_Size() int {
	return xxx_messageInfo_ReEncryptionKeyRequest.Size(m)
}
func (m *ReEncryptionKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReEncryptionKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReEncryptionKeyRequest proto.InternalMessageInfo

type ReEncryptionKeyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReEncryptionKeyResponse) Reset()         { *m = ReEncryptionKeyResponse{} }
func (m *ReEncryptionKeyResponse) String() string { return proto.CompactTextString(m) }
func (*ReEncryptionKeyResponse) ProtoMessage()    {}
func (*ReEncryptionKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_7ea4b82cd1fe5a40, []int{2}
}
func (m *ReEncryptionKeyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReEncryptionKeyResponse.Unmarshal(m, b)
}
func (m *ReEncryptionKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReEncryptionKeyResponse.Marshal(b, m, deterministic)
}
func (dst *ReEncryptionKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReEncryptionKeyResponse.Merge(dst, src)
}
func (m *ReEncryptionKeyResponse) XXX_Size() int {
	return xxx_messageInfo_ReEncryptionKeyResponse.Size(m)
}
func (m *ReEncryptionKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReEncryptionKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReEncryptionKeyResponse proto.InternalMessageInfo

type IdentityRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdentityRequest) Reset()         { *m = IdentityRequest{} }
func (m *IdentityRequest) String() string { return proto.CompactTextString(m) }
func (*IdentityRequest) ProtoMessage()    {}
func (*IdentityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_7ea4b82cd1fe5a40, []int{3}
}
func (m *IdentityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentityRequest.Unmarshal(m, b)
}
func (m *IdentityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentityRequest.Marshal(b, m, deterministic)
}
func (dst *IdentityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityRequest.Merge(dst, src)
}
func (m *IdentityRequest) XXX_Size() int {
	return xxx_messageInfo_IdentityRequest.Size(m)
}
func (m *IdentityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityRequest proto.InternalMessageInfo

type IdentityResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdentityResponse) Reset()         { *m = IdentityResponse{} }
func (m *IdentityResponse) String() string { return proto.CompactTextString(m) }
func (*IdentityResponse) ProtoMessage()    {}
func (*IdentityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_7ea4b82cd1fe5a40, []int{4}
}
func (m *IdentityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentityResponse.Unmarshal(m, b)
}
func (m *IdentityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentityResponse.Marshal(b, m, deterministic)
}
func (dst *IdentityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityResponse.Merge(dst, src)
}
func (m *IdentityResponse) XXX_Size() int {
	return xxx_messageInfo_IdentityResponse.Size(m)
}
func (m *IdentityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityResponse proto.InternalMessageInfo

type DatasyncRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatasyncRequest) Reset()         { *m = DatasyncRequest{} }
func (m *DatasyncRequest) String() string { return proto.CompactTextString(m) }
func (*DatasyncRequest) ProtoMessage()    {}
func (*DatasyncRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_7ea4b82cd1fe5a40, []int{5}
}
func (m *DatasyncRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatasyncRequest.Unmarshal(m, b)
}
func (m *DatasyncRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatasyncRequest.Marshal(b, m, deterministic)
}
func (dst *DatasyncRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatasyncRequest.Merge(dst, src)
}
func (m *DatasyncRequest) XXX_Size() int {
	return xxx_messageInfo_DatasyncRequest.Size(m)
}
func (m *DatasyncRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DatasyncRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DatasyncRequest proto.InternalMessageInfo

type DatasyncResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatasyncResponse) Reset()         { *m = DatasyncResponse{} }
func (m *DatasyncResponse) String() string { return proto.CompactTextString(m) }
func (*DatasyncResponse) ProtoMessage()    {}
func (*DatasyncResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_7ea4b82cd1fe5a40, []int{6}
}
func (m *DatasyncResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatasyncResponse.Unmarshal(m, b)
}
func (m *DatasyncResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatasyncResponse.Marshal(b, m, deterministic)
}
func (dst *DatasyncResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatasyncResponse.Merge(dst, src)
}
func (m *DatasyncResponse) XXX_Size() int {
	return xxx_messageInfo_DatasyncResponse.Size(m)
}
func (m *DatasyncResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DatasyncResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DatasyncResponse proto.InternalMessageInfo

type DACRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DACRequest) Reset()         { *m = DACRequest{} }
func (m *DACRequest) String() string { return proto.CompactTextString(m) }
func (*DACRequest) ProtoMessage()    {}
func (*DACRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_7ea4b82cd1fe5a40, []int{7}
}
func (m *DACRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DACRequest.Unmarshal(m, b)
}
func (m *DACRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DACRequest.Marshal(b, m, deterministic)
}
func (dst *DACRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DACRequest.Merge(dst, src)
}
func (m *DACRequest) XXX_Size() int {
	return xxx_messageInfo_DACRequest.Size(m)
}
func (m *DACRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DACRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DACRequest proto.InternalMessageInfo

type DACResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DACResponse) Reset()         { *m = DACResponse{} }
func (m *DACResponse) String() string { return proto.CompactTextString(m) }
func (*DACResponse) ProtoMessage()    {}
func (*DACResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_7ea4b82cd1fe5a40, []int{8}
}
func (m *DACResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DACResponse.Unmarshal(m, b)
}
func (m *DACResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DACResponse.Marshal(b, m, deterministic)
}
func (dst *DACResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DACResponse.Merge(dst, src)
}
func (m *DACResponse) XXX_Size() int {
	return xxx_messageInfo_DACResponse.Size(m)
}
func (m *DACResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DACResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DACResponse proto.InternalMessageInfo

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
	return fileDescriptor_message_7ea4b82cd1fe5a40, []int{9}
}
func (m *TestPing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestPing.Unmarshal(m, b)
}
func (m *TestPing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestPing.Marshal(b, m, deterministic)
}
func (dst *TestPing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestPing.Merge(dst, src)
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
	return fileDescriptor_message_7ea4b82cd1fe5a40, []int{10}
}
func (m *TestPong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestPong.Unmarshal(m, b)
}
func (m *TestPong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestPong.Marshal(b, m, deterministic)
}
func (dst *TestPong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestPong.Merge(dst, src)
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
	proto.RegisterType((*ReEncryptionKeyRequest)(nil), "airbloc.p2p.v1.ReEncryptionKeyRequest")
	proto.RegisterType((*ReEncryptionKeyResponse)(nil), "airbloc.p2p.v1.ReEncryptionKeyResponse")
	proto.RegisterType((*IdentityRequest)(nil), "airbloc.p2p.v1.IdentityRequest")
	proto.RegisterType((*IdentityResponse)(nil), "airbloc.p2p.v1.IdentityResponse")
	proto.RegisterType((*DatasyncRequest)(nil), "airbloc.p2p.v1.DatasyncRequest")
	proto.RegisterType((*DatasyncResponse)(nil), "airbloc.p2p.v1.DatasyncResponse")
	proto.RegisterType((*DACRequest)(nil), "airbloc.p2p.v1.DACRequest")
	proto.RegisterType((*DACResponse)(nil), "airbloc.p2p.v1.DACResponse")
	proto.RegisterType((*TestPing)(nil), "airbloc.p2p.v1.TestPing")
	proto.RegisterType((*TestPong)(nil), "airbloc.p2p.v1.TestPong")
	proto.RegisterEnum("airbloc.p2p.v1.CID", CID_name, CID_value)
	proto.RegisterEnum("airbloc.p2p.v1.Topic", Topic_name, Topic_value)
}

func init() { proto.RegisterFile("proto/p2p/message.proto", fileDescriptor_message_7ea4b82cd1fe5a40) }

var fileDescriptor_message_7ea4b82cd1fe5a40 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdf, 0x6a, 0xdb, 0x30,
	0x14, 0xc6, 0xeb, 0x26, 0x69, 0x92, 0xd3, 0xb4, 0x55, 0x0f, 0xeb, 0xe2, 0xfd, 0xb9, 0x28, 0x61,
	0x17, 0xa5, 0x03, 0x97, 0x65, 0x4f, 0x90, 0xd9, 0x5a, 0x31, 0xd9, 0x64, 0x4f, 0x56, 0x19, 0xb9,
	0x0a, 0xa9, 0x27, 0x8a, 0xa1, 0xb5, 0xb4, 0x58, 0x1b, 0x64, 0x8f, 0xbc, 0xa7, 0x18, 0x52, 0x5c,
	0x3b, 0xa3, 0xd0, 0xbb, 0xf3, 0xfd, 0xbe, 0x9f, 0xb1, 0xf9, 0x30, 0x8c, 0xf5, 0x5a, 0x19, 0x75,
	0xa5, 0xa7, 0xfa, 0xea, 0x41, 0x56, 0xd5, 0xea, 0x4e, 0x06, 0x8e, 0xe0, 0xf1, 0xaa, 0x58, 0xdf,
	0xde, 0xab, 0x3c, 0xd0, 0x53, 0x1d, 0xfc, 0xfe, 0x30, 0xf9, 0x03, 0xfd, 0xaf, 0x5b, 0x01, 0xdf,
	0x43, 0xcf, 0x28, 0x5d, 0xe4, 0xbe, 0x77, 0xee, 0x5d, 0x1c, 0x4f, 0xcf, 0x82, 0xff, 0xd5, 0x40,
	0xd8, 0x92, 0x6f, 0x1d, 0x44, 0xe8, 0x46, 0x2b, 0xb3, 0xf2, 0xf7, 0xcf, 0xbd, 0x8b, 0x11, 0x77,
	0xb7, 0x65, 0x9f, 0xd7, 0xea, 0xc1, 0xef, 0x6c, 0x99, 0xbd, 0xf1, 0x35, 0x0c, 0x52, 0xfb, 0xe2,
	0x5c, 0xdd, 0xfb, 0x5d, 0xc7, 0x9b, 0x3c, 0xf1, 0xe1, 0x25, 0x97, 0xb4, 0xcc, 0xd7, 0x1b, 0x6d,
	0x0a, 0x55, 0xce, 0xe5, 0x86, 0xcb, 0x9f, 0xbf, 0x64, 0x65, 0x26, 0xaf, 0x60, 0xfc, 0xa4, 0xa9,
	0xb4, 0x2a, 0x2b, 0x39, 0x39, 0x85, 0x93, 0xf8, 0x87, 0x2c, 0x4d, 0x61, 0x1a, 0x1b, 0x81, 0xb4,
	0xa8, 0xd5, 0xec, 0x37, 0x55, 0x9b, 0x32, 0xdf, 0xd1, 0x5a, 0x54, 0x6b, 0x23, 0x80, 0x68, 0x16,
	0x3e, 0x1a, 0x47, 0x70, 0xe8, 0x52, 0x5d, 0xbe, 0x83, 0x81, 0x90, 0x95, 0x49, 0x8b, 0xf2, 0x0e,
	0x7d, 0xe8, 0xd7, 0x43, 0xba, 0x79, 0x86, 0xfc, 0x31, 0x36, 0x96, 0x7a, 0xce, 0xba, 0x44, 0xe8,
	0x84, 0x71, 0x84, 0x87, 0xd0, 0x9f, 0xc5, 0xfc, 0xd3, 0x97, 0x24, 0x24, 0x7b, 0x97, 0x7f, 0x3d,
	0xe8, 0xb9, 0x51, 0x2d, 0xbe, 0x61, 0x73, 0x96, 0x7c, 0x67, 0x64, 0x0f, 0xdf, 0xc0, 0x98, 0x53,
	0x16, 0xf2, 0x45, 0x2a, 0xe2, 0x84, 0x2d, 0xe7, 0x74, 0xb1, 0xe4, 0xf4, 0xdb, 0x0d, 0xcd, 0x04,
	0xf1, 0xf0, 0x2d, 0xf8, 0x4f, 0xcb, 0x2c, 0x4d, 0x58, 0x46, 0xc9, 0x3e, 0xbe, 0x00, 0x12, 0x47,
	0x94, 0x89, 0x58, 0xb4, 0xcf, 0x74, 0xf0, 0x0c, 0x4e, 0x77, 0x68, 0x2d, 0x77, 0xad, 0x1c, 0xcd,
	0xc4, 0x2c, 0x5b, 0xb0, 0xb0, 0x91, 0x7b, 0x56, 0xde, 0xa1, 0xb5, 0x7c, 0x80, 0x27, 0x6e, 0x9a,
	0xc6, 0xeb, 0x23, 0x81, 0xd1, 0x16, 0xd4, 0xca, 0x00, 0x8f, 0x60, 0x28, 0x68, 0x26, 0x96, 0x69,
	0xcc, 0xae, 0xc9, 0xb0, 0x8d, 0x09, 0xbb, 0x26, 0x70, 0x7b, 0xe0, 0xfe, 0xbf, 0x8f, 0xff, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x36, 0xb1, 0x8f, 0x91, 0x9a, 0x02, 0x00, 0x00,
}
