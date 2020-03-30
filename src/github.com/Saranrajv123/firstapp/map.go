package main

import (
	"fmt"
)

func main() {
	// other way to create a map
	// statePopulation := make(map[string]int)
	// creating map
	statePopulation := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
		"New York":   19745289,
		"Ohio":       11614373,
	}

	fmt.Println("statePopulation before add georgia", statePopulation)

	statePopulation["Georgia"] = 10371
	fmt.Println("statePopulation after added georgia", statePopulation)

	fmt.Println("statePopulation before Delete", statePopulation)

	delete(statePopulation, "Georgia")
	fmt.Println("statePopulation after delete", statePopulation)

	fmt.Println(statePopulation, "before delete state population")
	deleteStatePopulation := statePopulation
	fmt.Println(deleteStatePopulation, "after delete state population")
}
