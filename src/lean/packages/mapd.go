package main

import "fmt"

/**
切片当量用完，会双倍扩大
*/
func main() {
	slice := make([]map[string]int, 5, 10)
	slice0 := make([]map[string]int, 0, 10)
	fmt.Println(slice)
	fmt.Println(slice0)
	slice = append(slice, make(map[string]int), make(map[string]int), make(map[string]int), make(map[string]int))
	fmt.Println(slice)
	fmt.Println(cap(slice))
	slice[0] = make(map[string]int)
	slice[0] = map[string]int{}
	slice[0]["a"] = 1

}
