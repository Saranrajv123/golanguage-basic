package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")


}

func panicker() {
	fmt.Println("panicker start")
	if err := recover(); err != nil {
		fmt.Println("error", err)
		panic("inside panic error")
	}

	panic("something bad happened ")
	fmt.Println("panicker end")
}