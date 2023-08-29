// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/cosmwasmpool/v1beta1/query.proto

package queryproto

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types1 "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	types "github.com/arnac-io/osmosis/x/cosmwasmpool/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// =============================== ContractInfoByPoolId
type ParamsRequest struct {
}

func (m *ParamsRequest) Reset()         { *m = ParamsRequest{} }
func (m *ParamsRequest) String() string { return proto.CompactTextString(m) }
func (*ParamsRequest) ProtoMessage()    {}
func (*ParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_733c758985c393b2, []int{0}
}
func (m *ParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamsRequest.Merge(m, src)
}
func (m *ParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *ParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ParamsRequest proto.InternalMessageInfo

type ParamsResponse struct {
	Params types.Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *ParamsResponse) Reset()         { *m = ParamsResponse{} }
func (m *ParamsResponse) String() string { return proto.CompactTextString(m) }
func (*ParamsResponse) ProtoMessage()    {}
func (*ParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_733c758985c393b2, []int{1}
}
func (m *ParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamsResponse.Merge(m, src)
}
func (m *ParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *ParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ParamsResponse proto.InternalMessageInfo

func (m *ParamsResponse) GetParams() types.Params {
	if m != nil {
		return m.Params
	}
	return types.Params{}
}

// =============================== Pools
type PoolsRequest struct {
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *PoolsRequest) Reset()         { *m = PoolsRequest{} }
func (m *PoolsRequest) String() string { return proto.CompactTextString(m) }
func (*PoolsRequest) ProtoMessage()    {}
func (*PoolsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_733c758985c393b2, []int{2}
}
func (m *PoolsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolsRequest.Merge(m, src)
}
func (m *PoolsRequest) XXX_Size() int {
	return m.Size()
}
func (m *PoolsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PoolsRequest proto.InternalMessageInfo

func (m *PoolsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type PoolsResponse struct {
	Pools []*types1.Any `protobuf:"bytes,1,rep,name=pools,proto3" json:"pools,omitempty"`
	// pagination defines the pagination in the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *PoolsResponse) Reset()         { *m = PoolsResponse{} }
func (m *PoolsResponse) String() string { return proto.CompactTextString(m) }
func (*PoolsResponse) ProtoMessage()    {}
func (*PoolsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_733c758985c393b2, []int{3}
}
func (m *PoolsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolsResponse.Merge(m, src)
}
func (m *PoolsResponse) XXX_Size() int {
	return m.Size()
}
func (m *PoolsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PoolsResponse proto.InternalMessageInfo

func (m *PoolsResponse) GetPools() []*types1.Any {
	if m != nil {
		return m.Pools
	}
	return nil
}

func (m *PoolsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// =============================== ContractInfoByPoolId
type ContractInfoByPoolIdRequest struct {
	// pool_id is the pool id of the requested pool.
	PoolId uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty" yaml:"pool_id"`
}

func (m *ContractInfoByPoolIdRequest) Reset()         { *m = ContractInfoByPoolIdRequest{} }
func (m *ContractInfoByPoolIdRequest) String() string { return proto.CompactTextString(m) }
func (*ContractInfoByPoolIdRequest) ProtoMessage()    {}
func (*ContractInfoByPoolIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_733c758985c393b2, []int{4}
}
func (m *ContractInfoByPoolIdRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractInfoByPoolIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractInfoByPoolIdRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractInfoByPoolIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractInfoByPoolIdRequest.Merge(m, src)
}
func (m *ContractInfoByPoolIdRequest) XXX_Size() int {
	return m.Size()
}
func (m *ContractInfoByPoolIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractInfoByPoolIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContractInfoByPoolIdRequest proto.InternalMessageInfo

func (m *ContractInfoByPoolIdRequest) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

type ContractInfoByPoolIdResponse struct {
	// contract_address is the pool address and contract address
	// of the requested pool id.
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty" yaml:"contract_address"`
	// code_id is the code id of the requested pool id.
	CodeId uint64 `protobuf:"varint,2,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty" yaml:"code_id"`
}

func (m *ContractInfoByPoolIdResponse) Reset()         { *m = ContractInfoByPoolIdResponse{} }
func (m *ContractInfoByPoolIdResponse) String() string { return proto.CompactTextString(m) }
func (*ContractInfoByPoolIdResponse) ProtoMessage()    {}
func (*ContractInfoByPoolIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_733c758985c393b2, []int{5}
}
func (m *ContractInfoByPoolIdResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractInfoByPoolIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractInfoByPoolIdResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractInfoByPoolIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractInfoByPoolIdResponse.Merge(m, src)
}
func (m *ContractInfoByPoolIdResponse) XXX_Size() int {
	return m.Size()
}
func (m *ContractInfoByPoolIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractInfoByPoolIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContractInfoByPoolIdResponse proto.InternalMessageInfo

func (m *ContractInfoByPoolIdResponse) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *ContractInfoByPoolIdResponse) GetCodeId() uint64 {
	if m != nil {
		return m.CodeId
	}
	return 0
}

func init() {
	proto.RegisterType((*ParamsRequest)(nil), "osmosis.cosmwasmpool.v1beta1.ParamsRequest")
	proto.RegisterType((*ParamsResponse)(nil), "osmosis.cosmwasmpool.v1beta1.ParamsResponse")
	proto.RegisterType((*PoolsRequest)(nil), "osmosis.cosmwasmpool.v1beta1.PoolsRequest")
	proto.RegisterType((*PoolsResponse)(nil), "osmosis.cosmwasmpool.v1beta1.PoolsResponse")
	proto.RegisterType((*ContractInfoByPoolIdRequest)(nil), "osmosis.cosmwasmpool.v1beta1.ContractInfoByPoolIdRequest")
	proto.RegisterType((*ContractInfoByPoolIdResponse)(nil), "osmosis.cosmwasmpool.v1beta1.ContractInfoByPoolIdResponse")
}

func init() {
	proto.RegisterFile("osmosis/cosmwasmpool/v1beta1/query.proto", fileDescriptor_733c758985c393b2)
}

var fileDescriptor_733c758985c393b2 = []byte{
	// 627 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xe3, 0x7e, 0x04, 0x75, 0x4b, 0x5b, 0xb4, 0xaa, 0x44, 0x49, 0x23, 0x07, 0x99, 0xb6,
	0x84, 0xa6, 0xb1, 0x95, 0x54, 0x08, 0xd1, 0x5b, 0x0d, 0x2a, 0x0a, 0xa7, 0x62, 0x21, 0x0e, 0x1c,
	0x88, 0x36, 0xf6, 0xd6, 0x58, 0xb2, 0xbd, 0x6e, 0x76, 0x53, 0x9a, 0x2b, 0x77, 0x04, 0xa8, 0xaf,
	0xc2, 0x13, 0x70, 0xaa, 0x38, 0x55, 0xe2, 0xc2, 0x29, 0x42, 0x09, 0x0f, 0x80, 0xfa, 0x04, 0x68,
	0x3f, 0x9c, 0xa6, 0x51, 0xe4, 0x46, 0x9c, 0x12, 0xef, 0xff, 0x3f, 0x33, 0xbf, 0x99, 0x1d, 0x1b,
	0x94, 0x09, 0x8d, 0x08, 0x0d, 0xa8, 0xe5, 0x12, 0x1a, 0x7d, 0x40, 0x34, 0x4a, 0x08, 0x09, 0xad,
	0x93, 0x5a, 0x0b, 0x33, 0x54, 0xb3, 0x8e, 0x3b, 0xb8, 0xdd, 0x35, 0x93, 0x36, 0x61, 0x04, 0x16,
	0x95, 0xd3, 0x1c, 0x75, 0x9a, 0xca, 0x59, 0x58, 0xf5, 0x89, 0x4f, 0x84, 0xd1, 0xe2, 0xff, 0x64,
	0x4c, 0xe1, 0x51, 0x66, 0xf6, 0x04, 0xb5, 0x51, 0x44, 0x95, 0x75, 0x33, 0xd3, 0xca, 0x4e, 0x95,
	0x4d, 0x77, 0x85, 0xcf, 0x6a, 0x21, 0x8a, 0x87, 0xaa, 0x4b, 0x82, 0x58, 0xe9, 0xdb, 0xa3, 0xba,
	0xc0, 0x1f, 0x29, 0xe7, 0x07, 0x31, 0x62, 0x01, 0x49, 0xbd, 0x45, 0x9f, 0x10, 0x3f, 0xc4, 0x16,
	0x4a, 0x02, 0x0b, 0xc5, 0x31, 0x61, 0x42, 0x4c, 0x81, 0xee, 0x29, 0x55, 0x3c, 0xb5, 0x3a, 0x47,
	0x16, 0x8a, 0xbb, 0xa9, 0x24, 0x8b, 0x34, 0x65, 0xbf, 0xf2, 0x41, 0x49, 0xa5, 0xf1, 0x28, 0x16,
	0x44, 0x98, 0x32, 0x14, 0x25, 0xd2, 0x60, 0xac, 0x80, 0xa5, 0x43, 0xd1, 0xb7, 0x83, 0x8f, 0x3b,
	0x98, 0x32, 0xe3, 0x35, 0x58, 0x4e, 0x0f, 0x68, 0x42, 0x62, 0x8a, 0xa1, 0x0d, 0xf2, 0x72, 0x34,
	0x6b, 0xda, 0x7d, 0xad, 0xbc, 0x58, 0xdf, 0x30, 0xb3, 0x46, 0x6f, 0xca, 0x68, 0x7b, 0xee, 0xbc,
	0x57, 0xca, 0x39, 0x2a, 0xd2, 0x78, 0x03, 0x6e, 0x1f, 0x12, 0x12, 0xa6, 0x55, 0xe0, 0x01, 0x00,
	0x57, 0xfd, 0xaf, 0xcd, 0x88, 0xbc, 0x5b, 0xa6, 0x42, 0xe7, 0xc3, 0x32, 0xe5, 0x5d, 0x5f, 0x25,
	0xf5, 0xb1, 0x8a, 0x75, 0x46, 0x22, 0x8d, 0xcf, 0x1a, 0x58, 0x52, 0x89, 0x15, 0xed, 0x63, 0x30,
	0xcf, 0x71, 0x38, 0xec, 0x6c, 0x79, 0xb1, 0xbe, 0x6a, 0xca, 0x09, 0x98, 0xe9, 0x04, 0xcc, 0xfd,
	0xb8, 0x6b, 0x2f, 0xfc, 0xf8, 0x56, 0x9d, 0xe7, 0x71, 0x0d, 0x47, 0xba, 0xe1, 0x8b, 0x09, 0x40,
	0x0f, 0x6f, 0x04, 0x92, 0x35, 0xaf, 0x11, 0xbd, 0x04, 0xeb, 0xcf, 0x48, 0xcc, 0xda, 0xc8, 0x65,
	0x8d, 0xf8, 0x88, 0xd8, 0x5d, 0x51, 0xc6, 0x4b, 0x1b, 0xaf, 0x80, 0x5b, 0xbc, 0x60, 0x33, 0xf0,
	0xc4, 0x34, 0xe7, 0x6c, 0x78, 0xd9, 0x2b, 0x2d, 0x77, 0x51, 0x14, 0xee, 0x19, 0x4a, 0x30, 0x9c,
	0x7c, 0x22, 0x62, 0x8c, 0x33, 0x0d, 0x14, 0x27, 0x27, 0x53, 0xcd, 0x1e, 0x80, 0x3b, 0xae, 0xd2,
	0x9b, 0xc8, 0xf3, 0xda, 0x98, 0xca, 0x4b, 0x5a, 0xb0, 0xd7, 0x2f, 0x7b, 0xa5, 0xbb, 0x32, 0xed,
	0xb8, 0xc3, 0x70, 0x56, 0xd2, 0xa3, 0x7d, 0x79, 0xc2, 0xa9, 0x5c, 0xe2, 0x61, 0x4e, 0x35, 0x33,
	0x4e, 0xa5, 0x04, 0xc3, 0xc9, 0xf3, 0x7f, 0x0d, 0xaf, 0xfe, 0x77, 0x16, 0xcc, 0xbf, 0xe2, 0xc3,
	0x80, 0x9f, 0x34, 0x20, 0xa6, 0x48, 0xe1, 0xf6, 0x0d, 0x3b, 0x31, 0x72, 0xf7, 0x85, 0xca, 0x54,
	0x5e, 0xd9, 0xa1, 0x51, 0xf9, 0xf8, 0xf3, 0xcf, 0xd9, 0xcc, 0x26, 0x7c, 0x60, 0x65, 0xbf, 0xbb,
	0x82, 0xe2, 0xab, 0x06, 0xf2, 0x72, 0xfd, 0x60, 0x65, 0x9a, 0x25, 0x4d, 0x89, 0x76, 0xa6, 0x33,
	0x2b, 0xa4, 0x1d, 0x81, 0xb4, 0x05, 0x37, 0xac, 0x29, 0x3e, 0x27, 0xf0, 0xbb, 0x06, 0x56, 0x27,
	0xdd, 0x21, 0x7c, 0x9a, 0x5d, 0x34, 0x63, 0x89, 0x0a, 0x7b, 0xff, 0x13, 0xaa, 0xe8, 0x77, 0x05,
	0x7d, 0x15, 0x56, 0xb2, 0xe9, 0x87, 0x4b, 0x13, 0xf0, 0x24, 0xef, 0xce, 0xfb, 0xba, 0x76, 0xd1,
	0xd7, 0xb5, 0xdf, 0x7d, 0x5d, 0xfb, 0x32, 0xd0, 0x73, 0x17, 0x03, 0x3d, 0xf7, 0x6b, 0xa0, 0xe7,
	0xde, 0x3e, 0xf7, 0x03, 0xf6, 0xbe, 0xd3, 0x32, 0x5d, 0x12, 0xa5, 0x09, 0xab, 0x21, 0x6a, 0xd1,
	0x61, 0xf6, 0x93, 0xda, 0x13, 0xeb, 0xf4, 0x7a, 0x0d, 0x37, 0x0c, 0x70, 0xcc, 0xe4, 0xe7, 0x50,
	0xbe, 0x96, 0x79, 0xf1, 0xb3, 0xfb, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x69, 0x97, 0xcf, 0x47, 0xff,
	0x05, 0x00, 0x00,
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
	// Pools returns all cosmwasm pools
	Pools(ctx context.Context, in *PoolsRequest, opts ...grpc.CallOption) (*PoolsResponse, error)
	// Params returns the parameters of the x/cosmwasmpool module.
	Params(ctx context.Context, in *ParamsRequest, opts ...grpc.CallOption) (*ParamsResponse, error)
	ContractInfoByPoolId(ctx context.Context, in *ContractInfoByPoolIdRequest, opts ...grpc.CallOption) (*ContractInfoByPoolIdResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Pools(ctx context.Context, in *PoolsRequest, opts ...grpc.CallOption) (*PoolsResponse, error) {
	out := new(PoolsResponse)
	err := c.cc.Invoke(ctx, "/osmosis.cosmwasmpool.v1beta1.Query/Pools", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *ParamsRequest, opts ...grpc.CallOption) (*ParamsResponse, error) {
	out := new(ParamsResponse)
	err := c.cc.Invoke(ctx, "/osmosis.cosmwasmpool.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ContractInfoByPoolId(ctx context.Context, in *ContractInfoByPoolIdRequest, opts ...grpc.CallOption) (*ContractInfoByPoolIdResponse, error) {
	out := new(ContractInfoByPoolIdResponse)
	err := c.cc.Invoke(ctx, "/osmosis.cosmwasmpool.v1beta1.Query/ContractInfoByPoolId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Pools returns all cosmwasm pools
	Pools(context.Context, *PoolsRequest) (*PoolsResponse, error)
	// Params returns the parameters of the x/cosmwasmpool module.
	Params(context.Context, *ParamsRequest) (*ParamsResponse, error)
	ContractInfoByPoolId(context.Context, *ContractInfoByPoolIdRequest) (*ContractInfoByPoolIdResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Pools(ctx context.Context, req *PoolsRequest) (*PoolsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pools not implemented")
}
func (*UnimplementedQueryServer) Params(ctx context.Context, req *ParamsRequest) (*ParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) ContractInfoByPoolId(ctx context.Context, req *ContractInfoByPoolIdRequest) (*ContractInfoByPoolIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContractInfoByPoolId not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Pools_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PoolsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Pools(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.cosmwasmpool.v1beta1.Query/Pools",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Pools(ctx, req.(*PoolsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.cosmwasmpool.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*ParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ContractInfoByPoolId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContractInfoByPoolIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ContractInfoByPoolId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.cosmwasmpool.v1beta1.Query/ContractInfoByPoolId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ContractInfoByPoolId(ctx, req.(*ContractInfoByPoolIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "osmosis.cosmwasmpool.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pools",
			Handler:    _Query_Pools_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "ContractInfoByPoolId",
			Handler:    _Query_ContractInfoByPoolId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "osmosis/cosmwasmpool/v1beta1/query.proto",
}

func (m *ParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *PoolsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *PoolsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Pools) > 0 {
		for iNdEx := len(m.Pools) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pools[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *ContractInfoByPoolIdRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractInfoByPoolIdRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractInfoByPoolIdRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PoolId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ContractInfoByPoolIdResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractInfoByPoolIdResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractInfoByPoolIdResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CodeId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.CodeId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ContractAddress)))
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
func (m *ParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *PoolsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *PoolsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Pools) > 0 {
		for _, e := range m.Pools {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *ContractInfoByPoolIdRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovQuery(uint64(m.PoolId))
	}
	return n
}

func (m *ContractInfoByPoolIdResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.CodeId != 0 {
		n += 1 + sovQuery(uint64(m.CodeId))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ParamsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *ParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *PoolsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: PoolsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *PoolsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: PoolsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pools", wireType)
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
			m.Pools = append(m.Pools, &types1.Any{})
			if err := m.Pools[len(m.Pools)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *ContractInfoByPoolIdRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ContractInfoByPoolIdRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractInfoByPoolIdRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
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
func (m *ContractInfoByPoolIdResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ContractInfoByPoolIdResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractInfoByPoolIdResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
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
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeId", wireType)
			}
			m.CodeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeId |= uint64(b&0x7F) << shift
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
