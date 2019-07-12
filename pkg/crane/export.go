// Copyright 2018 Google LLC All Rights Reserved.
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

package crane

import (
	"io"

	v1 "github.com/flant/go-containerregistry/pkg/v1"
	"github.com/flant/go-containerregistry/pkg/v1/mutate"
)

// Export writes the filesystem contents (as a tarball) of img to w.
func Export(img v1.Image, w io.WriteCloser) error {
	fs := mutate.Extract(img)
	_, err := io.Copy(w, fs)
	return err
}
