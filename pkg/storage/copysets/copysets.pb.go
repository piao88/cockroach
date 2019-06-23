// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/copysets/copysets.proto

package copysets

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import github_com_cockroachdb_cockroach_pkg_roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"

import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// CopysetStrategy has the set of supported copyset-store allocation strategies.
type CopysetStrategy int32

const (
	// MAXIMIZE_DIVERSITY is a strategy which tries to maximize locality diversity
	// when creating copysets from a store list.
	CopysetStrategy_MAXIMIZE_DIVERSITY CopysetStrategy = 0
	// MINIMIZE_MOVEMENT is a strategy which tries to minimize changes to
	// existing copysets when generating new copysets on store list changes.
	// It does not guarantee optimal locality diversity but tries to avoid
	// stores with same localities within copysets.
	CopysetStrategy_MINIMIZE_MOVEMENT CopysetStrategy = 1
)

var CopysetStrategy_name = map[int32]string{
	0: "MAXIMIZE_DIVERSITY",
	1: "MINIMIZE_MOVEMENT",
}
var CopysetStrategy_value = map[string]int32{
	"MAXIMIZE_DIVERSITY": 0,
	"MINIMIZE_MOVEMENT":  1,
}

func (x CopysetStrategy) String() string {
	return proto.EnumName(CopysetStrategy_name, int32(x))
}
func (CopysetStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_copysets_b03c788fb5799fc5, []int{0}
}

// AllCopysets contains the map between replication factor to
// its copysets.
type AllCopysets struct {
	// Map from replication factors to copysets.
	ByRf map[int32]Copysets `protobuf:"bytes,1,rep,name=by_rf,json=byRf,proto3" json:"by_rf" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Strategy used for store-copyset allocation.
	Strategy             CopysetStrategy `protobuf:"varint,2,opt,name=strategy,proto3,enum=cockroach.storage.copysets.CopysetStrategy" json:"strategy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AllCopysets) Reset()         { *m = AllCopysets{} }
func (m *AllCopysets) String() string { return proto.CompactTextString(m) }
func (*AllCopysets) ProtoMessage()    {}
func (*AllCopysets) Descriptor() ([]byte, []int) {
	return fileDescriptor_copysets_b03c788fb5799fc5, []int{0}
}
func (m *AllCopysets) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllCopysets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *AllCopysets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllCopysets.Merge(dst, src)
}
func (m *AllCopysets) XXX_Size() int {
	return m.Size()
}
func (m *AllCopysets) XXX_DiscardUnknown() {
	xxx_messageInfo_AllCopysets.DiscardUnknown(m)
}

var xxx_messageInfo_AllCopysets proto.InternalMessageInfo

// Copysets contains copysets for a particular replication factor.
// If copysets based rebalancing is enabled, the replicas of a range will
// be contained within a copy set. Each store belongs to a single copyset.
// Copyset based rebalancing significantly improves failure tolerance.
type Copysets struct {
	// Map from CopysetID to a Copyset (set of stores in the copyset).
	Sets map[CopysetID]Copyset `protobuf:"bytes,1,rep,name=sets,proto3,castkey=CopysetID" json:"sets" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Replication factor of copy sets.
	ReplicationFactor    int32    `protobuf:"varint,2,opt,name=replication_factor,json=replicationFactor,proto3" json:"replication_factor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Copysets) Reset()         { *m = Copysets{} }
func (m *Copysets) String() string { return proto.CompactTextString(m) }
func (*Copysets) ProtoMessage()    {}
func (*Copysets) Descriptor() ([]byte, []int) {
	return fileDescriptor_copysets_b03c788fb5799fc5, []int{1}
}
func (m *Copysets) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Copysets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *Copysets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Copysets.Merge(dst, src)
}
func (m *Copysets) XXX_Size() int {
	return m.Size()
}
func (m *Copysets) XXX_DiscardUnknown() {
	xxx_messageInfo_Copysets.DiscardUnknown(m)
}

var xxx_messageInfo_Copysets proto.InternalMessageInfo

// Copyset contains the set of stores belonging to the same copyset.
type Copyset struct {
	// Map of StoreIDs.
	Ids                  map[github_com_cockroachdb_cockroach_pkg_roachpb.StoreID]bool `protobuf:"bytes,1,rep,name=ids,proto3,castkey=github.com/cockroachdb/cockroach/pkg/roachpb.StoreID" json:"ids,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                                      `json:"-"`
	XXX_sizecache        int32                                                         `json:"-"`
}

func (m *Copyset) Reset()         { *m = Copyset{} }
func (m *Copyset) String() string { return proto.CompactTextString(m) }
func (*Copyset) ProtoMessage()    {}
func (*Copyset) Descriptor() ([]byte, []int) {
	return fileDescriptor_copysets_b03c788fb5799fc5, []int{2}
}
func (m *Copyset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Copyset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *Copyset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Copyset.Merge(dst, src)
}
func (m *Copyset) XXX_Size() int {
	return m.Size()
}
func (m *Copyset) XXX_DiscardUnknown() {
	xxx_messageInfo_Copyset.DiscardUnknown(m)
}

var xxx_messageInfo_Copyset proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AllCopysets)(nil), "cockroach.storage.copysets.AllCopysets")
	proto.RegisterMapType((map[int32]Copysets)(nil), "cockroach.storage.copysets.AllCopysets.ByRfEntry")
	proto.RegisterType((*Copysets)(nil), "cockroach.storage.copysets.Copysets")
	proto.RegisterMapType((map[CopysetID]Copyset)(nil), "cockroach.storage.copysets.Copysets.SetsEntry")
	proto.RegisterType((*Copyset)(nil), "cockroach.storage.copysets.Copyset")
	proto.RegisterMapType((map[github_com_cockroachdb_cockroach_pkg_roachpb.StoreID]bool)(nil), "cockroach.storage.copysets.Copyset.IdsEntry")
	proto.RegisterEnum("cockroach.storage.copysets.CopysetStrategy", CopysetStrategy_name, CopysetStrategy_value)
}
func (m *AllCopysets) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllCopysets) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ByRf) > 0 {
		keysForByRf := make([]int32, 0, len(m.ByRf))
		for k := range m.ByRf {
			keysForByRf = append(keysForByRf, int32(k))
		}
		github_com_gogo_protobuf_sortkeys.Int32s(keysForByRf)
		for _, k := range keysForByRf {
			dAtA[i] = 0xa
			i++
			v := m.ByRf[int32(k)]
			msgSize := 0
			if (&v) != nil {
				msgSize = (&v).Size()
				msgSize += 1 + sovCopysets(uint64(msgSize))
			}
			mapSize := 1 + sovCopysets(uint64(k)) + msgSize
			i = encodeVarintCopysets(dAtA, i, uint64(mapSize))
			dAtA[i] = 0x8
			i++
			i = encodeVarintCopysets(dAtA, i, uint64(k))
			dAtA[i] = 0x12
			i++
			i = encodeVarintCopysets(dAtA, i, uint64((&v).Size()))
			n1, err := (&v).MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n1
		}
	}
	if m.Strategy != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintCopysets(dAtA, i, uint64(m.Strategy))
	}
	return i, nil
}

func (m *Copysets) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Copysets) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Sets) > 0 {
		keysForSets := make([]int32, 0, len(m.Sets))
		for k := range m.Sets {
			keysForSets = append(keysForSets, int32(k))
		}
		github_com_gogo_protobuf_sortkeys.Int32s(keysForSets)
		for _, k := range keysForSets {
			dAtA[i] = 0xa
			i++
			v := m.Sets[CopysetID(k)]
			msgSize := 0
			if (&v) != nil {
				msgSize = (&v).Size()
				msgSize += 1 + sovCopysets(uint64(msgSize))
			}
			mapSize := 1 + sovCopysets(uint64(k)) + msgSize
			i = encodeVarintCopysets(dAtA, i, uint64(mapSize))
			dAtA[i] = 0x8
			i++
			i = encodeVarintCopysets(dAtA, i, uint64(k))
			dAtA[i] = 0x12
			i++
			i = encodeVarintCopysets(dAtA, i, uint64((&v).Size()))
			n2, err := (&v).MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n2
		}
	}
	if m.ReplicationFactor != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintCopysets(dAtA, i, uint64(m.ReplicationFactor))
	}
	return i, nil
}

func (m *Copyset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Copyset) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Ids) > 0 {
		keysForIds := make([]int32, 0, len(m.Ids))
		for k := range m.Ids {
			keysForIds = append(keysForIds, int32(k))
		}
		github_com_gogo_protobuf_sortkeys.Int32s(keysForIds)
		for _, k := range keysForIds {
			dAtA[i] = 0xa
			i++
			v := m.Ids[github_com_cockroachdb_cockroach_pkg_roachpb.StoreID(k)]
			mapSize := 1 + sovCopysets(uint64(k)) + 1 + 1
			i = encodeVarintCopysets(dAtA, i, uint64(mapSize))
			dAtA[i] = 0x8
			i++
			i = encodeVarintCopysets(dAtA, i, uint64(k))
			dAtA[i] = 0x10
			i++
			if v {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i++
		}
	}
	return i, nil
}

func encodeVarintCopysets(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AllCopysets) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ByRf) > 0 {
		for k, v := range m.ByRf {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + sovCopysets(uint64(k)) + 1 + l + sovCopysets(uint64(l))
			n += mapEntrySize + 1 + sovCopysets(uint64(mapEntrySize))
		}
	}
	if m.Strategy != 0 {
		n += 1 + sovCopysets(uint64(m.Strategy))
	}
	return n
}

func (m *Copysets) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Sets) > 0 {
		for k, v := range m.Sets {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + sovCopysets(uint64(k)) + 1 + l + sovCopysets(uint64(l))
			n += mapEntrySize + 1 + sovCopysets(uint64(mapEntrySize))
		}
	}
	if m.ReplicationFactor != 0 {
		n += 1 + sovCopysets(uint64(m.ReplicationFactor))
	}
	return n
}

func (m *Copyset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Ids) > 0 {
		for k, v := range m.Ids {
			_ = k
			_ = v
			mapEntrySize := 1 + sovCopysets(uint64(k)) + 1 + 1
			n += mapEntrySize + 1 + sovCopysets(uint64(mapEntrySize))
		}
	}
	return n
}

func sovCopysets(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCopysets(x uint64) (n int) {
	return sovCopysets(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AllCopysets) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCopysets
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AllCopysets: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllCopysets: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ByRf", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCopysets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCopysets
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ByRf == nil {
				m.ByRf = make(map[int32]Copysets)
			}
			var mapkey int32
			mapvalue := &Copysets{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCopysets
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCopysets
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCopysets
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= (int(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthCopysets
					}
					postmsgIndex := iNdEx + mapmsglen
					if mapmsglen < 0 {
						return ErrInvalidLengthCopysets
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Copysets{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipCopysets(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthCopysets
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.ByRf[mapkey] = *mapvalue
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Strategy", wireType)
			}
			m.Strategy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCopysets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Strategy |= (CopysetStrategy(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCopysets(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCopysets
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
func (m *Copysets) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCopysets
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Copysets: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Copysets: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCopysets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCopysets
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Sets == nil {
				m.Sets = make(map[CopysetID]Copyset)
			}
			var mapkey int32
			mapvalue := &Copyset{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCopysets
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCopysets
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCopysets
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= (int(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthCopysets
					}
					postmsgIndex := iNdEx + mapmsglen
					if mapmsglen < 0 {
						return ErrInvalidLengthCopysets
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Copyset{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipCopysets(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthCopysets
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Sets[CopysetID(mapkey)] = *mapvalue
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplicationFactor", wireType)
			}
			m.ReplicationFactor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCopysets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReplicationFactor |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCopysets(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCopysets
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
func (m *Copyset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCopysets
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Copyset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Copyset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ids", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCopysets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCopysets
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Ids == nil {
				m.Ids = make(map[github_com_cockroachdb_cockroach_pkg_roachpb.StoreID]bool)
			}
			var mapkey int32
			var mapvalue bool
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCopysets
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCopysets
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapvaluetemp int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCopysets
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvaluetemp |= (int(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					mapvalue = bool(mapvaluetemp != 0)
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipCopysets(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthCopysets
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Ids[github_com_cockroachdb_cockroach_pkg_roachpb.StoreID(mapkey)] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCopysets(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCopysets
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
func skipCopysets(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCopysets
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
					return 0, ErrIntOverflowCopysets
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCopysets
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCopysets
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCopysets
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCopysets(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCopysets = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCopysets   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("storage/copysets/copysets.proto", fileDescriptor_copysets_b03c788fb5799fc5)
}

var fileDescriptor_copysets_b03c788fb5799fc5 = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0x33, 0xf9, 0x01, 0xe7, 0x46, 0x82, 0x64, 0x54, 0x50, 0xe4, 0x85, 0x13, 0x05, 0x16,
	0x11, 0x3f, 0x63, 0x11, 0x10, 0x2a, 0x5d, 0xd1, 0x10, 0x83, 0x2c, 0xe1, 0x22, 0x4d, 0x4a, 0x05,
	0x15, 0x28, 0xb2, 0x9d, 0x89, 0x1b, 0x25, 0x74, 0xac, 0xf1, 0x14, 0xe1, 0x6d, 0x79, 0x01, 0xde,
	0x87, 0x17, 0xc8, 0x92, 0x25, 0x2b, 0x0a, 0x61, 0xcf, 0x33, 0x20, 0x3b, 0x13, 0x63, 0x50, 0xa9,
	0xb2, 0xf2, 0xd1, 0xdc, 0x7b, 0xce, 0xbd, 0xdf, 0xc8, 0x03, 0xad, 0x48, 0x72, 0xe1, 0x06, 0xcc,
	0xf4, 0x79, 0x18, 0x47, 0x4c, 0x46, 0x99, 0x20, 0xa1, 0xe0, 0x92, 0x63, 0xdd, 0xe7, 0xfe, 0x4c,
	0x70, 0xd7, 0x3f, 0x22, 0xaa, 0x95, 0xac, 0x3b, 0xf4, 0xad, 0x80, 0x07, 0x3c, 0x6d, 0x33, 0x13,
	0xb5, 0x72, 0x74, 0x3e, 0x16, 0xa1, 0xb6, 0x3b, 0x9f, 0x3f, 0x51, 0x5d, 0xf8, 0x39, 0x54, 0xbc,
	0x78, 0x24, 0x26, 0x4d, 0xd4, 0x2e, 0x75, 0x6b, 0xbd, 0x7b, 0xe4, 0xff, 0x89, 0x24, 0xe7, 0x23,
	0xfd, 0x98, 0x4e, 0xac, 0x63, 0x29, 0xe2, 0x7e, 0x79, 0xf1, 0xad, 0x55, 0xa0, 0x65, 0x2f, 0xa6,
	0x13, 0xfc, 0x0c, 0xb4, 0x48, 0x0a, 0x57, 0xb2, 0x20, 0x6e, 0x16, 0xdb, 0xa8, 0x7b, 0xa5, 0x77,
	0xfb, 0xa2, 0x40, 0x95, 0x36, 0x54, 0x16, 0x9a, 0x99, 0xf5, 0xb7, 0x50, 0xcd, 0x26, 0xe0, 0x3a,
	0x94, 0x66, 0x2c, 0x6e, 0xa2, 0x36, 0xea, 0x56, 0x68, 0x22, 0xf1, 0x0e, 0x54, 0xde, 0xbb, 0xf3,
	0x13, 0x96, 0x0e, 0xa9, 0xf5, 0x6e, 0x6e, 0x30, 0x24, 0xa2, 0x2b, 0xcb, 0x4e, 0x71, 0x1b, 0x75,
	0x7e, 0x21, 0xd0, 0xb2, 0x2b, 0x78, 0x09, 0xe5, 0xe4, 0xab, 0x6e, 0x80, 0x6c, 0x92, 0x45, 0x86,
	0x4c, 0x46, 0x2b, 0xfc, 0x46, 0x82, 0x7f, 0x7a, 0xd6, 0xaa, 0xaa, 0x9a, 0x3d, 0xa0, 0x69, 0x1c,
	0xbe, 0x0b, 0x58, 0xb0, 0x70, 0x3e, 0xf5, 0x5d, 0x39, 0xe5, 0xc7, 0xa3, 0x89, 0xeb, 0x4b, 0x2e,
	0xd2, 0x85, 0x2b, 0xb4, 0x91, 0xab, 0x3c, 0x4d, 0x0b, 0xfa, 0x1b, 0xa8, 0x66, 0xa1, 0xe7, 0x10,
	0x3f, 0xfa, 0x9b, 0xf8, 0xc6, 0x06, 0x5b, 0xe6, 0x81, 0x3f, 0x23, 0xb8, 0xac, 0x8e, 0xf1, 0x07,
	0x28, 0x4d, 0xc7, 0x6b, 0xdc, 0x3b, 0x1b, 0x04, 0x11, 0x7b, 0xac, 0x60, 0xb7, 0x4f, 0xcf, 0x5a,
	0x0f, 0x82, 0xa9, 0x3c, 0x3a, 0xf1, 0x88, 0xcf, 0xdf, 0x99, 0x99, 0x77, 0xec, 0xfd, 0xd1, 0x66,
	0x38, 0x0b, 0xcc, 0x54, 0x85, 0x1e, 0x19, 0x4a, 0x2e, 0x98, 0x3d, 0xa0, 0xc9, 0x48, 0xfd, 0x21,
	0x68, 0xeb, 0xa8, 0x73, 0x10, 0xb7, 0xf2, 0x88, 0x5a, 0x6e, 0xfb, 0x5b, 0x8f, 0xe1, 0xea, 0x3f,
	0xbf, 0x0a, 0xbe, 0x0e, 0xd8, 0xd9, 0x7d, 0x65, 0x3b, 0xf6, 0xa1, 0x35, 0x1a, 0xd8, 0x07, 0x16,
	0x1d, 0xda, 0xfb, 0xaf, 0xeb, 0x05, 0x7c, 0x0d, 0x1a, 0x8e, 0xbd, 0xb7, 0x3a, 0x77, 0x5e, 0x1c,
	0x58, 0x8e, 0xb5, 0xb7, 0x5f, 0x47, 0xfd, 0xce, 0xe2, 0x87, 0x51, 0x58, 0x2c, 0x0d, 0xf4, 0x65,
	0x69, 0xa0, 0xaf, 0x4b, 0x03, 0x7d, 0x5f, 0x1a, 0xe8, 0xd3, 0x4f, 0xa3, 0x70, 0xa8, 0xad, 0x69,
	0xbd, 0x4b, 0xe9, 0x0b, 0xb9, 0xff, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xc8, 0x73, 0x00, 0x76,
	0x03, 0x00, 0x00,
}
