// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config/protocol/protocol.proto

package protocol

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	redis "github.com/samaritan-proxy/samaritan/pb/config/protocol/redis"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

// Protocol enum.
type Protocol int32

const (
	UNKNOWN Protocol = 0
	// TCP
	TCP Protocol = 1
	// MySQL
	MySQL Protocol = 2
	// Redis
	Redis Protocol = 3
)

var Protocol_name = map[int32]string{
	0: "UNKNOWN",
	1: "TCP",
	2: "MySQL",
	3: "Redis",
}

var Protocol_value = map[string]int32{
	"UNKNOWN": 0,
	"TCP":     1,
	"MySQL":   2,
	"Redis":   3,
}

func (x Protocol) String() string {
	return proto.EnumName(Protocol_name, int32(x))
}

func (Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_024f19dd69f1399f, []int{0}
}

// MySQL protocol option.
type MySQLOption struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MySQLOption) Reset()         { *m = MySQLOption{} }
func (m *MySQLOption) String() string { return proto.CompactTextString(m) }
func (*MySQLOption) ProtoMessage()    {}
func (*MySQLOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_024f19dd69f1399f, []int{0}
}
func (m *MySQLOption) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MySQLOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MySQLOption.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MySQLOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MySQLOption.Merge(m, src)
}
func (m *MySQLOption) XXX_Size() int {
	return m.Size()
}
func (m *MySQLOption) XXX_DiscardUnknown() {
	xxx_messageInfo_MySQLOption.DiscardUnknown(m)
}

var xxx_messageInfo_MySQLOption proto.InternalMessageInfo

// TCP protocol option.
type TCPOption struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TCPOption) Reset()         { *m = TCPOption{} }
func (m *TCPOption) String() string { return proto.CompactTextString(m) }
func (*TCPOption) ProtoMessage()    {}
func (*TCPOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_024f19dd69f1399f, []int{1}
}
func (m *TCPOption) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TCPOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TCPOption.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TCPOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TCPOption.Merge(m, src)
}
func (m *TCPOption) XXX_Size() int {
	return m.Size()
}
func (m *TCPOption) XXX_DiscardUnknown() {
	xxx_messageInfo_TCPOption.DiscardUnknown(m)
}

var xxx_messageInfo_TCPOption proto.InternalMessageInfo

// Redis protocol option.
type RedisOption struct {
	// Strategy of a read only command.
	ReadStrategy redis.ReadStrategy `protobuf:"varint,1,opt,name=read_strategy,json=readStrategy,proto3,enum=redis.ReadStrategy" json:"read_strategy,omitempty"`
	// Configuration of compression.
	Compression          *redis.Compression `protobuf:"bytes,2,opt,name=compression,proto3" json:"compression,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RedisOption) Reset()         { *m = RedisOption{} }
func (m *RedisOption) String() string { return proto.CompactTextString(m) }
func (*RedisOption) ProtoMessage()    {}
func (*RedisOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_024f19dd69f1399f, []int{2}
}
func (m *RedisOption) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RedisOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RedisOption.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RedisOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisOption.Merge(m, src)
}
func (m *RedisOption) XXX_Size() int {
	return m.Size()
}
func (m *RedisOption) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisOption.DiscardUnknown(m)
}

var xxx_messageInfo_RedisOption proto.InternalMessageInfo

func (m *RedisOption) GetReadStrategy() redis.ReadStrategy {
	if m != nil {
		return m.ReadStrategy
	}
	return redis.ReadStrategy_MASTER
}

func (m *RedisOption) GetCompression() *redis.Compression {
	if m != nil {
		return m.Compression
	}
	return nil
}

func init() {
	proto.RegisterEnum("protocol.Protocol", Protocol_name, Protocol_value)
	proto.RegisterType((*MySQLOption)(nil), "protocol.MySQLOption")
	proto.RegisterType((*TCPOption)(nil), "protocol.TCPOption")
	proto.RegisterType((*RedisOption)(nil), "protocol.RedisOption")
}

func init() { proto.RegisterFile("config/protocol/protocol.proto", fileDescriptor_024f19dd69f1399f) }

var fileDescriptor_024f19dd69f1399f = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xce, 0xcf, 0x4b,
	0xcb, 0x4c, 0xd7, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0x81, 0x33, 0xf4, 0xc0, 0x0c,
	0x21, 0x0e, 0x18, 0x5f, 0x4a, 0x24, 0x3d, 0x3f, 0x3d, 0x1f, 0xcc, 0xd3, 0x07, 0xb1, 0x20, 0xf2,
	0x52, 0xe2, 0x65, 0x89, 0x39, 0x99, 0x29, 0x89, 0x25, 0xa9, 0xfa, 0x30, 0x06, 0x54, 0x42, 0x11,
	0xdd, 0xe0, 0xa2, 0xd4, 0x94, 0xcc, 0x62, 0x08, 0x09, 0x51, 0xa2, 0xc4, 0xcb, 0xc5, 0xed, 0x5b,
	0x19, 0x1c, 0xe8, 0xe3, 0x5f, 0x50, 0x92, 0x99, 0x9f, 0xa7, 0xc4, 0xcd, 0xc5, 0x19, 0xe2, 0x1c,
	0x00, 0xe5, 0x34, 0x33, 0x72, 0x71, 0x07, 0x81, 0xd4, 0x42, 0xf8, 0x42, 0x16, 0x5c, 0xbc, 0x45,
	0xa9, 0x89, 0x29, 0xf1, 0xc5, 0x25, 0x45, 0x89, 0x25, 0xa9, 0xe9, 0x95, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x7c, 0x46, 0xc2, 0x7a, 0x10, 0x03, 0x83, 0x52, 0x13, 0x53, 0x82, 0xa1, 0x52, 0x41, 0x3c,
	0x45, 0x48, 0x3c, 0x21, 0x13, 0x2e, 0xee, 0xe4, 0xfc, 0xdc, 0x82, 0xa2, 0xd4, 0xe2, 0xe2, 0xcc,
	0xfc, 0x3c, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x21, 0xa8, 0x3e, 0x67, 0x84, 0x4c, 0x10,
	0xb2, 0x32, 0x2f, 0x16, 0x0e, 0x66, 0x01, 0x56, 0x2d, 0x1b, 0x2e, 0x8e, 0x00, 0xa8, 0xfb, 0x85,
	0xb8, 0xb9, 0xd8, 0x43, 0xfd, 0xbc, 0xfd, 0xfc, 0xc3, 0xfd, 0x04, 0x18, 0x84, 0xd8, 0xb9, 0x98,
	0x43, 0x9c, 0x03, 0x04, 0x18, 0x85, 0x38, 0xb9, 0x58, 0xc1, 0x7e, 0x10, 0x60, 0x02, 0x31, 0xc1,
	0x2e, 0x16, 0x60, 0x96, 0x62, 0xe9, 0x58, 0x2c, 0xc7, 0xe0, 0x14, 0xf9, 0xe0, 0xa1, 0x1c, 0xe3,
	0x87, 0x87, 0x72, 0x8c, 0x2b, 0x1e, 0xc9, 0x31, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c,
	0xe3, 0x83, 0x47, 0x72, 0x8c, 0x2f, 0x1e, 0xc9, 0x31, 0x7e, 0x78, 0x24, 0xc7, 0x10, 0x65, 0x9e,
	0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x5f, 0x9c, 0x98, 0x9b, 0x58, 0x94,
	0x59, 0x92, 0x98, 0xa7, 0x5b, 0x50, 0x94, 0x5f, 0x51, 0x89, 0xe0, 0xeb, 0x17, 0x24, 0xe9, 0xa3,
	0x05, 0x66, 0x12, 0x1b, 0x98, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x13, 0xb6, 0x2b,
	0xbf, 0x01, 0x00, 0x00,
}

func (this *MySQLOption) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*MySQLOption)
	if !ok {
		that2, ok := that.(MySQLOption)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *MySQLOption")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *MySQLOption but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *MySQLOption but is not nil && this == nil")
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *MySQLOption) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MySQLOption)
	if !ok {
		that2, ok := that.(MySQLOption)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TCPOption) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*TCPOption)
	if !ok {
		that2, ok := that.(TCPOption)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *TCPOption")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *TCPOption but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *TCPOption but is not nil && this == nil")
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *TCPOption) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TCPOption)
	if !ok {
		that2, ok := that.(TCPOption)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RedisOption) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*RedisOption)
	if !ok {
		that2, ok := that.(RedisOption)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *RedisOption")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *RedisOption but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *RedisOption but is not nil && this == nil")
	}
	if this.ReadStrategy != that1.ReadStrategy {
		return fmt.Errorf("ReadStrategy this(%v) Not Equal that(%v)", this.ReadStrategy, that1.ReadStrategy)
	}
	if !this.Compression.Equal(that1.Compression) {
		return fmt.Errorf("Compression this(%v) Not Equal that(%v)", this.Compression, that1.Compression)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *RedisOption) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RedisOption)
	if !ok {
		that2, ok := that.(RedisOption)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ReadStrategy != that1.ReadStrategy {
		return false
	}
	if !this.Compression.Equal(that1.Compression) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MySQLOption) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&protocol.MySQLOption{")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TCPOption) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&protocol.TCPOption{")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *RedisOption) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&protocol.RedisOption{")
	s = append(s, "ReadStrategy: "+fmt.Sprintf("%#v", this.ReadStrategy)+",\n")
	if this.Compression != nil {
		s = append(s, "Compression: "+fmt.Sprintf("%#v", this.Compression)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringProtocol(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *MySQLOption) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MySQLOption) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MySQLOption) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *TCPOption) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TCPOption) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TCPOption) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *RedisOption) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RedisOption) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RedisOption) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Compression != nil {
		{
			size, err := m.Compression.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintProtocol(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ReadStrategy != 0 {
		i = encodeVarintProtocol(dAtA, i, uint64(m.ReadStrategy))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProtocol(dAtA []byte, offset int, v uint64) int {
	offset -= sovProtocol(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MySQLOption) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TCPOption) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RedisOption) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ReadStrategy != 0 {
		n += 1 + sovProtocol(uint64(m.ReadStrategy))
	}
	if m.Compression != nil {
		l = m.Compression.Size()
		n += 1 + l + sovProtocol(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovProtocol(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProtocol(x uint64) (n int) {
	return sovProtocol(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MySQLOption) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
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
			return fmt.Errorf("proto: MySQLOption: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MySQLOption: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TCPOption) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
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
			return fmt.Errorf("proto: TCPOption: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TCPOption: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RedisOption) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
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
			return fmt.Errorf("proto: RedisOption: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RedisOption: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReadStrategy", wireType)
			}
			m.ReadStrategy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReadStrategy |= redis.ReadStrategy(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Compression", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
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
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProtocol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Compression == nil {
				m.Compression = &redis.Compression{}
			}
			if err := m.Compression.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipProtocol(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProtocol
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
					return 0, ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProtocol
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
				return 0, ErrInvalidLengthProtocol
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthProtocol
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowProtocol
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipProtocol(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthProtocol
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthProtocol = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProtocol   = fmt.Errorf("proto: integer overflow")
)
