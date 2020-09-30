package main

import (
	"fmt"
	"strconv"
)

/*
go不支持类型隐式转换
*/
func main() {
	var i int = 97
	var s string = "s"
	fmt.Println(string(i), s)
	fmt.Println(strconv.Itoa(i))
	formatInt := strconv.FormatInt(int64(1), 10)
	fmt.Println(formatInt)

}

type Dog struct {
}

func (v Dog) drive() {
	fmt.Println("dogdog")
}
