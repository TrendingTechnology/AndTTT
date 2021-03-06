// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package proto

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

type ApiStatus struct {
	NumGoroutines          int32    `protobuf:"varint,1,opt,name=num_goroutines,json=numGoroutines,proto3" json:"num_goroutines,omitempty"`
	NumRegisteredClients   int32    `protobuf:"varint,2,opt,name=num_registered_clients,json=numRegisteredClients,proto3" json:"num_registered_clients,omitempty"`
	NumUnregisteredClients int32    `protobuf:"varint,3,opt,name=num_unregistered_clients,json=numUnregisteredClients,proto3" json:"num_unregistered_clients,omitempty"`
	NumRooms               int32    `protobuf:"varint,4,opt,name=num_rooms,json=numRooms,proto3" json:"num_rooms,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *ApiStatus) Reset()         { *m = ApiStatus{} }
func (m *ApiStatus) String() string { return proto.CompactTextString(m) }
func (*ApiStatus) ProtoMessage()    {}
func (*ApiStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_80257ec225fb8ab0, []int{0}
}
func (m *ApiStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiStatus.Unmarshal(m, b)
}
func (m *ApiStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiStatus.Marshal(b, m, deterministic)
}
func (dst *ApiStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiStatus.Merge(dst, src)
}
func (m *ApiStatus) XXX_Size() int {
	return xxx_messageInfo_ApiStatus.Size(m)
}
func (m *ApiStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ApiStatus proto.InternalMessageInfo

func (m *ApiStatus) GetNumGoroutines() int32 {
	if m != nil {
		return m.NumGoroutines
	}
	return 0
}

func (m *ApiStatus) GetNumRegisteredClients() int32 {
	if m != nil {
		return m.NumRegisteredClients
	}
	return 0
}

func (m *ApiStatus) GetNumUnregisteredClients() int32 {
	if m != nil {
		return m.NumUnregisteredClients
	}
	return 0
}

func (m *ApiStatus) GetNumRooms() int32 {
	if m != nil {
		return m.NumRooms
	}
	return 0
}

type ApiRooms struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Rooms                []*Room  `protobuf:"bytes,2,rep,name=rooms,proto3" json:"rooms,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiRooms) Reset()         { *m = ApiRooms{} }
func (m *ApiRooms) String() string { return proto.CompactTextString(m) }
func (*ApiRooms) ProtoMessage()    {}
func (*ApiRooms) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_80257ec225fb8ab0, []int{1}
}
func (m *ApiRooms) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiRooms.Unmarshal(m, b)
}
func (m *ApiRooms) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiRooms.Marshal(b, m, deterministic)
}
func (dst *ApiRooms) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiRooms.Merge(dst, src)
}
func (m *ApiRooms) XXX_Size() int {
	return xxx_messageInfo_ApiRooms.Size(m)
}
func (m *ApiRooms) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiRooms.DiscardUnknown(m)
}

var xxx_messageInfo_ApiRooms proto.InternalMessageInfo

func (m *ApiRooms) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ApiRooms) GetRooms() []*Room {
	if m != nil {
		return m.Rooms
	}
	return nil
}

type ApiRoom struct {
	Room                 *Room    `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiRoom) Reset()         { *m = ApiRoom{} }
func (m *ApiRoom) String() string { return proto.CompactTextString(m) }
func (*ApiRoom) ProtoMessage()    {}
func (*ApiRoom) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_80257ec225fb8ab0, []int{2}
}
func (m *ApiRoom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiRoom.Unmarshal(m, b)
}
func (m *ApiRoom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiRoom.Marshal(b, m, deterministic)
}
func (dst *ApiRoom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiRoom.Merge(dst, src)
}
func (m *ApiRoom) XXX_Size() int {
	return xxx_messageInfo_ApiRoom.Size(m)
}
func (m *ApiRoom) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiRoom.DiscardUnknown(m)
}

var xxx_messageInfo_ApiRoom proto.InternalMessageInfo

func (m *ApiRoom) GetRoom() *Room {
	if m != nil {
		return m.Room
	}
	return nil
}

type ApiClients struct {
	Count                int32     `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Clients              []*Client `protobuf:"bytes,2,rep,name=clients,proto3" json:"clients,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ApiClients) Reset()         { *m = ApiClients{} }
func (m *ApiClients) String() string { return proto.CompactTextString(m) }
func (*ApiClients) ProtoMessage()    {}
func (*ApiClients) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_80257ec225fb8ab0, []int{3}
}
func (m *ApiClients) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiClients.Unmarshal(m, b)
}
func (m *ApiClients) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiClients.Marshal(b, m, deterministic)
}
func (dst *ApiClients) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiClients.Merge(dst, src)
}
func (m *ApiClients) XXX_Size() int {
	return xxx_messageInfo_ApiClients.Size(m)
}
func (m *ApiClients) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiClients.DiscardUnknown(m)
}

var xxx_messageInfo_ApiClients proto.InternalMessageInfo

func (m *ApiClients) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ApiClients) GetClients() []*Client {
	if m != nil {
		return m.Clients
	}
	return nil
}

type ApiClient struct {
	Client               *Client  `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiClient) Reset()         { *m = ApiClient{} }
func (m *ApiClient) String() string { return proto.CompactTextString(m) }
func (*ApiClient) ProtoMessage()    {}
func (*ApiClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_80257ec225fb8ab0, []int{4}
}
func (m *ApiClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiClient.Unmarshal(m, b)
}
func (m *ApiClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiClient.Marshal(b, m, deterministic)
}
func (dst *ApiClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiClient.Merge(dst, src)
}
func (m *ApiClient) XXX_Size() int {
	return xxx_messageInfo_ApiClient.Size(m)
}
func (m *ApiClient) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiClient.DiscardUnknown(m)
}

var xxx_messageInfo_ApiClient proto.InternalMessageInfo

func (m *ApiClient) GetClient() *Client {
	if m != nil {
		return m.Client
	}
	return nil
}

func init() {
	proto.RegisterType((*ApiStatus)(nil), "proto.ApiStatus")
	proto.RegisterType((*ApiRooms)(nil), "proto.ApiRooms")
	proto.RegisterType((*ApiRoom)(nil), "proto.ApiRoom")
	proto.RegisterType((*ApiClients)(nil), "proto.ApiClients")
	proto.RegisterType((*ApiClient)(nil), "proto.ApiClient")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_api_80257ec225fb8ab0) }

var fileDescriptor_api_80257ec225fb8ab0 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x4f, 0x4f, 0xc2, 0x30,
	0x14, 0xcf, 0x80, 0x01, 0x7b, 0x88, 0x87, 0x86, 0x98, 0x05, 0x0f, 0xe0, 0x0c, 0x91, 0x78, 0xd8,
	0x01, 0x8d, 0xf1, 0x3a, 0x39, 0x78, 0xf0, 0x62, 0x66, 0x3c, 0x93, 0xb1, 0x2d, 0xb3, 0x89, 0x6d,
	0x97, 0xad, 0x8d, 0x1f, 0xcf, 0xaf, 0x66, 0xfa, 0x5e, 0x87, 0x84, 0x70, 0x5a, 0xdf, 0xef, 0x6f,
	0xf7, 0x0a, 0x41, 0x56, 0xf3, 0xb8, 0x6e, 0x94, 0x56, 0xcc, 0xc7, 0xcf, 0x1c, 0x1a, 0xa5, 0x04,
	0x41, 0xf3, 0x8b, 0xfc, 0x9b, 0x97, 0x52, 0xd3, 0x14, 0xfd, 0x7a, 0x10, 0x24, 0x35, 0xff, 0xd0,
	0x99, 0x36, 0x2d, 0x5b, 0xc1, 0xa5, 0x34, 0x62, 0x57, 0xa9, 0x46, 0x19, 0xcd, 0x65, 0xd9, 0x86,
	0xde, 0xd2, 0x5b, 0xfb, 0xe9, 0x54, 0x1a, 0xf1, 0x7a, 0x00, 0xd9, 0x23, 0x5c, 0x59, 0x59, 0x53,
	0x56, 0xbc, 0xd5, 0x65, 0x53, 0x16, 0x3b, 0xca, 0x6c, 0xc3, 0x1e, 0xca, 0x67, 0xd2, 0x88, 0xf4,
	0x40, 0x6e, 0x89, 0x63, 0xcf, 0x10, 0x5a, 0x97, 0x91, 0x67, 0x7c, 0x7d, 0xf4, 0xd9, 0xd4, 0xcf,
	0x23, 0xba, 0x73, 0x5e, 0x43, 0x80, 0x7d, 0x4a, 0x89, 0x36, 0x1c, 0xa0, 0x74, 0x6c, 0x2b, 0xec,
	0x1c, 0x6d, 0x61, 0x9c, 0xd4, 0x1c, 0xcf, 0x6c, 0x06, 0x7e, 0xae, 0x8c, 0xd4, 0xee, 0xda, 0x34,
	0xb0, 0x1b, 0xf0, 0xc9, 0xda, 0x5b, 0xf6, 0xd7, 0x93, 0xcd, 0x84, 0x7e, 0x3d, 0xb6, 0x96, 0x94,
	0x98, 0xe8, 0x1e, 0x46, 0x2e, 0x84, 0x2d, 0x60, 0x60, 0x31, 0x8c, 0x38, 0x11, 0x23, 0x11, 0xbd,
	0x01, 0x24, 0x35, 0xef, 0xee, 0x76, 0xbe, 0xf2, 0x0e, 0x46, 0xff, 0x2b, 0xb1, 0xa5, 0x53, 0x97,
	0x43, 0xb6, 0xb4, 0x63, 0xa3, 0x0d, 0xae, 0x9f, 0x50, 0xb6, 0x82, 0x21, 0xe1, 0xae, 0xfc, 0xc4,
	0xe4, 0xc8, 0x97, 0x5b, 0x58, 0xe4, 0x4a, 0xc4, 0x15, 0xd7, 0x5f, 0x66, 0x1f, 0x17, 0xd9, 0x0f,
	0x2f, 0x8a, 0xa7, 0x38, 0x93, 0x85, 0xd6, 0xee, 0x59, 0xdf, 0xbd, 0xfd, 0x10, 0x0f, 0x0f, 0x7f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x56, 0xc0, 0x79, 0x94, 0x0d, 0x02, 0x00, 0x00,
}
