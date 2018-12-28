// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/raft.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Entry struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_raft_cc55a4f80e3246ed, []int{0}
}
func (m *Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entry.Unmarshal(m, b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
}
func (dst *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(dst, src)
}
func (m *Entry) XXX_Size() int {
	return xxx_messageInfo_Entry.Size(m)
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type AppendEntries struct {
	Term                 int64    `protobuf:"varint,1,opt,name=term" json:"term,omitempty"`
	LeaderId             string   `protobuf:"bytes,2,opt,name=leaderId" json:"leaderId,omitempty"`
	PrevLogIndex         int64    `protobuf:"varint,3,opt,name=prevLogIndex" json:"prevLogIndex,omitempty"`
	Entries              []*Entry `protobuf:"bytes,4,rep,name=entries" json:"entries,omitempty"`
	LeaderCommit         int64    `protobuf:"varint,5,opt,name=leaderCommit" json:"leaderCommit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppendEntries) Reset()         { *m = AppendEntries{} }
func (m *AppendEntries) String() string { return proto.CompactTextString(m) }
func (*AppendEntries) ProtoMessage()    {}
func (*AppendEntries) Descriptor() ([]byte, []int) {
	return fileDescriptor_raft_cc55a4f80e3246ed, []int{1}
}
func (m *AppendEntries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppendEntries.Unmarshal(m, b)
}
func (m *AppendEntries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppendEntries.Marshal(b, m, deterministic)
}
func (dst *AppendEntries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppendEntries.Merge(dst, src)
}
func (m *AppendEntries) XXX_Size() int {
	return xxx_messageInfo_AppendEntries.Size(m)
}
func (m *AppendEntries) XXX_DiscardUnknown() {
	xxx_messageInfo_AppendEntries.DiscardUnknown(m)
}

var xxx_messageInfo_AppendEntries proto.InternalMessageInfo

func (m *AppendEntries) GetTerm() int64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *AppendEntries) GetLeaderId() string {
	if m != nil {
		return m.LeaderId
	}
	return ""
}

func (m *AppendEntries) GetPrevLogIndex() int64 {
	if m != nil {
		return m.PrevLogIndex
	}
	return 0
}

func (m *AppendEntries) GetEntries() []*Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *AppendEntries) GetLeaderCommit() int64 {
	if m != nil {
		return m.LeaderCommit
	}
	return 0
}

type AppendEntriesResult struct {
	Term                 int64    `protobuf:"varint,1,opt,name=term" json:"term,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppendEntriesResult) Reset()         { *m = AppendEntriesResult{} }
func (m *AppendEntriesResult) String() string { return proto.CompactTextString(m) }
func (*AppendEntriesResult) ProtoMessage()    {}
func (*AppendEntriesResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_raft_cc55a4f80e3246ed, []int{2}
}
func (m *AppendEntriesResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppendEntriesResult.Unmarshal(m, b)
}
func (m *AppendEntriesResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppendEntriesResult.Marshal(b, m, deterministic)
}
func (dst *AppendEntriesResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppendEntriesResult.Merge(dst, src)
}
func (m *AppendEntriesResult) XXX_Size() int {
	return xxx_messageInfo_AppendEntriesResult.Size(m)
}
func (m *AppendEntriesResult) XXX_DiscardUnknown() {
	xxx_messageInfo_AppendEntriesResult.DiscardUnknown(m)
}

var xxx_messageInfo_AppendEntriesResult proto.InternalMessageInfo

func (m *AppendEntriesResult) GetTerm() int64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *AppendEntriesResult) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type RequestVote struct {
	Term                 int64    `protobuf:"varint,1,opt,name=term" json:"term,omitempty"`
	CandidateId          int64    `protobuf:"varint,2,opt,name=candidateId" json:"candidateId,omitempty"`
	LastLogIndex         int64    `protobuf:"varint,3,opt,name=lastLogIndex" json:"lastLogIndex,omitempty"`
	LastLogTerm          int64    `protobuf:"varint,4,opt,name=lastLogTerm" json:"lastLogTerm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestVote) Reset()         { *m = RequestVote{} }
func (m *RequestVote) String() string { return proto.CompactTextString(m) }
func (*RequestVote) ProtoMessage()    {}
func (*RequestVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_raft_cc55a4f80e3246ed, []int{3}
}
func (m *RequestVote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestVote.Unmarshal(m, b)
}
func (m *RequestVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestVote.Marshal(b, m, deterministic)
}
func (dst *RequestVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestVote.Merge(dst, src)
}
func (m *RequestVote) XXX_Size() int {
	return xxx_messageInfo_RequestVote.Size(m)
}
func (m *RequestVote) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestVote.DiscardUnknown(m)
}

var xxx_messageInfo_RequestVote proto.InternalMessageInfo

func (m *RequestVote) GetTerm() int64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RequestVote) GetCandidateId() int64 {
	if m != nil {
		return m.CandidateId
	}
	return 0
}

func (m *RequestVote) GetLastLogIndex() int64 {
	if m != nil {
		return m.LastLogIndex
	}
	return 0
}

func (m *RequestVote) GetLastLogTerm() int64 {
	if m != nil {
		return m.LastLogTerm
	}
	return 0
}

type RequestVoteResult struct {
	Term                 int64    `protobuf:"varint,1,opt,name=term" json:"term,omitempty"`
	VoteGranted          bool     `protobuf:"varint,2,opt,name=voteGranted" json:"voteGranted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestVoteResult) Reset()         { *m = RequestVoteResult{} }
func (m *RequestVoteResult) String() string { return proto.CompactTextString(m) }
func (*RequestVoteResult) ProtoMessage()    {}
func (*RequestVoteResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_raft_cc55a4f80e3246ed, []int{4}
}
func (m *RequestVoteResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestVoteResult.Unmarshal(m, b)
}
func (m *RequestVoteResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestVoteResult.Marshal(b, m, deterministic)
}
func (dst *RequestVoteResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestVoteResult.Merge(dst, src)
}
func (m *RequestVoteResult) XXX_Size() int {
	return xxx_messageInfo_RequestVoteResult.Size(m)
}
func (m *RequestVoteResult) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestVoteResult.DiscardUnknown(m)
}

var xxx_messageInfo_RequestVoteResult proto.InternalMessageInfo

func (m *RequestVoteResult) GetTerm() int64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RequestVoteResult) GetVoteGranted() bool {
	if m != nil {
		return m.VoteGranted
	}
	return false
}

func init() {
	proto.RegisterType((*Entry)(nil), "proto.Entry")
	proto.RegisterType((*AppendEntries)(nil), "proto.AppendEntries")
	proto.RegisterType((*AppendEntriesResult)(nil), "proto.AppendEntriesResult")
	proto.RegisterType((*RequestVote)(nil), "proto.RequestVote")
	proto.RegisterType((*RequestVoteResult)(nil), "proto.RequestVoteResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RaftClient is the client API for Raft service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RaftClient interface {
	AppendEntriesRPC(ctx context.Context, in *AppendEntries, opts ...grpc.CallOption) (*AppendEntriesResult, error)
	RequestVoteRPC(ctx context.Context, in *RequestVote, opts ...grpc.CallOption) (*RequestVoteResult, error)
}

type raftClient struct {
	cc *grpc.ClientConn
}

func NewRaftClient(cc *grpc.ClientConn) RaftClient {
	return &raftClient{cc}
}

func (c *raftClient) AppendEntriesRPC(ctx context.Context, in *AppendEntries, opts ...grpc.CallOption) (*AppendEntriesResult, error) {
	out := new(AppendEntriesResult)
	err := c.cc.Invoke(ctx, "/proto.Raft/AppendEntriesRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftClient) RequestVoteRPC(ctx context.Context, in *RequestVote, opts ...grpc.CallOption) (*RequestVoteResult, error) {
	out := new(RequestVoteResult)
	err := c.cc.Invoke(ctx, "/proto.Raft/RequestVoteRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Raft service

type RaftServer interface {
	AppendEntriesRPC(context.Context, *AppendEntries) (*AppendEntriesResult, error)
	RequestVoteRPC(context.Context, *RequestVote) (*RequestVoteResult, error)
}

func RegisterRaftServer(s *grpc.Server, srv RaftServer) {
	s.RegisterService(&_Raft_serviceDesc, srv)
}

func _Raft_AppendEntriesRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendEntries)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).AppendEntriesRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Raft/AppendEntriesRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).AppendEntriesRPC(ctx, req.(*AppendEntries))
	}
	return interceptor(ctx, in, info, handler)
}

func _Raft_RequestVoteRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVote)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).RequestVoteRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Raft/RequestVoteRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).RequestVoteRPC(ctx, req.(*RequestVote))
	}
	return interceptor(ctx, in, info, handler)
}

var _Raft_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Raft",
	HandlerType: (*RaftServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppendEntriesRPC",
			Handler:    _Raft_AppendEntriesRPC_Handler,
		},
		{
			MethodName: "RequestVoteRPC",
			Handler:    _Raft_RequestVoteRPC_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/raft.proto",
}

func init() { proto.RegisterFile("proto/raft.proto", fileDescriptor_raft_cc55a4f80e3246ed) }

var fileDescriptor_raft_cc55a4f80e3246ed = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0xc9, 0x3f, 0xe9, 0xbf, 0x75, 0x52, 0xa5, 0xae, 0x1e, 0x96, 0x78, 0x09, 0x39, 0x48,
	0x4f, 0x15, 0xea, 0x5d, 0x90, 0x28, 0x52, 0xf0, 0x50, 0x16, 0xf1, 0xbe, 0x76, 0xa7, 0x52, 0x68,
	0x93, 0xb8, 0x3b, 0x2d, 0xfa, 0x08, 0x82, 0xef, 0xe2, 0x2b, 0x4a, 0x36, 0x5b, 0xd9, 0x68, 0xf0,
	0x94, 0x99, 0x6f, 0x66, 0xbe, 0xfc, 0xbe, 0x85, 0x51, 0xa5, 0x4b, 0x2a, 0x2f, 0xb4, 0x5c, 0xd2,
	0xc4, 0x96, 0xac, 0x67, 0x3f, 0xd9, 0x19, 0xf4, 0x6e, 0x0b, 0xd2, 0x6f, 0x8c, 0x41, 0xa4, 0x24,
	0x49, 0x1e, 0xa4, 0xc1, 0x78, 0x28, 0x6c, 0x9d, 0x7d, 0x06, 0x70, 0x78, 0x5d, 0x55, 0x58, 0xa8,
	0x7a, 0x67, 0x85, 0xa6, 0xde, 0x22, 0xd4, 0x1b, 0xbb, 0x15, 0x0a, 0x5b, 0xb3, 0x04, 0x06, 0x6b,
	0x94, 0x0a, 0xf5, 0x4c, 0xf1, 0x7f, 0x69, 0x30, 0x3e, 0x10, 0xdf, 0x3d, 0xcb, 0x60, 0x58, 0x69,
	0xdc, 0xdd, 0x97, 0xcf, 0xb3, 0x42, 0xe1, 0x2b, 0x0f, 0xed, 0x5d, 0x4b, 0x63, 0xe7, 0xd0, 0xc7,
	0xc6, 0x9e, 0x47, 0x69, 0x38, 0x8e, 0xa7, 0xc3, 0x06, 0x71, 0x62, 0xc1, 0xc4, 0x7e, 0x58, 0x7b,
	0x35, 0xbe, 0x79, 0xb9, 0xd9, 0xac, 0x88, 0xf7, 0x1a, 0x2f, 0x5f, 0xcb, 0x72, 0x38, 0x69, 0x01,
	0x0b, 0x34, 0xdb, 0x35, 0x75, 0x62, 0x73, 0xe8, 0x9b, 0xed, 0x62, 0x81, 0xc6, 0x58, 0xea, 0x81,
	0xd8, 0xb7, 0xd9, 0x7b, 0x00, 0xb1, 0xc0, 0x97, 0x2d, 0x1a, 0x7a, 0x2c, 0x09, 0x3b, 0xaf, 0x53,
	0x88, 0x17, 0xb2, 0x50, 0x2b, 0x25, 0x09, 0x5d, 0xee, 0x50, 0xf8, 0x92, 0xc5, 0x95, 0x86, 0x7e,
	0x46, 0xf7, 0xb5, 0xda, 0xc5, 0xf5, 0x0f, 0xf5, 0x0f, 0xa2, 0xc6, 0xc5, 0x93, 0xb2, 0x19, 0x1c,
	0x7b, 0x28, 0x7f, 0xc4, 0x49, 0x21, 0xde, 0x95, 0x84, 0x77, 0x5a, 0x16, 0x84, 0xca, 0x45, 0xf2,
	0xa5, 0xe9, 0x47, 0x00, 0x91, 0x90, 0x4b, 0x62, 0x37, 0x30, 0x6a, 0x3f, 0xd2, 0x3c, 0x67, 0xa7,
	0xee, 0xcd, 0x5b, 0x83, 0x24, 0xe9, 0x52, 0x1d, 0xc4, 0x15, 0x1c, 0xf9, 0x64, 0xf3, 0x9c, 0x31,
	0xb7, 0xed, 0xc9, 0x09, 0xff, 0xad, 0x35, 0xf7, 0x4f, 0xff, 0xed, 0xe0, 0xf2, 0x2b, 0x00, 0x00,
	0xff, 0xff, 0x3a, 0xa5, 0x72, 0x88, 0x9b, 0x02, 0x00, 0x00,
}