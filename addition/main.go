package main

/*
	#include "add.h"
*/
import "C"
import "fmt"

func main() {
	fmt.Println(C.addition(1336.3, 1))
}
