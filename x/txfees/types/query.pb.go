// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/txfees/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QueryFeeTokensRequest struct {
}

func (m *QueryFeeTokensRequest) Reset()         { *m = QueryFeeTokensRequest{} }
func (m *QueryFeeTokensRequest) String() string { return proto.CompactTextString(m) }
func (*QueryFeeTokensRequest) ProtoMessage()    {}
func (*QueryFeeTokensRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cbc1b48c44dfdd6, []int{0}
}
func (m *QueryFeeTokensRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFeeTokensRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFeeTokensRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFeeTokensRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFeeTokensRequest.Merge(m, src)
}
func (m *QueryFeeTokensRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryFeeTokensRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFeeTokensRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFeeTokensRequest proto.InternalMessageInfo

type QueryFeeTokensResponse struct {
	FeeTokens []FeeToken `protobuf:"bytes,1,rep,name=fee_tokens,json=feeTokens,proto3" json:"fee_tokens" yaml:"fee_tokens"`
}

func (m *QueryFeeTokensResponse) Reset()         { *m = QueryFeeTokensResponse{} }
func (m *QueryFeeTokensResponse) String() string { return proto.CompactTextString(m) }
func (*QueryFeeTokensResponse) ProtoMessage()    {}
func (*QueryFeeTokensResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cbc1b48c44dfdd6, []int{1}
}
func (m *QueryFeeTokensResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFeeTokensResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFeeTokensResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFeeTokensResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFeeTokensResponse.Merge(m, src)
}
func (m *QueryFeeTokensResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryFeeTokensResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFeeTokensResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFeeTokensResponse proto.InternalMessageInfo

func (m *QueryFeeTokensResponse) GetFeeTokens() []FeeToken {
	if m != nil {
		return m.FeeTokens
	}
	return nil
}

type QueryDenomPoolIDRequest struct {
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty" yaml:"denom"`
}

func (m *QueryDenomPoolIDRequest) Reset()         { *m = QueryDenomPoolIDRequest{} }
func (m *QueryDenomPoolIDRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDenomPoolIDRequest) ProtoMessage()    {}
func (*QueryDenomPoolIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cbc1b48c44dfdd6, []int{2}
}
func (m *QueryDenomPoolIDRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomPoolIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomPoolIDRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomPoolIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomPoolIDRequest.Merge(m, src)
}
func (m *QueryDenomPoolIDRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomPoolIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomPoolIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomPoolIDRequest proto.InternalMessageInfo

func (m *QueryDenomPoolIDRequest) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

type QueryDenomPoolIDResponse struct {
	PoolID uint64 `protobuf:"varint,1,opt,name=poolID,proto3" json:"poolID,omitempty" yaml:"pool_id"`
}

func (m *QueryDenomPoolIDResponse) Reset()         { *m = QueryDenomPoolIDResponse{} }
func (m *QueryDenomPoolIDResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDenomPoolIDResponse) ProtoMessage()    {}
func (*QueryDenomPoolIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cbc1b48c44dfdd6, []int{3}
}
func (m *QueryDenomPoolIDResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomPoolIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomPoolIDResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomPoolIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomPoolIDResponse.Merge(m, src)
}
func (m *QueryDenomPoolIDResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomPoolIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomPoolIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomPoolIDResponse proto.InternalMessageInfo

func (m *QueryDenomPoolIDResponse) GetPoolID() uint64 {
	if m != nil {
		return m.PoolID
	}
	return 0
}

type QueryBaseDenomRequest struct {
}

func (m *QueryBaseDenomRequest) Reset()         { *m = QueryBaseDenomRequest{} }
func (m *QueryBaseDenomRequest) String() string { return proto.CompactTextString(m) }
func (*QueryBaseDenomRequest) ProtoMessage()    {}
func (*QueryBaseDenomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cbc1b48c44dfdd6, []int{4}
}
func (m *QueryBaseDenomRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBaseDenomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBaseDenomRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBaseDenomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBaseDenomRequest.Merge(m, src)
}
func (m *QueryBaseDenomRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryBaseDenomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBaseDenomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBaseDenomRequest proto.InternalMessageInfo

type QueryBaseDenomResponse struct {
	BaseDenom string `protobuf:"bytes,1,opt,name=base_denom,json=baseDenom,proto3" json:"base_denom,omitempty" yaml:"base_denom"`
}

func (m *QueryBaseDenomResponse) Reset()         { *m = QueryBaseDenomResponse{} }
func (m *QueryBaseDenomResponse) String() string { return proto.CompactTextString(m) }
func (*QueryBaseDenomResponse) ProtoMessage()    {}
func (*QueryBaseDenomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cbc1b48c44dfdd6, []int{5}
}
func (m *QueryBaseDenomResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBaseDenomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBaseDenomResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBaseDenomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBaseDenomResponse.Merge(m, src)
}
func (m *QueryBaseDenomResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryBaseDenomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBaseDenomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBaseDenomResponse proto.InternalMessageInfo

func (m *QueryBaseDenomResponse) GetBaseDenom() string {
	if m != nil {
		return m.BaseDenom
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryFeeTokensRequest)(nil), "osmosis.txfees.v1beta1.QueryFeeTokensRequest")
	proto.RegisterType((*QueryFeeTokensResponse)(nil), "osmosis.txfees.v1beta1.QueryFeeTokensResponse")
	proto.RegisterType((*QueryDenomPoolIDRequest)(nil), "osmosis.txfees.v1beta1.QueryDenomPoolIDRequest")
	proto.RegisterType((*QueryDenomPoolIDResponse)(nil), "osmosis.txfees.v1beta1.QueryDenomPoolIDResponse")
	proto.RegisterType((*QueryBaseDenomRequest)(nil), "osmosis.txfees.v1beta1.QueryBaseDenomRequest")
	proto.RegisterType((*QueryBaseDenomResponse)(nil), "osmosis.txfees.v1beta1.QueryBaseDenomResponse")
}

func init() {
	proto.RegisterFile("osmosis/txfees/v1beta1/query.proto", fileDescriptor_6cbc1b48c44dfdd6)
}

var fileDescriptor_6cbc1b48c44dfdd6 = []byte{
	// 499 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xd1, 0x6a, 0x13, 0x41,
	0x14, 0xcd, 0x5a, 0x5b, 0xc8, 0x54, 0x44, 0x07, 0xdb, 0xc6, 0x20, 0x9b, 0x30, 0xa8, 0x94, 0x42,
	0x76, 0x4c, 0x55, 0x04, 0xdf, 0x5c, 0x4a, 0x41, 0x04, 0xd1, 0xc5, 0xa7, 0xbe, 0x84, 0x59, 0x73,
	0x77, 0x5d, 0x4c, 0xf6, 0x6e, 0x33, 0xb3, 0xa5, 0x41, 0x7c, 0xf1, 0x0b, 0x04, 0xc1, 0x5f, 0xf0,
	0x57, 0xfa, 0x58, 0xf0, 0xc5, 0xa7, 0x20, 0x89, 0x5f, 0x90, 0x2f, 0x90, 0x9d, 0x9d, 0x4d, 0x42,
	0xd3, 0x6d, 0xfa, 0x96, 0xb9, 0xe7, 0xcc, 0xb9, 0x67, 0xce, 0xc9, 0x12, 0x86, 0xb2, 0x8f, 0x32,
	0x92, 0x5c, 0x9d, 0x06, 0x00, 0x92, 0x9f, 0xb4, 0x7d, 0x50, 0xa2, 0xcd, 0x8f, 0x53, 0x18, 0x0c,
	0x9d, 0x64, 0x80, 0x0a, 0xe9, 0xb6, 0xe1, 0x38, 0x39, 0xc7, 0x31, 0x9c, 0xfa, 0xbd, 0x10, 0x43,
	0xd4, 0x14, 0x9e, 0xfd, 0xca, 0xd9, 0xf5, 0x07, 0x21, 0x62, 0xd8, 0x03, 0x2e, 0x92, 0x88, 0x8b,
	0x38, 0x46, 0x25, 0x54, 0x84, 0xb1, 0x34, 0xa8, 0x6d, 0x50, 0x7d, 0xf2, 0xd3, 0x80, 0x77, 0xd3,
	0x81, 0x26, 0x18, 0xfc, 0x51, 0x89, 0x9f, 0x00, 0x40, 0xe1, 0x67, 0x30, 0x34, 0xb6, 0x43, 0xb6,
	0xde, 0x67, 0x0e, 0x0f, 0x01, 0x3e, 0x64, 0x63, 0xe9, 0xc1, 0x71, 0x0a, 0x52, 0x31, 0x45, 0xb6,
	0x2f, 0x02, 0x32, 0xc1, 0x58, 0x02, 0x3d, 0x22, 0x24, 0x00, 0xe8, 0x68, 0x15, 0x59, 0xb3, 0x9a,
	0x6b, 0xbb, 0x9b, 0xfb, 0x4d, 0xe7, 0xf2, 0xa7, 0x39, 0xc5, 0x75, 0xf7, 0xfe, 0xd9, 0xa8, 0x51,
	0x99, 0x8e, 0x1a, 0x77, 0x87, 0xa2, 0xdf, 0x7b, 0xc9, 0xe6, 0x0a, 0xcc, 0xab, 0x06, 0xc5, 0x0e,
	0xf6, 0x8a, 0xec, 0xe8, 0xad, 0x07, 0x10, 0x63, 0xff, 0x1d, 0x62, 0xef, 0xf5, 0x81, 0x31, 0x44,
	0x1f, 0x93, 0xf5, 0x6e, 0x36, 0xad, 0x59, 0x4d, 0x6b, 0xb7, 0xea, 0xde, 0x99, 0x8e, 0x1a, 0xb7,
	0x72, 0x2d, 0x3d, 0x66, 0x5e, 0x0e, 0xb3, 0x43, 0x52, 0x5b, 0x96, 0x30, 0xd6, 0xf7, 0xc8, 0x46,
	0xa2, 0x27, 0x5a, 0xe4, 0xa6, 0x4b, 0xa7, 0xa3, 0xc6, 0xed, 0x5c, 0x24, 0x9b, 0x77, 0xa2, 0x2e,
	0xf3, 0x0c, 0x63, 0x96, 0x8c, 0x2b, 0x24, 0x68, 0xad, 0x22, 0x99, 0xb7, 0x26, 0x99, 0x05, 0xc0,
	0xc8, 0x3f, 0x23, 0xc4, 0x17, 0x12, 0x3a, 0x8b, 0x3e, 0xb7, 0xe6, 0x6f, 0x9e, 0x63, 0xcc, 0xab,
	0xfa, 0xc5, 0xed, 0xfd, 0xc9, 0x1a, 0x59, 0xd7, 0x82, 0xf4, 0xa7, 0x45, 0xaa, 0xb3, 0xbc, 0x69,
	0xab, 0x2c, 0xd3, 0x4b, 0x0b, 0xab, 0x3b, 0xd7, 0xa5, 0xe7, 0x66, 0xd9, 0xde, 0xb7, 0xdf, 0xff,
	0x7e, 0xdc, 0x78, 0x48, 0x19, 0x2f, 0xff, 0xa7, 0x98, 0x8a, 0xe8, 0x2f, 0x8b, 0x6c, 0x2e, 0xe4,
	0x49, 0xf9, 0x95, 0xbb, 0x96, 0xcb, 0xab, 0x3f, 0xb9, 0xfe, 0x05, 0x63, 0xef, 0xb9, 0xb6, 0xc7,
	0x69, 0xab, 0xcc, 0x9e, 0x0e, 0xb2, 0x63, 0x6a, 0xe3, 0x5f, 0xf4, 0xf1, 0xab, 0x8e, 0x70, 0x56,
	0xcc, 0x8a, 0x08, 0x2f, 0x36, 0xbb, 0x22, 0xc2, 0xa5, 0xbe, 0x57, 0x47, 0x38, 0x6f, 0xdc, 0x7d,
	0x73, 0x36, 0xb6, 0xad, 0xf3, 0xb1, 0x6d, 0xfd, 0x1d, 0xdb, 0xd6, 0xf7, 0x89, 0x5d, 0x39, 0x9f,
	0xd8, 0x95, 0x3f, 0x13, 0xbb, 0x72, 0xd4, 0x0e, 0x23, 0xf5, 0x29, 0xf5, 0x9d, 0x8f, 0xd8, 0x2f,
	0x74, 0x5a, 0x3d, 0xe1, 0xcb, 0x99, 0xe8, 0xc9, 0x0b, 0x7e, 0x5a, 0x28, 0xab, 0x61, 0x02, 0xd2,
	0xdf, 0xd0, 0x1f, 0xef, 0xd3, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x96, 0x44, 0xae, 0x75,
	0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// FeeTokens returns a list of all the whitelisted fee tokens and their
	// corresponding pools It does not include the BaseDenom, which has its own
	// query endpoint
	FeeTokens(ctx context.Context, in *QueryFeeTokensRequest, opts ...grpc.CallOption) (*QueryFeeTokensResponse, error)
	DenomPoolID(ctx context.Context, in *QueryDenomPoolIDRequest, opts ...grpc.CallOption) (*QueryDenomPoolIDResponse, error)
	BaseDenom(ctx context.Context, in *QueryBaseDenomRequest, opts ...grpc.CallOption) (*QueryBaseDenomResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) FeeTokens(ctx context.Context, in *QueryFeeTokensRequest, opts ...grpc.CallOption) (*QueryFeeTokensResponse, error) {
	out := new(QueryFeeTokensResponse)
	err := c.cc.Invoke(ctx, "/osmosis.txfees.v1beta1.Query/FeeTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DenomPoolID(ctx context.Context, in *QueryDenomPoolIDRequest, opts ...grpc.CallOption) (*QueryDenomPoolIDResponse, error) {
	out := new(QueryDenomPoolIDResponse)
	err := c.cc.Invoke(ctx, "/osmosis.txfees.v1beta1.Query/DenomPoolID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BaseDenom(ctx context.Context, in *QueryBaseDenomRequest, opts ...grpc.CallOption) (*QueryBaseDenomResponse, error) {
	out := new(QueryBaseDenomResponse)
	err := c.cc.Invoke(ctx, "/osmosis.txfees.v1beta1.Query/BaseDenom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// FeeTokens returns a list of all the whitelisted fee tokens and their
	// corresponding pools It does not include the BaseDenom, which has its own
	// query endpoint
	FeeTokens(context.Context, *QueryFeeTokensRequest) (*QueryFeeTokensResponse, error)
	DenomPoolID(context.Context, *QueryDenomPoolIDRequest) (*QueryDenomPoolIDResponse, error)
	BaseDenom(context.Context, *QueryBaseDenomRequest) (*QueryBaseDenomResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) FeeTokens(ctx context.Context, req *QueryFeeTokensRequest) (*QueryFeeTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeeTokens not implemented")
}
func (*UnimplementedQueryServer) DenomPoolID(ctx context.Context, req *QueryDenomPoolIDRequest) (*QueryDenomPoolIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomPoolID not implemented")
}
func (*UnimplementedQueryServer) BaseDenom(ctx context.Context, req *QueryBaseDenomRequest) (*QueryBaseDenomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BaseDenom not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_FeeTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFeeTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).FeeTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.txfees.v1beta1.Query/FeeTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).FeeTokens(ctx, req.(*QueryFeeTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DenomPoolID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDenomPoolIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomPoolID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.txfees.v1beta1.Query/DenomPoolID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomPoolID(ctx, req.(*QueryDenomPoolIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BaseDenom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBaseDenomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BaseDenom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.txfees.v1beta1.Query/BaseDenom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BaseDenom(ctx, req.(*QueryBaseDenomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "osmosis.txfees.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FeeTokens",
			Handler:    _Query_FeeTokens_Handler,
		},
		{
			MethodName: "DenomPoolID",
			Handler:    _Query_DenomPoolID_Handler,
		},
		{
			MethodName: "BaseDenom",
			Handler:    _Query_BaseDenom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "osmosis/txfees/v1beta1/query.proto",
}

func (m *QueryFeeTokensRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFeeTokensRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFeeTokensRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryFeeTokensResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFeeTokensResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFeeTokensResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FeeTokens) > 0 {
		for iNdEx := len(m.FeeTokens) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeTokens[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryDenomPoolIDRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomPoolIDRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomPoolIDRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDenomPoolIDResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomPoolIDResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomPoolIDResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PoolID != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.PoolID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryBaseDenomRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBaseDenomRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBaseDenomRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryBaseDenomResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBaseDenomResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBaseDenomResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BaseDenom) > 0 {
		i -= len(m.BaseDenom)
		copy(dAtA[i:], m.BaseDenom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BaseDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryFeeTokensRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryFeeTokensResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.FeeTokens) > 0 {
		for _, e := range m.FeeTokens {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryDenomPoolIDRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDenomPoolIDResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolID != 0 {
		n += 1 + sovQuery(uint64(m.PoolID))
	}
	return n
}

func (m *QueryBaseDenomRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryBaseDenomResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BaseDenom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryFeeTokensRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryFeeTokensRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFeeTokensRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryFeeTokensResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryFeeTokensResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFeeTokensResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeTokens", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeTokens = append(m.FeeTokens, FeeToken{})
			if err := m.FeeTokens[len(m.FeeTokens)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryDenomPoolIDRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryDenomPoolIDRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomPoolIDRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryDenomPoolIDResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryDenomPoolIDResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomPoolIDResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolID", wireType)
			}
			m.PoolID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryBaseDenomRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryBaseDenomRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBaseDenomRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryBaseDenomResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryBaseDenomResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBaseDenomResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
