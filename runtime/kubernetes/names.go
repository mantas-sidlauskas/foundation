// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package kubernetes

import (
	"crypto/sha256"
	"fmt"
	"regexp"
	"strings"

	"namespacelabs.dev/go-ids"
)

var (
	simpleLabelRe = regexp.MustCompile("[^0-9a-zA-Z]")
	domainPartRe  = regexp.MustCompile("[^_0-9a-zA-Z]")
)

func domainFragLike(parts ...string) string {
	return cleanName(250, domainPartRe, ".", parts...)
}

func labelLike(parts ...string) string {
	return cleanName(63, simpleLabelRe, "--", parts...)
}

// It's fairly normal that Kubernetes only accepts keys which match the
// [-._a-zA-Z0-9]+ regex. The strategy here is to replace all non-letter
// non-number characters with "-", and then join each segment with ".".
//
// Example: foobar.com/bar:key becomes foobar-com-bar.key
func cleanName(maxSize int, re *regexp.Regexp, sep string, parts ...string) string {
	clean := make([]string, len(parts))
	for k, str := range parts {
		clean[k] = cleanOnePart(re, maxSize/len(parts), strings.ToLower(str))
	}
	return strings.Join(clean, sep)
}

func cleanOnePart(re *regexp.Regexp, max int, str string) string {
	if len(str) > max {
		parts := strings.Split(str, "/")
		if len(parts) > 1 {
			return cleanOnePart(re, max, parts[len(parts)-1])
		}

		h := sha256.New()
		fmt.Fprint(h, str)
		hash := ids.EncodeToBase32String(h.Sum(nil))

		return fmt.Sprintf("%s-%s", re.ReplaceAllLiteralString(str[:max-5], "-"), hash[:4])
	}

	return re.ReplaceAllLiteralString(str, "-")
}
