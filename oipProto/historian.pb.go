// Code generated by protoc-gen-go. DO NOT EDIT.
// source: historian.proto

package oipProto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HistorianDataPoint struct {
	Version                  int32   `protobuf:"varint,1,opt,name=Version" json:"Version,omitempty"`
	PubKey                   []byte  `protobuf:"bytes,2,opt,name=PubKey,proto3" json:"PubKey,omitempty"`
	MiningRigRentalsLast10   float64 `protobuf:"fixed64,3,opt,name=MiningRigRentalsLast10" json:"MiningRigRentalsLast10,omitempty"`
	MiningRigRentalsLast24Hr float64 `protobuf:"fixed64,4,opt,name=MiningRigRentalsLast24Hr" json:"MiningRigRentalsLast24Hr,omitempty"`
	AutominerPoolHashrate    float64 `protobuf:"fixed64,5,opt,name=AutominerPoolHashrate" json:"AutominerPoolHashrate,omitempty"`
	FloNetHashRate           float64 `protobuf:"fixed64,6,opt,name=FloNetHashRate" json:"FloNetHashRate,omitempty"`
	FloMarketPriceBTC        float64 `protobuf:"fixed64,7,opt,name=FloMarketPriceBTC" json:"FloMarketPriceBTC,omitempty"`
	FloMarketPriceUSD        float64 `protobuf:"fixed64,8,opt,name=FloMarketPriceUSD" json:"FloMarketPriceUSD,omitempty"`
	LtcMarketPriceUSD        float64 `protobuf:"fixed64,9,opt,name=LtcMarketPriceUSD" json:"LtcMarketPriceUSD,omitempty"`
	NiceHashLast             float64 `protobuf:"fixed64,10,opt,name=NiceHashLast" json:"NiceHashLast,omitempty"`
	NiceHash24Hr             float64 `protobuf:"fixed64,11,opt,name=NiceHash24hr" json:"NiceHash24hr,omitempty"`
}

func (m *HistorianDataPoint) Reset()                    { *m = HistorianDataPoint{} }
func (m *HistorianDataPoint) String() string            { return proto.CompactTextString(m) }
func (*HistorianDataPoint) ProtoMessage()               {}
func (*HistorianDataPoint) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *HistorianDataPoint) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *HistorianDataPoint) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *HistorianDataPoint) GetMiningRigRentalsLast10() float64 {
	if m != nil {
		return m.MiningRigRentalsLast10
	}
	return 0
}

func (m *HistorianDataPoint) GetMiningRigRentalsLast24Hr() float64 {
	if m != nil {
		return m.MiningRigRentalsLast24Hr
	}
	return 0
}

func (m *HistorianDataPoint) GetAutominerPoolHashrate() float64 {
	if m != nil {
		return m.AutominerPoolHashrate
	}
	return 0
}

func (m *HistorianDataPoint) GetFloNetHashRate() float64 {
	if m != nil {
		return m.FloNetHashRate
	}
	return 0
}

func (m *HistorianDataPoint) GetFloMarketPriceBTC() float64 {
	if m != nil {
		return m.FloMarketPriceBTC
	}
	return 0
}

func (m *HistorianDataPoint) GetFloMarketPriceUSD() float64 {
	if m != nil {
		return m.FloMarketPriceUSD
	}
	return 0
}

func (m *HistorianDataPoint) GetLtcMarketPriceUSD() float64 {
	if m != nil {
		return m.LtcMarketPriceUSD
	}
	return 0
}

func (m *HistorianDataPoint) GetNiceHashLast() float64 {
	if m != nil {
		return m.NiceHashLast
	}
	return 0
}

func (m *HistorianDataPoint) GetNiceHash24Hr() float64 {
	if m != nil {
		return m.NiceHash24Hr
	}
	return 0
}

type HistorianPayout struct {
	Version int32 `protobuf:"varint,1,opt,name=Version" json:"Version,omitempty"`
}

func (m *HistorianPayout) Reset()                    { *m = HistorianPayout{} }
func (m *HistorianPayout) String() string            { return proto.CompactTextString(m) }
func (*HistorianPayout) ProtoMessage()               {}
func (*HistorianPayout) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *HistorianPayout) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*HistorianDataPoint)(nil), "oipProto.HistorianDataPoint")
	proto.RegisterType((*HistorianPayout)(nil), "oipProto.HistorianPayout")
}

func init() { proto.RegisterFile("historian.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0x87, 0x59, 0x6b, 0xff, 0x38, 0x16, 0x8b, 0x01, 0x4b, 0x8e, 0xa5, 0x07, 0x29, 0x28, 0xa2,
	0xb5, 0x78, 0xf0, 0x66, 0x5d, 0xca, 0x82, 0x6d, 0x09, 0xa9, 0xf5, 0xe0, 0x2d, 0x5d, 0x42, 0x37,
	0xb8, 0x66, 0x4a, 0x36, 0x7b, 0xe8, 0x2b, 0xfb, 0x14, 0x92, 0x68, 0x8b, 0xed, 0x6e, 0xbd, 0x65,
	0x7e, 0xdf, 0x37, 0x90, 0x61, 0x06, 0x5a, 0x89, 0xca, 0x2c, 0x1a, 0x25, 0xf4, 0xcd, 0xca, 0xa0,
	0x45, 0xd2, 0x40, 0xb5, 0x62, 0xee, 0xd5, 0xfd, 0xaa, 0x00, 0x89, 0x36, 0x34, 0x14, 0x56, 0x30,
	0x54, 0xda, 0x12, 0x0a, 0xf5, 0x37, 0x69, 0x32, 0x85, 0x9a, 0x06, 0x9d, 0xa0, 0x57, 0xe5, 0x9b,
	0x92, 0xb4, 0xa1, 0xc6, 0xf2, 0xc5, 0x8b, 0x5c, 0xd3, 0xa3, 0x4e, 0xd0, 0x6b, 0xf2, 0xdf, 0x8a,
	0x3c, 0x40, 0x7b, 0xa2, 0xb4, 0xd2, 0x4b, 0xae, 0x96, 0x5c, 0x6a, 0x2b, 0xd2, 0x6c, 0x2c, 0x32,
	0x7b, 0x77, 0x4b, 0x2b, 0x9d, 0xa0, 0x17, 0xf0, 0x03, 0x94, 0x3c, 0x02, 0x2d, 0x23, 0xfd, 0x41,
	0x64, 0xe8, 0xb1, 0xef, 0x3c, 0xc8, 0xc9, 0x00, 0x2e, 0x9e, 0x72, 0x8b, 0x9f, 0x4a, 0x4b, 0xc3,
	0x10, 0xd3, 0x48, 0x64, 0x89, 0x11, 0x56, 0xd2, 0xaa, 0x6f, 0x2c, 0x87, 0xe4, 0x12, 0xce, 0x46,
	0x29, 0x4e, 0xa5, 0x75, 0x09, 0x77, 0x7a, 0xcd, 0xeb, 0x7b, 0x29, 0xb9, 0x86, 0xf3, 0x51, 0x8a,
	0x13, 0x61, 0x3e, 0xa4, 0x65, 0x46, 0xc5, 0x72, 0xf8, 0xfa, 0x4c, 0xeb, 0x5e, 0x2d, 0x82, 0xa2,
	0x3d, 0x9f, 0x85, 0xb4, 0x51, 0x66, 0xcf, 0x67, 0xa1, 0xb3, 0xc7, 0x36, 0xde, 0xb3, 0x4f, 0x7e,
	0xec, 0x02, 0x20, 0x5d, 0x68, 0x4e, 0x55, 0x2c, 0xdd, 0xcf, 0xdc, 0xec, 0x14, 0xbc, 0xb8, 0x93,
	0xfd, 0x75, 0xfa, 0x83, 0xc4, 0xd0, 0xd3, 0x5d, 0xc7, 0x65, 0xdd, 0x2b, 0x68, 0x6d, 0x77, 0xcd,
	0xc4, 0x1a, 0xf3, 0x7f, 0x16, 0x3d, 0x84, 0xf7, 0xed, 0x95, 0x2c, 0x6a, 0xfe, 0x6c, 0xee, 0xbf,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x67, 0x3f, 0x8e, 0x27, 0x49, 0x02, 0x00, 0x00,
}
