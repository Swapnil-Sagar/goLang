package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in the file - file1.txt"

	file, err := os.Create("./filesCode.txt")

	if err != nil {
		panic(err) //panic shuts down the program execution
	}
	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Printf("length is\n", length)
	defer file.Close()

	readFile("./filesCode.txt")
}

func readFile(filname string) {
	databyte, err :=ioutil.ReadFile(filname)
	checkNilError(err)
	fmt.Println(string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err) //panic shuts down the program execution
	}
}