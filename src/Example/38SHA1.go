package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {

	s := "sha1 aaaa bbb string"
	h := sha1.New()
	h.Write([]byte(s))
	bs:=h.Sum(nil)
	fmt.Printf("%x",bs)

}
