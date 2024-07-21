package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to slices")
	
	var fruitList = []string{"Apple", "Orange", "Mango", "Grapes"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Peach", "Banana")
	fmt.Println("Fruit list is:", fruitList)

	fruitList = append(fruitList[1:3])
	//last range is non inclusive
	fmt.Println(fruitList)

	//
	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 456
	highScores[3] = 567
	// highScores[4] = 777; // -error

	highScores = append(highScores, 555, 666, 321)

	fmt.Println(highScores)

	
	fmt.Println("IsSorted", sort.IntsAreSorted(highScores))
	sort.Ints(highScores);
	fmt.Println("Sorted", highScores)


	//removing a value from slice based on index
	var courses = []string{"react", "js", "go", "php",	"python"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)


}