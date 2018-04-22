// Package oomutil implements some read-only operations for determining the Out-Of-Memory (OOM) status of a process on Linux.
package oomutil

import (
	"github.com/shirou/gopsutil/process"
)

// ProcessOOM embeds process.Process.
type ProcessOOM struct {
	process.Process
}
