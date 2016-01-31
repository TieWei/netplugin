/***
Copyright 2014 Cisco Systems Inc. All rights reserved.

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

package drivers

import (
	"encoding/json"
	"fmt"

	"github.com/contiv/netplugin/core"
	"github.com/contiv/netplugin/netmaster/mastercfg"
)

// OvsOperEndpointState is the necessary data used to perform operations on endpoints.
type OvsOperEndpointState struct {
	core.CommonState
	NetID       string `json:"netID"`
	ContName    string `json:"contName"`
	ServiceName string `json:"serviceName"`
	ContUUID    string `json:"contUUID"`
	AttachUUID  string `json:"attachUUID"`
	IPAddress   string `json:"ipAddress"`
	MacAddress  string `json:"macAddress"`
	HomingHost  string `json:"homingHost"`
	IntfName    string `json:"intfName"`
	PortName    string `json:"portName"`
	VtepIP      string `json:"vtepIP"`
}

// Matches matches the fields updated from configuration state
func (s *OvsOperEndpointState) Matches(c *mastercfg.CfgEndpointState) bool {
	return s.NetID == c.NetID &&
		s.ContName == c.ContName &&
		s.AttachUUID == c.AttachUUID &&
		s.IPAddress == c.IPAddress &&
		s.MacAddress == c.MacAddress &&
		s.HomingHost == c.HomingHost &&
		s.IntfName == c.IntfName &&
		s.VtepIP == c.VtepIP
}

// Write the state.
func (s *OvsOperEndpointState) Write() error {
	key := fmt.Sprintf(endpointOperPath, s.ID)
	return s.StateDriver.WriteState(key, s, json.Marshal)
}

// Read the state for a given identifier.
func (s *OvsOperEndpointState) Read(id string) error {
	key := fmt.Sprintf(endpointOperPath, id)
	return s.StateDriver.ReadState(key, s, json.Unmarshal)
}

// ReadAll reads all state into separate objects.
func (s *OvsOperEndpointState) ReadAll() ([]core.State, error) {
	return s.StateDriver.ReadAllState(endpointOperPathPrefix, s, json.Unmarshal)
}

// Clear removes the state.
func (s *OvsOperEndpointState) Clear() error {
	key := fmt.Sprintf(endpointOperPath, s.ID)
	return s.StateDriver.ClearState(key)
}
