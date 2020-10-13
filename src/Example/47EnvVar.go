package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("jmx","true")
	fmt.Println(os.Getenv("jmx"))
	fmt.Println(os.Getenv("JAVA_HOME"))
	for _,env := range os.Environ(){
		fmt.Println(env)
	}

}
