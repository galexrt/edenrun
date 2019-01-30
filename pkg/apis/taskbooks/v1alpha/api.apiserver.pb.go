// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/galexrt/edenconfmgmt/pkg/apis/taskbooks/v1alpha/api.proto

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
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// TaskBooksService !!WARNING!! This client must only be used by internal code!
// All "external" clients must use the actual GRPC API Client.
type TaskBooksService struct {
	TaskBooksServer
	store *github_com_galexrt_edenconfmgmt_pkg_store_object.Store
}

// NewTaskBooksService returns a new client to be used to access TaskBooks API with the given storage.
func NewTaskBooksService(store *github_com_galexrt_edenconfmgmt_pkg_store_object.Store) *TaskBooksService {
	return &TaskBooksService{
		store: store,
	}
}

// Get
func (this *TaskBooksService) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	out, err := this.store.Get(ctx, req.GetOptions())
	if err != nil {
		return nil, err
	}
	if req.Options != nil {
		req.Options.Namespace = ""
	}
	target := &TaskBook{}
	err = target.Unmarshal(out)
	if err != nil {
		return nil, err
	}
	return &GetResponse{
		TaskBook: target,
	}, nil
}

// List
func (this *TaskBooksService) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	out, err := this.store.List(ctx, req.GetOptions())
	if err != nil {
		return nil, err
	}
	if req.Options != nil {
		req.Options.Namespace = ""
	}
	list := &TaskBookList{}
	list.Items = make([]*TaskBook, len(out))
	for k := range out {
		list.Items[k] = &TaskBook{}
		err = list.Items[k].Unmarshal(out[k])
		if err != nil {
			return nil, err
		}
	}
	return &ListResponse{
		TaskBookList: list,
	}, nil
}

// Create
func (this *TaskBooksService) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	if err := req.GetTaskBook().SetDefaults(github_com_galexrt_edenconfmgmt_pkg_grpc_plugins_apiserver.MethodCreate); err != nil {
		return nil, err
	}
	if req.Options != nil {
		req.Options.Namespace = ""
	}
	req.GetTaskBook().GetMetadata().Namespace = ""
	if err := this.store.Create(ctx, req.GetTaskBook(), req.GetOptions()); err != nil {
		return nil, err
	}
	return &CreateResponse{
		TaskBook: req.GetTaskBook(),
	}, nil
}

// Update
func (this *TaskBooksService) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	if err := req.GetTaskBook().SetDefaults(github_com_galexrt_edenconfmgmt_pkg_grpc_plugins_apiserver.MethodUpdate); err != nil {
		return nil, err
	}
	if req.Options != nil {
		req.Options.Namespace = ""
	}
	req.GetTaskBook().GetMetadata().Namespace = ""
	if err := this.store.Update(ctx, req.GetTaskBook(), req.GetOptions()); err != nil {
		return nil, err
	}
	return &UpdateResponse{
		TaskBook: req.GetTaskBook(),
	}, nil
}

// Delete
func (this *TaskBooksService) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	if req.Options != nil {
		req.Options.Namespace = ""
	}
	return &DeleteResponse{}, this.store.Delete(ctx, req.GetTaskBook(), req.GetOptions())
}

// Watch
func (this *TaskBooksService) Watch(req *WatchRequest, stream TaskBooks_WatchServer) error {
	req.GetOptions().Namespace = ""
	watch, err := this.store.Watch(stream.Context(), req.GetOptions())
	if err != nil {
		return nil
	}
	for {
		select {
		case out := <-watch:
			target := &TaskBook{}
			if err = target.Unmarshal(out.Value); err != nil {
				return err
			}
			if err = stream.Send(&WatchResponse{
				TaskBook: target,
			}); err != nil {
				return err
			}
		case <-stream.Context().Done():
			return stream.Context().Err()
		}
	}
}
