// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: event.proto

package event

import (
	topo "github.com/drivenets/kne/proto/topo"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Cluster_ClusterType int32

const (
	Cluster_CLUSTER_TYPE_UNSPECIFIED Cluster_ClusterType = 0
	Cluster_CLUSTER_TYPE_KIND        Cluster_ClusterType = 1
	Cluster_CLUSTER_TYPE_EXTERNAL    Cluster_ClusterType = 2
)

// Enum value maps for Cluster_ClusterType.
var (
	Cluster_ClusterType_name = map[int32]string{
		0: "CLUSTER_TYPE_UNSPECIFIED",
		1: "CLUSTER_TYPE_KIND",
		2: "CLUSTER_TYPE_EXTERNAL",
	}
	Cluster_ClusterType_value = map[string]int32{
		"CLUSTER_TYPE_UNSPECIFIED": 0,
		"CLUSTER_TYPE_KIND":        1,
		"CLUSTER_TYPE_EXTERNAL":    2,
	}
)

func (x Cluster_ClusterType) Enum() *Cluster_ClusterType {
	p := new(Cluster_ClusterType)
	*p = x
	return p
}

func (x Cluster_ClusterType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Cluster_ClusterType) Descriptor() protoreflect.EnumDescriptor {
	return file_event_proto_enumTypes[0].Descriptor()
}

func (Cluster_ClusterType) Type() protoreflect.EnumType {
	return &file_event_proto_enumTypes[0]
}

func (x Cluster_ClusterType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Cluster_ClusterType.Descriptor instead.
func (Cluster_ClusterType) EnumDescriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{5, 0}
}

type Cluster_IngressType int32

const (
	Cluster_INGRESS_TYPE_UNSPECIFIED Cluster_IngressType = 0
	Cluster_INGRESS_TYPE_METALLB     Cluster_IngressType = 1
)

// Enum value maps for Cluster_IngressType.
var (
	Cluster_IngressType_name = map[int32]string{
		0: "INGRESS_TYPE_UNSPECIFIED",
		1: "INGRESS_TYPE_METALLB",
	}
	Cluster_IngressType_value = map[string]int32{
		"INGRESS_TYPE_UNSPECIFIED": 0,
		"INGRESS_TYPE_METALLB":     1,
	}
)

func (x Cluster_IngressType) Enum() *Cluster_IngressType {
	p := new(Cluster_IngressType)
	*p = x
	return p
}

func (x Cluster_IngressType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Cluster_IngressType) Descriptor() protoreflect.EnumDescriptor {
	return file_event_proto_enumTypes[1].Descriptor()
}

func (Cluster_IngressType) Type() protoreflect.EnumType {
	return &file_event_proto_enumTypes[1]
}

func (x Cluster_IngressType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Cluster_IngressType.Descriptor instead.
func (Cluster_IngressType) EnumDescriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{5, 1}
}

type Cluster_CNIType int32

const (
	Cluster_CNI_TYPE_UNSPECIFIED Cluster_CNIType = 0
	Cluster_CNI_TYPE_MESHNET     Cluster_CNIType = 1
)

// Enum value maps for Cluster_CNIType.
var (
	Cluster_CNIType_name = map[int32]string{
		0: "CNI_TYPE_UNSPECIFIED",
		1: "CNI_TYPE_MESHNET",
	}
	Cluster_CNIType_value = map[string]int32{
		"CNI_TYPE_UNSPECIFIED": 0,
		"CNI_TYPE_MESHNET":     1,
	}
)

func (x Cluster_CNIType) Enum() *Cluster_CNIType {
	p := new(Cluster_CNIType)
	*p = x
	return p
}

func (x Cluster_CNIType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Cluster_CNIType) Descriptor() protoreflect.EnumDescriptor {
	return file_event_proto_enumTypes[2].Descriptor()
}

func (Cluster_CNIType) Type() protoreflect.EnumType {
	return &file_event_proto_enumTypes[2]
}

func (x Cluster_CNIType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Cluster_CNIType.Descriptor instead.
func (Cluster_CNIType) EnumDescriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{5, 2}
}

type Cluster_ControllerType int32

const (
	Cluster_CONTROLLER_TYPE_UNSPECIFIED Cluster_ControllerType = 0
	Cluster_CONTROLLER_TYPE_IXIATG      Cluster_ControllerType = 1
	Cluster_CONTROLLER_TYPE_SRLINUX     Cluster_ControllerType = 2
	Cluster_CONTROLLER_TYPE_CEOSLAB     Cluster_ControllerType = 3
	Cluster_CONTROLLER_TYPE_LEMMING     Cluster_ControllerType = 4
	Cluster_CONTROLLER_TYPE_CDNOS	    Cluster_ControllerType = 5
)

// Enum value maps for Cluster_ControllerType.
var (
	Cluster_ControllerType_name = map[int32]string{
		0: "CONTROLLER_TYPE_UNSPECIFIED",
		1: "CONTROLLER_TYPE_IXIATG",
		2: "CONTROLLER_TYPE_SRLINUX",
		3: "CONTROLLER_TYPE_CEOSLAB",
		4: "CONTROLLER_TYPE_LEMMING",
		5: "CONTROLLER_TYPE_CDNOS",
	}
	Cluster_ControllerType_value = map[string]int32{
		"CONTROLLER_TYPE_UNSPECIFIED": 0,
		"CONTROLLER_TYPE_IXIATG":      1,
		"CONTROLLER_TYPE_SRLINUX":     2,
		"CONTROLLER_TYPE_CEOSLAB":     3,
		"CONTROLLER_TYPE_LEMMING":     4,
		"CONTROLLER_TYPE_CDNOS":       5,
	}
)

func (x Cluster_ControllerType) Enum() *Cluster_ControllerType {
	p := new(Cluster_ControllerType)
	*p = x
	return p
}

func (x Cluster_ControllerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Cluster_ControllerType) Descriptor() protoreflect.EnumDescriptor {
	return file_event_proto_enumTypes[3].Descriptor()
}

func (Cluster_ControllerType) Type() protoreflect.EnumType {
	return &file_event_proto_enumTypes[3]
}

func (x Cluster_ControllerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Cluster_ControllerType.Descriptor instead.
func (Cluster_ControllerType) EnumDescriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{5, 3}
}

// KNEEvent is a wrapper around a specific event with a unique ID and timestamp.
type KNEEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid      string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are assignable to Event:
	//
	//	*KNEEvent_DeployClusterStart
	//	*KNEEvent_DeployClusterEnd
	//	*KNEEvent_CreateTopologyStart
	//	*KNEEvent_CreateTopologyEnd
	Event isKNEEvent_Event `protobuf_oneof:"event"`
}

func (x *KNEEvent) Reset() {
	*x = KNEEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KNEEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KNEEvent) ProtoMessage() {}

func (x *KNEEvent) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KNEEvent.ProtoReflect.Descriptor instead.
func (*KNEEvent) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{0}
}

func (x *KNEEvent) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *KNEEvent) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (m *KNEEvent) GetEvent() isKNEEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *KNEEvent) GetDeployClusterStart() *DeployClusterStart {
	if x, ok := x.GetEvent().(*KNEEvent_DeployClusterStart); ok {
		return x.DeployClusterStart
	}
	return nil
}

func (x *KNEEvent) GetDeployClusterEnd() *DeployClusterEnd {
	if x, ok := x.GetEvent().(*KNEEvent_DeployClusterEnd); ok {
		return x.DeployClusterEnd
	}
	return nil
}

func (x *KNEEvent) GetCreateTopologyStart() *CreateTopologyStart {
	if x, ok := x.GetEvent().(*KNEEvent_CreateTopologyStart); ok {
		return x.CreateTopologyStart
	}
	return nil
}

func (x *KNEEvent) GetCreateTopologyEnd() *CreateTopologyEnd {
	if x, ok := x.GetEvent().(*KNEEvent_CreateTopologyEnd); ok {
		return x.CreateTopologyEnd
	}
	return nil
}

type isKNEEvent_Event interface {
	isKNEEvent_Event()
}

type KNEEvent_DeployClusterStart struct {
	DeployClusterStart *DeployClusterStart `protobuf:"bytes,3,opt,name=deploy_cluster_start,json=deployClusterStart,proto3,oneof"`
}

type KNEEvent_DeployClusterEnd struct {
	DeployClusterEnd *DeployClusterEnd `protobuf:"bytes,4,opt,name=deploy_cluster_end,json=deployClusterEnd,proto3,oneof"`
}

type KNEEvent_CreateTopologyStart struct {
	CreateTopologyStart *CreateTopologyStart `protobuf:"bytes,5,opt,name=create_topology_start,json=createTopologyStart,proto3,oneof"`
}

type KNEEvent_CreateTopologyEnd struct {
	CreateTopologyEnd *CreateTopologyEnd `protobuf:"bytes,6,opt,name=create_topology_end,json=createTopologyEnd,proto3,oneof"`
}

func (*KNEEvent_DeployClusterStart) isKNEEvent_Event() {}

func (*KNEEvent_DeployClusterEnd) isKNEEvent_Event() {}

func (*KNEEvent_CreateTopologyStart) isKNEEvent_Event() {}

func (*KNEEvent_CreateTopologyEnd) isKNEEvent_Event() {}

// DeployClusterStart is an event indicating a cluster deployment was started.
type DeployClusterStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster *Cluster `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *DeployClusterStart) Reset() {
	*x = DeployClusterStart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployClusterStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployClusterStart) ProtoMessage() {}

func (x *DeployClusterStart) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployClusterStart.ProtoReflect.Descriptor instead.
func (*DeployClusterStart) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{1}
}

func (x *DeployClusterStart) GetCluster() *Cluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

// DeployClusterEnd is an event indicating a cluster deployment was ended.
type DeployClusterEnd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// error is a string containing the error message that caused the cluster
	// deployment to end unsuccessfully. Empty string if no error.
	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeployClusterEnd) Reset() {
	*x = DeployClusterEnd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployClusterEnd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployClusterEnd) ProtoMessage() {}

func (x *DeployClusterEnd) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployClusterEnd.ProtoReflect.Descriptor instead.
func (*DeployClusterEnd) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{2}
}

func (x *DeployClusterEnd) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// CreateTopologyStart is an event indicating a topology creation was started.
type CreateTopologyStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topology *Topology `protobuf:"bytes,1,opt,name=topology,proto3" json:"topology,omitempty"`
}

func (x *CreateTopologyStart) Reset() {
	*x = CreateTopologyStart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTopologyStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTopologyStart) ProtoMessage() {}

func (x *CreateTopologyStart) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTopologyStart.ProtoReflect.Descriptor instead.
func (*CreateTopologyStart) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTopologyStart) GetTopology() *Topology {
	if x != nil {
		return x.Topology
	}
	return nil
}

// CreateTopologyEnd is an event indicating a topology creation was ended.
type CreateTopologyEnd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// error is a string containing the error message that caused the topology
	// creation to end unsuccessfully. Empty string if no error.
	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateTopologyEnd) Reset() {
	*x = CreateTopologyEnd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTopologyEnd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTopologyEnd) ProtoMessage() {}

func (x *CreateTopologyEnd) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTopologyEnd.ProtoReflect.Descriptor instead.
func (*CreateTopologyEnd) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{4}
}

func (x *CreateTopologyEnd) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// Cluster holds information about a cluster.
type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster     Cluster_ClusterType      `protobuf:"varint,1,opt,name=cluster,proto3,enum=event.Cluster_ClusterType" json:"cluster,omitempty"`
	Ingress     Cluster_IngressType      `protobuf:"varint,2,opt,name=ingress,proto3,enum=event.Cluster_IngressType" json:"ingress,omitempty"`
	Cni         Cluster_CNIType          `protobuf:"varint,3,opt,name=cni,proto3,enum=event.Cluster_CNIType" json:"cni,omitempty"`
	Controllers []Cluster_ControllerType `protobuf:"varint,4,rep,packed,name=controllers,proto3,enum=event.Cluster_ControllerType" json:"controllers,omitempty"`
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{5}
}

func (x *Cluster) GetCluster() Cluster_ClusterType {
	if x != nil {
		return x.Cluster
	}
	return Cluster_CLUSTER_TYPE_UNSPECIFIED
}

func (x *Cluster) GetIngress() Cluster_IngressType {
	if x != nil {
		return x.Ingress
	}
	return Cluster_INGRESS_TYPE_UNSPECIFIED
}

func (x *Cluster) GetCni() Cluster_CNIType {
	if x != nil {
		return x.Cni
	}
	return Cluster_CNI_TYPE_UNSPECIFIED
}

func (x *Cluster) GetControllers() []Cluster_ControllerType {
	if x != nil {
		return x.Controllers
	}
	return nil
}

// Topology holds information about a topology.
type Topology struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// nodes in the topology.
	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	// number of links in the topology. A count is sufficient here as the link
	// details are not needed in topology metrics, in an effort to include as
	// little potentially identifying information as possible in the reporting.
	LinkCount int64 `protobuf:"varint,2,opt,name=link_count,json=linkCount,proto3" json:"link_count,omitempty"`
}

func (x *Topology) Reset() {
	*x = Topology{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Topology) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Topology) ProtoMessage() {}

func (x *Topology) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Topology.ProtoReflect.Descriptor instead.
func (*Topology) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{6}
}

func (x *Topology) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *Topology) GetLinkCount() int64 {
	if x != nil {
		return x.LinkCount
	}
	return 0
}

// Node holds information about a node in a topology. This is a simplified
// version of the topo.Node in an effort to include as little potentially
// identifying information as possible in reporting. The included fields
// will be used in aggregate to identify trends and determine reliability
// of certain node configurations.
type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// vendor is an enum of the node vendor.
	Vendor topo.Vendor `protobuf:"varint,1,opt,name=vendor,proto3,enum=topo.Vendor" json:"vendor,omitempty"`
	// model is the name of the model of the node.
	Model string `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{7}
}

func (x *Node) GetVendor() topo.Vendor {
	if x != nil {
		return x.Vendor
	}
	return topo.Vendor(0)
}

func (x *Node) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

var File_event_proto protoreflect.FileDescriptor

var file_event_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x1a, 0x0a, 0x74, 0x6f, 0x70, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x97, 0x03, 0x0a, 0x08, 0x4b, 0x4e, 0x45, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x4d, 0x0a, 0x14,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x48, 0x00, 0x52, 0x12, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x47, 0x0a, 0x12, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x65, 0x6e,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x64,
	0x48, 0x00, 0x52, 0x10, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x45, 0x6e, 0x64, 0x12, 0x50, 0x0a, 0x15, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x53, 0x74, 0x61, 0x72, 0x74, 0x48,
	0x00, 0x52, 0x13, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67,
	0x79, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x4a, 0x0a, 0x13, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x45, 0x6e, 0x64, 0x48, 0x00, 0x52,
	0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x45,
	0x6e, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x3e, 0x0a, 0x12, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x28, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x28, 0x0a, 0x10, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x42, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2b, 0x0a, 0x08,
	0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x52,
	0x08, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x22, 0x29, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x45, 0x6e, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0xe8, 0x04, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x34, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1a, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x07, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x07, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x28, 0x0a, 0x03,
	0x63, 0x6e, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x4e, 0x49, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x03, 0x63, 0x6e, 0x69, 0x12, 0x3f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x22, 0x5d, 0x0a, 0x0b, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45,
	0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x43,
	0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x58, 0x54, 0x45,
	0x52, 0x4e, 0x41, 0x4c, 0x10, 0x02, 0x22, 0x45, 0x0a, 0x0b, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x49, 0x4e, 0x47, 0x52, 0x45, 0x53, 0x53,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x4e, 0x47, 0x52, 0x45, 0x53, 0x53, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x4c, 0x4c, 0x42, 0x10, 0x01, 0x22, 0x39, 0x0a,
	0x07, 0x43, 0x4e, 0x49, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4e, 0x49, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4e, 0x49, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d,
	0x45, 0x53, 0x48, 0x4e, 0x45, 0x54, 0x10, 0x01, 0x22, 0xa4, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x43,
	0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x4c, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16,
	0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x4c, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x49, 0x58, 0x49, 0x41, 0x54, 0x47, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x4e, 0x54,
	0x52, 0x4f, 0x4c, 0x4c, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x52, 0x4c, 0x49,
	0x4e, 0x55, 0x58, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c,
	0x4c, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x45, 0x4f, 0x53, 0x4c, 0x41, 0x42,
	0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x4c, 0x45, 0x52,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x45, 0x4d, 0x4d, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x22,
	0x4c, 0x0a, 0x08, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x6c, 0x69, 0x6e, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x42, 0x0a,
	0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x2e, 0x56, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6b, 0x6e, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_event_proto_rawDescOnce sync.Once
	file_event_proto_rawDescData = file_event_proto_rawDesc
)

func file_event_proto_rawDescGZIP() []byte {
	file_event_proto_rawDescOnce.Do(func() {
		file_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_proto_rawDescData)
	})
	return file_event_proto_rawDescData
}

var file_event_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_event_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_event_proto_goTypes = []interface{}{
	(Cluster_ClusterType)(0),      // 0: event.Cluster.ClusterType
	(Cluster_IngressType)(0),      // 1: event.Cluster.IngressType
	(Cluster_CNIType)(0),          // 2: event.Cluster.CNIType
	(Cluster_ControllerType)(0),   // 3: event.Cluster.ControllerType
	(*KNEEvent)(nil),              // 4: event.KNEEvent
	(*DeployClusterStart)(nil),    // 5: event.DeployClusterStart
	(*DeployClusterEnd)(nil),      // 6: event.DeployClusterEnd
	(*CreateTopologyStart)(nil),   // 7: event.CreateTopologyStart
	(*CreateTopologyEnd)(nil),     // 8: event.CreateTopologyEnd
	(*Cluster)(nil),               // 9: event.Cluster
	(*Topology)(nil),              // 10: event.Topology
	(*Node)(nil),                  // 11: event.Node
	(*timestamppb.Timestamp)(nil), // 12: google.protobuf.Timestamp
	(topo.Vendor)(0),              // 13: topo.Vendor
}
var file_event_proto_depIdxs = []int32{
	12, // 0: event.KNEEvent.timestamp:type_name -> google.protobuf.Timestamp
	5,  // 1: event.KNEEvent.deploy_cluster_start:type_name -> event.DeployClusterStart
	6,  // 2: event.KNEEvent.deploy_cluster_end:type_name -> event.DeployClusterEnd
	7,  // 3: event.KNEEvent.create_topology_start:type_name -> event.CreateTopologyStart
	8,  // 4: event.KNEEvent.create_topology_end:type_name -> event.CreateTopologyEnd
	9,  // 5: event.DeployClusterStart.cluster:type_name -> event.Cluster
	10, // 6: event.CreateTopologyStart.topology:type_name -> event.Topology
	0,  // 7: event.Cluster.cluster:type_name -> event.Cluster.ClusterType
	1,  // 8: event.Cluster.ingress:type_name -> event.Cluster.IngressType
	2,  // 9: event.Cluster.cni:type_name -> event.Cluster.CNIType
	3,  // 10: event.Cluster.controllers:type_name -> event.Cluster.ControllerType
	11, // 11: event.Topology.nodes:type_name -> event.Node
	13, // 12: event.Node.vendor:type_name -> topo.Vendor
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_event_proto_init() }
func file_event_proto_init() {
	if File_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KNEEvent); i {
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
		file_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployClusterStart); i {
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
		file_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployClusterEnd); i {
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
		file_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTopologyStart); i {
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
		file_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTopologyEnd); i {
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
		file_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster); i {
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
		file_event_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Topology); i {
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
		file_event_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
	file_event_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*KNEEvent_DeployClusterStart)(nil),
		(*KNEEvent_DeployClusterEnd)(nil),
		(*KNEEvent_CreateTopologyStart)(nil),
		(*KNEEvent_CreateTopologyEnd)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_event_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_proto_goTypes,
		DependencyIndexes: file_event_proto_depIdxs,
		EnumInfos:         file_event_proto_enumTypes,
		MessageInfos:      file_event_proto_msgTypes,
	}.Build()
	File_event_proto = out.File
	file_event_proto_rawDesc = nil
	file_event_proto_goTypes = nil
	file_event_proto_depIdxs = nil
}
