package main 

import (
	"fmt"
	"math"	
)

type geometry interface {
	area() float64
	peri() float64
}

type rect struct {
	length float64
	breadth float64
}

type circle struct {
	radius float64
}

func (r *rect)area() float64 {
	return r.length * r.breadth 
}

func (r *rect)peri() float64 {
	return 2 * (r.length  + r.breadth)
}

func (c *circle)area() float64 {
	return math.Pi * c.radius * c. radius 
}

func (c *circle)peri() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	var g geometry
	var r1 = rect{3, 4}
	var c1 = circle{5}

	g = r1
	fmt.Println(g)
	fmt.Println(g.area(), g.peri())

	g = &c1 
	fmt.Println(g)
	fmt.Println(g.area(), g.peri())
}