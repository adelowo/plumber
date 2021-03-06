// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "events.User")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0xce, 0xb1, 0xaa, 0xc2, 0x30,
	0x14, 0xc6, 0x71, 0xda, 0xdb, 0x5b, 0xec, 0x19, 0x1c, 0x82, 0x43, 0x41, 0x04, 0x71, 0x72, 0x90,
	0x66, 0xf0, 0x0d, 0x7c, 0x00, 0x07, 0xc1, 0xc5, 0x45, 0x92, 0xf4, 0xd8, 0x04, 0x9a, 0xa6, 0xe4,
	0x24, 0x3e, 0xbf, 0x78, 0x8a, 0xe3, 0xf7, 0xff, 0x2d, 0x1f, 0x40, 0x26, 0x8c, 0xdd, 0x1c, 0x43,
	0x0a, 0xa2, 0xc6, 0x37, 0x4e, 0x89, 0x0e, 0x16, 0xaa, 0x3b, 0x61, 0x14, 0x6b, 0x28, 0x5d, 0xdf,
	0x16, 0xfb, 0xe2, 0xd8, 0xdc, 0x4a, 0xd7, 0x8b, 0x1d, 0xc0, 0xcb, 0x45, 0x4a, 0xcf, 0x49, 0x79,
	0x6c, 0x4b, 0xee, 0x0d, 0x97, 0xab, 0xf2, 0x28, 0xb6, 0xd0, 0x8c, 0xea, 0xa7, 0x7f, 0xac, 0xab,
	0x6f, 0x60, 0xdc, 0xc0, 0x3f, 0x7a, 0xe5, 0xc6, 0xb6, 0x62, 0x58, 0xc6, 0xa5, 0x7b, 0x9c, 0x06,
	0x97, 0x6c, 0xd6, 0x9d, 0x09, 0x5e, 0x6a, 0x95, 0x8c, 0x35, 0x21, 0xce, 0x92, 0x8c, 0x45, 0xaf,
	0x48, 0xea, 0xec, 0xc6, 0x5e, 0x0e, 0x41, 0x2e, 0xcf, 0x74, 0xcd, 0x47, 0xcf, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x9d, 0xd4, 0x7e, 0x66, 0xb6, 0x00, 0x00, 0x00,
}
