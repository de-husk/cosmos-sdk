// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/orm/testdata/codec.proto

package testdata

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GroupInfo represents the high-level on-chain information for a group.
type GroupInfo struct {
	GroupId     uint64                                        `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Description string                                        `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Admin       github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,3,opt,name=admin,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"admin,omitempty"`
}

func (m *GroupInfo) Reset()         { *m = GroupInfo{} }
func (m *GroupInfo) String() string { return proto.CompactTextString(m) }
func (*GroupInfo) ProtoMessage()    {}
func (*GroupInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9806bf430fc0d7f, []int{0}
}
func (m *GroupInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GroupInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GroupInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GroupInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupInfo.Merge(m, src)
}
func (m *GroupInfo) XXX_Size() int {
	return m.Size()
}
func (m *GroupInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GroupInfo proto.InternalMessageInfo

func (m *GroupInfo) GetGroupId() uint64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *GroupInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *GroupInfo) GetAdmin() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Admin
	}
	return nil
}

// GroupMember represents the relationship between a group and a member.
type GroupMember struct {
	Group  github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=group,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"group,omitempty"`
	Member github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=member,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"member,omitempty"`
	Weight uint64                                        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (m *GroupMember) Reset()         { *m = GroupMember{} }
func (m *GroupMember) String() string { return proto.CompactTextString(m) }
func (*GroupMember) ProtoMessage()    {}
func (*GroupMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9806bf430fc0d7f, []int{1}
}
func (m *GroupMember) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GroupMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GroupMember.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GroupMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupMember.Merge(m, src)
}
func (m *GroupMember) XXX_Size() int {
	return m.Size()
}
func (m *GroupMember) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupMember.DiscardUnknown(m)
}

var xxx_messageInfo_GroupMember proto.InternalMessageInfo

func (m *GroupMember) GetGroup() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *GroupMember) GetMember() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Member
	}
	return nil
}

func (m *GroupMember) GetWeight() uint64 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func init() {
	proto.RegisterType((*GroupInfo)(nil), "cosmos.orm.testdata.GroupInfo")
	proto.RegisterType((*GroupMember)(nil), "cosmos.orm.testdata.GroupMember")
}

func init() { proto.RegisterFile("cosmos/orm/testdata/codec.proto", fileDescriptor_c9806bf430fc0d7f) }

var fileDescriptor_c9806bf430fc0d7f = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x31, 0x4e, 0xf3, 0x30,
	0x14, 0xc7, 0xeb, 0xef, 0x2b, 0x85, 0xba, 0x4c, 0x06, 0xa1, 0xc0, 0xe0, 0x46, 0x9d, 0xc2, 0xd0,
	0x58, 0x88, 0x13, 0xb4, 0x0c, 0x55, 0x06, 0x96, 0x8c, 0x2c, 0x28, 0xb1, 0x4d, 0x6a, 0x21, 0xe7,
	0x45, 0xb6, 0x2b, 0xc4, 0x2d, 0xe0, 0x38, 0xdc, 0x80, 0xb1, 0x23, 0x13, 0x42, 0xc9, 0x2d, 0x98,
	0x50, 0x9c, 0x20, 0x75, 0x62, 0xe8, 0x64, 0x3f, 0xfb, 0xff, 0x7e, 0xfe, 0x59, 0x0f, 0x4f, 0x39,
	0x58, 0x0d, 0x96, 0x81, 0xd1, 0xcc, 0x49, 0xeb, 0x44, 0xe6, 0x32, 0xc6, 0x41, 0x48, 0x1e, 0x57,
	0x06, 0x1c, 0x90, 0x93, 0x2e, 0x10, 0x83, 0xd1, 0xf1, 0x6f, 0xe0, 0xe2, 0xb4, 0x80, 0x02, 0xfc,
	0x3d, 0x6b, 0x77, 0x5d, 0x74, 0xf6, 0x8a, 0xf0, 0x78, 0x65, 0x60, 0x53, 0x25, 0xe5, 0x03, 0x90,
	0x73, 0x7c, 0x54, 0xb4, 0xc5, 0xbd, 0x12, 0x01, 0x0a, 0x51, 0x34, 0x4c, 0x0f, 0x7d, 0x9d, 0x08,
	0x12, 0xe2, 0x89, 0x90, 0x96, 0x1b, 0x55, 0x39, 0x05, 0x65, 0xf0, 0x2f, 0x44, 0xd1, 0x38, 0xdd,
	0x3d, 0x22, 0x2b, 0x7c, 0x90, 0x09, 0xad, 0xca, 0xe0, 0x7f, 0x88, 0xa2, 0xe3, 0xe5, 0xd5, 0xf7,
	0xe7, 0x74, 0x5e, 0x28, 0xb7, 0xde, 0xe4, 0x31, 0x07, 0xcd, 0x7a, 0xe9, 0x6e, 0x99, 0x5b, 0xf1,
	0xc8, 0xdc, 0x73, 0x25, 0x6d, 0xbc, 0xe0, 0x7c, 0x21, 0x84, 0x91, 0xd6, 0xa6, 0x5d, 0xff, 0xec,
	0x0d, 0xe1, 0x89, 0x77, 0xba, 0x95, 0x3a, 0x97, 0xa6, 0x05, 0x7b, 0x0b, 0xaf, 0xb4, 0x1f, 0xd8,
	0xf7, 0x93, 0x04, 0x8f, 0xb4, 0x47, 0x7a, 0xfd, 0xbd, 0x48, 0x3d, 0x80, 0x9c, 0xe1, 0xd1, 0x93,
	0x54, 0xc5, 0xda, 0xf9, 0xdf, 0x0e, 0xd3, 0xbe, 0x5a, 0xde, 0xbc, 0xd7, 0x14, 0x6d, 0x6b, 0x8a,
	0xbe, 0x6a, 0x8a, 0x5e, 0x1a, 0x3a, 0xd8, 0x36, 0x74, 0xf0, 0xd1, 0xd0, 0xc1, 0xdd, 0xe5, 0x9f,
	0x0f, 0xed, 0xce, 0x32, 0x1f, 0xf9, 0xd9, 0x5c, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xcf, 0xdb,
	0x7e, 0x7a, 0xe9, 0x01, 0x00, 0x00,
}

func (m *GroupInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GroupInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GroupInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Admin) > 0 {
		i -= len(m.Admin)
		copy(dAtA[i:], m.Admin)
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Admin)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if m.GroupId != 0 {
		i = encodeVarintCodec(dAtA, i, uint64(m.GroupId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GroupMember) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GroupMember) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GroupMember) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Weight != 0 {
		i = encodeVarintCodec(dAtA, i, uint64(m.Weight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Member) > 0 {
		i -= len(m.Member)
		copy(dAtA[i:], m.Member)
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Member)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Group) > 0 {
		i -= len(m.Group)
		copy(dAtA[i:], m.Group)
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Group)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCodec(dAtA []byte, offset int, v uint64) int {
	offset -= sovCodec(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GroupInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GroupId != 0 {
		n += 1 + sovCodec(uint64(m.GroupId))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Admin)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func (m *GroupMember) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Group)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Member)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if m.Weight != 0 {
		n += 1 + sovCodec(uint64(m.Weight))
	}
	return n
}

func sovCodec(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCodec(x uint64) (n int) {
	return sovCodec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GroupInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: GroupInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
			}
			m.GroupId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroupId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Admin = append(m.Admin[:0], dAtA[iNdEx:postIndex]...)
			if m.Admin == nil {
				m.Admin = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCodec
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
func (m *GroupMember) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: GroupMember: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupMember: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Group", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Group = append(m.Group[:0], dAtA[iNdEx:postIndex]...)
			if m.Group == nil {
				m.Group = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Member", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Member = append(m.Member[:0], dAtA[iNdEx:postIndex]...)
			if m.Member == nil {
				m.Member = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			m.Weight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Weight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCodec
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
func skipCodec(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCodec
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
					return 0, ErrIntOverflowCodec
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
					return 0, ErrIntOverflowCodec
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
				return 0, ErrInvalidLengthCodec
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCodec
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCodec
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCodec        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCodec          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCodec = fmt.Errorf("proto: unexpected end of group")
)