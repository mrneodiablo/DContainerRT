package cgroup

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func CGroupMemory(containerName string, pid string, membytesize int64) error {
	cgroups := "/sys/fs/cgroup/"

	containerMem := filepath.Join(filepath.Join(cgroups, "memory"), containerName)
	os.Mkdir(containerMem, 0755)

	// Limit memory to 1mb
	byte_mem := strconv.FormatInt(membytesize, 10)
	err := ioutil.WriteFile(filepath.Join(containerName, "memory.limit_in_bytes"), []byte(byte_mem), 0700)

	// Cleanup cgroup when it is not being used
	err = ioutil.WriteFile(filepath.Join(containerName, "notify_on_release"), []byte("1"), 0700)

	// Apply this and any child process in this cgroup
	err = ioutil.WriteFile(filepath.Join(containerName, "cgroup.procs"), []byte(pid), 0700)

	return err
}

func CgroupCPU(containerName string, pid string, membytesize int64) {
}
