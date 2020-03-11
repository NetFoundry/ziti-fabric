/*
	Copyright 2019 NetFoundry, Inc.

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

package models

import (
	"github.com/netfoundry/ziti-foundation/storage/ast"
	"github.com/netfoundry/ziti-foundation/storage/boltz"
	"go.etcd.io/bbolt"
	"time"
)

const (
	ListLimitMax      = 500
	ListOffsetMax     = 100000
	ListLimitDefault  = 10
	ListOffsetDefault = 0
)

type EntityLister interface {
	BaseList(query string) (*EntityListResult, error)
	GetStore() boltz.CrudStore
}

type EntityLoader interface {
	BaseLoad(id string) (Entity, error)
	BaseLoadInTx(tx *bbolt.Tx, id string) (Entity, error)
}

type Entity interface {
	GetId() string
	SetId(string)
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetTags() map[string]interface{}
}

type BaseEntity struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Tags      map[string]interface{}
}

func (entity *BaseEntity) GetId() string {
	return entity.Id
}

func (entity *BaseEntity) SetId(id string) {
	entity.Id = id
}

func (entity *BaseEntity) GetCreatedAt() time.Time {
	return entity.CreatedAt
}

func (entity *BaseEntity) GetUpdatedAt() time.Time {
	return entity.UpdatedAt
}

func (entity *BaseEntity) GetTags() map[string]interface{} {
	return entity.Tags
}

func (entity *BaseEntity) FillCommon(boltEntity boltz.ExtEntity) {
	entity.Id = boltEntity.GetId()
	entity.CreatedAt = boltEntity.GetCreatedAt()
	entity.UpdatedAt = boltEntity.GetUpdatedAt()
	entity.Tags = boltEntity.GetTags()
}

type EntityListResult struct {
	Loader   EntityLoader
	Entities []Entity
	QueryMetaData
}

func (result *EntityListResult) GetEntities() []Entity {
	return result.Entities
}

func (result *EntityListResult) GetMetaData() *QueryMetaData {
	return &result.QueryMetaData
}

func (result *EntityListResult) Collect(tx *bbolt.Tx, ids []string, queryMetaData *QueryMetaData) error {
	result.QueryMetaData = *queryMetaData
	for _, key := range ids {
		entity, err := result.Loader.BaseLoadInTx(tx, key)
		if err != nil {
			return err
		}
		result.Entities = append(result.Entities, entity)
	}
	return nil
}

type QueryMetaData struct {
	Count            int64
	Limit            int64
	Offset           int64
	FilterableFields []string
}

type BaseController struct {
	Store boltz.CrudStore
}

func (ctrl *BaseController) GetStore() boltz.CrudStore {
	return ctrl.Store
}

type ListResultHandler func(tx *bbolt.Tx, ids []string, qmd *QueryMetaData) error

func (ctrl *BaseController) checkLimits(query ast.Query) {
	if query.GetLimit() == nil || *query.GetLimit() < -1 || *query.GetLimit() == 0 {
		query.SetLimit(ListLimitDefault)
	} else if *query.GetLimit() > ListLimitMax {
		query.SetLimit(ListLimitMax)
	}

	if query.GetSkip() == nil || *query.GetSkip() < 0 {
		query.SetSkip(ListOffsetDefault)
	} else if *query.GetSkip() > ListOffsetMax {
		query.SetSkip(ListOffsetMax)
	}
}

func (ctrl *BaseController) ListWithTx(tx *bbolt.Tx, queryString string, resultHandler ListResultHandler) error {
	query, err := ast.Parse(ctrl.Store, queryString)
	if err != nil {
		return err
	}

	ctrl.checkLimits(query)

	keys, count, err := ctrl.Store.QueryIdsC(tx, query)
	if err != nil {
		return err
	}
	qmd := &QueryMetaData{
		Count:            count,
		Limit:            *query.GetLimit(),
		Offset:           *query.GetSkip(),
		FilterableFields: ctrl.Store.GetPublicSymbols(),
	}
	return resultHandler(tx, keys, qmd)
}
