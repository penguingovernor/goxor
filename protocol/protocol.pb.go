// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol/protocol.proto

package protocol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type Key struct {
	Signature            []byte   `protobuf:"bytes,1,opt,name=Signature,proto3" json:"Signature,omitempty"`
	PayLoad              []byte   `protobuf:"bytes,2,opt,name=PayLoad,proto3" json:"PayLoad,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_87968d26f3046c60, []int{0}
}
func (m *Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key.Unmarshal(m, b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key.Marshal(b, m, deterministic)
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return xxx_messageInfo_Key.Size(m)
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func (m *Key) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Key) GetPayLoad() []byte {
	if m != nil {
		return m.PayLoad
	}
	return nil
}

type Data struct {
	Signature            []byte   `protobuf:"bytes,1,opt,name=Signature,proto3" json:"Signature,omitempty"`
	PayLoad              []byte   `protobuf:"bytes,2,opt,name=PayLoad,proto3" json:"PayLoad,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_87968d26f3046c60, []int{1}
}
func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Data) GetPayLoad() []byte {
	if m != nil {
		return m.PayLoad
	}
	return nil
}

func init() {
	proto.RegisterType((*Key)(nil), "protocol.Key")
	proto.RegisterType((*Data)(nil), "protocol.Data")
}

func init() { proto.RegisterFile("protocol/protocol.proto", fileDescriptor_87968d26f3046c60) }

var fileDescriptor_87968d26f3046c60 = []byte{
	// 104 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x87, 0x31, 0xf4, 0xc0, 0x0c, 0x21, 0x0e, 0x18, 0x5f, 0xc9, 0x96,
	0x8b, 0xd9, 0x3b, 0xb5, 0x52, 0x48, 0x86, 0x8b, 0x33, 0x38, 0x33, 0x3d, 0x2f, 0xb1, 0xa4, 0xb4,
	0x28, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x27, 0x08, 0x21, 0x20, 0x24, 0xc1, 0xc5, 0x1e, 0x90,
	0x58, 0xe9, 0x93, 0x9f, 0x98, 0x22, 0xc1, 0x04, 0x96, 0x83, 0x71, 0x95, 0xec, 0xb8, 0x58, 0x5c,
	0x12, 0x4b, 0x12, 0xc9, 0xd5, 0x9f, 0xc4, 0x06, 0x76, 0x88, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0x7f, 0x2f, 0x3f, 0xd7, 0xaa, 0x00, 0x00, 0x00,
}