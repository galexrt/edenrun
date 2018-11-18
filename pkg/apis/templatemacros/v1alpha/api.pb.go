// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/galexrt/edenconfmgmt/pkg/apis/templatemacros/v1alpha/api.proto

package v1alpha

import (
	context "context"
	fmt "fmt"
	v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha"
	v1alpha1 "github.com/galexrt/edenconfmgmt/pkg/apis/events/v1alpha"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

// DataStore templatemacros for the data store.
type TemplateMacro struct {
	// Metadata for TemplateMacro object.
	Metadata *v1alpha.ObjectMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Spec for TemplateMacro.
	Spec                 *TemplateMacroSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TemplateMacro) Reset()         { *m = TemplateMacro{} }
func (m *TemplateMacro) String() string { return proto.CompactTextString(m) }
func (*TemplateMacro) ProtoMessage()    {}
func (*TemplateMacro) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6f457176fb6fb5, []int{0}
}

func (m *TemplateMacro) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemplateMacro.Unmarshal(m, b)
}
func (m *TemplateMacro) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemplateMacro.Marshal(b, m, deterministic)
}
func (m *TemplateMacro) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemplateMacro.Merge(m, src)
}
func (m *TemplateMacro) XXX_Size() int {
	return xxx_messageInfo_TemplateMacro.Size(m)
}
func (m *TemplateMacro) XXX_DiscardUnknown() {
	xxx_messageInfo_TemplateMacro.DiscardUnknown(m)
}

var xxx_messageInfo_TemplateMacro proto.InternalMessageInfo

func (m *TemplateMacro) GetMetadata() *v1alpha.ObjectMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *TemplateMacro) GetSpec() *TemplateMacroSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type TemplateMacroSpec struct {
	TemplateMacros       []*TemplateMacroItem `protobuf:"bytes,1,rep,name=templateMacros,proto3" json:"templateMacros,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TemplateMacroSpec) Reset()         { *m = TemplateMacroSpec{} }
func (m *TemplateMacroSpec) String() string { return proto.CompactTextString(m) }
func (*TemplateMacroSpec) ProtoMessage()    {}
func (*TemplateMacroSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6f457176fb6fb5, []int{1}
}

func (m *TemplateMacroSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemplateMacroSpec.Unmarshal(m, b)
}
func (m *TemplateMacroSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemplateMacroSpec.Marshal(b, m, deterministic)
}
func (m *TemplateMacroSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemplateMacroSpec.Merge(m, src)
}
func (m *TemplateMacroSpec) XXX_Size() int {
	return xxx_messageInfo_TemplateMacroSpec.Size(m)
}
func (m *TemplateMacroSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_TemplateMacroSpec.DiscardUnknown(m)
}

var xxx_messageInfo_TemplateMacroSpec proto.InternalMessageInfo

func (m *TemplateMacroSpec) GetTemplateMacros() []*TemplateMacroItem {
	if m != nil {
		return m.TemplateMacros
	}
	return nil
}

type TemplateMacroItem struct {
	// Action.
	Action string `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	// Parameters.
	Parameters           []*any.Any `protobuf:"bytes,2,rep,name=parameters,proto3" json:"parameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TemplateMacroItem) Reset()         { *m = TemplateMacroItem{} }
func (m *TemplateMacroItem) String() string { return proto.CompactTextString(m) }
func (*TemplateMacroItem) ProtoMessage()    {}
func (*TemplateMacroItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6f457176fb6fb5, []int{2}
}

func (m *TemplateMacroItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemplateMacroItem.Unmarshal(m, b)
}
func (m *TemplateMacroItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemplateMacroItem.Marshal(b, m, deterministic)
}
func (m *TemplateMacroItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemplateMacroItem.Merge(m, src)
}
func (m *TemplateMacroItem) XXX_Size() int {
	return xxx_messageInfo_TemplateMacroItem.Size(m)
}
func (m *TemplateMacroItem) XXX_DiscardUnknown() {
	xxx_messageInfo_TemplateMacroItem.DiscardUnknown(m)
}

var xxx_messageInfo_TemplateMacroItem proto.InternalMessageInfo

func (m *TemplateMacroItem) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *TemplateMacroItem) GetParameters() []*any.Any {
	if m != nil {
		return m.Parameters
	}
	return nil
}

// Get
type GetRequest struct {
	// GetOptions options for a GetRequest.
	GetOptions           *v1alpha.GetOptions `protobuf:"bytes,1,opt,name=getOptions,proto3" json:"getOptions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6f457176fb6fb5, []int{3}
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
	// TemplateMacros object.
	TemplateMacro *TemplateMacro `protobuf:"bytes,1,opt,name=TemplateMacro,proto3" json:"TemplateMacro,omitempty"`
	// Error object.
	Error                *v1alpha.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6f457176fb6fb5, []int{4}
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

func (m *GetResponse) GetTemplateMacro() *TemplateMacro {
	if m != nil {
		return m.TemplateMacro
	}
	return nil
}

func (m *GetResponse) GetError() *v1alpha.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

// List
type ListRequest struct {
	// ListOptions options for a ListRequest.
	ListOptions          *v1alpha.ListOptions `protobuf:"bytes,1,opt,name=listOptions,proto3" json:"listOptions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6f457176fb6fb5, []int{5}
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
	// TemplateMacro list.
	TemplateMacros []*TemplateMacro `protobuf:"bytes,1,rep,name=TemplateMacros,proto3" json:"TemplateMacros,omitempty"`
	// Error object.
	Error                *v1alpha.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6f457176fb6fb5, []int{6}
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

func (m *ListResponse) GetTemplateMacros() []*TemplateMacro {
	if m != nil {
		return m.TemplateMacros
	}
	return nil
}

func (m *ListResponse) GetError() *v1alpha.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

// Add
type AddRequest struct {
	// TemplateMacro object.
	TemplateMacro        *TemplateMacro `protobuf:"bytes,1,opt,name=TemplateMacro,proto3" json:"TemplateMacro,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6f457176fb6fb5, []int{7}
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

func (m *AddRequest) GetTemplateMacro() *TemplateMacro {
	if m != nil {
		return m.TemplateMacro
	}
	return nil
}

type AddResponse struct {
	// TemplateMacro object.
	TemplateMacro *TemplateMacro `protobuf:"bytes,1,opt,name=TemplateMacro,proto3" json:"TemplateMacro,omitempty"`
	// Error object.
	Error                *v1alpha.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()    {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6f457176fb6fb5, []int{8}
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

func (m *AddResponse) GetTemplateMacro() *TemplateMacro {
	if m != nil {
		return m.TemplateMacro
	}
	return nil
}

func (m *AddResponse) GetError() *v1alpha.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

// Update
type UpdateRequest struct {
	// TemplateMacro object.
	TemplateMacro        *TemplateMacro `protobuf:"bytes,1,opt,name=TemplateMacro,proto3" json:"TemplateMacro,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6f457176fb6fb5, []int{9}
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

func (m *UpdateRequest) GetTemplateMacro() *TemplateMacro {
	if m != nil {
		return m.TemplateMacro
	}
	return nil
}

type UpdateResponse struct {
	// TemplateMacro object.
	TemplateMacro *TemplateMacro `protobuf:"bytes,1,opt,name=TemplateMacro,proto3" json:"TemplateMacro,omitempty"`
	// Error object.
	Error                *v1alpha.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6f457176fb6fb5, []int{10}
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

func (m *UpdateResponse) GetTemplateMacro() *TemplateMacro {
	if m != nil {
		return m.TemplateMacro
	}
	return nil
}

func (m *UpdateResponse) GetError() *v1alpha.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

// Delete
type DeleteRequest struct {
	// TemplateMacro object.
	TemplateMacro        *TemplateMacro `protobuf:"bytes,1,opt,name=TemplateMacro,proto3" json:"TemplateMacro,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6f457176fb6fb5, []int{11}
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

func (m *DeleteRequest) GetTemplateMacro() *TemplateMacro {
	if m != nil {
		return m.TemplateMacro
	}
	return nil
}

type DeleteResponse struct {
	// TemplateMacro object.
	TemplateMacro *TemplateMacro `protobuf:"bytes,1,opt,name=TemplateMacro,proto3" json:"TemplateMacro,omitempty"`
	// Error object.
	Error                *v1alpha.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6f457176fb6fb5, []int{12}
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

func (m *DeleteResponse) GetTemplateMacro() *TemplateMacro {
	if m != nil {
		return m.TemplateMacro
	}
	return nil
}

func (m *DeleteResponse) GetError() *v1alpha.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

// Watch
type WatchRequest struct {
	// WatchOptions options for WatchRequest.
	WatchOptions         *v1alpha.WatchOptions `protobuf:"bytes,1,opt,name=watchOptions,proto3" json:"watchOptions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *WatchRequest) Reset()         { *m = WatchRequest{} }
func (m *WatchRequest) String() string { return proto.CompactTextString(m) }
func (*WatchRequest) ProtoMessage()    {}
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6f457176fb6fb5, []int{13}
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
	// TemplateMacro info for watch response.
	Event *v1alpha1.Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	// TemplateMacro for watch response.
	TemplateMacro *TemplateMacro `protobuf:"bytes,2,opt,name=TemplateMacro,proto3" json:"TemplateMacro,omitempty"`
	// Error object.
	Error                *v1alpha.Error `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *WatchResponse) Reset()         { *m = WatchResponse{} }
func (m *WatchResponse) String() string { return proto.CompactTextString(m) }
func (*WatchResponse) ProtoMessage()    {}
func (*WatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6f457176fb6fb5, []int{14}
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

func (m *WatchResponse) GetTemplateMacro() *TemplateMacro {
	if m != nil {
		return m.TemplateMacro
	}
	return nil
}

func (m *WatchResponse) GetError() *v1alpha.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*TemplateMacro)(nil), "templatemacros.v1alpha.TemplateMacro")
	proto.RegisterType((*TemplateMacroSpec)(nil), "templatemacros.v1alpha.TemplateMacroSpec")
	proto.RegisterType((*TemplateMacroItem)(nil), "templatemacros.v1alpha.TemplateMacroItem")
	proto.RegisterType((*GetRequest)(nil), "templatemacros.v1alpha.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "templatemacros.v1alpha.GetResponse")
	proto.RegisterType((*ListRequest)(nil), "templatemacros.v1alpha.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "templatemacros.v1alpha.ListResponse")
	proto.RegisterType((*AddRequest)(nil), "templatemacros.v1alpha.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "templatemacros.v1alpha.AddResponse")
	proto.RegisterType((*UpdateRequest)(nil), "templatemacros.v1alpha.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "templatemacros.v1alpha.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "templatemacros.v1alpha.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "templatemacros.v1alpha.DeleteResponse")
	proto.RegisterType((*WatchRequest)(nil), "templatemacros.v1alpha.WatchRequest")
	proto.RegisterType((*WatchResponse)(nil), "templatemacros.v1alpha.WatchResponse")
}

func init() {
	proto.RegisterFile("github.com/galexrt/edenconfmgmt/pkg/apis/templatemacros/v1alpha/api.proto", fileDescriptor_ab6f457176fb6fb5)
}

var fileDescriptor_ab6f457176fb6fb5 = []byte{
	// 689 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xcf, 0x4f, 0xd4, 0x40,
	0x14, 0x66, 0x58, 0x20, 0xfa, 0x16, 0x48, 0x1c, 0x95, 0x2c, 0x8d, 0x69, 0x48, 0x15, 0x23, 0x31,
	0xb6, 0x8a, 0x1e, 0x48, 0x8c, 0x92, 0x35, 0x2a, 0x41, 0xf9, 0x11, 0x2a, 0x4a, 0x34, 0x7a, 0x98,
	0x6d, 0x1f, 0x65, 0x75, 0xdb, 0xa9, 0xed, 0x80, 0x72, 0xf3, 0x62, 0xc2, 0xc5, 0xc4, 0x3f, 0xc6,
	0x3f, 0x82, 0xa3, 0x47, 0x8f, 0xb2, 0xfe, 0x23, 0xa6, 0xd3, 0x59, 0xb6, 0x2d, 0x14, 0xd6, 0x08,
	0xe1, 0xb6, 0xb3, 0xf3, 0xcd, 0xf7, 0xde, 0xf7, 0xbd, 0x79, 0x6f, 0x0a, 0xf3, 0x5e, 0x53, 0x6c,
	0x6c, 0x36, 0x4c, 0x87, 0xfb, 0x96, 0xc7, 0x5a, 0xf8, 0x39, 0x12, 0x16, 0xba, 0x18, 0x38, 0x3c,
	0x58, 0xf7, 0x3d, 0x5f, 0x58, 0xe1, 0x07, 0xcf, 0x62, 0x61, 0x33, 0xb6, 0x04, 0xfa, 0x61, 0x8b,
	0x09, 0xf4, 0x99, 0x13, 0xf1, 0xd8, 0xda, 0xba, 0xc3, 0x5a, 0xe1, 0x06, 0x4b, 0xf6, 0xcc, 0x30,
	0xe2, 0x82, 0xd3, 0xb1, 0x3c, 0xc2, 0x54, 0x08, 0xed, 0x56, 0x36, 0x04, 0xf7, 0xb8, 0x25, 0xe1,
	0x8d, 0xcd, 0x75, 0xb9, 0x92, 0x0b, 0xf9, 0x2b, 0xa5, 0xd1, 0xc6, 0x3d, 0xce, 0xbd, 0x16, 0x76,
	0x51, 0x2c, 0xd8, 0x56, 0x5b, 0xb3, 0x3d, 0x27, 0xeb, 0xf0, 0x08, 0x0f, 0xa6, 0xa8, 0xd5, 0x7b,
	0x26, 0xc0, 0x2d, 0x0c, 0xc4, 0x21, 0x2a, 0x8d, 0x1d, 0x02, 0x23, 0xab, 0x4a, 0xe8, 0x62, 0x22,
	0x94, 0xce, 0xc0, 0x39, 0x1f, 0x05, 0x73, 0x99, 0x60, 0x35, 0x32, 0x41, 0x6e, 0x54, 0xa7, 0xaf,
	0x98, 0x49, 0xfc, 0x8e, 0x01, 0xe6, 0x72, 0xe3, 0x3d, 0x3a, 0x62, 0x51, 0x61, 0xec, 0x7d, 0x34,
	0x7d, 0x00, 0x03, 0x71, 0x88, 0x4e, 0xad, 0x5f, 0x9e, 0x9a, 0x32, 0x0f, 0x37, 0xd0, 0xcc, 0x85,
	0x7b, 0x11, 0xa2, 0x63, 0xcb, 0x63, 0xc6, 0x3a, 0x5c, 0x38, 0xb0, 0x45, 0x57, 0x60, 0x54, 0x64,
	0xff, 0x8c, 0x6b, 0x64, 0xa2, 0xd2, 0x33, 0xfb, 0xbc, 0x40, 0xdf, 0x2e, 0x10, 0x18, 0xac, 0x10,
	0x27, 0x01, 0xd1, 0x31, 0x18, 0x62, 0x8e, 0x68, 0xf2, 0x40, 0x6a, 0x3e, 0x6f, 0xab, 0x15, 0xbd,
	0x07, 0x10, 0xb2, 0x88, 0xf9, 0x28, 0x30, 0x8a, 0x6b, 0xfd, 0x32, 0xf6, 0x25, 0x33, 0xad, 0xa9,
	0xd9, 0xa9, 0xa9, 0x59, 0x0f, 0xb6, 0xed, 0x0c, 0xce, 0x78, 0x0a, 0x30, 0x87, 0xc2, 0xc6, 0x8f,
	0x9b, 0x18, 0x0b, 0x3a, 0x03, 0xe0, 0xa1, 0x58, 0x0e, 0x13, 0xc2, 0x58, 0x79, 0x5a, 0xcb, 0x7b,
	0x3a, 0xb7, 0xbf, 0x6f, 0x67, 0xb0, 0xc6, 0x57, 0x02, 0x55, 0x49, 0x14, 0x87, 0x3c, 0x88, 0x91,
	0x3e, 0x2f, 0x14, 0x4b, 0x91, 0x4d, 0xf6, 0x64, 0x86, 0x5d, 0x28, 0xf4, 0x14, 0x0c, 0x62, 0x14,
	0xf1, 0x48, 0xd5, 0xeb, 0x62, 0x3e, 0xa3, 0x27, 0xc9, 0x96, 0x9d, 0x22, 0x8c, 0x67, 0x50, 0x5d,
	0x68, 0xc6, 0xfb, 0x82, 0xee, 0x43, 0xb5, 0xd5, 0x8c, 0x0b, 0x8a, 0xc6, 0xf3, 0xe7, 0x17, 0xba,
	0x00, 0x3b, 0x8b, 0x4e, 0x6e, 0xdc, 0x70, 0x4a, 0xa6, 0x44, 0x2d, 0xc2, 0xe8, 0xea, 0x61, 0x25,
	0xee, 0x51, 0x55, 0xe1, 0xf0, 0xbf, 0xc8, 0x7a, 0x0d, 0x50, 0x77, 0xdd, 0x8e, 0xaa, 0x93, 0x34,
	0x57, 0x56, 0x4e, 0x72, 0x9f, 0x71, 0xe5, 0xde, 0xc2, 0xc8, 0xcb, 0xd0, 0x65, 0x02, 0x4f, 0x45,
	0xe5, 0x0e, 0x81, 0xd1, 0x0e, 0xfd, 0xd9, 0x0b, 0x7d, 0x8c, 0x2d, 0x3c, 0x45, 0xa1, 0x1d, 0xfa,
	0x33, 0x16, 0xba, 0x04, 0xc3, 0x6b, 0x4c, 0x38, 0x1b, 0x1d, 0x9d, 0x0f, 0x61, 0xf8, 0x53, 0xb2,
	0xce, 0x77, 0xa3, 0x96, 0x67, 0x58, 0xcb, 0x20, 0xec, 0x1c, 0xde, 0xf8, 0x41, 0x60, 0x44, 0x11,
	0x2a, 0x65, 0x37, 0x61, 0x50, 0xbe, 0x17, 0x8a, 0xea, 0xb2, 0x99, 0xbe, 0x1e, 0xdd, 0x74, 0x92,
	0xa5, 0x9d, 0x62, 0x0e, 0xda, 0xd0, 0x7f, 0x12, 0x36, 0x54, 0x8e, 0xb3, 0x61, 0xfa, 0xdb, 0x40,
	0x71, 0x6c, 0xd0, 0x25, 0xa8, 0xcc, 0xa1, 0xa0, 0x46, 0x59, 0xe8, 0xee, 0x48, 0xd6, 0xae, 0x1e,
	0x89, 0x51, 0x3e, 0xac, 0xc0, 0x40, 0x32, 0xa8, 0x68, 0x29, 0x38, 0x33, 0x13, 0xb5, 0x6b, 0x47,
	0x83, 0x14, 0xe5, 0x12, 0x54, 0xea, 0xae, 0x5b, 0x9e, 0x62, 0x77, 0x1c, 0x95, 0xa7, 0x98, 0x1d,
	0x2b, 0x6b, 0x30, 0x94, 0xf6, 0x1f, 0x2d, 0x35, 0x3c, 0xd7, 0xfe, 0xda, 0xf5, 0xe3, 0x60, 0x5d,
	0xe2, 0xf4, 0xbe, 0x97, 0x13, 0xe7, 0xda, 0xad, 0x9c, 0xb8, 0xd0, 0x36, 0xaf, 0x60, 0x50, 0xde,
	0x36, 0x5a, 0x6a, 0x58, 0xf6, 0x76, 0x6b, 0x93, 0xc7, 0xa0, 0x52, 0xd6, 0xdb, 0xe4, 0xd1, 0xbb,
	0xdd, 0x3d, 0x9d, 0xfc, 0xda, 0xd3, 0xfb, 0xbe, 0xb4, 0x75, 0xb2, 0xdb, 0xd6, 0xc9, 0xcf, 0xb6,
	0x4e, 0x7e, 0xb7, 0x75, 0xf2, 0xfd, 0x8f, 0xde, 0xf7, 0x66, 0xf6, 0x3f, 0xbf, 0x0b, 0x1b, 0x43,
	0xf2, 0xad, 0xbf, 0xfb, 0x37, 0x00, 0x00, 0xff, 0xff, 0x44, 0x98, 0xa9, 0x1b, 0x61, 0x0a, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TemplateMacrosClient is the client API for TemplateMacros service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TemplateMacrosClient interface {
	// Get a TemplateMacro.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// List TemplateMacros.
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// Add a TemplateMacro.
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	// Update a TemplateMacro.
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Delete a TemplateMacro.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// Watch TemplateMacros.
	Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (TemplateMacros_WatchClient, error)
}

type templateMacrosClient struct {
	cc *grpc.ClientConn
}

func NewTemplateMacrosClient(cc *grpc.ClientConn) TemplateMacrosClient {
	return &templateMacrosClient{cc}
}

func (c *templateMacrosClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/templatemacros.v1alpha.TemplateMacros/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateMacrosClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/templatemacros.v1alpha.TemplateMacros/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateMacrosClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/templatemacros.v1alpha.TemplateMacros/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateMacrosClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/templatemacros.v1alpha.TemplateMacros/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateMacrosClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/templatemacros.v1alpha.TemplateMacros/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateMacrosClient) Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (TemplateMacros_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TemplateMacros_serviceDesc.Streams[0], "/templatemacros.v1alpha.TemplateMacros/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &templateMacrosWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TemplateMacros_WatchClient interface {
	Recv() (*WatchResponse, error)
	grpc.ClientStream
}

type templateMacrosWatchClient struct {
	grpc.ClientStream
}

func (x *templateMacrosWatchClient) Recv() (*WatchResponse, error) {
	m := new(WatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TemplateMacrosServer is the server API for TemplateMacros service.
type TemplateMacrosServer interface {
	// Get a TemplateMacro.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// List TemplateMacros.
	List(context.Context, *ListRequest) (*ListResponse, error)
	// Add a TemplateMacro.
	Add(context.Context, *AddRequest) (*AddResponse, error)
	// Update a TemplateMacro.
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Delete a TemplateMacro.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// Watch TemplateMacros.
	Watch(*WatchRequest, TemplateMacros_WatchServer) error
}

func RegisterTemplateMacrosServer(s *grpc.Server, srv TemplateMacrosServer) {
	s.RegisterService(&_TemplateMacros_serviceDesc, srv)
}

func _TemplateMacros_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateMacrosServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/templatemacros.v1alpha.TemplateMacros/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateMacrosServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TemplateMacros_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateMacrosServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/templatemacros.v1alpha.TemplateMacros/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateMacrosServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TemplateMacros_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateMacrosServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/templatemacros.v1alpha.TemplateMacros/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateMacrosServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TemplateMacros_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateMacrosServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/templatemacros.v1alpha.TemplateMacros/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateMacrosServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TemplateMacros_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateMacrosServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/templatemacros.v1alpha.TemplateMacros/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateMacrosServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TemplateMacros_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TemplateMacrosServer).Watch(m, &templateMacrosWatchServer{stream})
}

type TemplateMacros_WatchServer interface {
	Send(*WatchResponse) error
	grpc.ServerStream
}

type templateMacrosWatchServer struct {
	grpc.ServerStream
}

func (x *templateMacrosWatchServer) Send(m *WatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _TemplateMacros_serviceDesc = grpc.ServiceDesc{
	ServiceName: "templatemacros.v1alpha.TemplateMacros",
	HandlerType: (*TemplateMacrosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _TemplateMacros_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _TemplateMacros_List_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _TemplateMacros_Add_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TemplateMacros_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TemplateMacros_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _TemplateMacros_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/galexrt/edenconfmgmt/pkg/apis/templatemacros/v1alpha/api.proto",
}
