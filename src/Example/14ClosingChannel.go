package main

import (
	"fmt"
)

func main() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			job, more := <-jobs
			//fmt.Println(more)
			if more {
				fmt.Println("res job", job)
			} else {
				fmt.Println("no jobs to res")
				done <- true
			}
		}
	}()
	for i := 0; i < 5; i++ {
		jobs <- i
		fmt.Println("sent job ", i)
	}

	fmt.Println("jobs closed")

	close(jobs)

	<-done

}
