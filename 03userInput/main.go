package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to user input"
	fmt.Println(welcome);

	//user input
	reader := bufio.NewReader(os.Stdin);
	fmt.Println("Enter rating");

	//comma ok || err ok
	input, _ := reader.ReadString('\n');
	//reading as soon as hit /n

	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input)

} 