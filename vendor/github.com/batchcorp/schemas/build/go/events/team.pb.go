// Code generated by protoc-gen-go. DO NOT EDIT.
// source: team.proto

package events

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Team struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	User                 *User    `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Team) Reset()         { *m = Team{} }
func (m *Team) String() string { return proto.CompactTextString(m) }
func (*Team) ProtoMessage()    {}
func (*Team) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b4e9e93d7b2c6bb, []int{0}
}

func (m *Team) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Team.Unmarshal(m, b)
}
func (m *Team) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Team.Marshal(b, m, deterministic)
}
func (m *Team) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Team.Merge(m, src)
}
func (m *Team) XXX_Size() int {
	return xxx_messageInfo_Team.Size(m)
}
func (m *Team) XXX_DiscardUnknown() {
	xxx_messageInfo_Team.DiscardUnknown(m)
}

var xxx_messageInfo_Team proto.InternalMessageInfo

func (m *Team) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Team) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Team) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*Team)(nil), "events.Team")
}

func init() { proto.RegisterFile("team.proto", fileDescriptor_8b4e9e93d7b2c6bb) }

var fileDescriptor_8b4e9e93d7b2c6bb = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x24, 0x8e, 0xbd, 0xca, 0xc2, 0x30,
	0x18, 0x85, 0x69, 0xbf, 0x52, 0xf8, 0x5e, 0xc5, 0x21, 0x53, 0x70, 0x0a, 0x4e, 0x1d, 0x24, 0x01,
	0xbd, 0x03, 0x67, 0xa7, 0xa2, 0x8b, 0x5b, 0x7e, 0x5e, 0x9a, 0x80, 0x69, 0x4a, 0x7e, 0xbc, 0x7e,
	0x31, 0xd9, 0x0e, 0xcf, 0xe1, 0x39, 0x1c, 0x80, 0x8c, 0xd2, 0xf3, 0x2d, 0x86, 0x1c, 0xc8, 0x88,
	0x1f, 0x5c, 0x73, 0x3a, 0x42, 0x49, 0x18, 0x1b, 0x3b, 0xdd, 0x61, 0x78, 0xa0, 0xf4, 0xe4, 0x00,
	0xbd, 0x33, 0xb4, 0x63, 0xdd, 0xf4, 0x3f, 0xf7, 0xce, 0x10, 0x02, 0xc3, 0x2a, 0x3d, 0xd2, 0xbe,
	0x92, 0x9a, 0x09, 0x83, 0xe1, 0x67, 0xd2, 0x3f, 0xd6, 0x4d, 0xbb, 0xcb, 0x9e, 0xb7, 0x39, 0xfe,
	0x4c, 0x18, 0xe7, 0xda, 0xdc, 0xf8, 0xeb, 0xbc, 0xb8, 0x6c, 0x8b, 0xe2, 0x3a, 0x78, 0xa1, 0x64,
	0xd6, 0x56, 0x87, 0xb8, 0x89, 0xa4, 0x2d, 0x7a, 0x99, 0x84, 0x2a, 0xee, 0x6d, 0xc4, 0x12, 0x44,
	0x53, 0xd5, 0x58, 0x4f, 0x5c, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x30, 0xba, 0x22, 0x36, 0xa6,
	0x00, 0x00, 0x00,
}
