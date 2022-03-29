// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package kubernetes

import (
	"crypto/sha256"
	"encoding/base32"
	"path/filepath"
	"regexp"
	"strings"

	"namespacelabs.dev/foundation/schema"
)

const (
	lowerCaseEncodeBase32 = "0123456789abcdefghijklmnopqrstuv"
)

var (
	validChars     = regexp.MustCompile("[a-z0-9]+")
	base32encoding = base32.NewEncoding(lowerCaseEncodeBase32).WithPadding(base32.NoPadding)
)

// Only lowercase alphanumeric characters or '-', may not start/end with '-', max length 63 chars
// https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#dns-label-names

func labelName(command []string) string {
	var parts []string
	for _, c := range command {
		parts = append(parts, validChars.FindAllString(strings.ToLower(c), -1)...)
	}

	return strings.Join(parts, "-")
}

// We use namespaces to isolate deployments per workspace and environment.
// Using the path base plus a digest provides short, memorable names and avoids collision.
// TODO add knob to allow namespace overwrites if the need arises.
func namespace(ws *schema.Workspace, env *schema.Environment) string {
	parts := []string{strings.ToLower(env.Name)}
	parts = append(parts, validChars.FindAllString(filepath.Base(ws.ModuleName), -1)...)

	h := sha256.New()
	h.Write([]byte(ws.ModuleName)) // Write to a sha256 hash never fails.
	digest := h.Sum(nil)

	// A SHA256 is 32 bytes long, we're guarantee to always have at least 5 characters.
	parts = append(parts, base32encoding.EncodeToString(digest)[:5])
	return strings.Join(parts, "-")
}