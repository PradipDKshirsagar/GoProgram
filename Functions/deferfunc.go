/*it guarantees that you will never forget to close the file, a mistake that's 
easy to make if you later edit the function to add a new return path. Second, 
it means that the close sits near the open, which is much clearer than 
placing it at the end of the function. 
*/
package main 

import ("fmt"
		"reflect"
		)
func reverseArray(arr []int) {
	for _, x  := range arr {
		defer fmt.Println(x)
	}
	//fmt.Println("complete")
}

func printArray(arr []int) {
	for _,x := range arr {
		fmt.Println(x)
	}
	//fmt.Println("complete")
}

func main() {
	x := []int{1,3,5,7}	//it is slice 
	x = append(x,9)
	fmt.Println(reflect.TypeOf(x))
	fmt.Println("Array in same order : ")
	printArray(x)

	fmt.Println("Array in reverse order : ")
	reverseArray(x)
}