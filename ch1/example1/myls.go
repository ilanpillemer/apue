package main

/*
#include <dirent.h>
*/
import "C"

import (
	"fmt"
	"os"
)

func main() {
	dir, err := C.opendir(C.CString(os.Args[1]))
	defer C.closedir(dir)

	if err != nil {
		fmt.Printf("can't open %s: %v\n", os.Args[1], err)
		return
	}

	for dirent := C.readdir(dir); dirent != nil; dirent = C.readdir(dir) {
		name := C.GoString(&dirent.d_name[0])
		fmt.Println(name)
	}
}
