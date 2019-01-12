package main

import (
	"fmt"
	"math"
)

const s = "constant"


func main() {
	fmt.Println("s = ", s)

	const n = 12345
	fmt.Println("n = ", n)

	fmt.Println(math.Sin(n))
}
