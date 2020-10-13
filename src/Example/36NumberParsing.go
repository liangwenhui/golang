package main

import (
	"fmt"
	"strconv"
)

func main() {
	b, _ := strconv.ParseBool("1")
	fmt.Println(b)

	f,_ :=strconv.ParseFloat("459.234",64)
	fmt.Println(f)
	//根据进制转换
	i, _ := strconv.ParseInt("10", 0, 64)
	fmt.Println(i)
	//十进制
	k,_ := strconv.Atoi("132")
	fmt.Println(k)

	_,e:= strconv.Atoi("wwqw132")
	fmt.Println(e)
}
