//pass array to functions

package main 

import "fmt"

func Sum(a *[3]float64) (sum float64) {
    for _, v := range *a {
        sum += v
    }
    return
}

func main() {
	array := [...]float64{7.0, 8.5, 9.1} // Another Declartion
	x := Sum(&array)	
	fmt.Println(x)
}