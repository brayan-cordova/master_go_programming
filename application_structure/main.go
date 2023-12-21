// main is a special name declaring an executable rather than a library (package)
package main

// import declaration declares packages used in this file
import (
	"fmt"
)

// package scoped variables and constants
const secondsInHour = 3600

// a function declaration. main is a special function name
// it is the entry point for the executable program
// Go uses brace brackets to delimit code blocks
func main(){
	
	fmt.Println("Hello Go World...!!")
	distance := 60.8 // distance in km
	fmt.Printf("The distance in miles is %f \n", distance * 0.621)
}