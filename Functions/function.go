package main 

import "fmt"

func check(a, b int) {
	if a == b {
		fmt.Println("a and b are equal and values are a = ", a, " b= ", b)
	}else{
		fmt.Println("a and b are different")
	}
}

func swap(a, b float64) (x, y float64) {
	x = b
	y = a
	return 
}	

func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {
	a, b := 10, 20
	x, y := 30, 30 
	check(a, b)

	check(x, y)

	p, q := 10.242, 30434.2
	fmt.Println(p, q)
	p, q = swap(p ,q)
	fmt.Println(p, q)

	sum(1, 2)
    sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
    sum(nums...)
}