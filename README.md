# oomutil
Package oomutil implements some read-only operations for determining the Out-Of-Memory (OOM) status of a process on Linux.

This package retrieves information from:
* /proc/(pid)/oom_score
* /proc/(pid)/oom_score_adj
* /proc/sys/vm/overcommit_memory

### Installation
Install using standard `go get`:

```shell

$ go get github.com/dastergon/oomutil

```

### How to use

This is an example how to use this package:

```golang
pid := os.Getpid()
ps, err := oomutil.NewOOMProcess(int32(pid))
if err != nil {
	log.Fatal(err)
}

oomScore, err := ps.OOMScore()
if err != nil {
	log.Fatal(err)
}

oomScoreAdj, err := ps.OOMScoreAdj()
if err != nil {
    log.Fatal(err)
}

memoryOvercommit, err := ps.MemoryOvercommit()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Memory Overcommit: %d\nPID: %d => OOM Score: %d, OOM Score Adj: %d\n", memoryOvercommit, pid, oomScore, oomScoreAdj)
```

## Determining OOM Status
The OOM-killer checks if the systems is truly out-of-memory, and selects a process to kill.

### /proc/[pid]/oom_score
According to [proc(5)](https://linux.die.net/man/5/proc) man page:
```
/proc/[pid]/oom_score (since Linux 2.6.11)
This file displays the current score that the kernel gives to this process for the purpose of selecting a process for the OOM-killer. A higher score means that the process is more likely to be selected by the OOM-killer. The basis for this score is the amount of memory used by the process, with increases (+) or decreases (-) for factors including:

* whether the process creates a lot of children using fork(2) (+);
* whether the process has been running a long time, or has used a lot of CPU time (-);
* whether the process has a low nice value (i.e., > 0) (+);
* whether the process is privileged (-); and
* whether the process is making direct hardware access (-).

The oom_score also reflects the adjustment specified by the oom_score_adj or oom_adj setting for the process.

The higher the value of oom_score of any process, the higher is its likelihood of getting killed by the OOM Killer when the system is running out of memory.
```

### /proc/[pid]/oom_score_adj
According to [proc(5)](https://linux.die.net/man/5/proc) man page:

```
/proc/[pid]/oom_score_adj (since Linux 2.6.36)
This file can be used to adjust the badness heuristic used to select which process gets killed in out-of-memory conditions.
The badness heuristic assigns a value to each candidate task ranging from 0 (never kill) to 1000 (always kill) to determine which process is targeted. The units are roughly a proportion along that range of allowed memory the process may allocate from, based on an estimation of its current memory and swap use. For example, if a task is using all allowed memory, its badness score will be 1000. If it is using half of its allowed memory, its score will be 500.

There is an additional factor included in the badness score: root processes are given 3% extra memory over other tasks.

The amount of "allowed" memory depends on the context in which the OOM-killer was called. If it is due to the memory assigned to the allocating task's cpuset being exhausted, the allowed memory represents the set of mems assigned to that cpuset (see cpuset(7)). If it is due to a mempolicy's node(s) being exhausted, the allowed memory represents the set of mempolicy nodes. If it is due to a memory limit (or swap limit) being reached, the allowed memory is that configured limit. Finally, if it is due to the entire system being out of memory, the allowed memory represents all allocatable resources.

The value of oom_score_adj is added to the badness score before it is used to determine which task to kill. Acceptable values range from -1000 (OOM_SCORE_ADJ_MIN) to +1000 (OOM_SCORE_ADJ_MAX). This allows user space to control the preference for OOM-killing, ranging from always preferring a certain task or completely disabling it from OOM-killing. The lowest possible value, -1000, is equivalent to disabling OOM-killing entirely for that task, since it will always report a badness score of 0.

Consequently, it is very simple for user space to define the amount of memory to consider for each task. Setting a oom_score_adj value of +500, for example, is roughly equivalent to allowing the remainder of tasks sharing the same system, cpuset, mempolicy, or memory controller resources to use at least 50% more memory. A value of -500, on the other hand, would be roughly equivalent to discounting 50% of the task's allowed memory from being considered as scoring against the task.

For backward compatibility with previous kernels, /proc/[pid]/oom_adj can still be used to tune the badness score. Its value is scaled linearly with oom_score_adj.

Writing to /proc/[pid]/oom_score_adj or /proc/[pid]/oom_adj will change the other with its scaled value.
```

### /proc/sys/vm/overcommit_memory
According to [proc(5)](https://linux.die.net/man/5/proc) man page:
```
This file contains the kernel virtual memory accounting mode. Values are:
0: heuristic overcommit (this is the default)
1: always overcommit, never check
2: always check, never overcommit
```
/proc/sys/vm/overcommit_memory set to 2, disables the OOM-Killer.
