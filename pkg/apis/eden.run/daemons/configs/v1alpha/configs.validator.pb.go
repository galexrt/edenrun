// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/galexrt/edenrun/pkg/apis/eden.run/daemons/configs/v1alpha/configs.proto

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

func (this *Config) Validate() error {
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
func (this *ConfigList) Validate() error {
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
func (this *ConfigSpec) Validate() error {
	if this.Logging != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Logging); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Logging", err)
		}
	}
	if this.Store != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Store); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Store", err)
		}
	}
	return nil
}
func (this *Logging) Validate() error {
	return nil
}
func (this *Store) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	if this.Cache != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Cache); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Cache", err)
		}
	}
	return nil
}
func (this *Data) Validate() error {
	if this.Etcd != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Etcd); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Etcd", err)
		}
	}
	return nil
}
func (this *Cache) Validate() error {
	return nil
}
func (this *ETCD) Validate() error {
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
	if this.Config != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Config); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Config", err)
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
	if this.ConfigList != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ConfigList); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ConfigList", err)
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
	if this.Config != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Config); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Config", err)
		}
	}
	return nil
}
func (this *CreateResponse) Validate() error {
	if this.Config != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Config); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Config", err)
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
	if this.Config != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Config); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Config", err)
		}
	}
	return nil
}
func (this *UpdateResponse) Validate() error {
	if this.Config != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Config); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Config", err)
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
	if this.Config != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Config); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Config", err)
		}
	}
	return nil
}
func (this *DeleteResponse) Validate() error {
	if this.Config != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Config); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Config", err)
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
	if this.Config != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Config); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Config", err)
		}
	}
	return nil
}