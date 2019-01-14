/*
 A method is a function with a special receiver argument.

The receiver appears in its own argument list between the func keyword and the method name.

In this example, the Abs method has a receiver of type Vertex named v
*/
package main 

import "fmt"

type Vertex struct {
	X int 
	Y int
}

func (v Vertex) Abs() int {
	return v.X * v.X + 
}

func main() {
	
}