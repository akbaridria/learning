package main

import "fmt"

func main() {
	var String string = "oke gan disini"
	var number int = 123
	var floatnumber float64 = 0.000
	var this_is_bool bool = false
	fmt.Println("==== primitive data ====")
	fmt.Println("name: ", String, "number: ", number, "float number: ", floatnumber, "this is bool:", this_is_bool)
	fmt.Printf("name: %s, number: %d, floatNumber: %f, this is bool: %t", String, number, floatnumber, this_is_bool)
}
