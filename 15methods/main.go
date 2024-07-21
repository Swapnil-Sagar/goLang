package main

import "fmt"

func main() {

	fmt.Println("Welcome to structs in golang")
	// no inheritance in golang; No super or parent


	swapnil := User{"Swapnil", "6kS5s@example.com", true, 25}

	fmt.Println(swapnil);
	fmt.Printf("details are %+v\n", swapnil);
	fmt.Printf("name is %v and email is %v.\n", swapnil.Name, swapnil.Email);
    swapnil.GetStatus()
    swapnil.NewMail()
}

type User struct {
	Name 	string
	Email 	string
	Status 	bool
	Age 	int

}


//method
func (u User) GetStatus(){
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail(){
	u.Email = "newmail@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}