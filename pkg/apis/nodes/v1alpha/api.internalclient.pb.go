// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/galexrt/edenconfmgmt/pkg/apis/nodes/v1alpha/api.proto

package v1alpha

import (
	fmt "fmt"
	_ "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha"
	github_com_galexrt_edenconfmgmt_pkg_apis_core_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha"
	_ "github.com/galexrt/edenconfmgmt/pkg/apis/events/v1alpha"
	_ "github.com/galexrt/edenconfmgmt/pkg/grpc/plugins/internalclient"
	github_com_galexrt_edenconfmgmt_pkg_store_cache "github.com/galexrt/edenconfmgmt/pkg/store/cache"
	github_com_galexrt_edenconfmgmt_pkg_utils_api "github.com/galexrt/edenconfmgmt/pkg/utils/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Client !!WARNING!! This client must only be used by internal code!
// All "external" clients must use the actual GRPC API Client.
type Client struct {
	store *github_com_galexrt_edenconfmgmt_pkg_store_cache.Store
}

// NewClient returns a new Client to be used with the given storage.
func NewClient(store *github_com_galexrt_edenconfmgmt_pkg_store_cache.Store) *Client {
	return &Client{
		store: store,
	}
}

// Get
func (this *Client) Get(ctx context.Context, opts *github_com_galexrt_edenconfmgmt_pkg_apis_core_v1alpha.GetOptions) (*Node, error) {
	out, err := this.store.Get(context.Background(), github_com_galexrt_edenconfmgmt_pkg_utils_api.ObjectPath(APIPath, opts.Name))
	if err != nil {
		return nil, nil
	}
	node := &Node{}
	err = proto.Unmarshal(out, node)
	return node, err
}

// List
func (this *Client) List(name string) error {
	return nil
}

// Create
func (this *Client) Create(name string) error {
	return nil
}

// Update
func (this *Client) Update(name string) error {
	return nil
}

// Delete
func (this *Client) Delete(name string) error {
	return nil
}