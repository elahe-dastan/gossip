// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gossip.proto

package protocol

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type ID struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_878fa4887b90140c, []int{0}
}

func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (m *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(m, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Data struct {
	Id                   *ID      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_878fa4887b90140c, []int{1}
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

func (m *Data) GetId() *ID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Data) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*ID)(nil), "protocol.ID")
	proto.RegisterType((*Data)(nil), "protocol.Data")
}

func init() {
	proto.RegisterFile("gossip.proto", fileDescriptor_878fa4887b90140c)
}

var fileDescriptor_878fa4887b90140c = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xcf, 0x2f, 0x2e,
	0xce, 0x2c, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39, 0x52,
	0xd2, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x60, 0x81, 0xa4, 0xd2, 0x34, 0xfd, 0xd4, 0xdc,
	0x82, 0x92, 0x4a, 0x88, 0x32, 0x25, 0x11, 0x2e, 0x26, 0x4f, 0x17, 0x21, 0x3e, 0x2e, 0xa6, 0xcc,
	0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0xa6, 0xcc, 0x14, 0x25, 0x0b, 0x2e, 0x16, 0x97,
	0xc4, 0x92, 0x44, 0x21, 0x19, 0xb8, 0x38, 0xb7, 0x11, 0x8f, 0x1e, 0xcc, 0x44, 0x3d, 0x4f, 0x17,
	0x90, 0x2a, 0x21, 0x21, 0x2e, 0x96, 0x92, 0xd4, 0x8a, 0x12, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x30, 0xdb, 0x68, 0x0a, 0x23, 0x17, 0x8b, 0x73, 0x46, 0x62, 0x89, 0x90, 0x01, 0x17, 0x4b,
	0x70, 0x6a, 0x5e, 0x8a, 0x10, 0x1f, 0x42, 0x1b, 0xc8, 0x48, 0x29, 0x31, 0x3d, 0x88, 0x73, 0xf4,
	0x60, 0xce, 0xd1, 0x73, 0x05, 0x39, 0x47, 0x89, 0x41, 0x48, 0x9b, 0x8b, 0x3d, 0x28, 0x35, 0x39,
	0x35, 0xb3, 0x2c, 0x55, 0x08, 0xc5, 0x2e, 0x29, 0x34, 0x23, 0x94, 0x18, 0x0c, 0x18, 0x85, 0x74,
	0xb9, 0x98, 0xc3, 0x33, 0xf2, 0x85, 0x70, 0x98, 0x26, 0x85, 0x62, 0x80, 0x12, 0x43, 0x12, 0x1b,
	0x98, 0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xf6, 0xac, 0xa9, 0x24, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatClient interface {
	Send(ctx context.Context, in *Data, opts ...grpc.CallOption) (*empty.Empty, error)
	Receive(ctx context.Context, in *ID, opts ...grpc.CallOption) (Chat_ReceiveClient, error)
	Who(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ID, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) Send(ctx context.Context, in *Data, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/protocol.Chat/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) Receive(ctx context.Context, in *ID, opts ...grpc.CallOption) (Chat_ReceiveClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chat_serviceDesc.Streams[0], "/protocol.Chat/Receive", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatReceiveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chat_ReceiveClient interface {
	Recv() (*Data, error)
	grpc.ClientStream
}

type chatReceiveClient struct {
	grpc.ClientStream
}

func (x *chatReceiveClient) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatClient) Who(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/protocol.Chat/Who", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServer is the server API for Chat service.
type ChatServer interface {
	Send(context.Context, *Data) (*empty.Empty, error)
	Receive(*ID, Chat_ReceiveServer) error
	Who(context.Context, *empty.Empty) (*ID, error)
}

// UnimplementedChatServer can be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (*UnimplementedChatServer) Send(ctx context.Context, req *Data) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (*UnimplementedChatServer) Receive(req *ID, srv Chat_ReceiveServer) error {
	return status.Errorf(codes.Unimplemented, "method Receive not implemented")
}
func (*UnimplementedChatServer) Who(ctx context.Context, req *empty.Empty) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Who not implemented")
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Chat/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Send(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_Receive_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ID)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServer).Receive(m, &chatReceiveServer{stream})
}

type Chat_ReceiveServer interface {
	Send(*Data) error
	grpc.ServerStream
}

type chatReceiveServer struct {
	grpc.ServerStream
}

func (x *chatReceiveServer) Send(m *Data) error {
	return x.ServerStream.SendMsg(m)
}

func _Chat_Who_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Who(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Chat/Who",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Who(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Chat_Send_Handler,
		},
		{
			MethodName: "Who",
			Handler:    _Chat_Who_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Receive",
			Handler:       _Chat_Receive_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gossip.proto",
}
