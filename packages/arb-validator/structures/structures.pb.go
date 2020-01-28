// Code generated by protoc-gen-go. DO NOT EDIT.
// source: structures.proto

package structures

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/offchainlabs/arbitrum/packages/arb-util/common"
	protocol "github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	valprotocol "github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"
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

type ChainParamsBuf struct {
	StakeRequirement        *common.BigIntegerBuf `protobuf:"bytes,1,opt,name=stakeRequirement,proto3" json:"stakeRequirement,omitempty"`
	GracePeriod             *common.TimeTicksBuf  `protobuf:"bytes,2,opt,name=gracePeriod,proto3" json:"gracePeriod,omitempty"`
	MaxExecutionSteps       uint64                `protobuf:"varint,3,opt,name=maxExecutionSteps,proto3" json:"maxExecutionSteps,omitempty"`
	ArbGasSpeedLimitPerTick uint64                `protobuf:"varint,4,opt,name=ArbGasSpeedLimitPerTick,proto3" json:"ArbGasSpeedLimitPerTick,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}              `json:"-"`
	XXX_unrecognized        []byte                `json:"-"`
	XXX_sizecache           int32                 `json:"-"`
}

func (m *ChainParamsBuf) Reset()         { *m = ChainParamsBuf{} }
func (m *ChainParamsBuf) String() string { return proto.CompactTextString(m) }
func (*ChainParamsBuf) ProtoMessage()    {}
func (*ChainParamsBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_66ea84bc81126bff, []int{0}
}

func (m *ChainParamsBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainParamsBuf.Unmarshal(m, b)
}
func (m *ChainParamsBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainParamsBuf.Marshal(b, m, deterministic)
}
func (m *ChainParamsBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainParamsBuf.Merge(m, src)
}
func (m *ChainParamsBuf) XXX_Size() int {
	return xxx_messageInfo_ChainParamsBuf.Size(m)
}
func (m *ChainParamsBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainParamsBuf.DiscardUnknown(m)
}

var xxx_messageInfo_ChainParamsBuf proto.InternalMessageInfo

func (m *ChainParamsBuf) GetStakeRequirement() *common.BigIntegerBuf {
	if m != nil {
		return m.StakeRequirement
	}
	return nil
}

func (m *ChainParamsBuf) GetGracePeriod() *common.TimeTicksBuf {
	if m != nil {
		return m.GracePeriod
	}
	return nil
}

func (m *ChainParamsBuf) GetMaxExecutionSteps() uint64 {
	if m != nil {
		return m.MaxExecutionSteps
	}
	return 0
}

func (m *ChainParamsBuf) GetArbGasSpeedLimitPerTick() uint64 {
	if m != nil {
		return m.ArbGasSpeedLimitPerTick
	}
	return 0
}

type VMProtoDataBuf struct {
	MachineHash          *common.HashBuf       `protobuf:"bytes,1,opt,name=machineHash,proto3" json:"machineHash,omitempty"`
	PendingTop           *common.HashBuf       `protobuf:"bytes,2,opt,name=pendingTop,proto3" json:"pendingTop,omitempty"`
	PendingCount         *common.BigIntegerBuf `protobuf:"bytes,3,opt,name=pendingCount,proto3" json:"pendingCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *VMProtoDataBuf) Reset()         { *m = VMProtoDataBuf{} }
func (m *VMProtoDataBuf) String() string { return proto.CompactTextString(m) }
func (*VMProtoDataBuf) ProtoMessage()    {}
func (*VMProtoDataBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_66ea84bc81126bff, []int{1}
}

func (m *VMProtoDataBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VMProtoDataBuf.Unmarshal(m, b)
}
func (m *VMProtoDataBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VMProtoDataBuf.Marshal(b, m, deterministic)
}
func (m *VMProtoDataBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VMProtoDataBuf.Merge(m, src)
}
func (m *VMProtoDataBuf) XXX_Size() int {
	return xxx_messageInfo_VMProtoDataBuf.Size(m)
}
func (m *VMProtoDataBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_VMProtoDataBuf.DiscardUnknown(m)
}

var xxx_messageInfo_VMProtoDataBuf proto.InternalMessageInfo

func (m *VMProtoDataBuf) GetMachineHash() *common.HashBuf {
	if m != nil {
		return m.MachineHash
	}
	return nil
}

func (m *VMProtoDataBuf) GetPendingTop() *common.HashBuf {
	if m != nil {
		return m.PendingTop
	}
	return nil
}

func (m *VMProtoDataBuf) GetPendingCount() *common.BigIntegerBuf {
	if m != nil {
		return m.PendingCount
	}
	return nil
}

type AssertionParamsBuf struct {
	NumSteps             uint64                        `protobuf:"varint,1,opt,name=numSteps,proto3" json:"numSteps,omitempty"`
	TimeBounds           *protocol.TimeBoundsBlocksBuf `protobuf:"bytes,2,opt,name=timeBounds,proto3" json:"timeBounds,omitempty"`
	ImportedMessageCount *common.BigIntegerBuf         `protobuf:"bytes,3,opt,name=importedMessageCount,proto3" json:"importedMessageCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *AssertionParamsBuf) Reset()         { *m = AssertionParamsBuf{} }
func (m *AssertionParamsBuf) String() string { return proto.CompactTextString(m) }
func (*AssertionParamsBuf) ProtoMessage()    {}
func (*AssertionParamsBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_66ea84bc81126bff, []int{2}
}

func (m *AssertionParamsBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssertionParamsBuf.Unmarshal(m, b)
}
func (m *AssertionParamsBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssertionParamsBuf.Marshal(b, m, deterministic)
}
func (m *AssertionParamsBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssertionParamsBuf.Merge(m, src)
}
func (m *AssertionParamsBuf) XXX_Size() int {
	return xxx_messageInfo_AssertionParamsBuf.Size(m)
}
func (m *AssertionParamsBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_AssertionParamsBuf.DiscardUnknown(m)
}

var xxx_messageInfo_AssertionParamsBuf proto.InternalMessageInfo

func (m *AssertionParamsBuf) GetNumSteps() uint64 {
	if m != nil {
		return m.NumSteps
	}
	return 0
}

func (m *AssertionParamsBuf) GetTimeBounds() *protocol.TimeBoundsBlocksBuf {
	if m != nil {
		return m.TimeBounds
	}
	return nil
}

func (m *AssertionParamsBuf) GetImportedMessageCount() *common.BigIntegerBuf {
	if m != nil {
		return m.ImportedMessageCount
	}
	return nil
}

type AssertionClaimBuf struct {
	AfterPendingTop       *common.HashBuf                        `protobuf:"bytes,1,opt,name=afterPendingTop,proto3" json:"afterPendingTop,omitempty"`
	ImportedMessagesSlice *common.HashBuf                        `protobuf:"bytes,2,opt,name=importedMessagesSlice,proto3" json:"importedMessagesSlice,omitempty"`
	AssertionStub         *valprotocol.ExecutionAssertionStubBuf `protobuf:"bytes,3,opt,name=assertionStub,proto3" json:"assertionStub,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                               `json:"-"`
	XXX_unrecognized      []byte                                 `json:"-"`
	XXX_sizecache         int32                                  `json:"-"`
}

func (m *AssertionClaimBuf) Reset()         { *m = AssertionClaimBuf{} }
func (m *AssertionClaimBuf) String() string { return proto.CompactTextString(m) }
func (*AssertionClaimBuf) ProtoMessage()    {}
func (*AssertionClaimBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_66ea84bc81126bff, []int{3}
}

func (m *AssertionClaimBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssertionClaimBuf.Unmarshal(m, b)
}
func (m *AssertionClaimBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssertionClaimBuf.Marshal(b, m, deterministic)
}
func (m *AssertionClaimBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssertionClaimBuf.Merge(m, src)
}
func (m *AssertionClaimBuf) XXX_Size() int {
	return xxx_messageInfo_AssertionClaimBuf.Size(m)
}
func (m *AssertionClaimBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_AssertionClaimBuf.DiscardUnknown(m)
}

var xxx_messageInfo_AssertionClaimBuf proto.InternalMessageInfo

func (m *AssertionClaimBuf) GetAfterPendingTop() *common.HashBuf {
	if m != nil {
		return m.AfterPendingTop
	}
	return nil
}

func (m *AssertionClaimBuf) GetImportedMessagesSlice() *common.HashBuf {
	if m != nil {
		return m.ImportedMessagesSlice
	}
	return nil
}

func (m *AssertionClaimBuf) GetAssertionStub() *valprotocol.ExecutionAssertionStubBuf {
	if m != nil {
		return m.AssertionStub
	}
	return nil
}

type DisputableNodeBuf struct {
	AssertionParams      *AssertionParamsBuf   `protobuf:"bytes,1,opt,name=assertionParams,proto3" json:"assertionParams,omitempty"`
	AssertionClaim       *AssertionClaimBuf    `protobuf:"bytes,2,opt,name=assertionClaim,proto3" json:"assertionClaim,omitempty"`
	MaxPendingTop        *common.HashBuf       `protobuf:"bytes,3,opt,name=maxPendingTop,proto3" json:"maxPendingTop,omitempty"`
	MaxPendingCount      *common.BigIntegerBuf `protobuf:"bytes,4,opt,name=maxPendingCount,proto3" json:"maxPendingCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DisputableNodeBuf) Reset()         { *m = DisputableNodeBuf{} }
func (m *DisputableNodeBuf) String() string { return proto.CompactTextString(m) }
func (*DisputableNodeBuf) ProtoMessage()    {}
func (*DisputableNodeBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_66ea84bc81126bff, []int{4}
}

func (m *DisputableNodeBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisputableNodeBuf.Unmarshal(m, b)
}
func (m *DisputableNodeBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisputableNodeBuf.Marshal(b, m, deterministic)
}
func (m *DisputableNodeBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisputableNodeBuf.Merge(m, src)
}
func (m *DisputableNodeBuf) XXX_Size() int {
	return xxx_messageInfo_DisputableNodeBuf.Size(m)
}
func (m *DisputableNodeBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_DisputableNodeBuf.DiscardUnknown(m)
}

var xxx_messageInfo_DisputableNodeBuf proto.InternalMessageInfo

func (m *DisputableNodeBuf) GetAssertionParams() *AssertionParamsBuf {
	if m != nil {
		return m.AssertionParams
	}
	return nil
}

func (m *DisputableNodeBuf) GetAssertionClaim() *AssertionClaimBuf {
	if m != nil {
		return m.AssertionClaim
	}
	return nil
}

func (m *DisputableNodeBuf) GetMaxPendingTop() *common.HashBuf {
	if m != nil {
		return m.MaxPendingTop
	}
	return nil
}

func (m *DisputableNodeBuf) GetMaxPendingCount() *common.BigIntegerBuf {
	if m != nil {
		return m.MaxPendingCount
	}
	return nil
}

type PendingInboxBuf struct {
	TopCount             *common.BigIntegerBuf `protobuf:"bytes,1,opt,name=topCount,proto3" json:"topCount,omitempty"`
	ItemHashes           []*common.HashBuf     `protobuf:"bytes,2,rep,name=itemHashes,proto3" json:"itemHashes,omitempty"`
	HashOfRest           *common.HashBuf       `protobuf:"bytes,3,opt,name=hashOfRest,proto3" json:"hashOfRest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PendingInboxBuf) Reset()         { *m = PendingInboxBuf{} }
func (m *PendingInboxBuf) String() string { return proto.CompactTextString(m) }
func (*PendingInboxBuf) ProtoMessage()    {}
func (*PendingInboxBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_66ea84bc81126bff, []int{5}
}

func (m *PendingInboxBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PendingInboxBuf.Unmarshal(m, b)
}
func (m *PendingInboxBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PendingInboxBuf.Marshal(b, m, deterministic)
}
func (m *PendingInboxBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PendingInboxBuf.Merge(m, src)
}
func (m *PendingInboxBuf) XXX_Size() int {
	return xxx_messageInfo_PendingInboxBuf.Size(m)
}
func (m *PendingInboxBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_PendingInboxBuf.DiscardUnknown(m)
}

var xxx_messageInfo_PendingInboxBuf proto.InternalMessageInfo

func (m *PendingInboxBuf) GetTopCount() *common.BigIntegerBuf {
	if m != nil {
		return m.TopCount
	}
	return nil
}

func (m *PendingInboxBuf) GetItemHashes() []*common.HashBuf {
	if m != nil {
		return m.ItemHashes
	}
	return nil
}

func (m *PendingInboxBuf) GetHashOfRest() *common.HashBuf {
	if m != nil {
		return m.HashOfRest
	}
	return nil
}

type CheckpointManifest struct {
	Values               []*common.HashBuf `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	Machines             []*common.HashBuf `protobuf:"bytes,2,rep,name=machines,proto3" json:"machines,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CheckpointManifest) Reset()         { *m = CheckpointManifest{} }
func (m *CheckpointManifest) String() string { return proto.CompactTextString(m) }
func (*CheckpointManifest) ProtoMessage()    {}
func (*CheckpointManifest) Descriptor() ([]byte, []int) {
	return fileDescriptor_66ea84bc81126bff, []int{6}
}

func (m *CheckpointManifest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckpointManifest.Unmarshal(m, b)
}
func (m *CheckpointManifest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckpointManifest.Marshal(b, m, deterministic)
}
func (m *CheckpointManifest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckpointManifest.Merge(m, src)
}
func (m *CheckpointManifest) XXX_Size() int {
	return xxx_messageInfo_CheckpointManifest.Size(m)
}
func (m *CheckpointManifest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckpointManifest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckpointManifest proto.InternalMessageInfo

func (m *CheckpointManifest) GetValues() []*common.HashBuf {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *CheckpointManifest) GetMachines() []*common.HashBuf {
	if m != nil {
		return m.Machines
	}
	return nil
}

type CheckpointMetadata struct {
	FormatVersion        uint64      `protobuf:"varint,1,opt,name=formatVersion,proto3" json:"formatVersion,omitempty"`
	Oldest               *BlockIdBuf `protobuf:"bytes,2,opt,name=oldest,proto3" json:"oldest,omitempty"`
	Newest               *BlockIdBuf `protobuf:"bytes,3,opt,name=newest,proto3" json:"newest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CheckpointMetadata) Reset()         { *m = CheckpointMetadata{} }
func (m *CheckpointMetadata) String() string { return proto.CompactTextString(m) }
func (*CheckpointMetadata) ProtoMessage()    {}
func (*CheckpointMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_66ea84bc81126bff, []int{7}
}

func (m *CheckpointMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckpointMetadata.Unmarshal(m, b)
}
func (m *CheckpointMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckpointMetadata.Marshal(b, m, deterministic)
}
func (m *CheckpointMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckpointMetadata.Merge(m, src)
}
func (m *CheckpointMetadata) XXX_Size() int {
	return xxx_messageInfo_CheckpointMetadata.Size(m)
}
func (m *CheckpointMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckpointMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_CheckpointMetadata proto.InternalMessageInfo

func (m *CheckpointMetadata) GetFormatVersion() uint64 {
	if m != nil {
		return m.FormatVersion
	}
	return 0
}

func (m *CheckpointMetadata) GetOldest() *BlockIdBuf {
	if m != nil {
		return m.Oldest
	}
	return nil
}

func (m *CheckpointMetadata) GetNewest() *BlockIdBuf {
	if m != nil {
		return m.Newest
	}
	return nil
}

type CheckpointLinks struct {
	Next                 *BlockIdBuf `protobuf:"bytes,1,opt,name=next,proto3" json:"next,omitempty"`
	Prev                 *BlockIdBuf `protobuf:"bytes,2,opt,name=prev,proto3" json:"prev,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CheckpointLinks) Reset()         { *m = CheckpointLinks{} }
func (m *CheckpointLinks) String() string { return proto.CompactTextString(m) }
func (*CheckpointLinks) ProtoMessage()    {}
func (*CheckpointLinks) Descriptor() ([]byte, []int) {
	return fileDescriptor_66ea84bc81126bff, []int{8}
}

func (m *CheckpointLinks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckpointLinks.Unmarshal(m, b)
}
func (m *CheckpointLinks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckpointLinks.Marshal(b, m, deterministic)
}
func (m *CheckpointLinks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckpointLinks.Merge(m, src)
}
func (m *CheckpointLinks) XXX_Size() int {
	return xxx_messageInfo_CheckpointLinks.Size(m)
}
func (m *CheckpointLinks) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckpointLinks.DiscardUnknown(m)
}

var xxx_messageInfo_CheckpointLinks proto.InternalMessageInfo

func (m *CheckpointLinks) GetNext() *BlockIdBuf {
	if m != nil {
		return m.Next
	}
	return nil
}

func (m *CheckpointLinks) GetPrev() *BlockIdBuf {
	if m != nil {
		return m.Prev
	}
	return nil
}

type BlockIdBuf struct {
	Height               *common.TimeBlocksBuf `protobuf:"bytes,1,opt,name=height,proto3" json:"height,omitempty"`
	HeaderHash           *common.HashBuf       `protobuf:"bytes,2,opt,name=headerHash,proto3" json:"headerHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *BlockIdBuf) Reset()         { *m = BlockIdBuf{} }
func (m *BlockIdBuf) String() string { return proto.CompactTextString(m) }
func (*BlockIdBuf) ProtoMessage()    {}
func (*BlockIdBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_66ea84bc81126bff, []int{9}
}

func (m *BlockIdBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockIdBuf.Unmarshal(m, b)
}
func (m *BlockIdBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockIdBuf.Marshal(b, m, deterministic)
}
func (m *BlockIdBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockIdBuf.Merge(m, src)
}
func (m *BlockIdBuf) XXX_Size() int {
	return xxx_messageInfo_BlockIdBuf.Size(m)
}
func (m *BlockIdBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockIdBuf.DiscardUnknown(m)
}

var xxx_messageInfo_BlockIdBuf proto.InternalMessageInfo

func (m *BlockIdBuf) GetHeight() *common.TimeBlocksBuf {
	if m != nil {
		return m.Height
	}
	return nil
}

func (m *BlockIdBuf) GetHeaderHash() *common.HashBuf {
	if m != nil {
		return m.HeaderHash
	}
	return nil
}

type BlockIdBufList struct {
	Bufs                 []*BlockIdBuf `protobuf:"bytes,1,rep,name=bufs,proto3" json:"bufs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *BlockIdBufList) Reset()         { *m = BlockIdBufList{} }
func (m *BlockIdBufList) String() string { return proto.CompactTextString(m) }
func (*BlockIdBufList) ProtoMessage()    {}
func (*BlockIdBufList) Descriptor() ([]byte, []int) {
	return fileDescriptor_66ea84bc81126bff, []int{10}
}

func (m *BlockIdBufList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockIdBufList.Unmarshal(m, b)
}
func (m *BlockIdBufList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockIdBufList.Marshal(b, m, deterministic)
}
func (m *BlockIdBufList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockIdBufList.Merge(m, src)
}
func (m *BlockIdBufList) XXX_Size() int {
	return xxx_messageInfo_BlockIdBufList.Size(m)
}
func (m *BlockIdBufList) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockIdBufList.DiscardUnknown(m)
}

var xxx_messageInfo_BlockIdBufList proto.InternalMessageInfo

func (m *BlockIdBufList) GetBufs() []*BlockIdBuf {
	if m != nil {
		return m.Bufs
	}
	return nil
}

func init() {
	proto.RegisterType((*ChainParamsBuf)(nil), "structures.ChainParamsBuf")
	proto.RegisterType((*VMProtoDataBuf)(nil), "structures.VMProtoDataBuf")
	proto.RegisterType((*AssertionParamsBuf)(nil), "structures.AssertionParamsBuf")
	proto.RegisterType((*AssertionClaimBuf)(nil), "structures.AssertionClaimBuf")
	proto.RegisterType((*DisputableNodeBuf)(nil), "structures.DisputableNodeBuf")
	proto.RegisterType((*PendingInboxBuf)(nil), "structures.PendingInboxBuf")
	proto.RegisterType((*CheckpointManifest)(nil), "structures.CheckpointManifest")
	proto.RegisterType((*CheckpointMetadata)(nil), "structures.CheckpointMetadata")
	proto.RegisterType((*CheckpointLinks)(nil), "structures.CheckpointLinks")
	proto.RegisterType((*BlockIdBuf)(nil), "structures.BlockIdBuf")
	proto.RegisterType((*BlockIdBufList)(nil), "structures.BlockIdBufList")
}

func init() { proto.RegisterFile("structures.proto", fileDescriptor_66ea84bc81126bff) }

var fileDescriptor_66ea84bc81126bff = []byte{
	// 802 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x5f, 0x6f, 0xe3, 0x44,
	0x10, 0x97, 0xaf, 0x55, 0x55, 0x4d, 0x69, 0x43, 0x97, 0x3b, 0xae, 0xaa, 0x74, 0xe8, 0x64, 0x21,
	0x38, 0x01, 0x97, 0x08, 0x10, 0x88, 0x93, 0x40, 0xa8, 0x69, 0x2b, 0xae, 0x52, 0x0a, 0x51, 0x12,
	0xdd, 0x03, 0x6f, 0x63, 0x7b, 0x1c, 0x2f, 0xf1, 0xee, 0x9a, 0xdd, 0x75, 0xc8, 0x67, 0x41, 0x3c,
	0xf3, 0xca, 0x3b, 0x5f, 0x87, 0x8f, 0xc0, 0x17, 0x40, 0x6b, 0x3b, 0xb6, 0xf3, 0xc7, 0xe5, 0x9e,
	0x6c, 0xcf, 0xfc, 0x66, 0xe6, 0x37, 0x7f, 0x0d, 0xef, 0x1a, 0xab, 0xf3, 0xd0, 0xe6, 0x9a, 0x4c,
	0x3f, 0xd3, 0xca, 0x2a, 0x06, 0x8d, 0xe4, 0xf2, 0xbd, 0x50, 0x09, 0xa1, 0xe4, 0xa0, 0x7c, 0x94,
	0x80, 0xcb, 0xa7, 0xc5, 0x23, 0x54, 0xe9, 0x60, 0xfd, 0x52, 0x29, 0x9e, 0x2d, 0x31, 0xad, 0x75,
	0xad, 0xf7, 0x52, 0xed, 0xff, 0xeb, 0xc1, 0xd9, 0x75, 0x82, 0x5c, 0x8e, 0x51, 0xa3, 0x30, 0xc3,
	0x3c, 0x66, 0x57, 0x2e, 0x3e, 0x2e, 0x68, 0x42, 0xbf, 0xe6, 0x5c, 0x93, 0x20, 0x69, 0x2f, 0xbc,
	0xe7, 0xde, 0x8b, 0x93, 0x2f, 0x9e, 0xf4, 0xab, 0x98, 0x43, 0x3e, 0xbf, 0x93, 0x96, 0xe6, 0xa4,
	0x87, 0x79, 0x3c, 0xd9, 0x81, 0xb3, 0xaf, 0xe1, 0x64, 0xae, 0x31, 0xa4, 0x31, 0x69, 0xae, 0xa2,
	0x8b, 0x47, 0x85, 0xf5, 0xe3, 0xb5, 0xf5, 0x8c, 0x0b, 0x9a, 0xf1, 0x70, 0xe1, 0xa2, 0x4d, 0xda,
	0x40, 0xf6, 0x19, 0x9c, 0x0b, 0x5c, 0xdd, 0xae, 0x28, 0xcc, 0x2d, 0x57, 0x72, 0x6a, 0x29, 0x33,
	0x17, 0x07, 0xcf, 0xbd, 0x17, 0x87, 0x93, 0x5d, 0x05, 0xfb, 0x06, 0x9e, 0x5e, 0xe9, 0xe0, 0x07,
	0x34, 0xd3, 0x8c, 0x28, 0x1a, 0x71, 0xc1, 0xed, 0x98, 0xb4, 0xf3, 0x7c, 0x71, 0x58, 0xd8, 0x74,
	0xa9, 0xfd, 0xbf, 0x3c, 0x38, 0x7b, 0x73, 0x3f, 0x76, 0x15, 0xb8, 0x41, 0x8b, 0x2e, 0xeb, 0xcf,
	0xe1, 0x44, 0x60, 0x98, 0x70, 0x49, 0xaf, 0xd1, 0x24, 0x55, 0xc2, 0xbd, 0x35, 0x65, 0x27, 0x2b,
	0xd8, 0xb6, 0x30, 0x6c, 0x00, 0x90, 0x91, 0x8c, 0xb8, 0x9c, 0xcf, 0x54, 0x56, 0x25, 0xb9, 0x63,
	0xd1, 0x82, 0xb0, 0x57, 0xf0, 0x4e, 0xf5, 0x75, 0xad, 0x72, 0x69, 0x8b, 0xcc, 0x3a, 0xab, 0xba,
	0x01, 0xf5, 0xff, 0xf6, 0x80, 0x5d, 0x19, 0x43, 0xda, 0xa5, 0xdf, 0xf4, 0xea, 0x12, 0x8e, 0x65,
	0x2e, 0xca, 0x3a, 0x79, 0x45, 0xce, 0xf5, 0x37, 0xfb, 0x0e, 0xc0, 0x72, 0x41, 0x43, 0x95, 0xcb,
	0xc8, 0x54, 0xf4, 0x9e, 0xf5, 0xeb, 0xfe, 0xcf, 0x6a, 0xdd, 0x30, 0x55, 0x55, 0x33, 0x5a, 0x06,
	0xec, 0x0e, 0x1e, 0x73, 0x91, 0x29, 0x6d, 0x29, 0xba, 0x27, 0x63, 0x70, 0x4e, 0x6f, 0x41, 0x7a,
	0xaf, 0x89, 0xff, 0x8f, 0x07, 0xe7, 0x35, 0xf9, 0xeb, 0x14, 0xb9, 0x70, 0xdc, 0x5f, 0x41, 0x0f,
	0x63, 0x4b, 0x7a, 0xdc, 0xd4, 0xb0, 0xa3, 0xea, 0xdb, 0x38, 0x76, 0x0b, 0x4f, 0xb6, 0x02, 0x99,
	0x69, 0xca, 0x43, 0xea, 0x6a, 0xc2, 0x7e, 0x34, 0x1b, 0xc1, 0x29, 0xae, 0x69, 0x4d, 0x6d, 0x1e,
	0x54, 0xb9, 0x7d, 0xd4, 0x6f, 0xef, 0x49, 0x3d, 0x74, 0x57, 0x6d, 0xa8, 0xf3, 0xba, 0x69, 0xec,
	0xff, 0xf1, 0x08, 0xce, 0x6f, 0xb8, 0xc9, 0x72, 0x8b, 0x41, 0x4a, 0x3f, 0xaa, 0x88, 0x5c, 0x96,
	0xaf, 0xa1, 0x87, 0x9b, 0x7d, 0xab, 0xb2, 0xfc, 0xa0, 0xdf, 0xda, 0xf2, 0xdd, 0xd6, 0x4e, 0xb6,
	0xcd, 0xd8, 0x2d, 0x9c, 0xe1, 0x46, 0x11, 0xeb, 0x9e, 0xee, 0x73, 0xb4, 0x2e, 0xf3, 0x64, 0xcb,
	0x88, 0x7d, 0x05, 0xa7, 0x02, 0x57, 0xad, 0xa2, 0x1f, 0xec, 0xaf, 0xd9, 0x26, 0x8a, 0x7d, 0x0f,
	0xbd, 0x46, 0x50, 0x4e, 0xc2, 0xe1, 0x43, 0x93, 0xb0, 0x8d, 0xf6, 0xff, 0xf4, 0xa0, 0x57, 0x09,
	0xee, 0x64, 0xa0, 0x56, 0xe5, 0xd2, 0x1d, 0x5b, 0x95, 0x95, 0xde, 0x1e, 0x3c, 0x31, 0x35, 0xcc,
	0x2d, 0x1d, 0xb7, 0x24, 0x1c, 0x4b, 0x72, 0x53, 0x7d, 0xb0, 0x77, 0xe9, 0x1a, 0x88, 0x33, 0x48,
	0xd0, 0x24, 0x3f, 0xc5, 0x13, 0x32, 0xb6, 0x2b, 0xd9, 0x16, 0xc4, 0xff, 0x05, 0xd8, 0x75, 0x42,
	0xe1, 0x22, 0x53, 0x5c, 0xda, 0x7b, 0x94, 0x3c, 0x26, 0x63, 0xd9, 0xc7, 0x70, 0xb4, 0xc4, 0x34,
	0x27, 0xd7, 0xbe, 0xbd, 0x31, 0x2b, 0x35, 0xfb, 0x14, 0x8e, 0xab, 0x23, 0xd1, 0x49, 0xaf, 0x06,
	0xf8, 0xbf, 0x7b, 0x1b, 0xc1, 0xc8, 0x62, 0x84, 0x16, 0xd9, 0x87, 0x70, 0x1a, 0x2b, 0x2d, 0xd0,
	0xbe, 0x21, 0x6d, 0xb8, 0x92, 0xd5, 0x6e, 0x6f, 0x0a, 0x59, 0x1f, 0x8e, 0x54, 0x1a, 0xb9, 0xac,
	0xca, 0x41, 0x78, 0xbf, 0x3d, 0x08, 0xc5, 0x52, 0xdf, 0x45, 0x05, 0xb3, 0x12, 0xe5, 0xf0, 0x92,
	0x7e, 0x6b, 0xaa, 0xd0, 0x89, 0x2f, 0x51, 0x3e, 0x87, 0x5e, 0xc3, 0x6d, 0xc4, 0xe5, 0xc2, 0xb0,
	0x4f, 0xe0, 0x50, 0xd2, 0x6a, 0xdd, 0xac, 0x2e, 0x07, 0x05, 0xc6, 0x61, 0x33, 0x4d, 0xcb, 0xff,
	0x21, 0x57, 0x60, 0xfc, 0x14, 0xa0, 0x91, 0xb1, 0x97, 0x70, 0x94, 0x10, 0x9f, 0x27, 0x3b, 0x43,
	0x51, 0xdc, 0xac, 0xfa, 0x5a, 0x55, 0xa0, 0xa2, 0xc3, 0x84, 0x11, 0xe9, 0xe2, 0x72, 0x77, 0xdd,
	0xe1, 0x06, 0xe2, 0x7f, 0x0b, 0x67, 0x4d, 0xb4, 0x11, 0x37, 0x05, 0xd7, 0x20, 0x8f, 0xd7, 0xbd,
	0xed, 0xe4, 0xea, 0x30, 0xc3, 0x9b, 0x9f, 0x87, 0x73, 0x6e, 0x93, 0x3c, 0x70, 0x21, 0x06, 0x2a,
	0x8e, 0x43, 0xf7, 0xff, 0x4c, 0x31, 0x30, 0x03, 0xd4, 0x01, 0xb7, 0x3a, 0x17, 0x83, 0x0c, 0xc3,
	0x85, 0xbb, 0x33, 0x4e, 0xf2, 0x72, 0x89, 0x29, 0x8f, 0xd0, 0x2a, 0x3d, 0x68, 0xdc, 0x06, 0x47,
	0xc5, 0x81, 0xf9, 0xf2, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xb1, 0x18, 0xfe, 0xec, 0x07,
	0x00, 0x00,
}
