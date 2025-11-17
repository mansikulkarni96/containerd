<<<<<<< HEAD:plugin/plugin_other.go
//go:build !go1.8 || windows || !amd64 || static_build || gccgo || no_dynamic_plugins

=======
>>>>>>> v2.0.7:core/mount/mount_darwin.go
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

package mount

<<<<<<< HEAD:plugin/plugin_other.go
func loadPlugins(path string) (int, error) {
	// plugins not supported until 1.8
	return 0, nil
=======
import "github.com/containerd/errdefs"

// Mount to the provided target.
func (m *Mount) mount(target string) error {
	return errdefs.ErrNotImplemented
>>>>>>> v2.0.7:core/mount/mount_darwin.go
}
