package main

import "fmt"

func main() {
fmt.Println("Welcome to a class on pointers")
 
// var ptr *int 
// fmt.Println("Value of pointer is ", ptr)

myNumber := 23

var ptr = &myNumber

fmt.Println("Value of pointer is ", ptr)
fmt.Println("Value of pointer is ", *ptr)

//user * for creating a pointer and & for referencing/getting the address of the variable

*ptr = *ptr + 2

fmt.Println("New Value of pointer is ", myNumber)

}