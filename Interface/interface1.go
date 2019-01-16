// An interface type is defined as a set of method signatures.

//A value of interface type can hold any value that implements those methods
package main 

import (
	"fmt"
	"math"
)

type MyFloat float64

type Abser interface{
	Abs() float64
}

type Vertex struct {
	X float64
	Y float64
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex)Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
	var a Abser
	f := MyFloat(math.Sqrt(3))
	v := &Vertex{3,4}

	a = f 
	fmt.Println(a.Abs())

	a = v
	fmt.Println(a.Abs())
}