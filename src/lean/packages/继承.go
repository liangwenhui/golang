package main

/*
 	Go 的继承是结构体嵌入，一个结构体可以嵌入多个结构体
但是多个结构体不能拥有相同的方法
*/
import (
	"fmt"
	"strconv"
)

type Computer struct {
	cpu string
}
type Computer2 struct {
	cpu string
}

func (c Computer2) compute(v1, v2 int) *Computer2 {
	c.cpu = strconv.Itoa(v1) + "核" + strconv.Itoa(v2) + "线程 CPU"
	return &c
}
func (c Computer) compute(v1, v2 int) *Computer {
	fmt.Println("ccccc")
	c.cpu = strconv.Itoa(v1) + "核" + strconv.Itoa(v2) + "线程 CPU"
	return &c
}

//func  (c NetComputer)compute(v1, v2 int) *NetComputer {
//	c.cpu = strconv.Itoa(v1)+"核"+strconv.Itoa(v2)+"线程 CPU"
//	return &c
//}
type NetComputer struct {
	Computer
	//Computer2
	ipconifg string
}

func main() {
	c := new(Computer)
	c = c.compute(2, 8)
	fmt.Println(c.cpu)

	nc := new(NetComputer)
	nc.ipconifg = "eth0"
	c = nc.compute(4, 12)
	fmt.Println(nc)
	fmt.Println(c)
}
