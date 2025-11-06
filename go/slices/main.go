package main

import "fmt"

func main() {
	slc := []int{}

	slc = append(slc, 0)
	slc = append(slc, 1)

	for i := range slc {
		fmt.Println("Index: ", i)
	}

	slc_two := slc[1:]

	fmt.Printf("slc two: %v \n", slc_two)
	fmt.Printf("slc two: %#v", slc_two)
}
