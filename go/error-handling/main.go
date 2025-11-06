package main

import (
	"errors"
	"fmt"
)

func Devide(a int, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot devided by zero")
	}

	return float64(a) / float64(b), nil
}

func main() {

	res, err := Devide(5, 3)
	if err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		fmt.Printf("Result: %f", res)
	}

}
