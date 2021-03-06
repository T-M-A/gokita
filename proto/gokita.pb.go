// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gokita.proto

package gokita

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

type ListCustomerRequest struct {
	Limit                int64    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page                 int64    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCustomerRequest) Reset()         { *m = ListCustomerRequest{} }
func (m *ListCustomerRequest) String() string { return proto.CompactTextString(m) }
func (*ListCustomerRequest) ProtoMessage()    {}
func (*ListCustomerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_44a041aab893fba1, []int{0}
}

func (m *ListCustomerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCustomerRequest.Unmarshal(m, b)
}
func (m *ListCustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCustomerRequest.Marshal(b, m, deterministic)
}
func (m *ListCustomerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCustomerRequest.Merge(m, src)
}
func (m *ListCustomerRequest) XXX_Size() int {
	return xxx_messageInfo_ListCustomerRequest.Size(m)
}
func (m *ListCustomerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCustomerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCustomerRequest proto.InternalMessageInfo

func (m *ListCustomerRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListCustomerRequest) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

type ListCustomerReply struct {
	Items                []*Customer `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total                int64       `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Page                 int64       `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Err                  string      `protobuf:"bytes,4,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListCustomerReply) Reset()         { *m = ListCustomerReply{} }
func (m *ListCustomerReply) String() string { return proto.CompactTextString(m) }
func (*ListCustomerReply) ProtoMessage()    {}
func (*ListCustomerReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_44a041aab893fba1, []int{1}
}

func (m *ListCustomerReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCustomerReply.Unmarshal(m, b)
}
func (m *ListCustomerReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCustomerReply.Marshal(b, m, deterministic)
}
func (m *ListCustomerReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCustomerReply.Merge(m, src)
}
func (m *ListCustomerReply) XXX_Size() int {
	return xxx_messageInfo_ListCustomerReply.Size(m)
}
func (m *ListCustomerReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCustomerReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListCustomerReply proto.InternalMessageInfo

func (m *ListCustomerReply) GetItems() []*Customer {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ListCustomerReply) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListCustomerReply) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListCustomerReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type Customer struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string   `protobuf:"bytes,6,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,7,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Customer) Reset()         { *m = Customer{} }
func (m *Customer) String() string { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()    {}
func (*Customer) Descriptor() ([]byte, []int) {
	return fileDescriptor_44a041aab893fba1, []int{2}
}

func (m *Customer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Customer.Unmarshal(m, b)
}
func (m *Customer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Customer.Marshal(b, m, deterministic)
}
func (m *Customer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Customer.Merge(m, src)
}
func (m *Customer) XXX_Size() int {
	return xxx_messageInfo_Customer.Size(m)
}
func (m *Customer) XXX_DiscardUnknown() {
	xxx_messageInfo_Customer.DiscardUnknown(m)
}

var xxx_messageInfo_Customer proto.InternalMessageInfo

func (m *Customer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Customer) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Customer) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func init() {
	proto.RegisterType((*ListCustomerRequest)(nil), "gokita.ListCustomerRequest")
	proto.RegisterType((*ListCustomerReply)(nil), "gokita.ListCustomerReply")
	proto.RegisterType((*Customer)(nil), "gokita.Customer")
}

func init() { proto.RegisterFile("gokita.proto", fileDescriptor_44a041aab893fba1) }

var fileDescriptor_44a041aab893fba1 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x25, 0x49, 0x1b, 0x93, 0xb1, 0x48, 0x1d, 0x3d, 0xac, 0x16, 0x21, 0xe4, 0x20, 0x39, 0xf5,
	0x50, 0x3f, 0x40, 0x50, 0xd0, 0x8b, 0x88, 0xe4, 0xe0, 0xc1, 0x8b, 0xac, 0x74, 0x2d, 0x8b, 0xbb,
	0x6e, 0xdc, 0x9d, 0x22, 0xfd, 0x7b, 0xc9, 0x24, 0xa9, 0x0a, 0xb9, 0xcd, 0x7b, 0x6f, 0x78, 0x6f,
	0xe6, 0xc1, 0x6c, 0xe3, 0x3e, 0x34, 0xc9, 0x65, 0xe3, 0x1d, 0x39, 0x4c, 0x3b, 0x54, 0x5e, 0xc3,
	0xc9, 0x83, 0x0e, 0x74, 0xbb, 0x0d, 0xe4, 0xac, 0xf2, 0xb5, 0xfa, 0xda, 0xaa, 0x40, 0x78, 0x0a,
	0x53, 0xa3, 0xad, 0x26, 0x11, 0x15, 0x51, 0x95, 0xd4, 0x1d, 0x40, 0x84, 0x49, 0x23, 0x37, 0x4a,
	0xc4, 0x4c, 0xf2, 0x5c, 0x7e, 0xc3, 0xf1, 0x7f, 0x83, 0xc6, 0xec, 0xf0, 0x12, 0xa6, 0x9a, 0x94,
	0x0d, 0x22, 0x2a, 0x92, 0xea, 0x70, 0x35, 0x5f, 0xf6, 0xd9, 0xfb, 0xad, 0x4e, 0x6e, 0x63, 0xc8,
	0x91, 0x34, 0xbd, 0x63, 0x07, 0xf6, 0x31, 0xc9, 0x6f, 0x0c, 0xce, 0x21, 0x51, 0xde, 0x8b, 0x49,
	0x11, 0x55, 0x79, 0xdd, 0x8e, 0xe5, 0x33, 0x64, 0x83, 0x1d, 0x1e, 0x41, 0xac, 0xd7, 0x7c, 0x6b,
	0x5e, 0xc7, 0x7a, 0x8d, 0x17, 0x00, 0xef, 0xda, 0x07, 0x7a, 0xfd, 0x94, 0x56, 0x89, 0x94, 0xf9,
	0x9c, 0x99, 0x47, 0x69, 0x15, 0x2e, 0x20, 0x37, 0x72, 0x50, 0x0f, 0x58, 0xcd, 0x5a, 0xa2, 0x15,
	0x57, 0x4f, 0x90, 0xde, 0xf3, 0xb5, 0x78, 0x07, 0xb3, 0xbf, 0xaf, 0xe1, 0x62, 0x78, 0x63, 0xa4,
	0xb1, 0xf3, 0xb3, 0x71, 0xb1, 0x31, 0xbb, 0x9b, 0xec, 0xa5, 0x6f, 0xfb, 0x2d, 0xe5, 0xf2, 0xaf,
	0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xec, 0x27, 0x73, 0x85, 0x8c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GokitaClient is the client API for Gokita service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GokitaClient interface {
	ListCustomer(ctx context.Context, in *ListCustomerRequest, opts ...grpc.CallOption) (*ListCustomerReply, error)
}

type gokitaClient struct {
	cc grpc.ClientConnInterface
}

func NewGokitaClient(cc grpc.ClientConnInterface) GokitaClient {
	return &gokitaClient{cc}
}

func (c *gokitaClient) ListCustomer(ctx context.Context, in *ListCustomerRequest, opts ...grpc.CallOption) (*ListCustomerReply, error) {
	out := new(ListCustomerReply)
	err := c.cc.Invoke(ctx, "/gokita.Gokita/ListCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GokitaServer is the server API for Gokita service.
type GokitaServer interface {
	ListCustomer(context.Context, *ListCustomerRequest) (*ListCustomerReply, error)
}

// UnimplementedGokitaServer can be embedded to have forward compatible implementations.
type UnimplementedGokitaServer struct {
}

func (*UnimplementedGokitaServer) ListCustomer(ctx context.Context, req *ListCustomerRequest) (*ListCustomerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCustomer not implemented")
}

func RegisterGokitaServer(s *grpc.Server, srv GokitaServer) {
	s.RegisterService(&_Gokita_serviceDesc, srv)
}

func _Gokita_ListCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GokitaServer).ListCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gokita.Gokita/ListCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GokitaServer).ListCustomer(ctx, req.(*ListCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gokita_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gokita.Gokita",
	HandlerType: (*GokitaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCustomer",
			Handler:    _Gokita_ListCustomer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gokita.proto",
}
