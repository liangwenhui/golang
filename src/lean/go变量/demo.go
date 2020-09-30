package main

import (
	"fmt"
)

/*
var 变量名 类型
可以一次性声明多个变量
*/

func f1() {
	var i1 int = 1
	var i2, i3 int = 2, 3
	i1 = i1 + i2 + i3
}

//dauft
func f2() {
	var s string
	var b bool
	var i int
	fmt.Printf("%q %v %v \n", s, b, i)
}
func f3() {
	var a *int
	var b []int
	var c map[string]int
	var d chan int
	var e func(v string) int
	var f error
	fmt.Printf("%v %v %v %v %v\n", a, b, c, d, e, f)
}

//自动判断类型
func f4() {
	var a = "string"
	//a := true 编译错误
	b := true
	fmt.Println(a, b)
}

//这种写法一般用在全局变量
func f5() {
	var (
		a int  = 9
		b bool = true
	)
	fmt.Println(a, b)
}

//交换值，类型必须相同
func f6() {
	var a, b int = 1, 2
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}

func main() {
	//f2()
	//f3()
	//f4()
	f6()

}
