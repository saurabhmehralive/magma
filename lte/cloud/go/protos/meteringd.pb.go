// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lte/protos/meteringd.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protos "magma/orc8r/lib/go/protos"
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

type FlowRecord struct {
	Id                   *FlowRecord_ID       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Sid                  string               `protobuf:"bytes,2,opt,name=sid,proto3" json:"sid,omitempty"`
	GatewayId            string               `protobuf:"bytes,3,opt,name=gateway_id,json=gatewayId,proto3" json:"gateway_id,omitempty"`
	BytesTx              uint64               `protobuf:"varint,5,opt,name=bytes_tx,json=bytesTx,proto3" json:"bytes_tx,omitempty"`
	BytesRx              uint64               `protobuf:"varint,6,opt,name=bytes_rx,json=bytesRx,proto3" json:"bytes_rx,omitempty"`
	PktsTx               uint64               `protobuf:"varint,7,opt,name=pkts_tx,json=pktsTx,proto3" json:"pkts_tx,omitempty"`
	PktsRx               uint64               `protobuf:"varint,8,opt,name=pkts_rx,json=pktsRx,proto3" json:"pkts_rx,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,9,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FlowRecord) Reset()         { *m = FlowRecord{} }
func (m *FlowRecord) String() string { return proto.CompactTextString(m) }
func (*FlowRecord) ProtoMessage()    {}
func (*FlowRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88a6fd3c74575d2, []int{0}
}

func (m *FlowRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowRecord.Unmarshal(m, b)
}
func (m *FlowRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowRecord.Marshal(b, m, deterministic)
}
func (m *FlowRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowRecord.Merge(m, src)
}
func (m *FlowRecord) XXX_Size() int {
	return xxx_messageInfo_FlowRecord.Size(m)
}
func (m *FlowRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowRecord.DiscardUnknown(m)
}

var xxx_messageInfo_FlowRecord proto.InternalMessageInfo

func (m *FlowRecord) GetId() *FlowRecord_ID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *FlowRecord) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *FlowRecord) GetGatewayId() string {
	if m != nil {
		return m.GatewayId
	}
	return ""
}

func (m *FlowRecord) GetBytesTx() uint64 {
	if m != nil {
		return m.BytesTx
	}
	return 0
}

func (m *FlowRecord) GetBytesRx() uint64 {
	if m != nil {
		return m.BytesRx
	}
	return 0
}

func (m *FlowRecord) GetPktsTx() uint64 {
	if m != nil {
		return m.PktsTx
	}
	return 0
}

func (m *FlowRecord) GetPktsRx() uint64 {
	if m != nil {
		return m.PktsRx
	}
	return 0
}

func (m *FlowRecord) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

type FlowRecord_ID struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowRecord_ID) Reset()         { *m = FlowRecord_ID{} }
func (m *FlowRecord_ID) String() string { return proto.CompactTextString(m) }
func (*FlowRecord_ID) ProtoMessage()    {}
func (*FlowRecord_ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88a6fd3c74575d2, []int{0, 0}
}

func (m *FlowRecord_ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowRecord_ID.Unmarshal(m, b)
}
func (m *FlowRecord_ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowRecord_ID.Marshal(b, m, deterministic)
}
func (m *FlowRecord_ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowRecord_ID.Merge(m, src)
}
func (m *FlowRecord_ID) XXX_Size() int {
	return xxx_messageInfo_FlowRecord_ID.Size(m)
}
func (m *FlowRecord_ID) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowRecord_ID.DiscardUnknown(m)
}

var xxx_messageInfo_FlowRecord_ID proto.InternalMessageInfo

func (m *FlowRecord_ID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type FlowRecordSet struct {
	RecordIds            []string `protobuf:"bytes,1,rep,name=record_ids,json=recordIds,proto3" json:"record_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowRecordSet) Reset()         { *m = FlowRecordSet{} }
func (m *FlowRecordSet) String() string { return proto.CompactTextString(m) }
func (*FlowRecordSet) ProtoMessage()    {}
func (*FlowRecordSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88a6fd3c74575d2, []int{1}
}

func (m *FlowRecordSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowRecordSet.Unmarshal(m, b)
}
func (m *FlowRecordSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowRecordSet.Marshal(b, m, deterministic)
}
func (m *FlowRecordSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowRecordSet.Merge(m, src)
}
func (m *FlowRecordSet) XXX_Size() int {
	return xxx_messageInfo_FlowRecordSet.Size(m)
}
func (m *FlowRecordSet) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowRecordSet.DiscardUnknown(m)
}

var xxx_messageInfo_FlowRecordSet proto.InternalMessageInfo

func (m *FlowRecordSet) GetRecordIds() []string {
	if m != nil {
		return m.RecordIds
	}
	return nil
}

type FlowTable struct {
	Flows                []*FlowRecord `protobuf:"bytes,1,rep,name=flows,proto3" json:"flows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FlowTable) Reset()         { *m = FlowTable{} }
func (m *FlowTable) String() string { return proto.CompactTextString(m) }
func (*FlowTable) ProtoMessage()    {}
func (*FlowTable) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88a6fd3c74575d2, []int{2}
}

func (m *FlowTable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowTable.Unmarshal(m, b)
}
func (m *FlowTable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowTable.Marshal(b, m, deterministic)
}
func (m *FlowTable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowTable.Merge(m, src)
}
func (m *FlowTable) XXX_Size() int {
	return xxx_messageInfo_FlowTable.Size(m)
}
func (m *FlowTable) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowTable.DiscardUnknown(m)
}

var xxx_messageInfo_FlowTable proto.InternalMessageInfo

func (m *FlowTable) GetFlows() []*FlowRecord {
	if m != nil {
		return m.Flows
	}
	return nil
}

// --------------------------------------------------------------------------
// Cloud controller for MeteringD
//
// The MeteringdRecordsController is an RPC service used by the gateway to
// upstream records, and thus requires a Gateway context for that operation.
// It also serves as the read interface for obsidian or other cloud services to
// read records and expects a Network context for such operations.
//
// There is also the record query wrappers that pass the network context
// for cloud orinated requests.
// --------------------------------------------------------------------------
type FlowRecordQuery struct {
	NetworkId string `protobuf:"bytes,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// Types that are valid to be assigned to Query:
	//	*FlowRecordQuery_RecordId
	//	*FlowRecordQuery_GatewayId
	//	*FlowRecordQuery_SubscriberId
	Query                isFlowRecordQuery_Query `protobuf_oneof:"query"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FlowRecordQuery) Reset()         { *m = FlowRecordQuery{} }
func (m *FlowRecordQuery) String() string { return proto.CompactTextString(m) }
func (*FlowRecordQuery) ProtoMessage()    {}
func (*FlowRecordQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88a6fd3c74575d2, []int{3}
}

func (m *FlowRecordQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowRecordQuery.Unmarshal(m, b)
}
func (m *FlowRecordQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowRecordQuery.Marshal(b, m, deterministic)
}
func (m *FlowRecordQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowRecordQuery.Merge(m, src)
}
func (m *FlowRecordQuery) XXX_Size() int {
	return xxx_messageInfo_FlowRecordQuery.Size(m)
}
func (m *FlowRecordQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowRecordQuery.DiscardUnknown(m)
}

var xxx_messageInfo_FlowRecordQuery proto.InternalMessageInfo

func (m *FlowRecordQuery) GetNetworkId() string {
	if m != nil {
		return m.NetworkId
	}
	return ""
}

type isFlowRecordQuery_Query interface {
	isFlowRecordQuery_Query()
}

type FlowRecordQuery_RecordId struct {
	RecordId string `protobuf:"bytes,2,opt,name=record_id,json=recordId,proto3,oneof"`
}

type FlowRecordQuery_GatewayId struct {
	GatewayId string `protobuf:"bytes,3,opt,name=gateway_id,json=gatewayId,proto3,oneof"`
}

type FlowRecordQuery_SubscriberId struct {
	SubscriberId string `protobuf:"bytes,4,opt,name=subscriber_id,json=subscriberId,proto3,oneof"`
}

func (*FlowRecordQuery_RecordId) isFlowRecordQuery_Query() {}

func (*FlowRecordQuery_GatewayId) isFlowRecordQuery_Query() {}

func (*FlowRecordQuery_SubscriberId) isFlowRecordQuery_Query() {}

func (m *FlowRecordQuery) GetQuery() isFlowRecordQuery_Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *FlowRecordQuery) GetRecordId() string {
	if x, ok := m.GetQuery().(*FlowRecordQuery_RecordId); ok {
		return x.RecordId
	}
	return ""
}

func (m *FlowRecordQuery) GetGatewayId() string {
	if x, ok := m.GetQuery().(*FlowRecordQuery_GatewayId); ok {
		return x.GatewayId
	}
	return ""
}

func (m *FlowRecordQuery) GetSubscriberId() string {
	if x, ok := m.GetQuery().(*FlowRecordQuery_SubscriberId); ok {
		return x.SubscriberId
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FlowRecordQuery) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FlowRecordQuery_RecordId)(nil),
		(*FlowRecordQuery_GatewayId)(nil),
		(*FlowRecordQuery_SubscriberId)(nil),
	}
}

func init() {
	proto.RegisterType((*FlowRecord)(nil), "magma.lte.FlowRecord")
	proto.RegisterType((*FlowRecord_ID)(nil), "magma.lte.FlowRecord.ID")
	proto.RegisterType((*FlowRecordSet)(nil), "magma.lte.FlowRecordSet")
	proto.RegisterType((*FlowTable)(nil), "magma.lte.FlowTable")
	proto.RegisterType((*FlowRecordQuery)(nil), "magma.lte.FlowRecordQuery")
}

func init() { proto.RegisterFile("lte/protos/meteringd.proto", fileDescriptor_b88a6fd3c74575d2) }

var fileDescriptor_b88a6fd3c74575d2 = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x6d, 0xda, 0xf5, 0xc3, 0xb7, 0x8c, 0x0f, 0x6b, 0x13, 0x69, 0x10, 0x5a, 0x15, 0x09, 0xa9,
	0x12, 0x92, 0x2b, 0x95, 0x97, 0xf2, 0x86, 0xc6, 0x04, 0x8b, 0x80, 0x07, 0xb2, 0xc2, 0x03, 0x2f,
	0x55, 0x12, 0x7b, 0x91, 0xb5, 0xa4, 0x2e, 0xb6, 0xab, 0xa6, 0xff, 0x66, 0x7f, 0x8e, 0xff, 0x81,
	0xec, 0x7c, 0x4d, 0x50, 0xed, 0xa9, 0xbd, 0xe7, 0x9c, 0xf8, 0xde, 0x73, 0xec, 0x0b, 0x5e, 0xa6,
	0xd9, 0x7c, 0x2b, 0x85, 0x16, 0x6a, 0x9e, 0x33, 0xcd, 0x24, 0xdf, 0xa4, 0x94, 0x58, 0x00, 0xa3,
	0x3c, 0x4a, 0xf3, 0x88, 0x64, 0x9a, 0x79, 0x13, 0x21, 0x93, 0xa5, 0xac, 0x85, 0x89, 0xc8, 0x73,
	0xb1, 0x29, 0x55, 0xde, 0x45, 0x2a, 0x44, 0x9a, 0x55, 0x87, 0xc4, 0xbb, 0xdb, 0xb9, 0xe6, 0x39,
	0x53, 0x3a, 0xca, 0xb7, 0xa5, 0xc0, 0xbf, 0xef, 0x02, 0x7c, 0xca, 0xc4, 0x3e, 0x64, 0x89, 0x90,
	0x14, 0xcf, 0xa0, 0xcb, 0xa9, 0xeb, 0x4c, 0x9d, 0xd9, 0x78, 0xe1, 0x92, 0xa6, 0x05, 0x69, 0x25,
	0x24, 0xb8, 0x0a, 0xbb, 0x9c, 0xe2, 0xe7, 0xd0, 0x53, 0x9c, 0xba, 0xdd, 0xa9, 0x33, 0x43, 0xa1,
	0xf9, 0x8b, 0x5f, 0x03, 0xa4, 0x91, 0x66, 0xfb, 0xe8, 0xb0, 0xe6, 0xd4, 0xed, 0x59, 0x02, 0x55,
	0x48, 0x40, 0xf1, 0x04, 0x46, 0xf1, 0x41, 0x33, 0xb5, 0xd6, 0x85, 0xdb, 0x9f, 0x3a, 0xb3, 0x93,
	0x70, 0x68, 0xeb, 0x55, 0xd1, 0x52, 0xb2, 0x70, 0x07, 0x0f, 0xa8, 0xb0, 0xc0, 0x2f, 0x61, 0xb8,
	0xbd, 0xd3, 0xf6, 0xa3, 0xa1, 0x65, 0x06, 0xa6, 0x5c, 0xb5, 0x84, 0x2c, 0xdc, 0x51, 0x4b, 0x84,
	0x05, 0x7e, 0x0f, 0xa0, 0x74, 0x24, 0xf5, 0xda, 0x58, 0x75, 0x91, 0xb5, 0xe2, 0x91, 0x32, 0x07,
	0x52, 0xe7, 0x40, 0x56, 0x75, 0x0e, 0x21, 0xb2, 0x6a, 0x53, 0x7b, 0x67, 0xd0, 0x0d, 0xae, 0xf0,
	0xd3, 0x26, 0x03, 0x64, 0x9c, 0xfa, 0x04, 0x4e, 0x5b, 0xfb, 0x37, 0x4c, 0x1b, 0xa3, 0xd2, 0x16,
	0x6b, 0x4e, 0x95, 0xeb, 0x4c, 0x7b, 0xc6, 0x68, 0x89, 0x04, 0x54, 0xf9, 0x4b, 0x40, 0x46, 0xbf,
	0x8a, 0xe2, 0x8c, 0xe1, 0xb7, 0xd0, 0xbf, 0xcd, 0xc4, 0xbe, 0x94, 0x8d, 0x17, 0xe7, 0x47, 0x33,
	0x0d, 0x4b, 0x8d, 0x7f, 0xef, 0xc0, 0xb3, 0x16, 0xfd, 0xbe, 0x63, 0xf2, 0x60, 0x9a, 0x6d, 0x98,
	0xde, 0x0b, 0x79, 0xb7, 0x6e, 0xa6, 0x42, 0x15, 0x12, 0x98, 0xd0, 0x51, 0x33, 0x4b, 0x79, 0x19,
	0xd7, 0x9d, 0x70, 0x54, 0x0f, 0x83, 0x2f, 0xfe, 0xbf, 0x93, 0xeb, 0xce, 0xc3, 0x5b, 0x79, 0x03,
	0xa7, 0x6a, 0x17, 0xab, 0x44, 0xf2, 0x98, 0x49, 0xa3, 0x39, 0xa9, 0x34, 0x4f, 0x5a, 0x38, 0xa0,
	0x97, 0x43, 0xe8, 0xff, 0x36, 0xe3, 0x2c, 0xfe, 0x38, 0xe0, 0x7d, 0xab, 0x9f, 0x62, 0x39, 0xa7,
	0xfa, 0x28, 0x36, 0x5a, 0x8a, 0x2c, 0x63, 0x12, 0x7f, 0x00, 0xf4, 0x99, 0xe9, 0xea, 0x31, 0x79,
	0x47, 0xcd, 0x5a, 0x5b, 0xde, 0xf1, 0x20, 0xfc, 0x0e, 0xfe, 0x02, 0xe7, 0x5f, 0xb9, 0xd2, 0x37,
	0x4d, 0xf7, 0xaa, 0xc9, 0xa3, 0xa7, 0x9d, 0xfd, 0xc3, 0xd9, 0xec, 0xfd, 0x0e, 0x5e, 0xc2, 0xf8,
	0xc7, 0x96, 0x46, 0x9a, 0x19, 0x50, 0xe1, 0xa3, 0x32, 0xef, 0x45, 0x85, 0xda, 0x2d, 0x22, 0x3f,
	0x05, 0xa7, 0x7e, 0xe7, 0xf2, 0xd5, 0xaf, 0x89, 0x45, 0xe7, 0x66, 0x05, 0x93, 0x4c, 0xec, 0xe8,
	0x3c, 0x15, 0xd5, 0x8a, 0xc5, 0x03, 0xfb, 0xfb, 0xee, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x42,
	0x8c, 0xb8, 0x9d, 0xa0, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MeteringdRecordsControllerClient is the client API for MeteringdRecordsController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MeteringdRecordsControllerClient interface {
	// Get a flow record
	GetRecord(ctx context.Context, in *FlowRecordQuery, opts ...grpc.CallOption) (*FlowRecord, error)
	// Get all flow records for a subscriber
	ListSubscriberRecords(ctx context.Context, in *FlowRecordQuery, opts ...grpc.CallOption) (*FlowTable, error)
	// Update record of flows from gateway (has identity context)
	UpdateFlows(ctx context.Context, in *FlowTable, opts ...grpc.CallOption) (*protos.Void, error)
}

type meteringdRecordsControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewMeteringdRecordsControllerClient(cc grpc.ClientConnInterface) MeteringdRecordsControllerClient {
	return &meteringdRecordsControllerClient{cc}
}

func (c *meteringdRecordsControllerClient) GetRecord(ctx context.Context, in *FlowRecordQuery, opts ...grpc.CallOption) (*FlowRecord, error) {
	out := new(FlowRecord)
	err := c.cc.Invoke(ctx, "/magma.lte.MeteringdRecordsController/GetRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meteringdRecordsControllerClient) ListSubscriberRecords(ctx context.Context, in *FlowRecordQuery, opts ...grpc.CallOption) (*FlowTable, error) {
	out := new(FlowTable)
	err := c.cc.Invoke(ctx, "/magma.lte.MeteringdRecordsController/ListSubscriberRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meteringdRecordsControllerClient) UpdateFlows(ctx context.Context, in *FlowTable, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.lte.MeteringdRecordsController/UpdateFlows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeteringdRecordsControllerServer is the server API for MeteringdRecordsController service.
type MeteringdRecordsControllerServer interface {
	// Get a flow record
	GetRecord(context.Context, *FlowRecordQuery) (*FlowRecord, error)
	// Get all flow records for a subscriber
	ListSubscriberRecords(context.Context, *FlowRecordQuery) (*FlowTable, error)
	// Update record of flows from gateway (has identity context)
	UpdateFlows(context.Context, *FlowTable) (*protos.Void, error)
}

// UnimplementedMeteringdRecordsControllerServer can be embedded to have forward compatible implementations.
type UnimplementedMeteringdRecordsControllerServer struct {
}

func (*UnimplementedMeteringdRecordsControllerServer) GetRecord(ctx context.Context, req *FlowRecordQuery) (*FlowRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecord not implemented")
}
func (*UnimplementedMeteringdRecordsControllerServer) ListSubscriberRecords(ctx context.Context, req *FlowRecordQuery) (*FlowTable, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSubscriberRecords not implemented")
}
func (*UnimplementedMeteringdRecordsControllerServer) UpdateFlows(ctx context.Context, req *FlowTable) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFlows not implemented")
}

func RegisterMeteringdRecordsControllerServer(s *grpc.Server, srv MeteringdRecordsControllerServer) {
	s.RegisterService(&_MeteringdRecordsController_serviceDesc, srv)
}

func _MeteringdRecordsController_GetRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlowRecordQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeteringdRecordsControllerServer).GetRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.MeteringdRecordsController/GetRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeteringdRecordsControllerServer).GetRecord(ctx, req.(*FlowRecordQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeteringdRecordsController_ListSubscriberRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlowRecordQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeteringdRecordsControllerServer).ListSubscriberRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.MeteringdRecordsController/ListSubscriberRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeteringdRecordsControllerServer).ListSubscriberRecords(ctx, req.(*FlowRecordQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeteringdRecordsController_UpdateFlows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlowTable)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeteringdRecordsControllerServer).UpdateFlows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.MeteringdRecordsController/UpdateFlows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeteringdRecordsControllerServer).UpdateFlows(ctx, req.(*FlowTable))
	}
	return interceptor(ctx, in, info, handler)
}

var _MeteringdRecordsController_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.lte.MeteringdRecordsController",
	HandlerType: (*MeteringdRecordsControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRecord",
			Handler:    _MeteringdRecordsController_GetRecord_Handler,
		},
		{
			MethodName: "ListSubscriberRecords",
			Handler:    _MeteringdRecordsController_ListSubscriberRecords_Handler,
		},
		{
			MethodName: "UpdateFlows",
			Handler:    _MeteringdRecordsController_UpdateFlows_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lte/protos/meteringd.proto",
}
