// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/galexrt/edenconfmgmt/pkg/apis/nodes/v1alpha/api.proto

package v1alpha

import (
	fmt "fmt"
	v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/events/v1alpha"
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

type VersionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionRequest) Reset()         { *m = VersionRequest{} }
func (m *VersionRequest) String() string { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()    {}
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e58667dd6394e3, []int{0}
}

func (m *VersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionRequest.Unmarshal(m, b)
}
func (m *VersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionRequest.Marshal(b, m, deterministic)
}
func (m *VersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionRequest.Merge(m, src)
}
func (m *VersionRequest) XXX_Size() int {
	return xxx_messageInfo_VersionRequest.Size(m)
}
func (m *VersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VersionRequest proto.InternalMessageInfo

type VersionResponse struct {
	// Semversion compatible version number.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Semversion major number.
	Major int64 `protobuf:"varint,2,opt,name=major,proto3" json:"major,omitempty"`
	// Semversion minor number.
	Minor int64 `protobuf:"varint,3,opt,name=minor,proto3" json:"minor,omitempty"`
	// Semversion patch number.
	Patch int64 `protobuf:"varint,4,opt,name=patch,proto3" json:"patch,omitempty"`
	// API state (e.g., alpha, beta, stable).
	State                string   `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e58667dd6394e3, []int{1}
}

func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (m *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(m, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *VersionResponse) GetMajor() int64 {
	if m != nil {
		return m.Major
	}
	return 0
}

func (m *VersionResponse) GetMinor() int64 {
	if m != nil {
		return m.Minor
	}
	return 0
}

func (m *VersionResponse) GetPatch() int64 {
	if m != nil {
		return m.Patch
	}
	return 0
}

func (m *VersionResponse) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type Node struct {
	// ID of Node.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of Node.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// FQDN of Node.
	Fqdn string `protobuf:"bytes,3,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	// IP addresses of Node.
	Addresses []string `protobuf:"bytes,4,rep,name=addresses,proto3" json:"addresses,omitempty"`
	// Labels of Node.
	Labels map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Annotations of Node.
	Annotations          map[string]string `protobuf:"bytes,6,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e58667dd6394e3, []int{2}
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

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Node) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *Node) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *Node) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Node) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

type AddNodeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddNodeResponse) Reset()         { *m = AddNodeResponse{} }
func (m *AddNodeResponse) String() string { return proto.CompactTextString(m) }
func (*AddNodeResponse) ProtoMessage()    {}
func (*AddNodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e58667dd6394e3, []int{3}
}

func (m *AddNodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddNodeResponse.Unmarshal(m, b)
}
func (m *AddNodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddNodeResponse.Marshal(b, m, deterministic)
}
func (m *AddNodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddNodeResponse.Merge(m, src)
}
func (m *AddNodeResponse) XXX_Size() int {
	return xxx_messageInfo_AddNodeResponse.Size(m)
}
func (m *AddNodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddNodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddNodeResponse proto.InternalMessageInfo

type UpdateNodeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateNodeResponse) Reset()         { *m = UpdateNodeResponse{} }
func (m *UpdateNodeResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateNodeResponse) ProtoMessage()    {}
func (*UpdateNodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e58667dd6394e3, []int{4}
}

func (m *UpdateNodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateNodeResponse.Unmarshal(m, b)
}
func (m *UpdateNodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateNodeResponse.Marshal(b, m, deterministic)
}
func (m *UpdateNodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateNodeResponse.Merge(m, src)
}
func (m *UpdateNodeResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateNodeResponse.Size(m)
}
func (m *UpdateNodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateNodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateNodeResponse proto.InternalMessageInfo

type DeleteNodeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteNodeResponse) Reset()         { *m = DeleteNodeResponse{} }
func (m *DeleteNodeResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteNodeResponse) ProtoMessage()    {}
func (*DeleteNodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e58667dd6394e3, []int{5}
}

func (m *DeleteNodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteNodeResponse.Unmarshal(m, b)
}
func (m *DeleteNodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteNodeResponse.Marshal(b, m, deterministic)
}
func (m *DeleteNodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteNodeResponse.Merge(m, src)
}
func (m *DeleteNodeResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteNodeResponse.Size(m)
}
func (m *DeleteNodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteNodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteNodeResponse proto.InternalMessageInfo

type NodeWatchResponse struct {
	// Event info for watch response.
	Event *v1alpha.Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	// Node for watch response.
	Node                 *Node    `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeWatchResponse) Reset()         { *m = NodeWatchResponse{} }
func (m *NodeWatchResponse) String() string { return proto.CompactTextString(m) }
func (*NodeWatchResponse) ProtoMessage()    {}
func (*NodeWatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e58667dd6394e3, []int{6}
}

func (m *NodeWatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeWatchResponse.Unmarshal(m, b)
}
func (m *NodeWatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeWatchResponse.Marshal(b, m, deterministic)
}
func (m *NodeWatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeWatchResponse.Merge(m, src)
}
func (m *NodeWatchResponse) XXX_Size() int {
	return xxx_messageInfo_NodeWatchResponse.Size(m)
}
func (m *NodeWatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeWatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NodeWatchResponse proto.InternalMessageInfo

func (m *NodeWatchResponse) GetEvent() *v1alpha.Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *NodeWatchResponse) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func init() {
	proto.RegisterType((*VersionRequest)(nil), "nodes.v1alpha.VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "nodes.v1alpha.VersionResponse")
	proto.RegisterType((*Node)(nil), "nodes.v1alpha.Node")
	proto.RegisterMapType((map[string]string)(nil), "nodes.v1alpha.Node.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "nodes.v1alpha.Node.LabelsEntry")
	proto.RegisterType((*AddNodeResponse)(nil), "nodes.v1alpha.AddNodeResponse")
	proto.RegisterType((*UpdateNodeResponse)(nil), "nodes.v1alpha.UpdateNodeResponse")
	proto.RegisterType((*DeleteNodeResponse)(nil), "nodes.v1alpha.DeleteNodeResponse")
	proto.RegisterType((*NodeWatchResponse)(nil), "nodes.v1alpha.NodeWatchResponse")
}

func init() {
	proto.RegisterFile("github.com/galexrt/edenconfmgmt/pkg/apis/nodes/v1alpha/api.proto", fileDescriptor_72e58667dd6394e3)
}

var fileDescriptor_72e58667dd6394e3 = []byte{
	// 584 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0xe3, 0xa4, 0x91, 0x27, 0xd0, 0x9f, 0xa5, 0x48, 0x96, 0x05, 0x26, 0x44, 0x20, 0x22,
	0x21, 0x6c, 0x48, 0x0f, 0xfc, 0x1c, 0x2a, 0x8a, 0x5a, 0xe0, 0x80, 0x38, 0x18, 0x01, 0x12, 0xb7,
	0x4d, 0x3c, 0x49, 0x4c, 0x9d, 0x5d, 0xd7, 0xbb, 0x89, 0xe8, 0x0d, 0x89, 0x17, 0xe0, 0x0d, 0x78,
	0x9d, 0x1e, 0x39, 0xf6, 0x48, 0xc3, 0x8b, 0xa0, 0x5d, 0x3b, 0x7f, 0xae, 0x2b, 0x95, 0xdb, 0x7c,
	0xdf, 0xce, 0x7c, 0x33, 0x9e, 0x6f, 0x64, 0x78, 0x39, 0x88, 0xe4, 0x70, 0xdc, 0xf5, 0x7a, 0x7c,
	0xe4, 0x0f, 0x68, 0x8c, 0xdf, 0x52, 0xe9, 0x63, 0x88, 0xac, 0xc7, 0x59, 0x7f, 0x34, 0x18, 0x49,
	0x3f, 0x39, 0x1a, 0xf8, 0x34, 0x89, 0x84, 0xcf, 0x78, 0x88, 0xc2, 0x9f, 0x3c, 0xa1, 0x71, 0x32,
	0xa4, 0x8a, 0xf2, 0x92, 0x94, 0x4b, 0x4e, 0xae, 0xeb, 0x07, 0x2f, 0x7f, 0x70, 0x1e, 0x2d, 0x0b,
	0xf2, 0x01, 0xf7, 0x75, 0x56, 0x77, 0xdc, 0xd7, 0x48, 0x03, 0x1d, 0x65, 0xd5, 0xce, 0xfe, 0x95,
	0xfb, 0xe3, 0x04, 0x99, 0x2c, 0x19, 0xa0, 0xb5, 0x05, 0x1b, 0x9f, 0x30, 0x15, 0x11, 0x67, 0x01,
	0x1e, 0x8f, 0x51, 0xc8, 0xd6, 0x0f, 0x03, 0x36, 0xe7, 0x94, 0x48, 0x38, 0x13, 0x48, 0x6c, 0xa8,
	0x4f, 0x32, 0xca, 0x36, 0x9a, 0x46, 0xdb, 0x0a, 0x66, 0x90, 0xec, 0x40, 0x6d, 0x44, 0xbf, 0xf2,
	0xd4, 0xae, 0x34, 0x8d, 0xb6, 0x19, 0x64, 0x40, 0xb3, 0x11, 0xe3, 0xa9, 0x6d, 0xe6, 0xac, 0x02,
	0x8a, 0x4d, 0xa8, 0xec, 0x0d, 0xed, 0x6a, 0xc6, 0x6a, 0xa0, 0x58, 0x21, 0xa9, 0x44, 0xbb, 0xa6,
	0x95, 0x33, 0xd0, 0x3a, 0xab, 0x40, 0xf5, 0x3d, 0x0f, 0x91, 0x6c, 0x40, 0x25, 0x0a, 0xf3, 0xae,
	0x95, 0x28, 0x24, 0x04, 0xaa, 0x8c, 0x8e, 0x50, 0xf7, 0xb3, 0x02, 0x1d, 0x2b, 0xae, 0x7f, 0x1c,
	0x32, 0xdd, 0xcd, 0x0a, 0x74, 0x4c, 0x6e, 0x81, 0x45, 0xc3, 0x30, 0x45, 0x21, 0x50, 0xd8, 0xd5,
	0xa6, 0xd9, 0xb6, 0x82, 0x05, 0x41, 0x9e, 0xc2, 0x7a, 0x4c, 0xbb, 0x18, 0x0b, 0xbb, 0xd6, 0x34,
	0xdb, 0x8d, 0xce, 0x1d, 0x6f, 0xc5, 0x08, 0x4f, 0xb5, 0xf6, 0xde, 0xe9, 0x8c, 0x43, 0x26, 0xd3,
	0x93, 0x20, 0x4f, 0x27, 0xaf, 0xa1, 0x41, 0x19, 0xe3, 0x92, 0xca, 0x88, 0x33, 0x61, 0xaf, 0xeb,
	0xea, 0x7b, 0x65, 0xd5, 0xfb, 0x8b, 0xb4, 0x4c, 0x62, 0xb9, 0xd0, 0x79, 0x0e, 0x8d, 0x25, 0x79,
	0xb2, 0x05, 0xe6, 0x11, 0x9e, 0xe4, 0x9f, 0xa9, 0x42, 0xb5, 0x96, 0x09, 0x8d, 0xc7, 0xb3, 0x0f,
	0xcd, 0xc0, 0x8b, 0xca, 0x33, 0xc3, 0xd9, 0x83, 0xad, 0xa2, 0xf6, 0xff, 0xd4, 0xb7, 0xb6, 0x61,
	0x73, 0x3f, 0x0c, 0xd5, 0x8c, 0x33, 0x7f, 0x5b, 0x3b, 0x40, 0x3e, 0x26, 0x21, 0x95, 0x58, 0x64,
	0x0f, 0x30, 0xc6, 0x02, 0x1b, 0xc1, 0xb6, 0xc2, 0x9f, 0x95, 0x79, 0xf3, 0x03, 0x79, 0x08, 0x35,
	0x7d, 0x62, 0x7a, 0x82, 0x46, 0xe7, 0xa6, 0x97, 0x1d, 0xdc, 0x7c, 0x23, 0x87, 0x0a, 0x06, 0x59,
	0x0e, 0x79, 0x00, 0x55, 0xb5, 0x2f, 0x3d, 0x59, 0xa3, 0x73, 0xa3, 0x64, 0x79, 0x81, 0x4e, 0xe8,
	0xfc, 0x32, 0xe1, 0x9a, 0x82, 0xe2, 0x03, 0xa6, 0x93, 0xa8, 0x87, 0xe4, 0x2d, 0xd4, 0xf3, 0xd3,
	0x24, 0xb7, 0x0b, 0x65, 0xab, 0x57, 0xec, 0xb8, 0x97, 0x3d, 0xe7, 0x03, 0xef, 0x42, 0xfd, 0x0d,
	0x4a, 0x7d, 0x61, 0x65, 0x03, 0x38, 0x65, 0x24, 0xd9, 0x83, 0x7a, 0xbe, 0xb9, 0xf2, 0xa2, 0x62,
	0xd3, 0xc2, 0x9a, 0xc9, 0x01, 0xc0, 0x62, 0xcd, 0xe5, 0x12, 0x77, 0x0b, 0xe4, 0x45, 0x5b, 0x94,
	0xca, 0xc2, 0x96, 0xab, 0xa9, 0x5c, 0xb4, 0x91, 0x1c, 0x80, 0xa5, 0x2d, 0xbc, 0x5c, 0xa4, 0x59,
	0x42, 0xae, 0xb8, 0xfe, 0xd8, 0x78, 0x75, 0xff, 0xf4, 0xdc, 0x35, 0xce, 0xce, 0xdd, 0xb5, 0xef,
	0x53, 0xd7, 0x38, 0x9d, 0xba, 0xc6, 0xef, 0xa9, 0x6b, 0xfc, 0x99, 0xba, 0xc6, 0xcf, 0xbf, 0xee,
	0xda, 0x97, 0x7a, 0x5e, 0xda, 0x5d, 0xd7, 0x3f, 0x9b, 0xdd, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x0b, 0xd7, 0xbf, 0xd4, 0x31, 0x05, 0x00, 0x00,
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
	// Version returns the API version.
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
	// Get a Node.
	GetNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Node, error)
	// Add a Node.
	AddNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*AddNodeResponse, error)
	// Update a Node.
	UpdateNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*UpdateNodeResponse, error)
	// Delete a Node.
	DeleteNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*DeleteNodeResponse, error)
	// Watch a Node.
	WatchNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (NodesService_WatchNodeClient, error)
}

type nodesServiceClient struct {
	cc *grpc.ClientConn
}

func NewNodesServiceClient(cc *grpc.ClientConn) NodesServiceClient {
	return &nodesServiceClient{cc}
}

func (c *nodesServiceClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/nodes.v1alpha.NodesService/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodesServiceClient) GetNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/nodes.v1alpha.NodesService/GetNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodesServiceClient) AddNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*AddNodeResponse, error) {
	out := new(AddNodeResponse)
	err := c.cc.Invoke(ctx, "/nodes.v1alpha.NodesService/AddNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodesServiceClient) UpdateNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*UpdateNodeResponse, error) {
	out := new(UpdateNodeResponse)
	err := c.cc.Invoke(ctx, "/nodes.v1alpha.NodesService/UpdateNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodesServiceClient) DeleteNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*DeleteNodeResponse, error) {
	out := new(DeleteNodeResponse)
	err := c.cc.Invoke(ctx, "/nodes.v1alpha.NodesService/DeleteNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodesServiceClient) WatchNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (NodesService_WatchNodeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NodesService_serviceDesc.Streams[0], "/nodes.v1alpha.NodesService/WatchNode", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodesServiceWatchNodeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodesService_WatchNodeClient interface {
	Recv() (*NodeWatchResponse, error)
	grpc.ClientStream
}

type nodesServiceWatchNodeClient struct {
	grpc.ClientStream
}

func (x *nodesServiceWatchNodeClient) Recv() (*NodeWatchResponse, error) {
	m := new(NodeWatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NodesServiceServer is the server API for NodesService service.
type NodesServiceServer interface {
	// Version returns the API version.
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
	// Get a Node.
	GetNode(context.Context, *Node) (*Node, error)
	// Add a Node.
	AddNode(context.Context, *Node) (*AddNodeResponse, error)
	// Update a Node.
	UpdateNode(context.Context, *Node) (*UpdateNodeResponse, error)
	// Delete a Node.
	DeleteNode(context.Context, *Node) (*DeleteNodeResponse, error)
	// Watch a Node.
	WatchNode(*Node, NodesService_WatchNodeServer) error
}

func RegisterNodesServiceServer(s *grpc.Server, srv NodesServiceServer) {
	s.RegisterService(&_NodesService_serviceDesc, srv)
}

func _NodesService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodesServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodes.v1alpha.NodesService/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodesServiceServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodesService_GetNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodesServiceServer).GetNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodes.v1alpha.NodesService/GetNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodesServiceServer).GetNode(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodesService_AddNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodesServiceServer).AddNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodes.v1alpha.NodesService/AddNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodesServiceServer).AddNode(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodesService_UpdateNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodesServiceServer).UpdateNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodes.v1alpha.NodesService/UpdateNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodesServiceServer).UpdateNode(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodesService_DeleteNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodesServiceServer).DeleteNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodes.v1alpha.NodesService/DeleteNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodesServiceServer).DeleteNode(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodesService_WatchNode_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Node)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodesServiceServer).WatchNode(m, &nodesServiceWatchNodeServer{stream})
}

type NodesService_WatchNodeServer interface {
	Send(*NodeWatchResponse) error
	grpc.ServerStream
}

type nodesServiceWatchNodeServer struct {
	grpc.ServerStream
}

func (x *nodesServiceWatchNodeServer) Send(m *NodeWatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _NodesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nodes.v1alpha.NodesService",
	HandlerType: (*NodesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _NodesService_Version_Handler,
		},
		{
			MethodName: "GetNode",
			Handler:    _NodesService_GetNode_Handler,
		},
		{
			MethodName: "AddNode",
			Handler:    _NodesService_AddNode_Handler,
		},
		{
			MethodName: "UpdateNode",
			Handler:    _NodesService_UpdateNode_Handler,
		},
		{
			MethodName: "DeleteNode",
			Handler:    _NodesService_DeleteNode_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchNode",
			Handler:       _NodesService_WatchNode_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/galexrt/edenconfmgmt/pkg/apis/nodes/v1alpha/api.proto",
}
