// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cheese.proto

package cheesefarm

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

type CheeseType int32

const (
	CheeseType_EMMENTAL  CheeseType = 0
	CheeseType_BRIE      CheeseType = 1
	CheeseType_PECORINO  CheeseType = 2
	CheeseType_ROQUEFORT CheeseType = 3
	CheeseType_CANASTRA  CheeseType = 4
)

var CheeseType_name = map[int32]string{
	0: "EMMENTAL",
	1: "BRIE",
	2: "PECORINO",
	3: "ROQUEFORT",
	4: "CANASTRA",
}

var CheeseType_value = map[string]int32{
	"EMMENTAL":  0,
	"BRIE":      1,
	"PECORINO":  2,
	"ROQUEFORT": 3,
	"CANASTRA":  4,
}

func (x CheeseType) String() string {
	return proto.EnumName(CheeseType_name, int32(x))
}

func (CheeseType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2ac2cfb993fc6f6e, []int{0}
}

type Cheese struct {
	Age                  int32      `protobuf:"varint,1,opt,name=age,proto3" json:"age,omitempty"`
	Type                 CheeseType `protobuf:"varint,2,opt,name=type,proto3,enum=cheesefarm.CheeseType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Cheese) Reset()         { *m = Cheese{} }
func (m *Cheese) String() string { return proto.CompactTextString(m) }
func (*Cheese) ProtoMessage()    {}
func (*Cheese) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ac2cfb993fc6f6e, []int{0}
}

func (m *Cheese) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cheese.Unmarshal(m, b)
}
func (m *Cheese) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cheese.Marshal(b, m, deterministic)
}
func (m *Cheese) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cheese.Merge(m, src)
}
func (m *Cheese) XXX_Size() int {
	return xxx_messageInfo_Cheese.Size(m)
}
func (m *Cheese) XXX_DiscardUnknown() {
	xxx_messageInfo_Cheese.DiscardUnknown(m)
}

var xxx_messageInfo_Cheese proto.InternalMessageInfo

func (m *Cheese) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Cheese) GetType() CheeseType {
	if m != nil {
		return m.Type
	}
	return CheeseType_EMMENTAL
}

type CheeseRequest struct {
	Type                 CheeseType `protobuf:"varint,1,opt,name=type,proto3,enum=cheesefarm.CheeseType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CheeseRequest) Reset()         { *m = CheeseRequest{} }
func (m *CheeseRequest) String() string { return proto.CompactTextString(m) }
func (*CheeseRequest) ProtoMessage()    {}
func (*CheeseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ac2cfb993fc6f6e, []int{1}
}

func (m *CheeseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheeseRequest.Unmarshal(m, b)
}
func (m *CheeseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheeseRequest.Marshal(b, m, deterministic)
}
func (m *CheeseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheeseRequest.Merge(m, src)
}
func (m *CheeseRequest) XXX_Size() int {
	return xxx_messageInfo_CheeseRequest.Size(m)
}
func (m *CheeseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheeseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheeseRequest proto.InternalMessageInfo

func (m *CheeseRequest) GetType() CheeseType {
	if m != nil {
		return m.Type
	}
	return CheeseType_EMMENTAL
}

func init() {
	proto.RegisterEnum("cheesefarm.CheeseType", CheeseType_name, CheeseType_value)
	proto.RegisterType((*Cheese)(nil), "cheesefarm.Cheese")
	proto.RegisterType((*CheeseRequest)(nil), "cheesefarm.CheeseRequest")
}

func init() { proto.RegisterFile("cheese.proto", fileDescriptor_2ac2cfb993fc6f6e) }

var fileDescriptor_2ac2cfb993fc6f6e = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0x48, 0x4d,
	0x2d, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x82, 0xf0, 0xd2, 0x12, 0x8b, 0x72,
	0x95, 0xdc, 0xb8, 0xd8, 0x9c, 0xc1, 0x3c, 0x21, 0x01, 0x2e, 0xe6, 0xc4, 0xf4, 0x54, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xd6, 0x20, 0x10, 0x53, 0x48, 0x8b, 0x8b, 0xa5, 0xa4, 0xb2, 0x20, 0x55, 0x82,
	0x49, 0x81, 0x51, 0x83, 0xcf, 0x48, 0x4c, 0x0f, 0xa1, 0x4d, 0x0f, 0xa2, 0x27, 0xa4, 0xb2, 0x20,
	0x35, 0x08, 0xac, 0x46, 0xc9, 0x9a, 0x8b, 0x17, 0x22, 0x16, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c,
	0x02, 0xd7, 0xcc, 0x48, 0x58, 0xb3, 0x96, 0x3f, 0x17, 0x17, 0x42, 0x4c, 0x88, 0x87, 0x8b, 0xc3,
	0xd5, 0xd7, 0xd7, 0xd5, 0x2f, 0xc4, 0xd1, 0x47, 0x80, 0x41, 0x88, 0x83, 0x8b, 0xc5, 0x29, 0xc8,
	0xd3, 0x55, 0x80, 0x11, 0x24, 0x1e, 0xe0, 0xea, 0xec, 0x1f, 0xe4, 0xe9, 0xe7, 0x2f, 0xc0, 0x24,
	0xc4, 0xcb, 0xc5, 0x19, 0xe4, 0x1f, 0x18, 0xea, 0xea, 0xe6, 0x1f, 0x14, 0x22, 0xc0, 0x0c, 0x92,
	0x74, 0x76, 0xf4, 0x73, 0x0c, 0x0e, 0x09, 0x72, 0x14, 0x60, 0x31, 0x72, 0x87, 0xb9, 0x26, 0x38,
	0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0xc8, 0x8c, 0x8b, 0xd5, 0xbf, 0x28, 0x25, 0xb5, 0x48, 0x48,
	0x12, 0xd3, 0x21, 0x50, 0x17, 0x4b, 0x09, 0x61, 0x4a, 0x25, 0xb1, 0x81, 0x43, 0xcc, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x88, 0xc3, 0xe3, 0x27, 0x41, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CheeseServiceClient is the client API for CheeseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CheeseServiceClient interface {
	Order(ctx context.Context, in *CheeseRequest, opts ...grpc.CallOption) (*Cheese, error)
}

type cheeseServiceClient struct {
	cc *grpc.ClientConn
}

func NewCheeseServiceClient(cc *grpc.ClientConn) CheeseServiceClient {
	return &cheeseServiceClient{cc}
}

func (c *cheeseServiceClient) Order(ctx context.Context, in *CheeseRequest, opts ...grpc.CallOption) (*Cheese, error) {
	out := new(Cheese)
	err := c.cc.Invoke(ctx, "/cheesefarm.CheeseService/Order", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheeseServiceServer is the server API for CheeseService service.
type CheeseServiceServer interface {
	Order(context.Context, *CheeseRequest) (*Cheese, error)
}

func RegisterCheeseServiceServer(s *grpc.Server, srv CheeseServiceServer) {
	s.RegisterService(&_CheeseService_serviceDesc, srv)
}

func _CheeseService_Order_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheeseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheeseServiceServer).Order(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheesefarm.CheeseService/Order",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheeseServiceServer).Order(ctx, req.(*CheeseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CheeseService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cheesefarm.CheeseService",
	HandlerType: (*CheeseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Order",
			Handler:    _CheeseService_Order_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cheese.proto",
}
