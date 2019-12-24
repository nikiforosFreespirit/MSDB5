// Code generated by protoc-gen-go. DO NOT EDIT.
// source: points.proto

package pb

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

type PointsRequest struct {
	Number               uint32   `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PointsRequest) Reset()         { *m = PointsRequest{} }
func (m *PointsRequest) String() string { return proto.CompactTextString(m) }
func (*PointsRequest) ProtoMessage()    {}
func (*PointsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eba4b8ba17061946, []int{0}
}

func (m *PointsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PointsRequest.Unmarshal(m, b)
}
func (m *PointsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PointsRequest.Marshal(b, m, deterministic)
}
func (m *PointsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PointsRequest.Merge(m, src)
}
func (m *PointsRequest) XXX_Size() int {
	return xxx_messageInfo_PointsRequest.Size(m)
}
func (m *PointsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PointsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PointsRequest proto.InternalMessageInfo

func (m *PointsRequest) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type PointsResponse struct {
	Points               uint32   `protobuf:"varint,1,opt,name=points,proto3" json:"points,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PointsResponse) Reset()         { *m = PointsResponse{} }
func (m *PointsResponse) String() string { return proto.CompactTextString(m) }
func (*PointsResponse) ProtoMessage()    {}
func (*PointsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eba4b8ba17061946, []int{1}
}

func (m *PointsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PointsResponse.Unmarshal(m, b)
}
func (m *PointsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PointsResponse.Marshal(b, m, deterministic)
}
func (m *PointsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PointsResponse.Merge(m, src)
}
func (m *PointsResponse) XXX_Size() int {
	return xxx_messageInfo_PointsResponse.Size(m)
}
func (m *PointsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PointsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PointsResponse proto.InternalMessageInfo

func (m *PointsResponse) GetPoints() uint32 {
	if m != nil {
		return m.Points
	}
	return 0
}

func init() {
	proto.RegisterType((*PointsRequest)(nil), "pb.PointsRequest")
	proto.RegisterType((*PointsResponse)(nil), "pb.PointsResponse")
}

func init() { proto.RegisterFile("points.proto", fileDescriptor_eba4b8ba17061946) }

var fileDescriptor_eba4b8ba17061946 = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xc8, 0xcf, 0xcc,
	0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x52, 0xe7, 0xe2,
	0x0d, 0x00, 0x8b, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x89, 0x71, 0xb1, 0xe5, 0x95,
	0xe6, 0x26, 0xa5, 0x16, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x06, 0x41, 0x79, 0x4a, 0x1a, 0x5c,
	0x7c, 0x30, 0x85, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x20, 0x95, 0x10, 0xe3, 0x60, 0x2a, 0x21,
	0x3c, 0x23, 0x6b, 0x2e, 0x36, 0x88, 0x4a, 0x21, 0x43, 0x38, 0x4b, 0x50, 0xaf, 0x20, 0x49, 0x0f,
	0xc5, 0x22, 0x29, 0x21, 0x64, 0x21, 0x88, 0x91, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0xa7, 0x19, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x04, 0x55, 0xf6, 0xaa, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PointsClient is the client API for Points service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PointsClient interface {
	Points(ctx context.Context, in *PointsRequest, opts ...grpc.CallOption) (*PointsResponse, error)
}

type pointsClient struct {
	cc *grpc.ClientConn
}

func NewPointsClient(cc *grpc.ClientConn) PointsClient {
	return &pointsClient{cc}
}

func (c *pointsClient) Points(ctx context.Context, in *PointsRequest, opts ...grpc.CallOption) (*PointsResponse, error) {
	out := new(PointsResponse)
	err := c.cc.Invoke(ctx, "/pb.Points/Points", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PointsServer is the server API for Points service.
type PointsServer interface {
	Points(context.Context, *PointsRequest) (*PointsResponse, error)
}

// UnimplementedPointsServer can be embedded to have forward compatible implementations.
type UnimplementedPointsServer struct {
}

func (*UnimplementedPointsServer) Points(ctx context.Context, req *PointsRequest) (*PointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Points not implemented")
}

func RegisterPointsServer(s *grpc.Server, srv PointsServer) {
	s.RegisterService(&_Points_serviceDesc, srv)
}

func _Points_Points_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).Points(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Points/Points",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).Points(ctx, req.(*PointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Points_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Points",
	HandlerType: (*PointsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Points",
			Handler:    _Points_Points_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "points.proto",
}