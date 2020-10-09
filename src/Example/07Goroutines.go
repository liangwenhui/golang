package main

import (
	"fmt"
	"time"
)

func main() {
	f := func(s string) {
		for _, char := range s {
			fmt.Println(char)
		}
	}
	f("12345")
	go f("12345")
	go func(s string) {
		fmt.Println(s)
	}("going")
	time.Sleep(time.Second)
}
