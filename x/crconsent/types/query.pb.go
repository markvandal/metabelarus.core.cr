// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crconsent/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// this line is used by starport scaffolding # 3
type QueryGetRequestRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryGetRequestRequest) Reset()         { *m = QueryGetRequestRequest{} }
func (m *QueryGetRequestRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetRequestRequest) ProtoMessage()    {}
func (*QueryGetRequestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb15e1ace51238a3, []int{0}
}
func (m *QueryGetRequestRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetRequestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetRequestRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetRequestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetRequestRequest.Merge(m, src)
}
func (m *QueryGetRequestRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetRequestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetRequestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetRequestRequest proto.InternalMessageInfo

func (m *QueryGetRequestRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type QueryGetRequestResponse struct {
	Request *Request `protobuf:"bytes,1,opt,name=Request,proto3" json:"Request,omitempty"`
}

func (m *QueryGetRequestResponse) Reset()         { *m = QueryGetRequestResponse{} }
func (m *QueryGetRequestResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetRequestResponse) ProtoMessage()    {}
func (*QueryGetRequestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb15e1ace51238a3, []int{1}
}
func (m *QueryGetRequestResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetRequestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetRequestResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetRequestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetRequestResponse.Merge(m, src)
}
func (m *QueryGetRequestResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetRequestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetRequestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetRequestResponse proto.InternalMessageInfo

func (m *QueryGetRequestResponse) GetRequest() *Request {
	if m != nil {
		return m.Request
	}
	return nil
}

type QueryAllRequestRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllRequestRequest) Reset()         { *m = QueryAllRequestRequest{} }
func (m *QueryAllRequestRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllRequestRequest) ProtoMessage()    {}
func (*QueryAllRequestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb15e1ace51238a3, []int{2}
}
func (m *QueryAllRequestRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllRequestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllRequestRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllRequestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllRequestRequest.Merge(m, src)
}
func (m *QueryAllRequestRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllRequestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllRequestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllRequestRequest proto.InternalMessageInfo

func (m *QueryAllRequestRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllRequestResponse struct {
	Request    []*Request          `protobuf:"bytes,1,rep,name=Request,proto3" json:"Request,omitempty"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllRequestResponse) Reset()         { *m = QueryAllRequestResponse{} }
func (m *QueryAllRequestResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllRequestResponse) ProtoMessage()    {}
func (*QueryAllRequestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb15e1ace51238a3, []int{3}
}
func (m *QueryAllRequestResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllRequestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllRequestResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllRequestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllRequestResponse.Merge(m, src)
}
func (m *QueryAllRequestResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllRequestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllRequestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllRequestResponse proto.InternalMessageInfo

func (m *QueryAllRequestResponse) GetRequest() []*Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *QueryAllRequestResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetRequestRequest)(nil), "metabelarus.mbcorecr.crconsent.QueryGetRequestRequest")
	proto.RegisterType((*QueryGetRequestResponse)(nil), "metabelarus.mbcorecr.crconsent.QueryGetRequestResponse")
	proto.RegisterType((*QueryAllRequestRequest)(nil), "metabelarus.mbcorecr.crconsent.QueryAllRequestRequest")
	proto.RegisterType((*QueryAllRequestResponse)(nil), "metabelarus.mbcorecr.crconsent.QueryAllRequestResponse")
}

func init() { proto.RegisterFile("crconsent/query.proto", fileDescriptor_cb15e1ace51238a3) }

var fileDescriptor_cb15e1ace51238a3 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x4b, 0xeb, 0x30,
	0x1c, 0xc7, 0x97, 0x3e, 0xde, 0x7b, 0xbc, 0x3c, 0xf0, 0x10, 0xd0, 0xc9, 0x90, 0x22, 0x3d, 0xb8,
	0x29, 0x92, 0xb0, 0x29, 0x7a, 0x9e, 0x07, 0x77, 0xf0, 0xa2, 0x3b, 0x8a, 0x07, 0xd3, 0x2e, 0xd4,
	0x42, 0xdb, 0x74, 0x4d, 0x2a, 0x0e, 0xf1, 0xe2, 0x5f, 0x20, 0xf8, 0x37, 0x78, 0x15, 0xff, 0x0c,
	0x8f, 0x03, 0x2f, 0x1e, 0x65, 0xf3, 0xe6, 0x3f, 0x21, 0x6b, 0xd2, 0xad, 0xda, 0xe1, 0xa6, 0xa7,
	0x42, 0xf3, 0xfb, 0x7e, 0xf2, 0xf9, 0x26, 0x2d, 0x5c, 0x74, 0x62, 0x87, 0x87, 0x82, 0x85, 0x92,
	0x74, 0x13, 0x16, 0xf7, 0x70, 0x14, 0x73, 0xc9, 0x91, 0x19, 0x30, 0x49, 0x6d, 0xe6, 0xd3, 0x38,
	0x11, 0x38, 0xb0, 0x1d, 0x1e, 0x33, 0x27, 0xc6, 0xe3, 0xd9, 0xca, 0x8a, 0xcb, 0xb9, 0xeb, 0x33,
	0x42, 0x23, 0x8f, 0xd0, 0x30, 0xe4, 0x92, 0x4a, 0x8f, 0x87, 0x42, 0xa5, 0x2b, 0x1b, 0x0e, 0x17,
	0x01, 0x17, 0xc4, 0xa6, 0x82, 0x29, 0x2c, 0x39, 0xaf, 0xdb, 0x4c, 0xd2, 0x3a, 0x89, 0xa8, 0xeb,
	0x85, 0xe9, 0xb0, 0x9e, 0x2d, 0x4f, 0x04, 0x62, 0xd6, 0x4d, 0x98, 0x90, 0x6a, 0xc1, 0xaa, 0xc1,
	0xa5, 0xa3, 0x51, 0xb4, 0xc5, 0x64, 0x5b, 0x2d, 0xe8, 0x07, 0x5a, 0x80, 0x86, 0xd7, 0x59, 0x06,
	0xab, 0xa0, 0xf6, 0xaf, 0x6d, 0x78, 0x1d, 0xeb, 0x04, 0x96, 0x0b, 0x93, 0x22, 0x1a, 0x41, 0x51,
	0x13, 0xfe, 0xd5, 0xaf, 0xd2, 0xf9, 0xff, 0x8d, 0x2a, 0xfe, 0xba, 0x19, 0xce, 0x08, 0x59, 0xce,
	0x3a, 0xd5, 0x1e, 0x4d, 0xdf, 0xff, 0xe4, 0xb1, 0x0f, 0xe1, 0xa4, 0x8e, 0xe6, 0xaf, 0x61, 0xd5,
	0x1d, 0x8f, 0xba, 0x63, 0x75, 0xa4, 0xba, 0x3b, 0x3e, 0xa4, 0x2e, 0xcb, 0x10, 0xb9, 0xa4, 0x75,
	0x07, 0x74, 0x81, 0xfc, 0x16, 0xd3, 0x0a, 0xfc, 0xfa, 0x49, 0x01, 0xd4, 0xfa, 0xa0, 0x69, 0xe8,
	0x63, 0x98, 0xa5, 0xa9, 0xf6, 0xcf, 0x7b, 0x36, 0xde, 0x0c, 0xf8, 0x3b, 0xf5, 0x44, 0x0f, 0x60,
	0xac, 0x85, 0x76, 0x66, 0x09, 0x4d, 0xbf, 0xc5, 0xca, 0xee, 0xb7, 0x73, 0x4a, 0xc9, 0xda, 0xbe,
	0x7e, 0x7a, 0xbd, 0x35, 0x30, 0xda, 0x24, 0x39, 0x00, 0xc9, 0x00, 0xa4, 0xf0, 0x3d, 0x91, 0x4b,
	0xaf, 0x73, 0x85, 0xee, 0x01, 0x84, 0x9a, 0xd4, 0xf4, 0xfd, 0x39, 0xad, 0x0b, 0x77, 0x3e, 0xa7,
	0x75, 0xf1, 0x22, 0x2d, 0x92, 0x5a, 0xaf, 0xa3, 0xea, 0x9c, 0xd6, 0x7b, 0x07, 0x8f, 0x03, 0x13,
	0xf4, 0x07, 0x26, 0x78, 0x19, 0x98, 0xe0, 0x66, 0x68, 0x96, 0xfa, 0x43, 0xb3, 0xf4, 0x3c, 0x34,
	0x4b, 0xc7, 0x75, 0xd7, 0x93, 0x67, 0x89, 0x8d, 0x1d, 0x1e, 0x4c, 0x87, 0x5d, 0xe4, 0x70, 0xb2,
	0x17, 0x31, 0x61, 0xff, 0x49, 0xff, 0xa9, 0xad, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x21, 0x5b,
	0x78, 0x74, 0xef, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// this line is used by starport scaffolding # 2
	Request(ctx context.Context, in *QueryGetRequestRequest, opts ...grpc.CallOption) (*QueryGetRequestResponse, error)
	RequestAll(ctx context.Context, in *QueryAllRequestRequest, opts ...grpc.CallOption) (*QueryAllRequestResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Request(ctx context.Context, in *QueryGetRequestRequest, opts ...grpc.CallOption) (*QueryGetRequestResponse, error) {
	out := new(QueryGetRequestResponse)
	err := c.cc.Invoke(ctx, "/metabelarus.mbcorecr.crconsent.Query/Request", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) RequestAll(ctx context.Context, in *QueryAllRequestRequest, opts ...grpc.CallOption) (*QueryAllRequestResponse, error) {
	out := new(QueryAllRequestResponse)
	err := c.cc.Invoke(ctx, "/metabelarus.mbcorecr.crconsent.Query/RequestAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// this line is used by starport scaffolding # 2
	Request(context.Context, *QueryGetRequestRequest) (*QueryGetRequestResponse, error)
	RequestAll(context.Context, *QueryAllRequestRequest) (*QueryAllRequestResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Request(ctx context.Context, req *QueryGetRequestRequest) (*QueryGetRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Request not implemented")
}
func (*UnimplementedQueryServer) RequestAll(ctx context.Context, req *QueryAllRequestRequest) (*QueryAllRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metabelarus.mbcorecr.crconsent.Query/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Request(ctx, req.(*QueryGetRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_RequestAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RequestAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metabelarus.mbcorecr.crconsent.Query/RequestAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).RequestAll(ctx, req.(*QueryAllRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "metabelarus.mbcorecr.crconsent.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _Query_Request_Handler,
		},
		{
			MethodName: "RequestAll",
			Handler:    _Query_RequestAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crconsent/query.proto",
}

func (m *QueryGetRequestRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetRequestRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetRequestRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetRequestResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetRequestResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetRequestResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Request != nil {
		{
			size, err := m.Request.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllRequestRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllRequestRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllRequestRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllRequestResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllRequestResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllRequestResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Request) > 0 {
		for iNdEx := len(m.Request) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Request[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetRequestRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetRequestResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Request != nil {
		l = m.Request.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllRequestRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllRequestResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Request) > 0 {
		for _, e := range m.Request {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetRequestRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetRequestRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetRequestRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetRequestResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetRequestResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetRequestResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Request == nil {
				m.Request = &Request{}
			}
			if err := m.Request.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllRequestRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllRequestRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllRequestRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllRequestResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllRequestResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllRequestResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Request = append(m.Request, &Request{})
			if err := m.Request[len(m.Request)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)