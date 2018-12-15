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

package v1alpha

import (
	api "github.com/galexrt/edenconfmgmt/pkg/apis/triggers"
)

const (
	// Kind the name of the object kind (singular).
	Kind = "Trigger"
	// KindPlural the name of the object kind (plural).
	KindPlural = "Triggers"
	// APIVersion of the API.
	APIVersion = "v1alpha"
	// DataStorePath API object path in the data store.
	DataStorePath = api.APIName + "/" + APIVersion
)