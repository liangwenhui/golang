package main

import "fmt"

/*
运算符优先级
5 * / % << >> & &^
4 + - | ^
3 == != <= >= >
2 &&
1 ||
*/
func main() {
	var a int32 = 1
	var b int32 = 0x7fffffff
	fmt.Println(a << 31)
	fmt.Println(b)
}

func ys1() {
	var a = true && false
	var b = true || false
	fmt.Println(a, b)
}
