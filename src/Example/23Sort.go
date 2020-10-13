package main

import (
	"fmt"
	"sort"
)

func main() {

	var arr []int = []int{2,6,8,1,3,5,4}
	sort.Ints(arr)
	fmt.Println(arr)

}
