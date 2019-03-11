// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pipe.proto

package pipe

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Record struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1979e1a5fdc07ed, []int{0}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Record) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Record) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Record) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type Ack struct {
	Err                  bool     `protobuf:"varint,1,opt,name=err,proto3" json:"err,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ack) Reset()         { *m = Ack{} }
func (m *Ack) String() string { return proto.CompactTextString(m) }
func (*Ack) ProtoMessage()    {}
func (*Ack) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1979e1a5fdc07ed, []int{1}
}

func (m *Ack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ack.Unmarshal(m, b)
}
func (m *Ack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ack.Marshal(b, m, deterministic)
}
func (m *Ack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ack.Merge(m, src)
}
func (m *Ack) XXX_Size() int {
	return xxx_messageInfo_Ack.Size(m)
}
func (m *Ack) XXX_DiscardUnknown() {
	xxx_messageInfo_Ack.DiscardUnknown(m)
}

var xxx_messageInfo_Ack proto.InternalMessageInfo

func (m *Ack) GetErr() bool {
	if m != nil {
		return m.Err
	}
	return false
}

func (m *Ack) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Record)(nil), "pipe.Record")
	proto.RegisterType((*Ack)(nil), "pipe.Ack")
}

func init() { proto.RegisterFile("pipe.proto", fileDescriptor_d1979e1a5fdc07ed) }

var fileDescriptor_d1979e1a5fdc07ed = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0x31, 0x6f, 0x83, 0x30,
	0x10, 0x46, 0x6b, 0x4c, 0x29, 0x9c, 0xaa, 0xb6, 0x3a, 0x75, 0xb0, 0x3a, 0x21, 0x26, 0x26, 0x44,
	0x93, 0x35, 0x0b, 0x7f, 0xc1, 0x4a, 0xa4, 0xac, 0x0e, 0xbe, 0x24, 0x16, 0x01, 0x5b, 0x86, 0xff,
	0xaf, 0x08, 0x3b, 0x6c, 0xef, 0xbd, 0xe1, 0xf4, 0x1d, 0x80, 0x33, 0x8e, 0x1a, 0xe7, 0xed, 0x62,
	0x31, 0x5d, 0xb9, 0x3a, 0x43, 0x26, 0xa9, 0xb7, 0x5e, 0xe3, 0x17, 0x24, 0x46, 0x0b, 0x56, 0xb2,
	0x9a, 0xcb, 0xc4, 0x68, 0x44, 0x48, 0x27, 0x35, 0x92, 0x48, 0x4a, 0x56, 0x17, 0x32, 0x30, 0xfe,
	0xc2, 0x3b, 0x8d, 0xca, 0x3c, 0x04, 0x0f, 0x31, 0xca, 0x5a, 0xdd, 0xdd, 0x4e, 0x24, 0xd2, 0x58,
	0x83, 0x54, 0xff, 0xc0, 0xbb, 0x7e, 0xc0, 0x1f, 0xe0, 0xe4, 0x7d, 0xb8, 0x9b, 0xcb, 0x15, 0x51,
	0xc0, 0xc7, 0x48, 0xf3, 0xac, 0x6e, 0xdb, 0xed, 0x4d, 0x77, 0x07, 0xc8, 0x8f, 0x5e, 0x4d, 0xf3,
	0x95, 0x3c, 0xb6, 0xf0, 0xdd, 0x69, 0x7d, 0x72, 0x5a, 0x2d, 0xf4, 0x5a, 0xf8, 0xd9, 0x84, 0xf9,
	0xd1, 0xfe, 0x8a, 0x68, 0x5d, 0x3f, 0x54, 0x6f, 0x35, 0x6b, 0xd9, 0x25, 0x0b, 0x7f, 0xed, 0x9f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x91, 0x10, 0xbc, 0x7d, 0xe5, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransferClient is the client API for Transfer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransferClient interface {
	AddUpdateRecord(ctx context.Context, opts ...grpc.CallOption) (Transfer_AddUpdateRecordClient, error)
}

type transferClient struct {
	cc *grpc.ClientConn
}

func NewTransferClient(cc *grpc.ClientConn) TransferClient {
	return &transferClient{cc}
}

func (c *transferClient) AddUpdateRecord(ctx context.Context, opts ...grpc.CallOption) (Transfer_AddUpdateRecordClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Transfer_serviceDesc.Streams[0], "/pipe.Transfer/AddUpdateRecord", opts...)
	if err != nil {
		return nil, err
	}
	x := &transferAddUpdateRecordClient{stream}
	return x, nil
}

type Transfer_AddUpdateRecordClient interface {
	Send(*Record) error
	Recv() (*Ack, error)
	grpc.ClientStream
}

type transferAddUpdateRecordClient struct {
	grpc.ClientStream
}

func (x *transferAddUpdateRecordClient) Send(m *Record) error {
	return x.ClientStream.SendMsg(m)
}

func (x *transferAddUpdateRecordClient) Recv() (*Ack, error) {
	m := new(Ack)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TransferServer is the server API for Transfer service.
type TransferServer interface {
	AddUpdateRecord(Transfer_AddUpdateRecordServer) error
}

func RegisterTransferServer(s *grpc.Server, srv TransferServer) {
	s.RegisterService(&_Transfer_serviceDesc, srv)
}

func _Transfer_AddUpdateRecord_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TransferServer).AddUpdateRecord(&transferAddUpdateRecordServer{stream})
}

type Transfer_AddUpdateRecordServer interface {
	Send(*Ack) error
	Recv() (*Record, error)
	grpc.ServerStream
}

type transferAddUpdateRecordServer struct {
	grpc.ServerStream
}

func (x *transferAddUpdateRecordServer) Send(m *Ack) error {
	return x.ServerStream.SendMsg(m)
}

func (x *transferAddUpdateRecordServer) Recv() (*Record, error) {
	m := new(Record)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Transfer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pipe.Transfer",
	HandlerType: (*TransferServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AddUpdateRecord",
			Handler:       _Transfer_AddUpdateRecord_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pipe.proto",
}