package main

import "fmt"

/**
go 中range关键字用于for循环中迭代数组（array），切片，channel
map 的元素range返回index,value
*/
func ranged() {
	kv := map[string]string{"a": "aa", "b": "bb"}
	kv2 := map[string]int{"a": 1, "b": 2}
	for k, v := range kv {
		fmt.Println(k, v)
	}
	for k, v := range kv2 {
		fmt.Println(k, v)
	}

	for i, v := range "abc" {
		fmt.Println(i, v)
	}
}

func mapd() {

	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	fmt.Println(m)
	delete(m, "a")
	fmt.Println(m)
}

func main() {
	mapd()

}
