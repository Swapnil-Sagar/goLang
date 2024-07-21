package main

import (
	"fmt"
	"net/url"
)

const myurl string = "http://localhost:3000/employees?limit=10&offset=20"

func main() {
	fmt.Println("Welcome to handling urls in golang")
    fmt.Println(myurl);

	//parsing
	result, _ := url.Parse(myurl);

	fmt.Println(result.Scheme); //http
	fmt.Println(result.Host); //localhost:3000
	fmt.Println(result.Path); //employees
	fmt.Println(result.Port()); //3000
	fmt.Println(result.RawQuery); //limit=10&offset=20

	qparams := result.Query();
	fmt.Printf("The type of query params are: %T\n",qparams); //map[limit:[10] offset:[20]]
    
	fmt.Println(qparams["limit"]);

	for _, val := range qparams {
		fmt.Println("Params are: ",val)
	}

	//Creating a url
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=swapnil",
	}

	anotherurl := partsOfUrl.String();
	fmt.Println(anotherurl);

}