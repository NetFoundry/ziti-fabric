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
	"fmt"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/channel/protobufs"
	"github.com/openziti/fabric/controller/xt"
	"github.com/openziti/fabric/logcontext"
	"github.com/openziti/fabric/pb/ctrl_pb"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"time"
)

type routeSenderController struct {
	senders cmap.ConcurrentMap // map[string]*routeSender
}

func newRouteSenderController() *routeSenderController {
	return &routeSenderController{senders: cmap.New()}
}

func (self *routeSenderController) forwardRouteResult(rs *RouteStatus) bool {
	v, found := self.senders.Get(rs.CircuitId)
	if found {
		routeSender := v.(*routeSender)
		routeSender.in <- rs
		return true
	}
	logrus.Warnf("did not find route sender for [s/%s]", rs.CircuitId)
	return false
}

func (self *routeSenderController) addRouteSender(rs *routeSender) {
	self.senders.Set(rs.circuitId, rs)
}

func (self *routeSenderController) removeRouteSender(rs *routeSender) {
	self.senders.Remove(rs.circuitId)
}

type routeSender struct {
	circuitId       string
	path            *Path
	routeMsgs       []*ctrl_pb.Route
	timeout         time.Duration
	in              chan *RouteStatus
	attendance      map[string]bool
	serviceCounters ServiceCounters
}

func newRouteSender(circuitId string, timeout time.Duration, serviceCounters ServiceCounters) *routeSender {
	return &routeSender{
		circuitId:       circuitId,
		timeout:         timeout,
		in:              make(chan *RouteStatus, 16),
		attendance:      make(map[string]bool),
		serviceCounters: serviceCounters,
	}
}

func (self *routeSender) route(attempt uint32, path *Path, routeMsgs []*ctrl_pb.Route, strategy xt.Strategy, terminator xt.Terminator, ctx logcontext.Context) (peerData xt.PeerData, cleanups map[string]struct{}, err error) {
	logger := pfxlog.ChannelLogger(logcontext.EstablishPath).Wire(ctx)

	// send route messages
	tr := path.Nodes[len(path.Nodes)-1]
	for i := 0; i < len(path.Nodes); i++ {
		r := path.Nodes[i]
		msg := routeMsgs[i]
		logger.Debugf("sending route message to [r/%s] for attempt [#%d]", r.Id, msg.Attempt)
		go self.sendRoute(r, msg, ctx)
		self.attendance[r.Id] = false
	}

	deadline := time.Now().Add(self.timeout)
	timeout := time.Until(deadline)
attendance:
	for {
		select {
		case status := <-self.in:
			if status.Success {
				if status.Attempt == attempt {
					logger.Debugf("received successful route status from [r/%s] for attempt [#%d] of [s/%s]", status.Router.Id, status.Attempt, status.CircuitId)

					self.attendance[status.Router.Id] = true
					if status.Router == tr {
						peerData = status.PeerData
						strategy.NotifyEvent(xt.NewDialSucceeded(terminator))
						self.serviceCounters.ServiceDialSuccess(terminator.GetServiceId(), terminator.GetId())
					}
				} else {
					logger.Warnf("received successful route status from [r/%s] for alien attempt [#%d (not #%d)] of [s/%s]", status.Router.Id, status.Attempt, attempt, status.CircuitId)
				}

			} else {
				if status.Attempt == attempt {
					logger.Warnf("received failed route status from [r/%s] for attempt [#%d] of [s/%s] (%v)", status.Router.Id, status.Attempt, status.CircuitId, status.Err)

					if status.Router == tr {
						strategy.NotifyEvent(xt.NewDialFailedEvent(terminator))
						self.serviceCounters.ServiceDialFail(terminator.GetServiceId(), terminator.GetId())
					}
					cleanups = self.cleanups(path)

					return nil, cleanups, errors.Errorf("error creating route for [s/%s] on [r/%s] (%v)", self.circuitId, status.Router.Id, status.Err)
				} else {
					logger.Warnf("received failed route status from [r/%s] for alien attempt [#%d (not #%d)] of [s/%s]", status.Router.Id, status.Attempt, attempt, status.CircuitId)
				}
			}

		case <-time.After(timeout):
			cleanups = self.cleanups(path)
			strategy.NotifyEvent(xt.NewDialFailedEvent(terminator))
			self.serviceCounters.ServiceDialTimeout(terminator.GetServiceId(), terminator.GetId())
			return nil, cleanups, &routeTimeoutError{circuitId: self.circuitId}
		}

		allPresent := true
		for _, v := range self.attendance {
			if !v {
				allPresent = false
			}
		}
		if allPresent {
			break attendance
		}

		timeout = time.Until(deadline)
	}

	return peerData, nil, nil
}

func (self *routeSender) sendRoute(r *Router, routeMsg *ctrl_pb.Route, ctx logcontext.Context) {
	logger := pfxlog.ChannelLogger(logcontext.EstablishPath).Wire(ctx).WithField("routerId", r.Id)

	envelope := protobufs.MarshalTyped(routeMsg).WithTimeout(3 * time.Second)
	if err := envelope.SendAndWaitForWire(r.Control); err != nil {
		logger.WithError(err).Error("failure sending route message")
	} else {
		logger.Debug("sent route message")
	}
}

func (self *routeSender) cleanups(path *Path) map[string]struct{} {
	cleanups := make(map[string]struct{})
	for _, r := range path.Nodes {
		success, found := self.attendance[r.Id]
		if found && success {
			cleanups[r.Id] = struct{}{}
		}
	}
	return cleanups
}

type RouteStatus struct {
	Router    *Router
	CircuitId string
	Attempt   uint32
	Success   bool
	Err       string
	PeerData  xt.PeerData
	ErrorCode byte
}

type routeTimeoutError struct {
	circuitId string
}

func (self routeTimeoutError) Error() string {
	return fmt.Sprintf("timeout creating routes for [s/%s]", self.circuitId)
}
