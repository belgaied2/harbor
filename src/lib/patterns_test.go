// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchManifestURLPattern(t *testing.T) {
	_, _, ok := MatchManifestURLPattern("")
	assert.False(t, ok)

	_, _, ok = MatchManifestURLPattern("/v2/")
	assert.False(t, ok)

	repository, reference, ok := MatchManifestURLPattern("/v2/library/hello-world/manifests/latest")
	assert.True(t, ok)
	assert.Equal(t, "library/hello-world", repository)
	assert.Equal(t, "latest", reference)

	repository, reference, ok = MatchManifestURLPattern("/v2/library/hello-world/manifests/sha256:e5785cb0c62cebbed4965129bae371f0589cadd6d84798fb58c2c5f9e237efd9")
	assert.True(t, ok)
	assert.Equal(t, "library/hello-world", repository)
	assert.Equal(t, "sha256:e5785cb0c62cebbed4965129bae371f0589cadd6d84798fb58c2c5f9e237efd9", reference)
}

func TestMatchBlobURLPattern(t *testing.T) {
	_, _, ok := MatchBlobURLPattern("")
	assert.False(t, ok)

	_, _, ok = MatchBlobURLPattern("/v2/")
	assert.False(t, ok)

	repository, digest, ok := MatchBlobURLPattern("/v2/library/hello-world/blobs/sha256:e5785cb0c62cebbed4965129bae371f0589cadd6d84798fb58c2c5f9e237efd9")
	assert.True(t, ok)
	assert.Equal(t, "library/hello-world", repository)
	assert.Equal(t, "sha256:e5785cb0c62cebbed4965129bae371f0589cadd6d84798fb58c2c5f9e237efd9", digest)
}

func TestMatchBlobUploadURLPattern(t *testing.T) {
	_, ok := MatchBlobUploadURLPattern("")
	assert.False(t, ok)

	_, ok = MatchBlobUploadURLPattern("/v2/")
	assert.False(t, ok)

	repository, ok := MatchBlobUploadURLPattern("/v2/library/hello-world/blobs/uploads/")
	assert.True(t, ok)
	assert.Equal(t, "library/hello-world", repository)

	repository, ok = MatchBlobUploadURLPattern("/v2/library/hello-world/blobs/uploads/uuid")
	assert.True(t, ok)
	assert.Equal(t, "library/hello-world", repository)
}
