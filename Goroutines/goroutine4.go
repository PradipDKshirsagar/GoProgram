package main

import (
	"fmt"
	"time"
)

func main() {
	var msg = "Hello"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)
}

func sayHello() {
	fmt.Println("hello Pradip")
}