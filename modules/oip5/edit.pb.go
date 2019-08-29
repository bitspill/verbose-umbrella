// Code generated by protoc-gen-go. DO NOT EDIT.
// source: edit.proto

/*
Package oip5 is a generated protocol buffer package.

It is generated from these files:
	edit.proto
	NormalizeRecord.proto
	NormalizeRecord.proto
	oip5.proto
	Record.proto
	RecordTemplateProto.proto

It has these top-level messages:
	EditProto
	Op
	Step
	NormalizeRecordProto
	NormalField
	Field
	NormalizeRecordProto
	NormalField
	Field
	OipFive
	Transfer
	Deactivate
	RecordProto
	Permissions
	Payment
	OipDetails
	RecordTemplateProto
*/
package oip5

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import oip "github.com/oipwg/oip/modules/oip"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Step_Action int32

const (
	// 0 is reserved for errors.
	Step_ACTION_ERROR        Step_Action = 0
	Step_ACTION_REPLACE_ALL  Step_Action = 1
	Step_ACTION_APPEND_ALL   Step_Action = 2
	Step_ACTION_REMOVE_ALL   Step_Action = 3
	Step_ACTION_REMOVE_ONE   Step_Action = 4
	Step_ACTION_REPLACE_ONE  Step_Action = 5
	Step_ACTION_STRING_PATCH Step_Action = 6
	Step_ACTION_STEP_INTO    Step_Action = 7
)

var Step_Action_name = map[int32]string{
	0: "ACTION_ERROR",
	1: "ACTION_REPLACE_ALL",
	2: "ACTION_APPEND_ALL",
	3: "ACTION_REMOVE_ALL",
	4: "ACTION_REMOVE_ONE",
	5: "ACTION_REPLACE_ONE",
	6: "ACTION_STRING_PATCH",
	7: "ACTION_STEP_INTO",
}
var Step_Action_value = map[string]int32{
	"ACTION_ERROR":        0,
	"ACTION_REPLACE_ALL":  1,
	"ACTION_APPEND_ALL":   2,
	"ACTION_REMOVE_ALL":   3,
	"ACTION_REMOVE_ONE":   4,
	"ACTION_REPLACE_ONE":  5,
	"ACTION_STRING_PATCH": 6,
	"ACTION_STEP_INTO":    7,
}

func (x Step_Action) String() string {
	return proto.EnumName(Step_Action_name, int32(x))
}
func (Step_Action) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type EditProto struct {
	Reference *oip.Txid    `protobuf:"bytes,1,opt,name=reference" json:"reference,omitempty"`
	NewValues *RecordProto `protobuf:"bytes,2,opt,name=newValues" json:"newValues,omitempty"`
	Ops       []*Op        `protobuf:"bytes,3,rep,name=ops" json:"ops,omitempty"`
}

func (m *EditProto) Reset()                    { *m = EditProto{} }
func (m *EditProto) String() string            { return proto.CompactTextString(m) }
func (*EditProto) ProtoMessage()               {}
func (*EditProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EditProto) GetReference() *oip.Txid {
	if m != nil {
		return m.Reference
	}
	return nil
}

func (m *EditProto) GetNewValues() *RecordProto {
	if m != nil {
		return m.NewValues
	}
	return nil
}

func (m *EditProto) GetOps() []*Op {
	if m != nil {
		return m.Ops
	}
	return nil
}

type Op struct {
	Path []*Step `protobuf:"bytes,1,rep,name=Path" json:"Path,omitempty"`
}

func (m *Op) Reset()                    { *m = Op{} }
func (m *Op) String() string            { return proto.CompactTextString(m) }
func (*Op) ProtoMessage()               {}
func (*Op) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Op) GetPath() []*Step {
	if m != nil {
		return m.Path
	}
	return nil
}

type Step struct {
	Tag      int32       `protobuf:"varint,1,opt,name=tag" json:"tag,omitempty"`
	Action   Step_Action `protobuf:"varint,2,opt,name=action,enum=oip5.Step_Action" json:"action,omitempty"`
	SrcIndex int32       `protobuf:"varint,3,opt,name=srcIndex" json:"srcIndex,omitempty"`
	DstIndex int32       `protobuf:"varint,4,opt,name=dstIndex" json:"dstIndex,omitempty"`
}

func (m *Step) Reset()                    { *m = Step{} }
func (m *Step) String() string            { return proto.CompactTextString(m) }
func (*Step) ProtoMessage()               {}
func (*Step) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Step) GetTag() int32 {
	if m != nil {
		return m.Tag
	}
	return 0
}

func (m *Step) GetAction() Step_Action {
	if m != nil {
		return m.Action
	}
	return Step_ACTION_ERROR
}

func (m *Step) GetSrcIndex() int32 {
	if m != nil {
		return m.SrcIndex
	}
	return 0
}

func (m *Step) GetDstIndex() int32 {
	if m != nil {
		return m.DstIndex
	}
	return 0
}

func init() {
	proto.RegisterType((*EditProto)(nil), "oip5.EditProto")
	proto.RegisterType((*Op)(nil), "oip5.Op")
	proto.RegisterType((*Step)(nil), "oip5.Step")
	proto.RegisterEnum("oip5.Step_Action", Step_Action_name, Step_Action_value)
}

func init() { proto.RegisterFile("edit.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4d, 0x6b, 0xdc, 0x30,
	0x10, 0x86, 0xeb, 0x8f, 0xb8, 0xf1, 0x24, 0x14, 0x45, 0xfd, 0x32, 0x3e, 0x94, 0x60, 0x4a, 0xbb,
	0xbd, 0x78, 0x21, 0x65, 0x4f, 0x3d, 0xb9, 0x5b, 0xd1, 0x1a, 0xb6, 0xb6, 0x51, 0x4c, 0x0e, 0xbd,
	0x18, 0xc7, 0x56, 0x37, 0x82, 0xc4, 0x12, 0xb6, 0x96, 0xec, 0xb5, 0xff, 0xab, 0xf4, 0xb7, 0x15,
	0x49, 0x4e, 0x52, 0x68, 0x2e, 0x62, 0xe6, 0x79, 0x1f, 0x5e, 0xcf, 0xc1, 0x00, 0xac, 0xe7, 0x2a,
	0x95, 0xa3, 0x50, 0x02, 0xfb, 0x82, 0xcb, 0x55, 0x7c, 0x4c, 0x59, 0x27, 0xc6, 0xde, 0xb2, 0x18,
	0xd4, 0x9e, 0xcf, 0x73, 0xf2, 0xcb, 0x81, 0x90, 0xf4, 0x5c, 0x55, 0xc6, 0x7e, 0x0f, 0xe1, 0xc8,
	0x7e, 0xb2, 0x91, 0x0d, 0x1d, 0x8b, 0x9c, 0x53, 0x67, 0x71, 0x74, 0x16, 0xa6, 0x82, 0xcb, 0xb4,
	0xde, 0xf3, 0x9e, 0x3e, 0x64, 0x78, 0x09, 0xe1, 0xc0, 0x6e, 0x2f, 0xda, 0xeb, 0x1d, 0x9b, 0x22,
	0xd7, 0x88, 0x27, 0x5a, 0x5c, 0xa5, 0xf6, 0x4b, 0xa6, 0x8e, 0x3e, 0x38, 0x38, 0x06, 0x4f, 0xc8,
	0x29, 0xf2, 0x4e, 0xbd, 0xc5, 0xd1, 0xd9, 0xa1, 0x55, 0x4b, 0x49, 0x35, 0x4c, 0xde, 0x82, 0x5b,
	0x4a, 0xfc, 0x06, 0xfc, 0xaa, 0x55, 0x57, 0x91, 0x63, 0x14, 0xb0, 0xca, 0xb9, 0x62, 0x92, 0x1a,
	0x9e, 0xfc, 0x76, 0xc1, 0xd7, 0x2b, 0x46, 0xe0, 0xa9, 0x76, 0x6b, 0xce, 0x3b, 0xa0, 0x7a, 0xc4,
	0x1f, 0x20, 0x68, 0x3b, 0xc5, 0xc5, 0x60, 0x4e, 0x79, 0x76, 0x77, 0x8a, 0xb6, 0xd3, 0xcc, 0x04,
	0x74, 0x16, 0x70, 0x0c, 0x87, 0xd3, 0xd8, 0xe5, 0x43, 0xcf, 0xf6, 0x91, 0x67, 0x1a, 0xee, 0x77,
	0x9d, 0xf5, 0x93, 0xb2, 0x99, 0x6f, 0xb3, 0xbb, 0x3d, 0xf9, 0xe3, 0x40, 0x60, 0xab, 0x30, 0x82,
	0xe3, 0x6c, 0x5d, 0xe7, 0x65, 0xd1, 0x10, 0x4a, 0x4b, 0x8a, 0x9e, 0xe0, 0x57, 0x80, 0x67, 0x42,
	0x49, 0xb5, 0xc9, 0xd6, 0xa4, 0xc9, 0x36, 0x1b, 0xe4, 0xe0, 0x97, 0x70, 0x32, 0xf3, 0xac, 0xaa,
	0x48, 0xf1, 0xc5, 0x60, 0xf7, 0x1f, 0x4c, 0xc9, 0xf7, 0xf2, 0xc2, 0xda, 0xde, 0xff, 0xb8, 0x2c,
	0x08, 0xf2, 0x1f, 0x29, 0xd7, 0xfc, 0x00, 0xbf, 0x86, 0xe7, 0x33, 0x3f, 0xaf, 0x69, 0x5e, 0x7c,
	0x6d, 0xaa, 0xac, 0x5e, 0x7f, 0x43, 0x01, 0x7e, 0x01, 0xe8, 0x3e, 0x20, 0x55, 0x93, 0x17, 0x75,
	0x89, 0x9e, 0x7e, 0x5e, 0xfc, 0x78, 0xb7, 0xe5, 0xea, 0x6a, 0x77, 0x99, 0x76, 0xe2, 0x66, 0x29,
	0xb8, 0xbc, 0xdd, 0xea, 0x77, 0x79, 0x23, 0xfa, 0xdd, 0x35, 0x9b, 0xf4, 0xbc, 0xfa, 0xa4, 0x9f,
	0xcb, 0xc0, 0xfc, 0x19, 0x1f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x88, 0x88, 0x99, 0xbd, 0x47,
	0x02, 0x00, 0x00,
}
