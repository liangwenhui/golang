package main

import (
	"encoding/base64"
	"fmt"
)

//e8bd767b5aed8bd0e18b37f59a186bc056e78902
func main() {
	s := "sha1 aaaa bbb string"
	b := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(b)
	decodeString, _ := base64.StdEncoding.DecodeString(b)
	fmt.Println(string(decodeString))

}
