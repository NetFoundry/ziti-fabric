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

package network

import (
	"github.com/netfoundry/ziti-foundation/identity/identity"
	"github.com/netfoundry/ziti-foundation/util/sequence"
	"github.com/orcaman/concurrent-map"
	"math"
)

type linkController struct {
	linkTable      *linkTable
	adjacencyTable *adjacencyTable
	sequence       *sequence.Sequence
}

func newLinkController() *linkController {
	return &linkController{
		linkTable:      newLinkTable(),
		adjacencyTable: newAdjacencyTable(),
		sequence:       sequence.NewSequence(),
	}
}

func (linkController *linkController) add(link *Link) {
	linkController.linkTable.add(link)
	linkController.adjacencyTable.add(link)
}

func (linkController *linkController) has(link *Link) bool {
	return linkController.linkTable.has(link)
}

func (linkController *linkController) get(linkId *identity.TokenId) (*Link, bool) {
	link, found := linkController.linkTable.get(linkId)
	return link, found
}

func (linkController *linkController) all() []*Link {
	return linkController.linkTable.all()
}

func (linkController *linkController) remove(link *Link) {
	linkController.linkTable.remove(link)
	linkController.adjacencyTable.remove(link)
}

func (linkController *linkController) allLinksForRouter(routerId string) []*Link {
	if rlt, found := linkController.adjacencyTable.get(routerId); found {
		return rlt.allLinksForAllRouters()
	}
	return nil
}

func (linkController *linkController) connectedNeighborsOfRouter(router *Router) []*Router {
	neighborMap := make(map[string]*Router)

	links := linkController.allLinksForRouter(router.Id)
	for _, link := range links {
		currentState := link.CurrentState()
		if currentState != nil && currentState.Mode == Connected && !link.Down {
			if link.Src != router {
				neighborMap[link.Src.Id] = link.Src
			}
			if link.Dst != router {
				neighborMap[link.Dst.Id] = link.Dst
			}
		}
	}

	neighbors := make([]*Router, 0)
	for _, r := range neighborMap {
		neighbors = append(neighbors, r)
	}
	return neighbors
}

func (linkController *linkController) leastExpensiveLink(a, b *Router) (*Link, bool) {
	var selected *Link
	var cost int64 = math.MaxInt64

	links := linkController.allLinksForRouter(a.Id)
	for _, link := range links {
		currentState := link.CurrentState()
		if currentState != nil && currentState.Mode == Connected && !link.Down {
			if link.Src == a && link.Dst == b {
				if link.SrcLatency+link.DstLatency < cost {
					selected = link
				}
			}
			if link.Dst == a && link.Src == b {
				if link.SrcLatency+link.DstLatency < cost {
					selected = link
				}
			}
		}
	}

	if selected != nil {
		return selected, true
	}

	return nil, false
}

func (linkController *linkController) missingLinks(routers []*Router) ([]*Link, error) {
	missingLinks := make([]*Link, 0)
	for _, srcR := range routers {
		for _, dstR := range routers {
			if srcR != dstR && dstR.AdvertisedListener != "" {
				if _, found := linkController.firstDirectedLink(srcR, dstR); !found {
					id, err := linkController.sequence.NextHash()
					if err != nil {
						return nil, err
					}
					link := newLink(&identity.TokenId{Token: id})
					link.Src = srcR
					link.Dst = dstR
					missingLinks = append(missingLinks, link)
				}
			}
		}
	}

	return missingLinks, nil
}

func (linkController *linkController) firstDirectedLink(a, b *Router) (*Link, bool) {
	// a->b
	if rlt, found := linkController.adjacencyTable.get(a.Id); found {
		if links, found := rlt.allLinksForRouter(b.Id); found {
			for _, link := range links {
				if link.Src == a && link.Dst == b && link.CurrentState().Mode == Connected {
					return link, true
				}
			}
		}
	}
	// b->a
	if rlt, found := linkController.adjacencyTable.get(b.Id); found {
		if links, found := rlt.allLinksForRouter(a.Id); found {
			for _, link := range links {
				if link.Src == b && link.Dst == a && link.CurrentState().Mode == Connected {
					return link, true
				}
			}
		}
	}
	return nil, false
}

func (linkController *linkController) linksInMode(mode LinkMode) []*Link {
	return linkController.linkTable.allInMode(mode)
}

/*
 * linkTable
 */

type linkTable struct {
	links cmap.ConcurrentMap // map[Link.Id.Token]*Link
}

func newLinkTable() *linkTable {
	return &linkTable{links: cmap.New()}
}

func (lt *linkTable) add(link *Link) {
	lt.links.Set(link.Id.Token, link)
}

func (lt *linkTable) get(linkId *identity.TokenId) (*Link, bool) {
	link, found := lt.links.Get(linkId.Token)
	if link != nil {
		return link.(*Link), found
	}
	return nil, found
}

func (lt *linkTable) has(link *Link) bool {
	if i, found := lt.links.Get(link.Id.Token); found {
		if i.(*Link) == link {
			return true
		}
	}
	return false
}

func (lt *linkTable) all() []*Link {
	links := make([]*Link, 0)
	for i := range lt.links.IterBuffered() {
		links = append(links, i.Val.(*Link))
	}
	return links
}

func (lt *linkTable) allInMode(mode LinkMode) []*Link {
	links := make([]*Link, 0)
	for i := range lt.links.IterBuffered() {
		link := i.Val.(*Link)
		if link.CurrentState().Mode == mode {
			links = append(links, link)
		}
	}
	return links
}

func (lt *linkTable) remove(link *Link) {
	lt.links.Remove(link.Id.Token)
}

/*
 * adjacencyTable
 */

type adjacencyTable struct {
	adjacency cmap.ConcurrentMap // map[Router.Id.Token]*routerLinksTable
}

func newAdjacencyTable() *adjacencyTable {
	return &adjacencyTable{adjacency: cmap.New()}
}

func (at *adjacencyTable) add(link *Link) {
	// src->dst
	var rlt *routerLinksTable
	if i, found := at.adjacency.Get(link.Src.Id); found {
		rlt = i.(*routerLinksTable)
	} else {
		rlt = newRouterLinksTable()
	}
	rlt.addLinkForRouter(link.Dst.Id, link)
	at.adjacency.Set(link.Src.Id, rlt)

	// dst->src
	if i, found := at.adjacency.Get(link.Dst.Id); found {
		rlt = i.(*routerLinksTable)
	} else {
		rlt = newRouterLinksTable()
	}
	rlt.addLinkForRouter(link.Src.Id, link)
	at.adjacency.Set(link.Dst.Id, rlt)
}

func (at *adjacencyTable) get(routerId string) (*routerLinksTable, bool) {
	i, found := at.adjacency.Get(routerId)
	if i != nil {
		return i.(*routerLinksTable), found
	}
	return nil, found
}

func (at *adjacencyTable) remove(link *Link) {
	// src->dst
	if i, found := at.adjacency.Get(link.Src.Id); found {
		rlt := i.(*routerLinksTable)
		rlt.removeLinkFromRouter(link.Dst.Id, link)
		if rlt.size() > 0 {
			at.adjacency.Set(link.Src.Id, rlt)
		} else {
			at.adjacency.Remove(link.Src.Id)
		}
	}

	// dst->src
	if i, found := at.adjacency.Get(link.Dst.Id); found {
		rlt := i.(*routerLinksTable)
		rlt.removeLinkFromRouter(link.Src.Id, link)
		if rlt.size() > 0 {
			at.adjacency.Set(link.Dst.Id, rlt)
		} else {
			at.adjacency.Remove(link.Dst.Id)
		}
	}
}

/*
 * routerLinksTable
 */

type routerLinksTable struct {
	routerLinks cmap.ConcurrentMap // map[Router.Id.Token][]*Link
}

func newRouterLinksTable() *routerLinksTable {
	return &routerLinksTable{routerLinks: cmap.New()}
}

func (rlt *routerLinksTable) addLinkForRouter(routerId string, link *Link) {
	var links []*Link
	if i, found := rlt.routerLinks.Get(routerId); found {
		links = i.([]*Link)
	} else {
		links = make([]*Link, 0)
	}

	links = append(links, link)
	rlt.routerLinks.Set(routerId, links)
}

func (rlt *routerLinksTable) allLinksForRouter(routerId string) ([]*Link, bool) {
	if i, found := rlt.routerLinks.Get(routerId); found {
		return i.([]*Link), true
	}
	return nil, false
}

func (rlt *routerLinksTable) allLinksForAllRouters() []*Link {
	linksMap := make(map[string]*Link)
	for i := range rlt.routerLinks.IterBuffered() {
		links := i.Val.([]*Link)
		for _, link := range links {
			linksMap[link.Id.Token] = link
		}
	}

	allLinks := make([]*Link, 0)
	for _, v := range linksMap {
		allLinks = append(allLinks, v)
	}

	return allLinks
}

func (rlt *routerLinksTable) size() int {
	return rlt.routerLinks.Count()
}

func (rlt *routerLinksTable) removeLinkFromRouter(routerId string, link *Link) {
	var links []*Link
	if i, found := rlt.routerLinks.Get(routerId); found {
		links = i.([]*Link)
		if len(links) == 1 && links[0] == link {
			rlt.routerLinks.Remove(routerId)

		} else {
			i := -1
			for j, jLink := range links {
				if jLink == link {
					i = j
					break
				}
			}
			if i != -1 {
				links = append(links[:i], links[i+1:]...)
			}
			rlt.routerLinks.Set(routerId, links)
		}
	}
}
