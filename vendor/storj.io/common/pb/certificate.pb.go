// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: certificate.proto

package pb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"

	drpc "storj.io/drpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SigningRequest struct {
	AuthToken            string   `protobuf:"bytes,1,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SigningRequest) Reset()         { *m = SigningRequest{} }
func (m *SigningRequest) String() string { return proto.CompactTextString(m) }
func (*SigningRequest) ProtoMessage()    {}
func (*SigningRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0d34c34dd33be4b, []int{0}
}
func (m *SigningRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SigningRequest.Unmarshal(m, b)
}
func (m *SigningRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SigningRequest.Marshal(b, m, deterministic)
}
func (m *SigningRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SigningRequest.Merge(m, src)
}
func (m *SigningRequest) XXX_Size() int {
	return xxx_messageInfo_SigningRequest.Size(m)
}
func (m *SigningRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SigningRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SigningRequest proto.InternalMessageInfo

func (m *SigningRequest) GetAuthToken() string {
	if m != nil {
		return m.AuthToken
	}
	return ""
}

func (m *SigningRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type SigningResponse struct {
	Chain                [][]byte `protobuf:"bytes,1,rep,name=chain,proto3" json:"chain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SigningResponse) Reset()         { *m = SigningResponse{} }
func (m *SigningResponse) String() string { return proto.CompactTextString(m) }
func (*SigningResponse) ProtoMessage()    {}
func (*SigningResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0d34c34dd33be4b, []int{1}
}
func (m *SigningResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SigningResponse.Unmarshal(m, b)
}
func (m *SigningResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SigningResponse.Marshal(b, m, deterministic)
}
func (m *SigningResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SigningResponse.Merge(m, src)
}
func (m *SigningResponse) XXX_Size() int {
	return xxx_messageInfo_SigningResponse.Size(m)
}
func (m *SigningResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SigningResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SigningResponse proto.InternalMessageInfo

func (m *SigningResponse) GetChain() [][]byte {
	if m != nil {
		return m.Chain
	}
	return nil
}

func init() {
	proto.RegisterType((*SigningRequest)(nil), "node.SigningRequest")
	proto.RegisterType((*SigningResponse)(nil), "node.SigningResponse")
}

func init() { proto.RegisterFile("certificate.proto", fileDescriptor_c0d34c34dd33be4b) }

var fileDescriptor_c0d34c34dd33be4b = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0xc1, 0x4a, 0xc6, 0x30,
	0x10, 0x06, 0xa9, 0xff, 0xaf, 0xd0, 0xa5, 0x28, 0x86, 0x0a, 0x45, 0x14, 0x4a, 0x2f, 0xf6, 0x94,
	0x82, 0x7d, 0x03, 0x7b, 0xf6, 0x12, 0x3d, 0x79, 0x91, 0x34, 0xae, 0x6d, 0x94, 0x64, 0x63, 0xb3,
	0x7d, 0x7f, 0x89, 0x05, 0x8b, 0x1e, 0x77, 0x60, 0x87, 0xf9, 0xe0, 0xd2, 0xe0, 0xc2, 0xf6, 0xdd,
	0x1a, 0xcd, 0x28, 0xc3, 0x42, 0x4c, 0xe2, 0xe8, 0xe9, 0x0d, 0x9b, 0x47, 0x38, 0x7f, 0xb2, 0x93,
	0xb7, 0x7e, 0x52, 0xf8, 0xb5, 0x62, 0x64, 0x71, 0x0b, 0xa0, 0x57, 0x9e, 0x5f, 0x99, 0x3e, 0xd1,
	0x57, 0x59, 0x9d, 0xb5, 0xb9, 0xca, 0x13, 0x79, 0x4e, 0x40, 0xdc, 0x40, 0xce, 0xd6, 0x61, 0x64,
	0xed, 0x42, 0x75, 0x52, 0x67, 0xed, 0x41, 0xed, 0xa0, 0xb9, 0x83, 0x8b, 0x5f, 0x5d, 0x0c, 0xe4,
	0x23, 0x8a, 0x12, 0x4e, 0xcd, 0xac, 0x6d, 0x52, 0x1d, 0xda, 0x42, 0x6d, 0xc7, 0xfd, 0x00, 0xc5,
	0xb0, 0x27, 0x45, 0xd1, 0xc3, 0x31, 0x3d, 0x8a, 0x52, 0xa6, 0x2c, 0xf9, 0xb7, 0xe9, 0xfa, 0xea,
	0x1f, 0xdd, 0xd4, 0x0f, 0xe5, 0x8b, 0x88, 0x4c, 0xcb, 0x87, 0xb4, 0xd4, 0x19, 0x72, 0x8e, 0x7c,
	0x17, 0xc6, 0xf1, 0xec, 0x67, 0x5f, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x7b, 0xe3, 0xe4, 0xf2,
	0xf4, 0x00, 0x00, 0x00,
}

// --- DRPC BEGIN ---

type DRPCCertificatesClient interface {
	DRPCConn() drpc.Conn

	Sign(ctx context.Context, in *SigningRequest) (*SigningResponse, error)
}

type drpcCertificatesClient struct {
	cc drpc.Conn
}

func NewDRPCCertificatesClient(cc drpc.Conn) DRPCCertificatesClient {
	return &drpcCertificatesClient{cc}
}

func (c *drpcCertificatesClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcCertificatesClient) Sign(ctx context.Context, in *SigningRequest) (*SigningResponse, error) {
	out := new(SigningResponse)
	err := c.cc.Invoke(ctx, "/node.Certificates/Sign", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCCertificatesServer interface {
	Sign(context.Context, *SigningRequest) (*SigningResponse, error)
}

type DRPCCertificatesDescription struct{}

func (DRPCCertificatesDescription) NumMethods() int { return 1 }

func (DRPCCertificatesDescription) Method(n int) (string, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/node.Certificates/Sign",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCCertificatesServer).
					Sign(
						ctx,
						in1.(*SigningRequest),
					)
			}, DRPCCertificatesServer.Sign, true
	default:
		return "", nil, nil, false
	}
}

func DRPCRegisterCertificates(mux drpc.Mux, impl DRPCCertificatesServer) error {
	return mux.Register(impl, DRPCCertificatesDescription{})
}

type DRPCCertificates_SignStream interface {
	drpc.Stream
	SendAndClose(*SigningResponse) error
}

type drpcCertificatesSignStream struct {
	drpc.Stream
}

func (x *drpcCertificatesSignStream) SendAndClose(m *SigningResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

// --- DRPC END ---
