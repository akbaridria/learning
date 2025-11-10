package main

import (
	"fmt"
	"testing"
)

func TestArea(t *testing.T) {
	tests := []struct {
		shape                Shape
		name                 string
		expected_area_amount float64
	}{
		{name: "Test rectangle", shape: &Rectangle{width: 10, height: 10}, expected_area_amount: 100},
		{name: "Test circle", shape: &Circle{radius: 10}, expected_area_amount: 161},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			fmt.Println("got: ", got, test.expected_area_amount)
			if got < test.expected_area_amount {
				t.Errorf("%s: got %f expected %f", test.name, got, test.expected_area_amount)
			}

		})
	}

}
