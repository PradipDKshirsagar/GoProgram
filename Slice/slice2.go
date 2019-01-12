package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s1 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s1)

	//When slicing, you may omit the high or low bounds to use their defaults instead. 

	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s,len(s),cap(s))
	
	s = s[1:4]
	fmt.Println(s,len(s),cap(s))

	s = s[:2]
	fmt.Println(s,len(s),cap(s))

	s = s[1:]
	fmt.Println(s,len(s),cap(s))
}
