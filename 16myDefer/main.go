package main

import "fmt"

func main() {
	//defer executes the code at end at LIFO order
	defer fmt.Println("I'm done")
	defer fmt.Println("ONe")
	defer fmt.Println("Two")
	fmt.Println("Welcome to defer in golang")
 myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}