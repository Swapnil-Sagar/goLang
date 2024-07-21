package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days);

	//for loop
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	//OR

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	//OR

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 2 {
			goto dev;
		}

		if rogueValue == 5 {
			break;
		}
		if rogueValue == 3 {
			rogueValue++
			continue
		}

		fmt.Println("Rogue Value is: ", rogueValue)
		rogueValue++
	}

	dev:
	 fmt.Println("Jumping at DEV")

}