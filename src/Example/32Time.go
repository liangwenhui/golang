package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now)
	that := time.Date(
		2020,12,24,
		00,00,00,00,
		time.Local)
	fmt.Println(that)
	fmt.Println(that.Weekday())
	fmt.Println(that.Before(now))
	fmt.Println(that.Equal(now))
	diff := now.Sub(that)
	fmt.Println(diff)
	fmt.Println(now.Add(-diff))


}
