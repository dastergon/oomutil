// +build linux
package oomutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// NewOOMProcess returns a pointer to the ProcessOOM struct and an error.
func NewOOMProcess(pid int32) (*ProcessOOM, error) {
	p := &ProcessOOM{}
	p.Pid = pid
	file, err := os.Open(filepath.Join("/proc", strconv.Itoa(int(p.Pid))))
	defer file.Close()
	return p, err
}

// OOMScore returns the /proc/(pid)/oom_score of a specific process at particular point in time.
// The oom_score is is a dynamic value which changes with time.
// The higher the value of oom_score of any process, the higher is
// its likelihood of getting killed by the OOM-killer
// when the system is running out of memory.
func (p *ProcessOOM) OOMScore() (int32, error) {
	return p.fillFromOOMKiller("oom_score")
}

// OOMScoreAdj returns the /proc/(pid)/oom_score_adj value of a specific process.
// This file is used to prevent processes in the system from being killed.
func (p *ProcessOOM) OOMScoreAdj() (int32, error) {
	return p.fillFromOOMKiller("oom_score_adj")
}

// MemoryOvercommit return the /proc/sys/vm/overcommit_memory value.
// This file contains the kernel virtual memory accounting mode.
func (p *ProcessOOM) MemoryOvercommit() (int32, error) {
	return p.fillFromProcSys("vm/overcommit_memory")
}

// fillFromProcSys retrieves values from OOM related files in /proc/sys.
func (p *ProcessOOM) fillFromProcSys(fpath string) (int32, error) {
	return p.fillFrom(filepath.Join("/proc/sys/", fpath))
}

// fillFromOOMKiller retrieves values from OOM related files in /proc.
func (p *ProcessOOM) fillFromOOMKiller(fname string) (int32, error) {
	return p.fillFrom(filepath.Join("/proc", strconv.Itoa(int(p.Pid)), fname))
}

// fillFrom retrieves the values from certain paths in /proc.
func (p *ProcessOOM) fillFrom(fpath string) (int32, error) {
	data, err := ioutil.ReadFile(fpath)
	if err != nil {
		return -1, err
	}
	val, err := strconv.ParseInt(strings.TrimSuffix(string(data), "\n"), 10, 32)
	if err != nil {
		return -1, err
	}
	return int32(val), nil
}
