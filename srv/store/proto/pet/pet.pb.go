// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/pet/pet.proto

package store

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

// PetStatus defines pet status
type PetStatus int32

const (
	PetStatus_AVAILABLE PetStatus = 0
	PetStatus_PENDING   PetStatus = 1
	PetStatus_SOLD      PetStatus = 2
)

var PetStatus_name = map[int32]string{
	0: "AVAILABLE",
	1: "PENDING",
	2: "SOLD",
}

var PetStatus_value = map[string]int32{
	"AVAILABLE": 0,
	"PENDING":   1,
	"SOLD":      2,
}

func (x PetStatus) String() string {
	return proto.EnumName(PetStatus_name, int32(x))
}

func (PetStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b7b87aff2c076168, []int{0}
}

// UpsertPetsRequest defines request for RPC call to upsert pets.
type UpsertPetsRequest struct {
	Pets                 []*Pet   `protobuf:"bytes,1,rep,name=pets,proto3" json:"pets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpsertPetsRequest) Reset()         { *m = UpsertPetsRequest{} }
func (m *UpsertPetsRequest) String() string { return proto.CompactTextString(m) }
func (*UpsertPetsRequest) ProtoMessage()    {}
func (*UpsertPetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7b87aff2c076168, []int{0}
}

func (m *UpsertPetsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpsertPetsRequest.Unmarshal(m, b)
}
func (m *UpsertPetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpsertPetsRequest.Marshal(b, m, deterministic)
}
func (m *UpsertPetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpsertPetsRequest.Merge(m, src)
}
func (m *UpsertPetsRequest) XXX_Size() int {
	return xxx_messageInfo_UpsertPetsRequest.Size(m)
}
func (m *UpsertPetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpsertPetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpsertPetsRequest proto.InternalMessageInfo

func (m *UpsertPetsRequest) GetPets() []*Pet {
	if m != nil {
		return m.Pets
	}
	return nil
}

// UpsertPetsResponse defines response for RPC call.
type UpsertPetsResponse struct {
	Pets                 []*Pet   `protobuf:"bytes,1,rep,name=pets,proto3" json:"pets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpsertPetsResponse) Reset()         { *m = UpsertPetsResponse{} }
func (m *UpsertPetsResponse) String() string { return proto.CompactTextString(m) }
func (*UpsertPetsResponse) ProtoMessage()    {}
func (*UpsertPetsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7b87aff2c076168, []int{1}
}

func (m *UpsertPetsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpsertPetsResponse.Unmarshal(m, b)
}
func (m *UpsertPetsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpsertPetsResponse.Marshal(b, m, deterministic)
}
func (m *UpsertPetsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpsertPetsResponse.Merge(m, src)
}
func (m *UpsertPetsResponse) XXX_Size() int {
	return xxx_messageInfo_UpsertPetsResponse.Size(m)
}
func (m *UpsertPetsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpsertPetsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpsertPetsResponse proto.InternalMessageInfo

func (m *UpsertPetsResponse) GetPets() []*Pet {
	if m != nil {
		return m.Pets
	}
	return nil
}

// Pet define structure for message Pet.
type Pet struct {
	Id                   uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status               PetStatus `protobuf:"varint,4,opt,name=status,proto3,enum=store.PetStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Pet) Reset()         { *m = Pet{} }
func (m *Pet) String() string { return proto.CompactTextString(m) }
func (*Pet) ProtoMessage()    {}
func (*Pet) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7b87aff2c076168, []int{2}
}

func (m *Pet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pet.Unmarshal(m, b)
}
func (m *Pet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pet.Marshal(b, m, deterministic)
}
func (m *Pet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pet.Merge(m, src)
}
func (m *Pet) XXX_Size() int {
	return xxx_messageInfo_Pet.Size(m)
}
func (m *Pet) XXX_DiscardUnknown() {
	xxx_messageInfo_Pet.DiscardUnknown(m)
}

var xxx_messageInfo_Pet proto.InternalMessageInfo

func (m *Pet) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Pet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Pet) GetStatus() PetStatus {
	if m != nil {
		return m.Status
	}
	return PetStatus_AVAILABLE
}

func init() {
	proto.RegisterEnum("store.PetStatus", PetStatus_name, PetStatus_value)
	proto.RegisterType((*UpsertPetsRequest)(nil), "store.UpsertPetsRequest")
	proto.RegisterType((*UpsertPetsResponse)(nil), "store.UpsertPetsResponse")
	proto.RegisterType((*Pet)(nil), "store.Pet")
}

func init() { proto.RegisterFile("proto/pet/pet.proto", fileDescriptor_b7b87aff2c076168) }

var fileDescriptor_b7b87aff2c076168 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x50, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0xed, 0xa6, 0x6b, 0xb5, 0x53, 0x2c, 0x71, 0xbc, 0xac, 0x1e, 0x24, 0xe4, 0xb4, 0x78, 0x88,
	0x98, 0xfa, 0x03, 0xd1, 0x16, 0x29, 0x84, 0x1a, 0x37, 0xe8, 0xbd, 0xda, 0x39, 0xe4, 0x60, 0xb3,
	0x66, 0xa6, 0x7e, 0xbf, 0x74, 0x2d, 0x56, 0x50, 0xf0, 0x30, 0x30, 0xf3, 0xde, 0xbc, 0x19, 0xde,
	0x83, 0x53, 0xdf, 0xb5, 0xd2, 0x5e, 0x79, 0x92, 0x6d, 0x65, 0x61, 0xc2, 0x03, 0x96, 0xb6, 0xa3,
	0x74, 0x02, 0x27, 0x4f, 0x9e, 0xa9, 0x93, 0x8a, 0x84, 0x1d, 0xbd, 0x6f, 0x88, 0x05, 0x2f, 0x40,
	0x7b, 0x12, 0x36, 0x2a, 0xe9, 0xdb, 0x51, 0x0e, 0x59, 0x58, 0xcd, 0x2a, 0x12, 0x17, 0xf0, 0xf4,
	0x06, 0xf0, 0xa7, 0x88, 0x7d, 0xbb, 0x66, 0xfa, 0x57, 0x55, 0x43, 0xbf, 0x22, 0xc1, 0x31, 0x44,
	0xcd, 0xca, 0xa8, 0x44, 0x59, 0xed, 0xa2, 0x66, 0x85, 0x08, 0x7a, 0xbd, 0x7c, 0x23, 0x13, 0x25,
	0xca, 0x0e, 0x5d, 0xe8, 0xd1, 0xc2, 0x80, 0x65, 0x29, 0x1b, 0x36, 0x3a, 0x51, 0x76, 0x9c, 0xc7,
	0xfb, 0x63, 0x75, 0xc0, 0xdd, 0x8e, 0xbf, 0xbc, 0x86, 0xe1, 0x37, 0x88, 0xc7, 0x30, 0x2c, 0x9e,
	0x8b, 0x79, 0x59, 0xdc, 0x96, 0xb3, 0xb8, 0x87, 0x23, 0x38, 0xac, 0x66, 0x8b, 0xe9, 0x7c, 0x71,
	0x1f, 0x2b, 0x3c, 0x02, 0x5d, 0x3f, 0x94, 0xd3, 0x38, 0xca, 0x1f, 0x01, 0xb6, 0x12, 0xea, 0x3e,
	0x9a, 0x57, 0xc2, 0x3b, 0x80, 0xbd, 0x17, 0x34, 0xbb, 0x47, 0xbf, 0x32, 0x39, 0x3f, 0xfb, 0x83,
	0xf9, 0x32, 0x9e, 0xf6, 0x5e, 0x06, 0x21, 0xd3, 0xc9, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x95,
	0x7f, 0x54, 0x7b, 0x6a, 0x01, 0x00, 0x00,
}