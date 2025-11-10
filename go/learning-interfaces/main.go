package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return math.Phi * c.radius * c.radius
}

func main() {
	reactange := &Rectangle{width: 10, height: 10}
	circle := &Circle{radius: 10}

	rectangle_area := reactange.Area()
	fmt.Println("recrtangle area: ", rectangle_area)

	circle_area := circle.Area()
	fmt.Println("circle area: ", circle_area)

}
