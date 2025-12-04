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

package compression

<<<<<<< HEAD:defaults/defaults_snapshotter_windows.go
const (
	// DefaultSnapshotter will set the default snapshotter for the platform.
	// This will be based on the client compilation target, so take that into
	// account when choosing this value.
	DefaultSnapshotter = "windows"
	// DefaultDiffer will set the default differ for the platform.
	DefaultDiffer = "walking"
=======
import (
	"bytes"
	"testing"
>>>>>>> v2.1.0:pkg/archive/compression/compression_fuzzer_test.go
)

func FuzzDecompressStream(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		_, _ = DecompressStream(bytes.NewReader(data))
	})
}
