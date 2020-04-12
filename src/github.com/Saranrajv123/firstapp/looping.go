package main

import (
	"fmt"
)

func main() {
	// for i := 0; i <= 5; i++ {
	// 	fmt.Println(i)

	// 	if i%2 == 0 {
	// 		// i = i / 2
	// 		fmt.Println(i, "vefore inside if ")
	// 		i /= 2
	// 		fmt.Println(i, "inside if ")
	// 	} else {
	// 		fmt.Println(i, "after inside if ")
	// 		i = 2*i + 1
	// 		fmt.Println(i, "inside else")
	// 	}
	// }

	// for i, j := 0, 1; i < 10; i, j = i+1, j+1 {
	// 	fmt.Println(i, j)
	// }

	// global initailizer
	// you can remove semicolon and increment (i++) from loop it's work as well - infinite loop
	// we can also access i vlaue outside scope
	// for i < 5 { do something } this is more like while
	// if you remove 1 < 5 in for loop again it's keep on run - infinite loop

	// i := 0
	// for i < 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// fmt.Println(i, "uotside scope")

	// for {
	// 	fmt.Println("indise for loop", i)
	// 	i++
	// 	if i == 3 {
	// 		fmt.Println("inside if condition", i)
	// 		break
	// 	}
	// }

	// nested loops
	// Use Loop keyword to break parent loop below example use Lopp to break i
	// inside j break statement only break the j value as soon enter 3

	// Loop:
	// 	for i := 0; i < 5; i++ {
	// 		fmt.Println(i, "inside i loop")
	// 		for j := 0; j < 5; j++ {
	// 			fmt.Println("inside j", j)
	// 			fmt.Println(i * j)
	// 			if i*j == 3 {
	// 				break Loop
	// 			}

	// 		}
	// 	}

	// loop in collection, array, string, map,
	// with help range key word
	// remove value in for it only print keys of array
	// remove key in loop it's cause error to avaiod use _ operator so that we can get only value's from array

	// arrayaValue := []int{2, 3, 4, 5, 6}
	// arrayaValue := [10]int{2, 3, 4, 5, 6}
	// arrayaValue := [...]int{2, 3, 4, 5, 6}
	arrayaValue := "string looping"

	statePopulation := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
		"New York":   19745289,
		"Ohio":       11614373,
	}

	for k, v := range arrayaValue {
		fmt.Println("k and v", k, string(v))
	}

	for key, value := range statePopulation {
		fmt.Println("state and population", key, value)
	}
}
