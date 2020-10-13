package main

import (
	"fmt"
	"testing"
)
// go test -v
func Test(t *testing.T){
	i :=0
	t.Fail()//直接退出测试
	i++
	t.Error("2 err",i)//失败，但继续执行

}
func Test2(t *testing.T){
	i :=0

	i++
	fmt.Println(i)

}
