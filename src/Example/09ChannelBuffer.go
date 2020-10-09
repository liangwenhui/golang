package main

import "fmt"

func main() {

	//channel := make(chan string)
	buff_channel := make(chan string, 2) //可缓存的信道
	//阻塞，等待消费
	//channel <- "hello"
	//channel <- "world"
	//
	//fmt.Println(<-channel)
	//fmt.Println(<-channel)
	fmt.Println("---------------------")

	buff_channel <- "hello"
	buff_channel <- "world"
	fmt.Println(<-buff_channel)
	fmt.Println(<-buff_channel)

}
