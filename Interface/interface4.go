/*
If the concrete value inside the interface itself is nil, 
the method will be called with a nil receiver. 
*/
package main 

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return 
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	fmt.Printf("(%v, %T)\n", i, i)
	i.M()

	t1 := &T{"Pradip"}
	i = t1
	fmt.Printf("(%v, %T)\n", i, i)
	i.M()
}