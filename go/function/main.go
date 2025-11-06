package main

import (
	"errors"
	"fmt"
)

func Devide(a int, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("can't devided by zero")
	}

	result := float64(a) / float64(b)

	return result, nil
}

func main() {
	res, err := Devide(3, 4)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("Result: %f", res)
	}
}
