package main

import (
	"fmt"
	"math/rand"
	//"time"
)

func main() {
	//var seedValue int64 = time.Now().UnixNano()
	//fmt.Println(seedValue)
	//rand.Seed(1)

	/*Unix returns the local Time corresponding to the given Unix time, 
	sec seconds and nsec nanoseconds since January 1, 1970 UTC. 
	It is valid to pass nsec outside the range [0, 999999999]. 
	Not all sec values have a corresponding time value. 
	One such value is 1<<63-1 (the largest int64 value). 
	*/
	i := rand.Intn(100)

	fmt.Println(i)
}