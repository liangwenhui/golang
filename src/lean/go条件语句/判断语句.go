package main

import "fmt"

func main() {
	//var b bool = true
	//if(b){
	//	//
	//}else{
	//	//
	//}
	switchd("1")
	//switchd("2")
	//switchd("d")
	fallthroughd(2)
	var i int
	typeswitch(i)
	var a int8 = 1
	typeswitch(a)
	var b int32 = 1
	typeswitch(b)

}

/**
type switch
*/
func typeswitch(v interface{}) {
	switch i := v.(type) {
	case nil:
		fmt.Println("v 的类型是：%T", i)
	case int:
		fmt.Println("v 的类型是：int", i)
	case float64:
		fmt.Println("v 的类型是：float64", i)
	case func():
		fmt.Println("v 的类型是：func", i)
	case bool, string:
		fmt.Println("v 的类型是：bool,string", i)
	default:
		fmt.Println("v 的类型是：unknow")
	}
}

/**
fallthrough 命中会无条件执行下一个case
*/
func fallthroughd(v int) {
	switch v {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
		fallthrough
	case 4:
		fmt.Println("4")
	case 5:
		fmt.Println("5")
		fallthrough
	default:
		fmt.Println("d")
	}

}

func switchd(v string) {
	switch v {
	case "1":
		fmt.Println(1)
	case "2":
		fmt.Println(2)
		//break 不需要break
	case "3":
		fmt.Println(3)
	default:
		fmt.Println("def")

	}
}
