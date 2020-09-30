package main

import (
	"fmt"
	"reflect"
)

func main() {
	handleVar(1)
	handleVar(true)
	handleVar(T{"lwh"})

}

type T struct {
	name string
}

func handleVar(data interface{}) {
	d := reflect.ValueOf(data)
	fmt.Println(d)
	fmt.Println(d.Kind())

}
