package main

import (
	"fmt"
	"os"
	"unsafe"
)

/*
#include <unistd.h>
*/
import "C"

const bufferSize = 4096

func main() {
	var buf [bufferSize]byte
	n, _ := C.read(C.STDIN_FILENO, unsafe.Pointer(&buf[0]), C.ulong(bufferSize))
	for n != 0 {
		if w := C.write(C.STDOUT_FILENO, unsafe.Pointer(&buf[0]), C.ulong(n)); w != n {
			fmt.Println("write error")
			os.Exit(1)
		}
		n = C.read(C.STDIN_FILENO, unsafe.Pointer(&buf[0]), C.ulong(bufferSize))
		if n < 0 {
			fmt.Println("read error")
			os.Exit(1)
		}
	}
}
