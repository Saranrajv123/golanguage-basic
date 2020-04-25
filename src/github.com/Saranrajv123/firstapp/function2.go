package main

import (
	"fmt"
)

// anonymous function
//we are use anonymous function as variable
type greeting struct {
	greetingMessage string
	name string
}

func (g *greeting) greet() {
	fmt.Println(g.name, g.greetingMessage)
	g.greetingMessage = "welcome"
}
func main() {
	//for i := 0; i < 5; i++ {
	//	func(i int) {
	//		fmt.Println(i)
	//	}(i)
	//}

	//f := func() { // we can declare function as variable
	//	fmt.Println("Hello world")
	//}
	//f()

	//var divide func(float64, float64) (float64, error)
	//divide = func(a, b float64) (float64, error) {
	//	if b == 0.0 {
	//		return 0.0, fmt.Errorf("Cannot divide by zero")
	//	} else {
	//		return a / b, nil
	//	}
	//}
	//d, err := divide(1.0, 0.0)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(d)

	gree := greeting{
		greetingMessage: "Hello",
		name: "go lang",
	}

	gree.greet()
	fmt.Println("The greeting Message: ", gree.greetingMessage)



}
