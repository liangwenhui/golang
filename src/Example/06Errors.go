package main

import (
	"errors"
	"fmt"
)

func main() {

	f1 := func(i int) (int, error) {
		if i == 100 {
			return -1, errors.New("overload")
		} else {
			return i - 1, nil
		}
	}
	var r, e = f1(100)
	if e != nil {
		fmt.Println("error:", e)
	} else {
		fmt.Println("success:", r)
	}
	r, e = f1(99)
	if e != nil {
		fmt.Println("error", e)
	} else {
		fmt.Println("success:", r)
	}

	f2 := func(i int) (int, error) {
		if i == 100 {
			return -1, &argError{i, "overload"}
		} else {
			return i - 1, nil
		}
	}

	f2(99)
	r, e = f2(99)
	if e != nil {
		fmt.Println("error", e)
	} else {
		fmt.Println("success:", r)
	}
	r, e = f2(100)
	if e != nil {
		fmt.Println("error", e)
	} else {
		fmt.Println("success:", r)
	}

}

type argError struct {
	code int
	msg  string
}

//实现 Error方法
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.code, e.msg)
}
