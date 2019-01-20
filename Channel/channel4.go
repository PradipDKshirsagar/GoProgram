package main 

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 3)
	wg.Add(2)

	go func(ch <- chan int ) {
		i := <- ch 
		j := <- ch
		fmt.Println(i, j) 
		wg.Done()
	}(ch)
	go func(ch chan <- int) {
		ch <- 42
		ch <- 27
		wg.Done()
	}(ch)
	wg.Wait()
}