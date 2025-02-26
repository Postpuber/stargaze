// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stargaze/claim/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the claim module's genesis state.
type GenesisState struct {
	// this line is used by starport scaffolding # genesis/proto/state
	// balance of the claim module's account
	ModuleAccountBalance types.Coin `protobuf:"bytes,1,opt,name=module_account_balance,json=moduleAccountBalance,proto3" json:"module_account_balance" yaml:"module_account_balance"`
	// params defines all the parameters of the module.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params" yaml:"params"`
	// list of claim records, one for every airdrop recipient
	ClaimRecords []ClaimRecord `protobuf:"bytes,3,rep,name=claim_records,json=claimRecords,proto3" json:"claim_records" yaml:"claim_records"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac1d615666a45cf, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetModuleAccountBalance() types.Coin {
	if m != nil {
		return m.ModuleAccountBalance
	}
	return types.Coin{}
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetClaimRecords() []ClaimRecord {
	if m != nil {
		return m.ClaimRecords
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "publicawesome.stargaze.claim.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("stargaze/claim/v1beta1/genesis.proto", fileDescriptor_cac1d615666a45cf)
}

var fileDescriptor_cac1d615666a45cf = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0x5b, 0x48, 0x58, 0x14, 0xd8, 0x34, 0xdc, 0x1b, 0x2e, 0xb9, 0x0e, 0xa4, 0x62, 0x82,
	0x89, 0xce, 0x08, 0xee, 0xdc, 0x59, 0x16, 0x26, 0xae, 0x4c, 0xdd, 0xe9, 0x82, 0x4c, 0x87, 0x49,
	0x6d, 0xd2, 0xe9, 0x34, 0x9d, 0x29, 0x8a, 0x4f, 0xe1, 0x63, 0xb1, 0x64, 0xa9, 0x1b, 0x62, 0xe0,
	0x0d, 0x7c, 0x02, 0xc3, 0xcc, 0x80, 0x9a, 0x48, 0xc2, 0xae, 0x73, 0xce, 0xf9, 0xff, 0xff, 0xeb,
	0x39, 0x4e, 0x57, 0x48, 0x9c, 0x47, 0xf8, 0x99, 0x22, 0x92, 0xe0, 0x98, 0xa1, 0x49, 0x3f, 0xa4,
	0x12, 0xf7, 0x51, 0x44, 0x53, 0x2a, 0x62, 0x01, 0xb3, 0x9c, 0x4b, 0xee, 0x76, 0xb3, 0x22, 0x4c,
	0x62, 0x82, 0x1f, 0xa9, 0xe0, 0x8c, 0xc2, 0x8d, 0x06, 0x2a, 0x0d, 0x34, 0x9a, 0x56, 0x23, 0xe2,
	0x11, 0x57, 0x02, 0xb4, 0xfe, 0xd2, 0xda, 0x16, 0x20, 0x5c, 0x30, 0x2e, 0x50, 0x88, 0x05, 0xdd,
	0xda, 0x13, 0x1e, 0xa7, 0xa6, 0x7f, 0xbc, 0x83, 0x40, 0xbd, 0x46, 0x39, 0x25, 0x3c, 0x1f, 0x9b,
	0xd1, 0xc3, 0x1d, 0xa3, 0x19, 0xce, 0x31, 0x33, 0xac, 0xde, 0x5b, 0xc9, 0xa9, 0x5d, 0x69, 0xfa,
	0x5b, 0x89, 0x25, 0x75, 0x27, 0xce, 0x5f, 0xc6, 0xc7, 0x45, 0x42, 0x47, 0x98, 0x10, 0x5e, 0xa4,
	0x72, 0x14, 0xe2, 0x04, 0xa7, 0x84, 0x36, 0xed, 0x8e, 0xdd, 0xab, 0x0e, 0xfe, 0x41, 0x4d, 0x08,
	0xd7, 0x84, 0x9b, 0x9f, 0x81, 0x43, 0x1e, 0xa7, 0xfe, 0xd1, 0x6c, 0xd1, 0xb6, 0x3e, 0x16, 0xed,
	0x83, 0x29, 0x66, 0xc9, 0x85, 0xf7, 0xbb, 0x8d, 0x17, 0x34, 0x74, 0xe3, 0x52, 0xd7, 0x7d, 0x5d,
	0x76, 0xef, 0x9d, 0x8a, 0x06, 0x6b, 0x96, 0x54, 0xce, 0x09, 0xdc, 0x67, 0x8b, 0xf0, 0x46, 0x69,
	0xfc, 0x3f, 0x26, 0xba, 0xae, 0xa3, 0xb5, 0x93, 0x17, 0x18, 0x4b, 0x57, 0x3a, 0xf5, 0xef, 0x0b,
	0x12, 0xcd, 0x72, 0xa7, 0xdc, 0xab, 0x0e, 0xfa, 0xfb, 0x65, 0x0c, 0xd7, 0xaf, 0x40, 0x29, 0xfd,
	0xff, 0x26, 0xa8, 0xa1, 0x83, 0x7e, 0xb8, 0x7a, 0x41, 0x8d, 0x7c, 0x8d, 0x0a, 0xff, 0x7a, 0xb6,
	0x04, 0xf6, 0x7c, 0x09, 0xec, 0xf7, 0x25, 0xb0, 0x5f, 0x56, 0xc0, 0x9a, 0xaf, 0x80, 0xf5, 0xba,
	0x02, 0xd6, 0xdd, 0x59, 0x14, 0xcb, 0x87, 0x22, 0x84, 0x84, 0x33, 0xa4, 0x11, 0x4e, 0x0d, 0x03,
	0xda, 0x1e, 0xed, 0xc9, 0x9c, 0x4d, 0x4e, 0x33, 0x2a, 0xc2, 0x8a, 0x3a, 0xd7, 0xf9, 0x67, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xcc, 0x74, 0x9b, 0x0a, 0x82, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClaimRecords) > 0 {
		for iNdEx := len(m.ClaimRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClaimRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ModuleAccountBalance.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ModuleAccountBalance.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.ClaimRecords) > 0 {
		for _, e := range m.ClaimRecords {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleAccountBalance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ModuleAccountBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimRecords = append(m.ClaimRecords, ClaimRecord{})
			if err := m.ClaimRecords[len(m.ClaimRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
