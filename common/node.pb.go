// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node.proto

package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Node struct {
	RPCAddr              string   `protobuf:"bytes,1,opt,name=RPCAddr,proto3" json:"RPCAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_8ae738097d8d3879, []int{0}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (dst *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(dst, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetRPCAddr() string {
	if m != nil {
		return m.RPCAddr
	}
	return ""
}

func init() {
	proto.RegisterType((*Node)(nil), "common.Node")
}

func init() { proto.RegisterFile("node.proto", fileDescriptor_node_8ae738097d8d3879) }

var fileDescriptor_node_8ae738097d8d3879 = []byte{
	// 77 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0xcb, 0x4f, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x53, 0x52,
	0xe0, 0x62, 0xf1, 0xcb, 0x4f, 0x49, 0x15, 0x92, 0xe0, 0x62, 0x0f, 0x0a, 0x70, 0x76, 0x4c, 0x49,
	0x29, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x93, 0xd8, 0xc0, 0x1a, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x08, 0x3f, 0x82, 0x28, 0x3e, 0x00, 0x00, 0x00,
}
