package main

import "fmt"

func main(){
	fmt.Println("Functions in golang!")
	greeter()


// Err - Cannot write a function inside a function
// 	func greeterTwo() {
// 		fmt.Println("Hello, again!")
// 	}

result:= adder(3,5)
proresult, message:= proAdder(3,5,6,7,8)

 fmt.Println("Result is:",result)
 fmt.Println("PeoResult is:",proresult, message)

}

func adder(a,b int) int { //int - return type
    return a+b
}

func proAdder(values ...int) (int, string){
	total:=0
	for _,value:=range values{
		total+=value
	}
	return total, "Hi Str"
}

func greeter(){
	fmt.Println("Hello, World!")
}