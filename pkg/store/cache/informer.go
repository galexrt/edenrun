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
	"path"
	"strings"
	"sync"

	"github.com/galexrt/edenconfmgmt/pkg/store/data"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"go.uber.org/zap"
)

// Informer
type Informer struct {
	channel    map[string]*chanContainer
	receivers  map[string]*receiverList
	dataStore  data.Store
	cacheStore data.Store
	wg         sync.WaitGroup
	logger     *zap.Logger
}

type chanContainer struct {
	watch clientv3.WatchChan
	link  *chanContainer
}

type receiverList struct {
	list       []chan *InformerResult
	usageCount uint64
	sync.RWMutex
}

// NewInformer returns a new Informer.
func NewInformer(dataStore data.Store, cacheStore data.Store, logger *zap.Logger) *Informer {
	return &Informer{
		channel:    map[string]*chanContainer{},
		receivers:  map[string]*receiverList{},
		dataStore:  dataStore,
		cacheStore: cacheStore,
		logger:     logger.Named("pkg/store/cache:Informer"),
	}
}

// Start
func (inf *Informer) Start(stopCh chan struct{}) error {
	var err error
	for {
		select {
		// TODO use `inf.dataStore` to get a more "compact" watch (e.g., replace two single key watches with one directory watch)
		case <-stopCh:
			return err
		}
	}
}

// DataStoreChExists
func (inf *Informer) DataStoreChExists(key string) bool {
	if st := inf.getDataStoreCh(key); st != nil {
		return true
	}
	return false
}

func (inf *Informer) getDataStoreCh(key string) *chanContainer {
	/*
		Example `inf.channel` keys:
		* /registry/myapi/
		* /registry/otherapi/
		When asked for `key: /registry/otherapi/my-object`, the second key's
		value would get returned.
	*/
	currentPath := "/"
	components := strings.Split(key, "/")
	for _, c := range components {
		currentPath = path.Join(currentPath, c) + "/"
		if _, ok := inf.channel[currentPath]; ok {
			return inf.channel[currentPath]
		}
	}

	return nil
}

func (inf *Informer) getReceiverChs(key string) []*receiverList {
	/*
		Example `inf.receivers` keys:
		* /registry/myapi/
		* /registry/otherapi/
		When asked for `key: /registry/otherapi/my-object`, the second key's
		value would get returned.
	*/
	receivers := []*receiverList{}

	for recvPath := range inf.receivers {
		if recvPath == key {
			receivers = append(receivers, inf.receivers[recvPath])
		} else {
			components := strings.Split(key, "/")
			for k := len(components); k > 0; k-- {
				currentPath := strings.Join(components[0:k], "/")
				if recvPath == currentPath || strings.TrimRight(recvPath, "/") == currentPath {
					receivers = append(receivers, inf.receivers[recvPath])
					break
				}
			}
		}
	}

	return receivers
}

// Watch
func (inf *Informer) Watch(ctx context.Context, key string, dataStoreCh clientv3.WatchChan) (chan *InformerResult, error) {
	if _, ok := inf.channel[key]; !ok {
		inf.channel[key] = &chanContainer{
			watch: dataStoreCh,
		}
		go inf.watch(ctx, key)
	}
	if _, ok := inf.receivers[key]; !ok {
		inf.receivers[key] = &receiverList{
			list: []chan *InformerResult{},
		}
	}
	ch := make(chan *InformerResult)
	inf.receivers[key].Lock()
	inf.receivers[key].list = append(inf.receivers[key].list, ch)
	inf.receivers[key].Unlock()
	return ch, nil
}

// watch
func (inf *Informer) watch(ctx context.Context, key string) {
	dataStoreCh := inf.getDataStoreCh(key)
	for {
		select {
		case <-ctx.Done():
			return
		case resp := <-dataStoreCh.watch:
			var errs []error
			if err := resp.Err(); err != nil {
				errs = append(errs, err)
			}
			for _, event := range resp.Events {
				key := string(event.Kv.Key)
				state := convertETCDtoResultState(event)
				value := event.Kv.Value
				version := event.Kv.Version
				inf.wg.Add(1)
				go func(key string, state ResultState, value []byte) {
					switch state {
					case ResultStateCreated:
						fallthrough
					case ResultStateUpdated:
						inf.cacheStore.Put(ctx, key, value)
					case ResultStateDeleted:
						inf.cacheStore.Delete(ctx, key)
					default:
						inf.logger.Warn("got dataStore event with no or unknown ResultState", zap.String("key", key), zap.Int64("keyVersion", version), zap.Any("resultState", state))
					}
				}(key, state, value)
				inf.wg.Add(1)
				go func(key string, state ResultState, value []byte, version int64) {
					result := &InformerResult{
						Errors:  errs,
						Closed:  resp.Canceled,
						Key:     key,
						State:   state,
						Value:   value,
						Version: version,
					}
					// TODO should we lock the list here?
					receivers := inf.getReceiverChs(key)
					for _, rcvs := range receivers {
						rcvs.RLock()
						for _, rcv := range rcvs.list {
							rcv <- result
						}
						rcvs.RUnlock()
					}
				}(key, state, value, version)
			}
			if resp.Canceled {
				inf.logger.Warn("dataStore event has been canceled", zap.Error(resp.Err()))
				return
			}
		}
	}
}

func convertETCDtoResultState(event *clientv3.Event) ResultState {
	if event.Type == mvccpb.DELETE {
		return ResultStateDeleted
	}
	if event.Type == mvccpb.PUT {
		if event.Kv.CreateRevision == event.Kv.Version {
			return ResultStateUpdated
		}
		return ResultStateCreated
	}
	return ResultStateUnknown
}
