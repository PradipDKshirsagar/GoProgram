/*
Interfaces are implemented implicitly
A type implements an interface by implementing its methods. 
There is no explicit declaration of intent, no "implements" keyword. 
*/
package main 

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}

