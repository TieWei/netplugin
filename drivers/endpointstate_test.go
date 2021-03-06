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
	"testing"

	"github.com/contiv/netplugin/core"
)

const (
	testEpID  = "testEp"
	epOperKey = endpointOperPathPrefix + testEpID
)

type testEpStateDriver struct{}

var epStateDriver = &testEpStateDriver{}

func (d *testEpStateDriver) Init(instInfo *core.InstanceInfo) error {
	return core.Errorf("Shouldn't be called!")
}

func (d *testEpStateDriver) Deinit() {
}

func (d *testEpStateDriver) Write(key string, value []byte) error {
	return core.Errorf("Shouldn't be called!")
}

func (d *testEpStateDriver) Read(key string) ([]byte, error) {
	return []byte{}, core.Errorf("Shouldn't be called!")
}

func (d *testEpStateDriver) ReadAll(baseKey string) ([][]byte, error) {
	return [][]byte{}, core.Errorf("Shouldn't be called!")
}

func (d *testEpStateDriver) WatchAll(baseKey string, rsps chan [2][]byte) error {
	return core.Errorf("not supported")
}

func (d *testEpStateDriver) validateKey(key string) error {
	if key != epOperKey {
		return core.Errorf("Unexpected key. recvd: %s expected: %s ",
			key, epOperKey)
	}

	return nil
}

func (d *testEpStateDriver) ClearState(key string) error {
	return d.validateKey(key)
}

func (d *testEpStateDriver) ReadState(key string, value core.State,
	unmarshal func([]byte, interface{}) error) error {
	return d.validateKey(key)
}

func (d *testEpStateDriver) ReadAllState(key string, value core.State,
	unmarshal func([]byte, interface{}) error) ([]core.State, error) {
	return nil, core.Errorf("Shouldn't be called!")
}

func (d *testEpStateDriver) WatchAllState(baseKey string, sType core.State,
	unmarshal func([]byte, interface{}) error, rsps chan core.WatchState) error {
	return core.Errorf("not supported")
}

func (d *testEpStateDriver) WriteState(key string, value core.State,
	marshal func(interface{}) ([]byte, error)) error {
	return d.validateKey(key)
}

func TestOperEndpointStateRead(t *testing.T) {
	epOper := &OperEndpointState{}
	epOper.StateDriver = epStateDriver

	err := epOper.Read(testEpID)
	if err != nil {
		t.Fatalf("read oper state failed. Error: %s", err)
	}
}

func TestOperEndpointStateWrite(t *testing.T) {
	epOper := &OperEndpointState{}
	epOper.StateDriver = epStateDriver
	epOper.ID = testEpID

	err := epOper.Write()
	if err != nil {
		t.Fatalf("write oper state failed. Error: %s", err)
	}
}

func TestOperEndpointStateClear(t *testing.T) {
	epOper := &OperEndpointState{}
	epOper.StateDriver = epStateDriver
	epOper.ID = testEpID

	err := epOper.Clear()
	if err != nil {
		t.Fatalf("clear oper state failed. Error: %s", err)
	}
}
