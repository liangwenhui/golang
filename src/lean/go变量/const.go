package main

import (
	"fmt"
	"unsafe"
)

const (
	A int = 1
	B int = 2
)

const (
	C = "abc"
	D = len(C)
	E = unsafe.Sizeof(C)
)

func const1() {
	const F1 = "f1"
}
func main() {
	fmt.Println(C, D, E)
}
