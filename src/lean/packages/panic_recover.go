package main

import "fmt"

func main() {
	p1()
	p2()
	p3()

}

func p1() {
	fmt.Println(1)
}

//抛出异常 与 异常捕捉处理
func p2() {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	panic("throw err")
	fmt.Println(2)
}

func p3() {
	fmt.Println(3)
}
