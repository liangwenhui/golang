package main

import (
	"fmt"
	"os"
)

type stu2 struct {
	age int
	name string
}

func main() {


	var pf = func (format string,i *int,s ...interface{}){
		*i++
		fmt.Printf(format,*i,s)
	}
	var d = stu2{17,"kkop"}
	 i := 0
	pf("%v %v \n",&i,d)
	 //打印 +v k v
	pf("%v %+v \n",&i,d)
	 //打印 #v type{data}
	pf("%v %#v \n",&i,d)
	 //T type
	pf("%v %T \n",&i,d)
	 //t 格式化 bool
	pf("%v %t \n",&i,true)
	 //d 格式化 整形 10进制
	pf("%v %d \n",&i,0x12)
	// b 格式输出 二进制
	pf("%v %b \n",&i,6)
	// c 输出整形对应字符
	pf("%v %c \n",&i,33)
	// x 输出16进制
	 pf("%v %x \n",&i,16*16-1)
	 pf("%v %x \n",&i,"16精致")
	 // f 格式化浮点型
	 pf("%v %f \n",&i,123.4)
	 // E e 科学计数法
	 pf("%v %e \n",&i,123.4)
	 pf("%v %E \n",&i,123.4)
	 // s 替换字符串
	 pf("%v %s \n",&i,"\"zifuc\"")
	 // q 整个字符串打印，外加双引号
	 pf("%v %q \n",&i,"\"zifuc\"")
	 // p 输出指针 不用p输出的是构造体
	 pf("%v %p \n",&i,&d)
	 i++
	 //格式化 表格输出
	 fmt.Printf("%v |%12d|%6d| \n",i,12,15)
	 // 浮点型12.2 12表示宽，.2表示保留精度
	 fmt.Printf("%v |%6.2f|%12.2f| \n",i,13.1,16.1367)
	 // -号左对齐
	 fmt.Printf("%v |%-6.2f|%-12.2f| \n",i,13.1,16.1367)
	 fmt.Printf("%v |%-6s|%-12s| \n",i,"food","see")

	 fmt.Fprintf(os.Stderr,"an %s \n" , " bad request")


}
