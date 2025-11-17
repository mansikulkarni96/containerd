//go:build !windows

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

package oci

import (
	"context"

<<<<<<< HEAD:oci/spec_opts_nonwindows.go
	"github.com/containerd/containerd/containers"
=======
	"github.com/containerd/containerd/v2/core/containers"
>>>>>>> v2.0.7:pkg/oci/spec_opts_nonwindows.go
)

// WithDefaultPathEnv sets the $PATH environment variable to the
// default PATH defined in this package.
func WithDefaultPathEnv(_ context.Context, _ Client, _ *containers.Container, s *Spec) error {
	s.Process.Env = replaceOrAppendEnvValues(s.Process.Env, defaultUnixEnv)
	return nil
}
