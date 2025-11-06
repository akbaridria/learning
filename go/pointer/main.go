package main

import "fmt"

func Increment(i *int) {
	*i = *i + 1
}

func main() {
	x := 10
	fmt.Println("Before increment the value of x is ", x)
	Increment(&x)
	fmt.Println("After increment the value of x is ", x)
}
