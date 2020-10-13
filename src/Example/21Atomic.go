package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var i int32
	var wg sync.WaitGroup
	for j:=0;j<50;j++ {
		wg.Add(1)
		go func(){
			for k:=0;k<1000;k++ {
				//i++
				atomic.AddInt32(&i,1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(i)

}
