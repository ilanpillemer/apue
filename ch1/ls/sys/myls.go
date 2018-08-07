package main

import (
	"fmt"
	"log"
	"os"
	"unsafe"

	"golang.org/x/sys/unix"
)

type Dirent struct {
	Ino       uint64
	Seekoff   uint64
	Reclen    uint16
	Namlen    uint16
	Type      uint8
	Name      [1024]int8
	Pad_cgo_0 [3]byte
}

func main() {
	var perm uint32
	perm = unix.O_DIRECTORY

	fd, err := unix.Open(os.Args[1], unix.O_DIRECTORY, perm)
	if err != nil {
		log.Fatal(err)
	}

	buff := make([]byte, 256) // magic number. but needs to big enough to prevent infinite loops into buffer.
	for n, _ := unix.ReadDirent(fd, buff); n != 0; n, _ = unix.ReadDirent(fd, buff) {
		var p []byte
		p = append(p, buff...)
		for len(p) > 0 {
			rl_offset := unsafe.Offsetof(Dirent{}.Reclen)
			//rl_size := unsafe.Sizeof(Dirent{}.Reclen) // will always be 2, as its a uint16
			if int(rl_offset) > len(p) || int(rl_offset+1) > len(p) {
				break
			}
			rl := uint64(p[rl_offset]) | uint64(p[rl_offset+1])<<8 // assuming big endian
			if int(rl) > len(p) {
				break
			}
			record := p[:rl]
			//nl_size := unsafe.Sizeof(Dirent{}.Namlen)
			if len(record) == 0 {
				break
			} else {
				nl_offset := unsafe.Offsetof(Dirent{}.Namlen)
				nl := uint64(record[nl_offset]) | uint64(record[nl_offset+1])<<8 // assuming big endian
				name_offset := uint64(unsafe.Offsetof(Dirent{}.Name))

				if name_offset+nl > uint64(len(p)) {
					break
				}

				name := record[name_offset : name_offset+nl]
				fmt.Println(string(name))

				p = p[rl:]
			}
		}

	}
}
