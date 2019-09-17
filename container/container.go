package container

import (
	"os"
	"os/exec"
	"syscall"
	"fmt"
)	

func parent() {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	// Cloneflags is only available in Linux
	// CLONE_NEWUTS namespace isolates hostname
	// CLONE_NEWPID namespace isolates processes
	// CLONE_NEWNS namespace isolates mounts
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}

}

func child() {

	// Create cgroup

	syscall.Sethostname([]byte("container"))

	//
	syscall.Mount("rootfs", "rootfs", "", syscall.MS_BIND, "")
	syscall.Chroot("rootfs")
	os.Chdir("/")
	// // Mount /proc inside container so that `ps` command works
	syscall.Mount("proc", "proc", "proc", 0, "")

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}

	// Cleanup mount
	syscall.Unmount("/proc", 0)
	syscall.Unmount("rootfs", 0)

}
