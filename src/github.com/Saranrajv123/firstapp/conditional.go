package main

import (
	"fmt"
)

func main() {
	statePopulation := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
		"New York":   19745289,
		"Ohio":       11614373,
	}

	// can,t access pop outside the loop

	if pop, ok := statePopulation["Ohio"]; ok {
		fmt.Println("inside the if condition")
		fmt.Println("pop", pop)

	}

	// simple number guessing game
	number := 50
	guess := 30

	if number < guess {
		fmt.Println("Too big")
	}

	if number > guess {
		fmt.Println("Too small")
	}

	if number == guess {
		fmt.Println("you got it")
	}

	// have >=, <=, != in condition comparision
	fmt.Println(number <= guess)

	// switch case
	// case value must be unique

	switch 7 {
	case 1, 3, 5:
		fmt.Println("odd number")
	case 2, 4, 6:
		fmt.Println("even number")
	default:
		fmt.Println("default")

	}

	// like if condition switch can also initialiser
	switch i := 5; i {
	case 1, 3, 5:
		fmt.Println("odd number")
	case 2, 4, 6:
		fmt.Println("even number")
	default:
		fmt.Println("default")

	}

	// tag less swtich statement

	// instead of break og new key called fallthrough in switch case
	// falll trhough all cases if use fallthrough

	i := 9

	switch {
	case i%2 != 0:
		fmt.Println("tag less odd number")
		fallthrough
	case i%2 == 0:
		fmt.Println(" tage less even number")
		fallthrough
	case i < 10:
		fmt.Println("yep")
		fallthrough
	default:
		fmt.Println("default")

	}
}
