// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/galexrt/edenconfmgmt/pkg/apis/secrets/v1alpha/api.proto

package v1alpha

import (
	fmt "fmt"
	_ "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha"
	_ "github.com/galexrt/edenconfmgmt/pkg/apis/events/v1alpha"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Secret) Validate() error {
	if this.Metadata != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Metadata); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Metadata", err)
		}
	}
	if this.Spec != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Spec); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Spec", err)
		}
	}
	return nil
}
func (this *SecretSpec) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *GetRequest) Validate() error {
	if this.GetOptions != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.GetOptions); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("GetOptions", err)
		}
	}
	return nil
}
func (this *GetResponse) Validate() error {
	if this.Secret != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Secret); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Secret", err)
		}
	}
	if this.Error != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Error); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Error", err)
		}
	}
	return nil
}
func (this *ListRequest) Validate() error {
	if this.ListOptions != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ListOptions); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ListOptions", err)
		}
	}
	return nil
}
func (this *ListResponse) Validate() error {
	for _, item := range this.Secrets {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Secrets", err)
			}
		}
	}
	if this.Error != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Error); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Error", err)
		}
	}
	return nil
}
func (this *AddRequest) Validate() error {
	if this.Secret != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Secret); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Secret", err)
		}
	}
	return nil
}
func (this *AddResponse) Validate() error {
	if this.Secret != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Secret); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Secret", err)
		}
	}
	if this.Error != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Error); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Error", err)
		}
	}
	return nil
}
func (this *UpdateRequest) Validate() error {
	if this.Secret != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Secret); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Secret", err)
		}
	}
	return nil
}
func (this *UpdateResponse) Validate() error {
	if this.Secret != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Secret); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Secret", err)
		}
	}
	if this.Error != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Error); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Error", err)
		}
	}
	return nil
}
func (this *DeleteRequest) Validate() error {
	if this.Secret != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Secret); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Secret", err)
		}
	}
	return nil
}
func (this *DeleteResponse) Validate() error {
	if this.Secret != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Secret); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Secret", err)
		}
	}
	if this.Error != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Error); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Error", err)
		}
	}
	return nil
}
func (this *WatchRequest) Validate() error {
	if this.WatchOptions != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.WatchOptions); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("WatchOptions", err)
		}
	}
	return nil
}
func (this *WatchResponse) Validate() error {
	if this.Event != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Event); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Event", err)
		}
	}
	if this.Secret != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Secret); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Secret", err)
		}
	}
	if this.Error != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Error); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Error", err)
		}
	}
	return nil
}