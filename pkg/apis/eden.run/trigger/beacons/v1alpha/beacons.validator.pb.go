// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/galexrt/edenrun/pkg/apis/eden.run/trigger/beacons/v1alpha/beacons.proto

package v1alpha

import (
	fmt "fmt"
	_ "github.com/galexrt/edenrun/pkg/apis/eden.run/core/meta/v1"
	_ "github.com/galexrt/edenrun/pkg/grpc/plugins/apiserver"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Beacon) Validate() error {
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
func (this *BeaconList) Validate() error {
	if nil == this.Metadata {
		return github_com_mwitkow_go_proto_validators.FieldError("Metadata", fmt.Errorf("message must exist"))
	}
	if this.Metadata != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Metadata); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Metadata", err)
		}
	}
	for _, item := range this.Items {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Items", err)
			}
		}
	}
	return nil
}
func (this *BeaconSpec) Validate() error {
	return nil
}
func (this *GetRequest) Validate() error {
	if this.Options != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Options); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Options", err)
		}
	}
	return nil
}
func (this *GetResponse) Validate() error {
	if this.Beacon != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Beacon); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Beacon", err)
		}
	}
	return nil
}
func (this *ListRequest) Validate() error {
	if this.Options != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Options); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Options", err)
		}
	}
	return nil
}
func (this *ListResponse) Validate() error {
	if this.BeaconList != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.BeaconList); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("BeaconList", err)
		}
	}
	return nil
}
func (this *CreateRequest) Validate() error {
	if this.Options != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Options); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Options", err)
		}
	}
	if this.Beacon != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Beacon); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Beacon", err)
		}
	}
	return nil
}
func (this *CreateResponse) Validate() error {
	if this.Beacon != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Beacon); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Beacon", err)
		}
	}
	return nil
}
func (this *UpdateRequest) Validate() error {
	if this.Options != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Options); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Options", err)
		}
	}
	if this.Beacon != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Beacon); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Beacon", err)
		}
	}
	return nil
}
func (this *UpdateResponse) Validate() error {
	if this.Beacon != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Beacon); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Beacon", err)
		}
	}
	return nil
}
func (this *DeleteRequest) Validate() error {
	if this.Options != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Options); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Options", err)
		}
	}
	if this.Beacon != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Beacon); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Beacon", err)
		}
	}
	return nil
}
func (this *DeleteResponse) Validate() error {
	if this.Beacon != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Beacon); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Beacon", err)
		}
	}
	return nil
}
func (this *WatchRequest) Validate() error {
	if this.Options != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Options); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Options", err)
		}
	}
	return nil
}
func (this *WatchResponse) Validate() error {
	if this.Beacon != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Beacon); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Beacon", err)
		}
	}
	return nil
}