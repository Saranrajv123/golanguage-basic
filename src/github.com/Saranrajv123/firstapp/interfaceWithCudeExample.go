package main

import (
	"fmt"
)

type shape interface {
	Area() float64
}

type object interface {
	volume() float64
}

type cube struct {
	side float64
}

func (c cube) Area() float64 {
	return 6 * (c.side * c.side)
}

func (c cube) volume() float64 {
	return c.side * c.side * c.side
}

func main() {
	var sh shape
	sh = cube{side: 3}

	fmt.Println(sh.Area(), "sh")

	var obj object
	obj = cube{side: 3}

	fmt.Println(obj.volume())

}
