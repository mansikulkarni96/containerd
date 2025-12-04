//go:build gofuzz

/*
   Copyright The containerd Authors.

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

<<<<<<<< HEAD:vendor/github.com/containerd/containerd/pkg/cri/sbserver/fuzz.go
package sbserver

import (
	"fmt"

	"github.com/containerd/containerd/pkg/cri/server"
	"github.com/containerd/containerd/pkg/cri/store/sandbox"
)

func SandboxStore(cs server.CRIService) (*sandbox.Store, error) {
	s, ok := cs.(*criService)
	if !ok {
		return nil, fmt.Errorf("%+v is not sbserver.criService", cs)
	}
	return s.sandboxStore, nil
========
package main

import (
	"fmt"
	"log"
	"net"

	"github.com/vishvananda/netlink"
)

// isLoInterfaceUp validates whether the lo interface is up
func isLoInterfaceUp() (bool, error) {
	link, err := netlink.LinkByName("lo")
	if err != nil {
		return false, fmt.Errorf("could not find interface lo: %w", err)
	}
	return link.Attrs().Flags&net.FlagUp != 0, nil
}

func main() {
	up, err := isLoInterfaceUp()
	if err != nil {
		log.Fatalf("could not check lo interface status: %v", err)
	}
	fmt.Printf("Loopback interface is %s\n", map[bool]string{true: "UP", false: "DOWN"}[up])
>>>>>>>> v2.1.0:integration/failpoint/cmd/loopback-v2/main.go
}
