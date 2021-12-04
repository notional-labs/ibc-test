// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/core/channel/v1/channel_id.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/ibc-go/modules/core/02-client/types"
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

// channel id and its respective port id
type ChannelIdentifiers struct {
	PortIds    []string `protobuf:"bytes,1,rep,name=port_ids,json=portIds,proto3" json:"port_ids,omitempty"`
	ChannelIds []string `protobuf:"bytes,2,rep,name=channel_ids,json=channelIds,proto3" json:"channel_ids,omitempty"`
}

func (m *ChannelIdentifiers) Reset()         { *m = ChannelIdentifiers{} }
func (m *ChannelIdentifiers) String() string { return proto.CompactTextString(m) }
func (*ChannelIdentifiers) ProtoMessage()    {}
func (*ChannelIdentifiers) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e000420c0d52bf5, []int{0}
}
func (m *ChannelIdentifiers) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChannelIdentifiers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChannelIdentifiers.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChannelIdentifiers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelIdentifiers.Merge(m, src)
}
func (m *ChannelIdentifiers) XXX_Size() int {
	return m.Size()
}
func (m *ChannelIdentifiers) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelIdentifiers.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelIdentifiers proto.InternalMessageInfo

func (m *ChannelIdentifiers) GetPortIds() []string {
	if m != nil {
		return m.PortIds
	}
	return nil
}

func (m *ChannelIdentifiers) GetChannelIds() []string {
	if m != nil {
		return m.ChannelIds
	}
	return nil
}

func init() {
	proto.RegisterType((*ChannelIdentifiers)(nil), "ibc.core.channel.v1.ChannelIdentifiers")
}

func init() {
	proto.RegisterFile("ibc/core/channel/v1/channel_id.proto", fileDescriptor_3e000420c0d52bf5)
}

var fileDescriptor_3e000420c0d52bf5 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc9, 0x4c, 0x4a, 0xd6,
	0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0xce, 0x48, 0xcc, 0xcb, 0x4b, 0xcd, 0xd1, 0x2f, 0x33, 0x84,
	0x31, 0xe3, 0x33, 0x53, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x84, 0x33, 0x93, 0x92, 0xf5,
	0x40, 0xaa, 0xf4, 0xa0, 0x52, 0x7a, 0x65, 0x86, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x79,
	0x7d, 0x10, 0x0b, 0xa2, 0x54, 0x4a, 0x1e, 0x61, 0x60, 0x4e, 0x66, 0x6a, 0x5e, 0x09, 0xd8, 0x3c,
	0x30, 0x0b, 0xa2, 0x40, 0x29, 0x80, 0x4b, 0xc8, 0x19, 0x62, 0x88, 0x67, 0x4a, 0x6a, 0x5e, 0x49,
	0x66, 0x5a, 0x66, 0x6a, 0x51, 0xb1, 0x90, 0x24, 0x17, 0x47, 0x41, 0x7e, 0x51, 0x49, 0x7c, 0x66,
	0x4a, 0xb1, 0x04, 0xa3, 0x02, 0xb3, 0x06, 0x67, 0x10, 0x3b, 0x88, 0xef, 0x99, 0x52, 0x2c, 0x24,
	0xcf, 0xc5, 0x8d, 0x70, 0x50, 0xb1, 0x04, 0x13, 0x58, 0x96, 0x2b, 0x19, 0x66, 0x46, 0xb1, 0x53,
	0xc0, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1,
	0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x99, 0xa5, 0x67, 0x96, 0x64,
	0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x27, 0xe7, 0x17, 0xe7, 0xe6, 0x17, 0xeb, 0x67, 0x26,
	0x25, 0xeb, 0xa6, 0xe7, 0xeb, 0xe7, 0xe6, 0xa7, 0x94, 0xe6, 0xa4, 0x16, 0x43, 0x5c, 0x6a, 0x60,
	0xa2, 0x0b, 0xf3, 0x7d, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0xa9, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xd2, 0xf0, 0x8c, 0xd2, 0x1e, 0x01, 0x00, 0x00,
}

func (m *ChannelIdentifiers) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChannelIdentifiers) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChannelIdentifiers) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChannelIds) > 0 {
		for iNdEx := len(m.ChannelIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ChannelIds[iNdEx])
			copy(dAtA[i:], m.ChannelIds[iNdEx])
			i = encodeVarintChannelId(dAtA, i, uint64(len(m.ChannelIds[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.PortIds) > 0 {
		for iNdEx := len(m.PortIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PortIds[iNdEx])
			copy(dAtA[i:], m.PortIds[iNdEx])
			i = encodeVarintChannelId(dAtA, i, uint64(len(m.PortIds[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintChannelId(dAtA []byte, offset int, v uint64) int {
	offset -= sovChannelId(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChannelIdentifiers) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PortIds) > 0 {
		for _, s := range m.PortIds {
			l = len(s)
			n += 1 + l + sovChannelId(uint64(l))
		}
	}
	if len(m.ChannelIds) > 0 {
		for _, s := range m.ChannelIds {
			l = len(s)
			n += 1 + l + sovChannelId(uint64(l))
		}
	}
	return n
}

func sovChannelId(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozChannelId(x uint64) (n int) {
	return sovChannelId(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChannelIdentifiers) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChannelId
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
			return fmt.Errorf("proto: ChannelIdentifiers: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChannelIdentifiers: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortIds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChannelId
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
				return ErrInvalidLengthChannelId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChannelId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortIds = append(m.PortIds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelIds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChannelId
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
				return ErrInvalidLengthChannelId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChannelId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelIds = append(m.ChannelIds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChannelId(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChannelId
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
func skipChannelId(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChannelId
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
					return 0, ErrIntOverflowChannelId
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
					return 0, ErrIntOverflowChannelId
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
				return 0, ErrInvalidLengthChannelId
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupChannelId
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthChannelId
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthChannelId        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChannelId          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupChannelId = fmt.Errorf("proto: unexpected end of group")
)
