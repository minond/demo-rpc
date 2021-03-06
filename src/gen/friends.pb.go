// Code generated by protoc-gen-go. DO NOT EDIT.
// source: friends.proto

/*
Package gen is a generated protocol buffer package.

It is generated from these files:
	friends.proto

It has these top-level messages:
	Friend
	Ack
	SearchRequest
*/
package gen

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Friend struct {
	Id   uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Friend) Reset()                    { *m = Friend{} }
func (m *Friend) String() string            { return proto.CompactTextString(m) }
func (*Friend) ProtoMessage()               {}
func (*Friend) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Friend) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Friend) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Ack struct {
	Ok  bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *Ack) Reset()                    { *m = Ack{} }
func (m *Ack) String() string            { return proto.CompactTextString(m) }
func (*Ack) ProtoMessage()               {}
func (*Ack) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ack) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *Ack) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type SearchRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SearchRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Friend)(nil), "gen.Friend")
	proto.RegisterType((*Ack)(nil), "gen.Ack")
	proto.RegisterType((*SearchRequest)(nil), "gen.SearchRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Friends service

type FriendsClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (Friends_SearchClient, error)
	Create(ctx context.Context, in *Friend, opts ...grpc.CallOption) (*Ack, error)
}

type friendsClient struct {
	cc *grpc.ClientConn
}

func NewFriendsClient(cc *grpc.ClientConn) FriendsClient {
	return &friendsClient{cc}
}

func (c *friendsClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (Friends_SearchClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Friends_serviceDesc.Streams[0], c.cc, "/gen.Friends/Search", opts...)
	if err != nil {
		return nil, err
	}
	x := &friendsSearchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Friends_SearchClient interface {
	Recv() (*Friend, error)
	grpc.ClientStream
}

type friendsSearchClient struct {
	grpc.ClientStream
}

func (x *friendsSearchClient) Recv() (*Friend, error) {
	m := new(Friend)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *friendsClient) Create(ctx context.Context, in *Friend, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := grpc.Invoke(ctx, "/gen.Friends/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Friends service

type FriendsServer interface {
	Search(*SearchRequest, Friends_SearchServer) error
	Create(context.Context, *Friend) (*Ack, error)
}

func RegisterFriendsServer(s *grpc.Server, srv FriendsServer) {
	s.RegisterService(&_Friends_serviceDesc, srv)
}

func _Friends_Search_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FriendsServer).Search(m, &friendsSearchServer{stream})
}

type Friends_SearchServer interface {
	Send(*Friend) error
	grpc.ServerStream
}

type friendsSearchServer struct {
	grpc.ServerStream
}

func (x *friendsSearchServer) Send(m *Friend) error {
	return x.ServerStream.SendMsg(m)
}

func _Friends_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Friend)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gen.Friends/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendsServer).Create(ctx, req.(*Friend))
	}
	return interceptor(ctx, in, info, handler)
}

var _Friends_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gen.Friends",
	HandlerType: (*FriendsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Friends_Create_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Search",
			Handler:       _Friends_Search_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "friends.proto",
}

func init() { proto.RegisterFile("friends.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2b, 0xca, 0x4c,
	0xcd, 0x4b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x4f, 0xcd, 0x53, 0xd2,
	0xe1, 0x62, 0x73, 0x03, 0x8b, 0x0a, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a,
	0xf0, 0x06, 0x31, 0x65, 0xa6, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x29,
	0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0xea, 0x5c, 0xcc, 0x8e, 0xc9, 0xd9, 0x20, 0xa5, 0xf9,
	0xd9, 0x60, 0xa5, 0x1c, 0x41, 0x4c, 0xf9, 0xd9, 0x42, 0x02, 0x5c, 0xcc, 0xb9, 0xc5, 0xe9, 0x50,
	0x95, 0x20, 0xa6, 0x92, 0x32, 0x17, 0x6f, 0x70, 0x6a, 0x62, 0x51, 0x72, 0x46, 0x50, 0x6a, 0x61,
	0x69, 0x6a, 0x71, 0x09, 0xdc, 0x34, 0x46, 0x84, 0x69, 0x46, 0xd1, 0x5c, 0xec, 0x10, 0xbb, 0x8b,
	0x85, 0x74, 0xb9, 0xd8, 0x20, 0xea, 0x85, 0x84, 0xf4, 0xd2, 0x53, 0xf3, 0xf4, 0x50, 0x34, 0x4b,
	0x71, 0x83, 0xc5, 0x20, 0x6a, 0x95, 0x18, 0x0c, 0x18, 0x85, 0x14, 0xb9, 0xd8, 0x9c, 0x8b, 0x52,
	0x13, 0x4b, 0x52, 0x85, 0x90, 0xa5, 0xa4, 0x38, 0xc0, 0x1c, 0xc7, 0xe4, 0x6c, 0x25, 0x86, 0x24,
	0x36, 0xb0, 0x27, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xc4, 0xff, 0xd3, 0xf5, 0x00,
	0x00, 0x00,
}
