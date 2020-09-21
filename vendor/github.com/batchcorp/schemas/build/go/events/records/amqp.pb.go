// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/records/amqp.proto

package records

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

// See: https://godoc.org/github.com/streadway/amqp#Delivery
type AMQPSinkRecord struct {
	Body                 []byte   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Exchange             string   `protobuf:"bytes,4,opt,name=exchange,proto3" json:"exchange,omitempty"`
	RoutingKey           string   `protobuf:"bytes,5,opt,name=routing_key,json=routingKey,proto3" json:"routing_key,omitempty"`
	ContentType          string   `protobuf:"bytes,6,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	ContentEncoding      string   `protobuf:"bytes,7,opt,name=content_encoding,json=contentEncoding,proto3" json:"content_encoding,omitempty"`
	Priority             int32    `protobuf:"varint,8,opt,name=priority,proto3" json:"priority,omitempty"`
	Expiration           string   `protobuf:"bytes,9,opt,name=expiration,proto3" json:"expiration,omitempty"`
	MessageId            string   `protobuf:"bytes,10,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	UserId               string   `protobuf:"bytes,11,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AppId                string   `protobuf:"bytes,12,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	ReplyTo              string   `protobuf:"bytes,13,opt,name=reply_to,json=replyTo,proto3" json:"reply_to,omitempty"`
	CorrelationId        string   `protobuf:"bytes,14,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AMQPSinkRecord) Reset()         { *m = AMQPSinkRecord{} }
func (m *AMQPSinkRecord) String() string { return proto.CompactTextString(m) }
func (*AMQPSinkRecord) ProtoMessage()    {}
func (*AMQPSinkRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_f84ad538798beb46, []int{0}
}

func (m *AMQPSinkRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AMQPSinkRecord.Unmarshal(m, b)
}
func (m *AMQPSinkRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AMQPSinkRecord.Marshal(b, m, deterministic)
}
func (m *AMQPSinkRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AMQPSinkRecord.Merge(m, src)
}
func (m *AMQPSinkRecord) XXX_Size() int {
	return xxx_messageInfo_AMQPSinkRecord.Size(m)
}
func (m *AMQPSinkRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_AMQPSinkRecord.DiscardUnknown(m)
}

var xxx_messageInfo_AMQPSinkRecord proto.InternalMessageInfo

func (m *AMQPSinkRecord) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *AMQPSinkRecord) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *AMQPSinkRecord) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AMQPSinkRecord) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *AMQPSinkRecord) GetRoutingKey() string {
	if m != nil {
		return m.RoutingKey
	}
	return ""
}

func (m *AMQPSinkRecord) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *AMQPSinkRecord) GetContentEncoding() string {
	if m != nil {
		return m.ContentEncoding
	}
	return ""
}

func (m *AMQPSinkRecord) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *AMQPSinkRecord) GetExpiration() string {
	if m != nil {
		return m.Expiration
	}
	return ""
}

func (m *AMQPSinkRecord) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *AMQPSinkRecord) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *AMQPSinkRecord) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *AMQPSinkRecord) GetReplyTo() string {
	if m != nil {
		return m.ReplyTo
	}
	return ""
}

func (m *AMQPSinkRecord) GetCorrelationId() string {
	if m != nil {
		return m.CorrelationId
	}
	return ""
}

func init() {
	proto.RegisterType((*AMQPSinkRecord)(nil), "records.AMQPSinkRecord")
}

func init() { proto.RegisterFile("events/records/amqp.proto", fileDescriptor_f84ad538798beb46) }

var fileDescriptor_f84ad538798beb46 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x4f, 0x8b, 0xd4, 0x40,
	0x10, 0xc5, 0x89, 0xb3, 0x93, 0x4c, 0x6a, 0x66, 0x47, 0x69, 0x10, 0x7b, 0xc5, 0x3f, 0x51, 0x10,
	0xe2, 0x65, 0x73, 0x50, 0xbc, 0x2b, 0x78, 0x08, 0x22, 0x68, 0xdc, 0x93, 0x97, 0xd0, 0xe9, 0x2e,
	0x92, 0x66, 0x27, 0xdd, 0x6d, 0x77, 0x47, 0x36, 0x5f, 0xca, 0xcf, 0x28, 0x29, 0xe3, 0xae, 0xde,
	0xea, 0xfd, 0xde, 0xab, 0x3f, 0x50, 0x70, 0x81, 0x3f, 0xd1, 0xc4, 0x50, 0x79, 0x94, 0xd6, 0xab,
	0x50, 0x89, 0xf1, 0x87, 0xbb, 0x74, 0xde, 0x46, 0xcb, 0xb2, 0x95, 0xbd, 0xfc, 0xb5, 0x81, 0xe3,
	0xfb, 0xcf, 0x5f, 0xbf, 0x7c, 0xd3, 0xe6, 0xba, 0x21, 0xc6, 0x18, 0x9c, 0x75, 0x56, 0xcd, 0x3c,
	0x29, 0x92, 0xf2, 0xd0, 0x50, 0xcd, 0x9e, 0x40, 0x1e, 0xf5, 0x88, 0x21, 0x8a, 0xd1, 0xf1, 0x7b,
	0x45, 0x52, 0x6e, 0x9a, 0x3b, 0xb0, 0x74, 0xc4, 0xd9, 0x21, 0xdf, 0x14, 0x49, 0x99, 0x37, 0x54,
	0xb3, 0xc7, 0xb0, 0xc3, 0x1b, 0x39, 0x08, 0xd3, 0x23, 0x3f, 0x23, 0x7e, 0xab, 0xd9, 0x73, 0xd8,
	0x7b, 0x3b, 0x45, 0x6d, 0xfa, 0xf6, 0x1a, 0x67, 0xbe, 0x25, 0x1b, 0x56, 0xf4, 0x09, 0x67, 0xf6,
	0x02, 0x0e, 0xd2, 0x9a, 0x88, 0x26, 0xb6, 0x34, 0x38, 0xa5, 0xc4, 0x7e, 0x65, 0x57, 0xcb, 0xfc,
	0xd7, 0xf0, 0xe0, 0x6f, 0x04, 0x8d, 0xb4, 0x4a, 0x9b, 0x9e, 0x67, 0x14, 0xbb, 0xbf, 0xf2, 0x8f,
	0x2b, 0x5e, 0x4e, 0x71, 0x5e, 0x5b, 0xaf, 0xe3, 0xcc, 0x77, 0x45, 0x52, 0x6e, 0x9b, 0x5b, 0xcd,
	0x9e, 0x01, 0xe0, 0x8d, 0xd3, 0x5e, 0x44, 0x6d, 0x0d, 0xcf, 0xff, 0x5c, 0x72, 0x47, 0xd8, 0x53,
	0x80, 0x11, 0x43, 0x10, 0x3d, 0xb6, 0x5a, 0x71, 0x20, 0x3f, 0x5f, 0x49, 0xad, 0xd8, 0x23, 0xc8,
	0xa6, 0x80, 0x7e, 0xf1, 0xf6, 0xe4, 0xa5, 0x8b, 0xac, 0x15, 0x7b, 0x08, 0xa9, 0x70, 0x6e, 0xe1,
	0x07, 0xe2, 0x5b, 0xe1, 0x5c, 0xad, 0xd8, 0x05, 0xec, 0x3c, 0xba, 0xd3, 0xdc, 0x46, 0xcb, 0xcf,
	0xc9, 0xc8, 0x48, 0x5f, 0x59, 0xf6, 0x0a, 0x8e, 0xd2, 0x7a, 0x8f, 0x27, 0x5a, 0xbc, 0x74, 0x1e,
	0x29, 0x70, 0xfe, 0x0f, 0xad, 0xd5, 0x87, 0x77, 0xdf, 0xdf, 0xf6, 0x3a, 0x0e, 0x53, 0x77, 0x29,
	0xed, 0x58, 0x75, 0x22, 0xca, 0x41, 0x5a, 0xef, 0xaa, 0x20, 0x07, 0x1c, 0x45, 0xa8, 0xba, 0x49,
	0x9f, 0x54, 0xd5, 0xdb, 0xea, 0xff, 0xe7, 0x77, 0x29, 0x3d, 0xfe, 0xcd, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x82, 0x1e, 0xed, 0xd3, 0x15, 0x02, 0x00, 0x00,
}
