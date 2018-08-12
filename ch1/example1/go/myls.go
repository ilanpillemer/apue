package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// dirent.h provides opendir, readdir, closedir
// stdlib.h provides free

// the equivalent standard library calls in
// golang are provide by io.ioutil

// the golang authors decided not to
// return the "." and ".." directory entries

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: myls directory_name")
		os.Exit(1)
	}

	dir, err := ioutil.ReadDir(os.Args[1])
	if err != nil {
		fmt.Printf("can't open %s: %v\n", os.Args[1], err)
		os.Exit(1)
	}

	fmt.Println(".")
	fmt.Println("..")
	for _, f := range dir {
		name := f.Name()
		fmt.Println(name)
	}
}
