/*
	Copyright 2020 NetFoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package db

import (
	"github.com/netfoundry/ziti-foundation/storage/boltz"
	"reflect"
)

type Stores struct {
	Endpoint EndpointStore
	Router   RouterStore
	Service  ServiceStore
	storeMap map[string]boltz.CrudStore
}

func (stores *Stores) buildStoreMap() {
	stores.storeMap = map[string]boltz.CrudStore{}
	val := reflect.ValueOf(stores).Elem()
	for i := 0; i < val.NumField(); i++ {
		f := val.Field(i)
		if f.CanInterface() {
			if store, ok := f.Interface().(boltz.CrudStore); ok {
				stores.storeMap[store.GetEntityType()] = store
			}
		}
	}
}

func (stores *Stores) GetStoreList() []boltz.CrudStore {
	var result []boltz.CrudStore
	for _, store := range stores.storeMap {
		result = append(result, store)
	}
	return result
}

func (stores *Stores) GetStoreForEntity(entity boltz.Entity) boltz.CrudStore {
	return stores.storeMap[entity.GetEntityType()]
}

type stores struct {
	endpoint *endpointStoreImpl
	router   *routerStoreImpl
	service  *serviceStoreImpl
}

func InitStores(db boltz.Db) (*Stores, error) {
	internalStores := &stores{}

	internalStores.endpoint = newEndpointStore(internalStores)
	internalStores.router = newRouterStore(internalStores)
	internalStores.service = newServiceStore(internalStores)

	stores := &Stores{
		Endpoint: internalStores.endpoint,
		Router:   internalStores.router,
		Service:  internalStores.service,
	}

	stores.buildStoreMap()

	internalStores.endpoint.initializeLinked()
	internalStores.router.initializeLinked()
	internalStores.service.initializeLinked()

	mm := boltz.NewMigratorManager(db)
	if err := mm.Migrate("fabric", CurrentDbVersion, internalStores.migrate); err != nil {
		return nil, err
	}

	return stores, nil
}