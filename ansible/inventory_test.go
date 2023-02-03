/*
Copyright 2022.

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

package ansible

import (
	"testing"
)

func TestInventoryMarshall(t *testing.T) {
	inventory := MakeInventory()
	all := inventory.AddGroup("all")
	host := all.AddHost("testing")
	host.Vars["ansible_host"] = "host.test"
	childTest := all.AddChild(MakeGroup("child_test"))
	childHost := childTest.AddHost("child_testing")
	childHost.Vars["ansible_host"] = "child.host.test"

	invData, err := inventory.MarshalYAML()
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
	expected := `all:
    hosts:
        testing:
            ansible_host: host.test
    children:
        all:
            hosts:
                child_testing:
                    ansible_host: child.host.test
`
	result := string(invData)
	if result != expected {
		t.Log("error should be"+expected+", but got", result)
		t.Fail()
	}
}
