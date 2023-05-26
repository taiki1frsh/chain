// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ecosystem-incentive/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_UnUniFi_chain_types "github.com/UnUniFi/chain/types"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

type MsgRegister struct {
	Sender          github_com_UnUniFi_chain_types.StringAccAddress   `protobuf:"bytes,1,opt,name=sender,proto3,customtype=github.com/UnUniFi/chain/types.StringAccAddress" json:"sender" yaml:"sender"`
	IncentiveUnitId string                                            `protobuf:"bytes,2,opt,name=incentive_unit_id,json=incentiveUnitId,proto3" json:"incentive_unit_id,omitempty" yaml:"incentive_unit_id"`
	SubjectAddrs    []github_com_UnUniFi_chain_types.StringAccAddress `protobuf:"bytes,3,rep,name=subject_addrs,json=subjectAddrs,proto3,customtype=github.com/UnUniFi/chain/types.StringAccAddress" json:"subject_addrs" yaml:"subject_addrs"`
	Weights         []github_com_cosmos_cosmos_sdk_types.Dec          `protobuf:"bytes,4,rep,name=weights,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"weights" yaml:"weights"`
}

func (m *MsgRegister) Reset()         { *m = MsgRegister{} }
func (m *MsgRegister) String() string { return proto.CompactTextString(m) }
func (*MsgRegister) ProtoMessage()    {}
func (*MsgRegister) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f4f2600114c5509, []int{0}
}
func (m *MsgRegister) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegister.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegister.Merge(m, src)
}
func (m *MsgRegister) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegister) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegister.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegister proto.InternalMessageInfo

func (m *MsgRegister) GetIncentiveUnitId() string {
	if m != nil {
		return m.IncentiveUnitId
	}
	return ""
}

type MsgRegisterResponse struct {
}

func (m *MsgRegisterResponse) Reset()         { *m = MsgRegisterResponse{} }
func (m *MsgRegisterResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterResponse) ProtoMessage()    {}
func (*MsgRegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f4f2600114c5509, []int{1}
}
func (m *MsgRegisterResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterResponse.Merge(m, src)
}
func (m *MsgRegisterResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterResponse proto.InternalMessageInfo

type MsgWithdrawAllRewards struct {
	Sender github_com_UnUniFi_chain_types.StringAccAddress `protobuf:"bytes,1,opt,name=sender,proto3,customtype=github.com/UnUniFi/chain/types.StringAccAddress" json:"sender" yaml:"sender"`
}

func (m *MsgWithdrawAllRewards) Reset()         { *m = MsgWithdrawAllRewards{} }
func (m *MsgWithdrawAllRewards) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawAllRewards) ProtoMessage()    {}
func (*MsgWithdrawAllRewards) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f4f2600114c5509, []int{2}
}
func (m *MsgWithdrawAllRewards) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawAllRewards) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawAllRewards.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawAllRewards) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawAllRewards.Merge(m, src)
}
func (m *MsgWithdrawAllRewards) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawAllRewards) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawAllRewards.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawAllRewards proto.InternalMessageInfo

type MsgWithdrawAllRewardsResponse struct {
}

func (m *MsgWithdrawAllRewardsResponse) Reset()         { *m = MsgWithdrawAllRewardsResponse{} }
func (m *MsgWithdrawAllRewardsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawAllRewardsResponse) ProtoMessage()    {}
func (*MsgWithdrawAllRewardsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f4f2600114c5509, []int{3}
}
func (m *MsgWithdrawAllRewardsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawAllRewardsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawAllRewardsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawAllRewardsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawAllRewardsResponse.Merge(m, src)
}
func (m *MsgWithdrawAllRewardsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawAllRewardsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawAllRewardsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawAllRewardsResponse proto.InternalMessageInfo

type MsgWithdrawReward struct {
	Sender github_com_UnUniFi_chain_types.StringAccAddress `protobuf:"bytes,1,opt,name=sender,proto3,customtype=github.com/UnUniFi/chain/types.StringAccAddress" json:"sender" yaml:"sender"`
	Denom  string                                          `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty" yaml:"denom"`
}

func (m *MsgWithdrawReward) Reset()         { *m = MsgWithdrawReward{} }
func (m *MsgWithdrawReward) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawReward) ProtoMessage()    {}
func (*MsgWithdrawReward) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f4f2600114c5509, []int{4}
}
func (m *MsgWithdrawReward) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawReward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawReward.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawReward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawReward.Merge(m, src)
}
func (m *MsgWithdrawReward) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawReward) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawReward.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawReward proto.InternalMessageInfo

func (m *MsgWithdrawReward) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

type MsgWithdrawRewardResponse struct {
}

func (m *MsgWithdrawRewardResponse) Reset()         { *m = MsgWithdrawRewardResponse{} }
func (m *MsgWithdrawRewardResponse) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawRewardResponse) ProtoMessage()    {}
func (*MsgWithdrawRewardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f4f2600114c5509, []int{5}
}
func (m *MsgWithdrawRewardResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawRewardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawRewardResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawRewardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawRewardResponse.Merge(m, src)
}
func (m *MsgWithdrawRewardResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawRewardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawRewardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawRewardResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRegister)(nil), "ununifi.ecosystemincentive.MsgRegister")
	proto.RegisterType((*MsgRegisterResponse)(nil), "ununifi.ecosystemincentive.MsgRegisterResponse")
	proto.RegisterType((*MsgWithdrawAllRewards)(nil), "ununifi.ecosystemincentive.MsgWithdrawAllRewards")
	proto.RegisterType((*MsgWithdrawAllRewardsResponse)(nil), "ununifi.ecosystemincentive.MsgWithdrawAllRewardsResponse")
	proto.RegisterType((*MsgWithdrawReward)(nil), "ununifi.ecosystemincentive.MsgWithdrawReward")
	proto.RegisterType((*MsgWithdrawRewardResponse)(nil), "ununifi.ecosystemincentive.MsgWithdrawRewardResponse")
}

func init() { proto.RegisterFile("ecosystem-incentive/tx.proto", fileDescriptor_5f4f2600114c5509) }

var fileDescriptor_5f4f2600114c5509 = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xbf, 0x6f, 0xd3, 0x40,
	0x18, 0x8d, 0x1b, 0x28, 0xf4, 0x68, 0x0b, 0x35, 0xad, 0xe4, 0x9a, 0x62, 0x57, 0x37, 0x94, 0x2e,
	0xb1, 0xc5, 0xaf, 0x01, 0x26, 0x12, 0x21, 0x54, 0x86, 0x80, 0x64, 0x14, 0x21, 0x75, 0x09, 0x8e,
	0xef, 0x70, 0x0e, 0xe2, 0xbb, 0xc8, 0xdf, 0xb9, 0x69, 0x56, 0x56, 0x16, 0xfe, 0x01, 0xfe, 0x9f,
	0x8e, 0x1d, 0x11, 0x83, 0x85, 0x92, 0x85, 0x39, 0x7f, 0x01, 0x8a, 0xcf, 0xb1, 0x2c, 0x92, 0x8a,
	0x06, 0x09, 0x26, 0xfb, 0xee, 0x7b, 0xdf, 0x7b, 0x4f, 0xf7, 0xbe, 0x3b, 0xb4, 0x47, 0x03, 0x01,
	0x43, 0x90, 0x34, 0xaa, 0x31, 0x1e, 0x50, 0x2e, 0xd9, 0x09, 0x75, 0xe5, 0xa9, 0xd3, 0x8f, 0x85,
	0x14, 0xba, 0x99, 0xf0, 0x84, 0xb3, 0xf7, 0xcc, 0x29, 0x50, 0x05, 0xc8, 0xdc, 0x0e, 0x45, 0x28,
	0x32, 0x98, 0x3b, 0xfd, 0x53, 0x1d, 0xe6, 0x6e, 0x20, 0x20, 0x12, 0xd0, 0x56, 0x05, 0xb5, 0x50,
	0x25, 0xfc, 0xb9, 0x8a, 0x6e, 0x34, 0x21, 0xf4, 0x68, 0xc8, 0x40, 0xd2, 0x58, 0x7f, 0x87, 0x56,
	0x81, 0x72, 0x42, 0x63, 0x43, 0xdb, 0xd7, 0x0e, 0xd7, 0x1a, 0x47, 0x67, 0xa9, 0x5d, 0xf9, 0x9e,
	0xda, 0x6e, 0xc8, 0x64, 0x37, 0xe9, 0x38, 0x81, 0x88, 0xdc, 0x16, 0x6f, 0x71, 0xf6, 0x82, 0xb9,
	0x41, 0xd7, 0x67, 0xdc, 0x95, 0xc3, 0x3e, 0x05, 0xe7, 0x8d, 0x8c, 0x19, 0x0f, 0xeb, 0x41, 0x50,
	0x27, 0x24, 0xa6, 0x00, 0x93, 0xd4, 0xde, 0x18, 0xfa, 0x51, 0xef, 0x29, 0x56, 0x74, 0xd8, 0xcb,
	0x79, 0xf5, 0x23, 0xb4, 0x55, 0xf8, 0x6d, 0x27, 0x9c, 0xc9, 0x36, 0x23, 0xc6, 0x4a, 0x26, 0xb6,
	0x37, 0x49, 0x6d, 0x43, 0x75, 0xcd, 0x41, 0xb0, 0x77, 0xb3, 0xd8, 0x6b, 0x71, 0x26, 0x5f, 0x12,
	0x5d, 0xa2, 0x0d, 0x48, 0x3a, 0x1f, 0x68, 0x20, 0xdb, 0x3e, 0x21, 0x31, 0x18, 0xd5, 0xfd, 0xea,
	0xe1, 0x5a, 0xe3, 0xf5, 0xdf, 0x5b, 0xde, 0xce, 0x2d, 0x97, 0x59, 0xb1, 0xb7, 0x9e, 0xaf, 0xa7,
	0x38, 0xd0, 0x8f, 0xd1, 0xb5, 0x01, 0x65, 0x61, 0x57, 0x82, 0x71, 0x25, 0xd3, 0x7b, 0x96, 0xeb,
	0x1d, 0x94, 0xf4, 0xd4, 0x19, 0xe7, 0x9f, 0x1a, 0x90, 0x8f, 0xb9, 0xe6, 0x73, 0x1a, 0x4c, 0x52,
	0x7b, 0x53, 0xc9, 0xe4, 0x34, 0xd8, 0x9b, 0x11, 0xe2, 0x1d, 0x74, 0xbb, 0x14, 0x86, 0x47, 0xa1,
	0x2f, 0x38, 0x50, 0x3c, 0x44, 0x3b, 0x4d, 0x08, 0xdf, 0x32, 0xd9, 0x25, 0xb1, 0x3f, 0xa8, 0xf7,
	0x7a, 0x1e, 0x1d, 0xf8, 0x31, 0x81, 0x7f, 0x9f, 0x16, 0xb6, 0xd1, 0xdd, 0x85, 0xd2, 0x85, 0xb7,
	0xaf, 0x1a, 0xda, 0x2a, 0x21, 0x54, 0xf9, 0x3f, 0x8c, 0xd1, 0x01, 0xba, 0x4a, 0x28, 0x17, 0x51,
	0x3e, 0x3a, 0xb7, 0x26, 0xa9, 0xbd, 0xae, 0x90, 0xd9, 0x36, 0xf6, 0x54, 0x19, 0xdf, 0x41, 0xbb,
	0x73, 0xf6, 0x66, 0xe6, 0x1f, 0xfc, 0x5c, 0x41, 0xd5, 0x26, 0x84, 0x3a, 0x41, 0xd7, 0x8b, 0x1b,
	0x70, 0xcf, 0xb9, 0xf8, 0x7e, 0x39, 0xa5, 0x74, 0x4c, 0xf7, 0x92, 0xc0, 0x99, 0x9a, 0xfe, 0x49,
	0x43, 0xfa, 0x82, 0x10, 0xef, 0xff, 0x81, 0x67, 0xbe, 0xc5, 0x7c, 0xb2, 0x74, 0x4b, 0x61, 0xe2,
	0x04, 0x6d, 0xfe, 0x96, 0x55, 0xed, 0x92, 0x64, 0x0a, 0x6e, 0x3e, 0x5e, 0x0a, 0x3e, 0xd3, 0x6d,
	0xbc, 0x3a, 0x1b, 0x59, 0xda, 0xf9, 0xc8, 0xd2, 0x7e, 0x8c, 0x2c, 0xed, 0xcb, 0xd8, 0xaa, 0x9c,
	0x8f, 0xad, 0xca, 0xb7, 0xb1, 0x55, 0x39, 0x7e, 0x74, 0xe1, 0x4c, 0x9c, 0xba, 0x0b, 0x1f, 0xc2,
	0xe9, 0xa4, 0x74, 0x56, 0xb3, 0xf7, 0xeb, 0xe1, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x89,
	0x7d, 0x5c, 0x2c, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	Register(ctx context.Context, in *MsgRegister, opts ...grpc.CallOption) (*MsgRegisterResponse, error)
	WithdrawAllRewards(ctx context.Context, in *MsgWithdrawAllRewards, opts ...grpc.CallOption) (*MsgWithdrawAllRewardsResponse, error)
	WithdrawReward(ctx context.Context, in *MsgWithdrawReward, opts ...grpc.CallOption) (*MsgWithdrawRewardResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Register(ctx context.Context, in *MsgRegister, opts ...grpc.CallOption) (*MsgRegisterResponse, error) {
	out := new(MsgRegisterResponse)
	err := c.cc.Invoke(ctx, "/ununifi.ecosystemincentive.Msg/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) WithdrawAllRewards(ctx context.Context, in *MsgWithdrawAllRewards, opts ...grpc.CallOption) (*MsgWithdrawAllRewardsResponse, error) {
	out := new(MsgWithdrawAllRewardsResponse)
	err := c.cc.Invoke(ctx, "/ununifi.ecosystemincentive.Msg/WithdrawAllRewards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) WithdrawReward(ctx context.Context, in *MsgWithdrawReward, opts ...grpc.CallOption) (*MsgWithdrawRewardResponse, error) {
	out := new(MsgWithdrawRewardResponse)
	err := c.cc.Invoke(ctx, "/ununifi.ecosystemincentive.Msg/WithdrawReward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Register(context.Context, *MsgRegister) (*MsgRegisterResponse, error)
	WithdrawAllRewards(context.Context, *MsgWithdrawAllRewards) (*MsgWithdrawAllRewardsResponse, error)
	WithdrawReward(context.Context, *MsgWithdrawReward) (*MsgWithdrawRewardResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Register(ctx context.Context, req *MsgRegister) (*MsgRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedMsgServer) WithdrawAllRewards(ctx context.Context, req *MsgWithdrawAllRewards) (*MsgWithdrawAllRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawAllRewards not implemented")
}
func (*UnimplementedMsgServer) WithdrawReward(ctx context.Context, req *MsgWithdrawReward) (*MsgWithdrawRewardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawReward not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegister)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ununifi.ecosystemincentive.Msg/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Register(ctx, req.(*MsgRegister))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_WithdrawAllRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdrawAllRewards)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).WithdrawAllRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ununifi.ecosystemincentive.Msg/WithdrawAllRewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).WithdrawAllRewards(ctx, req.(*MsgWithdrawAllRewards))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_WithdrawReward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdrawReward)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).WithdrawReward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ununifi.ecosystemincentive.Msg/WithdrawReward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).WithdrawReward(ctx, req.(*MsgWithdrawReward))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ununifi.ecosystemincentive.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Msg_Register_Handler,
		},
		{
			MethodName: "WithdrawAllRewards",
			Handler:    _Msg_WithdrawAllRewards_Handler,
		},
		{
			MethodName: "WithdrawReward",
			Handler:    _Msg_WithdrawReward_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ecosystem-incentive/tx.proto",
}

func (m *MsgRegister) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegister) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegister) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Weights) > 0 {
		for iNdEx := len(m.Weights) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Weights[iNdEx].Size()
				i -= size
				if _, err := m.Weights[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.SubjectAddrs) > 0 {
		for iNdEx := len(m.SubjectAddrs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.SubjectAddrs[iNdEx].Size()
				i -= size
				if _, err := m.SubjectAddrs[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.IncentiveUnitId) > 0 {
		i -= len(m.IncentiveUnitId)
		copy(dAtA[i:], m.IncentiveUnitId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.IncentiveUnitId)))
		i--
		dAtA[i] = 0x12
	}
	{
		size := m.Sender.Size()
		i -= size
		if _, err := m.Sender.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgRegisterResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawAllRewards) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawAllRewards) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawAllRewards) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Sender.Size()
		i -= size
		if _, err := m.Sender.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawAllRewardsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawAllRewardsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawAllRewardsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawReward) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawReward) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawReward) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	{
		size := m.Sender.Size()
		i -= size
		if _, err := m.Sender.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawRewardResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawRewardResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawRewardResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRegister) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Sender.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.IncentiveUnitId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.SubjectAddrs) > 0 {
		for _, e := range m.SubjectAddrs {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.Weights) > 0 {
		for _, e := range m.Weights {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgRegisterResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgWithdrawAllRewards) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Sender.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgWithdrawAllRewardsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgWithdrawReward) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Sender.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgWithdrawRewardResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRegister) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRegister: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegister: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Sender.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncentiveUnitId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IncentiveUnitId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubjectAddrs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_UnUniFi_chain_types.StringAccAddress
			m.SubjectAddrs = append(m.SubjectAddrs, v)
			if err := m.SubjectAddrs[len(m.SubjectAddrs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weights", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.Weights = append(m.Weights, v)
			if err := m.Weights[len(m.Weights)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgRegisterResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRegisterResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgWithdrawAllRewards) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgWithdrawAllRewards: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawAllRewards: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Sender.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgWithdrawAllRewardsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgWithdrawAllRewardsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawAllRewardsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgWithdrawReward) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgWithdrawReward: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawReward: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Sender.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgWithdrawRewardResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgWithdrawRewardResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawRewardResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
