package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "go"
	queue <- "home"
	close(queue)
	for item := range queue {
		fmt.Println(item)
	}
}
