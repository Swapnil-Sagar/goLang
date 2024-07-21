package main

import "fmt"

func main() {

	fmt.Println("Welcome to structs in golang")
	// no inheritance in golang; No super or parent


	swapnil := User{"Swapnil", "6kS5s@example.com", true, 25}

	fmt.Println(swapnil);
	fmt.Printf("details are %+v\n", swapnil);
	fmt.Printf("name is %v and email is %v", swapnil.Name, swapnil.Email);
}

type User struct {
	Name 	string
	Email 	string
	Status 	bool
	Age 	int
}