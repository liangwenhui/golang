package main

import (
	"fmt"
	"time"
)

func main() {
	acceptWork := func(id int, jobs chan int, res chan int) {
		for item := range jobs {
			fmt.Println("worker ", id, "accept job", item)
			time.Sleep(time.Second)
			fmt.Println("worker ", id, "finished job", item)
			res <- item
		}

	}
	nums := 5
	jobs := make(chan int, nums)
	res := make(chan int, nums)

	for w := 0; w < nums; w++ {
		go acceptWork(w, jobs, res)
	}

	for j := 1; j <= nums+1; j++ {
		jobs <- j
	}
	close(jobs)
	for j := 1; j <= nums+1; j++ {
		<-res
	}

}
