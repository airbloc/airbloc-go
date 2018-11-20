// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/collections.proto

package api

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
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
	return fileDescriptor_00f79846e51b2b98, []int{0}
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
	return fileDescriptor_00f79846e51b2b98, []int{1}
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
	return fileDescriptor_00f79846e51b2b98, []int{2}
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
	return fileDescriptor_00f79846e51b2b98, []int{3}
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
	return fileDescriptor_00f79846e51b2b98, []int{4}
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
	return fileDescriptor_00f79846e51b2b98, []int{5}
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
	proto.RegisterType((*Policy)(nil), "airbloc.collections.Policy")
	proto.RegisterType((*CreateCollectionRequest)(nil), "airbloc.collections.CreateCollectionRequest")
	proto.RegisterType((*CreateCollectionResponse)(nil), "airbloc.collections.CreateCollectionResponse")
	proto.RegisterType((*ListCollectionRequest)(nil), "airbloc.collections.ListCollectionRequest")
	proto.RegisterType((*CollectionResponse)(nil), "airbloc.collections.CollectionResponse")
	proto.RegisterType((*ListCollectionResponse)(nil), "airbloc.collections.ListCollectionResponse")
}

func init() { proto.RegisterFile("proto/collections.proto", fileDescriptor_00f79846e51b2b98) }

var fileDescriptor_00f79846e51b2b98 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0xce, 0x93, 0x40,
	0x14, 0x0d, 0x5f, 0x5b, 0xf2, 0x71, 0xd1, 0xc4, 0x8c, 0x3f, 0x1f, 0xa9, 0x5d, 0x10, 0x62, 0x62,
	0xa3, 0x16, 0x92, 0x76, 0xef, 0xc2, 0x1a, 0x13, 0x8c, 0x89, 0x64, 0x76, 0xba, 0x31, 0x53, 0xb8,
	0xa1, 0x93, 0x50, 0x06, 0x67, 0xa6, 0x6a, 0x5f, 0xc0, 0x07, 0xf0, 0xd9, 0x7c, 0x20, 0xc3, 0x80,
	0x05, 0x2c, 0x8b, 0xba, 0x22, 0xf7, 0xcc, 0x39, 0x97, 0x73, 0x0e, 0x0c, 0xdc, 0x55, 0x52, 0x68,
	0x11, 0xa5, 0xa2, 0x28, 0x30, 0xd5, 0x5c, 0x94, 0x2a, 0x34, 0x08, 0x79, 0xc8, 0xb8, 0xdc, 0x15,
	0x22, 0x0d, 0x7b, 0x47, 0xc1, 0x2f, 0x0b, 0xec, 0x44, 0x14, 0x3c, 0x3d, 0x91, 0x00, 0xee, 0xbd,
	0x65, 0x9a, 0x25, 0x52, 0x7c, 0xe3, 0x19, 0x4a, 0xcf, 0xf2, 0xad, 0xa5, 0x45, 0x07, 0x18, 0x79,
	0x06, 0xf7, 0xdb, 0x39, 0x45, 0xa5, 0x84, 0xf4, 0x6e, 0x0c, 0x69, 0x08, 0x12, 0x1f, 0xdc, 0x1a,
	0xa0, 0x58, 0xb0, 0x13, 0x4a, 0x6f, 0x62, 0x38, 0x7d, 0x88, 0x2c, 0xc0, 0xa9, 0xc7, 0x8f, 0xdf,
	0x4b, 0x94, 0xde, 0xd4, 0x9c, 0x77, 0x40, 0xf0, 0xd3, 0x82, 0xbb, 0xad, 0x44, 0xa6, 0x71, 0x7b,
	0xb6, 0x4a, 0xf1, 0xeb, 0x11, 0x95, 0x26, 0x8f, 0x60, 0xc6, 0xaa, 0x2a, 0xce, 0x8c, 0x3d, 0x87,
	0x36, 0x03, 0x99, 0xc3, 0xad, 0x4a, 0xf7, 0x78, 0x60, 0x71, 0x66, 0x2c, 0x39, 0xf4, 0x3c, 0x93,
	0x0d, 0xd8, 0x95, 0x49, 0xe8, 0xb9, 0xbe, 0xb5, 0x74, 0xd7, 0x4f, 0xc3, 0x91, 0x22, 0xc2, 0xa6,
	0x04, 0xda, 0x52, 0xdf, 0x4f, 0x6f, 0x27, 0x0f, 0xdc, 0xe0, 0x35, 0x78, 0x97, 0x3e, 0x54, 0x25,
	0x4a, 0x85, 0x75, 0x5d, 0x9d, 0xfe, 0xec, 0x67, 0x80, 0x05, 0x9f, 0xe0, 0xf1, 0x07, 0xae, 0xf4,
	0x65, 0x8a, 0x05, 0x38, 0x15, 0xcb, 0x31, 0x2e, 0x33, 0xfc, 0x61, 0x94, 0x33, 0xda, 0x01, 0xf5,
	0x6a, 0xae, 0xf1, 0xa0, 0x12, 0x94, 0x09, 0xcb, 0xd1, 0x24, 0x9a, 0xd1, 0x01, 0x16, 0xbc, 0x03,
	0x32, 0x62, 0xea, 0xbf, 0xdb, 0x09, 0x4e, 0xf0, 0xe4, 0x5f, 0x8b, 0xdd, 0x2e, 0x2d, 0x34, 0x2b,
	0x5a, 0x7f, 0xcd, 0x40, 0x62, 0x70, 0x7b, 0xb5, 0x79, 0x37, 0xfe, 0x64, 0xe9, 0xae, 0x9f, 0x8f,
	0x56, 0x7a, 0xb9, 0x93, 0xf6, 0xb5, 0xeb, 0xdf, 0x16, 0x40, 0xc7, 0x21, 0x08, 0x76, 0x53, 0x36,
	0x79, 0x35, 0xbe, 0x6e, 0xfc, 0x8f, 0x98, 0xaf, 0xae, 0x64, 0xb7, 0xb1, 0xbe, 0xc0, 0xb4, 0x0e,
	0x4c, 0x5e, 0x8c, 0xca, 0x46, 0x3f, 0xd7, 0xfc, 0xe5, 0x55, 0xdc, 0xe6, 0x05, 0x6f, 0xa2, 0xcf,
	0xab, 0x9c, 0xeb, 0xfd, 0x71, 0x17, 0xa6, 0xe2, 0x10, 0xb5, 0xc2, 0xbf, 0xcf, 0x55, 0x3e, 0xb8,
	0x9a, 0x11, 0xab, 0xf8, 0xce, 0x36, 0xf7, 0x73, 0xf3, 0x27, 0x00, 0x00, 0xff, 0xff, 0x21, 0x03,
	0x59, 0x89, 0xba, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/airbloc.collections.Collection/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionClient) List(ctx context.Context, in *ListCollectionRequest, opts ...grpc.CallOption) (*ListCollectionResponse, error) {
	out := new(ListCollectionResponse)
	err := c.cc.Invoke(ctx, "/airbloc.collections.Collection/List", in, out, opts...)
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
		FullMethod: "/airbloc.collections.Collection/Create",
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
		FullMethod: "/airbloc.collections.Collection/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServer).List(ctx, req.(*ListCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Collection_serviceDesc = grpc.ServiceDesc{
	ServiceName: "airbloc.collections.Collection",
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
	Metadata: "proto/collections.proto",
}
