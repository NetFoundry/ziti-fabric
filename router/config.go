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

package router

import (
	"bytes"
	"fmt"
	"github.com/michaelquigley/pfxlog"
	"github.com/netfoundry/ziti-fabric/pb/ctrl_pb"
	"github.com/netfoundry/ziti-fabric/router/forwarder"
	"github.com/netfoundry/ziti-fabric/router/xgress"
	"github.com/netfoundry/ziti-foundation/channel2"
	"github.com/netfoundry/ziti-foundation/config"
	"github.com/netfoundry/ziti-foundation/identity/identity"
	"github.com/netfoundry/ziti-foundation/transport"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

const (
	FlagsCfgMapKey = "@flags"
)

func LoadConfigMap(path string) (map[interface{}]interface{}, error) {
	yamlBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfgmap := make(map[interface{}]interface{})
	if err = yaml.NewDecoder(bytes.NewReader(yamlBytes)).Decode(&cfgmap); err != nil {
		return nil, err
	}

	config.InjectEnv(cfgmap)

	return cfgmap, nil
}

func SetConfigMapFlags(cfgmap map[interface{}]interface{}, flags map[string]*pflag.Flag) {
	cfgmap[FlagsCfgMapKey] = flags
}

type Config struct {
	Id        *identity.TokenId
	Forwarder *forwarder.Options
	Trace     struct {
		Handler *channel2.TraceHandler
	}
	Profile struct {
		Memory struct {
			Path     string
			Interval time.Duration
		}
		CPU struct {
			Path string
		}
	}
	Ctrl struct {
		Endpoint transport.Address
		Options  *channel2.Options
	}
	Link struct {
		Listeners []map[interface{}]interface{}
		Dialers   []map[interface{}]interface{}
	}
	Dialers   map[string]xgress.OptionsData
	Listeners []listenerBinding
	src       map[interface{}]interface{}
}

func (config *Config) Configure(sub config.Subconfig) error {
	return sub.LoadConfig(config.src)
}

func (config *Config) SetFlags(flags map[string]*pflag.Flag) {
	SetConfigMapFlags(config.src, flags)
}

func LoadConfig(path string) (*Config, error) {
	cfgmap, err := LoadConfigMap(path)

	if err != nil {
		return nil, err
	}

	if value, found := cfgmap["v"]; found {
		if value.(int) != 3 {
			panic("config version mismatch: see docs for information on config updates")
		}
	} else {
		panic("config version mismatch: no configuration version specified")
	}

	identityConfig := identity.IdentityConfig{}
	if value, found := cfgmap["identity"]; found {
		submap := value.(map[interface{}]interface{})
		if value, found := submap["key"]; found {
			identityConfig.Key = value.(string)
		}
		if value, found := submap["cert"]; found {
			identityConfig.Cert = value.(string)
		}
		if value, found := submap["server_cert"]; found {
			identityConfig.ServerCert = value.(string)
		}
		if value, found := submap["server_key"]; found {
			identityConfig.ServerKey = value.(string)
		}
		if value, found := submap["ca"]; found {
			identityConfig.CA = value.(string)
		}
	}

	cfg := &Config{src: cfgmap}

	if id, err := identity.LoadIdentity(identityConfig); err != nil {
		return nil, fmt.Errorf("unable to load identity (%w)", err)
	} else {
		cfg.Id = identity.NewIdentity(id)
	}

	cfg.Forwarder = forwarder.DefaultOptions()
	if value, found := cfgmap["forwarder"]; found {
		if submap, ok := value.(map[interface{}]interface{}); ok {
			if options, err := forwarder.LoadOptions(submap); err == nil {
				cfg.Forwarder = options
			} else {
				return nil, fmt.Errorf("invalid 'forwarder' stanza (%w)", err)
			}
		} else {
			pfxlog.Logger().Warn("invalid or empty 'forwarder' stanza")
		}
	}

	if value, found := cfgmap["trace"]; found {
		submap := value.(map[interface{}]interface{})
		if value, found := submap["path"]; found {
			handler, err := channel2.NewTraceHandler(value.(string), cfg.Id)
			if err != nil {
				return nil, err
			}
			handler.AddDecoder(&channel2.Decoder{})
			handler.AddDecoder(&xgress.Decoder{})
			handler.AddDecoder(&ctrl_pb.Decoder{})
			cfg.Trace.Handler = handler
		}
	}

	if value, found := cfgmap["profile"]; found {
		if submap, ok := value.(map[interface{}]interface{}); ok {
			if value, found := submap["memory"]; found {
				if submap, ok := value.(map[interface{}]interface{}); ok {
					if value, found := submap["path"]; found {
						cfg.Profile.Memory.Path = value.(string)
					}
					if value, found := submap["intervalMs"]; found {
						cfg.Profile.Memory.Interval = time.Duration(value.(int)) * time.Millisecond
					} else {
						cfg.Profile.Memory.Interval = 15 * time.Second
					}
				}
			}
			if value, found := submap["cpu"]; found {
				if submap, ok := value.(map[interface{}]interface{}); ok {
					if value, found := submap["path"]; found {
						cfg.Profile.CPU.Path = value.(string)
					}
				}
			}
		}
	}

	cfg.Ctrl.Options = channel2.DefaultOptions()
	if value, found := cfgmap["ctrl"]; found {
		if submap, ok := value.(map[interface{}]interface{}); ok {
			if value, found := submap["endpoint"]; found {
				address, err := transport.ParseAddress(value.(string))
				if err != nil {
					return nil, fmt.Errorf("cannot parse [ctrl/endpoint] (%s)", err)
				}
				cfg.Ctrl.Endpoint = address
			}
			if value, found := submap["options"]; found {
				if optionsMap, ok := value.(map[interface{}]interface{}); ok {
					cfg.Ctrl.Options = channel2.LoadOptions(optionsMap)
				}
			}
			if cfg.Trace.Handler != nil {
				cfg.Ctrl.Options.PeekHandlers = append(cfg.Ctrl.Options.PeekHandlers, cfg.Trace.Handler)
			}
		}
	}

	if value, found := cfgmap["link"]; found {
		if submap, ok := value.(map[interface{}]interface{}); ok {
			if value, found := submap["listeners"]; found {
				if subarr, ok := value.([]interface{}); ok {
					for _, value := range subarr {
						if lmap, ok := value.(map[interface{}]interface{}); ok {
							if value, found := lmap["binding"]; found {
								if _, ok := value.(string); ok {
									cfg.Link.Listeners = append(cfg.Link.Listeners, lmap)
								} else {
									return nil, fmt.Errorf("[link/listeners] must provide string [binding] (%v)", lmap)
								}
							} else {
								return nil, fmt.Errorf("[link/listeners] must provide [binding] (%v)", lmap)
							}
						} else {
							return nil, fmt.Errorf("[link/listeners] must express a map (%v)", value)
						}
					}
				} else {
					return nil, fmt.Errorf("[link/listenrs] must provide at least one listener (%v)", value)
				}
			}
			if value, found := submap["dialers"]; found {
				if subarr, ok := value.([]interface{}); ok {
					for _, value := range subarr {
						if lmap, ok := value.(map[interface{}]interface{}); ok {
							if value, found := lmap["binding"]; found {
								if _, ok := value.(string); ok {
									cfg.Link.Dialers = append(cfg.Link.Dialers, lmap)
								} else {
									return nil, fmt.Errorf("[link/dialers] must provide string [binding] (%v)", lmap)
								}
							} else {
								return nil, fmt.Errorf("[link/dialers] must provide [binding] (%v)", lmap)
							}
						} else {
							return nil, fmt.Errorf("[link/dialers] must express a map (%v)", value)
						}
					}
				} else {
					return nil, fmt.Errorf("[link/dialers] must provide at least one dialer (%v)", value)
				}
			}
		}
	}

	if value, found := cfgmap["listeners"]; found {
		if subarr, ok := value.([]interface{}); ok {
			for _, value := range subarr {
				if submap, ok := value.(map[interface{}]interface{}); ok {
					binding, found := submap["binding"].(string)
					if !found {
						return nil, fmt.Errorf("[listener] must provide [binding] (%v)", submap)
					}
					_, found = submap["address"]
					if !found {
						return nil, fmt.Errorf("[listener] must provide [address] (%v)", submap)
					}
					cfg.Listeners = append(cfg.Listeners, listenerBinding{name: binding, options: submap})
				}
			}
		}
	}

	if value, found := cfgmap["dialers"]; found {
		if subarr, ok := value.([]interface{}); ok {
			for _, value := range subarr {
				if submap, ok := value.(map[interface{}]interface{}); ok {
					if value, found := submap["binding"]; found {
						binding := value.(string)
						if cfg.Dialers == nil {
							cfg.Dialers = make(map[string]xgress.OptionsData)
						}
						cfg.Dialers[binding] = submap
					} else {
						return nil, fmt.Errorf("[dialer] must provide [binding] (%v)", submap)
					}
				}
			}
		}
	}

	return cfg, nil
}

type listenerBinding struct {
	name    string
	options xgress.OptionsData
}
