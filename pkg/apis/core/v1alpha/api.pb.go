// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha/api.proto

package v1alpha

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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

type ObjectMetadata struct {
	// ID of object.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of object.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Labels of object.
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Annotations of object.
	Annotations          map[string]string `protobuf:"bytes,4,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ObjectMetadata) Reset()         { *m = ObjectMetadata{} }
func (m *ObjectMetadata) String() string { return proto.CompactTextString(m) }
func (*ObjectMetadata) ProtoMessage()    {}
func (*ObjectMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_30dbcc7c2d394087, []int{0}
}

func (m *ObjectMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectMetadata.Unmarshal(m, b)
}
func (m *ObjectMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectMetadata.Marshal(b, m, deterministic)
}
func (m *ObjectMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectMetadata.Merge(m, src)
}
func (m *ObjectMetadata) XXX_Size() int {
	return xxx_messageInfo_ObjectMetadata.Size(m)
}
func (m *ObjectMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectMetadata proto.InternalMessageInfo

func (m *ObjectMetadata) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ObjectMetadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ObjectMetadata) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ObjectMetadata) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

type VersionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionRequest) Reset()         { *m = VersionRequest{} }
func (m *VersionRequest) String() string { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()    {}
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30dbcc7c2d394087, []int{1}
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
	return fileDescriptor_30dbcc7c2d394087, []int{2}
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

func init() {
	proto.RegisterType((*ObjectMetadata)(nil), "core.v1alpha.ObjectMetadata")
	proto.RegisterMapType((map[string]string)(nil), "core.v1alpha.ObjectMetadata.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "core.v1alpha.ObjectMetadata.LabelsEntry")
	proto.RegisterType((*VersionRequest)(nil), "core.v1alpha.VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "core.v1alpha.VersionResponse")
}

func init() {
	proto.RegisterFile("github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha/api.proto", fileDescriptor_30dbcc7c2d394087)
}

var fileDescriptor_30dbcc7c2d394087 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xbf, 0x4b, 0xc3, 0x40,
	0x14, 0xc7, 0x7b, 0x49, 0x7f, 0xe0, 0x55, 0x6a, 0x39, 0x1c, 0x8e, 0x0e, 0xa1, 0x14, 0x84, 0x2e,
	0x4d, 0x50, 0x17, 0x75, 0xf0, 0x17, 0xb8, 0x29, 0x85, 0x0c, 0x0e, 0x6e, 0x97, 0xe4, 0x9a, 0xa6,
	0x4d, 0xee, 0x62, 0xee, 0x52, 0xec, 0x26, 0xf8, 0x0f, 0xf8, 0x37, 0x39, 0x75, 0x74, 0x74, 0xb4,
	0xf1, 0x1f, 0x91, 0x5c, 0x52, 0x8d, 0x0e, 0x82, 0xdb, 0xfb, 0x7e, 0x78, 0xef, 0x93, 0xf7, 0xc2,
	0xc1, 0x33, 0x3f, 0x90, 0xd3, 0xd4, 0x31, 0x5d, 0x1e, 0x59, 0x3e, 0x09, 0xe9, 0x43, 0x22, 0x2d,
	0xea, 0x51, 0xe6, 0x72, 0x36, 0x89, 0xfc, 0x48, 0x5a, 0xf1, 0xdc, 0xb7, 0x48, 0x1c, 0x08, 0xcb,
	0xe5, 0x09, 0xb5, 0x16, 0xfb, 0x24, 0x8c, 0xa7, 0x24, 0x27, 0x66, 0x9c, 0x70, 0xc9, 0xd1, 0x76,
	0xce, 0xcd, 0x92, 0xf7, 0x46, 0x55, 0x1d, 0xf7, 0xb9, 0xa5, 0x9a, 0x9c, 0x74, 0xa2, 0x92, 0x0a,
	0xaa, 0x2a, 0x86, 0x07, 0x2f, 0x1a, 0xec, 0x8c, 0x9d, 0x19, 0x75, 0xe5, 0x0d, 0x95, 0xc4, 0x23,
	0x92, 0xa0, 0x0e, 0xd4, 0x02, 0x0f, 0x83, 0x3e, 0x18, 0x6e, 0xd9, 0x5a, 0xe0, 0x21, 0x04, 0xeb,
	0x8c, 0x44, 0x14, 0x6b, 0x8a, 0xa8, 0x1a, 0x9d, 0xc3, 0x66, 0x48, 0x1c, 0x1a, 0x0a, 0xac, 0xf7,
	0xf5, 0x61, 0xfb, 0x60, 0x68, 0x56, 0x97, 0x30, 0x7f, 0x1a, 0xcd, 0x6b, 0xd5, 0x7a, 0xc5, 0x64,
	0xb2, 0xb4, 0xcb, 0x39, 0x34, 0x86, 0x6d, 0xc2, 0x18, 0x97, 0x44, 0x06, 0x9c, 0x09, 0x5c, 0x57,
	0x9a, 0xd1, 0x9f, 0x9a, 0x8b, 0xef, 0xfe, 0xc2, 0x55, 0x35, 0xf4, 0x8e, 0x61, 0xbb, 0xf2, 0x1d,
	0xd4, 0x85, 0xfa, 0x9c, 0x2e, 0xcb, 0x33, 0xf2, 0x12, 0xed, 0xc2, 0xc6, 0x82, 0x84, 0xe9, 0xe6,
	0x90, 0x22, 0x9c, 0x68, 0x47, 0xa0, 0x77, 0x0a, 0xbb, 0xbf, 0xdd, 0xff, 0x99, 0x1f, 0x74, 0x61,
	0xe7, 0x96, 0x26, 0x22, 0xe0, 0xcc, 0xa6, 0xf7, 0x29, 0x15, 0x72, 0xf0, 0x04, 0xe0, 0xce, 0x17,
	0x12, 0x31, 0x67, 0x82, 0x22, 0x0c, 0x5b, 0x8b, 0x02, 0x95, 0xd6, 0x4d, 0xcc, 0xcd, 0x11, 0x99,
	0xf1, 0x44, 0x99, 0x75, 0xbb, 0x08, 0x8a, 0x06, 0x8c, 0x27, 0x58, 0x2f, 0x69, 0x1e, 0x72, 0x1a,
	0x13, 0xe9, 0x4e, 0x71, 0xbd, 0xa0, 0x2a, 0xe4, 0x54, 0x48, 0x22, 0x29, 0x6e, 0x14, 0xbb, 0xa9,
	0x70, 0xb9, 0xb7, 0x5a, 0x1b, 0xe0, 0x6d, 0x6d, 0xd4, 0x1e, 0x33, 0x03, 0xac, 0x32, 0x03, 0xbc,
	0x66, 0x06, 0x78, 0xcf, 0x0c, 0xf0, 0xfc, 0x61, 0xd4, 0xee, 0x5a, 0xe5, 0x6f, 0x76, 0x9a, 0xea,
	0x29, 0x1c, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x99, 0xda, 0xe6, 0xeb, 0x8a, 0x02, 0x00, 0x00,
}
