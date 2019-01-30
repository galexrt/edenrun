// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/galexrt/edenconfmgmt/pkg/apis/agentconfigs/v1alpha/api.proto

package v1alpha

import (
	context "context"
	fmt "fmt"
	_ "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1"
	_ "github.com/galexrt/edenconfmgmt/pkg/apis/events/v1alpha"
	_ "github.com/galexrt/edenconfmgmt/pkg/grpc/plugins/apiserver"
	github_com_galexrt_edenconfmgmt_pkg_grpc_plugins_apiserver "github.com/galexrt/edenconfmgmt/pkg/grpc/plugins/apiserver"
	github_com_galexrt_edenconfmgmt_pkg_store_object "github.com/galexrt/edenconfmgmt/pkg/store/object"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// AgentConfigsService !!WARNING!! This client must only be used by internal code!
// All "external" clients must use the actual GRPC API Client.
type AgentConfigsService struct {
	AgentConfigsServer
	store *github_com_galexrt_edenconfmgmt_pkg_store_object.Store
}

// NewAgentConfigsService returns a new client to be used to access AgentConfigs API with the given storage.
func NewAgentConfigsService(store *github_com_galexrt_edenconfmgmt_pkg_store_object.Store) *AgentConfigsService {
	return &AgentConfigsService{
		store: store,
	}
}

// Get
func (this *AgentConfigsService) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	out, err := this.store.Get(ctx, req.GetOptions())
	if err != nil {
		return nil, err
	}
	if req.Options != nil {
		req.Options.Namespace = ""
	}
	target := &AgentConfig{}
	err = target.Unmarshal(out)
	if err != nil {
		return nil, err
	}
	return &GetResponse{
		AgentConfig: target,
	}, nil
}

// List
func (this *AgentConfigsService) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	out, err := this.store.List(ctx, req.GetOptions())
	if err != nil {
		return nil, err
	}
	if req.Options != nil {
		req.Options.Namespace = ""
	}
	list := &AgentConfigList{}
	list.Items = make([]*AgentConfig, len(out))
	for k := range out {
		list.Items[k] = &AgentConfig{}
		err = list.Items[k].Unmarshal(out[k])
		if err != nil {
			return nil, err
		}
	}
	return &ListResponse{
		AgentConfigList: list,
	}, nil
}

// Create
func (this *AgentConfigsService) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	if err := req.GetAgentConfig().SetDefaults(github_com_galexrt_edenconfmgmt_pkg_grpc_plugins_apiserver.MethodCreate); err != nil {
		return nil, err
	}
	if req.Options != nil {
		req.Options.Namespace = ""
	}
	req.GetAgentConfig().GetMetadata().Namespace = ""
	if err := this.store.Create(ctx, req.GetAgentConfig(), req.GetOptions()); err != nil {
		return nil, err
	}
	return &CreateResponse{
		AgentConfig: req.GetAgentConfig(),
	}, nil
}

// Update
func (this *AgentConfigsService) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	if err := req.GetAgentConfig().SetDefaults(github_com_galexrt_edenconfmgmt_pkg_grpc_plugins_apiserver.MethodUpdate); err != nil {
		return nil, err
	}
	if req.Options != nil {
		req.Options.Namespace = ""
	}
	req.GetAgentConfig().GetMetadata().Namespace = ""
	if err := this.store.Update(ctx, req.GetAgentConfig(), req.GetOptions()); err != nil {
		return nil, err
	}
	return &UpdateResponse{
		AgentConfig: req.GetAgentConfig(),
	}, nil
}

// Delete
func (this *AgentConfigsService) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	if req.Options != nil {
		req.Options.Namespace = ""
	}
	return &DeleteResponse{}, this.store.Delete(ctx, req.GetAgentConfig(), req.GetOptions())
}

// Watch
func (this *AgentConfigsService) Watch(req *WatchRequest, stream AgentConfigs_WatchServer) error {
	req.GetOptions().Namespace = ""
	watch, err := this.store.Watch(stream.Context(), req.GetOptions())
	if err != nil {
		return nil
	}
	for {
		select {
		case out := <-watch:
			target := &AgentConfig{}
			if err = target.Unmarshal(out.Value); err != nil {
				return err
			}
			if err = stream.Send(&WatchResponse{
				AgentConfig: target,
			}); err != nil {
				return err
			}
		case <-stream.Context().Done():
			return stream.Context().Err()
		}
	}
}
