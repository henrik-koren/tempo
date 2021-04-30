// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/tempopb/tempo.proto

package tempopb

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	v1 "github.com/grafana/tempo/pkg/tempopb/trace/v1"
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

type TraceByIDRequest struct {
	TraceID    []byte `protobuf:"bytes,1,opt,name=traceID,proto3" json:"traceID,omitempty"`
	BlockStart string `protobuf:"bytes,2,opt,name=blockStart,proto3" json:"blockStart,omitempty"`
	BlockEnd   string `protobuf:"bytes,3,opt,name=blockEnd,proto3" json:"blockEnd,omitempty"`
	QueryMode  string `protobuf:"bytes,5,opt,name=queryMode,proto3" json:"queryMode,omitempty"`
}

func (m *TraceByIDRequest) Reset()         { *m = TraceByIDRequest{} }
func (m *TraceByIDRequest) String() string { return proto.CompactTextString(m) }
func (*TraceByIDRequest) ProtoMessage()    {}
func (*TraceByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f22805646f4f62b6, []int{0}
}
func (m *TraceByIDRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TraceByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TraceByIDRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TraceByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceByIDRequest.Merge(m, src)
}
func (m *TraceByIDRequest) XXX_Size() int {
	return m.Size()
}
func (m *TraceByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TraceByIDRequest proto.InternalMessageInfo

func (m *TraceByIDRequest) GetTraceID() []byte {
	if m != nil {
		return m.TraceID
	}
	return nil
}

func (m *TraceByIDRequest) GetBlockStart() string {
	if m != nil {
		return m.BlockStart
	}
	return ""
}

func (m *TraceByIDRequest) GetBlockEnd() string {
	if m != nil {
		return m.BlockEnd
	}
	return ""
}

func (m *TraceByIDRequest) GetQueryMode() string {
	if m != nil {
		return m.QueryMode
	}
	return ""
}

type TraceByIDResponse struct {
	Trace *Trace `protobuf:"bytes,1,opt,name=trace,proto3" json:"trace,omitempty"`
}

func (m *TraceByIDResponse) Reset()         { *m = TraceByIDResponse{} }
func (m *TraceByIDResponse) String() string { return proto.CompactTextString(m) }
func (*TraceByIDResponse) ProtoMessage()    {}
func (*TraceByIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f22805646f4f62b6, []int{1}
}
func (m *TraceByIDResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TraceByIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TraceByIDResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TraceByIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceByIDResponse.Merge(m, src)
}
func (m *TraceByIDResponse) XXX_Size() int {
	return m.Size()
}
func (m *TraceByIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceByIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TraceByIDResponse proto.InternalMessageInfo

func (m *TraceByIDResponse) GetTrace() *Trace {
	if m != nil {
		return m.Trace
	}
	return nil
}

type Trace struct {
	Batches []*v1.ResourceSpans `protobuf:"bytes,1,rep,name=batches,proto3" json:"batches,omitempty"`
}

func (m *Trace) Reset()         { *m = Trace{} }
func (m *Trace) String() string { return proto.CompactTextString(m) }
func (*Trace) ProtoMessage()    {}
func (*Trace) Descriptor() ([]byte, []int) {
	return fileDescriptor_f22805646f4f62b6, []int{2}
}
func (m *Trace) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Trace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Trace.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Trace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trace.Merge(m, src)
}
func (m *Trace) XXX_Size() int {
	return m.Size()
}
func (m *Trace) XXX_DiscardUnknown() {
	xxx_messageInfo_Trace.DiscardUnknown(m)
}

var xxx_messageInfo_Trace proto.InternalMessageInfo

func (m *Trace) GetBatches() []*v1.ResourceSpans {
	if m != nil {
		return m.Batches
	}
	return nil
}

type PushRequest struct {
	Batch *v1.ResourceSpans `protobuf:"bytes,1,opt,name=batch,proto3" json:"batch,omitempty"`
}

func (m *PushRequest) Reset()         { *m = PushRequest{} }
func (m *PushRequest) String() string { return proto.CompactTextString(m) }
func (*PushRequest) ProtoMessage()    {}
func (*PushRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f22805646f4f62b6, []int{3}
}
func (m *PushRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PushRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PushRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PushRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushRequest.Merge(m, src)
}
func (m *PushRequest) XXX_Size() int {
	return m.Size()
}
func (m *PushRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PushRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PushRequest proto.InternalMessageInfo

func (m *PushRequest) GetBatch() *v1.ResourceSpans {
	if m != nil {
		return m.Batch
	}
	return nil
}

type PushResponse struct {
}

func (m *PushResponse) Reset()         { *m = PushResponse{} }
func (m *PushResponse) String() string { return proto.CompactTextString(m) }
func (*PushResponse) ProtoMessage()    {}
func (*PushResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f22805646f4f62b6, []int{4}
}
func (m *PushResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PushResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PushResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PushResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushResponse.Merge(m, src)
}
func (m *PushResponse) XXX_Size() int {
	return m.Size()
}
func (m *PushResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PushResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PushResponse proto.InternalMessageInfo

type PushBytesRequest struct {
	// pre-serialized PushRequests
	Requests []PreallocRequest `protobuf:"bytes,1,rep,name=requests,proto3,customtype=PreallocRequest" json:"requests"`
}

func (m *PushBytesRequest) Reset()         { *m = PushBytesRequest{} }
func (m *PushBytesRequest) String() string { return proto.CompactTextString(m) }
func (*PushBytesRequest) ProtoMessage()    {}
func (*PushBytesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f22805646f4f62b6, []int{5}
}
func (m *PushBytesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PushBytesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PushBytesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PushBytesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushBytesRequest.Merge(m, src)
}
func (m *PushBytesRequest) XXX_Size() int {
	return m.Size()
}
func (m *PushBytesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PushBytesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PushBytesRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TraceByIDRequest)(nil), "tempopb.TraceByIDRequest")
	proto.RegisterType((*TraceByIDResponse)(nil), "tempopb.TraceByIDResponse")
	proto.RegisterType((*Trace)(nil), "tempopb.Trace")
	proto.RegisterType((*PushRequest)(nil), "tempopb.PushRequest")
	proto.RegisterType((*PushResponse)(nil), "tempopb.PushResponse")
	proto.RegisterType((*PushBytesRequest)(nil), "tempopb.PushBytesRequest")
}

func init() { proto.RegisterFile("pkg/tempopb/tempo.proto", fileDescriptor_f22805646f4f62b6) }

var fileDescriptor_f22805646f4f62b6 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0x8e, 0x19, 0x59, 0xd6, 0xb7, 0x65, 0x0c, 0x6b, 0x68, 0x21, 0x42, 0x69, 0x15, 0x71, 0xe8,
	0x85, 0x44, 0xcb, 0xb4, 0xc3, 0x4e, 0x48, 0x51, 0xf9, 0xd8, 0x01, 0x69, 0xa4, 0xfc, 0x81, 0x24,
	0x35, 0x69, 0xb5, 0x2e, 0xce, 0x6c, 0xa7, 0x52, 0x6f, 0x9c, 0x38, 0xf3, 0xb3, 0x7a, 0xec, 0x11,
	0x71, 0xa8, 0x50, 0xfb, 0x47, 0x50, 0xec, 0x24, 0x84, 0x0a, 0xb4, 0x53, 0xde, 0xe7, 0xc3, 0x8f,
	0x9f, 0xd8, 0x86, 0xb3, 0xfc, 0x36, 0xf5, 0x04, 0xb9, 0xcb, 0x69, 0x1e, 0xab, 0xaf, 0x9b, 0x33,
	0x2a, 0x28, 0x36, 0x2a, 0xd2, 0x3a, 0x15, 0x2c, 0x4a, 0x88, 0xb7, 0x38, 0xf7, 0xe4, 0xa0, 0x64,
	0xeb, 0x75, 0x3a, 0x13, 0xd3, 0x22, 0x76, 0x13, 0x7a, 0xe7, 0xa5, 0x34, 0xa5, 0x9e, 0xa4, 0xe3,
	0xe2, 0x8b, 0x44, 0x12, 0xc8, 0x49, 0xd9, 0x9d, 0x6f, 0x08, 0x4e, 0x3e, 0x97, 0xcb, 0x83, 0xe5,
	0xf5, 0x28, 0x24, 0xf7, 0x05, 0xe1, 0x02, 0x9b, 0x60, 0xc8, 0xc8, 0xeb, 0x91, 0x89, 0x06, 0x68,
	0xd8, 0x0b, 0x6b, 0x88, 0x6d, 0x80, 0x78, 0x4e, 0x93, 0xdb, 0xb1, 0x88, 0x98, 0x30, 0x1f, 0x0d,
	0xd0, 0xb0, 0x13, 0xb6, 0x18, 0x6c, 0xc1, 0x91, 0x44, 0x6f, 0xb3, 0x89, 0x79, 0x20, 0xd5, 0x06,
	0xe3, 0x97, 0xd0, 0xb9, 0x2f, 0x08, 0x5b, 0x7e, 0xa4, 0x13, 0x62, 0xea, 0x52, 0xfc, 0x43, 0x38,
	0x57, 0xf0, 0xac, 0xd5, 0x83, 0xe7, 0x34, 0xe3, 0x04, 0xbf, 0x02, 0x5d, 0xee, 0x2c, 0x6b, 0x74,
	0xfd, 0x63, 0xb7, 0xfa, 0x77, 0x57, 0x5a, 0x43, 0x25, 0x3a, 0x01, 0xe8, 0x12, 0xe3, 0x2b, 0x30,
	0xe2, 0x48, 0x24, 0x53, 0xc2, 0x4d, 0x34, 0x38, 0x18, 0x76, 0xfd, 0x7e, 0xb3, 0x40, 0x1d, 0xd1,
	0xe2, 0xdc, 0x0d, 0x09, 0xa7, 0x05, 0x4b, 0xc8, 0x38, 0x8f, 0x32, 0x1e, 0xd6, 0x7e, 0x67, 0x04,
	0xdd, 0x9b, 0x82, 0x4f, 0xeb, 0x13, 0xb8, 0x04, 0x5d, 0x2a, 0xd5, 0xc6, 0x0f, 0xe6, 0x28, 0xb7,
	0x73, 0x0c, 0x3d, 0x95, 0xa2, 0xfa, 0x3b, 0xef, 0xe1, 0xa4, 0xc4, 0xc1, 0x52, 0x10, 0x5e, 0x47,
	0x5f, 0xc0, 0x11, 0x53, 0xa3, 0x6a, 0xd9, 0x0b, 0xce, 0x56, 0x9b, 0xbe, 0xf6, 0x73, 0xd3, 0x7f,
	0x7a, 0xc3, 0x48, 0x34, 0x9f, 0xd3, 0xa4, 0xb2, 0x86, 0x8d, 0xd1, 0xff, 0x8a, 0xe0, 0xb0, 0x4c,
	0x22, 0x0c, 0x5f, 0xc2, 0xe3, 0x72, 0xc2, 0xa7, 0x4d, 0xa7, 0x56, 0x71, 0xeb, 0xf9, 0x1e, 0x5b,
	0x15, 0xd1, 0xf0, 0x1b, 0xe8, 0x34, 0x55, 0xf0, 0x8b, 0xbf, 0x5c, 0xed, 0x7a, 0xff, 0x0d, 0xf0,
	0xc7, 0x60, 0x7c, 0x2a, 0x08, 0x9b, 0x11, 0x86, 0x3f, 0xc0, 0x93, 0x77, 0xb3, 0x6c, 0xd2, 0xdc,
	0x57, 0x2b, 0x6f, 0xff, 0x2d, 0x59, 0xd6, 0xbf, 0xa4, 0x3a, 0x34, 0x30, 0x57, 0x5b, 0x1b, 0xad,
	0xb7, 0x36, 0xfa, 0xb5, 0xb5, 0xd1, 0xf7, 0x9d, 0xad, 0xad, 0x77, 0xb6, 0xf6, 0x63, 0x67, 0x6b,
	0xf1, 0xa1, 0x7c, 0x9f, 0x17, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x97, 0x5f, 0xaf, 0xe4, 0x08,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PusherClient is the client API for Pusher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PusherClient interface {
	Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResponse, error)
	PushBytes(ctx context.Context, in *PushBytesRequest, opts ...grpc.CallOption) (*PushResponse, error)
}

type pusherClient struct {
	cc *grpc.ClientConn
}

func NewPusherClient(cc *grpc.ClientConn) PusherClient {
	return &pusherClient{cc}
}

func (c *pusherClient) Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResponse, error) {
	out := new(PushResponse)
	err := c.cc.Invoke(ctx, "/tempopb.Pusher/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pusherClient) PushBytes(ctx context.Context, in *PushBytesRequest, opts ...grpc.CallOption) (*PushResponse, error) {
	out := new(PushResponse)
	err := c.cc.Invoke(ctx, "/tempopb.Pusher/PushBytes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PusherServer is the server API for Pusher service.
type PusherServer interface {
	Push(context.Context, *PushRequest) (*PushResponse, error)
	PushBytes(context.Context, *PushBytesRequest) (*PushResponse, error)
}

// UnimplementedPusherServer can be embedded to have forward compatible implementations.
type UnimplementedPusherServer struct {
}

func (*UnimplementedPusherServer) Push(ctx context.Context, req *PushRequest) (*PushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (*UnimplementedPusherServer) PushBytes(ctx context.Context, req *PushBytesRequest) (*PushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushBytes not implemented")
}

func RegisterPusherServer(s *grpc.Server, srv PusherServer) {
	s.RegisterService(&_Pusher_serviceDesc, srv)
}

func _Pusher_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PusherServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tempopb.Pusher/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PusherServer).Push(ctx, req.(*PushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pusher_PushBytes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushBytesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PusherServer).PushBytes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tempopb.Pusher/PushBytes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PusherServer).PushBytes(ctx, req.(*PushBytesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Pusher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tempopb.Pusher",
	HandlerType: (*PusherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Push",
			Handler:    _Pusher_Push_Handler,
		},
		{
			MethodName: "PushBytes",
			Handler:    _Pusher_PushBytes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/tempopb/tempo.proto",
}

// QuerierClient is the client API for Querier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QuerierClient interface {
	FindTraceByID(ctx context.Context, in *TraceByIDRequest, opts ...grpc.CallOption) (*TraceByIDResponse, error)
}

type querierClient struct {
	cc *grpc.ClientConn
}

func NewQuerierClient(cc *grpc.ClientConn) QuerierClient {
	return &querierClient{cc}
}

func (c *querierClient) FindTraceByID(ctx context.Context, in *TraceByIDRequest, opts ...grpc.CallOption) (*TraceByIDResponse, error) {
	out := new(TraceByIDResponse)
	err := c.cc.Invoke(ctx, "/tempopb.Querier/FindTraceByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuerierServer is the server API for Querier service.
type QuerierServer interface {
	FindTraceByID(context.Context, *TraceByIDRequest) (*TraceByIDResponse, error)
}

// UnimplementedQuerierServer can be embedded to have forward compatible implementations.
type UnimplementedQuerierServer struct {
}

func (*UnimplementedQuerierServer) FindTraceByID(ctx context.Context, req *TraceByIDRequest) (*TraceByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindTraceByID not implemented")
}

func RegisterQuerierServer(s *grpc.Server, srv QuerierServer) {
	s.RegisterService(&_Querier_serviceDesc, srv)
}

func _Querier_FindTraceByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TraceByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuerierServer).FindTraceByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tempopb.Querier/FindTraceByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuerierServer).FindTraceByID(ctx, req.(*TraceByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Querier_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tempopb.Querier",
	HandlerType: (*QuerierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindTraceByID",
			Handler:    _Querier_FindTraceByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/tempopb/tempo.proto",
}

func (m *TraceByIDRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TraceByIDRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TraceByIDRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.QueryMode) > 0 {
		i -= len(m.QueryMode)
		copy(dAtA[i:], m.QueryMode)
		i = encodeVarintTempo(dAtA, i, uint64(len(m.QueryMode)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.BlockEnd) > 0 {
		i -= len(m.BlockEnd)
		copy(dAtA[i:], m.BlockEnd)
		i = encodeVarintTempo(dAtA, i, uint64(len(m.BlockEnd)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BlockStart) > 0 {
		i -= len(m.BlockStart)
		copy(dAtA[i:], m.BlockStart)
		i = encodeVarintTempo(dAtA, i, uint64(len(m.BlockStart)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TraceID) > 0 {
		i -= len(m.TraceID)
		copy(dAtA[i:], m.TraceID)
		i = encodeVarintTempo(dAtA, i, uint64(len(m.TraceID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TraceByIDResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TraceByIDResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TraceByIDResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Trace != nil {
		{
			size, err := m.Trace.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTempo(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Trace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Trace) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Trace) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Batches) > 0 {
		for iNdEx := len(m.Batches) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Batches[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTempo(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *PushRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PushRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PushRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Batch != nil {
		{
			size, err := m.Batch.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTempo(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PushResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PushResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PushResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *PushBytesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PushBytesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PushBytesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Requests) > 0 {
		for iNdEx := len(m.Requests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Requests[iNdEx].Size()
				i -= size
				if _, err := m.Requests[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintTempo(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTempo(dAtA []byte, offset int, v uint64) int {
	offset -= sovTempo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TraceByIDRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TraceID)
	if l > 0 {
		n += 1 + l + sovTempo(uint64(l))
	}
	l = len(m.BlockStart)
	if l > 0 {
		n += 1 + l + sovTempo(uint64(l))
	}
	l = len(m.BlockEnd)
	if l > 0 {
		n += 1 + l + sovTempo(uint64(l))
	}
	l = len(m.QueryMode)
	if l > 0 {
		n += 1 + l + sovTempo(uint64(l))
	}
	return n
}

func (m *TraceByIDResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Trace != nil {
		l = m.Trace.Size()
		n += 1 + l + sovTempo(uint64(l))
	}
	return n
}

func (m *Trace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Batches) > 0 {
		for _, e := range m.Batches {
			l = e.Size()
			n += 1 + l + sovTempo(uint64(l))
		}
	}
	return n
}

func (m *PushRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Batch != nil {
		l = m.Batch.Size()
		n += 1 + l + sovTempo(uint64(l))
	}
	return n
}

func (m *PushResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *PushBytesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Requests) > 0 {
		for _, e := range m.Requests {
			l = e.Size()
			n += 1 + l + sovTempo(uint64(l))
		}
	}
	return n
}

func sovTempo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTempo(x uint64) (n int) {
	return sovTempo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TraceByIDRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTempo
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
			return fmt.Errorf("proto: TraceByIDRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TraceByIDRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TraceID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTempo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTempo
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTempo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TraceID = append(m.TraceID[:0], dAtA[iNdEx:postIndex]...)
			if m.TraceID == nil {
				m.TraceID = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockStart", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTempo
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
				return ErrInvalidLengthTempo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTempo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockStart = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockEnd", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTempo
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
				return ErrInvalidLengthTempo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTempo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockEnd = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryMode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTempo
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
				return ErrInvalidLengthTempo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTempo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryMode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTempo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTempo
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
func (m *TraceByIDResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTempo
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
			return fmt.Errorf("proto: TraceByIDResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TraceByIDResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trace", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTempo
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
				return ErrInvalidLengthTempo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTempo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Trace == nil {
				m.Trace = &Trace{}
			}
			if err := m.Trace.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTempo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTempo
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
func (m *Trace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTempo
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
			return fmt.Errorf("proto: Trace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Trace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Batches", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTempo
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
				return ErrInvalidLengthTempo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTempo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Batches = append(m.Batches, &v1.ResourceSpans{})
			if err := m.Batches[len(m.Batches)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTempo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTempo
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
func (m *PushRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTempo
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
			return fmt.Errorf("proto: PushRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PushRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Batch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTempo
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
				return ErrInvalidLengthTempo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTempo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Batch == nil {
				m.Batch = &v1.ResourceSpans{}
			}
			if err := m.Batch.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTempo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTempo
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
func (m *PushResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTempo
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
			return fmt.Errorf("proto: PushResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PushResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTempo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTempo
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
func (m *PushBytesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTempo
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
			return fmt.Errorf("proto: PushBytesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PushBytesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requests", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTempo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTempo
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTempo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v PreallocRequest
			m.Requests = append(m.Requests, v)
			if err := m.Requests[len(m.Requests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTempo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTempo
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
func skipTempo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTempo
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
					return 0, ErrIntOverflowTempo
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
					return 0, ErrIntOverflowTempo
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
				return 0, ErrInvalidLengthTempo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTempo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTempo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTempo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTempo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTempo = fmt.Errorf("proto: unexpected end of group")
)
