// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha/api.proto

package v1alpha

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ObjectMetadata) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *VersionRequest) Validate() error {
	return nil
}
func (this *VersionResponse) Validate() error {
	return nil
}
func (this *Include) Validate() error {
	return nil
}
func (this *Conditions) Validate() error {
	if this.When != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.When); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("When", err)
		}
	}
	if this.Success != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Success); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Success", err)
		}
	}
	return nil
}
func (this *Condition) Validate() error {
	if this.Retry != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Retry); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Retry", err)
		}
	}
	return nil
}
func (this *Retry) Validate() error {
	if this.Interval != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Interval); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Interval", err)
		}
	}
	return nil
}
func (this *GetOptions) Validate() error {
	return nil
}
func (this *ListOptions) Validate() error {
	return nil
}
func (this *WatchOptions) Validate() error {
	return nil
}
func (this *Error) Validate() error {
	return nil
}
