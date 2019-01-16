package main

import (
	"fmt"
	"math"
)

type ErrNegative float64

func (err ErrNegative)Error() string {
	return fmt.Sprintf("Sqrt: negative number %g", err)
}

const delta = 1e-10

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegative(f)
	}
	z := f
	for {
		n := z - (z*z-f)/(2*z)
		fmt.Println(n,z,f)
		if math.Abs(n-z) < delta {
			break
		}
		z = n
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(9))
	fmt.Println(Sqrt(-2))
}
