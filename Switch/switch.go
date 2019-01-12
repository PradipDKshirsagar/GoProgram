package main 

import "fmt"

func main() {
	i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    a := 3
    b := 5

    switch {
    	case a == 4:
    		fmt.Println("a == 4")
    	case b == 5:
    		fmt.Println("b == 5") 
    }

	whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("whatAmI is bool")
        case int:
            fmt.Println("whatAmI is int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
	whatAmI(true)
    whatAmI(1)
    whatAmI("hey")

}