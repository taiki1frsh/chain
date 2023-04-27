// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pricefeed/pricefeed.proto

package types

import (
	fmt "fmt"
	github_com_UnUniFi_chain_types "github.com/UnUniFi/chain/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Market struct {
	MarketId   string                                            `protobuf:"bytes,1,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty" yaml:"market_id"`
	BaseAsset  string                                            `protobuf:"bytes,2,opt,name=base_asset,json=baseAsset,proto3" json:"base_asset,omitempty" yaml:"base_asset"`
	QuoteAsset string                                            `protobuf:"bytes,3,opt,name=quote_asset,json=quoteAsset,proto3" json:"quote_asset,omitempty" yaml:"quote_asset"`
	Oracles    []github_com_UnUniFi_chain_types.StringAccAddress `protobuf:"bytes,4,rep,name=oracles,proto3,customtype=github.com/UnUniFi/chain/types.StringAccAddress" json:"oracles" yaml:"oracles"`
	Active     bool                                              `protobuf:"varint,5,opt,name=active,proto3" json:"active,omitempty" yaml:"active"`
}

func (m *Market) Reset()         { *m = Market{} }
func (m *Market) String() string { return proto.CompactTextString(m) }
func (*Market) ProtoMessage()    {}
func (*Market) Descriptor() ([]byte, []int) {
	return fileDescriptor_c81c86fbf8838e8b, []int{0}
}
func (m *Market) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Market) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Market.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Market) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Market.Merge(m, src)
}
func (m *Market) XXX_Size() int {
	return m.Size()
}
func (m *Market) XXX_DiscardUnknown() {
	xxx_messageInfo_Market.DiscardUnknown(m)
}

var xxx_messageInfo_Market proto.InternalMessageInfo

func (m *Market) GetMarketId() string {
	if m != nil {
		return m.MarketId
	}
	return ""
}

func (m *Market) GetBaseAsset() string {
	if m != nil {
		return m.BaseAsset
	}
	return ""
}

func (m *Market) GetQuoteAsset() string {
	if m != nil {
		return m.QuoteAsset
	}
	return ""
}

func (m *Market) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

type CurrentPrice struct {
	MarketId string                                 `protobuf:"bytes,1,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty" yaml:"market_id"`
	Price    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price" yaml:"price"`
}

func (m *CurrentPrice) Reset()         { *m = CurrentPrice{} }
func (m *CurrentPrice) String() string { return proto.CompactTextString(m) }
func (*CurrentPrice) ProtoMessage()    {}
func (*CurrentPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_c81c86fbf8838e8b, []int{1}
}
func (m *CurrentPrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CurrentPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CurrentPrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CurrentPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrentPrice.Merge(m, src)
}
func (m *CurrentPrice) XXX_Size() int {
	return m.Size()
}
func (m *CurrentPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrentPrice.DiscardUnknown(m)
}

var xxx_messageInfo_CurrentPrice proto.InternalMessageInfo

func (m *CurrentPrice) GetMarketId() string {
	if m != nil {
		return m.MarketId
	}
	return ""
}

type PostedPrice struct {
	MarketId      string                                          `protobuf:"bytes,1,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty" yaml:"market_id"`
	OracleAddress github_com_UnUniFi_chain_types.StringAccAddress `protobuf:"bytes,2,opt,name=oracle_address,json=oracleAddress,proto3,customtype=github.com/UnUniFi/chain/types.StringAccAddress" json:"oracle_address" yaml:"oracle_address"`
	Price         github_com_cosmos_cosmos_sdk_types.Dec          `protobuf:"bytes,3,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price" yaml:"price"`
	Expiry        time.Time                                       `protobuf:"bytes,4,opt,name=expiry,proto3,stdtime" json:"expiry" yaml:"expiry"`
}

func (m *PostedPrice) Reset()         { *m = PostedPrice{} }
func (m *PostedPrice) String() string { return proto.CompactTextString(m) }
func (*PostedPrice) ProtoMessage()    {}
func (*PostedPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_c81c86fbf8838e8b, []int{2}
}
func (m *PostedPrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PostedPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PostedPrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PostedPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostedPrice.Merge(m, src)
}
func (m *PostedPrice) XXX_Size() int {
	return m.Size()
}
func (m *PostedPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_PostedPrice.DiscardUnknown(m)
}

var xxx_messageInfo_PostedPrice proto.InternalMessageInfo

func (m *PostedPrice) GetMarketId() string {
	if m != nil {
		return m.MarketId
	}
	return ""
}

func (m *PostedPrice) GetExpiry() time.Time {
	if m != nil {
		return m.Expiry
	}
	return time.Time{}
}

type Params struct {
	Markets []Market `protobuf:"bytes,1,rep,name=markets,proto3" json:"markets" yaml:"markets"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_c81c86fbf8838e8b, []int{3}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMarkets() []Market {
	if m != nil {
		return m.Markets
	}
	return nil
}

func init() {
	proto.RegisterType((*Market)(nil), "ununifi.pricefeed.Market")
	proto.RegisterType((*CurrentPrice)(nil), "ununifi.pricefeed.CurrentPrice")
	proto.RegisterType((*PostedPrice)(nil), "ununifi.pricefeed.PostedPrice")
	proto.RegisterType((*Params)(nil), "ununifi.pricefeed.Params")
}

func init() { proto.RegisterFile("pricefeed/pricefeed.proto", fileDescriptor_c81c86fbf8838e8b) }

var fileDescriptor_c81c86fbf8838e8b = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6a, 0xdb, 0x4c,
	0x14, 0xb5, 0xec, 0xc4, 0x89, 0xc7, 0x89, 0xf9, 0x2c, 0x92, 0x20, 0x7b, 0x21, 0x99, 0x59, 0x7c,
	0xb8, 0x85, 0x4a, 0x38, 0x2d, 0x14, 0xba, 0x28, 0x58, 0x0d, 0x85, 0x50, 0x02, 0x46, 0x8d, 0x37,
	0xdd, 0x98, 0xf1, 0x68, 0xac, 0x0c, 0xb1, 0x34, 0xae, 0x66, 0x94, 0xc6, 0x6f, 0x91, 0x5d, 0xdf,
	0xa5, 0x4f, 0x90, 0x65, 0x96, 0xa5, 0x0b, 0xb5, 0xd8, 0x6f, 0xa0, 0x27, 0x28, 0xd2, 0x8c, 0x7f,
	0xa0, 0x74, 0x63, 0xba, 0xf2, 0xbd, 0x73, 0xee, 0x39, 0x33, 0xf7, 0x5c, 0x5f, 0x81, 0xd6, 0x2c,
	0xa6, 0x98, 0x4c, 0x08, 0xf1, 0x9d, 0x75, 0x64, 0xcf, 0x62, 0x26, 0x98, 0xde, 0x4c, 0xa2, 0x24,
	0xa2, 0x13, 0x6a, 0xaf, 0x81, 0xf6, 0x49, 0xc0, 0x02, 0x56, 0xa0, 0x4e, 0x1e, 0xc9, 0xc2, 0xb6,
	0x15, 0x30, 0x16, 0x4c, 0x89, 0x53, 0x64, 0xe3, 0x64, 0xe2, 0x08, 0x1a, 0x12, 0x2e, 0x50, 0x38,
	0x53, 0x05, 0x26, 0x66, 0x3c, 0x64, 0xdc, 0x19, 0x23, 0x4e, 0x9c, 0xbb, 0xde, 0x98, 0x08, 0xd4,
	0x73, 0x30, 0xa3, 0x91, 0xc4, 0xe1, 0xb7, 0x32, 0xa8, 0x5e, 0xa1, 0xf8, 0x96, 0x08, 0xbd, 0x07,
	0x6a, 0x61, 0x11, 0x8d, 0xa8, 0x6f, 0x68, 0x1d, 0xad, 0x5b, 0x73, 0x4f, 0xb2, 0xd4, 0xfa, 0x6f,
	0x8e, 0xc2, 0xe9, 0x1b, 0xb8, 0x86, 0xa0, 0x77, 0x28, 0xe3, 0x4b, 0x5f, 0x7f, 0x05, 0x40, 0x2e,
	0x3c, 0x42, 0x9c, 0x13, 0x61, 0x94, 0x0b, 0xce, 0x69, 0x96, 0x5a, 0x4d, 0xc9, 0xd9, 0x60, 0xd0,
	0xab, 0xe5, 0x49, 0x3f, 0x8f, 0xf5, 0xd7, 0xa0, 0xfe, 0x39, 0x61, 0x62, 0x45, 0xab, 0x14, 0xb4,
	0xb3, 0x2c, 0xb5, 0x74, 0x49, 0xdb, 0x02, 0xa1, 0x07, 0x8a, 0x4c, 0x12, 0x31, 0x38, 0x60, 0x31,
	0xc2, 0x53, 0xc2, 0x8d, 0xbd, 0x4e, 0xa5, 0x5b, 0x73, 0x2f, 0x1f, 0x53, 0xab, 0xf4, 0x23, 0xb5,
	0x9c, 0x80, 0x8a, 0x9b, 0x64, 0x6c, 0x63, 0x16, 0x3a, 0xc3, 0x68, 0x18, 0xd1, 0xf7, 0xd4, 0xc1,
	0x37, 0x88, 0x46, 0x8e, 0x98, 0xcf, 0x08, 0xb7, 0x3f, 0x8a, 0x98, 0x46, 0x41, 0x1f, 0xe3, 0xbe,
	0xef, 0xc7, 0x84, 0xf3, 0x2c, 0xb5, 0x1a, 0xf2, 0x2e, 0xa5, 0x07, 0xbd, 0x95, 0xb2, 0xfe, 0x0c,
	0x54, 0x11, 0x16, 0xf4, 0x8e, 0x18, 0xfb, 0x1d, 0xad, 0x7b, 0xe8, 0x36, 0xb3, 0xd4, 0x3a, 0x96,
	0xc5, 0xf2, 0x1c, 0x7a, 0xaa, 0x00, 0x7e, 0xd5, 0xc0, 0xd1, 0xbb, 0x24, 0x8e, 0x49, 0x24, 0x06,
	0xf9, 0xa0, 0x76, 0xb1, 0xf0, 0x1a, 0xec, 0x17, 0x43, 0x56, 0xee, 0xbd, 0x55, 0x1d, 0xfd, 0xbf,
	0xd5, 0x91, 0x1a, 0xa1, 0xfc, 0x79, 0xc1, 0xfd, 0x5b, 0xd5, 0xd5, 0x05, 0xc1, 0x59, 0x6a, 0x1d,
	0x49, 0xf1, 0x42, 0x04, 0x7a, 0x52, 0x0c, 0xa6, 0x65, 0x50, 0x1f, 0x30, 0x2e, 0x88, 0xbf, 0xf3,
	0xc3, 0xbe, 0x80, 0x86, 0xb4, 0x64, 0x84, 0xa4, 0x69, 0xea, 0x85, 0x83, 0xdd, 0x3d, 0x3f, 0xdd,
	0xf6, 0x7c, 0x25, 0x0b, 0xbd, 0x63, 0x79, 0xa0, 0xea, 0x36, 0x8e, 0x54, 0xfe, 0xa1, 0x23, 0xfa,
	0x15, 0xa8, 0x92, 0xfb, 0x19, 0x8d, 0xe7, 0xc6, 0x5e, 0x47, 0xeb, 0xd6, 0xcf, 0xdb, 0xb6, 0x5c,
	0x1d, 0x7b, 0xb5, 0x3a, 0xf6, 0xf5, 0x6a, 0x75, 0xdc, 0x56, 0x7e, 0xe5, 0x66, 0xec, 0x92, 0x07,
	0x1f, 0x7e, 0x5a, 0x9a, 0xa7, 0x44, 0xe0, 0x10, 0x54, 0x07, 0x28, 0x46, 0x21, 0xd7, 0x3f, 0x80,
	0x03, 0xe9, 0x19, 0x37, 0xb4, 0x4e, 0xa5, 0x5b, 0x3f, 0x6f, 0xd9, 0x7f, 0x6c, 0xaf, 0x2d, 0x57,
	0xcc, 0x3d, 0x53, 0xc2, 0x8d, 0x6d, 0xdf, 0xf3, 0x3f, 0x9f, 0x8a, 0xdc, 0x8b, 0xc7, 0x85, 0xa9,
	0x3d, 0x2d, 0x4c, 0xed, 0xd7, 0xc2, 0xd4, 0x1e, 0x96, 0x66, 0xe9, 0x69, 0x69, 0x96, 0xbe, 0x2f,
	0xcd, 0xd2, 0xa7, 0xe7, 0x7f, 0xb5, 0xfb, 0x7e, 0xf3, 0xf9, 0x90, 0x36, 0x8c, 0xab, 0x45, 0x4f,
	0x2f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xb7, 0x0d, 0x5e, 0x62, 0x04, 0x00, 0x00,
}

func (m *Market) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Market) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Market) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Active {
		i--
		if m.Active {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.Oracles) > 0 {
		for iNdEx := len(m.Oracles) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Oracles[iNdEx].Size()
				i -= size
				if _, err := m.Oracles[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintPricefeed(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.QuoteAsset) > 0 {
		i -= len(m.QuoteAsset)
		copy(dAtA[i:], m.QuoteAsset)
		i = encodeVarintPricefeed(dAtA, i, uint64(len(m.QuoteAsset)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BaseAsset) > 0 {
		i -= len(m.BaseAsset)
		copy(dAtA[i:], m.BaseAsset)
		i = encodeVarintPricefeed(dAtA, i, uint64(len(m.BaseAsset)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MarketId) > 0 {
		i -= len(m.MarketId)
		copy(dAtA[i:], m.MarketId)
		i = encodeVarintPricefeed(dAtA, i, uint64(len(m.MarketId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CurrentPrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CurrentPrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CurrentPrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPricefeed(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.MarketId) > 0 {
		i -= len(m.MarketId)
		copy(dAtA[i:], m.MarketId)
		i = encodeVarintPricefeed(dAtA, i, uint64(len(m.MarketId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PostedPrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PostedPrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PostedPrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Expiry, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Expiry):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintPricefeed(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPricefeed(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.OracleAddress.Size()
		i -= size
		if _, err := m.OracleAddress.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPricefeed(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.MarketId) > 0 {
		i -= len(m.MarketId)
		copy(dAtA[i:], m.MarketId)
		i = encodeVarintPricefeed(dAtA, i, uint64(len(m.MarketId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Markets) > 0 {
		for iNdEx := len(m.Markets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Markets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPricefeed(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintPricefeed(dAtA []byte, offset int, v uint64) int {
	offset -= sovPricefeed(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Market) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MarketId)
	if l > 0 {
		n += 1 + l + sovPricefeed(uint64(l))
	}
	l = len(m.BaseAsset)
	if l > 0 {
		n += 1 + l + sovPricefeed(uint64(l))
	}
	l = len(m.QuoteAsset)
	if l > 0 {
		n += 1 + l + sovPricefeed(uint64(l))
	}
	if len(m.Oracles) > 0 {
		for _, e := range m.Oracles {
			l = e.Size()
			n += 1 + l + sovPricefeed(uint64(l))
		}
	}
	if m.Active {
		n += 2
	}
	return n
}

func (m *CurrentPrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MarketId)
	if l > 0 {
		n += 1 + l + sovPricefeed(uint64(l))
	}
	l = m.Price.Size()
	n += 1 + l + sovPricefeed(uint64(l))
	return n
}

func (m *PostedPrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MarketId)
	if l > 0 {
		n += 1 + l + sovPricefeed(uint64(l))
	}
	l = m.OracleAddress.Size()
	n += 1 + l + sovPricefeed(uint64(l))
	l = m.Price.Size()
	n += 1 + l + sovPricefeed(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Expiry)
	n += 1 + l + sovPricefeed(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Markets) > 0 {
		for _, e := range m.Markets {
			l = e.Size()
			n += 1 + l + sovPricefeed(uint64(l))
		}
	}
	return n
}

func sovPricefeed(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPricefeed(x uint64) (n int) {
	return sovPricefeed(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Market) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPricefeed
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
			return fmt.Errorf("proto: Market: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Market: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPricefeed
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
				return ErrInvalidLengthPricefeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPricefeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarketId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPricefeed
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
				return ErrInvalidLengthPricefeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPricefeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseAsset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuoteAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPricefeed
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
				return ErrInvalidLengthPricefeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPricefeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QuoteAsset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Oracles", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPricefeed
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
				return ErrInvalidLengthPricefeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPricefeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_UnUniFi_chain_types.StringAccAddress
			m.Oracles = append(m.Oracles, v)
			if err := m.Oracles[len(m.Oracles)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Active", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPricefeed
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Active = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPricefeed(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPricefeed
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
func (m *CurrentPrice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPricefeed
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
			return fmt.Errorf("proto: CurrentPrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CurrentPrice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPricefeed
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
				return ErrInvalidLengthPricefeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPricefeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarketId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPricefeed
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
				return ErrInvalidLengthPricefeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPricefeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPricefeed(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPricefeed
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
func (m *PostedPrice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPricefeed
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
			return fmt.Errorf("proto: PostedPrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PostedPrice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPricefeed
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
				return ErrInvalidLengthPricefeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPricefeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarketId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPricefeed
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
				return ErrInvalidLengthPricefeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPricefeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OracleAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPricefeed
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
				return ErrInvalidLengthPricefeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPricefeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPricefeed
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
				return ErrInvalidLengthPricefeed
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPricefeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Expiry, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPricefeed(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPricefeed
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPricefeed
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Markets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPricefeed
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
				return ErrInvalidLengthPricefeed
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPricefeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Markets = append(m.Markets, Market{})
			if err := m.Markets[len(m.Markets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPricefeed(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPricefeed
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
func skipPricefeed(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPricefeed
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
					return 0, ErrIntOverflowPricefeed
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
					return 0, ErrIntOverflowPricefeed
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
				return 0, ErrInvalidLengthPricefeed
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPricefeed
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPricefeed
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPricefeed        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPricefeed          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPricefeed = fmt.Errorf("proto: unexpected end of group")
)
