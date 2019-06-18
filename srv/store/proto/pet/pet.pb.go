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

// Pet define structure for message Pet.
type Pet struct {
	Id                   uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status               PetStatus `protobuf:"varint,3,opt,name=status,proto3,enum=store.PetStatus" json:"status,omitempty"`
	PhotoUrls            string    `protobuf:"bytes,4,opt,name=photoUrls,proto3" json:"photoUrls,omitempty"`
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

func (m *Pet) GetPhotoUrls() string {
	if m != nil {
		return m.PhotoUrls
	}
	return ""
}

func init() {
	proto.RegisterEnum("store.PetStatus", PetStatus_name, PetStatus_value)
	proto.RegisterType((*UpsertPetsRequest)(nil), "store.UpsertPetsRequest")
	proto.RegisterType((*UpsertPetsResponse)(nil), "store.UpsertPetsResponse")
	proto.RegisterType((*Pet)(nil), "store.Pet")
}

func init() { proto.RegisterFile("proto/pet/pet.proto", fileDescriptor_b7b87aff2c076168) }

var fileDescriptor_b7b87aff2c076168 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0xed, 0x26, 0xf9, 0xfa, 0x99, 0x29, 0x96, 0x38, 0x7a, 0x58, 0x45, 0x24, 0xe4, 0x14, 0x3c,
	0x44, 0x4c, 0x7f, 0x41, 0xb4, 0x45, 0x0a, 0xa1, 0xc6, 0x94, 0x7a, 0xaf, 0x76, 0xc0, 0x80, 0x76,
	0xb7, 0x99, 0xa9, 0xbf, 0x5f, 0xba, 0x04, 0x23, 0xe8, 0x61, 0x61, 0xe7, 0xbd, 0x79, 0x33, 0xf3,
	0x1e, 0x9c, 0xda, 0xd6, 0x88, 0xb9, 0xb1, 0x24, 0x87, 0x97, 0xb9, 0x0a, 0xff, 0xb1, 0x98, 0x96,
	0x92, 0x09, 0x9c, 0xac, 0x2c, 0x53, 0x2b, 0x15, 0x09, 0xd7, 0xb4, 0xdb, 0x13, 0x0b, 0x5e, 0x41,
	0x60, 0x49, 0x58, 0xab, 0xd8, 0x4f, 0x47, 0x39, 0x64, 0xae, 0x35, 0xab, 0x48, 0x6a, 0x87, 0x27,
	0x67, 0x80, 0x3f, 0x45, 0x6c, 0xcd, 0x96, 0x29, 0xd9, 0x81, 0x5f, 0x91, 0xe0, 0x18, 0xbc, 0x66,
	0xa3, 0x55, 0xac, 0xd2, 0xa0, 0xf6, 0x9a, 0x0d, 0x22, 0x04, 0xdb, 0xf5, 0x07, 0x69, 0x2f, 0x56,
	0x69, 0x58, 0xbb, 0x3f, 0xa6, 0x30, 0x64, 0x59, 0xcb, 0x9e, 0xb5, 0x1f, 0xab, 0x74, 0x9c, 0x47,
	0xfd, 0x8a, 0xa5, 0xc3, 0xeb, 0x8e, 0xc7, 0x4b, 0x08, 0xed, 0x9b, 0x11, 0xb3, 0x6a, 0xdf, 0x59,
	0x07, 0x6e, 0x44, 0x0f, 0x5c, 0xdf, 0x42, 0xf8, 0x2d, 0xc1, 0x63, 0x08, 0x8b, 0xe7, 0x62, 0x5e,
	0x16, 0x77, 0xe5, 0x2c, 0x1a, 0xe0, 0x08, 0xfe, 0x57, 0xb3, 0xc5, 0x74, 0xbe, 0x78, 0x88, 0x14,
	0x1e, 0x41, 0xb0, 0x7c, 0x2c, 0xa7, 0x91, 0x97, 0x3f, 0x01, 0x1c, 0x24, 0xd4, 0x7e, 0x36, 0xaf,
	0x84, 0xf7, 0x00, 0xbd, 0x13, 0xd4, 0xdd, 0x19, 0xbf, 0x12, 0xb9, 0x38, 0xff, 0x83, 0xe9, 0x6c,
	0x0f, 0x5e, 0x86, 0x2e, 0xd1, 0xc9, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x35, 0x04, 0x7e,
	0x68, 0x01, 0x00, 0x00,
}
