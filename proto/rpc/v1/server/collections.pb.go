// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/rpc/v1/server/collections.proto

package server

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Policy struct {
	DataProvider         float64  `protobuf:"fixed64,1,opt,name=DataProvider,proto3" json:"DataProvider,omitempty"`
	DataProcessor        float64  `protobuf:"fixed64,2,opt,name=DataProcessor,proto3" json:"DataProcessor,omitempty"`
	DataRelayer          float64  `protobuf:"fixed64,3,opt,name=DataRelayer,proto3" json:"DataRelayer,omitempty"`
	DataOwner            float64  `protobuf:"fixed64,4,opt,name=DataOwner,proto3" json:"DataOwner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_902b54d9242d285e, []int{0}
}

func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetDataProvider() float64 {
	if m != nil {
		return m.DataProvider
	}
	return 0
}

func (m *Policy) GetDataProcessor() float64 {
	if m != nil {
		return m.DataProcessor
	}
	return 0
}

func (m *Policy) GetDataRelayer() float64 {
	if m != nil {
		return m.DataRelayer
	}
	return 0
}

func (m *Policy) GetDataOwner() float64 {
	if m != nil {
		return m.DataOwner
	}
	return 0
}

type CreateCollectionRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	SchemaId             string   `protobuf:"bytes,2,opt,name=schemaId,proto3" json:"schemaId,omitempty"`
	Policy               *Policy  `protobuf:"bytes,11,opt,name=policy,proto3" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCollectionRequest) Reset()         { *m = CreateCollectionRequest{} }
func (m *CreateCollectionRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCollectionRequest) ProtoMessage()    {}
func (*CreateCollectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_902b54d9242d285e, []int{1}
}

func (m *CreateCollectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCollectionRequest.Unmarshal(m, b)
}
func (m *CreateCollectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCollectionRequest.Marshal(b, m, deterministic)
}
func (m *CreateCollectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCollectionRequest.Merge(m, src)
}
func (m *CreateCollectionRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCollectionRequest.Size(m)
}
func (m *CreateCollectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCollectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCollectionRequest proto.InternalMessageInfo

func (m *CreateCollectionRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *CreateCollectionRequest) GetSchemaId() string {
	if m != nil {
		return m.SchemaId
	}
	return ""
}

func (m *CreateCollectionRequest) GetPolicy() *Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type CreateCollectionResponse struct {
	CollectionId         string   `protobuf:"bytes,1,opt,name=collectionId,proto3" json:"collectionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCollectionResponse) Reset()         { *m = CreateCollectionResponse{} }
func (m *CreateCollectionResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCollectionResponse) ProtoMessage()    {}
func (*CreateCollectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_902b54d9242d285e, []int{2}
}

func (m *CreateCollectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCollectionResponse.Unmarshal(m, b)
}
func (m *CreateCollectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCollectionResponse.Marshal(b, m, deterministic)
}
func (m *CreateCollectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCollectionResponse.Merge(m, src)
}
func (m *CreateCollectionResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCollectionResponse.Size(m)
}
func (m *CreateCollectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCollectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCollectionResponse proto.InternalMessageInfo

func (m *CreateCollectionResponse) GetCollectionId() string {
	if m != nil {
		return m.CollectionId
	}
	return ""
}

type ListCollectionRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCollectionRequest) Reset()         { *m = ListCollectionRequest{} }
func (m *ListCollectionRequest) String() string { return proto.CompactTextString(m) }
func (*ListCollectionRequest) ProtoMessage()    {}
func (*ListCollectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_902b54d9242d285e, []int{3}
}

func (m *ListCollectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCollectionRequest.Unmarshal(m, b)
}
func (m *ListCollectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCollectionRequest.Marshal(b, m, deterministic)
}
func (m *ListCollectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCollectionRequest.Merge(m, src)
}
func (m *ListCollectionRequest) XXX_Size() int {
	return xxx_messageInfo_ListCollectionRequest.Size(m)
}
func (m *ListCollectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCollectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCollectionRequest proto.InternalMessageInfo

func (m *ListCollectionRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

type ListCollectionResponse struct {
	Total                int32                                `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Collections          []*ListCollectionResponse_Collection `protobuf:"bytes,2,rep,name=collections,proto3" json:"collections,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *ListCollectionResponse) Reset()         { *m = ListCollectionResponse{} }
func (m *ListCollectionResponse) String() string { return proto.CompactTextString(m) }
func (*ListCollectionResponse) ProtoMessage()    {}
func (*ListCollectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_902b54d9242d285e, []int{4}
}

func (m *ListCollectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCollectionResponse.Unmarshal(m, b)
}
func (m *ListCollectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCollectionResponse.Marshal(b, m, deterministic)
}
func (m *ListCollectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCollectionResponse.Merge(m, src)
}
func (m *ListCollectionResponse) XXX_Size() int {
	return xxx_messageInfo_ListCollectionResponse.Size(m)
}
func (m *ListCollectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCollectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCollectionResponse proto.InternalMessageInfo

func (m *ListCollectionResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListCollectionResponse) GetCollections() []*ListCollectionResponse_Collection {
	if m != nil {
		return m.Collections
	}
	return nil
}

type ListCollectionResponse_Collection struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SchemaId             string   `protobuf:"bytes,2,opt,name=schemaId,proto3" json:"schemaId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCollectionResponse_Collection) Reset()         { *m = ListCollectionResponse_Collection{} }
func (m *ListCollectionResponse_Collection) String() string { return proto.CompactTextString(m) }
func (*ListCollectionResponse_Collection) ProtoMessage()    {}
func (*ListCollectionResponse_Collection) Descriptor() ([]byte, []int) {
	return fileDescriptor_902b54d9242d285e, []int{4, 0}
}

func (m *ListCollectionResponse_Collection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCollectionResponse_Collection.Unmarshal(m, b)
}
func (m *ListCollectionResponse_Collection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCollectionResponse_Collection.Marshal(b, m, deterministic)
}
func (m *ListCollectionResponse_Collection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCollectionResponse_Collection.Merge(m, src)
}
func (m *ListCollectionResponse_Collection) XXX_Size() int {
	return xxx_messageInfo_ListCollectionResponse_Collection.Size(m)
}
func (m *ListCollectionResponse_Collection) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCollectionResponse_Collection.DiscardUnknown(m)
}

var xxx_messageInfo_ListCollectionResponse_Collection proto.InternalMessageInfo

func (m *ListCollectionResponse_Collection) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ListCollectionResponse_Collection) GetSchemaId() string {
	if m != nil {
		return m.SchemaId
	}
	return ""
}

func init() {
	proto.RegisterType((*Policy)(nil), "airbloc.rpc.v1.Policy")
	proto.RegisterType((*CreateCollectionRequest)(nil), "airbloc.rpc.v1.CreateCollectionRequest")
	proto.RegisterType((*CreateCollectionResponse)(nil), "airbloc.rpc.v1.CreateCollectionResponse")
	proto.RegisterType((*ListCollectionRequest)(nil), "airbloc.rpc.v1.ListCollectionRequest")
	proto.RegisterType((*ListCollectionResponse)(nil), "airbloc.rpc.v1.ListCollectionResponse")
	proto.RegisterType((*ListCollectionResponse_Collection)(nil), "airbloc.rpc.v1.ListCollectionResponse.Collection")
}

func init() {
	proto.RegisterFile("proto/rpc/v1/server/collections.proto", fileDescriptor_902b54d9242d285e)
}

var fileDescriptor_902b54d9242d285e = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0xae, 0xd2, 0x40,
	0x14, 0xcd, 0x14, 0x68, 0x1e, 0xb7, 0xfa, 0x62, 0x26, 0xcf, 0x67, 0x43, 0x5c, 0x34, 0x8d, 0x4f,
	0xbb, 0x61, 0x9a, 0xc2, 0xc6, 0x95, 0x0b, 0x71, 0x83, 0x31, 0x91, 0xd4, 0xb8, 0xd1, 0xd5, 0x30,
	0x4c, 0x60, 0x92, 0xc2, 0xd4, 0x99, 0xa1, 0x86, 0xa5, 0xbf, 0xe0, 0xff, 0x98, 0xf8, 0x69, 0xa6,
	0xd3, 0x02, 0x2d, 0xa2, 0xb0, 0x22, 0xe7, 0xdc, 0x73, 0x2f, 0xe7, 0x9c, 0x49, 0xe1, 0x21, 0x57,
	0xd2, 0xc8, 0x58, 0xe5, 0x2c, 0x2e, 0x92, 0x58, 0x73, 0x55, 0x70, 0x15, 0x33, 0x99, 0x65, 0x9c,
	0x19, 0x21, 0x37, 0x9a, 0xd8, 0x39, 0xbe, 0xa5, 0x42, 0xcd, 0x33, 0xc9, 0x88, 0xca, 0x19, 0x29,
	0x92, 0xf0, 0x27, 0x02, 0x77, 0x26, 0x33, 0xc1, 0x76, 0x38, 0x84, 0x47, 0xef, 0xa8, 0xa1, 0x33,
	0x25, 0x0b, 0xb1, 0xe0, 0xca, 0x47, 0x01, 0x8a, 0x50, 0xda, 0xe2, 0xf0, 0x0b, 0x78, 0x5c, 0x63,
	0xc6, 0xb5, 0x96, 0xca, 0x77, 0xac, 0xa8, 0x4d, 0xe2, 0x00, 0xbc, 0x92, 0x48, 0x79, 0x46, 0x77,
	0x5c, 0xf9, 0x1d, 0xab, 0x69, 0x52, 0xf8, 0x39, 0xf4, 0x4b, 0xf8, 0xf1, 0xfb, 0x86, 0x2b, 0xbf,
	0x6b, 0xe7, 0x47, 0x22, 0xfc, 0x81, 0xe0, 0xd9, 0x44, 0x71, 0x6a, 0xf8, 0xe4, 0x10, 0x20, 0xe5,
	0xdf, 0xb6, 0x5c, 0x1b, 0x7c, 0x07, 0x3d, 0x9a, 0xe7, 0xd3, 0x85, 0xb5, 0xd7, 0x4f, 0x2b, 0x80,
	0x07, 0x70, 0xa3, 0xd9, 0x8a, 0xaf, 0xe9, 0x74, 0x61, 0x2d, 0xf5, 0xd3, 0x03, 0xc6, 0x04, 0xdc,
	0xdc, 0x26, 0xf4, 0xbd, 0x00, 0x45, 0xde, 0xe8, 0x9e, 0xb4, 0x3b, 0x20, 0x55, 0xfe, 0xb4, 0x56,
	0xbd, 0xef, 0xde, 0x74, 0x9e, 0x78, 0xe1, 0x1b, 0xf0, 0xff, 0xb6, 0xa0, 0x73, 0xb9, 0xd1, 0xbc,
	0x6c, 0xea, 0xd8, 0xec, 0xc1, 0x4a, 0x8b, 0x0b, 0x87, 0xf0, 0xf4, 0x83, 0xd0, 0xe6, 0xca, 0x00,
	0xe1, 0x2f, 0x04, 0xf7, 0xa7, 0xfa, 0xfa, 0xdf, 0xee, 0xa0, 0x67, 0xa4, 0xa1, 0x99, 0x5d, 0xe8,
	0xa5, 0x15, 0xc0, 0x9f, 0xc0, 0x6b, 0xbc, 0xae, 0xef, 0x04, 0x9d, 0xc8, 0x1b, 0x25, 0xa7, 0xd1,
	0xce, 0x9f, 0x24, 0x0d, 0xaa, 0x79, 0x65, 0xf0, 0x1a, 0xe0, 0x38, 0xc2, 0xb7, 0xe0, 0x88, 0xbd,
	0x4d, 0x47, 0xfc, 0xb7, 0xe4, 0xd1, 0x6f, 0xd4, 0x5a, 0xfd, 0x0a, 0x6e, 0xd5, 0x1e, 0x7e, 0x75,
	0x6a, 0xe9, 0x1f, 0x0f, 0x3b, 0x88, 0x2e, 0x0b, 0xeb, 0x42, 0x3e, 0x43, 0xb7, 0xcc, 0x85, 0x1f,
	0x2e, 0xa5, 0xad, 0x0e, 0xbf, 0xbc, 0xae, 0x94, 0xb7, 0xe3, 0x2f, 0xc9, 0x52, 0x98, 0xd5, 0x76,
	0x4e, 0x98, 0x5c, 0xc7, 0xf5, 0xce, 0xfe, 0x77, 0xb8, 0x94, 0xf1, 0x99, 0x2f, 0x6c, 0xee, 0x5a,
	0x72, 0xfc, 0x27, 0x00, 0x00, 0xff, 0xff, 0xad, 0x67, 0x9e, 0x7f, 0x7f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CollectionClient is the client API for Collection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CollectionClient interface {
	Create(ctx context.Context, in *CreateCollectionRequest, opts ...grpc.CallOption) (*CreateCollectionResponse, error)
	List(ctx context.Context, in *ListCollectionRequest, opts ...grpc.CallOption) (*ListCollectionResponse, error)
}

type collectionClient struct {
	cc *grpc.ClientConn
}

func NewCollectionClient(cc *grpc.ClientConn) CollectionClient {
	return &collectionClient{cc}
}

func (c *collectionClient) Create(ctx context.Context, in *CreateCollectionRequest, opts ...grpc.CallOption) (*CreateCollectionResponse, error) {
	out := new(CreateCollectionResponse)
	err := c.cc.Invoke(ctx, "/airbloc.rpc.v1.Collection/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionClient) List(ctx context.Context, in *ListCollectionRequest, opts ...grpc.CallOption) (*ListCollectionResponse, error) {
	out := new(ListCollectionResponse)
	err := c.cc.Invoke(ctx, "/airbloc.rpc.v1.Collection/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectionServer is the server API for Collection service.
type CollectionServer interface {
	Create(context.Context, *CreateCollectionRequest) (*CreateCollectionResponse, error)
	List(context.Context, *ListCollectionRequest) (*ListCollectionResponse, error)
}

func RegisterCollectionServer(s *grpc.Server, srv CollectionServer) {
	s.RegisterService(&_Collection_serviceDesc, srv)
}

func _Collection_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.rpc.v1.Collection/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServer).Create(ctx, req.(*CreateCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collection_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.rpc.v1.Collection/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServer).List(ctx, req.(*ListCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Collection_serviceDesc = grpc.ServiceDesc{
	ServiceName: "airbloc.rpc.v1.Collection",
	HandlerType: (*CollectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Collection_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Collection_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rpc/v1/server/collections.proto",
}
