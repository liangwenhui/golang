package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//不管执行多少次都是一样的随机数

	fmt.Print(rand.Intn(100),",")
	fmt.Println(rand.Intn(100))
	source := rand.NewSource(123)
	r := rand.New(source)
	fmt.Print(r.Intn(100),",")
	fmt.Println(r.Intn(100))

	source = rand.NewSource(time.Now().UnixNano())
	r = rand.New(source)
	fmt.Print(r.Intn(100),",")
	fmt.Println(r.Intn(100))
}
