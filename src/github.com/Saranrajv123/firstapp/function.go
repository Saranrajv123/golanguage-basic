package main

import (
	"fmt"
)

func main() {
	greeting := "hello"
	name := "saranraj"

	sayGreeting(&greeting, &name)
	fmt.Println(name, "inside main function")

	// sum(["saran", "raj", 1, 2, 3, 4, 5, 6]) /* find this how to do this way */
	value := sum(1, 2, 3, 4, 5, 6)
	fmt.Println("values", value)

	sub := sub(3, 4, 5, 6, 7, 8)
	fmt.Println(sub, "sub")

	div, err := divideHandlingFunc(4.0, 4.0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(div, "div")

}

// if ypu want to return a function  value than
// we need tell to function what type function should return

func sum(values ...int) int { /* specific return type */
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

func sub(values ...int) (result int) { /* this way you can return value explicity */
	for _, v := range values {
		result -= v
	}

	return /* no need specific the return varaible in return - result  */
}

// func sum(values ...int) { /*  veridict parameter like spread operators */
// 	fmt.Println(values, "values")
// 	result := 0
// 	for _, v := range values { /* if you are not use _ in for loop it takes only 5 values  */
// 		result += v
// 		fmt.Println(v, "v")
// 	}

// 	fmt.Println(name, result, " name, result")

// }

func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "go lang"
	fmt.Println(*name, "name")
}

func divideHandlingFunc(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by 0 ")
	}
	return a / b, nil
}
