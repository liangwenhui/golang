package main

import (
	"fmt"
)

/*
定义切片
var identifier []type
或者使用make函数创建
:= make([]type,len)
:= make([]type,len,capacity)
*/
func main() {
	s := []int{1, 2, 3}
	ss := s[:] //copy 把元素指针复制
	fmt.Println(&ss[1])
	fmt.Println(&s[1])
	ss2 := s[1:2]
	fmt.Println(ss2)
	fmt.Println(&ss2)
	//len() cap() 方法
	var _ = len(s)
	var _ = cap(s)
	//var a interface{} = []int{1,2,4}
	//getType(a)
	fmt.Println("----------------------------")
	var numbers []int = []int{1, 2, 3, 4, 5}
	fmt.Println(&numbers[1])
	var _ = append(numbers, 6)
	//numbers = append(numbers,6) //数组append的返回结果是另一个相同值，不同地址的数组
	fmt.Println(numbers)
	fmt.Println(&numbers[1])

	number2 := make([]int, 2)
	//copy(number2,numbers)
	fmt.Println(number2)
	fmt.Println(&number2[1])
	var _ = append(number2, 3)
	//number2[4] = 5 //报错
	fmt.Println(number2)
	fmt.Println("---------------")
	var number3 []int
	fmt.Println(number3)
	var _ = append(number3, 1, 2, 3, 4, 5)

}
func getType(v interface{}) {
	switch i := v.(type) {
	default:
		fmt.Println(i)
	}

}
