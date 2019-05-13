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
	"log"

	"github.com/flant/go-containerregistry/pkg/authn"
	"github.com/flant/go-containerregistry/pkg/name"
	"github.com/flant/go-containerregistry/pkg/v1/remote"
	"github.com/flant/go-containerregistry/pkg/v1/types"
	"github.com/spf13/cobra"
)

func init() { Root.AddCommand(NewCmdCopy()) }

// NewCmdCopy creates a new cobra.Command for the copy subcommand.
func NewCmdCopy() *cobra.Command {
	return &cobra.Command{
		Use:     "copy",
		Aliases: []string{"cp"},
		Short:   "Efficiently copy a remote image from src to dst",
		Args:    cobra.ExactArgs(2),
		Run:     doCopy,
	}
}

func doCopy(_ *cobra.Command, args []string) {
	src, dst := args[0], args[1]
	srcRef, err := name.ParseReference(src)
	if err != nil {
		log.Fatalf("parsing reference %q: %v", src, err)
	}
	log.Printf("Pulling %v", srcRef)

	desc, err := remote.Get(srcRef, remote.WithAuthFromKeychain(authn.DefaultKeychain))
	if err != nil {
		log.Fatalf("fetching image %q: %v", srcRef, err)
	}

	dstRef, err := name.ParseReference(dst)
	if err != nil {
		log.Fatalf("parsing reference %q: %v", dst, err)
	}
	log.Printf("Pushing %v", dstRef)

	switch desc.MediaType {
	case types.OCIImageIndex, types.DockerManifestList:
		// Handle indexes separately.
		if err := copyIndex(desc, dstRef); err != nil {
			log.Fatalf("failed to copy index: %v", err)
		}
	default:
		// Assume anything else is an image, since some registries don't set mediaTypes properly.
		if err := copyImage(desc, dstRef); err != nil {
			log.Fatalf("failed to copy image: %v", err)
		}
	}
}

func copyImage(desc *remote.Descriptor, dstRef name.Reference) error {
	img, err := desc.Image()
	if err != nil {
		return err
	}
	return remote.Write(dstRef, img, remote.WithAuthFromKeychain(authn.DefaultKeychain))
}

func copyIndex(desc *remote.Descriptor, dstRef name.Reference) error {
	idx, err := desc.ImageIndex()
	if err != nil {
		return err
	}
	return remote.WriteIndex(dstRef, idx, remote.WithAuthFromKeychain(authn.DefaultKeychain))
}
