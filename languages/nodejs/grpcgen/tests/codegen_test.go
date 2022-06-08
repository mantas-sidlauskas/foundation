// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"gotest.tools/assert"
	"namespacelabs.dev/foundation/internal/fnfs/memfs"
	"namespacelabs.dev/foundation/languages/nodejs/grpcgen"
	"namespacelabs.dev/foundation/workspace/source/protos"
)

const (
	testFile1 = "test1.proto"
)

var (
	testFiles = []string{
		testFile1,
		"nested/test2.proto",
	}
)

func TestCodegen(t *testing.T) {
	var fsys memfs.FS
	for _, fn := range testFiles {
		file, err := os.ReadFile(fn)
		assert.NilError(t, err)
		fsys.Add(fn, file)
	}

	fds, err := protos.Parse(&fsys, testFiles)
	assert.NilError(t, err)

	for _, fd := range fds.File {
		generatedCode, err := grpcgen.Generate(fd, fds)
		assert.NilError(t, err)
		// Adding ".generated" extension so the ".ts" files don't appear broken (due to missing dependencies) in the IDE.
		assert.NilError(t, cupaloy.SnapshotMulti(fmt.Sprintf("%s.ts.generated", filepath.Base(fd.GetName())), generatedCode))
	}
}