package main

import "fmt"

func main() {
	sum := func(nums ...int) int {
		fmt.Println("nums:", nums)
		total := 0
		for _, v := range nums {
			total += v
		}
		return total
	}
	fmt.Println(sum(1, 1, 1))
	fmt.Println(sum(2, 2, 2))
	nums := []int{3, 3, 3}
	fmt.Println(sum(nums...))

}
