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
		// this does not work when being redirected..
		// but does not work when run normally, but does not write to screen though..
		wtf, err := C.putc(char, C.stdout)
		fmt.Printf("%c", wtf) // this does work when redirected....
		if err != nil {
			log.Fatalf("output error: %v\n", err)
		}
	}
}
