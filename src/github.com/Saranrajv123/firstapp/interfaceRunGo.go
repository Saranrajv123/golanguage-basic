package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type rectangle struct {
	width, height float64
}

func (r rectangle) Area() float64 {
	return r.width * r.height
}

func (r rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type circle struct {
	radius float64
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	var s Shape
	s = rectangle{
		width:  5.0,
		height: 4.0,
	}
	fmt.Println(s, "s")
	fmt.Println(s.Area(), "s")
	fmt.Println(s.Perimeter(), "s")

	s = circle{radius: 10}

	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())

}
