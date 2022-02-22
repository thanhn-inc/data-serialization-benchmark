// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tx_ver2.proto

package proto_test

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PbScalar struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PbScalar) Reset()         { *m = PbScalar{} }
func (m *PbScalar) String() string { return proto.CompactTextString(m) }
func (*PbScalar) ProtoMessage()    {}
func (*PbScalar) Descriptor() ([]byte, []int) {
	return fileDescriptor_753663838b47014e, []int{0}
}

func (m *PbScalar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbScalar.Unmarshal(m, b)
}
func (m *PbScalar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbScalar.Marshal(b, m, deterministic)
}
func (m *PbScalar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbScalar.Merge(m, src)
}
func (m *PbScalar) XXX_Size() int {
	return xxx_messageInfo_PbScalar.Size(m)
}
func (m *PbScalar) XXX_DiscardUnknown() {
	xxx_messageInfo_PbScalar.DiscardUnknown(m)
}

var xxx_messageInfo_PbScalar proto.InternalMessageInfo

func (m *PbScalar) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type PbPoint struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PbPoint) Reset()         { *m = PbPoint{} }
func (m *PbPoint) String() string { return proto.CompactTextString(m) }
func (*PbPoint) ProtoMessage()    {}
func (*PbPoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_753663838b47014e, []int{1}
}

func (m *PbPoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbPoint.Unmarshal(m, b)
}
func (m *PbPoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbPoint.Marshal(b, m, deterministic)
}
func (m *PbPoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbPoint.Merge(m, src)
}
func (m *PbPoint) XXX_Size() int {
	return xxx_messageInfo_PbPoint.Size(m)
}
func (m *PbPoint) XXX_DiscardUnknown() {
	xxx_messageInfo_PbPoint.DiscardUnknown(m)
}

var xxx_messageInfo_PbPoint proto.InternalMessageInfo

func (m *PbPoint) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type PbCoinV2 struct {
	Version              int32     `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	Info                 []byte    `protobuf:"bytes,2,opt,name=Info,proto3" json:"Info,omitempty"`
	PublicKey            *PbPoint  `protobuf:"bytes,3,opt,name=PublicKey,proto3" json:"PublicKey,omitempty"`
	Commitment           *PbPoint  `protobuf:"bytes,4,opt,name=Commitment,proto3" json:"Commitment,omitempty"`
	KeyImage             *PbPoint  `protobuf:"bytes,5,opt,name=KeyImage,proto3" json:"KeyImage,omitempty"`
	SharedConcealRandom  *PbScalar `protobuf:"bytes,6,opt,name=SharedConcealRandom,proto3" json:"SharedConcealRandom,omitempty"`
	SharedRandom         *PbScalar `protobuf:"bytes,7,opt,name=SharedRandom,proto3" json:"SharedRandom,omitempty"`
	TxRandom             []byte    `protobuf:"bytes,8,opt,name=TxRandom,proto3" json:"TxRandom,omitempty"`
	Mask                 *PbScalar `protobuf:"bytes,9,opt,name=Mask,proto3" json:"Mask,omitempty"`
	Amount               *PbScalar `protobuf:"bytes,10,opt,name=Amount,proto3" json:"Amount,omitempty"`
	AssetTag             *PbPoint  `protobuf:"bytes,11,opt,name=AssetTag,proto3" json:"AssetTag,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PbCoinV2) Reset()         { *m = PbCoinV2{} }
func (m *PbCoinV2) String() string { return proto.CompactTextString(m) }
func (*PbCoinV2) ProtoMessage()    {}
func (*PbCoinV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_753663838b47014e, []int{2}
}

func (m *PbCoinV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbCoinV2.Unmarshal(m, b)
}
func (m *PbCoinV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbCoinV2.Marshal(b, m, deterministic)
}
func (m *PbCoinV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbCoinV2.Merge(m, src)
}
func (m *PbCoinV2) XXX_Size() int {
	return xxx_messageInfo_PbCoinV2.Size(m)
}
func (m *PbCoinV2) XXX_DiscardUnknown() {
	xxx_messageInfo_PbCoinV2.DiscardUnknown(m)
}

var xxx_messageInfo_PbCoinV2 proto.InternalMessageInfo

func (m *PbCoinV2) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *PbCoinV2) GetInfo() []byte {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *PbCoinV2) GetPublicKey() *PbPoint {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *PbCoinV2) GetCommitment() *PbPoint {
	if m != nil {
		return m.Commitment
	}
	return nil
}

func (m *PbCoinV2) GetKeyImage() *PbPoint {
	if m != nil {
		return m.KeyImage
	}
	return nil
}

func (m *PbCoinV2) GetSharedConcealRandom() *PbScalar {
	if m != nil {
		return m.SharedConcealRandom
	}
	return nil
}

func (m *PbCoinV2) GetSharedRandom() *PbScalar {
	if m != nil {
		return m.SharedRandom
	}
	return nil
}

func (m *PbCoinV2) GetTxRandom() []byte {
	if m != nil {
		return m.TxRandom
	}
	return nil
}

func (m *PbCoinV2) GetMask() *PbScalar {
	if m != nil {
		return m.Mask
	}
	return nil
}

func (m *PbCoinV2) GetAmount() *PbScalar {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *PbCoinV2) GetAssetTag() *PbPoint {
	if m != nil {
		return m.AssetTag
	}
	return nil
}

type PbInnerProductProof struct {
	L                    []*PbPoint `protobuf:"bytes,1,rep,name=L,proto3" json:"L,omitempty"`
	R                    []*PbPoint `protobuf:"bytes,2,rep,name=R,proto3" json:"R,omitempty"`
	A                    *PbScalar  `protobuf:"bytes,3,opt,name=A,proto3" json:"A,omitempty"`
	B                    *PbScalar  `protobuf:"bytes,4,opt,name=B,proto3" json:"B,omitempty"`
	P                    *PbPoint   `protobuf:"bytes,5,opt,name=P,proto3" json:"P,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PbInnerProductProof) Reset()         { *m = PbInnerProductProof{} }
func (m *PbInnerProductProof) String() string { return proto.CompactTextString(m) }
func (*PbInnerProductProof) ProtoMessage()    {}
func (*PbInnerProductProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_753663838b47014e, []int{3}
}

func (m *PbInnerProductProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbInnerProductProof.Unmarshal(m, b)
}
func (m *PbInnerProductProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbInnerProductProof.Marshal(b, m, deterministic)
}
func (m *PbInnerProductProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbInnerProductProof.Merge(m, src)
}
func (m *PbInnerProductProof) XXX_Size() int {
	return xxx_messageInfo_PbInnerProductProof.Size(m)
}
func (m *PbInnerProductProof) XXX_DiscardUnknown() {
	xxx_messageInfo_PbInnerProductProof.DiscardUnknown(m)
}

var xxx_messageInfo_PbInnerProductProof proto.InternalMessageInfo

func (m *PbInnerProductProof) GetL() []*PbPoint {
	if m != nil {
		return m.L
	}
	return nil
}

func (m *PbInnerProductProof) GetR() []*PbPoint {
	if m != nil {
		return m.R
	}
	return nil
}

func (m *PbInnerProductProof) GetA() *PbScalar {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *PbInnerProductProof) GetB() *PbScalar {
	if m != nil {
		return m.B
	}
	return nil
}

func (m *PbInnerProductProof) GetP() *PbPoint {
	if m != nil {
		return m.P
	}
	return nil
}

type PbRangeProof struct {
	CmsValues            []*PbPoint           `protobuf:"bytes,1,rep,name=CmsValues,proto3" json:"CmsValues,omitempty"`
	A                    *PbPoint             `protobuf:"bytes,2,opt,name=A,proto3" json:"A,omitempty"`
	S                    *PbPoint             `protobuf:"bytes,3,opt,name=S,proto3" json:"S,omitempty"`
	T1                   *PbPoint             `protobuf:"bytes,4,opt,name=T1,proto3" json:"T1,omitempty"`
	T2                   *PbPoint             `protobuf:"bytes,5,opt,name=T2,proto3" json:"T2,omitempty"`
	TauX                 *PbScalar            `protobuf:"bytes,6,opt,name=TauX,proto3" json:"TauX,omitempty"`
	THat                 *PbScalar            `protobuf:"bytes,7,opt,name=THat,proto3" json:"THat,omitempty"`
	Mu                   *PbScalar            `protobuf:"bytes,8,opt,name=Mu,proto3" json:"Mu,omitempty"`
	InnerProof           *PbInnerProductProof `protobuf:"bytes,9,opt,name=InnerProof,proto3" json:"InnerProof,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PbRangeProof) Reset()         { *m = PbRangeProof{} }
func (m *PbRangeProof) String() string { return proto.CompactTextString(m) }
func (*PbRangeProof) ProtoMessage()    {}
func (*PbRangeProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_753663838b47014e, []int{4}
}

func (m *PbRangeProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbRangeProof.Unmarshal(m, b)
}
func (m *PbRangeProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbRangeProof.Marshal(b, m, deterministic)
}
func (m *PbRangeProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbRangeProof.Merge(m, src)
}
func (m *PbRangeProof) XXX_Size() int {
	return xxx_messageInfo_PbRangeProof.Size(m)
}
func (m *PbRangeProof) XXX_DiscardUnknown() {
	xxx_messageInfo_PbRangeProof.DiscardUnknown(m)
}

var xxx_messageInfo_PbRangeProof proto.InternalMessageInfo

func (m *PbRangeProof) GetCmsValues() []*PbPoint {
	if m != nil {
		return m.CmsValues
	}
	return nil
}

func (m *PbRangeProof) GetA() *PbPoint {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *PbRangeProof) GetS() *PbPoint {
	if m != nil {
		return m.S
	}
	return nil
}

func (m *PbRangeProof) GetT1() *PbPoint {
	if m != nil {
		return m.T1
	}
	return nil
}

func (m *PbRangeProof) GetT2() *PbPoint {
	if m != nil {
		return m.T2
	}
	return nil
}

func (m *PbRangeProof) GetTauX() *PbScalar {
	if m != nil {
		return m.TauX
	}
	return nil
}

func (m *PbRangeProof) GetTHat() *PbScalar {
	if m != nil {
		return m.THat
	}
	return nil
}

func (m *PbRangeProof) GetMu() *PbScalar {
	if m != nil {
		return m.Mu
	}
	return nil
}

func (m *PbRangeProof) GetInnerProof() *PbInnerProductProof {
	if m != nil {
		return m.InnerProof
	}
	return nil
}

type PbProofV2 struct {
	Version              int32         `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	RangeProof           *PbRangeProof `protobuf:"bytes,2,opt,name=RangeProof,proto3" json:"RangeProof,omitempty"`
	InputCoins           []*PbCoinV2   `protobuf:"bytes,3,rep,name=InputCoins,proto3" json:"InputCoins,omitempty"`
	OutputCoins          []*PbCoinV2   `protobuf:"bytes,4,rep,name=OutputCoins,proto3" json:"OutputCoins,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PbProofV2) Reset()         { *m = PbProofV2{} }
func (m *PbProofV2) String() string { return proto.CompactTextString(m) }
func (*PbProofV2) ProtoMessage()    {}
func (*PbProofV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_753663838b47014e, []int{5}
}

func (m *PbProofV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbProofV2.Unmarshal(m, b)
}
func (m *PbProofV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbProofV2.Marshal(b, m, deterministic)
}
func (m *PbProofV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbProofV2.Merge(m, src)
}
func (m *PbProofV2) XXX_Size() int {
	return xxx_messageInfo_PbProofV2.Size(m)
}
func (m *PbProofV2) XXX_DiscardUnknown() {
	xxx_messageInfo_PbProofV2.DiscardUnknown(m)
}

var xxx_messageInfo_PbProofV2 proto.InternalMessageInfo

func (m *PbProofV2) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *PbProofV2) GetRangeProof() *PbRangeProof {
	if m != nil {
		return m.RangeProof
	}
	return nil
}

func (m *PbProofV2) GetInputCoins() []*PbCoinV2 {
	if m != nil {
		return m.InputCoins
	}
	return nil
}

func (m *PbProofV2) GetOutputCoins() []*PbCoinV2 {
	if m != nil {
		return m.OutputCoins
	}
	return nil
}

type PbTxVer2 struct {
	Version              int32      `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	Type                 string     `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	LockTime             int64      `protobuf:"varint,3,opt,name=LockTime,proto3" json:"LockTime,omitempty"`
	Fee                  uint64     `protobuf:"varint,4,opt,name=Fee,proto3" json:"Fee,omitempty"`
	Info                 []byte     `protobuf:"bytes,5,opt,name=Info,proto3" json:"Info,omitempty"`
	SigPubKey            []byte     `protobuf:"bytes,6,opt,name=SigPubKey,proto3" json:"SigPubKey,omitempty"`
	Sig                  []byte     `protobuf:"bytes,7,opt,name=Sig,proto3" json:"Sig,omitempty"`
	Proof                *PbProofV2 `protobuf:"bytes,8,opt,name=Proof,proto3" json:"Proof,omitempty"`
	LastByte             int32      `protobuf:"varint,9,opt,name=LastByte,proto3" json:"LastByte,omitempty"`
	Metadata             []byte     `protobuf:"bytes,10,opt,name=Metadata,proto3" json:"Metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PbTxVer2) Reset()         { *m = PbTxVer2{} }
func (m *PbTxVer2) String() string { return proto.CompactTextString(m) }
func (*PbTxVer2) ProtoMessage()    {}
func (*PbTxVer2) Descriptor() ([]byte, []int) {
	return fileDescriptor_753663838b47014e, []int{6}
}

func (m *PbTxVer2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbTxVer2.Unmarshal(m, b)
}
func (m *PbTxVer2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbTxVer2.Marshal(b, m, deterministic)
}
func (m *PbTxVer2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbTxVer2.Merge(m, src)
}
func (m *PbTxVer2) XXX_Size() int {
	return xxx_messageInfo_PbTxVer2.Size(m)
}
func (m *PbTxVer2) XXX_DiscardUnknown() {
	xxx_messageInfo_PbTxVer2.DiscardUnknown(m)
}

var xxx_messageInfo_PbTxVer2 proto.InternalMessageInfo

func (m *PbTxVer2) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *PbTxVer2) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PbTxVer2) GetLockTime() int64 {
	if m != nil {
		return m.LockTime
	}
	return 0
}

func (m *PbTxVer2) GetFee() uint64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *PbTxVer2) GetInfo() []byte {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *PbTxVer2) GetSigPubKey() []byte {
	if m != nil {
		return m.SigPubKey
	}
	return nil
}

func (m *PbTxVer2) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

func (m *PbTxVer2) GetProof() *PbProofV2 {
	if m != nil {
		return m.Proof
	}
	return nil
}

func (m *PbTxVer2) GetLastByte() int32 {
	if m != nil {
		return m.LastByte
	}
	return 0
}

func (m *PbTxVer2) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*PbScalar)(nil), "PbScalar")
	proto.RegisterType((*PbPoint)(nil), "PbPoint")
	proto.RegisterType((*PbCoinV2)(nil), "PbCoinV2")
	proto.RegisterType((*PbInnerProductProof)(nil), "PbInnerProductProof")
	proto.RegisterType((*PbRangeProof)(nil), "PbRangeProof")
	proto.RegisterType((*PbProofV2)(nil), "PbProofV2")
	proto.RegisterType((*PbTxVer2)(nil), "PbTxVer2")
}

func init() { proto.RegisterFile("tx_ver2.proto", fileDescriptor_753663838b47014e) }

var fileDescriptor_753663838b47014e = []byte{
	// 630 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xdf, 0x6b, 0x13, 0x41,
	0x10, 0x66, 0x2f, 0x3f, 0x9a, 0x9b, 0xa6, 0x20, 0x5b, 0xd1, 0x55, 0x2b, 0xc4, 0x20, 0xa5, 0x22,
	0x0d, 0x34, 0xfa, 0xe6, 0x53, 0x12, 0x10, 0x4b, 0x1b, 0x5c, 0x36, 0x47, 0x10, 0x5f, 0xca, 0x5e,
	0xb2, 0x8d, 0x47, 0x73, 0xb7, 0xe5, 0x6e, 0x4f, 0x9a, 0x77, 0xdf, 0xfd, 0x17, 0x7c, 0x15, 0xfc,
	0x23, 0x65, 0x27, 0x97, 0xcb, 0xe5, 0x2e, 0xe8, 0xdb, 0xcc, 0x7c, 0xdf, 0xcc, 0xcd, 0xcd, 0xb7,
	0x33, 0x70, 0x64, 0x1e, 0x6e, 0xbe, 0xab, 0xb8, 0xdf, 0xbb, 0x8f, 0xb5, 0xd1, 0xdd, 0x13, 0x68,
	0x71, 0x7f, 0x32, 0x93, 0x4b, 0x19, 0xd3, 0x47, 0x50, 0xbb, 0x53, 0x2b, 0x46, 0x3a, 0xe4, 0xac,
	0x2d, 0xac, 0xd9, 0x7d, 0x01, 0x07, 0xdc, 0xe7, 0x3a, 0x88, 0xcc, 0x1e, 0xf0, 0x57, 0xcd, 0xe6,
	0x8e, 0x74, 0x10, 0x4d, 0xfb, 0x94, 0xc1, 0xc1, 0x54, 0xc5, 0x49, 0xa0, 0x23, 0xa4, 0x34, 0xc4,
	0xc6, 0xa5, 0x14, 0xea, 0x97, 0xd1, 0xad, 0x66, 0x0e, 0x66, 0xa2, 0x4d, 0x4f, 0xc1, 0xe5, 0xa9,
	0xbf, 0x0c, 0x66, 0x57, 0x6a, 0xc5, 0x6a, 0x1d, 0x72, 0x76, 0xd8, 0x6f, 0xf5, 0xb2, 0x2f, 0x89,
	0x2d, 0x44, 0xcf, 0x00, 0x46, 0x3a, 0x0c, 0x03, 0x13, 0xaa, 0xc8, 0xb0, 0x7a, 0x89, 0x58, 0xc0,
	0xe8, 0x6b, 0x68, 0x5d, 0xa9, 0xd5, 0x65, 0x28, 0x17, 0x8a, 0x35, 0x4a, 0xbc, 0x1c, 0xa1, 0x1f,
	0xe0, 0x78, 0xf2, 0x4d, 0xc6, 0x6a, 0x3e, 0xd2, 0xd1, 0x4c, 0xc9, 0xa5, 0x90, 0xd1, 0x5c, 0x87,
	0xac, 0x89, 0x09, 0x6e, 0x6f, 0x33, 0x09, 0xb1, 0x8f, 0x45, 0xcf, 0xa1, 0xbd, 0x0e, 0x67, 0x59,
	0x07, 0xe5, 0xac, 0x1d, 0x98, 0x3e, 0x87, 0x96, 0xf7, 0x90, 0x51, 0x5b, 0xf8, 0xef, 0xb9, 0x4f,
	0x5f, 0x42, 0x7d, 0x2c, 0x93, 0x3b, 0xe6, 0x96, 0x4b, 0x60, 0x98, 0xbe, 0x82, 0xe6, 0x20, 0xd4,
	0x69, 0x64, 0x18, 0x94, 0x09, 0x19, 0x60, 0xff, 0x77, 0x90, 0x24, 0xca, 0x78, 0x72, 0xc1, 0x0e,
	0xcb, 0xff, 0xbb, 0x41, 0xba, 0x3f, 0x09, 0x1c, 0x73, 0xff, 0x32, 0x8a, 0x54, 0xcc, 0x63, 0x3d,
	0x4f, 0x67, 0x86, 0xc7, 0x5a, 0xdf, 0xd2, 0x27, 0x40, 0xae, 0x19, 0xe9, 0xd4, 0x76, 0xd2, 0xc8,
	0xb5, 0x8d, 0x0b, 0xe6, 0x94, 0xe3, 0x82, 0x3e, 0x05, 0x32, 0xc8, 0x74, 0x2a, 0xf4, 0x42, 0x06,
	0x16, 0x18, 0x66, 0xba, 0x14, 0x81, 0xa1, 0xad, 0xc4, 0x2b, 0x42, 0x10, 0xde, 0xfd, 0xed, 0x40,
	0x9b, 0xfb, 0x42, 0x46, 0x0b, 0xb5, 0x6e, 0xe5, 0x14, 0xdc, 0x51, 0x98, 0x4c, 0xe5, 0x32, 0x55,
	0x49, 0xa5, 0xa5, 0x2d, 0x64, 0x0b, 0x0e, 0xf0, 0x0d, 0xed, 0x14, 0x1c, 0xd8, 0xf8, 0xa4, 0xf2,
	0x84, 0xc8, 0x84, 0x32, 0x70, 0xbc, 0x8b, 0xca, 0x93, 0x71, 0xbc, 0x0b, 0x44, 0xfa, 0x95, 0xde,
	0x1c, 0xaf, 0x6f, 0x65, 0xf1, 0x64, 0xfa, 0xa5, 0xfa, 0x1e, 0x30, 0x8c, 0xf0, 0x27, 0x69, 0xaa,
	0xc2, 0x63, 0x98, 0x3e, 0x03, 0x67, 0x9c, 0xa2, 0xd4, 0x3b, 0xa0, 0x33, 0x4e, 0xe9, 0x7b, 0x80,
	0x8d, 0x08, 0xfa, 0x36, 0x53, 0xfd, 0x71, 0x6f, 0x8f, 0x32, 0xa2, 0xc0, 0xeb, 0xfe, 0x21, 0xe0,
	0x72, 0x1f, 0xed, 0x7f, 0x6e, 0xd8, 0x39, 0xc0, 0x76, 0xa0, 0xd9, 0x8c, 0x8e, 0x7a, 0xc5, 0x29,
	0x8b, 0x02, 0x81, 0xbe, 0xb1, 0xcd, 0xdc, 0xa7, 0xc6, 0x6e, 0x6e, 0xc2, 0x6a, 0x38, 0x72, 0xdb,
	0xef, 0x7a, 0x93, 0x45, 0x01, 0xa4, 0x6f, 0xe1, 0xf0, 0x73, 0x6a, 0x72, 0x6e, 0xbd, 0xcc, 0x2d,
	0xa2, 0xdd, 0x1f, 0x8e, 0xbd, 0x07, 0xde, 0xc3, 0x54, 0xc5, 0xff, 0xb9, 0x07, 0xde, 0xea, 0x5e,
	0x61, 0x9f, 0xae, 0x40, 0xdb, 0xee, 0xca, 0xb5, 0x9e, 0xdd, 0x79, 0x41, 0xa8, 0x50, 0xcb, 0x9a,
	0xc8, 0x7d, 0x7b, 0x78, 0x3e, 0x2a, 0x85, 0x4a, 0xd6, 0x85, 0x35, 0xf3, 0x8b, 0xd2, 0x28, 0x5c,
	0x94, 0x13, 0x70, 0x27, 0xc1, 0x82, 0xa7, 0xbe, 0xbd, 0x28, 0x4d, 0x04, 0xb6, 0x01, 0x5b, 0x63,
	0x12, 0x2c, 0x50, 0xb8, 0xb6, 0xb0, 0x26, 0xed, 0x40, 0x63, 0x3d, 0xae, 0xb5, 0x5e, 0xd0, 0xcb,
	0x07, 0x2d, 0xd6, 0x00, 0xf6, 0x24, 0x13, 0x33, 0x5c, 0x19, 0x85, 0x8a, 0x35, 0x44, 0xee, 0x5b,
	0x6c, 0xac, 0x8c, 0x9c, 0x4b, 0x23, 0x71, 0x45, 0xdb, 0x22, 0xf7, 0x87, 0xed, 0xaf, 0x80, 0xa7,
	0xf5, 0xc6, 0xa8, 0xc4, 0xf8, 0x4d, 0xb4, 0xdf, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x10, 0x2e,
	0x35, 0xdb, 0x77, 0x05, 0x00, 0x00,
}
