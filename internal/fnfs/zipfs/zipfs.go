// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package zipfs

import (
	"archive/zip"
	"context"
	"io/fs"
	"math"

	"namespacelabs.dev/foundation/internal/bytestream"
	"namespacelabs.dev/foundation/internal/compute"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/fnfs"
	"namespacelabs.dev/foundation/internal/fnfs/memfs"
	"namespacelabs.dev/foundation/std/tasks"
)

func Unzip(contents compute.Computable[bytestream.ByteStream]) compute.Computable[fs.FS] {
	return compute.Map(tasks.Action("zip.extract"),
		compute.Inputs().Computable("contents", contents),
		compute.Output{},
		func(ctx context.Context, r compute.Resolved) (fs.FS, error) {
			blob := compute.MustGetDepValue(r, contents, "contents")

			if blob.ContentLength() >= math.MaxInt64 {
				return nil, fnerrors.InternalError("blob is too big")
			}

			bsr, err := bytestream.ReaderAt(blob)
			if err != nil {
				return nil, err
			}

			defer bsr.Close()

			zipr, err := zip.NewReader(bsr, int64(blob.ContentLength()))
			if err != nil {
				return nil, err
			}

			var snapshot memfs.FS
			if err := fnfs.CopyTo(ctx, &snapshot, ".", zipr); err != nil {
				return nil, err
			}

			return &snapshot, nil
		})
}
