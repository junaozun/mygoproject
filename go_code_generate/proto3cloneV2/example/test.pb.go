// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package example

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	sometype "go_gen/proto3cloneV2/example/sometype"
	example2 "go_gen/proto3cloneV2/example2"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 请求Debug命令
// |RequestServiceDebugCmd2
type RequestServiceDebugCmd2 struct {
	Cmd                  string                              `protobuf:"bytes,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Commanders           [][]byte                            `protobuf:"bytes,6,rep,name=commanders,proto3" json:"commanders,omitempty"`
	Apple                *sometype.ApplesServerin            `protobuf:"bytes,7,opt,name=apple,proto3" json:"apple,omitempty"`
	Tem                  map[string]int32                    `protobuf:"bytes,8,rep,name=tem,proto3" json:"tem,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	AddrEntry            map[string]*sometype.ApplesServerin `protobuf:"bytes,10,rep,name=addrEntry,proto3" json:"addrEntry,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Hero                 *example2.Hero                      `protobuf:"bytes,12,opt,name=hero,proto3" json:"hero,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *RequestServiceDebugCmd2) Reset()         { *m = RequestServiceDebugCmd2{} }
func (m *RequestServiceDebugCmd2) String() string { return proto.CompactTextString(m) }
func (*RequestServiceDebugCmd2) ProtoMessage()    {}
func (*RequestServiceDebugCmd2) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *RequestServiceDebugCmd2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestServiceDebugCmd2.Unmarshal(m, b)
}
func (m *RequestServiceDebugCmd2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestServiceDebugCmd2.Marshal(b, m, deterministic)
}
func (m *RequestServiceDebugCmd2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestServiceDebugCmd2.Merge(m, src)
}
func (m *RequestServiceDebugCmd2) XXX_Size() int {
	return xxx_messageInfo_RequestServiceDebugCmd2.Size(m)
}
func (m *RequestServiceDebugCmd2) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestServiceDebugCmd2.DiscardUnknown(m)
}

var xxx_messageInfo_RequestServiceDebugCmd2 proto.InternalMessageInfo

func (m *RequestServiceDebugCmd2) GetCmd() string {
	if m != nil {
		return m.Cmd
	}
	return ""
}

func (m *RequestServiceDebugCmd2) GetCommanders() [][]byte {
	if m != nil {
		return m.Commanders
	}
	return nil
}

func (m *RequestServiceDebugCmd2) GetApple() *sometype.ApplesServerin {
	if m != nil {
		return m.Apple
	}
	return nil
}

func (m *RequestServiceDebugCmd2) GetTem() map[string]int32 {
	if m != nil {
		return m.Tem
	}
	return nil
}

func (m *RequestServiceDebugCmd2) GetAddrEntry() map[string]*sometype.ApplesServerin {
	if m != nil {
		return m.AddrEntry
	}
	return nil
}

func (m *RequestServiceDebugCmd2) GetHero() *example2.Hero {
	if m != nil {
		return m.Hero
	}
	return nil
}

func init() {
	proto.RegisterType((*RequestServiceDebugCmd2)(nil), "example.RequestServiceDebugCmd2")
	proto.RegisterMapType((map[string]*sometype.ApplesServerin)(nil), "example.RequestServiceDebugCmd2.AddrEntryEntry")
	proto.RegisterMapType((map[string]int32)(nil), "example.RequestServiceDebugCmd2.TemEntry")
}

func init() {
	proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e)
}

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0x4b, 0x4b, 0xf4, 0x30,
	0x14, 0xa5, 0xd3, 0xaf, 0xf3, 0xb8, 0x33, 0x94, 0x8f, 0x20, 0x18, 0xba, 0x90, 0x32, 0xab, 0xba,
	0x89, 0x50, 0x41, 0x44, 0x57, 0xa2, 0x82, 0x1b, 0x37, 0x51, 0xdc, 0x77, 0xda, 0x8b, 0x0e, 0xd3,
	0x34, 0x35, 0x49, 0x07, 0xfb, 0xdb, 0xfc, 0x73, 0xd2, 0x64, 0x5a, 0x47, 0xf0, 0xb1, 0xeb, 0x3d,
	0xa7, 0xe7, 0xc1, 0x09, 0x80, 0x41, 0x6d, 0x58, 0xad, 0xa4, 0x91, 0x64, 0x82, 0x6f, 0x99, 0xa8,
	0x4b, 0x8c, 0x42, 0x2d, 0x05, 0x9a, 0xb6, 0x46, 0x47, 0x44, 0xe1, 0x8e, 0x48, 0xdd, 0xbd, 0x7c,
	0xf7, 0xe1, 0x90, 0xe3, 0x6b, 0x83, 0xda, 0x3c, 0xa0, 0xda, 0xae, 0x73, 0xbc, 0xc1, 0x55, 0xf3,
	0x7c, 0x2d, 0x8a, 0x94, 0xfc, 0x07, 0x3f, 0x17, 0x05, 0xf5, 0x62, 0x2f, 0x99, 0xf1, 0xee, 0x93,
	0x1c, 0x01, 0xe4, 0x52, 0x88, 0xac, 0x2a, 0x50, 0x69, 0x3a, 0x8e, 0xfd, 0x64, 0xc1, 0xf7, 0x10,
	0xc2, 0x20, 0xc8, 0xea, 0xba, 0x44, 0x3a, 0x89, 0xbd, 0x64, 0x9e, 0x52, 0x36, 0xa4, 0x5b, 0x58,
	0x77, 0x11, 0xa8, 0xd6, 0x15, 0x77, 0xbf, 0x91, 0x4b, 0xf0, 0x0d, 0x0a, 0x3a, 0x8d, 0xfd, 0x64,
	0x9e, 0x1e, 0xb3, 0x5d, 0x37, 0xf6, 0x43, 0x21, 0xf6, 0x88, 0xe2, 0xb6, 0x32, 0xaa, 0xe5, 0x9d,
	0x8a, 0xdc, 0xc3, 0x2c, 0x2b, 0x0a, 0x65, 0x11, 0x0a, 0xd6, 0xe2, 0xe4, 0x4f, 0x8b, 0xab, 0x5e,
	0xe1, 0x8c, 0x3e, 0x1d, 0xc8, 0x12, 0xfe, 0xbd, 0xa0, 0x92, 0x74, 0x61, 0xab, 0x87, 0x6c, 0x18,
	0xea, 0x0e, 0x95, 0xe4, 0x96, 0x8b, 0xce, 0x60, 0xda, 0x77, 0xe8, 0xd6, 0xd9, 0x60, 0xdb, 0xaf,
	0xb3, 0xc1, 0x96, 0x1c, 0x40, 0xb0, 0xcd, 0xca, 0x06, 0xe9, 0x28, 0xf6, 0x92, 0x80, 0xbb, 0xe3,
	0x62, 0x74, 0xee, 0x45, 0x4f, 0x10, 0x7e, 0x0d, 0xfe, 0x46, 0xcd, 0xf6, 0xd5, 0xbf, 0x6e, 0x37,
	0xf8, 0xae, 0xc6, 0xf6, 0x11, 0x4f, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1e, 0xdb, 0x9a, 0x27,
	0xfb, 0x01, 0x00, 0x00,
}