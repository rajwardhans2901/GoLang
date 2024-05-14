package main

import (
    "fmt"
)

// const declares a constant value.

// const statement can appear anywhere a var statement can.
const s string = "constant"

func main() {
    fmt.Println(s)

    const d = 4

	// d = 1  //This throws error as constant cant be reassigned

    fmt.Println(d) 

    
}