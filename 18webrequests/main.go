package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://gorest.co.in/public/v2/users";

func main() {
	fmt.Println("Welcome to web requests in golang");

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response);

	defer response.Body.Close() // closing the connection is important

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content);
}
