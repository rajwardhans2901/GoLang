package main

import "fmt"

func main() {

    var a = "initial"
    fmt.Println(a)

	// Go will infer the type of initialized variables.
    
	var b, c int = 1, 2
    fmt.Println(b, c)

	// You can declare multiple variables at once
    
	var d = true
    fmt.Println(d)
 

    var e int = '4'
    fmt.Println(e)

	e = 899
	fmt.Println(e)

    f := "apple"   // The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case.
    fmt.Println(f)
	//'apple' throws error as single quote is used for char

	// Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
	
	var g string  // empty string
    fmt.Println(g)

	var g bool  // false
    fmt.Println(g)

}

