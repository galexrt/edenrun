/*
Copyright 2019 Alexander Trost. All rights reserved.

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

package cache

import (
	"context"
	"fmt"
	"sync"

	"github.com/galexrt/edenconfmgmt/pkg/store/data"
	"github.com/galexrt/edenconfmgmt/pkg/utils/errors"
	"go.uber.org/zap"
)

// Store using the dataStore and a second cacheStore for caching.
type Store struct {
	dataStore  data.Store
	cacheStore data.Store
	informer   *Informer
	prefix     string
	wg         sync.WaitGroup
	logger     *zap.Logger
}

// TODO Rewrite to save objects protobuf encoded as this will definitly be faster than to always parse JSON again.

// New return new Store.
func New(dataStore data.Store, cacheStore data.Store, logger *zap.Logger) *Store {
	inf := NewInformer(dataStore, cacheStore, logger)
	return &Store{
		dataStore:  dataStore,
		cacheStore: cacheStore,
		informer:   inf,
		logger:     logger.Named("pkg/store/cache:Store"),
	}
}

// Start start the logic behind the cache store.
func (st *Store) Start(stopCh chan struct{}) error {
	var err error

	st.wg.Add(1)
	go func() {
		defer st.wg.Done()
		st.informer.Start(stopCh)
	}()

	for {
		select {
		case <-stopCh:
			st.wg.Wait()
			return err
		}
	}
}

// SetKeyPrefix set the prefix to prefix all given keys with.
func (st *Store) SetKeyPrefix(prefix string) {
	st.dataStore.SetKeyPrefix(prefix)
	st.cacheStore.SetKeyPrefix(prefix)
	st.prefix = prefix
}

// Get get a value for a key.
func (st *Store) Get(ctx context.Context, key string) ([]byte, error) {
	result, err := st.cacheStore.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		result, err = st.dataStore.Get(ctx, key)
		if err != nil {
			return result, err
		}
	}
	err = st.cacheStore.Put(ctx, key, result)
	return result, err
}

// GetRecursive return a list of keys with values at the given key arg.
func (st *Store) GetRecursive(ctx context.Context, key string) (map[string][]byte, error) {
	results, err := st.cacheStore.GetRecursive(ctx, key)
	if err != nil {
		return results, err
	}
	if len(results) == 0 {
		results, err = st.dataStore.GetRecursive(ctx, key)
		if err != nil {
			return results, err
		}
	}

	if len(results) == 0 {
		return results, err
	}

	var errs []error
	for k, v := range results {
		if err = st.cacheStore.Put(ctx, k, v); err != nil {
			// TODO Handle error (e.g., build concated error)
			errs = append(errs, err)
		}
	}
	return results, errors.Concat(errs...)
}

// Put put a key value pair.
func (st *Store) Put(ctx context.Context, key string, value []byte) error {
	if err := st.dataStore.Put(ctx, key, value); err != nil {
		return err
	}
	return st.cacheStore.Put(ctx, key, value)
}

// Delete delete a key value pair.
func (st *Store) Delete(ctx context.Context, key string) error {
	if err := st.dataStore.Delete(ctx, key); err != nil {
		return err
	}
	return st.cacheStore.Delete(ctx, key)
}

// Watch watch a key or directory for creation, changes and deletion.
func (st *Store) Watch(ctx context.Context, key string) (chan *InformerResult, error) {
	// TODO call st.informer.DataStoreChExists() to see if a datastore watchchan is needed.
	watch, err := st.dataStore.Watch(ctx, key)
	if err != nil {
		return nil, err
	}
	fmt.Printf("store.Watch key: %s\n", key)
	return st.informer.Watch(ctx, key, watch)
}

// Close adapter.
func (st *Store) Close() error {
	return errors.Concat(st.dataStore.Close(), st.cacheStore.Close())
}