package main

import (
	"fmt"
	"time"
)

//时间戳
func main() {
	now := time.Now()
	secs := now.Unix()
	nacos := now.UnixNano()
	millis := nacos/ 1000000
	fmt.Println(now)
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nacos)
	fmt.Println(time.Unix(secs,0))
	fmt.Println(time.Unix(0,nacos))

}
