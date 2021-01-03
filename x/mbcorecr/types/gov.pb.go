// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mbcorecr/gov.proto

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

type MsgCreateSuperIdentity struct {
	Creator    string     `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Uid        string     `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	CreationDt *time.Time `protobuf:"bytes,3,opt,name=creationDt,proto3,stdtime" json:"creationDt,omitempty"`
}

func (m *MsgCreateSuperIdentity) Reset()         { *m = MsgCreateSuperIdentity{} }
func (m *MsgCreateSuperIdentity) String() string { return proto.CompactTextString(m) }
func (*MsgCreateSuperIdentity) ProtoMessage()    {}
func (*MsgCreateSuperIdentity) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef5434c92c47b57d, []int{0}
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
	return fileDescriptor_ef5434c92c47b57d, []int{1}
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
	proto.RegisterType((*MsgCreateSuperIdentity)(nil), "metabelarus.mbcorecr.mbcorecr.MsgCreateSuperIdentity")
	proto.RegisterType((*IdentityAccount)(nil), "metabelarus.mbcorecr.mbcorecr.IdentityAccount")
}

func init() { proto.RegisterFile("mbcorecr/gov.proto", fileDescriptor_ef5434c92c47b57d) }

var fileDescriptor_ef5434c92c47b57d = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x41, 0x4e, 0x02, 0x31,
	0x18, 0x85, 0xa9, 0x20, 0x48, 0x5d, 0x68, 0x1a, 0x43, 0x9a, 0x49, 0x2c, 0x84, 0x15, 0xab, 0x19,
	0xa3, 0x17, 0x50, 0x64, 0x83, 0xc6, 0x0d, 0xba, 0x72, 0x37, 0xd3, 0xa9, 0xb5, 0x09, 0x9d, 0x4e,
	0x3a, 0x2d, 0x91, 0x0b, 0xb8, 0xe6, 0x1e, 0x5e, 0xc4, 0x25, 0x4b, 0x77, 0x1a, 0xb8, 0x88, 0x99,
	0xce, 0x14, 0xd8, 0xbd, 0xf7, 0xbf, 0xff, 0xe5, 0x6b, 0xfa, 0x43, 0x24, 0x13, 0xaa, 0x34, 0xa3,
	0x3a, 0xe2, 0x6a, 0x11, 0xe6, 0x5a, 0x19, 0x85, 0x2e, 0x25, 0x33, 0x71, 0xc2, 0xe6, 0xb1, 0xb6,
	0x45, 0xe8, 0xf3, 0x9d, 0x08, 0x2e, 0xb8, 0xe2, 0xca, 0x6d, 0x46, 0xa5, 0xaa, 0x4a, 0x41, 0x9f,
	0x2b, 0xc5, 0xe7, 0x2c, 0x72, 0x2e, 0xb1, 0x6f, 0x91, 0x11, 0x92, 0x15, 0x26, 0x96, 0x79, 0xb5,
	0x30, 0xfc, 0x04, 0xb0, 0xf7, 0x54, 0xf0, 0x7b, 0xcd, 0x62, 0xc3, 0x9e, 0x6d, 0xce, 0xf4, 0x34,
	0x65, 0x99, 0x11, 0x66, 0x89, 0x30, 0xec, 0xd0, 0x72, 0xac, 0x34, 0x06, 0x03, 0x30, 0xea, 0xce,
	0xbc, 0x45, 0xe7, 0xb0, 0x69, 0x45, 0x8a, 0x8f, 0xdc, 0xb4, 0x94, 0xe8, 0x16, 0x42, 0x17, 0x0a,
	0x95, 0x4d, 0x0c, 0x6e, 0x0e, 0xc0, 0xe8, 0xf4, 0x3a, 0x08, 0x2b, 0x78, 0xe8, 0xe1, 0xe1, 0x8b,
	0x87, 0x8f, 0x5b, 0xab, 0xdf, 0x3e, 0x98, 0x1d, 0x74, 0x86, 0x5f, 0x00, 0x9e, 0x79, 0xf4, 0x1d,
	0xa5, 0xca, 0x66, 0xc6, 0x73, 0xc0, 0x9e, 0x83, 0x61, 0x27, 0x4e, 0x53, 0xcd, 0x8a, 0xa2, 0xa6,
	0x7b, 0x8b, 0x7a, 0xb0, 0x9d, 0xdb, 0xe4, 0x91, 0x2d, 0x1d, 0xbd, 0x3b, 0xab, 0x5d, 0xd9, 0xc8,
	0xb5, 0x58, 0x94, 0x41, 0xab, 0x6a, 0xd4, 0x16, 0x05, 0xf0, 0x44, 0x66, 0x4c, 0xaa, 0x4c, 0x50,
	0x7c, 0xec, 0xa2, 0x9d, 0x47, 0x04, 0x42, 0x51, 0x3f, 0x66, 0x3a, 0xc1, 0x6d, 0x97, 0x1e, 0x4c,
	0xc6, 0x0f, 0xdf, 0x1b, 0x02, 0xd6, 0x1b, 0x02, 0xfe, 0x36, 0x04, 0xac, 0xb6, 0xa4, 0xb1, 0xde,
	0x92, 0xc6, 0xcf, 0x96, 0x34, 0x5e, 0xaf, 0xb8, 0x30, 0xef, 0x36, 0x09, 0xa9, 0x92, 0xd1, 0xc1,
	0xc5, 0xa2, 0xdd, 0x45, 0x3f, 0xf6, 0xd2, 0x2c, 0x73, 0x56, 0x24, 0x6d, 0xf7, 0x3f, 0x37, 0xff,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x87, 0xed, 0xef, 0xf5, 0x01, 0x00, 0x00,
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
		n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.CreationDt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreationDt):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintGov(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Uid) > 0 {
		i -= len(m.Uid)
		copy(dAtA[i:], m.Uid)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Uid)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Creator)))
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
		i = encodeVarintGov(dAtA, i, uint64(len(m.IdentityID)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Mnemonic) > 0 {
		i -= len(m.Mnemonic)
		copy(dAtA[i:], m.Mnemonic)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Mnemonic)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PrivKey) > 0 {
		i -= len(m.PrivKey)
		copy(dAtA[i:], m.PrivKey)
		i = encodeVarintGov(dAtA, i, uint64(len(m.PrivKey)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintGov(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Uid) > 0 {
		i -= len(m.Uid)
		copy(dAtA[i:], m.Uid)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Uid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateSuperIdentity) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Uid)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if m.CreationDt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreationDt)
		n += 1 + l + sovGov(uint64(l))
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
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.PrivKey)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Mnemonic)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.IdentityID)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateSuperIdentity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
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
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
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
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
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
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGov
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGov
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
				return ErrIntOverflowGov
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
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
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
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
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
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
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
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
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
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
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
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IdentityID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGov
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)
