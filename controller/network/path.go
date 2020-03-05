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

package network

import (
	"errors"
	"math"
)

func (network *Network) shortestPath(srcR *Router, dstR *Router) ([]*Router, error) {
	if srcR == nil || dstR == nil {
		return nil, errors.New("not routable (!srcR||!dstR)")
	}

	if srcR == dstR {
		return []*Router{srcR}, nil
	}

	dist := make(map[*Router]int64)
	prev := make(map[*Router]*Router)
	unvisited := make(map[*Router]bool)

	for _, r := range network.routerController.allConnected() {
		dist[r] = math.MaxInt32
		unvisited[r] = true
	}
	dist[srcR] = 0

	for len(unvisited) > 0 {
		u := minCost(unvisited, dist)
		delete(unvisited, u)

		neighbors := network.linkController.connectedNeighborsOfRouter(u)
		for _, r := range neighbors {
			if _, found := unvisited[r]; found {
				cost := int64(r.CostFactor)
				if l, found := network.linkController.leastExpensiveLink(r, u); found {
					cost += l.SrcLatency
					cost += l.DstLatency
				}

				alt := dist[u] + cost
				if alt < dist[r] {
					dist[r] = alt
					prev[r] = u
				}
			}
		}
	}

	/*
	 * dist: (r2->r1->r0)
	 *		r0 = 2 <- r1
	 *		r1 = 1 <- r2
	 *		r2 = 0 <- nil
	 */

	routerPath := make([]*Router, 0)
	p := prev[dstR]
	for p != nil {
		routerPath = append([]*Router{p}, routerPath...)
		p = prev[p]
	}
	routerPath = append(routerPath, dstR)

	if routerPath[0] != srcR {
		return nil, errors.New("not routable (~srcR)")
	}
	if routerPath[len(routerPath)-1] != dstR {
		return nil, errors.New("not routable (~dstR)")
	}

	return routerPath, nil
}

func minCost(q map[*Router]bool, dist map[*Router]int64) *Router {
	if dist == nil || len(dist) < 1 {
		return nil
	}

	min := int64(math.MaxInt64)
	var selected *Router
	for r := range q {
		d := dist[r]
		if d <= min {
			selected = r
			min = d
		}
	}
	return selected
}
