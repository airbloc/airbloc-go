// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/rpc/v1/server/user.proto

package server

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

// TODO: pagination
type DataRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	From                 int64    `protobuf:"varint,2,opt,name=from,proto3" json:"from,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataRequest) Reset()         { *m = DataRequest{} }
func (m *DataRequest) String() string { return proto.CompactTextString(m) }
func (*DataRequest) ProtoMessage()    {}
func (*DataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_93739d9e0922a9b9, []int{0}
}

func (m *DataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataRequest.Unmarshal(m, b)
}
func (m *DataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataRequest.Marshal(b, m, deterministic)
}
func (m *DataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataRequest.Merge(m, src)
}
func (m *DataRequest) XXX_Size() int {
	return xxx_messageInfo_DataRequest.Size(m)
}
func (m *DataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataRequest proto.InternalMessageInfo

func (m *DataRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *DataRequest) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

// TODO: pagination
type DataIdRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataIdRequest) Reset()         { *m = DataIdRequest{} }
func (m *DataIdRequest) String() string { return proto.CompactTextString(m) }
func (*DataIdRequest) ProtoMessage()    {}
func (*DataIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_93739d9e0922a9b9, []int{1}
}

func (m *DataIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataIdRequest.Unmarshal(m, b)
}
func (m *DataIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataIdRequest.Marshal(b, m, deterministic)
}
func (m *DataIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataIdRequest.Merge(m, src)
}
func (m *DataIdRequest) XXX_Size() int {
	return xxx_messageInfo_DataIdRequest.Size(m)
}
func (m *DataIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataIdRequest proto.InternalMessageInfo

func (m *DataIdRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetDataReponse struct {
	Collections          []*GetDataReponse_Collection `protobuf:"bytes,1,rep,name=collections,proto3" json:"collections,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *GetDataReponse) Reset()         { *m = GetDataReponse{} }
func (m *GetDataReponse) String() string { return proto.CompactTextString(m) }
func (*GetDataReponse) ProtoMessage()    {}
func (*GetDataReponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_93739d9e0922a9b9, []int{2}
}

func (m *GetDataReponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDataReponse.Unmarshal(m, b)
}
func (m *GetDataReponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDataReponse.Marshal(b, m, deterministic)
}
func (m *GetDataReponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDataReponse.Merge(m, src)
}
func (m *GetDataReponse) XXX_Size() int {
	return xxx_messageInfo_GetDataReponse.Size(m)
}
func (m *GetDataReponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDataReponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDataReponse proto.InternalMessageInfo

func (m *GetDataReponse) GetCollections() []*GetDataReponse_Collection {
	if m != nil {
		return m.Collections
	}
	return nil
}

type GetDataReponse_Data struct {
	CollectedAt          int64    `protobuf:"varint,1,opt,name=collectedAt,proto3" json:"collectedAt,omitempty"`
	IngestedAt           int64    `protobuf:"varint,2,opt,name=ingestedAt,proto3" json:"ingestedAt,omitempty"`
	Payload              string   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDataReponse_Data) Reset()         { *m = GetDataReponse_Data{} }
func (m *GetDataReponse_Data) String() string { return proto.CompactTextString(m) }
func (*GetDataReponse_Data) ProtoMessage()    {}
func (*GetDataReponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_93739d9e0922a9b9, []int{2, 0}
}

func (m *GetDataReponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDataReponse_Data.Unmarshal(m, b)
}
func (m *GetDataReponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDataReponse_Data.Marshal(b, m, deterministic)
}
func (m *GetDataReponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDataReponse_Data.Merge(m, src)
}
func (m *GetDataReponse_Data) XXX_Size() int {
	return xxx_messageInfo_GetDataReponse_Data.Size(m)
}
func (m *GetDataReponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDataReponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_GetDataReponse_Data proto.InternalMessageInfo

func (m *GetDataReponse_Data) GetCollectedAt() int64 {
	if m != nil {
		return m.CollectedAt
	}
	return 0
}

func (m *GetDataReponse_Data) GetIngestedAt() int64 {
	if m != nil {
		return m.IngestedAt
	}
	return 0
}

func (m *GetDataReponse_Data) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type GetDataReponse_Collection struct {
	Id                   string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Data                 []*GetDataReponse_Data `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GetDataReponse_Collection) Reset()         { *m = GetDataReponse_Collection{} }
func (m *GetDataReponse_Collection) String() string { return proto.CompactTextString(m) }
func (*GetDataReponse_Collection) ProtoMessage()    {}
func (*GetDataReponse_Collection) Descriptor() ([]byte, []int) {
	return fileDescriptor_93739d9e0922a9b9, []int{2, 1}
}

func (m *GetDataReponse_Collection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDataReponse_Collection.Unmarshal(m, b)
}
func (m *GetDataReponse_Collection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDataReponse_Collection.Marshal(b, m, deterministic)
}
func (m *GetDataReponse_Collection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDataReponse_Collection.Merge(m, src)
}
func (m *GetDataReponse_Collection) XXX_Size() int {
	return xxx_messageInfo_GetDataReponse_Collection.Size(m)
}
func (m *GetDataReponse_Collection) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDataReponse_Collection.DiscardUnknown(m)
}

var xxx_messageInfo_GetDataReponse_Collection proto.InternalMessageInfo

func (m *GetDataReponse_Collection) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetDataReponse_Collection) GetData() []*GetDataReponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetDataIdsResponse struct {
	Collections          []*GetDataIdsResponse_Collection `protobuf:"bytes,1,rep,name=collections,proto3" json:"collections,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *GetDataIdsResponse) Reset()         { *m = GetDataIdsResponse{} }
func (m *GetDataIdsResponse) String() string { return proto.CompactTextString(m) }
func (*GetDataIdsResponse) ProtoMessage()    {}
func (*GetDataIdsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_93739d9e0922a9b9, []int{3}
}

func (m *GetDataIdsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDataIdsResponse.Unmarshal(m, b)
}
func (m *GetDataIdsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDataIdsResponse.Marshal(b, m, deterministic)
}
func (m *GetDataIdsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDataIdsResponse.Merge(m, src)
}
func (m *GetDataIdsResponse) XXX_Size() int {
	return xxx_messageInfo_GetDataIdsResponse.Size(m)
}
func (m *GetDataIdsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDataIdsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDataIdsResponse proto.InternalMessageInfo

func (m *GetDataIdsResponse) GetCollections() []*GetDataIdsResponse_Collection {
	if m != nil {
		return m.Collections
	}
	return nil
}

type GetDataIdsResponse_DataInfo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CollectedAt          int64    `protobuf:"varint,2,opt,name=collectedAt,proto3" json:"collectedAt,omitempty"`
	IngestedAt           int64    `protobuf:"varint,3,opt,name=ingestedAt,proto3" json:"ingestedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDataIdsResponse_DataInfo) Reset()         { *m = GetDataIdsResponse_DataInfo{} }
func (m *GetDataIdsResponse_DataInfo) String() string { return proto.CompactTextString(m) }
func (*GetDataIdsResponse_DataInfo) ProtoMessage()    {}
func (*GetDataIdsResponse_DataInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_93739d9e0922a9b9, []int{3, 0}
}

func (m *GetDataIdsResponse_DataInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDataIdsResponse_DataInfo.Unmarshal(m, b)
}
func (m *GetDataIdsResponse_DataInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDataIdsResponse_DataInfo.Marshal(b, m, deterministic)
}
func (m *GetDataIdsResponse_DataInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDataIdsResponse_DataInfo.Merge(m, src)
}
func (m *GetDataIdsResponse_DataInfo) XXX_Size() int {
	return xxx_messageInfo_GetDataIdsResponse_DataInfo.Size(m)
}
func (m *GetDataIdsResponse_DataInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDataIdsResponse_DataInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GetDataIdsResponse_DataInfo proto.InternalMessageInfo

func (m *GetDataIdsResponse_DataInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetDataIdsResponse_DataInfo) GetCollectedAt() int64 {
	if m != nil {
		return m.CollectedAt
	}
	return 0
}

func (m *GetDataIdsResponse_DataInfo) GetIngestedAt() int64 {
	if m != nil {
		return m.IngestedAt
	}
	return 0
}

type GetDataIdsResponse_Collection struct {
	CollectionId         string                         `protobuf:"bytes,3,opt,name=collectionId,proto3" json:"collectionId,omitempty"`
	DataInfoes           []*GetDataIdsResponse_DataInfo `protobuf:"bytes,4,rep,name=dataInfoes,proto3" json:"dataInfoes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *GetDataIdsResponse_Collection) Reset()         { *m = GetDataIdsResponse_Collection{} }
func (m *GetDataIdsResponse_Collection) String() string { return proto.CompactTextString(m) }
func (*GetDataIdsResponse_Collection) ProtoMessage()    {}
func (*GetDataIdsResponse_Collection) Descriptor() ([]byte, []int) {
	return fileDescriptor_93739d9e0922a9b9, []int{3, 1}
}

func (m *GetDataIdsResponse_Collection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDataIdsResponse_Collection.Unmarshal(m, b)
}
func (m *GetDataIdsResponse_Collection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDataIdsResponse_Collection.Marshal(b, m, deterministic)
}
func (m *GetDataIdsResponse_Collection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDataIdsResponse_Collection.Merge(m, src)
}
func (m *GetDataIdsResponse_Collection) XXX_Size() int {
	return xxx_messageInfo_GetDataIdsResponse_Collection.Size(m)
}
func (m *GetDataIdsResponse_Collection) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDataIdsResponse_Collection.DiscardUnknown(m)
}

var xxx_messageInfo_GetDataIdsResponse_Collection proto.InternalMessageInfo

func (m *GetDataIdsResponse_Collection) GetCollectionId() string {
	if m != nil {
		return m.CollectionId
	}
	return ""
}

func (m *GetDataIdsResponse_Collection) GetDataInfoes() []*GetDataIdsResponse_DataInfo {
	if m != nil {
		return m.DataInfoes
	}
	return nil
}

func init() {
	proto.RegisterType((*DataRequest)(nil), "airbloc.rpc.v1.DataRequest")
	proto.RegisterType((*DataIdRequest)(nil), "airbloc.rpc.v1.DataIdRequest")
	proto.RegisterType((*GetDataReponse)(nil), "airbloc.rpc.v1.GetDataReponse")
	proto.RegisterType((*GetDataReponse_Data)(nil), "airbloc.rpc.v1.GetDataReponse.Data")
	proto.RegisterType((*GetDataReponse_Collection)(nil), "airbloc.rpc.v1.GetDataReponse.Collection")
	proto.RegisterType((*GetDataIdsResponse)(nil), "airbloc.rpc.v1.GetDataIdsResponse")
	proto.RegisterType((*GetDataIdsResponse_DataInfo)(nil), "airbloc.rpc.v1.GetDataIdsResponse.DataInfo")
	proto.RegisterType((*GetDataIdsResponse_Collection)(nil), "airbloc.rpc.v1.GetDataIdsResponse.Collection")
}

func init() { proto.RegisterFile("proto/rpc/v1/server/user.proto", fileDescriptor_93739d9e0922a9b9) }

var fileDescriptor_93739d9e0922a9b9 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x8f, 0xda, 0x30,
	0x10, 0x55, 0x3e, 0x04, 0xed, 0xd0, 0x72, 0x98, 0x43, 0x15, 0xa5, 0x2a, 0x8a, 0xd2, 0x43, 0xa9,
	0x2a, 0x12, 0x01, 0x87, 0xaa, 0xc7, 0x7e, 0xa8, 0x55, 0xc4, 0x01, 0x29, 0x12, 0x97, 0xaa, 0x97,
	0x24, 0x36, 0x34, 0x52, 0x88, 0x53, 0xdb, 0x20, 0xf5, 0xb0, 0xff, 0x63, 0x6f, 0xfb, 0x5f, 0xf6,
	0x97, 0xad, 0x70, 0x12, 0x48, 0x58, 0xd8, 0xec, 0x29, 0xf6, 0xf8, 0xbd, 0x99, 0xf7, 0x5e, 0x6c,
	0x18, 0x15, 0x9c, 0x49, 0xe6, 0xf3, 0x22, 0xf1, 0xf7, 0x53, 0x5f, 0x50, 0xbe, 0xa7, 0xdc, 0xdf,
	0x09, 0xca, 0x3d, 0x75, 0x80, 0xc3, 0x28, 0xe5, 0x71, 0xc6, 0x12, 0x8f, 0x17, 0x89, 0xb7, 0x9f,
	0xba, 0x5f, 0x60, 0xf0, 0x23, 0x92, 0x51, 0x48, 0xff, 0xed, 0xa8, 0x90, 0xf8, 0x06, 0x7a, 0x07,
	0x70, 0x40, 0x2c, 0xcd, 0xd1, 0xc6, 0x2f, 0xc3, 0x6a, 0x87, 0x08, 0xe6, 0x9a, 0xb3, 0xad, 0xa5,
	0x3b, 0xda, 0xd8, 0x08, 0xd5, 0xda, 0xfd, 0x00, 0xaf, 0x0f, 0xd4, 0x80, 0x74, 0x90, 0xdd, 0x5b,
	0x1d, 0x86, 0xbf, 0xa8, 0x2c, 0xe7, 0x14, 0x2c, 0x17, 0x14, 0x17, 0x30, 0x48, 0x58, 0x96, 0xd1,
	0x44, 0xa6, 0x2c, 0x17, 0x96, 0xe6, 0x18, 0xe3, 0xc1, 0xec, 0xa3, 0xd7, 0x16, 0xe7, 0xb5, 0x49,
	0xde, 0xf7, 0x23, 0x23, 0x6c, 0xb2, 0xed, 0x18, 0xcc, 0x03, 0x0c, 0x9d, 0x63, 0x53, 0x4a, 0xbe,
	0x4a, 0x25, 0xc2, 0x08, 0x9b, 0x25, 0x1c, 0x01, 0xa4, 0xf9, 0x86, 0x8a, 0x12, 0x50, 0x9a, 0x69,
	0x54, 0xd0, 0x82, 0x7e, 0x11, 0xfd, 0xcf, 0x58, 0x44, 0x2c, 0x43, 0x59, 0xa8, 0xb7, 0xf6, 0x0a,
	0xe0, 0x34, 0x1e, 0x87, 0xa0, 0xa7, 0xb5, 0x4b, 0x3d, 0x25, 0xf8, 0x19, 0x4c, 0x12, 0xc9, 0xc8,
	0x32, 0x94, 0x8f, 0xf7, 0x1d, 0x3e, 0xd4, 0x5a, 0x11, 0xdc, 0x7b, 0x1d, 0xb0, 0x3a, 0x0d, 0x88,
	0x08, 0xa9, 0x28, 0xe3, 0x59, 0x5e, 0x8a, 0x67, 0x72, 0xa5, 0x6d, 0x83, 0x78, 0x35, 0xa2, 0x3f,
	0xf0, 0x42, 0x41, 0xf3, 0x35, 0x7b, 0x24, 0xfe, 0x2c, 0x36, 0xbd, 0x2b, 0x36, 0xe3, 0x3c, 0x36,
	0xfb, 0xa6, 0x15, 0x8e, 0x0b, 0xaf, 0x4e, 0xa3, 0x83, 0x3a, 0xc9, 0x56, 0x0d, 0x17, 0x00, 0xa4,
	0xd2, 0x43, 0x85, 0x65, 0x2a, 0x7f, 0x9f, 0x9e, 0xe1, 0xaf, 0x36, 0x11, 0x36, 0xe8, 0xb3, 0x3b,
	0x0d, 0xcc, 0x95, 0xa0, 0x1c, 0x7f, 0x42, 0xbf, 0xe2, 0xe0, 0xdb, 0xf3, 0x66, 0x8d, 0x5b, 0x6e,
	0x8f, 0x9e, 0xfe, 0x41, 0xb8, 0x04, 0x38, 0xcd, 0xc6, 0x77, 0x97, 0x5a, 0x1d, 0x6f, 0xbd, 0xed,
	0x76, 0xcb, 0xfe, 0x36, 0xff, 0x3d, 0xdd, 0xa4, 0xf2, 0xef, 0x2e, 0xf6, 0x12, 0xb6, 0xf5, 0x2b,
	0x7c, 0xfd, 0x9d, 0x6c, 0x98, 0x7f, 0xe1, 0xd5, 0xc6, 0x3d, 0x55, 0x9c, 0x3f, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x39, 0xa5, 0x84, 0x69, 0xd3, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	GetData(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*GetDataReponse, error)
	GetDataIds(ctx context.Context, in *DataIdRequest, opts ...grpc.CallOption) (*GetDataIdsResponse, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetData(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*GetDataReponse, error) {
	out := new(GetDataReponse)
	err := c.cc.Invoke(ctx, "/airbloc.rpc.v1.User/GetData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetDataIds(ctx context.Context, in *DataIdRequest, opts ...grpc.CallOption) (*GetDataIdsResponse, error) {
	out := new(GetDataIdsResponse)
	err := c.cc.Invoke(ctx, "/airbloc.rpc.v1.User/GetDataIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	GetData(context.Context, *DataRequest) (*GetDataReponse, error)
	GetDataIds(context.Context, *DataIdRequest) (*GetDataIdsResponse, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.rpc.v1.User/GetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetData(ctx, req.(*DataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetDataIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetDataIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.rpc.v1.User/GetDataIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetDataIds(ctx, req.(*DataIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "airbloc.rpc.v1.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetData",
			Handler:    _User_GetData_Handler,
		},
		{
			MethodName: "GetDataIds",
			Handler:    _User_GetDataIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rpc/v1/server/user.proto",
}
