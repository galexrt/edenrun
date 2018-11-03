// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/galexrt/edenconfmgmt/pkg/apis/nodes/v1alpha/api.proto

package v1alpha

import (
	fmt "fmt"
	v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha"
	v1alpha1 "github.com/galexrt/edenconfmgmt/pkg/apis/events/v1alpha"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Node struct {
	// Metadata for Node object.
	Metadata *v1alpha.ObjectMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Spec for Node.
	Spec                 *Spec    `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e58667dd6394e3, []int{0}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetMetadata() *v1alpha.ObjectMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Node) GetSpec() *Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type Spec struct {
	// FQDN of Node.
	Fqdn string `protobuf:"bytes,3,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	// IP addresses of Node.
	Addresses            []string `protobuf:"bytes,4,rep,name=addresses,proto3" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Spec) Reset()         { *m = Spec{} }
func (m *Spec) String() string { return proto.CompactTextString(m) }
func (*Spec) ProtoMessage()    {}
func (*Spec) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e58667dd6394e3, []int{1}
}

func (m *Spec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Spec.Unmarshal(m, b)
}
func (m *Spec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Spec.Marshal(b, m, deterministic)
}
func (m *Spec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Spec.Merge(m, src)
}
func (m *Spec) XXX_Size() int {
	return xxx_messageInfo_Spec.Size(m)
}
func (m *Spec) XXX_DiscardUnknown() {
	xxx_messageInfo_Spec.DiscardUnknown(m)
}

var xxx_messageInfo_Spec proto.InternalMessageInfo

func (m *Spec) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *Spec) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

// Get
type GetRequest struct {
	GetOptions           *v1alpha.GetOptions `protobuf:"bytes,1,opt,name=getOptions,proto3" json:"getOptions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e58667dd6394e3, []int{2}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetGetOptions() *v1alpha.GetOptions {
	if m != nil {
		return m.GetOptions
	}
	return nil
}

type GetResponse struct {
	// Node object.
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e58667dd6394e3, []int{3}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

// List
type ListRequest struct {
	ListOptions          *v1alpha.ListOptions `protobuf:"bytes,1,opt,name=listOptions,proto3" json:"listOptions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e58667dd6394e3, []int{4}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetListOptions() *v1alpha.ListOptions {
	if m != nil {
		return m.ListOptions
	}
	return nil
}

type ListResponse struct {
	// Node list.
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e58667dd6394e3, []int{5}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

// Add
type AddRequest struct {
	// Node object.
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e58667dd6394e3, []int{6}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

type AddResponse struct {
	// Node object.
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()    {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e58667dd6394e3, []int{7}
}

func (m *AddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddResponse.Unmarshal(m, b)
}
func (m *AddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddResponse.Marshal(b, m, deterministic)
}
func (m *AddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResponse.Merge(m, src)
}
func (m *AddResponse) XXX_Size() int {
	return xxx_messageInfo_AddResponse.Size(m)
}
func (m *AddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddResponse proto.InternalMessageInfo

func (m *AddResponse) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

// Update
type UpdateRequest struct {
	// Node object.
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e58667dd6394e3, []int{8}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

type UpdateResponse struct {
	// Node object.
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e58667dd6394e3, []int{9}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

// Delete
type DeleteRequest struct {
	// Node object.
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e58667dd6394e3, []int{10}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

type DeleteResponse struct {
	// Node object.
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e58667dd6394e3, []int{11}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

// Watch
type WatchRequest struct {
	WatchOptions         *v1alpha.WatchOptions `protobuf:"bytes,1,opt,name=watchOptions,proto3" json:"watchOptions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *WatchRequest) Reset()         { *m = WatchRequest{} }
func (m *WatchRequest) String() string { return proto.CompactTextString(m) }
func (*WatchRequest) ProtoMessage()    {}
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e58667dd6394e3, []int{12}
}

func (m *WatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchRequest.Unmarshal(m, b)
}
func (m *WatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchRequest.Marshal(b, m, deterministic)
}
func (m *WatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchRequest.Merge(m, src)
}
func (m *WatchRequest) XXX_Size() int {
	return xxx_messageInfo_WatchRequest.Size(m)
}
func (m *WatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WatchRequest proto.InternalMessageInfo

func (m *WatchRequest) GetWatchOptions() *v1alpha.WatchOptions {
	if m != nil {
		return m.WatchOptions
	}
	return nil
}

type WatchResponse struct {
	// Event info for watch response.
	Event *v1alpha1.Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	// Node for watch response.
	Node                 *Node    `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WatchResponse) Reset()         { *m = WatchResponse{} }
func (m *WatchResponse) String() string { return proto.CompactTextString(m) }
func (*WatchResponse) ProtoMessage()    {}
func (*WatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e58667dd6394e3, []int{13}
}

func (m *WatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchResponse.Unmarshal(m, b)
}
func (m *WatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchResponse.Marshal(b, m, deterministic)
}
func (m *WatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchResponse.Merge(m, src)
}
func (m *WatchResponse) XXX_Size() int {
	return xxx_messageInfo_WatchResponse.Size(m)
}
func (m *WatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WatchResponse proto.InternalMessageInfo

func (m *WatchResponse) GetEvent() *v1alpha1.Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *WatchResponse) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func init() {
	proto.RegisterType((*Node)(nil), "nodes.v1alpha.Node")
	proto.RegisterType((*Spec)(nil), "nodes.v1alpha.Spec")
	proto.RegisterType((*GetRequest)(nil), "nodes.v1alpha.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "nodes.v1alpha.GetResponse")
	proto.RegisterType((*ListRequest)(nil), "nodes.v1alpha.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "nodes.v1alpha.ListResponse")
	proto.RegisterType((*AddRequest)(nil), "nodes.v1alpha.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "nodes.v1alpha.AddResponse")
	proto.RegisterType((*UpdateRequest)(nil), "nodes.v1alpha.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "nodes.v1alpha.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "nodes.v1alpha.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "nodes.v1alpha.DeleteResponse")
	proto.RegisterType((*WatchRequest)(nil), "nodes.v1alpha.WatchRequest")
	proto.RegisterType((*WatchResponse)(nil), "nodes.v1alpha.WatchResponse")
}

func init() {
	proto.RegisterFile("github.com/galexrt/edenconfmgmt/pkg/apis/nodes/v1alpha/api.proto", fileDescriptor_72e58667dd6394e3)
}

var fileDescriptor_72e58667dd6394e3 = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0xad, 0x6b, 0xb7, 0xfa, 0x3a, 0x49, 0xbf, 0xc3, 0x22, 0x24, 0xd7, 0x04, 0x2b, 0xf2, 0x85,
	0x20, 0x84, 0x0d, 0x45, 0x44, 0xa9, 0xa8, 0x28, 0x41, 0x2d, 0x91, 0x50, 0x69, 0x25, 0x17, 0x54,
	0x89, 0x9b, 0xe3, 0x9d, 0x38, 0x86, 0xc4, 0x76, 0xb3, 0x9b, 0xc0, 0x91, 0x9f, 0xc0, 0x3f, 0xe0,
	0xef, 0xf4, 0xc8, 0x91, 0x23, 0x0d, 0x7f, 0x04, 0x79, 0xed, 0x24, 0xb6, 0x93, 0x8a, 0x24, 0x37,
	0x8f, 0xe7, 0xcd, 0x7b, 0x6f, 0xbc, 0x6f, 0x13, 0x78, 0xe5, 0xf9, 0xbc, 0x3b, 0x6c, 0x9b, 0x6e,
	0xd8, 0xb7, 0x3c, 0xa7, 0x87, 0x5f, 0x07, 0xdc, 0x42, 0x8a, 0x81, 0x1b, 0x06, 0x9d, 0xbe, 0xd7,
	0xe7, 0x56, 0xf4, 0xd9, 0xb3, 0x9c, 0xc8, 0x67, 0x56, 0x10, 0x52, 0x64, 0xd6, 0xe8, 0xa9, 0xd3,
	0x8b, 0xba, 0x4e, 0xfc, 0xca, 0x8c, 0x06, 0x21, 0x0f, 0xc9, 0xae, 0x68, 0x98, 0x69, 0x43, 0x7b,
	0x9c, 0x25, 0x0c, 0xbd, 0xd0, 0x12, 0xa8, 0xf6, 0xb0, 0x23, 0x2a, 0x51, 0x88, 0xa7, 0x64, 0x5a,
	0x3b, 0x5a, 0x5a, 0xdf, 0x0d, 0x07, 0x38, 0x2f, 0xaf, 0x35, 0x97, 0x26, 0xc0, 0x11, 0x06, 0x7c,
	0xc1, 0x06, 0x86, 0x0f, 0xca, 0x59, 0x48, 0x91, 0x34, 0xe0, 0xbf, 0x3e, 0x72, 0x87, 0x3a, 0xdc,
	0x51, 0xa5, 0xaa, 0x54, 0x2b, 0xed, 0x57, 0xcc, 0x58, 0x75, 0xb2, 0x9b, 0x79, 0xde, 0xfe, 0x84,
	0x2e, 0x7f, 0x97, 0x62, 0xec, 0x29, 0x9a, 0x3c, 0x00, 0x85, 0x45, 0xe8, 0xaa, 0x9b, 0x62, 0xea,
	0x8e, 0x99, 0xfb, 0x24, 0xe6, 0x45, 0x84, 0xae, 0x2d, 0x00, 0x46, 0x03, 0x94, 0xb8, 0x22, 0x04,
	0x94, 0xce, 0x15, 0x0d, 0x54, 0xb9, 0x2a, 0xd5, 0x76, 0x6c, 0xf1, 0x4c, 0x2a, 0xb0, 0xe3, 0x50,
	0x3a, 0x40, 0xc6, 0x90, 0xa9, 0x4a, 0x55, 0xae, 0xed, 0xd8, 0xb3, 0x17, 0xc6, 0x1b, 0x80, 0x16,
	0x72, 0x1b, 0xaf, 0x86, 0xc8, 0x38, 0x69, 0x00, 0x78, 0xc8, 0xcf, 0x23, 0xee, 0x87, 0x01, 0x4b,
	0xcd, 0xaa, 0x79, 0xb3, 0xad, 0x69, 0xdf, 0xce, 0x60, 0x8d, 0x3a, 0x94, 0x04, 0x0f, 0x8b, 0xc2,
	0x80, 0x61, 0xec, 0x3c, 0x36, 0x9b, 0x52, 0x14, 0x9d, 0xc7, 0x9f, 0xc5, 0x16, 0x00, 0xe3, 0x2d,
	0x94, 0x4e, 0x7d, 0x36, 0x35, 0xf0, 0x02, 0x4a, 0x3d, 0x9f, 0x15, 0x1c, 0xec, 0xe5, 0x1d, 0x9c,
	0xce, 0x00, 0x76, 0x16, 0x6d, 0x1c, 0x40, 0x39, 0xe1, 0x4a, 0x4d, 0x3c, 0x84, 0x2d, 0xa1, 0xab,
	0x4a, 0x55, 0xf9, 0x36, 0x17, 0x09, 0xc2, 0x78, 0x0e, 0xd0, 0xa4, 0x74, 0xe2, 0x62, 0x69, 0xf7,
	0x75, 0x28, 0x89, 0xb1, 0x55, 0xb7, 0x6e, 0xc0, 0xee, 0x87, 0x88, 0x3a, 0x1c, 0x57, 0x56, 0x3c,
	0x80, 0xff, 0x27, 0x93, 0x6b, 0x88, 0x1e, 0x63, 0x0f, 0xd7, 0x13, 0x9d, 0x4c, 0xae, 0x2a, 0x7a,
	0x06, 0xe5, 0x4b, 0x87, 0xbb, 0xdd, 0x89, 0xe6, 0x4b, 0x28, 0x7f, 0x89, 0xeb, 0xfc, 0x09, 0x6b,
	0xf9, 0x13, 0xbe, 0xcc, 0x20, 0xec, 0x1c, 0xde, 0x40, 0xd8, 0x4d, 0xf9, 0x52, 0x27, 0x8f, 0x60,
	0x4b, 0xdc, 0xc0, 0x94, 0xe9, 0xae, 0x99, 0xdc, 0xc7, 0x29, 0xd7, 0x49, 0x5c, 0xda, 0x09, 0x66,
	0x6a, 0x7b, 0xf3, 0x1f, 0xb6, 0xf7, 0x7f, 0xc8, 0x50, 0x8e, 0x4b, 0x76, 0x81, 0x83, 0x91, 0xef,
	0x22, 0x39, 0x04, 0xb9, 0x85, 0x9c, 0xec, 0x15, 0x46, 0x66, 0x77, 0x47, 0xd3, 0x16, 0xb5, 0x52,
	0x93, 0x47, 0xa0, 0xc4, 0xc9, 0x24, 0x45, 0x4c, 0x26, 0xfa, 0xda, 0xbd, 0x85, 0xbd, 0x94, 0xe0,
	0x10, 0xe4, 0x26, 0xa5, 0x73, 0xf2, 0xb3, 0xcc, 0xce, 0xc9, 0x67, 0x73, 0x79, 0x02, 0xdb, 0x49,
	0x68, 0x48, 0xa5, 0x80, 0xca, 0xa5, 0x50, 0xbb, 0x7f, 0x4b, 0x77, 0x46, 0x93, 0xc4, 0x60, 0x8e,
	0x26, 0x97, 0xab, 0x39, 0x9a, 0x42, 0x76, 0x8e, 0x61, 0x4b, 0x1c, 0x21, 0x29, 0x6e, 0x9c, 0x0d,
	0x8a, 0x56, 0x59, 0xdc, 0x4c, 0x38, 0x9e, 0x48, 0xaf, 0xdf, 0x5f, 0xdf, 0xe8, 0xd2, 0xaf, 0x1b,
	0x7d, 0xe3, 0xdb, 0x58, 0x97, 0xae, 0xc7, 0xba, 0xf4, 0x73, 0xac, 0x4b, 0xbf, 0xc7, 0xba, 0xf4,
	0xfd, 0x8f, 0xbe, 0xf1, 0xb1, 0xbe, 0xde, 0xff, 0x4f, 0x7b, 0x5b, 0xfc, 0x74, 0x3f, 0xfb, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0xe6, 0x37, 0x9f, 0xc5, 0xc0, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodesServiceClient is the client API for NodesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodesServiceClient interface {
	// Get a Node.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// List Nodes.
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// Add a Node.
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	// Update a Node.
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Delete a Node.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// Watch Node.
	Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (NodesService_WatchClient, error)
}

type nodesServiceClient struct {
	cc *grpc.ClientConn
}

func NewNodesServiceClient(cc *grpc.ClientConn) NodesServiceClient {
	return &nodesServiceClient{cc}
}

func (c *nodesServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/nodes.v1alpha.NodesService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodesServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/nodes.v1alpha.NodesService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodesServiceClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/nodes.v1alpha.NodesService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodesServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/nodes.v1alpha.NodesService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodesServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/nodes.v1alpha.NodesService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodesServiceClient) Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (NodesService_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NodesService_serviceDesc.Streams[0], "/nodes.v1alpha.NodesService/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodesServiceWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodesService_WatchClient interface {
	Recv() (*WatchResponse, error)
	grpc.ClientStream
}

type nodesServiceWatchClient struct {
	grpc.ClientStream
}

func (x *nodesServiceWatchClient) Recv() (*WatchResponse, error) {
	m := new(WatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NodesServiceServer is the server API for NodesService service.
type NodesServiceServer interface {
	// Get a Node.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// List Nodes.
	List(context.Context, *ListRequest) (*ListResponse, error)
	// Add a Node.
	Add(context.Context, *AddRequest) (*AddResponse, error)
	// Update a Node.
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Delete a Node.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// Watch Node.
	Watch(*WatchRequest, NodesService_WatchServer) error
}

func RegisterNodesServiceServer(s *grpc.Server, srv NodesServiceServer) {
	s.RegisterService(&_NodesService_serviceDesc, srv)
}

func _NodesService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodesServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodes.v1alpha.NodesService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodesServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodesService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodesServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodes.v1alpha.NodesService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodesServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodesService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodesServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodes.v1alpha.NodesService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodesServiceServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodesService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodesServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodes.v1alpha.NodesService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodesServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodesService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodesServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodes.v1alpha.NodesService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodesServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodesService_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodesServiceServer).Watch(m, &nodesServiceWatchServer{stream})
}

type NodesService_WatchServer interface {
	Send(*WatchResponse) error
	grpc.ServerStream
}

type nodesServiceWatchServer struct {
	grpc.ServerStream
}

func (x *nodesServiceWatchServer) Send(m *WatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _NodesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nodes.v1alpha.NodesService",
	HandlerType: (*NodesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _NodesService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _NodesService_List_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _NodesService_Add_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _NodesService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NodesService_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _NodesService_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/galexrt/edenconfmgmt/pkg/apis/nodes/v1alpha/api.proto",
}
