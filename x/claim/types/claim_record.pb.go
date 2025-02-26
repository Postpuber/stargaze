// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stargaze/claim/v1beta1/claim_record.proto

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

type Action int32

const (
	ActionInitialClaim   Action = 0
	ActionBuySocialToken Action = 1
	ActionMintNFT        Action = 2
	ActionVote           Action = 3
	ActionDelegateStake  Action = 4
)

var Action_name = map[int32]string{
	0: "ActionInitialClaim",
	1: "ActionBuySocialToken",
	2: "ActionMintNFT",
	3: "ActionVote",
	4: "ActionDelegateStake",
}

var Action_value = map[string]int32{
	"ActionInitialClaim":   0,
	"ActionBuySocialToken": 1,
	"ActionMintNFT":        2,
	"ActionVote":           3,
	"ActionDelegateStake":  4,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}

func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8a75b9157744df4f, []int{0}
}

type ClaimRecord struct {
	// address of claim user
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	// total initial claimable amount for the user
	InitialClaimableAmount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=initial_claimable_amount,json=initialClaimableAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"initial_claimable_amount" yaml:"initial_claimable_amount"`
	// true if action is completed
	// index of bool in array refers to action enum #
	ActionCompleted []bool `protobuf:"varint,4,rep,packed,name=action_completed,json=actionCompleted,proto3" json:"action_completed,omitempty" yaml:"action_completed"`
}

func (m *ClaimRecord) Reset()         { *m = ClaimRecord{} }
func (m *ClaimRecord) String() string { return proto.CompactTextString(m) }
func (*ClaimRecord) ProtoMessage()    {}
func (*ClaimRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a75b9157744df4f, []int{0}
}
func (m *ClaimRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimRecord.Merge(m, src)
}
func (m *ClaimRecord) XXX_Size() int {
	return m.Size()
}
func (m *ClaimRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimRecord proto.InternalMessageInfo

func (m *ClaimRecord) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ClaimRecord) GetInitialClaimableAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.InitialClaimableAmount
	}
	return nil
}

func (m *ClaimRecord) GetActionCompleted() []bool {
	if m != nil {
		return m.ActionCompleted
	}
	return nil
}

func init() {
	proto.RegisterEnum("publicawesome.stargaze.claim.v1beta1.Action", Action_name, Action_value)
	proto.RegisterType((*ClaimRecord)(nil), "publicawesome.stargaze.claim.v1beta1.ClaimRecord")
}

func init() {
	proto.RegisterFile("stargaze/claim/v1beta1/claim_record.proto", fileDescriptor_8a75b9157744df4f)
}

var fileDescriptor_8a75b9157744df4f = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xb5, 0x93, 0xa8, 0xc0, 0x56, 0x14, 0xb3, 0x54, 0xad, 0xc9, 0xc1, 0x8e, 0x2c, 0x0e, 0x01,
	0x51, 0x9b, 0xc2, 0x8d, 0x5b, 0x1d, 0x84, 0x44, 0x25, 0x38, 0x24, 0x15, 0x07, 0x2e, 0xd1, 0xda,
	0x1e, 0x99, 0x55, 0x6c, 0x4f, 0xe4, 0x5d, 0x03, 0x41, 0x7c, 0x00, 0x47, 0xfe, 0x81, 0x0b, 0xe2,
	0xc8, 0x57, 0xf4, 0xd8, 0x23, 0xa7, 0x80, 0x92, 0x3f, 0xe8, 0x17, 0xa0, 0xec, 0xae, 0x21, 0x42,
	0xea, 0x69, 0x77, 0xdf, 0xcc, 0x7b, 0xf3, 0x76, 0xf4, 0xc8, 0x7d, 0x21, 0x59, 0x9d, 0xb3, 0x8f,
	0x10, 0xa5, 0x05, 0xe3, 0x65, 0xf4, 0xee, 0x38, 0x01, 0xc9, 0x8e, 0xf5, 0x6b, 0x5a, 0x43, 0x8a,
	0x75, 0x16, 0xce, 0x6b, 0x94, 0x48, 0xef, 0xcd, 0x9b, 0xa4, 0xe0, 0x29, 0x7b, 0x0f, 0x02, 0x4b,
	0x08, 0x5b, 0x62, 0xa8, 0x5a, 0x43, 0x43, 0xec, 0xef, 0xe7, 0x98, 0xa3, 0x22, 0x44, 0x9b, 0x9b,
	0xe6, 0xf6, 0xbd, 0x14, 0x45, 0x89, 0x22, 0x4a, 0x98, 0x80, 0x7f, 0x33, 0x90, 0x57, 0xba, 0x1e,
	0xfc, 0xe8, 0x90, 0xdd, 0xd1, 0x46, 0x67, 0xac, 0x26, 0xd2, 0x87, 0xe4, 0x1a, 0xcb, 0xb2, 0x1a,
	0x84, 0x70, 0xed, 0x81, 0x3d, 0xbc, 0x11, 0xd3, 0xcb, 0xa5, 0xbf, 0xb7, 0x60, 0x65, 0xf1, 0x34,
	0x30, 0x85, 0x60, 0xdc, 0xb6, 0xd0, 0x6f, 0x36, 0x71, 0x79, 0xc5, 0x25, 0x67, 0xc5, 0x54, 0xb9,
	0x61, 0x49, 0x01, 0x53, 0x56, 0x62, 0x53, 0x49, 0xb7, 0x33, 0xe8, 0x0e, 0x77, 0x1f, 0xdf, 0x0d,
	0xb5, 0x83, 0x70, 0xe3, 0xa0, 0x35, 0x1b, 0x8e, 0x90, 0x57, 0xf1, 0xe4, 0x7c, 0xe9, 0x5b, 0x97,
	0x4b, 0xdf, 0xd7, 0xf2, 0x57, 0x09, 0x05, 0xdf, 0x7f, 0xf9, 0xc3, 0x9c, 0xcb, 0xb7, 0x4d, 0x12,
	0xa6, 0x58, 0x46, 0xe6, 0x47, 0xfa, 0x38, 0x12, 0xd9, 0x2c, 0x92, 0x8b, 0x39, 0x08, 0xa5, 0x29,
	0xc6, 0x07, 0x46, 0x66, 0xd4, 0xaa, 0x9c, 0x28, 0x11, 0x7a, 0x4a, 0x1c, 0x96, 0x4a, 0x8e, 0xd5,
	0x34, 0xc5, 0x72, 0x5e, 0x80, 0x84, 0xcc, 0xed, 0x0d, 0xba, 0xc3, 0xeb, 0xb1, 0x6f, 0x6c, 0x1c,
	0x9a, 0x5f, 0xfe, 0xd7, 0x15, 0x8c, 0x6f, 0x69, 0x68, 0xd4, 0x22, 0x0f, 0x3e, 0x91, 0x9d, 0x13,
	0x05, 0xd1, 0x03, 0x42, 0xf5, 0xed, 0xc5, 0xd6, 0x54, 0xc7, 0xa2, 0x2e, 0xd9, 0xd7, 0x78, 0xdc,
	0x2c, 0x26, 0x98, 0x72, 0x56, 0x9c, 0xe1, 0x0c, 0x2a, 0xc7, 0xa6, 0xb7, 0xc9, 0x4d, 0x5d, 0x79,
	0xc9, 0x2b, 0xf9, 0xea, 0xf9, 0x99, 0xd3, 0xa1, 0x7b, 0x84, 0x68, 0xe8, 0x35, 0x4a, 0x70, 0xba,
	0xf4, 0x90, 0xdc, 0xd1, 0xef, 0x67, 0x50, 0x40, 0xce, 0x24, 0x4c, 0x24, 0x9b, 0x81, 0xd3, 0xeb,
	0xf7, 0x3e, 0x7f, 0xf5, 0xac, 0xf8, 0xf4, 0x7c, 0xe5, 0xd9, 0x17, 0x2b, 0xcf, 0xfe, 0xbd, 0xf2,
	0xec, 0x2f, 0x6b, 0xcf, 0xba, 0x58, 0x7b, 0xd6, 0xcf, 0xb5, 0x67, 0xbd, 0x79, 0xb4, 0xb5, 0x25,
	0x9d, 0x99, 0x23, 0x13, 0x9a, 0xe8, 0x6f, 0xda, 0x3e, 0x98, 0xbc, 0xa9, 0x9d, 0x25, 0x3b, 0x2a,
	0x05, 0x4f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x95, 0x86, 0xca, 0x8e, 0x02, 0x00, 0x00,
}

func (m *ClaimRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ActionCompleted) > 0 {
		for iNdEx := len(m.ActionCompleted) - 1; iNdEx >= 0; iNdEx-- {
			i--
			if m.ActionCompleted[iNdEx] {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
		}
		i = encodeVarintClaimRecord(dAtA, i, uint64(len(m.ActionCompleted)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.InitialClaimableAmount) > 0 {
		for iNdEx := len(m.InitialClaimableAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InitialClaimableAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintClaimRecord(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintClaimRecord(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintClaimRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovClaimRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClaimRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovClaimRecord(uint64(l))
	}
	if len(m.InitialClaimableAmount) > 0 {
		for _, e := range m.InitialClaimableAmount {
			l = e.Size()
			n += 1 + l + sovClaimRecord(uint64(l))
		}
	}
	if len(m.ActionCompleted) > 0 {
		n += 1 + sovClaimRecord(uint64(len(m.ActionCompleted))) + len(m.ActionCompleted)*1
	}
	return n
}

func sovClaimRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClaimRecord(x uint64) (n int) {
	return sovClaimRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClaimRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaimRecord
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
			return fmt.Errorf("proto: ClaimRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimRecord
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
				return ErrInvalidLengthClaimRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaimRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialClaimableAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimRecord
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
				return ErrInvalidLengthClaimRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaimRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialClaimableAmount = append(m.InitialClaimableAmount, types.Coin{})
			if err := m.InitialClaimableAmount[len(m.InitialClaimableAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType == 0 {
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaimRecord
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ActionCompleted = append(m.ActionCompleted, bool(v != 0))
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaimRecord
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClaimRecord
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthClaimRecord
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen
				if elementCount != 0 && len(m.ActionCompleted) == 0 {
					m.ActionCompleted = make([]bool, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaimRecord
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ActionCompleted = append(m.ActionCompleted, bool(v != 0))
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionCompleted", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClaimRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaimRecord
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
func skipClaimRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClaimRecord
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
					return 0, ErrIntOverflowClaimRecord
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
					return 0, ErrIntOverflowClaimRecord
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
				return 0, ErrInvalidLengthClaimRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClaimRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClaimRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClaimRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClaimRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClaimRecord = fmt.Errorf("proto: unexpected end of group")
)
