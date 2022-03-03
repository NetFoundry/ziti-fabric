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

package handler_mgmt

import (
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/channel"
	"github.com/openziti/fabric/controller/network"
	"github.com/openziti/fabric/controller/xmgmt"
)

type MgmtAccepter struct {
	listener channel.UnderlayListener
	options  *channel.Options
	network  *network.Network
	xmgmts   []xmgmt.Xmgmt
}

func NewMgmtAccepter(network *network.Network,
	xmgmts []xmgmt.Xmgmt,
	listener channel.UnderlayListener,
	options *channel.Options) *MgmtAccepter {
	return &MgmtAccepter{
		network:  network,
		xmgmts:   xmgmts,
		listener: listener,
		options:  options,
	}
}

func (self *MgmtAccepter) Run() {
	log := pfxlog.Logger()
	log.Info("started")
	defer log.Warn("exited")

	bindHandler := NewBindHandler(self.network, self.xmgmts)

	for {
		ch, err := channel.NewChannel("mgmt", self.listener, bindHandler, self.options)
		if err == nil {
			log.Debugf("accepted mgmt connection [%s]", ch.Id().Token)

		} else {
			log.Errorf("error accepting (%s)", err)
			if err.Error() == "closed" {
				return
			}
		}
	}
}
