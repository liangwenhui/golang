package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())

	timestamp := now.Unix()
	fmt.Println(timestamp)
	ntimestamp := now.UnixNano()
	fmt.Println(ntimestamp)
	//格式化日期，必须使用2006-01-02 15:04:05 数据，修改样式即可，很奇怪，可能这是纪念go出生
	fmt.Println(now.Format("2006-01-02 15:04:05"))

}
