package main

import "fmt"

func main() {
	fruits := map[string]int{}

	fruits["banana"] = 0
	fruits["apple"] = 1

	_, isOrangeExist := fruits["orange"]

	if !isOrangeExist {
		fmt.Println("Orange not exist")
	}

	delete(fruits, "banana")

	for k, v := range fruits {
		fmt.Printf("Name of the fruit %s with value %d", k, v)
	}
}
