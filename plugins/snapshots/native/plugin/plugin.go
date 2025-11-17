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

package plugin

import (
	"errors"

<<<<<<< HEAD:snapshots/native/plugin/plugin.go
	"github.com/containerd/containerd/plugin"
	"github.com/containerd/containerd/snapshots/native"
	"github.com/containerd/platforms"
=======
	"github.com/containerd/containerd/v2/plugins"
	"github.com/containerd/containerd/v2/plugins/snapshots/native"
	"github.com/containerd/platforms"
	"github.com/containerd/plugin"
	"github.com/containerd/plugin/registry"
>>>>>>> v2.0.7:plugins/snapshots/native/plugin/plugin.go
)

// Config represents configuration for the native plugin.
type Config struct {
	// Root directory for the plugin
	RootPath string `toml:"root_path"`
}

func init() {
	registry.Register(&plugin.Registration{
		Type:   plugins.SnapshotPlugin,
		ID:     "native",
		Config: &Config{},
		InitFn: func(ic *plugin.InitContext) (interface{}, error) {
			ic.Meta.Platforms = append(ic.Meta.Platforms, platforms.DefaultSpec())

			config, ok := ic.Config.(*Config)
			if !ok {
				return nil, errors.New("invalid native configuration")
			}

			root := ic.Properties[plugins.PropertyRootDir]
			if len(config.RootPath) != 0 {
				root = config.RootPath
			}

<<<<<<< HEAD:snapshots/native/plugin/plugin.go
			ic.Meta.Exports[plugin.SnapshotterRootDir] = root
=======
			ic.Meta.Exports[plugins.SnapshotterRootDir] = root
>>>>>>> v2.0.7:plugins/snapshots/native/plugin/plugin.go
			return native.NewSnapshotter(root)
		},
	})
}
