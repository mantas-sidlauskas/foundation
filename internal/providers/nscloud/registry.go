// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package nscloud

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"namespacelabs.dev/foundation/internal/artifacts/oci"
	"namespacelabs.dev/foundation/internal/artifacts/registry"
	"namespacelabs.dev/foundation/internal/compute"
	"namespacelabs.dev/foundation/internal/fnapi"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/std/cfg"
)

var DefaultKeychain oci.Keychain = defaultKeychain{}

const loginEndpoint = "login.namespace.so/token"

type nscloudRegistry struct{}

func RegisterRegistry() {
	registry.Register("nscloud", func(ctx context.Context, ck cfg.Configuration) (registry.Manager, error) {
		return nscloudRegistry{}, nil
	})

	oci.RegisterDomainKeychain(registryAddr, DefaultKeychain, oci.Keychain_UseAlways)
}

func (nscloudRegistry) IsInsecure() bool { return false }

func (r nscloudRegistry) AllocateName(repository string) compute.Computable[oci.AllocatedRepository] {
	url := registryAddr
	if strings.HasSuffix(url, "/") {
		url += repository
	} else {
		url += "/" + repository
	}

	imgid := oci.ImageID{Repository: url}

	// We need to make sure our keychain is attached to the name.
	return registry.StaticName(r, imgid, r.IsInsecure(), defaultKeychain{})
}

func (r nscloudRegistry) AttachKeychain(imgid oci.ImageID) (oci.AllocatedRepository, error) {
	return registry.AttachStaticKeychain(r, imgid, defaultKeychain{}), nil
}

type defaultKeychain struct{}

func (dk defaultKeychain) Resolve(ctx context.Context, r authn.Resource) (authn.Authenticator, error) {
	user, err := fnapi.LoadUser()
	if err != nil {
		return nil, err
	}

	ref, err := name.ParseReference(r.String())
	if err != nil {
		return nil, err
	}

	values := url.Values{}
	values.Add("scope", fmt.Sprintf("repository:%s:push,pull", ref.Context().RepositoryStr()))
	values.Add("service", "Authentication")

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://%s?%s", loginEndpoint, values.Encode()), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Namespace-Token", base64.RawStdEncoding.EncodeToString(user.Opaque))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fnerrors.InvocationError("%s: unexpected status when fetching an access token: %d", r, resp.StatusCode)
	}

	tokenData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fnerrors.InvocationError("%s: unexpected error when fetching an access token: %w", r, err)
	}

	var t Token
	if err := json.Unmarshal(tokenData, &t); err != nil {
		return nil, fnerrors.InvocationError("%s: unexpected error when unmarshalling an access token: %w", r, err)
	}

	return &authn.Bearer{Token: t.Token}, nil
}

type Token struct {
	Token string `json:"token"`
}
