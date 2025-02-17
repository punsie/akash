// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/deployment/v1beta2/deployment.proto

package v1beta2

import (
	fmt "fmt"
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

// State is an enum which refers to state of deployment
type Deployment_State int32

const (
	// Prefix should start with 0 in enum. So declaring dummy state
	DeploymentStateInvalid Deployment_State = 0
	// DeploymentActive denotes state for deployment active
	DeploymentActive Deployment_State = 1
	// DeploymentClosed denotes state for deployment closed
	DeploymentClosed Deployment_State = 2
)

var Deployment_State_name = map[int32]string{
	0: "invalid",
	1: "active",
	2: "closed",
}

var Deployment_State_value = map[string]int32{
	"invalid": 0,
	"active":  1,
	"closed":  2,
}

func (x Deployment_State) String() string {
	return proto.EnumName(Deployment_State_name, int32(x))
}

func (Deployment_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_897ec42830b2cbac, []int{1, 0}
}

// DeploymentID stores owner and sequence number
type DeploymentID struct {
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner" yaml:"owner"`
	DSeq  uint64 `protobuf:"varint,2,opt,name=dseq,proto3" json:"dseq" yaml:"dseq"`
}

func (m *DeploymentID) Reset()      { *m = DeploymentID{} }
func (*DeploymentID) ProtoMessage() {}
func (*DeploymentID) Descriptor() ([]byte, []int) {
	return fileDescriptor_897ec42830b2cbac, []int{0}
}
func (m *DeploymentID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeploymentID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeploymentID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeploymentID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeploymentID.Merge(m, src)
}
func (m *DeploymentID) XXX_Size() int {
	return m.Size()
}
func (m *DeploymentID) XXX_DiscardUnknown() {
	xxx_messageInfo_DeploymentID.DiscardUnknown(m)
}

var xxx_messageInfo_DeploymentID proto.InternalMessageInfo

func (m *DeploymentID) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *DeploymentID) GetDSeq() uint64 {
	if m != nil {
		return m.DSeq
	}
	return 0
}

// Deployment stores deploymentID, state and version details
type Deployment struct {
	DeploymentID DeploymentID     `protobuf:"bytes,1,opt,name=deployment_id,json=deploymentId,proto3" json:"id" yaml:"id"`
	State        Deployment_State `protobuf:"varint,2,opt,name=state,proto3,enum=akash.deployment.v1beta2.Deployment_State" json:"state" yaml:"state"`
	Version      []byte           `protobuf:"bytes,3,opt,name=version,proto3" json:"version" yaml:"version"`
	CreatedAt    int64            `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (m *Deployment) Reset()         { *m = Deployment{} }
func (m *Deployment) String() string { return proto.CompactTextString(m) }
func (*Deployment) ProtoMessage()    {}
func (*Deployment) Descriptor() ([]byte, []int) {
	return fileDescriptor_897ec42830b2cbac, []int{1}
}
func (m *Deployment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Deployment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Deployment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Deployment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deployment.Merge(m, src)
}
func (m *Deployment) XXX_Size() int {
	return m.Size()
}
func (m *Deployment) XXX_DiscardUnknown() {
	xxx_messageInfo_Deployment.DiscardUnknown(m)
}

var xxx_messageInfo_Deployment proto.InternalMessageInfo

func (m *Deployment) GetDeploymentID() DeploymentID {
	if m != nil {
		return m.DeploymentID
	}
	return DeploymentID{}
}

func (m *Deployment) GetState() Deployment_State {
	if m != nil {
		return m.State
	}
	return DeploymentStateInvalid
}

func (m *Deployment) GetVersion() []byte {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *Deployment) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

// DeploymentFilters defines filters used to filter deployments
type DeploymentFilters struct {
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner" yaml:"owner"`
	DSeq  uint64 `protobuf:"varint,2,opt,name=dseq,proto3" json:"dseq" yaml:"dseq"`
	State string `protobuf:"bytes,3,opt,name=state,proto3" json:"state" yaml:"state"`
}

func (m *DeploymentFilters) Reset()         { *m = DeploymentFilters{} }
func (m *DeploymentFilters) String() string { return proto.CompactTextString(m) }
func (*DeploymentFilters) ProtoMessage()    {}
func (*DeploymentFilters) Descriptor() ([]byte, []int) {
	return fileDescriptor_897ec42830b2cbac, []int{2}
}
func (m *DeploymentFilters) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeploymentFilters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeploymentFilters.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeploymentFilters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeploymentFilters.Merge(m, src)
}
func (m *DeploymentFilters) XXX_Size() int {
	return m.Size()
}
func (m *DeploymentFilters) XXX_DiscardUnknown() {
	xxx_messageInfo_DeploymentFilters.DiscardUnknown(m)
}

var xxx_messageInfo_DeploymentFilters proto.InternalMessageInfo

func (m *DeploymentFilters) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *DeploymentFilters) GetDSeq() uint64 {
	if m != nil {
		return m.DSeq
	}
	return 0
}

func (m *DeploymentFilters) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func init() {
	proto.RegisterEnum("akash.deployment.v1beta2.Deployment_State", Deployment_State_name, Deployment_State_value)
	proto.RegisterType((*DeploymentID)(nil), "akash.deployment.v1beta2.DeploymentID")
	proto.RegisterType((*Deployment)(nil), "akash.deployment.v1beta2.Deployment")
	proto.RegisterType((*DeploymentFilters)(nil), "akash.deployment.v1beta2.DeploymentFilters")
}

func init() {
	proto.RegisterFile("akash/deployment/v1beta2/deployment.proto", fileDescriptor_897ec42830b2cbac)
}

var fileDescriptor_897ec42830b2cbac = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x7d, 0x89, 0xd3, 0x92, 0x6b, 0xa8, 0x82, 0x55, 0x21, 0x63, 0xa9, 0x3e, 0xcb, 0x03,
	0x0d, 0x0c, 0xb6, 0x48, 0x07, 0xa4, 0x6c, 0x35, 0x11, 0x52, 0x24, 0x26, 0x77, 0x83, 0xa1, 0x72,
	0x7c, 0xa7, 0xf4, 0x54, 0x27, 0x97, 0xda, 0x87, 0x21, 0x0c, 0xcc, 0xa8, 0x13, 0x23, 0x4b, 0xa5,
	0x4a, 0x7c, 0x01, 0x56, 0xbe, 0x41, 0xc7, 0x8e, 0x4c, 0x27, 0xe4, 0x2c, 0x28, 0x63, 0x3e, 0x01,
	0xf2, 0x5d, 0x82, 0x03, 0x12, 0x88, 0xa9, 0x9b, 0xdf, 0xef, 0xfe, 0xcf, 0xef, 0xff, 0xee, 0xdd,
	0x83, 0x8f, 0xa2, 0xb3, 0x28, 0x3b, 0xf5, 0x31, 0x99, 0x26, 0x6c, 0x36, 0x26, 0x13, 0xee, 0xe7,
	0x4f, 0x86, 0x84, 0x47, 0xdd, 0x0d, 0xe4, 0x4d, 0x53, 0xc6, 0x99, 0x61, 0x4a, 0xa9, 0xb7, 0xc1,
	0x57, 0x52, 0x6b, 0x6f, 0xc4, 0x46, 0x4c, 0x8a, 0xfc, 0xf2, 0x4b, 0xe9, 0xdd, 0xf7, 0xb0, 0xd5,
	0xff, 0xa5, 0x1d, 0xf4, 0x0d, 0x1f, 0x36, 0xd8, 0x9b, 0x09, 0x49, 0x4d, 0xe0, 0x80, 0x4e, 0x33,
	0x78, 0xb0, 0x10, 0x48, 0x81, 0xa5, 0x40, 0xad, 0x59, 0x34, 0x4e, 0x7a, 0xae, 0x0c, 0xdd, 0x50,
	0x61, 0xe3, 0x10, 0xea, 0x38, 0x23, 0xe7, 0x66, 0xcd, 0x01, 0x1d, 0x3d, 0x40, 0x85, 0x40, 0x7a,
	0xff, 0x98, 0x9c, 0x2f, 0x04, 0x92, 0x7c, 0x29, 0xd0, 0x8e, 0x4a, 0x2b, 0x23, 0x37, 0x94, 0xb0,
	0x77, 0xe7, 0xd3, 0x15, 0xd2, 0x7e, 0x5c, 0x21, 0xcd, 0xfd, 0x5a, 0x87, 0xb0, 0x32, 0x60, 0x70,
	0x78, 0xb7, 0xb2, 0x7e, 0x42, 0xb1, 0xb4, 0xb1, 0xd3, 0x7d, 0xe8, 0xfd, 0xad, 0x2d, 0x6f, 0xd3,
	0x7d, 0x70, 0x70, 0x2d, 0x90, 0x56, 0x08, 0xf4, 0x5b, 0x4f, 0x0b, 0x81, 0x6a, 0x14, 0x2f, 0x05,
	0x6a, 0x2a, 0x23, 0x14, 0xbb, 0x61, 0xab, 0xfa, 0xd3, 0x00, 0x1b, 0xaf, 0x60, 0x23, 0xe3, 0x11,
	0x27, 0xb2, 0x89, 0xdd, 0xee, 0xe3, 0xff, 0xa9, 0xe6, 0x1d, 0x97, 0x19, 0xea, 0x82, 0x64, 0x72,
	0x75, 0x41, 0x32, 0x74, 0x43, 0x85, 0x8d, 0xa7, 0x70, 0x3b, 0x27, 0x69, 0x46, 0xd9, 0xc4, 0xac,
	0x3b, 0xa0, 0xd3, 0x0a, 0xf6, 0x17, 0x02, 0xad, 0xd1, 0x52, 0xa0, 0x5d, 0x95, 0xb4, 0x02, 0x6e,
	0xb8, 0x3e, 0x32, 0xf6, 0x21, 0x8c, 0x53, 0x12, 0x71, 0x82, 0x4f, 0x22, 0x6e, 0xea, 0x0e, 0xe8,
	0xd4, 0xc3, 0xe6, 0x8a, 0x1c, 0x71, 0xf7, 0x1d, 0x6c, 0x48, 0x0b, 0xc6, 0x01, 0xdc, 0xa6, 0x93,
	0x3c, 0x4a, 0x28, 0x6e, 0x6b, 0x96, 0x75, 0x71, 0xe9, 0xdc, 0xaf, 0x5c, 0x4a, 0xc5, 0x40, 0x9d,
	0x1a, 0x0e, 0xdc, 0x8a, 0x62, 0x4e, 0x73, 0xd2, 0x06, 0xd6, 0xde, 0xc5, 0xa5, 0xd3, 0xae, 0x74,
	0x47, 0x92, 0x97, 0x8a, 0x38, 0x61, 0x19, 0xc1, 0xed, 0xda, 0x9f, 0x8a, 0x67, 0x92, 0x5b, 0xfa,
	0x87, 0xcf, 0xb6, 0xd6, 0xd3, 0xe5, 0xec, 0xbe, 0x00, 0x78, 0xaf, 0x12, 0x3c, 0xa7, 0x09, 0x27,
	0x69, 0x76, 0x3b, 0x2f, 0xa8, 0xac, 0xa2, 0x46, 0x56, 0xaf, 0xaa, 0xfc, 0x6b, 0x0c, 0xca, 0x72,
	0xf0, 0xe2, 0xba, 0xb0, 0xc1, 0x4d, 0x61, 0x83, 0xef, 0x85, 0x0d, 0x3e, 0xce, 0x6d, 0xed, 0x66,
	0x6e, 0x6b, 0xdf, 0xe6, 0xb6, 0xf6, 0xb2, 0x3b, 0xa2, 0xfc, 0xf4, 0xf5, 0xd0, 0x8b, 0xd9, 0xd8,
	0x67, 0x79, 0x1a, 0x27, 0x67, 0xbe, 0xda, 0xba, 0xb7, 0x9b, 0x7b, 0xc7, 0x67, 0x53, 0x92, 0xad,
	0xb7, 0x6f, 0xb8, 0x25, 0x77, 0xe8, 0xf0, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x9a, 0x81,
	0x9b, 0xa0, 0x03, 0x00, 0x00,
}

func (m *DeploymentID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeploymentID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeploymentID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DSeq != 0 {
		i = encodeVarintDeployment(dAtA, i, uint64(m.DSeq))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintDeployment(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Deployment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Deployment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Deployment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreatedAt != 0 {
		i = encodeVarintDeployment(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintDeployment(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x1a
	}
	if m.State != 0 {
		i = encodeVarintDeployment(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.DeploymentID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintDeployment(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *DeploymentFilters) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeploymentFilters) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeploymentFilters) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintDeployment(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x1a
	}
	if m.DSeq != 0 {
		i = encodeVarintDeployment(dAtA, i, uint64(m.DSeq))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintDeployment(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDeployment(dAtA []byte, offset int, v uint64) int {
	offset -= sovDeployment(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DeploymentID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovDeployment(uint64(l))
	}
	if m.DSeq != 0 {
		n += 1 + sovDeployment(uint64(m.DSeq))
	}
	return n
}

func (m *Deployment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.DeploymentID.Size()
	n += 1 + l + sovDeployment(uint64(l))
	if m.State != 0 {
		n += 1 + sovDeployment(uint64(m.State))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovDeployment(uint64(l))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovDeployment(uint64(m.CreatedAt))
	}
	return n
}

func (m *DeploymentFilters) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovDeployment(uint64(l))
	}
	if m.DSeq != 0 {
		n += 1 + sovDeployment(uint64(m.DSeq))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovDeployment(uint64(l))
	}
	return n
}

func sovDeployment(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDeployment(x uint64) (n int) {
	return sovDeployment(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DeploymentID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeployment
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
			return fmt.Errorf("proto: DeploymentID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeploymentID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeployment
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
				return ErrInvalidLengthDeployment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeployment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DSeq", wireType)
			}
			m.DSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeployment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DSeq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDeployment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeployment
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
func (m *Deployment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeployment
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
			return fmt.Errorf("proto: Deployment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Deployment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeploymentID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeployment
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
				return ErrInvalidLengthDeployment
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeployment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DeploymentID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeployment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= Deployment_State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeployment
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
				return ErrInvalidLengthDeployment
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDeployment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = append(m.Version[:0], dAtA[iNdEx:postIndex]...)
			if m.Version == nil {
				m.Version = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeployment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDeployment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeployment
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
func (m *DeploymentFilters) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeployment
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
			return fmt.Errorf("proto: DeploymentFilters: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeploymentFilters: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeployment
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
				return ErrInvalidLengthDeployment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeployment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DSeq", wireType)
			}
			m.DSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeployment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DSeq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeployment
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
				return ErrInvalidLengthDeployment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeployment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeployment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeployment
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
func skipDeployment(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeployment
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
					return 0, ErrIntOverflowDeployment
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
					return 0, ErrIntOverflowDeployment
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
				return 0, ErrInvalidLengthDeployment
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDeployment
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDeployment
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDeployment        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeployment          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDeployment = fmt.Errorf("proto: unexpected end of group")
)
