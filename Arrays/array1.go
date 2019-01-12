package main

import "fmt"

func main() {
	const i int = 5
	var a [i]string
	a[0] = "Hello"
	a[4] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{ 2}
	fmt.Println(primes)
}
