package main

// "io/ioutil"
// "log"
// "os"
// "os/exec"
// "path/filepath"
// "strconv"
// "syscall"
// "fmt"
import (
	"fmt"
	"main/config"
	"main/images"
)

func main() {

	tmp, _ := config.Config()
	fmt.Printf(tmp.PathImages)

	a := images.InitContainerImage()
	print(a)

	// switch os.Args[1] {
	// case "run":
	// 	parent()
	// case "child":
	// 	child()
	// default:
	// 	log.Fatal("Unknown command. Use run <command_name>, like `run /bin/bash` or `run echo hello`")
	// }
}

// func parent() {
// 	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
// 	// Cloneflags is only available in Linux
// 	// CLONE_NEWUTS namespace isolates hostname
// 	// CLONE_NEWPID namespace isolates processes
// 	// CLONE_NEWNS namespace isolates mounts
// 	cmd.SysProcAttr = &syscall.SysProcAttr{
// 		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
// 	}
// 	cmd.Stdin = os.Stdin
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr

// 	if err := cmd.Run(); err != nil {
// 		fmt.Println("ERROR", err)
// 		os.Exit(1)
// 	}

// }

// func child() {
// 	// Create cgroup
// 	cg()

// 	must(syscall.Sethostname([]byte("container")))

// 	//
// 	must(syscall.Mount("rootfs", "rootfs", "", syscall.MS_BIND, ""))
// 	must(syscall.Chroot("rootfs"))
// 	must(os.Chdir("/"))
// 	// // Mount /proc inside container so that `ps` command works
// 	must(syscall.Mount("proc", "proc", "proc", 0, ""))

// 	cmd := exec.Command(os.Args[2], os.Args[3:]...)
// 	cmd.Stdin = os.Stdin
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	if err := cmd.Run(); err != nil {
// 		fmt.Println("ERROR", err)
// 		os.Exit(1)
// 	}

// 	// Cleanup mount
// 	must(syscall.Unmount("/proc", 0))
// 	must(syscall.Unmount("rootfs", 0))

// }

// func cg() {
// 	// cgroup location in Ubuntu
// 	cgroups := "/sys/fs/cgroup/"

// 	mem := filepath.Join(cgroups, "memory")
// 	kontainer := filepath.Join(mem, "kontainer")
// 	os.Mkdir(kontainer, 0755)
// 	// Limit memory to 1mb
// 	must(ioutil.WriteFile(filepath.Join(kontainer, "memory.limit_in_bytes"), []byte("999424"), 0700))
// 	// Cleanup cgroup when it is not being used
// 	must(ioutil.WriteFile(filepath.Join(kontainer, "notify_on_release"), []byte("1"), 0700))

// 	pid := strconv.Itoa(os.Getpid())
// 	// Apply this and any child process in this cgroup
// 	must(ioutil.WriteFile(filepath.Join(kontainer, "cgroup.procs"), []byte(pid), 0700))
// }

// func must(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
