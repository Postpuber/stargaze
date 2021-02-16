// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stargaze/user/v1beta1/user.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

type Vouch struct {
	Voucher string `protobuf:"bytes,1,opt,name=voucher,proto3" json:"voucher,omitempty" yaml:"voucher"`
	Vouched string `protobuf:"bytes,2,opt,name=vouched,proto3" json:"vouched,omitempty" yaml:"vouched"`
	Comment string `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty" yaml:"comment"`
}

func (m *Vouch) Reset()         { *m = Vouch{} }
func (m *Vouch) String() string { return proto.CompactTextString(m) }
func (*Vouch) ProtoMessage()    {}
func (*Vouch) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4a98d2d6b2606be, []int{0}
}
func (m *Vouch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vouch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vouch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Vouch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vouch.Merge(m, src)
}
func (m *Vouch) XXX_Size() int {
	return m.Size()
}
func (m *Vouch) XXX_DiscardUnknown() {
	xxx_messageInfo_Vouch.DiscardUnknown(m)
}

var xxx_messageInfo_Vouch proto.InternalMessageInfo

func (m *Vouch) GetVoucher() string {
	if m != nil {
		return m.Voucher
	}
	return ""
}

func (m *Vouch) GetVouched() string {
	if m != nil {
		return m.Vouched
	}
	return ""
}

func (m *Vouch) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type Params struct {
	ThresholdAmount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=threshold_amount,json=thresholdAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"threshold_amount" yaml:"threshold_amount"`
	VouchCount      uint32                                   `protobuf:"varint,2,opt,name=vouch_count,json=vouchCount,proto3" json:"vouch_count,omitempty" yaml:"vouch_count"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4a98d2d6b2606be, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetThresholdAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.ThresholdAmount
	}
	return nil
}

func (m *Params) GetVouchCount() uint32 {
	if m != nil {
		return m.VouchCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Vouch)(nil), "stargaze.user.v1beta1.Vouch")
	proto.RegisterType((*Params)(nil), "stargaze.user.v1beta1.Params")
}

func init() { proto.RegisterFile("stargaze/user/v1beta1/user.proto", fileDescriptor_c4a98d2d6b2606be) }

var fileDescriptor_c4a98d2d6b2606be = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xbb, 0x4e, 0xeb, 0x40,
	0x10, 0x86, 0xbd, 0x27, 0x87, 0x20, 0x1c, 0x71, 0x91, 0xc5, 0xc5, 0xa4, 0xb0, 0x23, 0x57, 0x29,
	0x88, 0x57, 0x81, 0x02, 0x29, 0x1d, 0x4e, 0x85, 0x68, 0x90, 0x0b, 0x0a, 0x9a, 0x68, 0x6d, 0xaf,
	0x6c, 0x8b, 0xac, 0x37, 0xf2, 0xae, 0x03, 0xe1, 0x29, 0x48, 0x47, 0x49, 0xcd, 0x93, 0xa4, 0x4c,
	0x89, 0x84, 0x64, 0x50, 0xf2, 0x06, 0x79, 0x02, 0xe4, 0xf5, 0x85, 0x28, 0xa2, 0xf2, 0x8c, 0xff,
	0x6f, 0xc6, 0xf3, 0x8f, 0x47, 0x6e, 0x31, 0x8e, 0x62, 0x1f, 0x3d, 0x63, 0x98, 0x30, 0x1c, 0xc3,
	0x71, 0xd7, 0xc1, 0x1c, 0x75, 0x45, 0x62, 0x8e, 0x62, 0xca, 0xa9, 0x72, 0x54, 0x12, 0xa6, 0x78,
	0x59, 0x10, 0xcd, 0x43, 0x9f, 0xfa, 0x54, 0x10, 0x30, 0x8b, 0x72, 0xb8, 0xa9, 0xb9, 0x94, 0x11,
	0xca, 0xa0, 0x83, 0x18, 0xae, 0x9a, 0xb9, 0x34, 0x8c, 0x72, 0xdd, 0x98, 0x02, 0x79, 0xeb, 0x8e,
	0x26, 0x6e, 0xa0, 0x9c, 0xc9, 0xdb, 0xe3, 0x2c, 0xc0, 0xb1, 0x0a, 0x5a, 0xa0, 0xbd, 0x63, 0x29,
	0xab, 0x54, 0xdf, 0x9b, 0x20, 0x32, 0xec, 0x19, 0x85, 0x60, 0xd8, 0x25, 0xf2, 0x4b, 0x7b, 0xea,
	0xbf, 0xbf, 0x69, 0xaf, 0xa2, 0xbd, 0x8c, 0x76, 0x29, 0x21, 0x38, 0xe2, 0x6a, 0x6d, 0x93, 0x2e,
	0x04, 0xc3, 0x2e, 0x11, 0xe3, 0x13, 0xc8, 0xf5, 0x5b, 0x14, 0x23, 0xc2, 0x94, 0x29, 0x90, 0x0f,
	0x78, 0x10, 0x63, 0x16, 0xd0, 0xa1, 0x37, 0x40, 0x84, 0x26, 0x11, 0x57, 0x41, 0xab, 0xd6, 0x6e,
	0x9c, 0x9f, 0x9a, 0xb9, 0x35, 0x33, 0xb3, 0x56, 0x6e, 0xc1, 0xec, 0xd3, 0x30, 0xb2, 0x6e, 0x66,
	0xa9, 0x2e, 0xad, 0x52, 0xfd, 0x24, 0xff, 0xc2, 0x66, 0x03, 0xe3, 0xfd, 0x4b, 0x6f, 0xfb, 0x21,
	0x0f, 0x12, 0xc7, 0x74, 0x29, 0x81, 0xc5, 0x8a, 0xf2, 0x47, 0x87, 0x79, 0x0f, 0x90, 0x4f, 0x46,
	0x98, 0x89, 0x5e, 0xcc, 0xde, 0xaf, 0xca, 0xaf, 0x44, 0xb5, 0x72, 0x29, 0x37, 0x84, 0xaf, 0x81,
	0x2b, 0xa6, 0xc9, 0xec, 0xef, 0x5a, 0xc7, 0xab, 0x54, 0x57, 0xd6, 0xec, 0xe7, 0xa2, 0x61, 0xcb,
	0x22, 0xeb, 0x67, 0x49, 0xef, 0xff, 0xeb, 0x9b, 0x2e, 0x59, 0xd7, 0xb3, 0x85, 0x06, 0xe6, 0x0b,
	0x0d, 0x7c, 0x2f, 0x34, 0xf0, 0xb2, 0xd4, 0xa4, 0xf9, 0x52, 0x93, 0x3e, 0x96, 0x9a, 0x74, 0x0f,
	0xd7, 0x66, 0x1a, 0x25, 0xce, 0x30, 0x74, 0x3b, 0xe8, 0x11, 0x33, 0x4a, 0x30, 0xac, 0x8e, 0xe2,
	0x29, 0x3f, 0x0b, 0x31, 0xa0, 0x53, 0x17, 0xff, 0xf0, 0xe2, 0x27, 0x00, 0x00, 0xff, 0xff, 0x5b,
	0xab, 0x0c, 0x98, 0x34, 0x02, 0x00, 0x00,
}

func (m *Vouch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vouch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vouch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Comment) > 0 {
		i -= len(m.Comment)
		copy(dAtA[i:], m.Comment)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Comment)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Vouched) > 0 {
		i -= len(m.Vouched)
		copy(dAtA[i:], m.Vouched)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Vouched)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Voucher) > 0 {
		i -= len(m.Voucher)
		copy(dAtA[i:], m.Voucher)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Voucher)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VouchCount != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.VouchCount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ThresholdAmount) > 0 {
		for iNdEx := len(m.ThresholdAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ThresholdAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUser(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintUser(dAtA []byte, offset int, v uint64) int {
	offset -= sovUser(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Vouch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Voucher)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Vouched)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Comment)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ThresholdAmount) > 0 {
		for _, e := range m.ThresholdAmount {
			l = e.Size()
			n += 1 + l + sovUser(uint64(l))
		}
	}
	if m.VouchCount != 0 {
		n += 1 + sovUser(uint64(m.VouchCount))
	}
	return n
}

func sovUser(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUser(x uint64) (n int) {
	return sovUser(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Vouch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: Vouch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vouch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voucher", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Voucher = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vouched", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vouched = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Comment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Comment = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUser
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ThresholdAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ThresholdAmount = append(m.ThresholdAmount, types.Coin{})
			if err := m.ThresholdAmount[len(m.ThresholdAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VouchCount", wireType)
			}
			m.VouchCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VouchCount |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUser
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
func skipUser(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
				return 0, ErrInvalidLengthUser
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUser
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUser
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUser        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUser          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUser = fmt.Errorf("proto: unexpected end of group")
)
