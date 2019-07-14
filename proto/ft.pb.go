// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/ft.proto

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

type DownloadRequest struct {
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadRequest) Reset()         { *m = DownloadRequest{} }
func (m *DownloadRequest) String() string { return proto.CompactTextString(m) }
func (*DownloadRequest) ProtoMessage()    {}
func (*DownloadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2011610b010dc524, []int{0}
}

func (m *DownloadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadRequest.Unmarshal(m, b)
}
func (m *DownloadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadRequest.Marshal(b, m, deterministic)
}
func (m *DownloadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadRequest.Merge(m, src)
}
func (m *DownloadRequest) XXX_Size() int {
	return xxx_messageInfo_DownloadRequest.Size(m)
}
func (m *DownloadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadRequest proto.InternalMessageInfo

func (m *DownloadRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

type DownloadResponse struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadResponse) Reset()         { *m = DownloadResponse{} }
func (m *DownloadResponse) String() string { return proto.CompactTextString(m) }
func (*DownloadResponse) ProtoMessage()    {}
func (*DownloadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2011610b010dc524, []int{1}
}

func (m *DownloadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadResponse.Unmarshal(m, b)
}
func (m *DownloadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadResponse.Marshal(b, m, deterministic)
}
func (m *DownloadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadResponse.Merge(m, src)
}
func (m *DownloadResponse) XXX_Size() int {
	return xxx_messageInfo_DownloadResponse.Size(m)
}
func (m *DownloadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadResponse proto.InternalMessageInfo

func (m *DownloadResponse) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type UploadRequest struct {
	To                   string   `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	Content              []byte   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadRequest) Reset()         { *m = UploadRequest{} }
func (m *UploadRequest) String() string { return proto.CompactTextString(m) }
func (*UploadRequest) ProtoMessage()    {}
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2011610b010dc524, []int{2}
}

func (m *UploadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadRequest.Unmarshal(m, b)
}
func (m *UploadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadRequest.Marshal(b, m, deterministic)
}
func (m *UploadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadRequest.Merge(m, src)
}
func (m *UploadRequest) XXX_Size() int {
	return xxx_messageInfo_UploadRequest.Size(m)
}
func (m *UploadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadRequest proto.InternalMessageInfo

func (m *UploadRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *UploadRequest) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type UploadResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadResponse) Reset()         { *m = UploadResponse{} }
func (m *UploadResponse) String() string { return proto.CompactTextString(m) }
func (*UploadResponse) ProtoMessage()    {}
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2011610b010dc524, []int{3}
}

func (m *UploadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadResponse.Unmarshal(m, b)
}
func (m *UploadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadResponse.Marshal(b, m, deterministic)
}
func (m *UploadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadResponse.Merge(m, src)
}
func (m *UploadResponse) XXX_Size() int {
	return xxx_messageInfo_UploadResponse.Size(m)
}
func (m *UploadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DownloadRequest)(nil), "DownloadRequest")
	proto.RegisterType((*DownloadResponse)(nil), "DownloadResponse")
	proto.RegisterType((*UploadRequest)(nil), "UploadRequest")
	proto.RegisterType((*UploadResponse)(nil), "UploadResponse")
}

func init() { proto.RegisterFile("proto/ft.proto", fileDescriptor_2011610b010dc524) }

var fileDescriptor_2011610b010dc524 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x2b, 0xd1, 0x03, 0x33, 0x94, 0x54, 0xb9, 0xf8, 0x5d, 0xf2, 0xcb, 0xf3, 0x72,
	0xf2, 0x13, 0x53, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0xd2, 0x8a,
	0xf2, 0x73, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x1d, 0x2e, 0x01, 0x84,
	0xb2, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x09, 0x2e, 0xf6, 0xe4, 0xfc, 0xbc, 0x92, 0xd4,
	0xbc, 0x12, 0xb0, 0x52, 0x9e, 0x20, 0x18, 0x57, 0xc9, 0x92, 0x8b, 0x37, 0xb4, 0x00, 0xd9, 0x48,
	0x3e, 0x2e, 0xa6, 0x92, 0x7c, 0xa8, 0x81, 0x4c, 0x25, 0xf9, 0xc8, 0x5a, 0x99, 0x51, 0xb5, 0x0a,
	0x70, 0xf1, 0xc1, 0xb4, 0x42, 0xac, 0x31, 0x2a, 0xe1, 0xe2, 0x73, 0xcb, 0xcc, 0x49, 0x0d, 0x29,
	0x4a, 0xcc, 0x2b, 0x4e, 0x4b, 0x2d, 0x4a, 0x2d, 0x12, 0x32, 0xe6, 0xe2, 0x80, 0x39, 0x46, 0x48,
	0x40, 0x0f, 0xcd, 0xf9, 0x52, 0x82, 0x7a, 0xe8, 0x2e, 0x55, 0x62, 0x30, 0x60, 0x14, 0xd2, 0xe5,
	0x62, 0x83, 0x18, 0x2c, 0xc4, 0xa7, 0x87, 0xe2, 0x38, 0x29, 0x7e, 0x3d, 0x54, 0x1b, 0x95, 0x18,
	0x34, 0x18, 0x9d, 0xd8, 0xa3, 0x58, 0xc1, 0x01, 0x94, 0xc4, 0x06, 0xa6, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x08, 0x4f, 0x51, 0x86, 0x39, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FileTransfererClient is the client API for FileTransferer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileTransfererClient interface {
	Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (FileTransferer_DownloadClient, error)
	Upload(ctx context.Context, opts ...grpc.CallOption) (FileTransferer_UploadClient, error)
}

type fileTransfererClient struct {
	cc *grpc.ClientConn
}

func NewFileTransfererClient(cc *grpc.ClientConn) FileTransfererClient {
	return &fileTransfererClient{cc}
}

func (c *fileTransfererClient) Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (FileTransferer_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileTransferer_serviceDesc.Streams[0], "/FileTransferer/Download", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileTransfererDownloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileTransferer_DownloadClient interface {
	Recv() (*DownloadResponse, error)
	grpc.ClientStream
}

type fileTransfererDownloadClient struct {
	grpc.ClientStream
}

func (x *fileTransfererDownloadClient) Recv() (*DownloadResponse, error) {
	m := new(DownloadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileTransfererClient) Upload(ctx context.Context, opts ...grpc.CallOption) (FileTransferer_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileTransferer_serviceDesc.Streams[1], "/FileTransferer/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileTransfererUploadClient{stream}
	return x, nil
}

type FileTransferer_UploadClient interface {
	Send(*UploadRequest) error
	CloseAndRecv() (*UploadResponse, error)
	grpc.ClientStream
}

type fileTransfererUploadClient struct {
	grpc.ClientStream
}

func (x *fileTransfererUploadClient) Send(m *UploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileTransfererUploadClient) CloseAndRecv() (*UploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileTransfererServer is the server API for FileTransferer service.
type FileTransfererServer interface {
	Download(*DownloadRequest, FileTransferer_DownloadServer) error
	Upload(FileTransferer_UploadServer) error
}

// UnimplementedFileTransfererServer can be embedded to have forward compatible implementations.
type UnimplementedFileTransfererServer struct {
}

func (*UnimplementedFileTransfererServer) Download(req *DownloadRequest, srv FileTransferer_DownloadServer) error {
	return status.Errorf(codes.Unimplemented, "method Download not implemented")
}
func (*UnimplementedFileTransfererServer) Upload(srv FileTransferer_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}

func RegisterFileTransfererServer(s *grpc.Server, srv FileTransfererServer) {
	s.RegisterService(&_FileTransferer_serviceDesc, srv)
}

func _FileTransferer_Download_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileTransfererServer).Download(m, &fileTransfererDownloadServer{stream})
}

type FileTransferer_DownloadServer interface {
	Send(*DownloadResponse) error
	grpc.ServerStream
}

type fileTransfererDownloadServer struct {
	grpc.ServerStream
}

func (x *fileTransfererDownloadServer) Send(m *DownloadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _FileTransferer_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileTransfererServer).Upload(&fileTransfererUploadServer{stream})
}

type FileTransferer_UploadServer interface {
	SendAndClose(*UploadResponse) error
	Recv() (*UploadRequest, error)
	grpc.ServerStream
}

type fileTransfererUploadServer struct {
	grpc.ServerStream
}

func (x *fileTransfererUploadServer) SendAndClose(m *UploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileTransfererUploadServer) Recv() (*UploadRequest, error) {
	m := new(UploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _FileTransferer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "FileTransferer",
	HandlerType: (*FileTransfererServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Download",
			Handler:       _FileTransferer_Download_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Upload",
			Handler:       _FileTransferer_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/ft.proto",
}
