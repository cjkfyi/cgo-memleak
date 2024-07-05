package main

/*
	#include <unistd.h>
	#include <sys/types.h>
*/
import "C"
import (
	"fmt"
)

func main() {
	uid := C.getuid()
	fmt.Println("User ID:", uid)
}
