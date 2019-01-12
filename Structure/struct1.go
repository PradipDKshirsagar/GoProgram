package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{10, 20})

	var s Vertex
	s.X = 4
	fmt.Println(s)

	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)

	var (
		v1 = Vertex{1, 2}  
		v2 = Vertex{Y: 1}  
		v3 = Vertex{}      
		p1 = &Vertex{1, 2} 
	)
	fmt.Println(v1, p1, v2, v3)


	sp := &v
    fmt.Println(sp.Y)

    sp.Y = 200
    fmt.Println(sp.Y)
}