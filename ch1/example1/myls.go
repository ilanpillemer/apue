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
	if len(os.Args) != 2 {
		fmt.Println("usage: myls directory_name")
		return
	}

	dir, err := C.opendir(C.CString(os.Args[1]))
	if err != nil {
		fmt.Printf("can't open %s: %v\n", os.Args[1], err)
		return
	}
	defer C.closedir(dir)

	for dirent := C.readdir(dir); dirent != nil; dirent = C.readdir(dir) {
		name := C.GoString(&dirent.d_name[0])
		fmt.Println(name)
	}
}
