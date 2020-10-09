package main

import "fmt"

func main() {

	i := 1
	switch i {
	case 0:
		fmt.Println("命中0")
	case 1:
		fmt.Println("命中1")
		fallthrough
	case 2:
		fmt.Println("命中2")
	case 3:
		fmt.Println("命中3")
	default:
		fmt.Println("未命中")

	}

}
