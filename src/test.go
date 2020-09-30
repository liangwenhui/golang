package main

import (
	"./myMath"
	"fmt"
)

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("hellop")
	var i int = mathClazz.Add(1, 1)
	fmt.Println(i)
}

func one() {

}

func tow(v string) {
	fmt.Println(v)
}
