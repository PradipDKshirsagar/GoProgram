/*
 A method is a function with a special receiver argument.

The receiver appears in its own argument list between the func keyword and the method name.

In this example, the Abs method has a receiver of type Vertex named v
*/
package main 

import (
	"fmt"
	"math"
)

type Vertex struct {
	X float64 
	Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


func AbsNormal(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	fmt.Println("Method Function : ")
	v1 := Vertex{3, 4}
	fmt.Println(v1.Abs())

	fmt.Println("Normal Function : ")
	v2 := Vertex{3, 4}
	fmt.Println(AbsNormal(v2))
}