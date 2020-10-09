package main

import "fmt"

func main() {
	i := 0
	for i <= 3 {
		fmt.Println("f1:", i)
		i++
	}

	for j := 0; j <= 3; j++ {
		fmt.Println("f2", j)
	}

	for {
		fmt.Println("loop")
		break
	}

}
