/*
Copyright 2018 Alexander Trost. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package store

import (
	"context"
)

// TODO Write comments for states of a Response (e.g. Deleted, Updated, Created).
const (
	ResponseStateCreated = "Created"
	ResponseStateUpdated = "Updated"
	ResponseStateDeleted = "Deleted"
)

// Response response of a watch event.
type Response struct {
	Error error
	Key   string
	Value string
	State string
}

// Watcher interface.
type Watcher interface {
	// Watch return channel which "streams" events as they come.
	Watch(ctx context.Context) (chan *Response, error)
}
