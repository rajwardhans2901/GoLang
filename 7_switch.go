package main
import (
    "fmt"
)
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
	default:
		fmt.Println("It's a weekday")
    }
	
	var num int = 10
	
	switch {
	case num < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	
	// n:= 10
	
	// switch n {
	// case n<11, n> 0:
	// 	fmt.Println("Positive under 10")
	// case n>10:
	// 	fmt.Println("Positive over 10")
	// }

	n := 10

	switch {
		case n < 11, n > 0:
			fmt.Println("Positive under 10")
		case n > 10:
			fmt.Println("Positive over 10")
	}

	n = 10

	switch {
		case n < 11 && n > 0:
			fmt.Println("Positive under 10")
		case n > 10:
			fmt.Println("Positive over 10")
	}
}