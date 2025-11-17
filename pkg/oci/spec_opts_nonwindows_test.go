<<<<<<< HEAD:oci/spec_opts_nonwindows_test.go
=======
//go:build !windows

>>>>>>> v2.0.7:pkg/oci/spec_opts_nonwindows_test.go
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
<<<<<<< HEAD:oci/spec_opts_nonwindows_test.go
<<<<<<<< HEAD:pkg/oci/spec_opts_nonwindows_test.go
	"context"
	"testing"

	"github.com/containerd/containerd/namespaces"
=======
	"context"
	"testing"

	"github.com/containerd/containerd/v2/pkg/namespaces"
>>>>>>> v2.0.7:pkg/oci/spec_opts_nonwindows_test.go
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

func TestWithDefaultPathEnv(t *testing.T) {
	t.Parallel()
	s := Spec{}
	s.Process = &specs.Process{
		Env: []string{},
	}
	var (
		defaultUnixEnv = "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
		ctx            = namespaces.WithNamespace(context.Background(), "test")
	)
	WithDefaultPathEnv(ctx, nil, nil, &s)
	if !Contains(s.Process.Env, defaultUnixEnv) {
		t.Fatal("default Unix Env not found")
<<<<<<< HEAD:oci/spec_opts_nonwindows_test.go
========
	"fmt"
	"runtime"

	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"golang.org/x/sys/windows"
)

// DefaultSpec returns the current platform's default platform specification.
func DefaultSpec() specs.Platform {
	major, minor, build := windows.RtlGetNtVersionNumbers()
	return specs.Platform{
		OS:           runtime.GOOS,
		Architecture: runtime.GOARCH,
		OSVersion:    fmt.Sprintf("%d.%d.%d", major, minor, build),
		// The Variant field will be empty if arch != ARM.
		Variant: cpuVariant(),
>>>>>>>> v2.0.7:vendor/github.com/containerd/platforms/defaults_windows.go
	}
}

// Default returns the current platform's default platform specification.
func Default() MatchComparer {
	return Only(DefaultSpec())
}
=======
	}
}
>>>>>>> v2.0.7:pkg/oci/spec_opts_nonwindows_test.go
