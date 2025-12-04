//go:build darwin
// +build darwin

/*
   Copyright © 2021 The CDI Authors

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

package cdi

<<<<<<<< HEAD:vendor/github.com/containerd/containerd/platforms/platforms_deprecated_other.go
func getWindowsOsVersion() string {
	return ""
========
import "syscall"

func osSync() {
	_ = syscall.Sync()
>>>>>>>> v2.1.0:vendor/tags.cncf.io/container-device-interface/pkg/cdi/cache_test_darwin.go
}
