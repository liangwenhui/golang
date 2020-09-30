package main

import "fmt"

/*
闭包
*/
func main() {
	var f = cp1(100)
	f(100)
	f(100)
	f(100)
}

func cp1(x int) func(int) {
	return func(y int) {
		x += y
		fmt.Println(x)
	}
}
