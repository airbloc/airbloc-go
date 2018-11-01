// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/exchange.proto

package exchange // import "github.com/airbloc/airbloc-go/exchange"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/airbloc/airbloc-go/common"
import data "github.com/airbloc/airbloc-go/data"

import (
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

type Contract_Type int32

const (
	Contract_RICHARDIAN Contract_Type = 0
	Contract_SMART      Contract_Type = 1
)

var Contract_Type_name = map[int32]string{
	0: "RICHARDIAN",
	1: "SMART",
}
var Contract_Type_value = map[string]int32{
	"RICHARDIAN": 0,
	"SMART":      1,
}

func (x Contract_Type) String() string {
	return proto.EnumName(Contract_Type_name, int32(x))
}
func (Contract_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_exchange_fd549e15ab19de12, []int{0, 0}
}

type Contract struct {
	Type                 Contract_Type   `protobuf:"varint,1,opt,name=type,proto3,enum=airbloc.exchange.Contract_Type" json:"type,omitempty"`
	SmartContractAddress *common.Address `protobuf:"bytes,2,opt,name=smartContractAddress,proto3" json:"smartContractAddress,omitempty"`
	RichardianHash       []byte          `protobuf:"bytes,3,opt,name=richardianHash,proto3" json:"richardianHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Contract) Reset()         { *m = Contract{} }
func (m *Contract) String() string { return proto.CompactTextString(m) }
func (*Contract) ProtoMessage()    {}
func (*Contract) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_fd549e15ab19de12, []int{0}
}
func (m *Contract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contract.Unmarshal(m, b)
}
func (m *Contract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contract.Marshal(b, m, deterministic)
}
func (dst *Contract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contract.Merge(dst, src)
}
func (m *Contract) XXX_Size() int {
	return xxx_messageInfo_Contract.Size(m)
}
func (m *Contract) XXX_DiscardUnknown() {
	xxx_messageInfo_Contract.DiscardUnknown(m)
}

var xxx_messageInfo_Contract proto.InternalMessageInfo

func (m *Contract) GetType() Contract_Type {
	if m != nil {
		return m.Type
	}
	return Contract_RICHARDIAN
}

func (m *Contract) GetSmartContractAddress() *common.Address {
	if m != nil {
		return m.SmartContractAddress
	}
	return nil
}

func (m *Contract) GetRichardianHash() []byte {
	if m != nil {
		return m.RichardianHash
	}
	return nil
}

type OrderRequest struct {
	Data                 *data.Batch `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Contract             *Contract   `protobuf:"bytes,2,opt,name=contract,proto3" json:"contract,omitempty"`
	Offeror              string      `protobuf:"bytes,3,opt,name=offeror,proto3" json:"offeror,omitempty"`
	Offeree              string      `protobuf:"bytes,4,opt,name=offeree,proto3" json:"offeree,omitempty"`
	Option               []string    `protobuf:"bytes,5,rep,name=option,proto3" json:"option,omitempty"`
	Amount               float64     `protobuf:"fixed64,6,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *OrderRequest) Reset()         { *m = OrderRequest{} }
func (m *OrderRequest) String() string { return proto.CompactTextString(m) }
func (*OrderRequest) ProtoMessage()    {}
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_fd549e15ab19de12, []int{1}
}
func (m *OrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderRequest.Unmarshal(m, b)
}
func (m *OrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderRequest.Marshal(b, m, deterministic)
}
func (dst *OrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderRequest.Merge(dst, src)
}
func (m *OrderRequest) XXX_Size() int {
	return xxx_messageInfo_OrderRequest.Size(m)
}
func (m *OrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OrderRequest proto.InternalMessageInfo

func (m *OrderRequest) GetData() *data.Batch {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *OrderRequest) GetContract() *Contract {
	if m != nil {
		return m.Contract
	}
	return nil
}

func (m *OrderRequest) GetOfferor() string {
	if m != nil {
		return m.Offeror
	}
	return ""
}

func (m *OrderRequest) GetOfferee() string {
	if m != nil {
		return m.Offeree
	}
	return ""
}

func (m *OrderRequest) GetOption() []string {
	if m != nil {
		return m.Option
	}
	return nil
}

func (m *OrderRequest) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type OrderId struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderId) Reset()         { *m = OrderId{} }
func (m *OrderId) String() string { return proto.CompactTextString(m) }
func (*OrderId) ProtoMessage()    {}
func (*OrderId) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_fd549e15ab19de12, []int{2}
}
func (m *OrderId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderId.Unmarshal(m, b)
}
func (m *OrderId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderId.Marshal(b, m, deterministic)
}
func (dst *OrderId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderId.Merge(dst, src)
}
func (m *OrderId) XXX_Size() int {
	return xxx_messageInfo_OrderId.Size(m)
}
func (m *OrderId) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderId.DiscardUnknown(m)
}

var xxx_messageInfo_OrderId proto.InternalMessageInfo

func (m *OrderId) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

type SettleMessage struct {
	OrderId              *OrderId `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SettleMessage) Reset()         { *m = SettleMessage{} }
func (m *SettleMessage) String() string { return proto.CompactTextString(m) }
func (*SettleMessage) ProtoMessage()    {}
func (*SettleMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_fd549e15ab19de12, []int{3}
}
func (m *SettleMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettleMessage.Unmarshal(m, b)
}
func (m *SettleMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettleMessage.Marshal(b, m, deterministic)
}
func (dst *SettleMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettleMessage.Merge(dst, src)
}
func (m *SettleMessage) XXX_Size() int {
	return xxx_messageInfo_SettleMessage.Size(m)
}
func (m *SettleMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SettleMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SettleMessage proto.InternalMessageInfo

func (m *SettleMessage) GetOrderId() *OrderId {
	if m != nil {
		return m.OrderId
	}
	return nil
}

type ContractId struct {
	ContractId           string   `protobuf:"bytes,1,opt,name=ContractId,proto3" json:"ContractId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContractId) Reset()         { *m = ContractId{} }
func (m *ContractId) String() string { return proto.CompactTextString(m) }
func (*ContractId) ProtoMessage()    {}
func (*ContractId) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_fd549e15ab19de12, []int{4}
}
func (m *ContractId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractId.Unmarshal(m, b)
}
func (m *ContractId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractId.Marshal(b, m, deterministic)
}
func (dst *ContractId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractId.Merge(dst, src)
}
func (m *ContractId) XXX_Size() int {
	return xxx_messageInfo_ContractId.Size(m)
}
func (m *ContractId) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractId.DiscardUnknown(m)
}

var xxx_messageInfo_ContractId proto.InternalMessageInfo

func (m *ContractId) GetContractId() string {
	if m != nil {
		return m.ContractId
	}
	return ""
}

type SettleResult struct {
	ContractId           *ContractId `protobuf:"bytes,1,opt,name=contractId,proto3" json:"contractId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SettleResult) Reset()         { *m = SettleResult{} }
func (m *SettleResult) String() string { return proto.CompactTextString(m) }
func (*SettleResult) ProtoMessage()    {}
func (*SettleResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_fd549e15ab19de12, []int{5}
}
func (m *SettleResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettleResult.Unmarshal(m, b)
}
func (m *SettleResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettleResult.Marshal(b, m, deterministic)
}
func (dst *SettleResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettleResult.Merge(dst, src)
}
func (m *SettleResult) XXX_Size() int {
	return xxx_messageInfo_SettleResult.Size(m)
}
func (m *SettleResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SettleResult.DiscardUnknown(m)
}

var xxx_messageInfo_SettleResult proto.InternalMessageInfo

func (m *SettleResult) GetContractId() *ContractId {
	if m != nil {
		return m.ContractId
	}
	return nil
}

func init() {
	proto.RegisterType((*Contract)(nil), "airbloc.exchange.Contract")
	proto.RegisterType((*OrderRequest)(nil), "airbloc.exchange.OrderRequest")
	proto.RegisterType((*OrderId)(nil), "airbloc.exchange.OrderId")
	proto.RegisterType((*SettleMessage)(nil), "airbloc.exchange.SettleMessage")
	proto.RegisterType((*ContractId)(nil), "airbloc.exchange.ContractId")
	proto.RegisterType((*SettleResult)(nil), "airbloc.exchange.SettleResult")
	proto.RegisterEnum("airbloc.exchange.Contract_Type", Contract_Type_name, Contract_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExchangeClient is the client API for Exchange service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExchangeClient interface {
	Order(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderId, error)
	Settle(ctx context.Context, in *SettleMessage, opts ...grpc.CallOption) (*SettleResult, error)
	Reject(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*common.Result, error)
	// after Open()
	CloseOrder(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*common.Result, error)
}

type exchangeClient struct {
	cc *grpc.ClientConn
}

func NewExchangeClient(cc *grpc.ClientConn) ExchangeClient {
	return &exchangeClient{cc}
}

func (c *exchangeClient) Order(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderId, error) {
	out := new(OrderId)
	err := c.cc.Invoke(ctx, "/airbloc.exchange.Exchange/Order", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) Settle(ctx context.Context, in *SettleMessage, opts ...grpc.CallOption) (*SettleResult, error) {
	out := new(SettleResult)
	err := c.cc.Invoke(ctx, "/airbloc.exchange.Exchange/Settle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) Reject(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*common.Result, error) {
	out := new(common.Result)
	err := c.cc.Invoke(ctx, "/airbloc.exchange.Exchange/Reject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) CloseOrder(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*common.Result, error) {
	out := new(common.Result)
	err := c.cc.Invoke(ctx, "/airbloc.exchange.Exchange/CloseOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExchangeServer is the server API for Exchange service.
type ExchangeServer interface {
	Order(context.Context, *OrderRequest) (*OrderId, error)
	Settle(context.Context, *SettleMessage) (*SettleResult, error)
	Reject(context.Context, *OrderId) (*common.Result, error)
	// after Open()
	CloseOrder(context.Context, *OrderId) (*common.Result, error)
}

func RegisterExchangeServer(s *grpc.Server, srv ExchangeServer) {
	s.RegisterService(&_Exchange_serviceDesc, srv)
}

func _Exchange_Order_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).Order(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.exchange.Exchange/Order",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).Order(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_Settle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SettleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).Settle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.exchange.Exchange/Settle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).Settle(ctx, req.(*SettleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_Reject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).Reject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.exchange.Exchange/Reject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).Reject(ctx, req.(*OrderId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_CloseOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).CloseOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.exchange.Exchange/CloseOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).CloseOrder(ctx, req.(*OrderId))
	}
	return interceptor(ctx, in, info, handler)
}

var _Exchange_serviceDesc = grpc.ServiceDesc{
	ServiceName: "airbloc.exchange.Exchange",
	HandlerType: (*ExchangeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Order",
			Handler:    _Exchange_Order_Handler,
		},
		{
			MethodName: "Settle",
			Handler:    _Exchange_Settle_Handler,
		},
		{
			MethodName: "Reject",
			Handler:    _Exchange_Reject_Handler,
		},
		{
			MethodName: "CloseOrder",
			Handler:    _Exchange_CloseOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/exchange.proto",
}

func init() { proto.RegisterFile("proto/exchange.proto", fileDescriptor_exchange_fd549e15ab19de12) }

var fileDescriptor_exchange_fd549e15ab19de12 = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x6d, 0x92, 0x26, 0x93, 0x10, 0x85, 0xa5, 0x2a, 0xc6, 0x42, 0xc1, 0x18, 0xa9, 0xf8,
	0x00, 0xae, 0xe4, 0x48, 0x5c, 0x40, 0x42, 0x49, 0x8a, 0x54, 0x0b, 0x0a, 0xd2, 0xb6, 0x27, 0x6e,
	0x1b, 0x7b, 0x1a, 0x1b, 0x25, 0xde, 0xb0, 0xbb, 0x91, 0xe8, 0x97, 0xf0, 0x61, 0x1c, 0xf8, 0x1d,
	0xe4, 0xf5, 0x3a, 0x4d, 0x42, 0x7d, 0xe0, 0x94, 0xcc, 0x9b, 0xb7, 0x6f, 0xde, 0x1b, 0xaf, 0x0d,
	0xc7, 0x2b, 0xc1, 0x15, 0x3f, 0xc3, 0x9f, 0x71, 0xca, 0xf2, 0x39, 0x06, 0xba, 0x24, 0x03, 0x96,
	0x89, 0xd9, 0x82, 0xc7, 0x41, 0x85, 0x3b, 0x83, 0x92, 0x97, 0x30, 0xc5, 0x4a, 0x8e, 0xf3, 0xa8,
	0x44, 0xd4, 0xed, 0x0a, 0x65, 0x09, 0x79, 0x7f, 0x2c, 0x68, 0x4f, 0x79, 0xae, 0x04, 0x8b, 0x15,
	0x19, 0x41, 0xa3, 0xe8, 0xd9, 0x96, 0x6b, 0xf9, 0xfd, 0xf0, 0x79, 0xb0, 0x2f, 0x19, 0x54, 0xcc,
	0xe0, 0xfa, 0x76, 0x85, 0x54, 0x93, 0xc9, 0x27, 0x38, 0x96, 0x4b, 0x26, 0x54, 0xd5, 0x1b, 0x27,
	0x89, 0x40, 0x29, 0xed, 0x03, 0xd7, 0xf2, 0xbb, 0xe1, 0x93, 0x8d, 0x48, 0xcc, 0x97, 0x4b, 0x9e,
	0x07, 0xa6, 0x4d, 0xef, 0x3d, 0x44, 0x4e, 0xa1, 0x2f, 0xb2, 0x38, 0x65, 0x22, 0xc9, 0x58, 0x7e,
	0xc1, 0x64, 0x6a, 0x1f, 0xba, 0x96, 0xdf, 0xa3, 0x7b, 0xa8, 0xf7, 0x02, 0x1a, 0x85, 0x05, 0xd2,
	0x07, 0xa0, 0xd1, 0xf4, 0x62, 0x4c, 0xcf, 0xa3, 0xf1, 0x97, 0xc1, 0x03, 0xd2, 0x81, 0xe6, 0xd5,
	0xe5, 0x98, 0x5e, 0x0f, 0x2c, 0xef, 0xb7, 0x05, 0xbd, 0xaf, 0x22, 0x41, 0x41, 0xf1, 0xc7, 0x1a,
	0xa5, 0x22, 0xaf, 0xa0, 0x51, 0xec, 0x42, 0xa7, 0xeb, 0x86, 0x8f, 0x37, 0xc6, 0xf4, 0x82, 0x26,
	0x4c, 0xc5, 0x29, 0xd5, 0x04, 0xf2, 0x16, 0xda, 0xb1, 0xf1, 0x65, 0x52, 0x38, 0xf5, 0xab, 0xa0,
	0x1b, 0x2e, 0xb1, 0xe1, 0x88, 0xdf, 0xdc, 0xa0, 0xe0, 0x42, 0xbb, 0xee, 0xd0, 0xaa, 0xdc, 0x74,
	0x10, 0xed, 0xc6, 0x56, 0x07, 0x91, 0x9c, 0x40, 0x8b, 0xaf, 0x54, 0xc6, 0x73, 0xbb, 0xe9, 0x1e,
	0xfa, 0x1d, 0x6a, 0xaa, 0x02, 0x67, 0x4b, 0xbe, 0xce, 0x95, 0xdd, 0x72, 0x2d, 0xdf, 0xa2, 0xa6,
	0xf2, 0x5e, 0xc2, 0x91, 0x0e, 0x15, 0x25, 0x5a, 0xb4, 0xfc, 0xab, 0x23, 0x15, 0xa2, 0x65, 0xe9,
	0x9d, 0xc3, 0xc3, 0x2b, 0x54, 0x6a, 0x81, 0x97, 0x28, 0x25, 0x9b, 0x23, 0x19, 0xed, 0x52, 0xbb,
	0xe1, 0xd3, 0x7f, 0x03, 0x19, 0xd9, 0x3b, 0x95, 0xd7, 0x00, 0x55, 0xc8, 0x28, 0x21, 0xc3, 0xed,
	0xca, 0x0c, 0xdc, 0x42, 0xbc, 0xcf, 0xd0, 0x2b, 0x67, 0x52, 0x94, 0xeb, 0x85, 0x22, 0xef, 0x01,
	0xe2, 0x5d, 0x7e, 0x37, 0x7c, 0x56, 0xbf, 0xc6, 0x28, 0xa1, 0x5b, 0xfc, 0xf0, 0xd7, 0x01, 0xb4,
	0x3f, 0x1a, 0x0e, 0x99, 0x40, 0x53, 0x9b, 0x23, 0xc3, 0x1a, 0xd7, 0xe6, 0x09, 0x3b, 0xf5, 0xa9,
	0x48, 0x04, 0xad, 0xd2, 0x1e, 0xb9, 0xe7, 0x5a, 0xef, 0x2c, 0xcb, 0x19, 0xd6, 0x11, 0x4c, 0xb2,
	0x77, 0xd0, 0xa2, 0xf8, 0x1d, 0x63, 0x45, 0xea, 0xe7, 0x39, 0x27, 0xfb, 0xf7, 0xde, 0x1c, 0xfe,
	0x00, 0x30, 0x5d, 0x70, 0x89, 0x65, 0xa0, 0xff, 0x17, 0x98, 0xf8, 0xdf, 0x4e, 0xe7, 0x99, 0x4a,
	0xd7, 0xb3, 0x02, 0x3f, 0x33, 0x9c, 0xea, 0xf7, 0xcd, 0xfc, 0xee, 0xbb, 0x30, 0x6b, 0xe9, 0x37,
	0x7c, 0xf4, 0x37, 0x00, 0x00, 0xff, 0xff, 0x05, 0xd8, 0xde, 0x72, 0x30, 0x04, 0x00, 0x00,
}
