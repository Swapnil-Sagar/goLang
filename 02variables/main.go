package main

import "fmt"

const LoginToken string = "login"; //Public

//If first letter is captial of a variable then it's a public variable  

func main() {
	var username string = "Swapnil Variable"
	fmt.Println("Hello, " + username)
	fmt.Printf("type: %T \n", username)
	

	//implicit type
	var name = "Swapnil Variable";
	fmt.Println(name)

	//no var style
	noOfUsers := 1000.123
	fmt.Println(noOfUsers)


	//you can only use : inside a method like above
	
	fmt.Println(LoginToken);
}