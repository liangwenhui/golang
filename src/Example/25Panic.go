package main

import (
	"fmt"
	"os"
)
//panic 抛出异常
func main() {
	panic("a problem")
	fmt.Println("after panic")
	_,err := os.Open("D:\\ok.txt")
	if err!=nil {
		panic(err)
	}
}
