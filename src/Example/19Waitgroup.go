package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	accept := func(id int, wg *sync.WaitGroup) {
		fmt.Printf("worker %d accept job \n", id)
		time.Sleep(time.Second)
		fmt.Printf("worker %d done job \n", id)
		wg.Done()
	}
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go accept(i, &wg)
	}
	wg.Wait()
	fmt.Println("done")

}
