package main 

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)

	go func(ch <- chan int ) {
		i := <- ch 
		fmt.Println(i)
	//	ch <- 27 
		wg.Done()
	}(ch)
	go func(ch chan <- int) {
		i := 42
		ch <- i
	//	i = <-ch
	//	fmt.Println(i)
		wg.Done()
	}(ch)
	wg.Wait()
}