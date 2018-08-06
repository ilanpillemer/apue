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
// http://man7.org/linux/man-pages/man3/readdir.3.html

// syscall/dirent.go

// The data returned by readdir() may be overwritten by subsequent calls
// to readdir() for the same directory stream.

// On success, readdir() returns a pointer to a dirent structure.  (This
// structure may be statically allocated; do not attempt to free(3) it.)

// If the end of the directory stream is reached, NULL is returned and
// errno is not changed.  If an error occurs, NULL is returned and errno
// is set appropriately.  To distinguish end of stream and from an
// error, set errno to zero before calling readdir() and then check the
// value of errno if NULL is returned.

// In the current POSIX.1 specification (POSIX.1-2008), readdir() is not
//       required to be thread-safe.  However, in modern implementations
//       (including the glibc implementation), concurrent calls to readdir()
//       that specify different directory streams are thread-safe.  In cases
//       where multiple threads must read from the same directory stream,
//       using readdir() with external synchronization is still preferable to
//       the use of the deprecated readdir_r(3) function.  It is expected that
//       a future version of POSIX.1 will require that readdir() be thread-
//       safe when concurrently employed on different directory streams.

// A directory stream is opened using opendir(3).

//        The order in which filenames are read by successive calls to
//        readdir() depends on the filesystem implementation; it is unlikely
//        that the names will be sorted in any fashion.

//        Only the fields d_name and (as an XSI extension) d_ino are specified
//        in POSIX.1.  Other than Linux, the d_type field is available mainly
//        only on BSD systems.  The remaining fields are available on many, but
//        not all systems.  Under glibc, programs can check for the
//        availability of the fields not defined in POSIX.1 by testing whether
//        the macros _DIRENT_HAVE_D_NAMLEN, _DIRENT_HAVE_D_RECLEN,
//        _DIRENT_HAVE_D_OFF, or _DIRENT_HAVE_D_TYPE are defined.

//    The d_name field
//        The dirent structure definition shown above is taken from the glibc
//        headers, and shows the d_name field with a fixed size.

//        Warning: applications should avoid any dependence on the size of the
//        d_name field.  POSIX defines it as char d_name[], a character array
//        of unspecified size, with at most NAME_MAX characters preceding the
//        terminating null byte ('\0').

//        POSIX.1 explicitly notes that this field should not be used as an
//        lvalue.  The standard also notes that the use of sizeof(d_name) is
//        incorrect; use strlen(d_name) instead.  (On some systems, this field
//        is defined as char d_name[1]!)  By implication, the use sizeof(struct
//        dirent) to capture the size of the record including the size of
//        d_name is also incorrect.

//        Note that while the call

//            fpathconf(fd, _PC_NAME_MAX)

//        returns the value 255 for most filesystems, on some filesystems
//        (e.g., CIFS, Windows SMB servers), the null-terminated filename that
//        is (correctly) returned in d_name can actually exceed this size.  In
//        such cases, the d_reclen field will contain a value that exceeds the
//        size of the glibc dirent structure shown above.

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

// https://www.freebsd.org/cgi/man.cgi?query=getdirentries&sektion=2&apropos=0&manpath=freebsd

// The actual number of bytes transferred is returned. The current
//      position pointer associated with fd is set to point to the
//      next block of entries. The pointer may not advance by the
//      number of bytes returned by getdirentries() or getdents(). A
//      value of zero is returned when the end of the directory has
//      been reached.

// ztypes_darwin_386.go

// ERRORS
//      The getdirentries() system	call will fail if:

//      [EBADF]		The fd argument	is not a valid file descriptor open
// 			for reading.

//      [EFAULT]		Either buf or non-NULL basep point outside the allo-
// 			cated address space.

//      [EINVAL]		The file referenced by fd is not a directory, or
// 			nbytes is too small for	returning a directory entry or
// 			block of entries, or the current position pointer is
// 			invalid.

//      [EIO]		An I/O error occurred while reading from or writing to
// 			the file system.

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
			if int(rl_offset) > len(p) || int(rl_offset+1) > len(p) {
				break
			}
			rl := uint64(p[rl_offset]) | uint64(p[rl_offset+1])<<8
			if int(rl) > len(p) {
				break
			}
			// put record into its own slice
			record := p[:rl]
			//nl_size := unsafe.Sizeof(Dirent{}.Namlen)
			if len(record) == 0 {
				//log.Println("We have drained... need more buffer")
				//n, _ = unix.ReadDirent(fd, p)
				break
			} else {
				nl_offset := unsafe.Offsetof(Dirent{}.Namlen)
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
