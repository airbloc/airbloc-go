// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/p2p/v1/dauth.proto

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

type DAuthSignUpRequest struct {
	Provider             string   `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	IdentityHash         string   `protobuf:"bytes,2,opt,name=identityHash,proto3" json:"identityHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DAuthSignUpRequest) Reset()         { *m = DAuthSignUpRequest{} }
func (m *DAuthSignUpRequest) String() string { return proto.CompactTextString(m) }
func (*DAuthSignUpRequest) ProtoMessage()    {}
func (*DAuthSignUpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6b4b5d564e7416f, []int{0}
}

func (m *DAuthSignUpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DAuthSignUpRequest.Unmarshal(m, b)
}
func (m *DAuthSignUpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DAuthSignUpRequest.Marshal(b, m, deterministic)
}
func (m *DAuthSignUpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DAuthSignUpRequest.Merge(m, src)
}
func (m *DAuthSignUpRequest) XXX_Size() int {
	return xxx_messageInfo_DAuthSignUpRequest.Size(m)
}
func (m *DAuthSignUpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DAuthSignUpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DAuthSignUpRequest proto.InternalMessageInfo

func (m *DAuthSignUpRequest) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *DAuthSignUpRequest) GetIdentityHash() string {
	if m != nil {
		return m.IdentityHash
	}
	return ""
}

type DAuthSignUpResponse struct {
	AccountId            string   `protobuf:"bytes,1,opt,name=accountId,proto3" json:"accountId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DAuthSignUpResponse) Reset()         { *m = DAuthSignUpResponse{} }
func (m *DAuthSignUpResponse) String() string { return proto.CompactTextString(m) }
func (*DAuthSignUpResponse) ProtoMessage()    {}
func (*DAuthSignUpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6b4b5d564e7416f, []int{1}
}

func (m *DAuthSignUpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DAuthSignUpResponse.Unmarshal(m, b)
}
func (m *DAuthSignUpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DAuthSignUpResponse.Marshal(b, m, deterministic)
}
func (m *DAuthSignUpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DAuthSignUpResponse.Merge(m, src)
}
func (m *DAuthSignUpResponse) XXX_Size() int {
	return xxx_messageInfo_DAuthSignUpResponse.Size(m)
}
func (m *DAuthSignUpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DAuthSignUpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DAuthSignUpResponse proto.InternalMessageInfo

func (m *DAuthSignUpResponse) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

//*
// DAuth (Data Collection Authentification) API.
type DAuthRequest struct {
	CollectionId         string   `protobuf:"bytes,1,opt,name=collectionId,proto3" json:"collectionId,omitempty"`
	Allow                bool     `protobuf:"varint,2,opt,name=allow,proto3" json:"allow,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DAuthRequest) Reset()         { *m = DAuthRequest{} }
func (m *DAuthRequest) String() string { return proto.CompactTextString(m) }
func (*DAuthRequest) ProtoMessage()    {}
func (*DAuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6b4b5d564e7416f, []int{2}
}

func (m *DAuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DAuthRequest.Unmarshal(m, b)
}
func (m *DAuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DAuthRequest.Marshal(b, m, deterministic)
}
func (m *DAuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DAuthRequest.Merge(m, src)
}
func (m *DAuthRequest) XXX_Size() int {
	return xxx_messageInfo_DAuthRequest.Size(m)
}
func (m *DAuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DAuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DAuthRequest proto.InternalMessageInfo

func (m *DAuthRequest) GetCollectionId() string {
	if m != nil {
		return m.CollectionId
	}
	return ""
}

func (m *DAuthRequest) GetAllow() bool {
	if m != nil {
		return m.Allow
	}
	return false
}

type DAuthResponse struct {
	CollectionId         string   `protobuf:"bytes,1,opt,name=collectionId,proto3" json:"collectionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DAuthResponse) Reset()         { *m = DAuthResponse{} }
func (m *DAuthResponse) String() string { return proto.CompactTextString(m) }
func (*DAuthResponse) ProtoMessage()    {}
func (*DAuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6b4b5d564e7416f, []int{3}
}

func (m *DAuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DAuthResponse.Unmarshal(m, b)
}
func (m *DAuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DAuthResponse.Marshal(b, m, deterministic)
}
func (m *DAuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DAuthResponse.Merge(m, src)
}
func (m *DAuthResponse) XXX_Size() int {
	return xxx_messageInfo_DAuthResponse.Size(m)
}
func (m *DAuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DAuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DAuthResponse proto.InternalMessageInfo

func (m *DAuthResponse) GetCollectionId() string {
	if m != nil {
		return m.CollectionId
	}
	return ""
}

func init() {
	proto.RegisterType((*DAuthSignUpRequest)(nil), "airbloc.p2p.v1.DAuthSignUpRequest")
	proto.RegisterType((*DAuthSignUpResponse)(nil), "airbloc.p2p.v1.DAuthSignUpResponse")
	proto.RegisterType((*DAuthRequest)(nil), "airbloc.p2p.v1.DAuthRequest")
	proto.RegisterType((*DAuthResponse)(nil), "airbloc.p2p.v1.DAuthResponse")
}

func init() { proto.RegisterFile("proto/p2p/v1/dauth.proto", fileDescriptor_d6b4b5d564e7416f) }

var fileDescriptor_d6b4b5d564e7416f = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x59, 0x41, 0xd9, 0x1d, 0xaa, 0x87, 0xe8, 0xa1, 0x88, 0x07, 0xc9, 0x49, 0x44, 0x1b,
	0x76, 0xfb, 0x04, 0x8a, 0x87, 0xf5, 0x5a, 0xf5, 0xe2, 0x2d, 0x4d, 0x43, 0x1b, 0x88, 0x99, 0xb1,
	0x9d, 0x54, 0x7c, 0x7b, 0x21, 0x76, 0x57, 0x7b, 0xdb, 0x53, 0xf8, 0xbf, 0x30, 0xdf, 0x0c, 0x3f,
	0xe4, 0xd4, 0x23, 0xa3, 0xa2, 0x0d, 0xa9, 0x71, 0xad, 0x1a, 0x1d, 0xb9, 0x2b, 0x12, 0x12, 0x67,
	0xda, 0xf5, 0xb5, 0x47, 0x53, 0xd0, 0x86, 0x8a, 0x71, 0x2d, 0x5f, 0x41, 0x3c, 0x3d, 0x44, 0xee,
	0x5e, 0x5c, 0x1b, 0xde, 0xa8, 0xb2, 0x9f, 0xd1, 0x0e, 0x2c, 0x2e, 0x61, 0x49, 0x3d, 0x8e, 0xae,
	0xb1, 0x7d, 0xbe, 0xb8, 0x5e, 0xdc, 0xac, 0xaa, 0x7d, 0x16, 0x12, 0x32, 0xd7, 0xd8, 0xc0, 0x8e,
	0xbf, 0xb7, 0x7a, 0xe8, 0xf2, 0xa3, 0xf4, 0x3f, 0x63, 0xb2, 0x84, 0xf3, 0x99, 0x75, 0x20, 0x0c,
	0x83, 0x15, 0x57, 0xb0, 0xd2, 0xc6, 0x60, 0x0c, 0xfc, 0xdc, 0x4c, 0xde, 0x3f, 0x20, 0xb7, 0x90,
	0xa5, 0xa1, 0xdd, 0x11, 0x12, 0x32, 0x83, 0xde, 0x5b, 0xc3, 0x0e, 0xc3, 0x7e, 0x60, 0xc6, 0xc4,
	0x05, 0x1c, 0x6b, 0xef, 0xf1, 0x2b, 0x5d, 0xb1, 0xac, 0x7e, 0x83, 0x2c, 0xe1, 0x74, 0x32, 0x4d,
	0x8b, 0x0f, 0x50, 0x3d, 0xde, 0xbd, 0xdf, 0xb6, 0x8e, 0xbb, 0x58, 0x17, 0x06, 0x3f, 0xd4, 0x54,
	0xd3, 0xee, 0xbd, 0x6f, 0x51, 0xfd, 0xef, 0xb4, 0x3e, 0x49, 0xa9, 0xfc, 0x09, 0x00, 0x00, 0xff,
	0xff, 0xdb, 0xd4, 0x61, 0x09, 0x6a, 0x01, 0x00, 0x00,
}
