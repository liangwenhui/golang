package main

import "fmt"

func main() {

	for {
		goto that
	}

	for i := 0; i < 10; i++ {

	}
that:
	strings := []string{"a", "b"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
	fmt.Println(strings[0])
}

type F struct {
	d D
}
type D struct {
}
