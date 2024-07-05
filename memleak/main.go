package main

//#include "printer.h"
import "C"
import "encoding/json"

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
