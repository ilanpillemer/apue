package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"log"
	"os"
	"syscall"
)

// http://pubs.opengroup.org/onlinepubs/7908799/xsh/dirent.h.html
// https://www.gnu.org/software/libc/manual/html_node/Directory-Entries.html

// https://linux.die.net/man/3/readdir

// http://man7.org/linux/man-pages/man3/getdirentries.3.html

// syscall/dirent.go

// getdirentries() returns the number of bytes read or zero when at the
// end of the directory.  If an error occurs, -1 is returned, and errno
// is set appropriately.

// syscall_bsd.go

// func ReadDirent(fd int, buf []byte) (n int, err error) {
// 	// Final argument is (basep *uintptr) and the syscall doesn't take nil.
// 	// 64 bits should be enough. (32 bits isn't even on 386). Since the
// 	// actual system call is getdirentries64, 64 is a good guess.
// 	// TODO(rsc): Can we use a single global basep for all calls?
// 	var base = (*uintptr)(unsafe.Pointer(new(uint64)))
// 	return Getdirentries(fd, buf, base)
// }

// http://man7.org/linux/man-pages/man2/open.2.html

// type Dirent struct {
// 	Ino       uint64
// 	Seekoff   uint64
// 	Reclen    uint16
// 	Namlen    uint16
// 	Type      uint8
// 	Name      [1024]int8
// 	Pad_cgo_0 [3]byte
// }

// file_unix.go

// type dirInfo struct {
// 	buf  []byte // buffer for directory I/O
// 	nbuf int    // length of buf; return value from Getdirentries
// 	bufp int    // location of next record in buf.
// }

// type file struct {
// 	pfd         poll.FD
// 	name        string
// 	dirinfo     *dirInfo // nil unless directory being read
// 	nonblock    bool     // whether we set nonblocking mode
// 	stdoutOrErr bool     // whether this is stdout or stderr
// }

func main() {
	var perm uint32
	perm = unix.O_DIRECTORY

	fd, err := unix.Open(os.Args[1], unix.O_DIRECTORY, perm)
	if err != nil {
		log.Fatal(err)
	}

	p := make([]byte, 128)

	for {
		n, err := unix.ReadDirent(fd, p)

		if n == 0 {
			break
		}

		if err != nil {
			log.Fatal("sys err : ", err.Error())
		}

		var names []string
		_, _, newnames := syscall.ParseDirent(p, 3, names)

		for _, v := range newnames {
			fmt.Println(v)
		}

	}
}
