/*
	Copyright NetFoundry, Inc.

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

package xt

import (
	cmap "github.com/orcaman/concurrent-map"
	"math"
)

const (
	failedIntervalStart   = 10_000
	defaultIntervalStart  = 5000
	requiredIntervalStart = 0
)

var globalCosts = &costs{
	costMap: cmap.New(),
}

func GlobalCosts() Costs {
	return globalCosts
}

type precedence struct {
	costIntervalStart uint16
}

func (p precedence) Unbias(cost uint16) uint16 {
	return cost - p.costIntervalStart
}

func (p precedence) IsFailed() bool {
	return p.costIntervalStart == failedIntervalStart
}

func (p precedence) IsDefault() bool {
	return p.costIntervalStart == defaultIntervalStart
}

func (p precedence) IsRequired() bool {
	return p.costIntervalStart == requiredIntervalStart
}

func (p precedence) getCostIntervalStart() uint16 {
	return p.costIntervalStart
}

// Precedences define the precedence levels
var Precedences = struct {
	// Required terminators should always be used in preference to non-required terminators
	// Example: The current active node in an active/passive setup. If the strategy determintes
	//          that the current active has failed, it will drop the precedence and mark some other
	//          terminator as Required
	Required Precedence

	// Default precedence is where terminators start
	Default Precedence

	// Failed means this terminator should only be used if all other terminators are also in a failed state
	// Example: A strategy might move a terminator to Failed if three dials in a row fail
	Failed Precedence
}{
	Required: precedence{costIntervalStart: requiredIntervalStart},
	Default:  precedence{costIntervalStart: defaultIntervalStart},
	Failed:   precedence{costIntervalStart: failedIntervalStart},
}

type terminatorStats struct {
	cost           uint16
	precedence     Precedence
	precedenceCost uint8
}

func (stats *terminatorStats) GetCost() uint16 {
	return stats.cost
}

func (stats *terminatorStats) GetPrecedence() Precedence {
	return stats.precedence
}

type costs struct {
	costMap cmap.ConcurrentMap
}

func (self *costs) ClearCost(terminatorId string) {
	self.costMap.Remove(terminatorId)
}

func (self *costs) GetCost(terminatorId string) uint16 {
	stats := self.getStats(terminatorId)
	if stats == nil {
		return 0
	}
	return stats.cost
}

func (self *costs) GetStats(terminatorId string) Stats {
	stats := self.getStats(terminatorId)
	if stats == nil {
		return &terminatorStats{
			cost:           Precedences.Default.getCostIntervalStart(),
			precedence:     Precedences.Default,
			precedenceCost: 0,
		}
	}
	return stats
}

func (self *costs) getStats(terminatorId string) *terminatorStats {
	val, found := self.costMap.Get(terminatorId)
	if !found {
		return nil
	}
	return val.(*terminatorStats)
}

func (self *costs) GetPrecedence(terminatorId string) Precedence {
	stats := self.getStats(terminatorId)
	if stats == nil {
		return Precedences.Default
	}
	return stats.precedence
}

func (self *costs) SetPrecedence(terminatorId string, precedence Precedence) {
	stats := self.getStats(terminatorId)
	var precedenceCost uint8
	if stats != nil {
		precedenceCost = stats.precedenceCost
	}
	self.costMap.Set(terminatorId, self.newStats(precedence, precedenceCost))
}

func (self *costs) SetPrecedenceCost(terminatorId string, cost uint8) {
	stats := self.getStats(terminatorId)
	var p Precedence
	if stats == nil {
		p = Precedences.Default
	} else {
		p = stats.precedence
	}
	self.costMap.Set(terminatorId, self.newStats(p, cost))
}

func (self *costs) GetPrecedenceCost(terminatorId string) uint8 {
	stats := self.getStats(terminatorId)
	if stats == nil {
		return 0
	}
	return stats.precedenceCost
}

func (self *costs) newStats(precedence Precedence, cost uint8) *terminatorStats {
	return &terminatorStats{
		cost:           self.calculateCost(precedence, cost),
		precedence:     precedence,
		precedenceCost: cost,
	}
}

func (self *costs) calculateCost(precedence Precedence, cost uint8) uint16 {
	nextCost := uint32(precedence.getCostIntervalStart()) + uint32(cost)
	if nextCost > math.MaxUint16 {
		return math.MaxUint16
	}
	return uint16(nextCost)
}

// In a list which is sorted by precedence, returns the terminators which have the
// same precedence as that of the first entry in the list
func GetRelatedTerminators(list []CostedTerminator) []CostedTerminator {
	first := list[0]
	var result = []CostedTerminator{first}
	for _, t := range list[1:] {
		if t.GetPrecedence() == first.GetPrecedence() {
			result = append(result, t)
		} else {
			break
		}
	}
	return result
}
