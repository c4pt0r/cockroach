// Code generated by protoc-gen-gogo.
// source: cockroach/proto/status.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "gogoproto/gogo.pb"

import io "io"
import fmt "fmt"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = math.Inf

// StoreStatus contains the stats needed to calculate the current status of a
// store.
// TODO(bram): add significantly more statistics.
type StoreStatus struct {
	// The storeID for this store.
	StoreID StoreID `protobuf:"varint,1,opt,name=store_id,customtype=StoreID" json:"store_id"`
	// The node id associated with this store.
	NodeID NodeID `protobuf:"varint,2,opt,name=node_id,customtype=NodeID" json:"node_id"`
	// A list of the raft_ids (ranges) available through this store.
	RaftIDs []int64 `protobuf:"varint,3,rep,name=raft_ids" json:"raft_ids"`
	// Time based statistics.
	// The last time this was updated, note that this is not the time it was
	// written to the DB, but when this measurement was taken.
	UpdatedAt int64 `protobuf:"varint,4,opt,name=updated_at" json:"updated_at"`
	// The last time this node was started.
	StartedAt int64 `protobuf:"varint,5,opt,name=started_at" json:"started_at"`
	// Size based statistics.
	// The number of bytes currently being used by this store.
	UsedBytes int64 `protobuf:"varint,6,opt,name=used_bytes" json:"used_bytes"`
	// The maximum number of bytes available.
	MaxBytes         int64  `protobuf:"varint,7,opt,name=max_bytes" json:"max_bytes"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *StoreStatus) Reset()         { *m = StoreStatus{} }
func (m *StoreStatus) String() string { return proto1.CompactTextString(m) }
func (*StoreStatus) ProtoMessage()    {}

func (m *StoreStatus) GetRaftIDs() []int64 {
	if m != nil {
		return m.RaftIDs
	}
	return nil
}

func (m *StoreStatus) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *StoreStatus) GetStartedAt() int64 {
	if m != nil {
		return m.StartedAt
	}
	return 0
}

func (m *StoreStatus) GetUsedBytes() int64 {
	if m != nil {
		return m.UsedBytes
	}
	return 0
}

func (m *StoreStatus) GetMaxBytes() int64 {
	if m != nil {
		return m.MaxBytes
	}
	return 0
}

func init() {
}
func (m *StoreStatus) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreID", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				m.StoreID |= (StoreID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				m.NodeID |= (NodeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaftIDs", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.RaftIDs = append(m.RaftIDs, v)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				m.UpdatedAt |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartedAt", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				m.StartedAt |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UsedBytes", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				m.UsedBytes |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxBytes", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				m.MaxBytes |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *StoreStatus) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovStatus(uint64(m.StoreID))
	n += 1 + sovStatus(uint64(m.NodeID))
	if len(m.RaftIDs) > 0 {
		for _, e := range m.RaftIDs {
			n += 1 + sovStatus(uint64(e))
		}
	}
	n += 1 + sovStatus(uint64(m.UpdatedAt))
	n += 1 + sovStatus(uint64(m.StartedAt))
	n += 1 + sovStatus(uint64(m.UsedBytes))
	n += 1 + sovStatus(uint64(m.MaxBytes))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovStatus(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStatus(x uint64) (n int) {
	return sovStatus(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoreStatus) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *StoreStatus) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintStatus(data, i, uint64(m.StoreID))
	data[i] = 0x10
	i++
	i = encodeVarintStatus(data, i, uint64(m.NodeID))
	if len(m.RaftIDs) > 0 {
		for _, num := range m.RaftIDs {
			data[i] = 0x18
			i++
			i = encodeVarintStatus(data, i, uint64(num))
		}
	}
	data[i] = 0x20
	i++
	i = encodeVarintStatus(data, i, uint64(m.UpdatedAt))
	data[i] = 0x28
	i++
	i = encodeVarintStatus(data, i, uint64(m.StartedAt))
	data[i] = 0x30
	i++
	i = encodeVarintStatus(data, i, uint64(m.UsedBytes))
	data[i] = 0x38
	i++
	i = encodeVarintStatus(data, i, uint64(m.MaxBytes))
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Status(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Status(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintStatus(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}