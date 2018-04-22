package oomutil

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

func testNewProcess(t *testing.T) {
	testPid := os.Getpid()
	ret, err := NewOOMProcess(int32(testPid))
	if err != nil {
		t.Errorf("A valid object initialization should not return error: %v", err)
	}
	if ret == nil {
		t.Errorf("A valid object initialization should not return nil: %T", ret)
	}
}

func testOOMScore(t *testing.T) {
	testPid := os.Getpid()
	ret, err := NewOOMProcess(int32(testPid))
	if err != nil {
		t.Errorf("A valid object initialization should not return error: %v", err)
	}
	if ret == nil {
		t.Errorf("A valid object initialization should not return nil: %T", ret)
	}

	oomScore, err := ret.OOMScore()
	if err != nil {
		t.Errorf("A valid call should not return an error: %v", err)
	}
	if _, err := strconv.ParseInt(string(oomScore), 10, 32); err != nil {
		t.Fatal("A valid result would be an int32 type")
	}
}

func testOOMScoreAdj(t *testing.T) {
	testPid := os.Getpid()
	ret, err := NewOOMProcess(int32(testPid))
	if err != nil {
		t.Errorf("A valid object initialization should not return error: %v", err)
	}
	if ret == nil {
		t.Errorf("A valid object initialization should not return nil: %T", ret)
	}

	oomScoreAdj, err := ret.OOMScoreAdj()
	if err != nil {
		t.Errorf("A valid call should not return an error: %v", err)
	}

	if _, err := strconv.ParseInt(string(oomScoreAdj), 10, 32); err != nil {
		t.Fatal("A valid result would be an int32 type")
	}
}

func testMemoryOvercommit(t *testing.T) {
	testPid := os.Getpid()
	ret, err := NewOOMProcess(int32(testPid))
	if err != nil {
		t.Errorf("A valid object initialization should not return error: %v", err)
	}
	if ret == nil {
		t.Errorf("A valid object initialization should not return nil: %T", ret)
	}

	memoryOvercommit, err := ret.MemoryOvercommit()
	if err != nil {
		t.Errorf("A valid call should not return an error: %v", err)
	}
	if _, err := strconv.ParseInt(string(memoryOvercommit), 10, 32); err != nil {
		t.Fatal("A valid result would be an int32 type")
	}
}

func testFillFrom(t *testing.T) {
	testPid := os.Getpid()
	ret, err := NewOOMProcess(int32(testPid))
	if err != nil {
		t.Errorf("A valid object initialization should not return error: %v", err)
	}
	if ret == nil {
		t.Errorf("A valid object initialization should not return nil: %T", ret)
	}

	val, err := ret.fillFrom(filepath.Join("/proc", strconv.Itoa(int(testPid)), "oom_score"))
	if err != nil {
		t.Errorf("A valid call to fillFrom should not return error: %v", err)
	}

	if _, err := strconv.ParseInt(string(val), 10, 32); err != nil {
		t.Fatal("A valid result would be an int32 type")
	}
}
