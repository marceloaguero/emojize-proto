// Code generated by protoc-gen-go. DO NOT EDIT.
// source: emoji.proto

package proto

import (
	context "context"
	fmt "fmt"
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

type EmojiRequest struct {
	InputText            string   `protobuf:"bytes,1,opt,name=input_text,json=inputText,proto3" json:"input_text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmojiRequest) Reset()         { *m = EmojiRequest{} }
func (m *EmojiRequest) String() string { return proto.CompactTextString(m) }
func (*EmojiRequest) ProtoMessage()    {}
func (*EmojiRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_890939a4de8c1374, []int{0}
}

func (m *EmojiRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmojiRequest.Unmarshal(m, b)
}
func (m *EmojiRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmojiRequest.Marshal(b, m, deterministic)
}
func (m *EmojiRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmojiRequest.Merge(m, src)
}
func (m *EmojiRequest) XXX_Size() int {
	return xxx_messageInfo_EmojiRequest.Size(m)
}
func (m *EmojiRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmojiRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmojiRequest proto.InternalMessageInfo

func (m *EmojiRequest) GetInputText() string {
	if m != nil {
		return m.InputText
	}
	return ""
}

type EmojiResponse struct {
	OutputText           string   `protobuf:"bytes,1,opt,name=output_text,json=outputText,proto3" json:"output_text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmojiResponse) Reset()         { *m = EmojiResponse{} }
func (m *EmojiResponse) String() string { return proto.CompactTextString(m) }
func (*EmojiResponse) ProtoMessage()    {}
func (*EmojiResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_890939a4de8c1374, []int{1}
}

func (m *EmojiResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmojiResponse.Unmarshal(m, b)
}
func (m *EmojiResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmojiResponse.Marshal(b, m, deterministic)
}
func (m *EmojiResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmojiResponse.Merge(m, src)
}
func (m *EmojiResponse) XXX_Size() int {
	return xxx_messageInfo_EmojiResponse.Size(m)
}
func (m *EmojiResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmojiResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmojiResponse proto.InternalMessageInfo

func (m *EmojiResponse) GetOutputText() string {
	if m != nil {
		return m.OutputText
	}
	return ""
}

func init() {
	proto.RegisterType((*EmojiRequest)(nil), "proto.EmojiRequest")
	proto.RegisterType((*EmojiResponse)(nil), "proto.EmojiResponse")
}

func init() { proto.RegisterFile("emoji.proto", fileDescriptor_890939a4de8c1374) }

var fileDescriptor_890939a4de8c1374 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xcd, 0xcd, 0xcf,
	0xca, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xba, 0x5c, 0x3c, 0xae,
	0x20, 0xd1, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x59, 0x2e, 0xae, 0xcc, 0xbc, 0x82,
	0xd2, 0x92, 0xf8, 0x92, 0xd4, 0x8a, 0x12, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x4e, 0xb0,
	0x48, 0x48, 0x6a, 0x45, 0x89, 0x92, 0x01, 0x17, 0x2f, 0x54, 0x79, 0x71, 0x41, 0x7e, 0x5e, 0x71,
	0xaa, 0x90, 0x3c, 0x17, 0x77, 0x7e, 0x69, 0x09, 0x9a, 0x06, 0x2e, 0x88, 0x10, 0x48, 0x87, 0x91,
	0x27, 0xd4, 0x82, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x4b, 0x2e, 0x1e, 0xcf, 0xbc,
	0xe2, 0xd4, 0xa2, 0x12, 0xb0, 0x68, 0xb1, 0x90, 0x30, 0xc4, 0x3d, 0x7a, 0xc8, 0xae, 0x90, 0x12,
	0x41, 0x15, 0x84, 0xd8, 0x95, 0xc4, 0x06, 0x16, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4d,
	0xcf, 0x74, 0xde, 0xc8, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EmojiServiceClient is the client API for EmojiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmojiServiceClient interface {
	InsertEmojis(ctx context.Context, in *EmojiRequest, opts ...grpc.CallOption) (*EmojiResponse, error)
}

type emojiServiceClient struct {
	cc *grpc.ClientConn
}

func NewEmojiServiceClient(cc *grpc.ClientConn) EmojiServiceClient {
	return &emojiServiceClient{cc}
}

func (c *emojiServiceClient) InsertEmojis(ctx context.Context, in *EmojiRequest, opts ...grpc.CallOption) (*EmojiResponse, error) {
	out := new(EmojiResponse)
	err := c.cc.Invoke(ctx, "/proto.EmojiService/InsertEmojis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmojiServiceServer is the server API for EmojiService service.
type EmojiServiceServer interface {
	InsertEmojis(context.Context, *EmojiRequest) (*EmojiResponse, error)
}

// UnimplementedEmojiServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEmojiServiceServer struct {
}

func (*UnimplementedEmojiServiceServer) InsertEmojis(ctx context.Context, req *EmojiRequest) (*EmojiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertEmojis not implemented")
}

func RegisterEmojiServiceServer(s *grpc.Server, srv EmojiServiceServer) {
	s.RegisterService(&_EmojiService_serviceDesc, srv)
}

func _EmojiService_InsertEmojis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmojiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmojiServiceServer).InsertEmojis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmojiService/InsertEmojis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmojiServiceServer).InsertEmojis(ctx, req.(*EmojiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EmojiService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EmojiService",
	HandlerType: (*EmojiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertEmojis",
			Handler:    _EmojiService_InsertEmojis_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "emoji.proto",
}
