package main

/*
	#include "hello.h"
*/
import "C"

func main() {
	C.wrapPrintf(C.CString("Hello World!\n"))
}
