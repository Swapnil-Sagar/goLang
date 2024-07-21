package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err!= nil {
        panic(err)
    }
	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder
	content,_ :=io.ReadAll(response.Body)

	byteCount,_ := responseString.Write(content)

	fmt.Println("Byte count is: ",byteCount)
	fmt.Println(responseString.String())

	//OR

	// fmt.Println(string(content))

}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	//fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename": "Let's go with golang",
			"price": 0,
			"platform": "golang bootcamp"
		}`)

		response, err := http.Post(myurl, "application/json", requestBody)
		if err!= nil {
			panic(err)
		}
		defer response.Body.Close()
		content,_ :=io.ReadAll(response.Body)
		fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

    // form data

	data := url.Values{}
	data.Add("firstname", "Swapnil")
	data.Add("lastname", "Sagar")
	data.Add("email", "<EMAIL>")
	response,err :=http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content,_ :=io.ReadAll(response.Body)
	fmt.Println(string(content))
}