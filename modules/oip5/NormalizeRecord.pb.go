// Code generated by protoc-gen-go. DO NOT EDIT.
// source: NormalizeRecord.proto

package oip5

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// https://github.com/golang/protobuf/blob/882cf97/protoc-gen-go/descriptor/descriptor.proto#L136-L168
type Field_Type int32

const (
	// 0 is reserved for errors.
	Field_TYPE_ERROR Field_Type = 0
	// Order is weird for historical reasons.
	Field_TYPE_DOUBLE Field_Type = 1
	Field_TYPE_FLOAT  Field_Type = 2
	// Not ZigZag encoded.  Negative numbers take 10 bytes.  Use TYPE_SINT64 if
	// negative values are likely.
	Field_TYPE_INT64  Field_Type = 3
	Field_TYPE_UINT64 Field_Type = 4
	// Not ZigZag encoded.  Negative numbers take 10 bytes.  Use TYPE_SINT32 if
	// negative values are likely.
	Field_TYPE_INT32   Field_Type = 5
	Field_TYPE_FIXED64 Field_Type = 6
	Field_TYPE_FIXED32 Field_Type = 7
	Field_TYPE_BOOL    Field_Type = 8
	Field_TYPE_STRING  Field_Type = 9
	// Tag-delimited aggregate.
	// Group type is deprecated and not supported in proto3. However, Proto3
	// implementations should still be able to parse the group wire format and
	// treat group fields as unknown fields.
	Field_TYPE_GROUP   Field_Type = 10
	Field_TYPE_MESSAGE Field_Type = 11
	// New in version 2.
	Field_TYPE_BYTES    Field_Type = 12
	Field_TYPE_UINT32   Field_Type = 13
	Field_TYPE_ENUM     Field_Type = 14
	Field_TYPE_SFIXED32 Field_Type = 15
	Field_TYPE_SFIXED64 Field_Type = 16
	Field_TYPE_SINT32   Field_Type = 17
	Field_TYPE_SINT64   Field_Type = 18
)

var Field_Type_name = map[int32]string{
	0:  "TYPE_ERROR",
	1:  "TYPE_DOUBLE",
	2:  "TYPE_FLOAT",
	3:  "TYPE_INT64",
	4:  "TYPE_UINT64",
	5:  "TYPE_INT32",
	6:  "TYPE_FIXED64",
	7:  "TYPE_FIXED32",
	8:  "TYPE_BOOL",
	9:  "TYPE_STRING",
	10: "TYPE_GROUP",
	11: "TYPE_MESSAGE",
	12: "TYPE_BYTES",
	13: "TYPE_UINT32",
	14: "TYPE_ENUM",
	15: "TYPE_SFIXED32",
	16: "TYPE_SFIXED64",
	17: "TYPE_SINT32",
	18: "TYPE_SINT64",
}
var Field_Type_value = map[string]int32{
	"TYPE_ERROR":    0,
	"TYPE_DOUBLE":   1,
	"TYPE_FLOAT":    2,
	"TYPE_INT64":    3,
	"TYPE_UINT64":   4,
	"TYPE_INT32":    5,
	"TYPE_FIXED64":  6,
	"TYPE_FIXED32":  7,
	"TYPE_BOOL":     8,
	"TYPE_STRING":   9,
	"TYPE_GROUP":    10,
	"TYPE_MESSAGE":  11,
	"TYPE_BYTES":    12,
	"TYPE_UINT32":   13,
	"TYPE_ENUM":     14,
	"TYPE_SFIXED32": 15,
	"TYPE_SFIXED64": 16,
	"TYPE_SINT32":   17,
	"TYPE_SINT64":   18,
}

func (x Field_Type) String() string {
	return proto.EnumName(Field_Type_name, int32(x))
}
func (Field_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{2, 0} }

type NormalizeRecordProto struct {
	// standard template ID which triggers this meta-type
	MainTemplate uint32 `protobuf:"fixed32,3,opt,name=mainTemplate" json:"mainTemplate,omitempty"`
	// fields comprising this new meta-type
	Fields []*NormalField `protobuf:"bytes,4,rep,name=fields" json:"fields,omitempty"`
}

func (m *NormalizeRecordProto) Reset()                    { *m = NormalizeRecordProto{} }
func (m *NormalizeRecordProto) String() string            { return proto.CompactTextString(m) }
func (*NormalizeRecordProto) ProtoMessage()               {}
func (*NormalizeRecordProto) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *NormalizeRecordProto) GetMainTemplate() uint32 {
	if m != nil {
		return m.MainTemplate
	}
	return 0
}

func (m *NormalizeRecordProto) GetFields() []*NormalField {
	if m != nil {
		return m.Fields
	}
	return nil
}

type NormalField struct {
	// new ES compatible name
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// path to the source value
	Path []*Field `protobuf:"bytes,2,rep,name=path" json:"path,omitempty"`
}

func (m *NormalField) Reset()                    { *m = NormalField{} }
func (m *NormalField) String() string            { return proto.CompactTextString(m) }
func (*NormalField) ProtoMessage()               {}
func (*NormalField) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *NormalField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NormalField) GetPath() []*Field {
	if m != nil {
		return m.Path
	}
	return nil
}

type Field struct {
	// Tag number in prior Proto message
	Tag int32 `protobuf:"varint,1,opt,name=tag" json:"tag,omitempty"`
	// Expected type
	// if type_message resolution proceeds into the message
	// if type_message and type_name == oip.txid resolves linked record before proceeding
	// if type doesn't match field is ignored
	Type Field_Type `protobuf:"varint,2,opt,name=type,enum=oip5.Field_Type" json:"type,omitempty"`
	// Enter specified details template
	// Set either tag/type OR template
	Template uint32 `protobuf:"fixed32,3,opt,name=template" json:"template,omitempty"`
}

func (m *Field) Reset()                    { *m = Field{} }
func (m *Field) String() string            { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()               {}
func (*Field) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *Field) GetTag() int32 {
	if m != nil {
		return m.Tag
	}
	return 0
}

func (m *Field) GetType() Field_Type {
	if m != nil {
		return m.Type
	}
	return Field_TYPE_ERROR
}

func (m *Field) GetTemplate() uint32 {
	if m != nil {
		return m.Template
	}
	return 0
}

func init() {
	proto.RegisterType((*NormalizeRecordProto)(nil), "oip5.NormalizeRecordProto")
	proto.RegisterType((*NormalField)(nil), "oip5.NormalField")
	proto.RegisterType((*Field)(nil), "oip5.Field")
	proto.RegisterEnum("oip5.Field_Type", Field_Type_name, Field_Type_value)
}

func init() { proto.RegisterFile("NormalizeRecord.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0x86, 0x2b, 0x4b, 0x76, 0xe2, 0x91, 0xed, 0x8c, 0x97, 0x16, 0x44, 0x2f, 0x35, 0xa2, 0x14,
	0xf5, 0xa2, 0x80, 0xec, 0xf8, 0xd2, 0x53, 0x44, 0x36, 0xc6, 0xe0, 0x58, 0x66, 0x25, 0x41, 0xd3,
	0x4b, 0x51, 0xe2, 0xad, 0x23, 0x90, 0xb2, 0xc2, 0x51, 0x28, 0xe9, 0xf3, 0xf4, 0x59, 0xfa, 0x5c,
	0x45, 0xbb, 0xed, 0xc6, 0xf2, 0x65, 0x99, 0xf9, 0xff, 0x7f, 0xbe, 0x19, 0x81, 0xe0, 0xdd, 0x5a,
	0xec, 0xcb, 0xac, 0xc8, 0x7f, 0x71, 0xc6, 0xef, 0xc5, 0x7e, 0xeb, 0x57, 0x7b, 0x51, 0x0b, 0x62,
	0x89, 0xbc, 0xba, 0x70, 0x39, 0xbc, 0x3d, 0xb2, 0x37, 0xd2, 0x75, 0x61, 0x50, 0x66, 0xf9, 0x63,
	0xc2, 0xcb, 0xaa, 0xc8, 0x6a, 0xee, 0x98, 0x13, 0xc3, 0x3b, 0x61, 0x2d, 0x8d, 0x7c, 0x86, 0xde,
	0x8f, 0x9c, 0x17, 0xdb, 0x27, 0xc7, 0x9a, 0x98, 0x9e, 0x1d, 0x8c, 0xfd, 0x06, 0xe9, 0x2b, 0xde,
	0x75, 0xe3, 0xb0, 0x7f, 0x01, 0x37, 0x04, 0xfb, 0x40, 0x26, 0x04, 0xac, 0xc7, 0xac, 0xe4, 0x8e,
	0x31, 0x31, 0xbc, 0x3e, 0x93, 0x35, 0xf9, 0x00, 0x56, 0x95, 0xd5, 0x0f, 0x4e, 0x47, 0xb2, 0x6c,
	0xc5, 0x52, 0x14, 0x69, 0xb8, 0xbf, 0x4d, 0xe8, 0xaa, 0x71, 0x04, 0xb3, 0xce, 0x76, 0x72, 0xba,
	0xcb, 0x9a, 0x92, 0x7c, 0x04, 0xab, 0x7e, 0xa9, 0xb8, 0xd3, 0x99, 0x18, 0xde, 0x28, 0xc0, 0x83,
	0x61, 0x3f, 0x79, 0xa9, 0x38, 0x93, 0x2e, 0x79, 0x0f, 0xa7, 0x75, 0xfb, 0x83, 0x74, 0xef, 0xfe,
	0xe9, 0x80, 0xd5, 0x44, 0xc9, 0x08, 0x20, 0xb9, 0xdd, 0xd0, 0xef, 0x94, 0xb1, 0x88, 0xe1, 0x1b,
	0x72, 0x06, 0xb6, 0xec, 0xaf, 0xa2, 0x34, 0x5c, 0x51, 0x34, 0x74, 0xe0, 0x7a, 0x15, 0x5d, 0x26,
	0xd8, 0xd1, 0xfd, 0x72, 0x9d, 0xcc, 0x67, 0x68, 0xea, 0x81, 0x54, 0x09, 0xd6, 0x61, 0x60, 0x1a,
	0x60, 0x97, 0x20, 0x0c, 0x14, 0x60, 0xf9, 0x95, 0x5e, 0xcd, 0x67, 0xd8, 0x6b, 0x2b, 0xd3, 0x00,
	0x4f, 0xc8, 0x10, 0xfa, 0x52, 0x09, 0xa3, 0x68, 0x85, 0xa7, 0x9a, 0x19, 0x27, 0x6c, 0xb9, 0x5e,
	0x60, 0x5f, 0x33, 0x17, 0x2c, 0x4a, 0x37, 0x08, 0x9a, 0x70, 0x43, 0xe3, 0xf8, 0x72, 0x41, 0xd1,
	0xd6, 0x89, 0xf0, 0x36, 0xa1, 0x31, 0x0e, 0x5a, 0x67, 0x4d, 0x03, 0x1c, 0xea, 0x15, 0x74, 0x9d,
	0xde, 0xe0, 0x88, 0x8c, 0x61, 0xa8, 0x56, 0xfc, 0x3f, 0xe2, 0xec, 0x48, 0x9a, 0xcf, 0x10, 0x5f,
	0x0f, 0x51, 0x94, 0x71, 0x4b, 0x98, 0xcf, 0x90, 0x84, 0xde, 0xb7, 0x4f, 0xbb, 0xbc, 0x7e, 0x78,
	0xbe, 0xf3, 0xef, 0x45, 0x79, 0x2e, 0xf2, 0xea, 0xe7, 0xae, 0x79, 0xcf, 0x4b, 0xb1, 0x7d, 0x2e,
	0xf8, 0x53, 0x53, 0x5f, 0x7c, 0x69, 0x9e, 0xbb, 0x9e, 0xfc, 0x11, 0xa7, 0x7f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x5b, 0x21, 0xc0, 0x16, 0xa1, 0x02, 0x00, 0x00,
}
