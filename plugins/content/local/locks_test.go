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

<<<<<<<< HEAD:vendor/github.com/containerd/log/logtest/context_deprecated.go
package logtest
========
package local
>>>>>>>> v2.0.7:plugins/content/local/locks_test.go

import (
	"testing"

<<<<<<<< HEAD:vendor/github.com/containerd/log/logtest/context_deprecated.go
	"github.com/containerd/log/logtest"
)

// WithT adds a logging hook for the given test
// Changes debug level to debug, clears output, and
// outputs all log messages as test logs.
//
// Deprecated: use [logtest.WithT].
func WithT(ctx context.Context, t testing.TB) context.Context {
	return logtest.WithT(ctx, t)
========
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTryLock(t *testing.T) {
	s := &store{locks: map[string]*lock{}}

	err := s.tryLock("testref")
	assert.NoError(t, err)
	defer s.unlock("testref")

	err = s.tryLock("testref")
	require.NotNil(t, err)
	assert.Contains(t, err.Error(), "ref testref locked for ")
>>>>>>>> v2.0.7:plugins/content/local/locks_test.go
}
