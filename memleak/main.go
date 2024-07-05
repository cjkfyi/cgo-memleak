package main

import (
	"encoding/json"
)

/*
	#include "printer.h"
*/
import "C"

//export EncodeF64
func EncodeF64(key *C.char, value float64) *C.char {
	m := map[string]float64{
		C.GoString(key): value,
	}
	b, _ := json.Marshal(&m)
	// Memory leak example here
	return C.CString(string(b))
}

func main() {
	C.printPiJSON()
}
