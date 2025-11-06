package main

import "fmt"

func ModifyArr(arr *[3]int) {
	arr[0] = 1
}

func LoopArr(arr *[3]int) {
	for i := range arr {
		fmt.Println("Number: ", i)
	}
}

func main() {
	arr := [3]int{0, 1, 2}
	LoopArr(&arr)
	ModifyArr(&arr)
	LoopArr(&arr)
	fmt.Println("length of an array is: ", len(arr))
}
