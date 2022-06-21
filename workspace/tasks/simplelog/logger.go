// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package simplelog

import (
	"io"

	"namespacelabs.dev/foundation/internal/console/consolesink"
	"namespacelabs.dev/foundation/workspace/tasks"
)

var AlsoReportStartEvents = false

func NewSink(w io.Writer, maxLevel int) tasks.ActionSink {
	return &logger{w, maxLevel}
}

type logger struct {
	out      io.Writer
	maxLevel int // Only display actions at this level or below (all actions are still computed).
}

func (sl logger) shouldLog(ev tasks.EventData) bool {
	return ev.Level <= sl.maxLevel
}

func (sl *logger) Waiting(ra *tasks.RunningAction) {
	// Do nothing.
}

func (sl *logger) Started(ra *tasks.RunningAction) {
	if !AlsoReportStartEvents {
		return
	}

	if !sl.shouldLog(ra.Data) {
		return
	}

	consolesink.NoColors.LogAction(sl.out, ra.Data)
}

func (sl *logger) Done(ra *tasks.RunningAction) {
	if !sl.shouldLog(ra.Data) {
		return
	}

	consolesink.NoColors.LogAction(sl.out, ra.Data)
}

func (sl *logger) Instant(ev *tasks.EventData) {
	if !sl.shouldLog(*ev) {
		return
	}

	consolesink.NoColors.LogAction(sl.out, *ev)
}

func (sl *logger) AttachmentsUpdated(tasks.ActionID, *tasks.ResultData) { /* nothing to do */ }