// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lte/protos/keyval.proto

package protos

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// IPs can be in 1 of 4 states during their lifecycle; see
// mobilityd.ip_allocator for full description
type IPDesc_IPState int32

const (
	IPDesc_FREE      IPDesc_IPState = 0
	IPDesc_ALLOCATED IPDesc_IPState = 1
	IPDesc_RELEASED  IPDesc_IPState = 2
	IPDesc_REAPED    IPDesc_IPState = 3
	IPDesc_RESERVED  IPDesc_IPState = 4
)

var IPDesc_IPState_name = map[int32]string{
	0: "FREE",
	1: "ALLOCATED",
	2: "RELEASED",
	3: "REAPED",
	4: "RESERVED",
}

var IPDesc_IPState_value = map[string]int32{
	"FREE":      0,
	"ALLOCATED": 1,
	"RELEASED":  2,
	"REAPED":    3,
	"RESERVED":  4,
}

func (x IPDesc_IPState) String() string {
	return proto.EnumName(IPDesc_IPState_name, int32(x))
}

func (IPDesc_IPState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7ec85b478524e2d6, []int{1, 0}
}

type IPDesc_IPType int32

const (
	IPDesc_STATIC  IPDesc_IPType = 0
	IPDesc_IP_POOL IPDesc_IPType = 1
	IPDesc_DHCP    IPDesc_IPType = 3
)

var IPDesc_IPType_name = map[int32]string{
	0: "STATIC",
	1: "IP_POOL",
	3: "DHCP",
}

var IPDesc_IPType_value = map[string]int32{
	"STATIC":  0,
	"IP_POOL": 1,
	"DHCP":    3,
}

func (x IPDesc_IPType) String() string {
	return proto.EnumName(IPDesc_IPType_name, int32(x))
}

func (IPDesc_IPType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7ec85b478524e2d6, []int{1, 1}
}

// --------------------------------------------------------------------------
// [mobilityd] List of assigned IP blocks
// --------------------------------------------------------------------------
type AssignedIPBlocks struct {
	IpBlockList          []*IPBlock `protobuf:"bytes,1,rep,name=ip_block_list,json=ipBlockList,proto3" json:"ip_block_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AssignedIPBlocks) Reset()         { *m = AssignedIPBlocks{} }
func (m *AssignedIPBlocks) String() string { return proto.CompactTextString(m) }
func (*AssignedIPBlocks) ProtoMessage()    {}
func (*AssignedIPBlocks) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ec85b478524e2d6, []int{0}
}

func (m *AssignedIPBlocks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignedIPBlocks.Unmarshal(m, b)
}
func (m *AssignedIPBlocks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignedIPBlocks.Marshal(b, m, deterministic)
}
func (m *AssignedIPBlocks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignedIPBlocks.Merge(m, src)
}
func (m *AssignedIPBlocks) XXX_Size() int {
	return xxx_messageInfo_AssignedIPBlocks.Size(m)
}
func (m *AssignedIPBlocks) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignedIPBlocks.DiscardUnknown(m)
}

var xxx_messageInfo_AssignedIPBlocks proto.InternalMessageInfo

func (m *AssignedIPBlocks) GetIpBlockList() []*IPBlock {
	if m != nil {
		return m.IpBlockList
	}
	return nil
}

// --------------------------------------------------------------------------
// [mobilityd] IP descriptor (IP desc) describing an assigned IP
// --------------------------------------------------------------------------
type IPDesc struct {
	Ip                   *IPAddress     `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	IpBlock              *IPBlock       `protobuf:"bytes,2,opt,name=ip_block,json=ipBlock,proto3" json:"ip_block,omitempty"`
	State                IPDesc_IPState `protobuf:"varint,3,opt,name=state,proto3,enum=magma.lte.IPDesc_IPState" json:"state,omitempty"`
	Sid                  *SubscriberID  `protobuf:"bytes,4,opt,name=sid,proto3" json:"sid,omitempty"`
	Type                 IPDesc_IPType  `protobuf:"varint,5,opt,name=type,proto3,enum=magma.lte.IPDesc_IPType" json:"type,omitempty"`
	VlanId               uint32         `protobuf:"varint,6,opt,name=vlan_id,json=vlanId,proto3" json:"vlan_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *IPDesc) Reset()         { *m = IPDesc{} }
func (m *IPDesc) String() string { return proto.CompactTextString(m) }
func (*IPDesc) ProtoMessage()    {}
func (*IPDesc) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ec85b478524e2d6, []int{1}
}

func (m *IPDesc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPDesc.Unmarshal(m, b)
}
func (m *IPDesc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPDesc.Marshal(b, m, deterministic)
}
func (m *IPDesc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPDesc.Merge(m, src)
}
func (m *IPDesc) XXX_Size() int {
	return xxx_messageInfo_IPDesc.Size(m)
}
func (m *IPDesc) XXX_DiscardUnknown() {
	xxx_messageInfo_IPDesc.DiscardUnknown(m)
}

var xxx_messageInfo_IPDesc proto.InternalMessageInfo

func (m *IPDesc) GetIp() *IPAddress {
	if m != nil {
		return m.Ip
	}
	return nil
}

func (m *IPDesc) GetIpBlock() *IPBlock {
	if m != nil {
		return m.IpBlock
	}
	return nil
}

func (m *IPDesc) GetState() IPDesc_IPState {
	if m != nil {
		return m.State
	}
	return IPDesc_FREE
}

func (m *IPDesc) GetSid() *SubscriberID {
	if m != nil {
		return m.Sid
	}
	return nil
}

func (m *IPDesc) GetType() IPDesc_IPType {
	if m != nil {
		return m.Type
	}
	return IPDesc_STATIC
}

func (m *IPDesc) GetVlanId() uint32 {
	if m != nil {
		return m.VlanId
	}
	return 0
}

// --------------------------------------------------------------------------
// [mobilityd] List of IP descriptors
// --------------------------------------------------------------------------
type IPDescs struct {
	IpDescs              []*IPDesc `protobuf:"bytes,1,rep,name=ip_descs,json=ipDescs,proto3" json:"ip_descs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *IPDescs) Reset()         { *m = IPDescs{} }
func (m *IPDescs) String() string { return proto.CompactTextString(m) }
func (*IPDescs) ProtoMessage()    {}
func (*IPDescs) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ec85b478524e2d6, []int{2}
}

func (m *IPDescs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPDescs.Unmarshal(m, b)
}
func (m *IPDescs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPDescs.Marshal(b, m, deterministic)
}
func (m *IPDescs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPDescs.Merge(m, src)
}
func (m *IPDescs) XXX_Size() int {
	return xxx_messageInfo_IPDescs.Size(m)
}
func (m *IPDescs) XXX_DiscardUnknown() {
	xxx_messageInfo_IPDescs.DiscardUnknown(m)
}

var xxx_messageInfo_IPDescs proto.InternalMessageInfo

func (m *IPDescs) GetIpDescs() []*IPDesc {
	if m != nil {
		return m.IpDescs
	}
	return nil
}

func init() {
	proto.RegisterEnum("magma.lte.IPDesc_IPState", IPDesc_IPState_name, IPDesc_IPState_value)
	proto.RegisterEnum("magma.lte.IPDesc_IPType", IPDesc_IPType_name, IPDesc_IPType_value)
	proto.RegisterType((*AssignedIPBlocks)(nil), "magma.lte.AssignedIPBlocks")
	proto.RegisterType((*IPDesc)(nil), "magma.lte.IPDesc")
	proto.RegisterType((*IPDescs)(nil), "magma.lte.IPDescs")
}

func init() { proto.RegisterFile("lte/protos/keyval.proto", fileDescriptor_7ec85b478524e2d6) }

var fileDescriptor_7ec85b478524e2d6 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x6b, 0xdb, 0x30,
	0x14, 0x87, 0x6b, 0x3b, 0x75, 0xd2, 0x97, 0x65, 0x68, 0x62, 0x10, 0x37, 0x63, 0x10, 0xcc, 0x0e,
	0x19, 0xeb, 0x6c, 0xc8, 0x60, 0x3b, 0xbb, 0xb1, 0xc6, 0x5c, 0x0c, 0x31, 0x72, 0xd8, 0x61, 0x97,
	0x60, 0x47, 0x22, 0x88, 0x2a, 0xb5, 0x89, 0xd4, 0x42, 0xfe, 0xf4, 0xdd, 0x86, 0x14, 0xaf, 0xa4,
	0x34, 0x27, 0xeb, 0xf9, 0x7d, 0x7e, 0x9f, 0xfc, 0x93, 0x60, 0x2c, 0x35, 0x8f, 0xdb, 0x7d, 0xa3,
	0x1b, 0x15, 0xdf, 0xf3, 0xc3, 0x53, 0x25, 0x23, 0x5b, 0xe1, 0xab, 0x5d, 0xb5, 0xdd, 0x55, 0x91,
	0xd4, 0x7c, 0xf2, 0xf1, 0x84, 0x51, 0x8f, 0xb5, 0xda, 0xec, 0x45, 0xcd, 0xf7, 0xac, 0x3e, 0x92,
	0x93, 0xc9, 0x49, 0x7b, 0xd7, 0xd4, 0x42, 0x0a, 0x7d, 0x60, 0xc7, 0x5e, 0x78, 0x07, 0x28, 0x51,
	0x4a, 0x6c, 0x1f, 0x38, 0xcb, 0x8a, 0x5b, 0xd9, 0x6c, 0xee, 0x15, 0xfe, 0x0e, 0x23, 0xd1, 0xae,
	0x6b, 0x53, 0xac, 0xa5, 0x50, 0x3a, 0x70, 0xa6, 0xde, 0x6c, 0x38, 0xc7, 0xd1, 0xb3, 0x31, 0xea,
	0x58, 0x3a, 0x14, 0xad, 0x5d, 0xe4, 0x42, 0xe9, 0xf0, 0xaf, 0x0b, 0x7e, 0x56, 0xa4, 0x5c, 0x6d,
	0xf0, 0x27, 0x70, 0x45, 0x1b, 0x38, 0x53, 0x67, 0x36, 0x9c, 0xbf, 0x7f, 0xf1, 0x5d, 0xc2, 0xd8,
	0x9e, 0x2b, 0x45, 0x5d, 0xd1, 0xe2, 0xaf, 0x30, 0xf8, 0x2f, 0x0a, 0x5c, 0xcb, 0x9e, 0x73, 0xf4,
	0x3b, 0x07, 0x8e, 0xe1, 0x52, 0xe9, 0x4a, 0xf3, 0xc0, 0x9b, 0x3a, 0xb3, 0xb7, 0xf3, 0xeb, 0x17,
	0xac, 0xd1, 0x46, 0x59, 0x51, 0x1a, 0x80, 0x1e, 0x39, 0xfc, 0x19, 0x3c, 0x25, 0x58, 0xd0, 0xb3,
	0xa3, 0xc7, 0x27, 0x78, 0xf9, 0x1c, 0x52, 0x96, 0x52, 0xc3, 0xe0, 0x1b, 0xe8, 0xe9, 0x43, 0xcb,
	0x83, 0x4b, 0x3b, 0x3a, 0x38, 0x37, 0x7a, 0x75, 0x68, 0x39, 0xb5, 0x14, 0x1e, 0x43, 0xff, 0x49,
	0x56, 0x0f, 0x6b, 0xc1, 0x02, 0x7f, 0xea, 0xcc, 0x46, 0xd4, 0x37, 0x65, 0xc6, 0xc2, 0x3b, 0xe8,
	0x77, 0x7b, 0xc0, 0x03, 0xe8, 0xfd, 0xa4, 0x84, 0xa0, 0x0b, 0x3c, 0x82, 0xab, 0x24, 0xcf, 0x97,
	0x8b, 0x64, 0x45, 0x52, 0xe4, 0xe0, 0x37, 0x30, 0xa0, 0x24, 0x27, 0x49, 0x49, 0x52, 0xe4, 0x62,
	0x00, 0x9f, 0x92, 0xa4, 0x20, 0x29, 0xf2, 0x8e, 0x9d, 0x92, 0xd0, 0xdf, 0x24, 0x45, 0xbd, 0xf0,
	0x8b, 0x49, 0xd3, 0x48, 0x0d, 0x53, 0xae, 0x92, 0x55, 0xb6, 0x40, 0x17, 0x78, 0x68, 0x0c, 0xeb,
	0x62, 0xb9, 0xcc, 0x91, 0x63, 0x1c, 0xe9, 0xaf, 0x45, 0x81, 0xbc, 0xf0, 0x87, 0x79, 0x6d, 0x36,
	0xaa, 0xf0, 0x8d, 0x4d, 0x95, 0x99, 0x75, 0x77, 0x72, 0xef, 0x5e, 0xfd, 0x8e, 0x09, 0xd5, 0xd2,
	0xb7, 0x1f, 0xfe, 0x5c, 0xdb, 0x66, 0x6c, 0x2e, 0xc9, 0x46, 0x36, 0x8f, 0x2c, 0xde, 0x36, 0xdd,
	0x6d, 0xa9, 0x7d, 0xfb, 0xfc, 0xf6, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x69, 0x4e, 0xc4, 0x85,
	0x02, 0x00, 0x00,
}
