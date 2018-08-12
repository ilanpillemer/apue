package main

/*
#include <dirent.h>
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"os"
	"unsafe"
)

// dirent.h provides opendir, readdir, closedir
// stdlib.h provides free

// these calls are not sys calls, but standard library calls
// directory entries are represented different in different
// unix systems, and the system calls vary.

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: myls directory_name")
		os.Exit(1)
	}
	arg := C.CString(os.Args[1])
	defer C.free(unsafe.Pointer(arg))

	dir, err := C.opendir(arg)
	if err != nil {
		fmt.Printf("can't open %s: %v\n", os.Args[1], err)
		os.Exit(1)
	}
	defer C.closedir(dir)

	for dirent := C.readdir(dir); dirent != nil; dirent = C.readdir(dir) {
		name := C.GoString(&dirent.d_name[0])
		fmt.Println(name)
	}
}
