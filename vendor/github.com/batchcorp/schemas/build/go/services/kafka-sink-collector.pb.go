// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kafka-sink-collector.proto

package services

import (
	context "context"
	fmt "fmt"
	events "github.com/batchcorp/schemas/build/go/events"
	records "github.com/batchcorp/schemas/build/go/events/records"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AddKafkaSinkRecordRequest struct {
	Records              []*records.KafkaSinkRecord `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *AddKafkaSinkRecordRequest) Reset()         { *m = AddKafkaSinkRecordRequest{} }
func (m *AddKafkaSinkRecordRequest) String() string { return proto.CompactTextString(m) }
func (*AddKafkaSinkRecordRequest) ProtoMessage()    {}
func (*AddKafkaSinkRecordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf58ab6229f42719, []int{0}
}

func (m *AddKafkaSinkRecordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddKafkaSinkRecordRequest.Unmarshal(m, b)
}
func (m *AddKafkaSinkRecordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddKafkaSinkRecordRequest.Marshal(b, m, deterministic)
}
func (m *AddKafkaSinkRecordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddKafkaSinkRecordRequest.Merge(m, src)
}
func (m *AddKafkaSinkRecordRequest) XXX_Size() int {
	return xxx_messageInfo_AddKafkaSinkRecordRequest.Size(m)
}
func (m *AddKafkaSinkRecordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddKafkaSinkRecordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddKafkaSinkRecordRequest proto.InternalMessageInfo

func (m *AddKafkaSinkRecordRequest) GetRecords() []*records.KafkaSinkRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

type AddKafkaSinkRecordResponse struct {
	NumRecordsProcessed  int64          `protobuf:"varint,1,opt,name=num_records_processed,json=numRecordsProcessed,proto3" json:"num_records_processed,omitempty"`
	Status               *events.Status `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AddKafkaSinkRecordResponse) Reset()         { *m = AddKafkaSinkRecordResponse{} }
func (m *AddKafkaSinkRecordResponse) String() string { return proto.CompactTextString(m) }
func (*AddKafkaSinkRecordResponse) ProtoMessage()    {}
func (*AddKafkaSinkRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf58ab6229f42719, []int{1}
}

func (m *AddKafkaSinkRecordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddKafkaSinkRecordResponse.Unmarshal(m, b)
}
func (m *AddKafkaSinkRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddKafkaSinkRecordResponse.Marshal(b, m, deterministic)
}
func (m *AddKafkaSinkRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddKafkaSinkRecordResponse.Merge(m, src)
}
func (m *AddKafkaSinkRecordResponse) XXX_Size() int {
	return xxx_messageInfo_AddKafkaSinkRecordResponse.Size(m)
}
func (m *AddKafkaSinkRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddKafkaSinkRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddKafkaSinkRecordResponse proto.InternalMessageInfo

func (m *AddKafkaSinkRecordResponse) GetNumRecordsProcessed() int64 {
	if m != nil {
		return m.NumRecordsProcessed
	}
	return 0
}

func (m *AddKafkaSinkRecordResponse) GetStatus() *events.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*AddKafkaSinkRecordRequest)(nil), "services.AddKafkaSinkRecordRequest")
	proto.RegisterType((*AddKafkaSinkRecordResponse)(nil), "services.AddKafkaSinkRecordResponse")
}

func init() { proto.RegisterFile("kafka-sink-collector.proto", fileDescriptor_cf58ab6229f42719) }

var fileDescriptor_cf58ab6229f42719 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x18, 0xc5, 0xa9, 0x83, 0xa9, 0x19, 0x08, 0x66, 0x08, 0x35, 0xa7, 0x31, 0x45, 0x7a, 0x59, 0x02,
	0xf5, 0x2f, 0x98, 0xe2, 0xc9, 0x83, 0xa3, 0x03, 0x0f, 0x5e, 0x46, 0x9b, 0x7c, 0xae, 0xa5, 0x6d,
	0x52, 0xf3, 0x25, 0xc3, 0x3f, 0x5f, 0x6c, 0xda, 0x1d, 0x44, 0x3d, 0xf6, 0xbd, 0xdf, 0x7b, 0xe5,
	0x7d, 0x21, 0xac, 0xce, 0xdf, 0xeb, 0x7c, 0x85, 0x95, 0xae, 0x57, 0xd2, 0x34, 0x0d, 0x48, 0x67,
	0x2c, 0xef, 0xac, 0x71, 0x86, 0x9e, 0x21, 0xd8, 0x43, 0x25, 0x01, 0xd9, 0x1c, 0x0e, 0xa0, 0x1d,
	0x0a, 0x74, 0xb9, 0xf3, 0x18, 0x6c, 0xc6, 0x06, 0xd1, 0x82, 0x34, 0x56, 0xa1, 0xe8, 0x9b, 0x82,
	0xb7, 0x7c, 0x21, 0xd7, 0x6b, 0xa5, 0x9e, 0xbf, 0x95, 0x6d, 0xa5, 0xeb, 0xac, 0x47, 0x32, 0xf8,
	0xf0, 0x80, 0x8e, 0xa6, 0xe4, 0x74, 0xc8, 0xc4, 0xd1, 0x62, 0x92, 0xcc, 0xd2, 0x98, 0x0f, 0xdf,
	0xfc, 0x67, 0x62, 0x04, 0x97, 0x9f, 0x84, 0xfd, 0x56, 0x88, 0x9d, 0xd1, 0x08, 0x34, 0x25, 0x57,
	0xda, 0xb7, 0xbb, 0x01, 0xde, 0x75, 0xd6, 0x48, 0x40, 0x04, 0x15, 0x47, 0x8b, 0x28, 0x99, 0x64,
	0x73, 0xed, 0xdb, 0x90, 0xc0, 0xcd, 0x68, 0xd1, 0x3b, 0x32, 0x0d, 0x73, 0xe2, 0x93, 0x45, 0x94,
	0xcc, 0xd2, 0x0b, 0x1e, 0xf6, 0xf0, 0x6d, 0xaf, 0x66, 0x83, 0x9b, 0x36, 0x84, 0x1e, 0x7f, 0xfb,
	0x38, 0x5e, 0x88, 0xbe, 0x92, 0xf3, 0xb5, 0x52, 0xa1, 0x94, 0xde, 0xf0, 0xf1, 0x52, 0xfc, 0xcf,
	0xd5, 0xec, 0xf6, 0x7f, 0x28, 0x2c, 0x79, 0x78, 0x22, 0x97, 0x58, 0xf2, 0x22, 0x77, 0xb2, 0x3c,
	0xf2, 0x9b, 0xe8, 0x8d, 0xef, 0x2b, 0x57, 0xfa, 0x82, 0x4b, 0xd3, 0x8a, 0xde, 0x94, 0xc6, 0x76,
	0x02, 0x65, 0x09, 0x6d, 0x8e, 0xa2, 0xf0, 0x55, 0xa3, 0xc4, 0xde, 0x88, 0x31, 0x51, 0x4c, 0xfb,
	0x67, 0xb8, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x4c, 0x11, 0xdc, 0xdf, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KafkaSinkCollectorClient is the client API for KafkaSinkCollector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KafkaSinkCollectorClient interface {
	AddRecord(ctx context.Context, in *AddKafkaSinkRecordRequest, opts ...grpc.CallOption) (*AddKafkaSinkRecordResponse, error)
}

type kafkaSinkCollectorClient struct {
	cc *grpc.ClientConn
}

func NewKafkaSinkCollectorClient(cc *grpc.ClientConn) KafkaSinkCollectorClient {
	return &kafkaSinkCollectorClient{cc}
}

func (c *kafkaSinkCollectorClient) AddRecord(ctx context.Context, in *AddKafkaSinkRecordRequest, opts ...grpc.CallOption) (*AddKafkaSinkRecordResponse, error) {
	out := new(AddKafkaSinkRecordResponse)
	err := c.cc.Invoke(ctx, "/services.KafkaSinkCollector/AddRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KafkaSinkCollectorServer is the server API for KafkaSinkCollector service.
type KafkaSinkCollectorServer interface {
	AddRecord(context.Context, *AddKafkaSinkRecordRequest) (*AddKafkaSinkRecordResponse, error)
}

// UnimplementedKafkaSinkCollectorServer can be embedded to have forward compatible implementations.
type UnimplementedKafkaSinkCollectorServer struct {
}

func (*UnimplementedKafkaSinkCollectorServer) AddRecord(ctx context.Context, req *AddKafkaSinkRecordRequest) (*AddKafkaSinkRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRecord not implemented")
}

func RegisterKafkaSinkCollectorServer(s *grpc.Server, srv KafkaSinkCollectorServer) {
	s.RegisterService(&_KafkaSinkCollector_serviceDesc, srv)
}

func _KafkaSinkCollector_AddRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddKafkaSinkRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaSinkCollectorServer).AddRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.KafkaSinkCollector/AddRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaSinkCollectorServer).AddRecord(ctx, req.(*AddKafkaSinkRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KafkaSinkCollector_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.KafkaSinkCollector",
	HandlerType: (*KafkaSinkCollectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddRecord",
			Handler:    _KafkaSinkCollector_AddRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kafka-sink-collector.proto",
}
