package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//Normal variable receiver
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


//Pointer receiver for changing value of original Vertex v
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
