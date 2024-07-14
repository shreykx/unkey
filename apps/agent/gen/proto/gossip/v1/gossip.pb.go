// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: proto/gossip/v1/gossip.proto

package gossipv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type State int32

const (
	State_State_UNSPECIFIED State = 0
	State_State_ALIVE       State = 1
	State_State_DEAD        State = 2
	State_State_LEFT        State = 3
	State_State_SUSPECT     State = 4
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "State_UNSPECIFIED",
		1: "State_ALIVE",
		2: "State_DEAD",
		3: "State_LEFT",
		4: "State_SUSPECT",
	}
	State_value = map[string]int32{
		"State_UNSPECIFIED": 0,
		"State_ALIVE":       1,
		"State_DEAD":        2,
		"State_LEFT":        3,
		"State_SUSPECT":     4,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_gossip_v1_gossip_proto_enumTypes[0].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_proto_gossip_v1_gossip_proto_enumTypes[0]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_proto_gossip_v1_gossip_proto_rawDescGZIP(), []int{0}
}

type Rumor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time int64 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *Rumor) Reset() {
	*x = Rumor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gossip_v1_gossip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rumor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rumor) ProtoMessage() {}

func (x *Rumor) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gossip_v1_gossip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rumor.ProtoReflect.Descriptor instead.
func (*Rumor) Descriptor() ([]byte, []int) {
	return file_proto_gossip_v1_gossip_proto_rawDescGZIP(), []int{0}
}

func (x *Rumor) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

type GossipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GossipRequest) Reset() {
	*x = GossipRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gossip_v1_gossip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GossipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GossipRequest) ProtoMessage() {}

func (x *GossipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gossip_v1_gossip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GossipRequest.ProtoReflect.Descriptor instead.
func (*GossipRequest) Descriptor() ([]byte, []int) {
	return file_proto_gossip_v1_gossip_proto_rawDescGZIP(), []int{1}
}

type GossipResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GossipResponse) Reset() {
	*x = GossipResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gossip_v1_gossip_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GossipResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GossipResponse) ProtoMessage() {}

func (x *GossipResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gossip_v1_gossip_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GossipResponse.ProtoReflect.Descriptor instead.
func (*GossipResponse) Descriptor() ([]byte, []int) {
	return file_proto_gossip_v1_gossip_proto_rawDescGZIP(), []int{2}
}

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gossip_v1_gossip_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gossip_v1_gossip_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_proto_gossip_v1_gossip_proto_rawDescGZIP(), []int{3}
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State State `protobuf:"varint,1,opt,name=state,proto3,enum=gossip.v1.State" json:"state,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gossip_v1_gossip_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gossip_v1_gossip_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_proto_gossip_v1_gossip_proto_rawDescGZIP(), []int{4}
}

func (x *PingResponse) GetState() State {
	if x != nil {
		return x.State
	}
	return State_State_UNSPECIFIED
}

type IndirectPingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId  string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	RpcAddr string `protobuf:"bytes,2,opt,name=rpc_addr,json=rpcAddr,proto3" json:"rpc_addr,omitempty"`
}

func (x *IndirectPingRequest) Reset() {
	*x = IndirectPingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gossip_v1_gossip_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndirectPingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndirectPingRequest) ProtoMessage() {}

func (x *IndirectPingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gossip_v1_gossip_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndirectPingRequest.ProtoReflect.Descriptor instead.
func (*IndirectPingRequest) Descriptor() ([]byte, []int) {
	return file_proto_gossip_v1_gossip_proto_rawDescGZIP(), []int{5}
}

func (x *IndirectPingRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *IndirectPingRequest) GetRpcAddr() string {
	if x != nil {
		return x.RpcAddr
	}
	return ""
}

type IndirectPingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State State `protobuf:"varint,1,opt,name=state,proto3,enum=gossip.v1.State" json:"state,omitempty"`
}

func (x *IndirectPingResponse) Reset() {
	*x = IndirectPingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gossip_v1_gossip_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndirectPingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndirectPingResponse) ProtoMessage() {}

func (x *IndirectPingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gossip_v1_gossip_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndirectPingResponse.ProtoReflect.Descriptor instead.
func (*IndirectPingResponse) Descriptor() ([]byte, []int) {
	return file_proto_gossip_v1_gossip_proto_rawDescGZIP(), []int{6}
}

func (x *IndirectPingResponse) GetState() State {
	if x != nil {
		return x.State
	}
	return State_State_UNSPECIFIED
}

type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId  string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	RpcAddr string `protobuf:"bytes,2,opt,name=rpc_addr,json=rpcAddr,proto3" json:"rpc_addr,omitempty"`
	State   State  `protobuf:"varint,3,opt,name=state,proto3,enum=gossip.v1.State" json:"state,omitempty"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gossip_v1_gossip_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gossip_v1_gossip_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_proto_gossip_v1_gossip_proto_rawDescGZIP(), []int{7}
}

func (x *Member) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *Member) GetRpcAddr() string {
	if x != nil {
		return x.RpcAddr
	}
	return ""
}

func (x *Member) GetState() State {
	if x != nil {
		return x.State
	}
	return State_State_UNSPECIFIED
}

type SyncMembersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The members that the sender knows about
	Members []*Member `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *SyncMembersRequest) Reset() {
	*x = SyncMembersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gossip_v1_gossip_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncMembersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncMembersRequest) ProtoMessage() {}

func (x *SyncMembersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gossip_v1_gossip_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncMembersRequest.ProtoReflect.Descriptor instead.
func (*SyncMembersRequest) Descriptor() ([]byte, []int) {
	return file_proto_gossip_v1_gossip_proto_rawDescGZIP(), []int{8}
}

func (x *SyncMembersRequest) GetMembers() []*Member {
	if x != nil {
		return x.Members
	}
	return nil
}

type SyncMembersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The members that the receiver knows about
	Members []*Member `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *SyncMembersResponse) Reset() {
	*x = SyncMembersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gossip_v1_gossip_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncMembersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncMembersResponse) ProtoMessage() {}

func (x *SyncMembersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gossip_v1_gossip_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncMembersResponse.ProtoReflect.Descriptor instead.
func (*SyncMembersResponse) Descriptor() ([]byte, []int) {
	return file_proto_gossip_v1_gossip_proto_rawDescGZIP(), []int{9}
}

func (x *SyncMembersResponse) GetMembers() []*Member {
	if x != nil {
		return x.Members
	}
	return nil
}

var File_proto_gossip_v1_gossip_proto protoreflect.FileDescriptor

var file_proto_gossip_v1_gossip_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2f, 0x76,
	0x31, 0x2f, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x76, 0x31, 0x22, 0x1b, 0x0a, 0x05, 0x52, 0x75, 0x6d,
	0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x6f, 0x73, 0x73, 0x69,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x50, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x36, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x22, 0x49, 0x0a, 0x13, 0x49, 0x6e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x72, 0x70, 0x63, 0x41, 0x64, 0x64, 0x72, 0x22, 0x3e, 0x0a, 0x14, 0x49,
	0x6e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x64, 0x0a, 0x06, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x72, 0x70, 0x63, 0x41, 0x64, 0x64, 0x72, 0x12, 0x26, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x73, 0x73, 0x69,
	0x70, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x22, 0x41, 0x0a, 0x12, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x73, 0x73, 0x69,
	0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x07, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x22, 0x42, 0x0a, 0x13, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67,
	0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2a, 0x62, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x5f, 0x41, 0x4c, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x5f, 0x44, 0x45, 0x41, 0x44, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x5f, 0x4c, 0x45, 0x46, 0x54, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x5f, 0x53, 0x55, 0x53, 0x50, 0x45, 0x43, 0x54, 0x10, 0x04, 0x32, 0xed, 0x01, 0x0a,
	0x0d, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39,
	0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0c, 0x49, 0x6e, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x73, 0x73,
	0x69, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x6f, 0x73, 0x73,
	0x69, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0b,
	0x53, 0x79, 0x6e, 0x63, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x2e, 0x67, 0x6f,
	0x73, 0x73, 0x69, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x6f, 0x73,
	0x73, 0x69, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x42, 0x5a, 0x40,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x6e, 0x6b, 0x65, 0x79,
	0x65, 0x64, 0x2f, 0x75, 0x6e, 0x6b, 0x65, 0x79, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x73, 0x73, 0x69, 0x70, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_gossip_v1_gossip_proto_rawDescOnce sync.Once
	file_proto_gossip_v1_gossip_proto_rawDescData = file_proto_gossip_v1_gossip_proto_rawDesc
)

func file_proto_gossip_v1_gossip_proto_rawDescGZIP() []byte {
	file_proto_gossip_v1_gossip_proto_rawDescOnce.Do(func() {
		file_proto_gossip_v1_gossip_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_gossip_v1_gossip_proto_rawDescData)
	})
	return file_proto_gossip_v1_gossip_proto_rawDescData
}

var file_proto_gossip_v1_gossip_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_gossip_v1_gossip_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_gossip_v1_gossip_proto_goTypes = []any{
	(State)(0),                   // 0: gossip.v1.State
	(*Rumor)(nil),                // 1: gossip.v1.Rumor
	(*GossipRequest)(nil),        // 2: gossip.v1.GossipRequest
	(*GossipResponse)(nil),       // 3: gossip.v1.GossipResponse
	(*PingRequest)(nil),          // 4: gossip.v1.PingRequest
	(*PingResponse)(nil),         // 5: gossip.v1.PingResponse
	(*IndirectPingRequest)(nil),  // 6: gossip.v1.IndirectPingRequest
	(*IndirectPingResponse)(nil), // 7: gossip.v1.IndirectPingResponse
	(*Member)(nil),               // 8: gossip.v1.Member
	(*SyncMembersRequest)(nil),   // 9: gossip.v1.SyncMembersRequest
	(*SyncMembersResponse)(nil),  // 10: gossip.v1.SyncMembersResponse
}
var file_proto_gossip_v1_gossip_proto_depIdxs = []int32{
	0,  // 0: gossip.v1.PingResponse.state:type_name -> gossip.v1.State
	0,  // 1: gossip.v1.IndirectPingResponse.state:type_name -> gossip.v1.State
	0,  // 2: gossip.v1.Member.state:type_name -> gossip.v1.State
	8,  // 3: gossip.v1.SyncMembersRequest.members:type_name -> gossip.v1.Member
	8,  // 4: gossip.v1.SyncMembersResponse.members:type_name -> gossip.v1.Member
	4,  // 5: gossip.v1.GossipService.Ping:input_type -> gossip.v1.PingRequest
	6,  // 6: gossip.v1.GossipService.IndirectPing:input_type -> gossip.v1.IndirectPingRequest
	9,  // 7: gossip.v1.GossipService.SyncMembers:input_type -> gossip.v1.SyncMembersRequest
	5,  // 8: gossip.v1.GossipService.Ping:output_type -> gossip.v1.PingResponse
	7,  // 9: gossip.v1.GossipService.IndirectPing:output_type -> gossip.v1.IndirectPingResponse
	10, // 10: gossip.v1.GossipService.SyncMembers:output_type -> gossip.v1.SyncMembersResponse
	8,  // [8:11] is the sub-list for method output_type
	5,  // [5:8] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_gossip_v1_gossip_proto_init() }
func file_proto_gossip_v1_gossip_proto_init() {
	if File_proto_gossip_v1_gossip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_gossip_v1_gossip_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Rumor); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_gossip_v1_gossip_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GossipRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_gossip_v1_gossip_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GossipResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_gossip_v1_gossip_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PingRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_gossip_v1_gossip_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*PingResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_gossip_v1_gossip_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*IndirectPingRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_gossip_v1_gossip_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*IndirectPingResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_gossip_v1_gossip_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Member); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_gossip_v1_gossip_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*SyncMembersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_gossip_v1_gossip_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*SyncMembersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_gossip_v1_gossip_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_gossip_v1_gossip_proto_goTypes,
		DependencyIndexes: file_proto_gossip_v1_gossip_proto_depIdxs,
		EnumInfos:         file_proto_gossip_v1_gossip_proto_enumTypes,
		MessageInfos:      file_proto_gossip_v1_gossip_proto_msgTypes,
	}.Build()
	File_proto_gossip_v1_gossip_proto = out.File
	file_proto_gossip_v1_gossip_proto_rawDesc = nil
	file_proto_gossip_v1_gossip_proto_goTypes = nil
	file_proto_gossip_v1_gossip_proto_depIdxs = nil
}
