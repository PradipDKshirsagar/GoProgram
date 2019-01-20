package main 

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50)
	wg.Add(2)
	go func(ch <- chan int ) {
		for {
			if i, ok := <- ch ; ok {
				fmt.Println(i) 	
			} else {
				break
			}
		}
		wg.Done()
	}(ch)
	go func(ch chan <- int) {
		ch <- 42
		//close(ch) // deadlock
		ch <- 27
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}
//ok syntax to not go for pani with condition check

//if you are not in a loop