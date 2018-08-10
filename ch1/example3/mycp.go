package main

/*
#include <stdio.h>
*/
import "C"
import (
	"fmt"
	"log"
)

func main() {
	for char, inErr := C.getc(C.stdin); char != C.EOF; char, inErr = C.getc(C.stdin) {
		if inErr != nil {
			log.Fatalf("input error: %v\n", inErr)
		}
		wtf, err := C.putc(char, C.stdout)
		fmt.Printf("%c", wtf) // this does work when redirected....
		if err != nil {
			log.Fatalf("output error: %v\n", err)
		}
	}

	// C only flushes when it detects that stdout is a terminal. So
	// this is needed when piping..
	C.fflush(C.stdout)
}
