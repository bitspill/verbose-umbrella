// Code generated by protoc-gen-go. DO NOT EDIT.
// source: multipart.proto

package oipProto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MultiPart struct {
	CurrentPart uint32 `protobuf:"varint,1,opt,name=currentPart" json:"currentPart,omitempty"`
	CountParts  uint32 `protobuf:"varint,2,opt,name=countParts" json:"countParts,omitempty"`
	RawData     []byte `protobuf:"bytes,3,opt,name=rawData,proto3" json:"rawData,omitempty"`
	Reference   *Txid  `protobuf:"bytes,4,opt,name=reference" json:"reference,omitempty"`
}

func (m *MultiPart) Reset()                    { *m = MultiPart{} }
func (m *MultiPart) String() string            { return proto.CompactTextString(m) }
func (*MultiPart) ProtoMessage()               {}
func (*MultiPart) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *MultiPart) GetCurrentPart() uint32 {
	if m != nil {
		return m.CurrentPart
	}
	return 0
}

func (m *MultiPart) GetCountParts() uint32 {
	if m != nil {
		return m.CountParts
	}
	return 0
}

func (m *MultiPart) GetRawData() []byte {
	if m != nil {
		return m.RawData
	}
	return nil
}

func (m *MultiPart) GetReference() *Txid {
	if m != nil {
		return m.Reference
	}
	return nil
}

func init() {
	proto.RegisterType((*MultiPart)(nil), "oipProto.MultiPart")
}

func init() { proto.RegisterFile("multipart.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x2d, 0xcd, 0x29,
	0xc9, 0x2c, 0x48, 0x2c, 0x2a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0xcf, 0x2c,
	0x08, 0x00, 0xb1, 0xa4, 0xb8, 0x4a, 0x2a, 0x32, 0x53, 0x20, 0xa2, 0x4a, 0x53, 0x19, 0xb9, 0x38,
	0x7d, 0x41, 0x2a, 0x03, 0x12, 0x8b, 0x4a, 0x84, 0x14, 0xb8, 0xb8, 0x93, 0x4b, 0x8b, 0x8a, 0x52,
	0xf3, 0x4a, 0x40, 0x5c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xde, 0x20, 0x64, 0x21, 0x21, 0x39, 0x2e,
	0xae, 0xe4, 0xfc, 0x52, 0x08, 0xa7, 0x58, 0x82, 0x09, 0xac, 0x00, 0x49, 0x44, 0x48, 0x82, 0x8b,
	0xbd, 0x28, 0xb1, 0xdc, 0x25, 0xb1, 0x24, 0x51, 0x82, 0x59, 0x81, 0x51, 0x83, 0x27, 0x08, 0xc6,
	0x15, 0xd2, 0xe1, 0xe2, 0x2c, 0x4a, 0x4d, 0x4b, 0x2d, 0x4a, 0xcd, 0x4b, 0x4e, 0x95, 0x60, 0x51,
	0x60, 0xd4, 0xe0, 0x36, 0xe2, 0xd3, 0x83, 0xb9, 0x49, 0x2f, 0xa4, 0x22, 0x33, 0x25, 0x08, 0xa1,
	0xc0, 0x89, 0x2b, 0x0a, 0xee, 0xde, 0x24, 0x36, 0xb0, 0x53, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x84, 0x22, 0x31, 0x4a, 0xd3, 0x00, 0x00, 0x00,
}
