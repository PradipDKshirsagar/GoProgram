/*
Go programs express error state with error values.

The error type is a built-in interface similar to fmt.Stringer: 
type error interface {
    Error() string
}
*/

package main 

import (
	"fmt"
	"time"
	"strconv"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError)Error() string{
	return fmt.Sprintf("%v %v", e.When,e.What)
}

func main() {
	var e = &MyError{time.Now(),"pradip"}
	if e != nil {
		fmt.Println(e)
	}
	i, err := strconv.Atoi("42")
	if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
  		return
	}
	fmt.Println("Converted integer:", i)
}