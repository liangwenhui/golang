package main

import (
	"fmt"
	"time"
)

func chann(ch chan int, ch2 chan bool) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	fmt.Println("chann() over")
	ch2 <- true
}

/*
每个 case 都必须是一个通信
所有 channel 表达式都会被求值
所有被发送的表达式都会被求值
如果任意某个通信可以进行，它就执行，其他被忽略。
如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。
否则：
如果有 default 子句，则执行该语句。
如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值。
*/
func main() {
	ch := make(chan int)
	ch2 := make(chan bool)
	go chann(ch, ch2)
	for {
		select {
		case c := <-ch:
			fmt.Println("Recvice from ch :", c)
		case s := <-ch:
			fmt.Println("Recvice  :", s)
		case _ = <-ch2:
			fmt.Println("over  :")
			goto finaly
		default:
			fmt.Println("?  ")
			time.Sleep(time.Second)
		}
	}
finaly:
}
