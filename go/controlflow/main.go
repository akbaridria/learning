package main

import "fmt"

func PrintDay(day string) {
	switch day {
	case "Monday":
		fmt.Println("Today is Monday")
	case "Tuesday":
		fmt.Println("Today is Tuesday")
	default:
		fmt.Println("Unknwon day")
	}

}

func main() {

	// integer if conditional things
	age := 30

	if age > 20 {
		fmt.Printf("you arge is %d and you are old \n", age)
	} else {
		fmt.Println("you are still young")
	}

	if x := 0; x > 20 {
		fmt.Printf("yor age is %d and you are old", x)
	} else {
		fmt.Println("you are young")
	}

	// switch statement
	PrintDay("Monday")
	PrintDay("Tuesday")
	PrintDay("Wednesday")
}
