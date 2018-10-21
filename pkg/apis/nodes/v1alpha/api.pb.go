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
	ObjectMetadata *v1alpha.ObjectMetadata `protobuf:"bytes,1,opt,name=objectMetadata,proto3" json:"objectMetadata,omitempty"`
	// Spec for Node.
	Spec                 *NodeSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
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

func (m *Node) GetObjectMetadata() *v1alpha.ObjectMetadata {
	if m != nil {
		return m.ObjectMetadata
	}
	return nil
}

func (m *Node) GetSpec() *NodeSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type NodeSpec struct {
	// FQDN of Node.
	Fqdn string `protobuf:"bytes,3,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	// IP addresses of Node.
	Addresses            []string `protobuf:"bytes,4,rep,name=addresses,proto3" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeSpec) Reset()         { *m = NodeSpec{} }
func (m *NodeSpec) String() string { return proto.CompactTextString(m) }
func (*NodeSpec) ProtoMessage()    {}
func (*NodeSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e58667dd6394e3, []int{1}
}

func (m *NodeSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeSpec.Unmarshal(m, b)
}
func (m *NodeSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeSpec.Marshal(b, m, deterministic)
}
func (m *NodeSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeSpec.Merge(m, src)
}
func (m *NodeSpec) XXX_Size() int {
	return xxx_messageInfo_NodeSpec.Size(m)
}
func (m *NodeSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeSpec.DiscardUnknown(m)
}

var xxx_messageInfo_NodeSpec proto.InternalMessageInfo

func (m *NodeSpec) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *NodeSpec) GetAddresses() []string {
	if m != nil {
		return m.Addresses
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
	return fileDescriptor_72e58667dd6394e3, []int{2}
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
	return fileDescriptor_72e58667dd6394e3, []int{3}
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
	return fileDescriptor_72e58667dd6394e3, []int{4}
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
	Event *v1alpha1.Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
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
	return fileDescriptor_72e58667dd6394e3, []int{5}
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

func (m *NodeWatchResponse) GetEvent() *v1alpha1.Event {
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
	proto.RegisterType((*Node)(nil), "nodes.v1alpha.Node")
	proto.RegisterType((*NodeSpec)(nil), "nodes.v1alpha.NodeSpec")
	proto.RegisterType((*AddNodeResponse)(nil), "nodes.v1alpha.AddNodeResponse")
	proto.RegisterType((*UpdateNodeResponse)(nil), "nodes.v1alpha.UpdateNodeResponse")
	proto.RegisterType((*DeleteNodeResponse)(nil), "nodes.v1alpha.DeleteNodeResponse")
	proto.RegisterType((*NodeWatchResponse)(nil), "nodes.v1alpha.NodeWatchResponse")
}

func init() {
	proto.RegisterFile("github.com/galexrt/edenconfmgmt/pkg/apis/nodes/v1alpha/api.proto", fileDescriptor_72e58667dd6394e3)
}

var fileDescriptor_72e58667dd6394e3 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xed, 0x12, 0x43, 0xc8, 0xf0, 0xa5, 0x2e, 0x20, 0x22, 0xab, 0x58, 0x21, 0x12, 0x22, 0x52,
	0x85, 0x0d, 0xed, 0x15, 0x01, 0x45, 0x01, 0x4e, 0x80, 0xe4, 0x0a, 0x90, 0xb8, 0x6d, 0x76, 0x27,
	0x8e, 0x21, 0xf1, 0x6e, 0xbd, 0x9b, 0x08, 0x6e, 0xfc, 0x04, 0xfe, 0x00, 0xff, 0xa7, 0x47, 0x8e,
	0x1c, 0xa9, 0xf9, 0x23, 0x68, 0x37, 0x0e, 0x69, 0x1c, 0x47, 0xca, 0x6d, 0xe7, 0xcd, 0xcc, 0x1b,
	0xbf, 0xf7, 0x64, 0x78, 0x9e, 0xa4, 0x66, 0x34, 0x1d, 0x84, 0x5c, 0x4e, 0xa2, 0x84, 0x8d, 0xf1,
	0x6b, 0x6e, 0x22, 0x14, 0x98, 0x71, 0x99, 0x0d, 0x27, 0xc9, 0xc4, 0x44, 0xea, 0x4b, 0x12, 0x31,
	0x95, 0xea, 0x28, 0x93, 0x02, 0x75, 0x34, 0x7b, 0xcc, 0xc6, 0x6a, 0xc4, 0x2c, 0x14, 0xaa, 0x5c,
	0x1a, 0x49, 0xaf, 0xb9, 0x46, 0x58, 0x36, 0xfc, 0x87, 0xe7, 0x09, 0x65, 0x22, 0x23, 0x37, 0x35,
	0x98, 0x0e, 0x5d, 0xe5, 0x0a, 0xf7, 0x9a, 0x6f, 0xfb, 0xcf, 0xb6, 0xbe, 0xcf, 0x65, 0x8e, 0xeb,
	0xe7, 0xfd, 0xa3, 0xad, 0x09, 0x70, 0x86, 0x99, 0xa9, 0x51, 0xd0, 0xfd, 0x06, 0xde, 0x5b, 0x29,
	0x90, 0xf6, 0xe1, 0xba, 0x1c, 0x7c, 0x46, 0x6e, 0xde, 0xa0, 0x61, 0x82, 0x19, 0xd6, 0x26, 0x1d,
	0xd2, 0xbb, 0x72, 0xb0, 0x17, 0xda, 0xdb, 0x0b, 0x85, 0xe1, 0xbb, 0x95, 0x99, 0xb8, 0xb2, 0x43,
	0xf7, 0xc1, 0xd3, 0x0a, 0x79, 0xfb, 0x82, 0xdb, 0xbd, 0x13, 0xae, 0xd8, 0x13, 0xda, 0x43, 0xc7,
	0x0a, 0x79, 0xec, 0x86, 0xba, 0x4f, 0xe0, 0xf2, 0x02, 0xa1, 0x14, 0xbc, 0xe1, 0x89, 0xc8, 0xda,
	0x8d, 0x0e, 0xe9, 0xb5, 0x62, 0xf7, 0xa6, 0x7b, 0xd0, 0x62, 0x42, 0xe4, 0xa8, 0x35, 0xea, 0xb6,
	0xd7, 0x69, 0xf4, 0x5a, 0xf1, 0x12, 0xe8, 0xee, 0xc2, 0x8d, 0x23, 0x21, 0x2c, 0x41, 0x8c, 0x5a,
	0xc9, 0x4c, 0x63, 0xf7, 0x16, 0xd0, 0xf7, 0x4a, 0x30, 0x83, 0x55, 0xb4, 0x8f, 0x63, 0xac, 0xa0,
	0x29, 0xec, 0xda, 0xfa, 0x23, 0x33, 0x7c, 0xb4, 0x00, 0xe9, 0x3e, 0x5c, 0x74, 0x46, 0x95, 0xda,
	0x6f, 0x87, 0x73, 0xdb, 0xfe, 0x0b, 0x78, 0x69, 0xcb, 0x78, 0x3e, 0x43, 0x1f, 0x80, 0x67, 0xe5,
	0x95, 0x5a, 0x6f, 0xd6, 0x68, 0x8d, 0xdd, 0xc0, 0xc1, 0xcf, 0x06, 0x5c, 0xb5, 0xa5, 0x3e, 0xc6,
	0x7c, 0x96, 0x72, 0xa4, 0xaf, 0xa0, 0xf9, 0x01, 0x73, 0x9d, 0xca, 0x8c, 0x56, 0xec, 0x2d, 0xe1,
	0x18, 0x4f, 0xa6, 0xa8, 0x8d, 0x7f, 0x77, 0x43, 0xb7, 0xfc, 0xdc, 0x43, 0x68, 0xbe, 0x46, 0xe3,
	0xe2, 0xab, 0x3b, 0xef, 0xd7, 0x81, 0xf4, 0x29, 0x34, 0x4b, 0xdf, 0xea, 0x97, 0x82, 0x0a, 0x58,
	0x31, 0x99, 0xf6, 0x01, 0x96, 0x26, 0xd7, 0x53, 0xdc, 0xab, 0x80, 0xeb, 0xa1, 0x58, 0x96, 0x65,
	0x28, 0xdb, 0xb1, 0xac, 0x87, 0x48, 0xfb, 0xd0, 0x72, 0x01, 0x6e, 0x26, 0xe9, 0xd4, 0x80, 0x2b,
	0x99, 0x3f, 0x22, 0x2f, 0xee, 0x9f, 0x9e, 0x05, 0xe4, 0xf7, 0x59, 0xb0, 0xf3, 0xbd, 0x08, 0xc8,
	0x69, 0x11, 0x90, 0x5f, 0x45, 0x40, 0xfe, 0x14, 0x01, 0xf9, 0xf1, 0x37, 0xd8, 0xf9, 0xd4, 0x2c,
	0x57, 0x07, 0x97, 0xdc, 0x0f, 0x73, 0xf8, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x61, 0x9d, 0x3d, 0x33,
	0x36, 0x04, 0x00, 0x00,
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
	Version(ctx context.Context, in *v1alpha.VersionRequest, opts ...grpc.CallOption) (*v1alpha.VersionResponse, error)
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

func (c *nodesServiceClient) Version(ctx context.Context, in *v1alpha.VersionRequest, opts ...grpc.CallOption) (*v1alpha.VersionResponse, error) {
	out := new(v1alpha.VersionResponse)
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
	Version(context.Context, *v1alpha.VersionRequest) (*v1alpha.VersionResponse, error)
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
	in := new(v1alpha.VersionRequest)
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
		return srv.(NodesServiceServer).Version(ctx, req.(*v1alpha.VersionRequest))
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
