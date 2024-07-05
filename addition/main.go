package main

/*
	#include "math.h"
	double add(double a, double b) {
		return a + b;
	}
*/
import "C"
import "fmt"

func main() {
	fmt.Println(C.floor(C.add(1336.3, 1)))
}
