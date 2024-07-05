package main

/*
	#include "stdio.h"
	void wrapPrintf(const char *s) {
		printf("%s", s);
	}
*/
import "C"

func main() {
	C.wrapPrintf(C.CString("Hello World!\n"))
}
