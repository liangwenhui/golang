package main

import (
	"fmt"
	"os"
)

func main(){

	var f1 = func(word string) string{
		fmt.Println(word)
		return "ok"
	}
	fmt.Println(f1)
	addBody(f1)
	var args interface{} = os.Args
	fmt.Println(args)
	
}

func addBody(f func(v string) string){
	f("take");
}