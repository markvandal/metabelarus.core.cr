// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mbcorecr/identity.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type IdentityType int32

const (
	IdentityType_CITIZEN         IdentityType = 0
	IdentityType_FOREIGNER       IdentityType = 1
	IdentityType_DIASPORA_MEMBER IdentityType = 2
	IdentityType_SERVICE         IdentityType = 3
)

var IdentityType_name = map[int32]string{
	0: "CITIZEN",
	1: "FOREIGNER",
	2: "DIASPORA_MEMBER",
	3: "SERVICE",
}

var IdentityType_value = map[string]int32{
	"CITIZEN":         0,
	"FOREIGNER":       1,
	"DIASPORA_MEMBER": 2,
	"SERVICE":         3,
}

func (x IdentityType) String() string {
	return proto.EnumName(IdentityType_name, int32(x))
}

func (IdentityType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f60e6993b4abcee2, []int{0}
}

type Identity struct {
	Id           string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AccountID    string       `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
	IdentityType IdentityType `protobuf:"varint,3,opt,name=identityType,proto3,enum=metabelarus.mbcorecr.mbcorecr.IdentityType" json:"identityType,omitempty"`
	Details      string       `protobuf:"bytes,4,opt,name=details,proto3" json:"details,omitempty"`
	InvitationId string       `protobuf:"bytes,5,opt,name=invitationId,proto3" json:"invitationId,omitempty"`
	CreationDt   *time.Time   `protobuf:"bytes,6,opt,name=creationDt,proto3,stdtime" json:"creationDt,omitempty"`
}

func (m *Identity) Reset()         { *m = Identity{} }
func (m *Identity) String() string { return proto.CompactTextString(m) }
func (*Identity) ProtoMessage()    {}
func (*Identity) Descriptor() ([]byte, []int) {
	return fileDescriptor_f60e6993b4abcee2, []int{0}
}
func (m *Identity) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Identity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Identity.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Identity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identity.Merge(m, src)
}
func (m *Identity) XXX_Size() int {
	return m.Size()
}
func (m *Identity) XXX_DiscardUnknown() {
	xxx_messageInfo_Identity.DiscardUnknown(m)
}

var xxx_messageInfo_Identity proto.InternalMessageInfo

func (m *Identity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Identity) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *Identity) GetIdentityType() IdentityType {
	if m != nil {
		return m.IdentityType
	}
	return IdentityType_CITIZEN
}

func (m *Identity) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

func (m *Identity) GetInvitationId() string {
	if m != nil {
		return m.InvitationId
	}
	return ""
}

func (m *Identity) GetCreationDt() *time.Time {
	if m != nil {
		return m.CreationDt
	}
	return nil
}

type MsgUpdateIdentity struct {
	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AccountID string `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
	Details   string `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
}

func (m *MsgUpdateIdentity) Reset()         { *m = MsgUpdateIdentity{} }
func (m *MsgUpdateIdentity) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateIdentity) ProtoMessage()    {}
func (*MsgUpdateIdentity) Descriptor() ([]byte, []int) {
	return fileDescriptor_f60e6993b4abcee2, []int{1}
}
func (m *MsgUpdateIdentity) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateIdentity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateIdentity.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateIdentity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateIdentity.Merge(m, src)
}
func (m *MsgUpdateIdentity) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateIdentity) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateIdentity.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateIdentity proto.InternalMessageInfo

func (m *MsgUpdateIdentity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MsgUpdateIdentity) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *MsgUpdateIdentity) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

type MsgCreateSuperIdentity struct {
	Creator    string     `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Uid        string     `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	CreationDt *time.Time `protobuf:"bytes,3,opt,name=creationDt,proto3,stdtime" json:"creationDt,omitempty"`
}

func (m *MsgCreateSuperIdentity) Reset()         { *m = MsgCreateSuperIdentity{} }
func (m *MsgCreateSuperIdentity) String() string { return proto.CompactTextString(m) }
func (*MsgCreateSuperIdentity) ProtoMessage()    {}
func (*MsgCreateSuperIdentity) Descriptor() ([]byte, []int) {
	return fileDescriptor_f60e6993b4abcee2, []int{2}
}
func (m *MsgCreateSuperIdentity) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateSuperIdentity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateSuperIdentity.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateSuperIdentity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateSuperIdentity.Merge(m, src)
}
func (m *MsgCreateSuperIdentity) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateSuperIdentity) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateSuperIdentity.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateSuperIdentity proto.InternalMessageInfo

func (m *MsgCreateSuperIdentity) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateSuperIdentity) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *MsgCreateSuperIdentity) GetCreationDt() *time.Time {
	if m != nil {
		return m.CreationDt
	}
	return nil
}

type IdentityAccount struct {
	Uid        string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Address    string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	PubKey     string `protobuf:"bytes,3,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	PrivKey    string `protobuf:"bytes,4,opt,name=privKey,proto3" json:"privKey,omitempty"`
	Mnemonic   string `protobuf:"bytes,5,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"`
	IdentityID string `protobuf:"bytes,6,opt,name=identityID,proto3" json:"identityID,omitempty"`
}

func (m *IdentityAccount) Reset()         { *m = IdentityAccount{} }
func (m *IdentityAccount) String() string { return proto.CompactTextString(m) }
func (*IdentityAccount) ProtoMessage()    {}
func (*IdentityAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_f60e6993b4abcee2, []int{3}
}
func (m *IdentityAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IdentityAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IdentityAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IdentityAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityAccount.Merge(m, src)
}
func (m *IdentityAccount) XXX_Size() int {
	return m.Size()
}
func (m *IdentityAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityAccount.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityAccount proto.InternalMessageInfo

func (m *IdentityAccount) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *IdentityAccount) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *IdentityAccount) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func (m *IdentityAccount) GetPrivKey() string {
	if m != nil {
		return m.PrivKey
	}
	return ""
}

func (m *IdentityAccount) GetMnemonic() string {
	if m != nil {
		return m.Mnemonic
	}
	return ""
}

func (m *IdentityAccount) GetIdentityID() string {
	if m != nil {
		return m.IdentityID
	}
	return ""
}

func init() {
	proto.RegisterEnum("metabelarus.mbcorecr.mbcorecr.IdentityType", IdentityType_name, IdentityType_value)
	proto.RegisterType((*Identity)(nil), "metabelarus.mbcorecr.mbcorecr.Identity")
	proto.RegisterType((*MsgUpdateIdentity)(nil), "metabelarus.mbcorecr.mbcorecr.MsgUpdateIdentity")
	proto.RegisterType((*MsgCreateSuperIdentity)(nil), "metabelarus.mbcorecr.mbcorecr.MsgCreateSuperIdentity")
	proto.RegisterType((*IdentityAccount)(nil), "metabelarus.mbcorecr.mbcorecr.IdentityAccount")
}

func init() { proto.RegisterFile("mbcorecr/identity.proto", fileDescriptor_f60e6993b4abcee2) }

var fileDescriptor_f60e6993b4abcee2 = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xb3, 0x71, 0x49, 0x9a, 0x69, 0x68, 0xcd, 0x82, 0x8a, 0x15, 0x81, 0x1b, 0xe5, 0x14,
	0x81, 0x64, 0xa3, 0xf2, 0x02, 0xe4, 0xc3, 0xa0, 0x05, 0xd2, 0x20, 0x27, 0x70, 0x28, 0x07, 0xe4,
	0x8f, 0xc5, 0xac, 0x14, 0x7b, 0x2d, 0x7b, 0x5d, 0x91, 0x17, 0x40, 0xe2, 0xd6, 0xf7, 0xe0, 0x45,
	0x38, 0xf6, 0xc8, 0x0d, 0x94, 0xbc, 0x08, 0xb2, 0xe3, 0x4d, 0x1c, 0x0e, 0x1c, 0x7a, 0x9b, 0xff,
	0xec, 0x8e, 0x7f, 0xf3, 0xdf, 0x19, 0xc3, 0xc3, 0xd0, 0xf5, 0x78, 0x42, 0xbd, 0xc4, 0x64, 0x3e,
	0x8d, 0x04, 0x13, 0x4b, 0x23, 0x4e, 0xb8, 0xe0, 0xf8, 0x71, 0x48, 0x85, 0xe3, 0xd2, 0x85, 0x93,
	0x64, 0xa9, 0x21, 0x2f, 0x6d, 0x83, 0xce, 0x83, 0x80, 0x07, 0xbc, 0xb8, 0x69, 0xe6, 0xd1, 0xa6,
	0xa8, 0x73, 0x16, 0x70, 0x1e, 0x2c, 0xa8, 0x59, 0x28, 0x37, 0xfb, 0x6c, 0x0a, 0x16, 0xd2, 0x54,
	0x38, 0x61, 0xbc, 0xb9, 0xd0, 0xfb, 0x5e, 0x87, 0x43, 0x52, 0x82, 0xf0, 0x31, 0xd4, 0x99, 0xaf,
	0xa1, 0x2e, 0xea, 0xb7, 0xec, 0x3a, 0xf3, 0xf1, 0x23, 0x68, 0x39, 0x9e, 0xc7, 0xb3, 0x48, 0x90,
	0xb1, 0x56, 0x2f, 0xd2, 0xbb, 0x04, 0x9e, 0x42, 0x5b, 0xb6, 0x38, 0x5f, 0xc6, 0x54, 0x53, 0xba,
	0xa8, 0x7f, 0x7c, 0xfe, 0xd4, 0xf8, 0x6f, 0x9f, 0x06, 0xa9, 0x94, 0xd8, 0x7b, 0x1f, 0xc0, 0x1a,
	0x34, 0x7d, 0x2a, 0x1c, 0xb6, 0x48, 0xb5, 0x83, 0x02, 0x26, 0x25, 0xee, 0x41, 0x9b, 0x45, 0x57,
	0x4c, 0x38, 0x82, 0xf1, 0x88, 0xf8, 0xda, 0x9d, 0xe2, 0x78, 0x2f, 0x87, 0x5f, 0x00, 0x78, 0x09,
	0x2d, 0xd4, 0x58, 0x68, 0x8d, 0x2e, 0xea, 0x1f, 0x9d, 0x77, 0x8c, 0x8d, 0x7f, 0x43, 0xfa, 0x37,
	0xe6, 0xd2, 0xff, 0xf0, 0xe0, 0xfa, 0xf7, 0x19, 0xb2, 0x2b, 0x35, 0xbd, 0x8f, 0x70, 0x6f, 0x92,
	0x06, 0xef, 0x63, 0xdf, 0x11, 0xf4, 0x96, 0x6f, 0x52, 0xb1, 0xa0, 0xec, 0x59, 0xe8, 0x7d, 0x43,
	0x70, 0x3a, 0x49, 0x83, 0x51, 0x8e, 0xa3, 0xb3, 0x2c, 0xa6, 0xc9, 0x16, 0xa1, 0x41, 0xb3, 0xe8,
	0x82, 0x27, 0x25, 0x47, 0x4a, 0xac, 0x82, 0x92, 0x31, 0xbf, 0xc4, 0xe4, 0xe1, 0x3f, 0x2e, 0x95,
	0x5b, 0xb8, 0xfc, 0x81, 0xe0, 0x44, 0xa2, 0x07, 0x9b, 0xc6, 0x25, 0x07, 0xed, 0x38, 0x1a, 0x34,
	0x1d, 0xdf, 0x4f, 0x68, 0x9a, 0x96, 0x74, 0x29, 0xf1, 0x29, 0x34, 0xe2, 0xcc, 0x7d, 0x43, 0x97,
	0xa5, 0xc3, 0x52, 0xe5, 0x15, 0x71, 0xc2, 0xae, 0xf2, 0x83, 0x72, 0x7a, 0xa5, 0xc4, 0x1d, 0x38,
	0x0c, 0x23, 0x1a, 0xf2, 0x88, 0x79, 0xe5, 0xe4, 0xb6, 0x1a, 0xeb, 0x00, 0x72, 0x07, 0xc8, 0xb8,
	0x98, 0x5a, 0xcb, 0xae, 0x64, 0x9e, 0xbc, 0x85, 0x76, 0x75, 0x63, 0xf0, 0x11, 0x34, 0x47, 0x64,
	0x4e, 0x2e, 0xad, 0x0b, 0xb5, 0x86, 0xef, 0x42, 0xeb, 0xe5, 0xd4, 0xb6, 0xc8, 0xab, 0x0b, 0xcb,
	0x56, 0x11, 0xbe, 0x0f, 0x27, 0x63, 0x32, 0x98, 0xbd, 0x9b, 0xda, 0x83, 0x4f, 0x13, 0x6b, 0x32,
	0xb4, 0x6c, 0xb5, 0x9e, 0x17, 0xcc, 0x2c, 0xfb, 0x03, 0x19, 0x59, 0xaa, 0x32, 0x7c, 0xfd, 0x73,
	0xa5, 0xa3, 0x9b, 0x95, 0x8e, 0xfe, 0xac, 0x74, 0x74, 0xbd, 0xd6, 0x6b, 0x37, 0x6b, 0xbd, 0xf6,
	0x6b, 0xad, 0xd7, 0x2e, 0x9f, 0x05, 0x4c, 0x7c, 0xc9, 0x5c, 0xc3, 0xe3, 0xa1, 0x59, 0x59, 0x60,
	0x73, 0xfb, 0x37, 0x7e, 0xdd, 0x85, 0x62, 0x19, 0xd3, 0xd4, 0x6d, 0x14, 0xaf, 0xfd, 0xfc, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x61, 0xb0, 0x34, 0xb1, 0x03, 0x00, 0x00,
}

func (m *Identity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Identity) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Identity) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreationDt != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.CreationDt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreationDt):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintIdentity(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x32
	}
	if len(m.InvitationId) > 0 {
		i -= len(m.InvitationId)
		copy(dAtA[i:], m.InvitationId)
		i = encodeVarintIdentity(dAtA, i, uint64(len(m.InvitationId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Details) > 0 {
		i -= len(m.Details)
		copy(dAtA[i:], m.Details)
		i = encodeVarintIdentity(dAtA, i, uint64(len(m.Details)))
		i--
		dAtA[i] = 0x22
	}
	if m.IdentityType != 0 {
		i = encodeVarintIdentity(dAtA, i, uint64(m.IdentityType))
		i--
		dAtA[i] = 0x18
	}
	if len(m.AccountID) > 0 {
		i -= len(m.AccountID)
		copy(dAtA[i:], m.AccountID)
		i = encodeVarintIdentity(dAtA, i, uint64(len(m.AccountID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintIdentity(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateIdentity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateIdentity) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateIdentity) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Details) > 0 {
		i -= len(m.Details)
		copy(dAtA[i:], m.Details)
		i = encodeVarintIdentity(dAtA, i, uint64(len(m.Details)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AccountID) > 0 {
		i -= len(m.AccountID)
		copy(dAtA[i:], m.AccountID)
		i = encodeVarintIdentity(dAtA, i, uint64(len(m.AccountID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintIdentity(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateSuperIdentity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateSuperIdentity) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateSuperIdentity) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreationDt != nil {
		n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.CreationDt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreationDt):])
		if err2 != nil {
			return 0, err2
		}
		i -= n2
		i = encodeVarintIdentity(dAtA, i, uint64(n2))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Uid) > 0 {
		i -= len(m.Uid)
		copy(dAtA[i:], m.Uid)
		i = encodeVarintIdentity(dAtA, i, uint64(len(m.Uid)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintIdentity(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IdentityAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IdentityAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IdentityAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IdentityID) > 0 {
		i -= len(m.IdentityID)
		copy(dAtA[i:], m.IdentityID)
		i = encodeVarintIdentity(dAtA, i, uint64(len(m.IdentityID)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Mnemonic) > 0 {
		i -= len(m.Mnemonic)
		copy(dAtA[i:], m.Mnemonic)
		i = encodeVarintIdentity(dAtA, i, uint64(len(m.Mnemonic)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PrivKey) > 0 {
		i -= len(m.PrivKey)
		copy(dAtA[i:], m.PrivKey)
		i = encodeVarintIdentity(dAtA, i, uint64(len(m.PrivKey)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintIdentity(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintIdentity(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Uid) > 0 {
		i -= len(m.Uid)
		copy(dAtA[i:], m.Uid)
		i = encodeVarintIdentity(dAtA, i, uint64(len(m.Uid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIdentity(dAtA []byte, offset int, v uint64) int {
	offset -= sovIdentity(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Identity) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovIdentity(uint64(l))
	}
	l = len(m.AccountID)
	if l > 0 {
		n += 1 + l + sovIdentity(uint64(l))
	}
	if m.IdentityType != 0 {
		n += 1 + sovIdentity(uint64(m.IdentityType))
	}
	l = len(m.Details)
	if l > 0 {
		n += 1 + l + sovIdentity(uint64(l))
	}
	l = len(m.InvitationId)
	if l > 0 {
		n += 1 + l + sovIdentity(uint64(l))
	}
	if m.CreationDt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreationDt)
		n += 1 + l + sovIdentity(uint64(l))
	}
	return n
}

func (m *MsgUpdateIdentity) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovIdentity(uint64(l))
	}
	l = len(m.AccountID)
	if l > 0 {
		n += 1 + l + sovIdentity(uint64(l))
	}
	l = len(m.Details)
	if l > 0 {
		n += 1 + l + sovIdentity(uint64(l))
	}
	return n
}

func (m *MsgCreateSuperIdentity) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovIdentity(uint64(l))
	}
	l = len(m.Uid)
	if l > 0 {
		n += 1 + l + sovIdentity(uint64(l))
	}
	if m.CreationDt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreationDt)
		n += 1 + l + sovIdentity(uint64(l))
	}
	return n
}

func (m *IdentityAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Uid)
	if l > 0 {
		n += 1 + l + sovIdentity(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovIdentity(uint64(l))
	}
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovIdentity(uint64(l))
	}
	l = len(m.PrivKey)
	if l > 0 {
		n += 1 + l + sovIdentity(uint64(l))
	}
	l = len(m.Mnemonic)
	if l > 0 {
		n += 1 + l + sovIdentity(uint64(l))
	}
	l = len(m.IdentityID)
	if l > 0 {
		n += 1 + l + sovIdentity(uint64(l))
	}
	return n
}

func sovIdentity(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIdentity(x uint64) (n int) {
	return sovIdentity(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Identity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIdentity
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
			return fmt.Errorf("proto: Identity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Identity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentity
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
				return ErrInvalidLengthIdentity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentity
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
				return ErrInvalidLengthIdentity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdentityType", wireType)
			}
			m.IdentityType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IdentityType |= IdentityType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentity
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
				return ErrInvalidLengthIdentity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Details = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InvitationId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentity
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
				return ErrInvalidLengthIdentity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InvitationId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationDt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentity
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
				return ErrInvalidLengthIdentity
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIdentity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreationDt == nil {
				m.CreationDt = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.CreationDt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIdentity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIdentity
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIdentity
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
func (m *MsgUpdateIdentity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIdentity
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
			return fmt.Errorf("proto: MsgUpdateIdentity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateIdentity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentity
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
				return ErrInvalidLengthIdentity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentity
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
				return ErrInvalidLengthIdentity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentity
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
				return ErrInvalidLengthIdentity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Details = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIdentity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIdentity
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIdentity
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
func (m *MsgCreateSuperIdentity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIdentity
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
			return fmt.Errorf("proto: MsgCreateSuperIdentity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateSuperIdentity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentity
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
				return ErrInvalidLengthIdentity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentity
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
				return ErrInvalidLengthIdentity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationDt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentity
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
				return ErrInvalidLengthIdentity
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIdentity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreationDt == nil {
				m.CreationDt = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.CreationDt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIdentity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIdentity
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIdentity
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
func (m *IdentityAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIdentity
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
			return fmt.Errorf("proto: IdentityAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IdentityAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentity
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
				return ErrInvalidLengthIdentity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentity
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
				return ErrInvalidLengthIdentity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentity
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
				return ErrInvalidLengthIdentity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrivKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentity
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
				return ErrInvalidLengthIdentity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrivKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mnemonic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentity
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
				return ErrInvalidLengthIdentity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mnemonic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdentityID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentity
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
				return ErrInvalidLengthIdentity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IdentityID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIdentity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIdentity
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIdentity
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
func skipIdentity(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIdentity
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
					return 0, ErrIntOverflowIdentity
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
					return 0, ErrIntOverflowIdentity
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
				return 0, ErrInvalidLengthIdentity
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIdentity
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIdentity
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIdentity        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIdentity          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIdentity = fmt.Errorf("proto: unexpected end of group")
)