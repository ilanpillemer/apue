package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"log"
	"os"
	//	"syscall"
	"unsafe"
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

// ztypes_darwin_386.go
type Dirent struct {
	Ino       uint64
	Seekoff   uint64
	Reclen    uint16
	Namlen    uint16
	Type      uint8
	Name      [1024]int8
	Pad_cgo_0 [3]byte
}

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

// func main() {
// 	var perm uint32
// 	perm = unix.O_DIRECTORY

// 	fd, err := unix.Open(os.Args[1], unix.O_DIRECTORY, perm)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	p := make([]byte, 128)

// 	for {
// 		n, err := unix.ReadDirent(fd, p)

// 		if n == 0 {
// 			break
// 		}

// 		if err != nil {
// 			log.Fatal("sys err : ", err.Error())
// 		}

// 		var names []string
// 		_, _, newnames := syscall.ParseDirent(p, 3, names)

// 		for _, v := range newnames {
// 			fmt.Println(v)
// 		}

// 	}
// }

func main() {
	var perm uint32
	perm = unix.O_DIRECTORY

	fd, err := unix.Open(os.Args[1], unix.O_DIRECTORY, perm)
	if err != nil {
		log.Fatal(err)
	}

	buff := make([]byte, 128) // magic number?

	// drain all directory entries buffer of directory entries...
	// each call to ReadDirent updates the buffer and returns the
	// number of bytes read or zero when at the end of the
	// directory.

	for n, _ := unix.ReadDirent(fd, buff); n != 0; n, _ = unix.ReadDirent(fd, buff) {
		p := buff
		// drain the data in p
		for len(p) > 0 {
			// get first directory entry record in the buffer

			// get actual length of record, requires
			// reading it from the c struct, which we know
			// we are at the beginning of

			rl_offset := unsafe.Offsetof(Dirent{}.Reclen)
			//rl_size := unsafe.Sizeof(Dirent{}.Reclen) // will always be 2, as its a uint16
			rl := uint64(p[rl_offset]) | uint64(p[rl_offset+1])<<8
			// put record into its own slice
			record := p[:rl]
			nl_offset := unsafe.Offsetof(Dirent{}.Namlen)
			//nl_size := unsafe.Sizeof(Dirent{}.Namlen)
			if len(record) == 0 {
				//log.Println("We have drained... need more buffer")
				//n, _ = unix.ReadDirent(fd, p)
				break
			} else {
				nl := uint64(record[nl_offset]) | uint64(record[nl_offset+1])<<8
				name_offset := uint64(unsafe.Offsetof(Dirent{}.Name))

				if name_offset+nl > uint64(len(p)) {
					break
				}

				name := record[name_offset : name_offset+nl]
				fmt.Println(string(name))

				// drain buffer
				p = p[rl:]
			}
		}

	}
}

// func ParseDirent(buf []byte, max int, names []string) (consumed int, count int, newnames []string) {
// 	origlen := len(buf)
// 	count = 0
// 	for max != 0 && len(buf) > 0 {
// 		reclen, ok := direntReclen(buf)
// 		if !ok || reclen > uint64(len(buf)) {
// 			return origlen, count, names
// 		}
// 		rec := buf[:reclen]
// 		buf = buf[reclen:]
// 		ino, ok := direntIno(rec)
// 		if !ok {
// 			break
// 		}
// 		if ino == 0 { // File absent in directory.
// 			continue
// 		}
// 		const namoff = uint64(unsafe.Offsetof(Dirent{}.Name))
// 		namlen, ok := direntNamlen(rec)
// 		if !ok || namoff+namlen > uint64(len(rec)) {
// 			break
// 		}
// 		name := rec[namoff : namoff+namlen]
// 		for i, c := range name {
// 			if c == 0 {
// 				name = name[:i]
// 				break
// 			}
// 		}
// 		// Check for useless names before allocating a string.
// 		if string(name) == "." || string(name) == ".." {
// 			continue
// 		}
// 		max--
// 		count++
// 		names = append(names, string(name))
// 	}
// 	return origlen - len(buf), count, names
// }
