// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: link.proto

package link

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Message struct {
	Context Context `protobuf:"bytes,1,opt,name=Context,proto3" json:"Context"`
	Content []byte  `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ee656911eb8a56a, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

type Context struct {
	ID          uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TS          uint64 `protobuf:"varint,2,opt,name=TS,proto3" json:"TS,omitempty"`
	QOS         uint32 `protobuf:"varint,3,opt,name=QOS,proto3" json:"QOS,omitempty"`
	Flags       uint32 `protobuf:"varint,4,opt,name=Flags,proto3" json:"Flags,omitempty"`
	Topic       string `protobuf:"bytes,5,opt,name=Topic,proto3" json:"Topic,omitempty"`
	Source      string `protobuf:"bytes,6,opt,name=Source,proto3" json:"Source,omitempty"`
	Destination string `protobuf:"bytes,7,opt,name=Destination,proto3" json:"Destination,omitempty"`
}

func (m *Context) Reset()         { *m = Context{} }
func (m *Context) String() string { return proto.CompactTextString(m) }
func (*Context) ProtoMessage()    {}
func (*Context) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ee656911eb8a56a, []int{1}
}
func (m *Context) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Context) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Context.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Context) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Context.Merge(m, src)
}
func (m *Context) XXX_Size() int {
	return m.Size()
}
func (m *Context) XXX_DiscardUnknown() {
	xxx_messageInfo_Context.DiscardUnknown(m)
}

var xxx_messageInfo_Context proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Message)(nil), "link.Message")
	proto.RegisterType((*Context)(nil), "link.Context")
}

func init() { proto.RegisterFile("link.proto", fileDescriptor_2ee656911eb8a56a) }

var fileDescriptor_2ee656911eb8a56a = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4a, 0xc3, 0x40,
	0x14, 0x86, 0xf3, 0xd2, 0x69, 0x8b, 0xaf, 0x56, 0x64, 0x10, 0x19, 0xba, 0x18, 0x43, 0x17, 0x12,
	0xc4, 0xb6, 0x52, 0x6f, 0xd0, 0x16, 0xa1, 0xa0, 0x88, 0x93, 0xac, 0xdc, 0xa5, 0x25, 0xc6, 0xd0,
	0x98, 0x29, 0xcd, 0x14, 0x3c, 0x86, 0x57, 0x70, 0xe7, 0x11, 0x5c, 0xba, 0xec, 0xb2, 0x4b, 0x57,
	0x62, 0xa7, 0x97, 0x70, 0x29, 0x99, 0xa6, 0xa2, 0x2b, 0x77, 0xff, 0xf7, 0xbf, 0xff, 0xfd, 0x30,
	0xf3, 0x10, 0x93, 0x38, 0x9d, 0xb4, 0xa7, 0x33, 0xa9, 0x24, 0x25, 0xb9, 0x6e, 0xb4, 0xa2, 0x58,
	0xdd, 0xcf, 0x47, 0xed, 0xb1, 0x7c, 0xe8, 0x44, 0x32, 0x92, 0x1d, 0x33, 0x1c, 0xcd, 0xef, 0x0c,
	0x19, 0x30, 0x6a, 0xb3, 0xd4, 0x14, 0x58, 0xbd, 0x0a, 0xb3, 0x2c, 0x88, 0x42, 0xda, 0xc2, 0x6a,
	0x5f, 0xa6, 0x2a, 0x7c, 0x54, 0x0c, 0x1c, 0x70, 0x6b, 0xdd, 0x7a, 0xdb, 0xb4, 0x17, 0x66, 0x8f,
	0x2c, 0x3e, 0x8e, 0x2c, 0xb1, 0xcd, 0x50, 0x56, 0xc4, 0x53, 0xc5, 0x6c, 0x07, 0xdc, 0x5d, 0xb1,
	0xc5, 0xe6, 0x33, 0xfc, 0x34, 0xd1, 0x3d, 0xb4, 0x87, 0x03, 0xd3, 0x47, 0x84, 0x3d, 0x1c, 0xe4,
	0xec, 0x7b, 0x66, 0x81, 0x08, 0xdb, 0xf7, 0xe8, 0x3e, 0x96, 0x6e, 0xae, 0x3d, 0x56, 0x72, 0xc0,
	0xad, 0x8b, 0x5c, 0xd2, 0x03, 0x2c, 0x5f, 0x24, 0x41, 0x94, 0x31, 0x62, 0xbc, 0x0d, 0xe4, 0xae,
	0x2f, 0xa7, 0xf1, 0x98, 0x95, 0x1d, 0x70, 0x77, 0xc4, 0x06, 0xe8, 0x21, 0x56, 0x3c, 0x39, 0x9f,
	0x8d, 0x43, 0x56, 0x31, 0x76, 0x41, 0xd4, 0xc1, 0xda, 0x20, 0xcc, 0x54, 0x9c, 0x06, 0x2a, 0x96,
	0x29, 0xab, 0x9a, 0xe1, 0x6f, 0xab, 0x7b, 0x8b, 0xe4, 0x32, 0x4e, 0x27, 0xf4, 0x04, 0x89, 0x1f,
	0x24, 0x13, 0x5a, 0xbc, 0xb5, 0xf8, 0x8b, 0xc6, 0x5f, 0x6c, 0x5a, 0x2e, 0x9c, 0x01, 0x3d, 0x46,
	0xd2, 0x0f, 0x92, 0xe4, 0xbf, 0x6c, 0xef, 0x74, 0xb1, 0xe2, 0xd6, 0xd7, 0x8a, 0xc3, 0x8b, 0xe6,
	0xf0, 0xaa, 0x39, 0xbc, 0x69, 0x0e, 0x0b, 0xcd, 0x61, 0xa9, 0x39, 0x7c, 0x6a, 0x0e, 0x4f, 0x6b,
	0x6e, 0x2d, 0xd7, 0xdc, 0x7a, 0x5f, 0x73, 0x6b, 0x54, 0x31, 0x87, 0x38, 0xff, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0x5a, 0x05, 0x3c, 0x3f, 0xcb, 0x01, 0x00, 0x00,
}

func (this *Message) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Message)
	if !ok {
		that2, ok := that.(Message)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Context.Equal(&that1.Context) {
		return false
	}
	if !bytes.Equal(this.Content, that1.Content) {
		return false
	}
	return true
}
func (this *Context) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Context)
	if !ok {
		that2, ok := that.(Context)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ID != that1.ID {
		return false
	}
	if this.TS != that1.TS {
		return false
	}
	if this.QOS != that1.QOS {
		return false
	}
	if this.Flags != that1.Flags {
		return false
	}
	if this.Topic != that1.Topic {
		return false
	}
	if this.Source != that1.Source {
		return false
	}
	if this.Destination != that1.Destination {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LinkClient is the client API for Link service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LinkClient interface {
	Talk(ctx context.Context, opts ...grpc.CallOption) (Link_TalkClient, error)
	Call(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type linkClient struct {
	cc *grpc.ClientConn
}

func NewLinkClient(cc *grpc.ClientConn) LinkClient {
	return &linkClient{cc}
}

func (c *linkClient) Talk(ctx context.Context, opts ...grpc.CallOption) (Link_TalkClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Link_serviceDesc.Streams[0], "/link.Link/Talk", opts...)
	if err != nil {
		return nil, err
	}
	x := &linkTalkClient{stream}
	return x, nil
}

type Link_TalkClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type linkTalkClient struct {
	grpc.ClientStream
}

func (x *linkTalkClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *linkTalkClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *linkClient) Call(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/link.Link/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinkServer is the server API for Link service.
type LinkServer interface {
	Talk(Link_TalkServer) error
	Call(context.Context, *Message) (*Message, error)
}

// UnimplementedLinkServer can be embedded to have forward compatible implementations.
type UnimplementedLinkServer struct {
}

func (*UnimplementedLinkServer) Talk(srv Link_TalkServer) error {
	return status.Errorf(codes.Unimplemented, "method Talk not implemented")
}
func (*UnimplementedLinkServer) Call(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}

func RegisterLinkServer(s *grpc.Server, srv LinkServer) {
	s.RegisterService(&_Link_serviceDesc, srv)
}

func _Link_Talk_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LinkServer).Talk(&linkTalkServer{stream})
}

type Link_TalkServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type linkTalkServer struct {
	grpc.ServerStream
}

func (x *linkTalkServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *linkTalkServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Link_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/link.Link/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkServer).Call(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Link_serviceDesc = grpc.ServiceDesc{
	ServiceName: "link.Link",
	HandlerType: (*LinkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Link_Call_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Talk",
			Handler:       _Link_Talk_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "link.proto",
}

func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintLink(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Context.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLink(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Context) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Context) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Context) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Destination) > 0 {
		i -= len(m.Destination)
		copy(dAtA[i:], m.Destination)
		i = encodeVarintLink(dAtA, i, uint64(len(m.Destination)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Source) > 0 {
		i -= len(m.Source)
		copy(dAtA[i:], m.Source)
		i = encodeVarintLink(dAtA, i, uint64(len(m.Source)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Topic) > 0 {
		i -= len(m.Topic)
		copy(dAtA[i:], m.Topic)
		i = encodeVarintLink(dAtA, i, uint64(len(m.Topic)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Flags != 0 {
		i = encodeVarintLink(dAtA, i, uint64(m.Flags))
		i--
		dAtA[i] = 0x20
	}
	if m.QOS != 0 {
		i = encodeVarintLink(dAtA, i, uint64(m.QOS))
		i--
		dAtA[i] = 0x18
	}
	if m.TS != 0 {
		i = encodeVarintLink(dAtA, i, uint64(m.TS))
		i--
		dAtA[i] = 0x10
	}
	if m.ID != 0 {
		i = encodeVarintLink(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLink(dAtA []byte, offset int, v uint64) int {
	offset -= sovLink(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedMessage(r randyLink, easy bool) *Message {
	this := &Message{}
	v1 := NewPopulatedContext(r, easy)
	this.Context = *v1
	v2 := r.Intn(100)
	this.Content = make([]byte, v2)
	for i := 0; i < v2; i++ {
		this.Content[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedContext(r randyLink, easy bool) *Context {
	this := &Context{}
	this.ID = uint64(uint64(r.Uint32()))
	this.TS = uint64(uint64(r.Uint32()))
	this.QOS = uint32(r.Uint32())
	this.Flags = uint32(r.Uint32())
	this.Topic = string(randStringLink(r))
	this.Source = string(randStringLink(r))
	this.Destination = string(randStringLink(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyLink interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneLink(r randyLink) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringLink(r randyLink) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneLink(r)
	}
	return string(tmps)
}
func randUnrecognizedLink(r randyLink, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldLink(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldLink(dAtA []byte, r randyLink, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateLink(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateLink(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateLink(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateLink(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateLink(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateLink(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateLink(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Context.Size()
	n += 1 + l + sovLink(uint64(l))
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovLink(uint64(l))
	}
	return n
}

func (m *Context) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovLink(uint64(m.ID))
	}
	if m.TS != 0 {
		n += 1 + sovLink(uint64(m.TS))
	}
	if m.QOS != 0 {
		n += 1 + sovLink(uint64(m.QOS))
	}
	if m.Flags != 0 {
		n += 1 + sovLink(uint64(m.Flags))
	}
	l = len(m.Topic)
	if l > 0 {
		n += 1 + l + sovLink(uint64(l))
	}
	l = len(m.Source)
	if l > 0 {
		n += 1 + l + sovLink(uint64(l))
	}
	l = len(m.Destination)
	if l > 0 {
		n += 1 + l + sovLink(uint64(l))
	}
	return n
}

func sovLink(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLink(x uint64) (n int) {
	return sovLink(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLink
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
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Context", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLink
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
				return ErrInvalidLengthLink
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLink
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Context.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLink
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
				return ErrInvalidLengthLink
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLink
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = append(m.Content[:0], dAtA[iNdEx:postIndex]...)
			if m.Content == nil {
				m.Content = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLink(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLink
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLink
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
func (m *Context) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLink
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
			return fmt.Errorf("proto: Context: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Context: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLink
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TS", wireType)
			}
			m.TS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLink
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TS |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field QOS", wireType)
			}
			m.QOS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLink
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.QOS |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flags", wireType)
			}
			m.Flags = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLink
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Flags |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLink
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
				return ErrInvalidLengthLink
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLink
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Topic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLink
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
				return ErrInvalidLengthLink
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLink
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Source = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Destination", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLink
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
				return ErrInvalidLengthLink
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLink
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Destination = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLink(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLink
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLink
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
func skipLink(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLink
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
					return 0, ErrIntOverflowLink
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
					return 0, ErrIntOverflowLink
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
				return 0, ErrInvalidLengthLink
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLink
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLink
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLink        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLink          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLink = fmt.Errorf("proto: unexpected end of group")
)
