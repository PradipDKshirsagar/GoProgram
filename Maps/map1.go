//Like slices, maps hold references to an underlying data structure. 
//If you pass a map to a function that changes the contents of the map, 
//the changes will be visible in the caller.
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}



func main() {
	var m1 = make(map[string]Vertex)
	m1["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m1["Bell Labs"])

	fmt.Println(m1)

	m2 := make(map[string]int) 

	m2["Answer"] = 42
	fmt.Println("The value:", m2["Answer"])

	m2["Answer"] = 48
	fmt.Println("The value:", m2["Answer"])

	delete(m2, "Answer")
	fmt.Println("The value:", m2["Answer"])

	v, ok := m2["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
