package main 

import (
	"fmt"
	"sync"
	"runtime"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	runtime.GOMAXPROCS(100) // no of ope sys thred : -1 and   100 - op sys thread
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	
	fmt.Println("Hello Pradip", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	
	counter++
	m.Unlock()
	wg.Done()
}
