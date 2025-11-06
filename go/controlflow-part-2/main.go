package main

import "fmt"

func main() {
	// classic
	for i := 0; i < 5; i++ {
		fmt.Println("The number is ", i)
	}

	// while
	j := 0
	for j < 5 {
		fmt.Println("The number is ", j)
		j++
	}

	// infinity
	counter := 0
	for {
		if counter > 5 {
			break
		}
		fmt.Println("The number is ", counter)
		counter++
	}

	// range
	for k := range 100 {
		fmt.Println("The number is ", k)
		k++
	}

}
