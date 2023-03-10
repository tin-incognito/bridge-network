// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bridge/bridge/pub_key_set.proto

package types

import (
	fmt "fmt"
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

type PubKeySet struct {
	Secp256K1 string `protobuf:"bytes,1,opt,name=secp256k1,proto3" json:"secp256k1,omitempty"`
	Ed25519   string `protobuf:"bytes,2,opt,name=ed25519,proto3" json:"ed25519,omitempty"`
}

func (m *PubKeySet) Reset()         { *m = PubKeySet{} }
func (m *PubKeySet) String() string { return proto.CompactTextString(m) }
func (*PubKeySet) ProtoMessage()    {}
func (*PubKeySet) Descriptor() ([]byte, []int) {
	return fileDescriptor_91677ddf6bc05d5a, []int{0}
}
func (m *PubKeySet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PubKeySet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PubKeySet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PubKeySet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubKeySet.Merge(m, src)
}
func (m *PubKeySet) XXX_Size() int {
	return m.Size()
}
func (m *PubKeySet) XXX_DiscardUnknown() {
	xxx_messageInfo_PubKeySet.DiscardUnknown(m)
}

var xxx_messageInfo_PubKeySet proto.InternalMessageInfo

func (m *PubKeySet) GetSecp256K1() string {
	if m != nil {
		return m.Secp256K1
	}
	return ""
}

func (m *PubKeySet) GetEd25519() string {
	if m != nil {
		return m.Ed25519
	}
	return ""
}

func init() {
	proto.RegisterType((*PubKeySet)(nil), "bridge.bridge.PubKeySet")
}

func init() { proto.RegisterFile("bridge/bridge/pub_key_set.proto", fileDescriptor_91677ddf6bc05d5a) }

var fileDescriptor_91677ddf6bc05d5a = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x2a, 0xca, 0x4c,
	0x49, 0x4f, 0xd5, 0x87, 0x52, 0x05, 0xa5, 0x49, 0xf1, 0xd9, 0xa9, 0x95, 0xf1, 0xc5, 0xa9, 0x25,
	0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xbc, 0x10, 0x19, 0x3d, 0x08, 0xa5, 0xe4, 0xcc, 0xc5,
	0x19, 0x50, 0x9a, 0xe4, 0x9d, 0x5a, 0x19, 0x9c, 0x5a, 0x22, 0x24, 0xc3, 0xc5, 0x59, 0x9c, 0x9a,
	0x5c, 0x60, 0x64, 0x6a, 0x96, 0x6d, 0x28, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x10, 0x10,
	0x92, 0xe0, 0x62, 0x4f, 0x4d, 0x31, 0x32, 0x35, 0x35, 0xb4, 0x94, 0x60, 0x02, 0xcb, 0xc1, 0xb8,
	0x4e, 0xfa, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84,
	0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x25, 0x0a, 0x75, 0x47,
	0x05, 0xcc, 0x41, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0xb7, 0x18, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x2e, 0x1f, 0x24, 0xf4, 0xae, 0x00, 0x00, 0x00,
}

func (m *PubKeySet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PubKeySet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PubKeySet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Ed25519) > 0 {
		i -= len(m.Ed25519)
		copy(dAtA[i:], m.Ed25519)
		i = encodeVarintPubKeySet(dAtA, i, uint64(len(m.Ed25519)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Secp256K1) > 0 {
		i -= len(m.Secp256K1)
		copy(dAtA[i:], m.Secp256K1)
		i = encodeVarintPubKeySet(dAtA, i, uint64(len(m.Secp256K1)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPubKeySet(dAtA []byte, offset int, v uint64) int {
	offset -= sovPubKeySet(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PubKeySet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Secp256K1)
	if l > 0 {
		n += 1 + l + sovPubKeySet(uint64(l))
	}
	l = len(m.Ed25519)
	if l > 0 {
		n += 1 + l + sovPubKeySet(uint64(l))
	}
	return n
}

func sovPubKeySet(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPubKeySet(x uint64) (n int) {
	return sovPubKeySet(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PubKeySet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPubKeySet
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
			return fmt.Errorf("proto: PubKeySet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PubKeySet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secp256K1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPubKeySet
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
				return ErrInvalidLengthPubKeySet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPubKeySet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Secp256K1 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ed25519", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPubKeySet
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
				return ErrInvalidLengthPubKeySet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPubKeySet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ed25519 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPubKeySet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPubKeySet
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
func skipPubKeySet(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPubKeySet
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
					return 0, ErrIntOverflowPubKeySet
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
					return 0, ErrIntOverflowPubKeySet
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
				return 0, ErrInvalidLengthPubKeySet
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPubKeySet
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPubKeySet
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPubKeySet        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPubKeySet          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPubKeySet = fmt.Errorf("proto: unexpected end of group")
)
