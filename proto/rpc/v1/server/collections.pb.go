// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/rpc/v1/server/collections.proto

package server

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

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
	PageIndex            int32    `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
	ItemsPerPage         int32    `protobuf:"varint,2,opt,name=itemsPerPage,proto3" json:"itemsPerPage,omitempty"`
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

func (m *ListCollectionRequest) GetPageIndex() int32 {
	if m != nil {
		return m.PageIndex
	}
	return 0
}

func (m *ListCollectionRequest) GetItemsPerPage() int32 {
	if m != nil {
		return m.ItemsPerPage
	}
	return 0
}

type CollectionResponse struct {
	AppId                string   `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	SchemaId             string   `protobuf:"bytes,2,opt,name=schemaId,proto3" json:"schemaId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectionResponse) Reset()         { *m = CollectionResponse{} }
func (m *CollectionResponse) String() string { return proto.CompactTextString(m) }
func (*CollectionResponse) ProtoMessage()    {}
func (*CollectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_902b54d9242d285e, []int{4}
}

func (m *CollectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionResponse.Unmarshal(m, b)
}
func (m *CollectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionResponse.Marshal(b, m, deterministic)
}
func (m *CollectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionResponse.Merge(m, src)
}
func (m *CollectionResponse) XXX_Size() int {
	return xxx_messageInfo_CollectionResponse.Size(m)
}
func (m *CollectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionResponse proto.InternalMessageInfo

func (m *CollectionResponse) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *CollectionResponse) GetSchemaId() string {
	if m != nil {
		return m.SchemaId
	}
	return ""
}

type ListCollectionResponse struct {
	Total                int32                 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Collections          []*CollectionResponse `protobuf:"bytes,2,rep,name=collections,proto3" json:"collections,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListCollectionResponse) Reset()         { *m = ListCollectionResponse{} }
func (m *ListCollectionResponse) String() string { return proto.CompactTextString(m) }
func (*ListCollectionResponse) ProtoMessage()    {}
func (*ListCollectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_902b54d9242d285e, []int{5}
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

func (m *ListCollectionResponse) GetCollections() []*CollectionResponse {
	if m != nil {
		return m.Collections
	}
	return nil
}

func init() {
	proto.RegisterType((*Policy)(nil), "airbloc.rpc.v1.Policy")
	proto.RegisterType((*CreateCollectionRequest)(nil), "airbloc.rpc.v1.CreateCollectionRequest")
	proto.RegisterType((*CreateCollectionResponse)(nil), "airbloc.rpc.v1.CreateCollectionResponse")
	proto.RegisterType((*ListCollectionRequest)(nil), "airbloc.rpc.v1.ListCollectionRequest")
	proto.RegisterType((*CollectionResponse)(nil), "airbloc.rpc.v1.CollectionResponse")
	proto.RegisterType((*ListCollectionResponse)(nil), "airbloc.rpc.v1.ListCollectionResponse")
}

func init() {
	proto.RegisterFile("proto/rpc/v1/server/collections.proto", fileDescriptor_902b54d9242d285e)
}

var fileDescriptor_902b54d9242d285e = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcf, 0x8b, 0xd3, 0x40,
	0x18, 0x25, 0xdb, 0x1f, 0x6c, 0xbf, 0xa8, 0xc8, 0xa0, 0x6b, 0x28, 0x7b, 0x28, 0x83, 0xab, 0xbd,
	0x98, 0xd0, 0xee, 0xdd, 0x83, 0xbb, 0x08, 0x15, 0xc1, 0x30, 0xe0, 0x41, 0x3d, 0x4d, 0x27, 0x1f,
	0xd9, 0x40, 0x9a, 0x19, 0x67, 0x66, 0xa3, 0x7b, 0xf4, 0x5f, 0xf0, 0x2f, 0xf2, 0x4f, 0x93, 0x4c,
	0x62, 0xd3, 0x34, 0x81, 0xe2, 0x29, 0x7c, 0x6f, 0xde, 0xf7, 0x78, 0xef, 0x65, 0x06, 0xae, 0x94,
	0x96, 0x56, 0x46, 0x5a, 0x89, 0xa8, 0x5c, 0x45, 0x06, 0x75, 0x89, 0x3a, 0x12, 0x32, 0xcf, 0x51,
	0xd8, 0x4c, 0x16, 0x26, 0x74, 0xe7, 0xe4, 0x09, 0xcf, 0xf4, 0x36, 0x97, 0x22, 0xd4, 0x4a, 0x84,
	0xe5, 0x8a, 0xfe, 0xf6, 0x60, 0x1a, 0xcb, 0x3c, 0x13, 0x0f, 0x84, 0xc2, 0xa3, 0x5b, 0x6e, 0x79,
	0xac, 0x65, 0x99, 0x25, 0xa8, 0x03, 0x6f, 0xe1, 0x2d, 0x3d, 0xd6, 0xc1, 0xc8, 0x4b, 0x78, 0xdc,
	0xcc, 0x02, 0x8d, 0x91, 0x3a, 0x38, 0x73, 0xa4, 0x2e, 0x48, 0x16, 0xe0, 0x57, 0x00, 0xc3, 0x9c,
	0x3f, 0xa0, 0x0e, 0x46, 0x8e, 0x73, 0x08, 0x91, 0x4b, 0x98, 0x55, 0xe3, 0xa7, 0x1f, 0x05, 0xea,
	0x60, 0xec, 0xce, 0x5b, 0x80, 0xfe, 0xf2, 0xe0, 0xc5, 0x8d, 0x46, 0x6e, 0xf1, 0x66, 0x1f, 0x80,
	0xe1, 0xf7, 0x7b, 0x34, 0x96, 0x3c, 0x83, 0x09, 0x57, 0x6a, 0x93, 0x38, 0x7b, 0x33, 0x56, 0x0f,
	0x64, 0x0e, 0xe7, 0x46, 0xdc, 0xe1, 0x8e, 0x6f, 0x12, 0x67, 0x69, 0xc6, 0xf6, 0x33, 0x09, 0x61,
	0xaa, 0x5c, 0xc2, 0xc0, 0x5f, 0x78, 0x4b, 0x7f, 0x7d, 0x11, 0x76, 0x3b, 0x08, 0xeb, 0xfc, 0xac,
	0x61, 0x7d, 0x18, 0x9f, 0x8f, 0x9e, 0xfa, 0xf4, 0x2d, 0x04, 0x7d, 0x0b, 0x46, 0xc9, 0xc2, 0x60,
	0xd5, 0x54, 0xdb, 0xec, 0xde, 0x4a, 0x07, 0xa3, 0x5f, 0xe0, 0xf9, 0xc7, 0xcc, 0xd8, 0x7e, 0x80,
	0x4b, 0x98, 0x29, 0x9e, 0xe2, 0xa6, 0x48, 0xf0, 0xa7, 0xdb, 0x9c, 0xb0, 0x16, 0xa8, 0xa4, 0x33,
	0x8b, 0x3b, 0x13, 0xa3, 0x8e, 0x79, 0x8a, 0x2e, 0xcc, 0x84, 0x75, 0x30, 0xfa, 0x1e, 0xc8, 0x80,
	0xa9, 0xff, 0x2e, 0x86, 0x5a, 0xb8, 0x38, 0xb6, 0xd8, 0x6a, 0x59, 0x69, 0x79, 0xde, 0xf8, 0xab,
	0x07, 0x72, 0x0b, 0xfe, 0xc1, 0x85, 0x0a, 0xce, 0x16, 0xa3, 0xa5, 0xbf, 0xa6, 0xc7, 0x6d, 0xf6,
	0xe5, 0xd8, 0xe1, 0xda, 0xfa, 0x8f, 0x07, 0xd0, 0x72, 0xc8, 0x37, 0x98, 0xd6, 0x3d, 0x93, 0xd7,
	0x3d, 0xa5, 0xe1, 0x2b, 0x30, 0x5f, 0x9e, 0x26, 0x36, 0x39, 0x3e, 0xc3, 0xb8, 0x4a, 0x48, 0xae,
	0x8e, 0x37, 0x06, 0x7f, 0xcd, 0xfc, 0xd5, 0x29, 0x5a, 0x2d, 0xfb, 0xee, 0xfa, 0xeb, 0x2a, 0xcd,
	0xec, 0xdd, 0xfd, 0x36, 0x14, 0x72, 0x17, 0x35, 0x3b, 0xff, 0xbe, 0x6f, 0x52, 0x19, 0x0d, 0xbc,
	0xc5, 0xed, 0xd4, 0x81, 0xd7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x2b, 0x69, 0xd6, 0xa9,
	0x03, 0x00, 0x00,
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
