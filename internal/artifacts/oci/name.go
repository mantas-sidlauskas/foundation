// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package oci

import (
	"context"
	"strings"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"namespacelabs.dev/foundation/schema"
)

type Keychain interface {
	Resolve(context.Context, authn.Resource) (authn.Authenticator, error)
}

type AllocatedName struct {
	InsecureRegistry bool
	Keychain         Keychain
	ImageID
}

func (t AllocatedName) ComputeDigest(context.Context) (schema.Digest, error) {
	return schema.DigestOf("insecureRegistry", t.InsecureRegistry, "repository", t.Repository, "tag", t.Tag, "digest", t.Digest)
}

func defaultTag(digest v1.Hash) string {
	// TODO revisit tag generation.
	// Registry protocol requires a tag. go-containerregistry uses "latest" by default.
	// ECR treats all tags as immutable. This does not combine well, so we infer a stable tag here.
	return strings.TrimPrefix(digest.String(), "sha256:")
}

func ParseTag(tag AllocatedName, digest v1.Hash) (name.Tag, error) {
	var opts []name.Option
	if tag.InsecureRegistry {
		opts = append(opts, name.Insecure)
	}

	opts = append(opts, name.WithDefaultTag(defaultTag(digest)))

	return name.NewTag(tag.ImageRef(), opts...)
}
