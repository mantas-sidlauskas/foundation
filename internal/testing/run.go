// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package testing

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"strings"

	"namespacelabs.dev/foundation/internal/artifacts/oci"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/engine/ops"
	"namespacelabs.dev/foundation/internal/executor"
	"namespacelabs.dev/foundation/internal/fnfs/memfs"
	"namespacelabs.dev/foundation/internal/syncbuffer"
	"namespacelabs.dev/foundation/internal/uniquestrings"
	"namespacelabs.dev/foundation/provision/deploy"
	"namespacelabs.dev/foundation/runtime"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/workspace/compute"
	"namespacelabs.dev/foundation/workspace/tasks"
)

type testRun struct {
	Env ops.Environment // Doesn't affect the output.

	TestName       string
	TestBinPkg     schema.PackageName
	TestBinCommand []string
	TestBinImageID compute.Computable[oci.ImageID]

	Stack *schema.Stack
	Focus []string // Package names.
	Plan  compute.Computable[*deploy.Plan]
	Debug bool

	compute.LocalScoped[fs.FS]
}

var _ compute.Computable[fs.FS] = &testRun{}

func (rt *testRun) Action() *tasks.ActionEvent {
	return tasks.Action("test").Arg("name", rt.TestName).Arg("package_name", rt.TestBinPkg)
}

func (rt *testRun) Inputs() *compute.In {
	return compute.Inputs().
		Str("testName", rt.TestName).
		Stringer("testBinPkg", rt.TestBinPkg).
		Strs("testBinCommand", rt.TestBinCommand).
		Computable("testBin", rt.TestBinImageID).
		Proto("stack", rt.Stack).
		Strs("focus", rt.Focus).
		Computable("plan", rt.Plan).
		Bool("debug", rt.Debug)
}

func (rt *testRun) Compute(ctx context.Context, r compute.Resolved) (fs.FS, error) {
	p := compute.GetDepValue(r, rt.Plan, "plan")

	compute.On(ctx).Cleanup(tasks.Action("test.cleanup"), func(ctx context.Context) error {
		if err := runtime.For(rt.Env).DeleteRecursively(ctx); err != nil {
			return err
		}

		return nil
	})

	waiters, err := p.Deployer.Apply(ctx, runtime.TaskServerDeploy, rt.Env)
	if err != nil {
		return nil, err
	}

	if err := ops.WaitMultiple(ctx, waiters, nil); err != nil {
		return nil, err
	}

	testRun := runtime.ServerRunOpts{
		Image:              compute.GetDepValue(r, rt.TestBinImageID, "testBin"),
		Command:            rt.TestBinCommand,
		Args:               nil,
		ReadOnlyFilesystem: true,
	}

	if rt.Debug {
		testRun.Args = append(testRun.Args, "--debug")
	}

	localCtx, cancelAll := context.WithCancel(ctx)
	defer cancelAll()

	ex, wait := executor.New(localCtx)

	var focus uniquestrings.List

	for _, srv := range rt.Focus {
		focus.Add(srv)
	}

	var serverLogs []*syncbuffer.ByteBuffer // Follows same indexing as rt.Focus.

	for _, entry := range rt.Stack.Entry {
		srv := entry.Server // Close on srv.

		w, serverLog := makeLog(ctx, srv.Name, focus.Has(srv.PackageName))
		serverLogs = append(serverLogs, serverLog)

		ex.Go(func(ctx context.Context) error {
			err := runtime.For(rt.Env).StreamLogsTo(ctx, w, srv, runtime.StreamLogsOpts{})
			if errors.Is(err, context.Canceled) {
				return nil
			}
			return err
		})
	}

	testLog, testLogBuf := makeLog(ctx, "testlog", true)

	ex.Go(func(ctx context.Context) error {
		defer cancelAll() // When the test is done, cancel logging.

		if err := runtime.For(rt.Env).RunOneShot(ctx, rt.TestBinPkg, testRun, testLog); err != nil {
			var e runtime.ErrContainerExitStatus
			if errors.As(err, &e) && e.ExitCode > 0 {
				return errors.New("test failed")
			} else {
				return err
			}
		}

		return nil
	})

	if err := wait(); err != nil {
		return nil, err
	}

	var fs memfs.FS
	fs.Add("test.log", testLogBuf.Seal().Bytes())

	for k, entry := range rt.Stack.Entry {
		fs.Add(fmt.Sprintf("server/%s.log", strings.ReplaceAll(entry.GetPackageName().String(), "/", "-")), serverLogs[k].Seal().Bytes())
	}

	return &fs, nil
}

func makeLog(ctx context.Context, name string, focus bool) (io.Writer, *syncbuffer.ByteBuffer) {
	buf := syncbuffer.NewByteBuffer()
	if !focus {
		return buf, buf
	}

	w := io.MultiWriter(console.Output(ctx, name), buf)
	return w, buf
}