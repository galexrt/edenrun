// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/galexrt/edenconfmgmt/pkg/apis/events/v1alpha/api.proto

package v1alpha

import (
	fmt "fmt"
	v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha"
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

type Event struct {
	// Metadata for the Event object.
	Metadata *v1alpha.ObjectMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Spec for Event.
	Spec                 *EventSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_57c26361c2d00172, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetMetadata() *v1alpha.ObjectMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Event) GetSpec() *EventSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type EventSpec struct {
	// Type of the Event.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Summary of Event.
	Summary              string   `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventSpec) Reset()         { *m = EventSpec{} }
func (m *EventSpec) String() string { return proto.CompactTextString(m) }
func (*EventSpec) ProtoMessage()    {}
func (*EventSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_57c26361c2d00172, []int{1}
}

func (m *EventSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventSpec.Unmarshal(m, b)
}
func (m *EventSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventSpec.Marshal(b, m, deterministic)
}
func (m *EventSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSpec.Merge(m, src)
}
func (m *EventSpec) XXX_Size() int {
	return xxx_messageInfo_EventSpec.Size(m)
}
func (m *EventSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSpec.DiscardUnknown(m)
}

var xxx_messageInfo_EventSpec proto.InternalMessageInfo

func (m *EventSpec) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *EventSpec) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
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
	return fileDescriptor_57c26361c2d00172, []int{2}
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
	// Event object.
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_57c26361c2d00172, []int{3}
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

func (m *GetResponse) GetEvent() *Event {
	if m != nil {
		return m.Event
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
	return fileDescriptor_57c26361c2d00172, []int{4}
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
	// Event list.
	Events               []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_57c26361c2d00172, []int{5}
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

func (m *ListResponse) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

// Add
type AddRequest struct {
	// Event object.
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_57c26361c2d00172, []int{6}
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

func (m *AddRequest) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type AddResponse struct {
	// Event object.
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()    {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_57c26361c2d00172, []int{7}
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

func (m *AddResponse) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

// Update
type UpdateRequest struct {
	// Event object.
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_57c26361c2d00172, []int{8}
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

func (m *UpdateRequest) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type UpdateResponse struct {
	// Event object.
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_57c26361c2d00172, []int{9}
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

func (m *UpdateResponse) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

// Delete
type DeleteRequest struct {
	// Event object.
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_57c26361c2d00172, []int{10}
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

func (m *DeleteRequest) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type DeleteResponse struct {
	// Event object.
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_57c26361c2d00172, []int{11}
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

func (m *DeleteResponse) GetEvent() *Event {
	if m != nil {
		return m.Event
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
	return fileDescriptor_57c26361c2d00172, []int{12}
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
	// Event for watch stream.
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WatchResponse) Reset()         { *m = WatchResponse{} }
func (m *WatchResponse) String() string { return proto.CompactTextString(m) }
func (*WatchResponse) ProtoMessage()    {}
func (*WatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_57c26361c2d00172, []int{13}
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

func (m *WatchResponse) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "events.v1alpha.Event")
	proto.RegisterType((*EventSpec)(nil), "events.v1alpha.EventSpec")
	proto.RegisterType((*GetRequest)(nil), "events.v1alpha.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "events.v1alpha.GetResponse")
	proto.RegisterType((*ListRequest)(nil), "events.v1alpha.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "events.v1alpha.ListResponse")
	proto.RegisterType((*AddRequest)(nil), "events.v1alpha.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "events.v1alpha.AddResponse")
	proto.RegisterType((*UpdateRequest)(nil), "events.v1alpha.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "events.v1alpha.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "events.v1alpha.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "events.v1alpha.DeleteResponse")
	proto.RegisterType((*WatchRequest)(nil), "events.v1alpha.WatchRequest")
	proto.RegisterType((*WatchResponse)(nil), "events.v1alpha.WatchResponse")
}

func init() {
	proto.RegisterFile("github.com/galexrt/edenconfmgmt/pkg/apis/events/v1alpha/api.proto", fileDescriptor_57c26361c2d00172)
}

var fileDescriptor_57c26361c2d00172 = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0xc9, 0x07, 0x64, 0x12, 0x7a, 0x58, 0x09, 0xc9, 0x35, 0xad, 0x55, 0xf9, 0x84, 0x84,
	0x62, 0x43, 0x39, 0x40, 0x0b, 0x0d, 0x0a, 0x82, 0x46, 0x42, 0x85, 0x4a, 0x46, 0x55, 0x25, 0x6e,
	0x1b, 0x7b, 0xea, 0x18, 0x62, 0x7b, 0xc9, 0x6e, 0x0a, 0xbd, 0xf1, 0x13, 0xf8, 0x13, 0xfc, 0x97,
	0x1e, 0x39, 0x72, 0xa4, 0xe1, 0x8f, 0x20, 0xaf, 0xed, 0xc4, 0x8e, 0x13, 0x89, 0xfa, 0xb6, 0xbb,
	0xf3, 0xe6, 0xbd, 0x37, 0x99, 0x97, 0x04, 0xfa, 0x9e, 0x2f, 0x46, 0xd3, 0xa1, 0xe9, 0x44, 0x81,
	0xe5, 0xd1, 0x31, 0x7e, 0x9b, 0x08, 0x0b, 0x5d, 0x0c, 0x9d, 0x28, 0x3c, 0x0f, 0xbc, 0x40, 0x58,
	0xec, 0xb3, 0x67, 0x51, 0xe6, 0x73, 0x0b, 0x2f, 0x30, 0x14, 0xdc, 0xba, 0x78, 0x4c, 0xc7, 0x6c,
	0x44, 0xe3, 0x37, 0x93, 0x4d, 0x22, 0x11, 0x91, 0xcd, 0xa4, 0x62, 0xa6, 0x15, 0xad, 0x9b, 0xa7,
	0x8c, 0xbc, 0xc8, 0x92, 0xb0, 0xe1, 0xf4, 0x5c, 0xde, 0xe4, 0x45, 0x9e, 0x92, 0x76, 0xed, 0xe5,
	0x7f, 0x3b, 0x70, 0xa2, 0x09, 0x96, 0xf5, 0x0d, 0x06, 0x8d, 0x37, 0xb1, 0x03, 0xf2, 0x0c, 0xee,
	0x04, 0x28, 0xa8, 0x4b, 0x05, 0x55, 0x95, 0x5d, 0xe5, 0x41, 0x7b, 0x6f, 0xdb, 0x8c, 0x7b, 0x32,
	0x67, 0xe6, 0xc9, 0xf0, 0x13, 0x3a, 0xe2, 0x5d, 0x8a, 0xb1, 0xe7, 0x68, 0xd2, 0x85, 0x3a, 0x67,
	0xe8, 0xa8, 0xb7, 0x64, 0xd7, 0x96, 0x59, 0x9c, 0xc8, 0x94, 0xf4, 0x1f, 0x18, 0x3a, 0xb6, 0x84,
	0x19, 0xfb, 0xd0, 0x9a, 0x3f, 0x11, 0x02, 0x75, 0x71, 0xc9, 0x50, 0x2a, 0xb6, 0x6c, 0x79, 0x26,
	0x2a, 0xdc, 0xe6, 0xd3, 0x20, 0xa0, 0x93, 0x4b, 0x49, 0xd9, 0xb2, 0xb3, 0xab, 0x71, 0x04, 0x30,
	0x40, 0x61, 0xe3, 0x97, 0x29, 0xf2, 0xd8, 0x31, 0x78, 0x28, 0x4e, 0x98, 0xf0, 0xa3, 0x90, 0xa7,
	0x9e, 0xd5, 0xa2, 0xe7, 0xc1, 0xbc, 0x6e, 0xe7, 0xb0, 0xc6, 0x01, 0xb4, 0x25, 0x0f, 0x67, 0x51,
	0xc8, 0x91, 0x3c, 0x84, 0x86, 0xf4, 0x9c, 0x72, 0xdc, 0x5b, 0x39, 0x81, 0x9d, 0x60, 0x8c, 0xb7,
	0xd0, 0x3e, 0xf6, 0xf9, 0xdc, 0xc4, 0x73, 0x68, 0x8f, 0x7d, 0xbe, 0xe4, 0x62, 0xab, 0xe8, 0xe2,
	0x78, 0x01, 0xb0, 0xf3, 0x68, 0xe3, 0x10, 0x3a, 0x09, 0x57, 0x6a, 0xa4, 0x0b, 0xcd, 0x44, 0x5a,
	0x55, 0x76, 0x6b, 0xeb, 0x9d, 0xa4, 0x20, 0x63, 0x1f, 0xa0, 0xef, 0xba, 0x99, 0x93, 0x1b, 0x4d,
	0x71, 0x00, 0x6d, 0xd9, 0x5a, 0xe5, 0x13, 0x78, 0x01, 0x77, 0x4f, 0x99, 0x4b, 0x05, 0x56, 0x52,
	0x3e, 0x84, 0xcd, 0xac, 0xbb, 0xa2, 0xf8, 0x6b, 0x1c, 0x63, 0x75, 0xf1, 0xac, 0xbb, 0x8a, 0xf8,
	0x7b, 0xe8, 0x9c, 0x51, 0xe1, 0x8c, 0x32, 0xed, 0x1e, 0x74, 0xbe, 0xc6, 0xf7, 0xe2, 0xf6, 0xb5,
	0xe2, 0xf6, 0xcf, 0x72, 0x08, 0xbb, 0x80, 0x8f, 0x87, 0x49, 0xf9, 0x2a, 0xb8, 0xd9, 0xfb, 0x59,
	0x83, 0xa6, 0x7c, 0xe0, 0xa4, 0x07, 0xb5, 0x01, 0x0a, 0xa2, 0x2d, 0xe3, 0x17, 0xdf, 0x16, 0xed,
	0xfe, 0xca, 0x5a, 0xaa, 0xdb, 0x87, 0x7a, 0x1c, 0x44, 0x52, 0x02, 0xe5, 0xa2, 0xae, 0x6d, 0xaf,
	0x2e, 0xa6, 0x14, 0x3d, 0xa8, 0xf5, 0x5d, 0xb7, 0x6c, 0x61, 0x91, 0xd0, 0xb2, 0x85, 0x7c, 0x04,
	0x07, 0xd0, 0x4c, 0x72, 0x41, 0x76, 0x96, 0x61, 0x85, 0xb4, 0x69, 0xfa, 0xba, 0xf2, 0x82, 0x28,
	0xd9, 0x71, 0x99, 0xa8, 0x90, 0x9c, 0x32, 0xd1, 0x52, 0x34, 0x8e, 0xa0, 0x21, 0xb7, 0x43, 0x4a,
	0x83, 0xe7, 0x43, 0xa0, 0xed, 0xac, 0xa9, 0x26, 0x2c, 0x8f, 0x94, 0x57, 0xa7, 0x57, 0xd7, 0xba,
	0xf2, 0xfb, 0x5a, 0xdf, 0xf8, 0x3e, 0xd3, 0x95, 0xab, 0x99, 0xae, 0xfc, 0x9a, 0xe9, 0xca, 0x9f,
	0x99, 0xae, 0xfc, 0xf8, 0xab, 0x6f, 0x7c, 0x7c, 0x5a, 0xf1, 0x3f, 0x64, 0xd8, 0x94, 0x3f, 0xe0,
	0x4f, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xc3, 0x6b, 0xe9, 0x85, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventsClient is the client API for Events service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventsClient interface {
	// Get a Event.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// List Events.
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// Add a Event.
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	// Update a Event.
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Delete a Event.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// Watch for Events.
	Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (Events_WatchClient, error)
}

type eventsClient struct {
	cc *grpc.ClientConn
}

func NewEventsClient(cc *grpc.ClientConn) EventsClient {
	return &eventsClient{cc}
}

func (c *eventsClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/events.v1alpha.Events/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/events.v1alpha.Events/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/events.v1alpha.Events/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/events.v1alpha.Events/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/events.v1alpha.Events/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsClient) Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (Events_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Events_serviceDesc.Streams[0], "/events.v1alpha.Events/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventsWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Events_WatchClient interface {
	Recv() (*WatchResponse, error)
	grpc.ClientStream
}

type eventsWatchClient struct {
	grpc.ClientStream
}

func (x *eventsWatchClient) Recv() (*WatchResponse, error) {
	m := new(WatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventsServer is the server API for Events service.
type EventsServer interface {
	// Get a Event.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// List Events.
	List(context.Context, *ListRequest) (*ListResponse, error)
	// Add a Event.
	Add(context.Context, *AddRequest) (*AddResponse, error)
	// Update a Event.
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Delete a Event.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// Watch for Events.
	Watch(*WatchRequest, Events_WatchServer) error
}

func RegisterEventsServer(s *grpc.Server, srv EventsServer) {
	s.RegisterService(&_Events_serviceDesc, srv)
}

func _Events_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/events.v1alpha.Events/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/events.v1alpha.Events/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/events.v1alpha.Events/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/events.v1alpha.Events/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/events.v1alpha.Events/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventsServer).Watch(m, &eventsWatchServer{stream})
}

type Events_WatchServer interface {
	Send(*WatchResponse) error
	grpc.ServerStream
}

type eventsWatchServer struct {
	grpc.ServerStream
}

func (x *eventsWatchServer) Send(m *WatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Events_serviceDesc = grpc.ServiceDesc{
	ServiceName: "events.v1alpha.Events",
	HandlerType: (*EventsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Events_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Events_List_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Events_Add_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Events_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Events_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Events_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/galexrt/edenconfmgmt/pkg/apis/events/v1alpha/api.proto",
}
