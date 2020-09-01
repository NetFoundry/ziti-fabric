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
	"github.com/openziti/fabric/controller/db"
	"github.com/openziti/foundation/identity/identity"
	"github.com/openziti/foundation/transport/tcp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortestPath(t *testing.T) {
	ctx := db.NewTestContext(t)
	defer ctx.Cleanup()

	req := assert.New(t)

	nodeId := &identity.TokenId{Token: "test"}
	network, err := NewNetwork(nodeId, nil, ctx.GetDb(), nil)
	req.NoError(err)

	addr := "tcp:0.0.0.0:0"
	transportAddr, err := tcp.AddressParser{}.Parse(addr)
	req.NoError(err)

	r0 := newRouterForTest("r0", "", transportAddr, nil)
	network.Routers.markConnected(r0)

	r1 := newRouterForTest("r1", "", transportAddr, nil)
	network.Routers.markConnected(r1)

	r2 := newRouterForTest("r2", "", transportAddr, nil)
	network.Routers.markConnected(r2)

	r3 := newRouterForTest("r3", "", transportAddr, nil)
	network.Routers.markConnected(r3)

	link := newLink(&identity.TokenId{Token: "l0"})
	link.Cost = 2
	link.DstLatency = 10
	link.SrcLatency = 11
	link.Src = r0
	link.Dst = r1
	link.addState(newLinkState(Connected))
	network.linkController.add(link)

	link = newLink(&identity.TokenId{Token: "l1"})
	link.Cost = 5
	link.DstLatency = 15
	link.SrcLatency = 16
	link.Src = r0
	link.Dst = r2
	link.addState(newLinkState(Connected))
	network.linkController.add(link)

	link = newLink(&identity.TokenId{Token: "l2"})
	link.Cost = 9
	link.DstLatency = 20
	link.SrcLatency = 21
	link.Src = r1
	link.Dst = r3
	link.addState(newLinkState(Connected))
	network.linkController.add(link)

	link = newLink(&identity.TokenId{Token: "l3"})
	link.Cost = 13
	link.DstLatency = 25
	link.SrcLatency = 26
	link.Src = r2
	link.Dst = r3
	link.addState(newLinkState(Connected))
	network.linkController.add(link)

	path, cost, err := network.shortestPath(r0, r3)
	req.NoError(err)
	req.NotNil(t, path)
	req.Equal(path[0], r0)
	req.Equal(path[1], r1)
	req.Equal(path[2], r3)

	expected := 1 + 1 + // start and end router cost
		10 + 11 + 2 + // link1 cost and src and dest latency
		9 + 20 + 21 // link2 cost and src and dest latency
	req.Equal(int64(expected), cost)
}
