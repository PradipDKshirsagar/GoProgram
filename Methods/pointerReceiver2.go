//Comparing the previous two programs, you might notice that functions with a pointer 
//argument must take a pointer
/*
 There are two reasons to use a pointer receiver.

The first is so that the method can modify the value that its receiver points to.

The second is to avoid copying the value on each method call. This can 
be more efficient if the receiver is a large struct
*/


package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//value affect is depends on methd declaration i.e. (v Vertex) and (v *Vertex) not what pass to method 
func (v *Vertex)Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v1 := Vertex{3, 4}
	v1.Scale(2)
	ScaleFunc(&v1, 10)

	p1 := &Vertex{4, 3}
	p1.Scale(3)
	ScaleFunc(p1, 8)

	fmt.Println(v1, p1)

	v2 := Vertex{3, 4}
	fmt.Println(v2.Abs())
	fmt.Println(AbsFunc(v2))

	p2 := &Vertex{4, 3}
	fmt.Println(p2.Abs())
	fmt.Println(AbsFunc(*p2))
}