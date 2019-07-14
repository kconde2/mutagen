// Code generated by protoc-gen-go. DO NOT EDIT.
// source: forwarding/endpoint/remote/protocol.proto

package remote

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	forwarding "github.com/havoc-io/mutagen/pkg/forwarding"
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

// InitializeForwardingRequest is the initialization request sent to remote
// forwarding endpoints.
type InitializeForwardingRequest struct {
	// Version is the session version.
	Version forwarding.Version `protobuf:"varint,1,opt,name=version,proto3,enum=forwarding.Version" json:"version,omitempty"`
	// Configuration is the session configuration.
	Configuration *forwarding.Configuration `protobuf:"bytes,2,opt,name=configuration,proto3" json:"configuration,omitempty"`
	// Protocol is the protocol specification for the listener/dialer.
	Protocol string `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// Address it the bind address for a listener or dial address for a dialer.
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	// Listener indicates whether this endpoint should function as a listener or
	// dialer for the associated address.
	Listener             bool     `protobuf:"varint,5,opt,name=listener,proto3" json:"listener,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitializeForwardingRequest) Reset()         { *m = InitializeForwardingRequest{} }
func (m *InitializeForwardingRequest) String() string { return proto.CompactTextString(m) }
func (*InitializeForwardingRequest) ProtoMessage()    {}
func (*InitializeForwardingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bb0a3c00a9f3e65, []int{0}
}

func (m *InitializeForwardingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitializeForwardingRequest.Unmarshal(m, b)
}
func (m *InitializeForwardingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitializeForwardingRequest.Marshal(b, m, deterministic)
}
func (m *InitializeForwardingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitializeForwardingRequest.Merge(m, src)
}
func (m *InitializeForwardingRequest) XXX_Size() int {
	return xxx_messageInfo_InitializeForwardingRequest.Size(m)
}
func (m *InitializeForwardingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InitializeForwardingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InitializeForwardingRequest proto.InternalMessageInfo

func (m *InitializeForwardingRequest) GetVersion() forwarding.Version {
	if m != nil {
		return m.Version
	}
	return forwarding.Version_Invalid
}

func (m *InitializeForwardingRequest) GetConfiguration() *forwarding.Configuration {
	if m != nil {
		return m.Configuration
	}
	return nil
}

func (m *InitializeForwardingRequest) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *InitializeForwardingRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *InitializeForwardingRequest) GetListener() bool {
	if m != nil {
		return m.Listener
	}
	return false
}

// InitializeForwardingResponse is the initialization response sent from remote
// forwarding endpoint.
type InitializeForwardingResponse struct {
	// Error is any error that occurred during initialization.
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitializeForwardingResponse) Reset()         { *m = InitializeForwardingResponse{} }
func (m *InitializeForwardingResponse) String() string { return proto.CompactTextString(m) }
func (*InitializeForwardingResponse) ProtoMessage()    {}
func (*InitializeForwardingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bb0a3c00a9f3e65, []int{1}
}

func (m *InitializeForwardingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitializeForwardingResponse.Unmarshal(m, b)
}
func (m *InitializeForwardingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitializeForwardingResponse.Marshal(b, m, deterministic)
}
func (m *InitializeForwardingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitializeForwardingResponse.Merge(m, src)
}
func (m *InitializeForwardingResponse) XXX_Size() int {
	return xxx_messageInfo_InitializeForwardingResponse.Size(m)
}
func (m *InitializeForwardingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InitializeForwardingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InitializeForwardingResponse proto.InternalMessageInfo

func (m *InitializeForwardingResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*InitializeForwardingRequest)(nil), "remote.InitializeForwardingRequest")
	proto.RegisterType((*InitializeForwardingResponse)(nil), "remote.InitializeForwardingResponse")
}

func init() {
	proto.RegisterFile("forwarding/endpoint/remote/protocol.proto", fileDescriptor_4bb0a3c00a9f3e65)
}

var fileDescriptor_4bb0a3c00a9f3e65 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4a, 0xf4, 0x30,
	0x14, 0xc5, 0xc9, 0xf7, 0x39, 0xff, 0x22, 0xba, 0x88, 0x2e, 0xe2, 0x28, 0x52, 0x66, 0x55, 0x17,
	0xd3, 0xc0, 0xe8, 0x4a, 0x04, 0x41, 0x41, 0x70, 0x9b, 0x85, 0x0b, 0x77, 0x99, 0xf6, 0x4e, 0x27,
	0x38, 0xcd, 0xad, 0x37, 0xe9, 0x08, 0xbe, 0xac, 0xaf, 0x22, 0xb6, 0xb6, 0x76, 0x40, 0x77, 0x39,
	0x39, 0xbf, 0x93, 0xdc, 0xc3, 0xe5, 0x17, 0x2b, 0xa4, 0x37, 0x43, 0x99, 0x75, 0xb9, 0x02, 0x97,
	0x95, 0x68, 0x5d, 0x50, 0x04, 0x05, 0x06, 0x50, 0x25, 0x61, 0xc0, 0x14, 0x37, 0x49, 0x7d, 0x10,
	0xc3, 0xe6, 0x7a, 0x7a, 0xde, 0x8b, 0xa4, 0xe8, 0x56, 0x36, 0xaf, 0xc8, 0x04, 0x8b, 0xae, 0xe1,
	0xa6, 0xb2, 0xe7, 0x6f, 0x81, 0x7c, 0xe7, 0xcc, 0x3e, 0x18, 0x3f, 0x7d, 0x74, 0x36, 0x58, 0xb3,
	0xb1, 0xef, 0xf0, 0xd0, 0x61, 0x1a, 0x5e, 0x2b, 0xf0, 0x41, 0xcc, 0xf9, 0xe8, 0x3b, 0x20, 0x59,
	0xc4, 0xe2, 0xc3, 0xc5, 0x51, 0xf2, 0xf3, 0x56, 0xf2, 0xd4, 0x58, 0xba, 0x65, 0xc4, 0x2d, 0x3f,
	0xd8, 0xf9, 0x5f, 0xfe, 0x8b, 0x58, 0xbc, 0xbf, 0x38, 0xe9, 0x87, 0xee, 0xfb, 0x80, 0xde, 0xe5,
	0xc5, 0x94, 0x8f, 0xdb, 0x8e, 0xf2, 0x7f, 0xc4, 0xe2, 0x89, 0xee, 0xb4, 0x90, 0x7c, 0x64, 0xb2,
	0x8c, 0xc0, 0x7b, 0xb9, 0x57, 0x5b, 0xad, 0xfc, 0x4a, 0x6d, 0xac, 0x0f, 0xe0, 0x80, 0xe4, 0x20,
	0x62, 0xf1, 0x58, 0x77, 0x7a, 0x76, 0xc5, 0xcf, 0x7e, 0x2f, 0xe8, 0x4b, 0x74, 0x1e, 0xc4, 0x31,
	0x1f, 0x00, 0x11, 0x52, 0xdd, 0x6f, 0xa2, 0x1b, 0x71, 0x77, 0xf3, 0x7c, 0x9d, 0xdb, 0xb0, 0xae,
	0x96, 0x49, 0x8a, 0x85, 0x5a, 0x9b, 0x2d, 0xa6, 0x73, 0x8b, 0xaa, 0xa8, 0x82, 0xc9, 0xc1, 0xa9,
	0xf2, 0x25, 0x57, 0x7f, 0xaf, 0x69, 0x39, 0xac, 0x67, 0xbe, 0xfc, 0x0c, 0x00, 0x00, 0xff, 0xff,
	0x21, 0xc8, 0x79, 0x43, 0xcb, 0x01, 0x00, 0x00,
}
