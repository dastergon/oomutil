package oomutil

import (
	"os"
	"testing"
)

func testNewProcess(t *testing.T) {
	testPid := os.Getpid()
	ret, err := NewOOMProcess(int32(testPid))
	if err != nil {
		t.Errorf("A valid object initialization should not return error: %v", err)
	}
	if *ret != &ProcessOOM{} {
		t.Errorf("A valid call to NewOOMProcess should return ProcessOOM")
	}
}

func testOOMScore(t *testing.T) {
	testPid := os.Getpid()
	ret, err := NewOOMProcess(int32(testPid))
	if err != nil {
		t.Errorf("A valid object initialization should not return error: %v", err)
	}
	if *ret != &ProcessOOM{} {
		t.Errorf("A valid call to NewOOMProcess should return ProcessOOM")
	}

	oomScore := ret.OOMScore()
	i, ok := oomScore.(int32)
	if !ok {
		t.Fatal("A valid result would be a int32 type")
	}
}

func testOOMScoreAdj(t *testing.T) {
	testPid := os.Getpid()
	ret, err := NewOOMProcess(int32(testPid))
	if err != nil {
		t.Errorf("A valid object initialization should not return error: %v", err)
	}

	if *ret != &ProcessOOM{} {
		t.Errorf("A valid call to NewOOMProcess should return ProcessOOM")
	}

	oomScoreAdj := ret.OOMScoreAdj()
	i, ok := oomScoreAdj.(int32)
	if !ok {
		t.Fatal("A valid result would be a int32 type")
	}
}

func testMemoryOvercommit(t *testing.T) {
	testPid := os.Getpid()
	ret, err := NewOOMProcess(int32(testPid))
	if err != nil {
		t.Errorf("A valid object initialization should not return error: %v", err)
	}

	if *ret != &ProcessOOM{} {
		t.Errorf("A valid call to NewOOMProcess should return ProcessOOM")
	}

	memoryOvercommit:= ret.MemoryOvercommit()
	i, ok := memoryOvercommit.(int32)
	if !ok {
		t.Fatal("A valid result would be a int32 type")
	}
}

func testFillFrom(t *testing.T) {
	testPid := os.Getpid()
	ret, err := NewOOMProcess(int32(testPid))
	if err != nil {
		t.Errorf("A valid object initialization should not return error: %v", err)
	}

	if *ret != &ProcessOOM{} {
		t.Errorf("A valid call to NewOOMProcess should return ProcessOOM")
	}

	val, err := ret.fillFrom(filepath.Join("/proc", strconv.Itoa(int(p.Pid)), fname))
	if err != nil {
		t.Errorf("A valid call to fillFrom should not return error: %v", err)
	}
	i, ok := val.(int32)
	if !ok {
		t.Fatal("A valid result would be a int32 type")
	}
}
