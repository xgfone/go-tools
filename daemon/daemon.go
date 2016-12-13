// Make the current process to the daemon process.
//
// NOTICE: DON NOT USE THIS FUNCTION.
//
// When using the daemon in Go, if your program has some goroutines, they won't
// work. One explanation is that the daemon only fork the main goroutine, not
// other goroutines. Furthermore, once the go program is started, all the
// goroutines about GC will run to work; and when forking the current process,
// all those will go away.
//
// Notice: The above only is the guess, but it's in fact that the goroutines
// except the main goroutine don't work when turn the go program into a daemon.
package daemon

import (
	"log"
	"os"
	"runtime"
	"syscall"
)

// Make it into the daemon process.
//
// If chdir is true, chroot to the root direcotry of "/".
// If _close is true, redirect STDIN, STDOUT, STDERR to "/dev/null".
//
// Return true if successfully, or return false.
func Daemon(chdir, _close bool) bool {
	var ret, ret2 uintptr
	var err syscall.Errno

	darwin := runtime.GOOS == "darwin"

	// already a daemon
	if syscall.Getppid() == 1 {
		return true
	}

	// fork off the parent process
	ret, ret2, err = syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)
	if err != 0 {
		return false
	}

	// failure
	if ret2 < 0 {
		os.Exit(-1)
	}

	// handle exception for darwin
	if darwin && ret2 == 1 {
		ret = 0
	}

	// if we got a good PID, then we call exit the parent process.
	if ret > 0 {
		os.Exit(0)
	}

	/* Change the file mode mask */
	_ = syscall.Umask(0)

	// create a new SID for the child process
	s_ret, s_errno := syscall.Setsid()
	if s_errno != nil {
		log.Printf("Error: syscall.Setsid errno: %d", s_errno)
	}
	if s_ret < 0 {
		return false
	}

	if chdir {
		os.Chdir("/")
	}

	if _close {
		f, e := os.OpenFile("/dev/null", os.O_RDWR, 0)
		if e == nil {
			fd := f.Fd()
			syscall.Dup2(int(fd), int(os.Stdin.Fd()))
			syscall.Dup2(int(fd), int(os.Stdout.Fd()))
			syscall.Dup2(int(fd), int(os.Stderr.Fd()))
		}
	}

	return true
}
