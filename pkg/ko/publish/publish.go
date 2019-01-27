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

package publish

import (
	"github.com/flant/go-containerregistry/pkg/name"
	v1 "github.com/flant/go-containerregistry/pkg/v1"
)

// Interface abstracts different methods for publishing images.
type Interface interface {
	// Publish uploads the given v1.Image to a registry incorporating the
	// provided string into the image's repository name.  Returns the digest
	// of the published image.
	Publish(v1.Image, string) (name.Reference, error)
}
