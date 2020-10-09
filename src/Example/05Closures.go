package main

import "fmt"

//闭包
func main() {
	seq := func(i int) func() int {
		return func() int {
			i++
			return i
		}
	}

	next := seq(0)
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

}
