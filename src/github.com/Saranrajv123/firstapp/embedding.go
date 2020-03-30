package main

import (
	"fmt"
)

type Animal struct {
	name   string
	origin string
}

type bird struct {
	Animal
	fly bool
	eat string
}

type cars struct {
	name string `required`
	engine string
}

func main() {

	b := bird{
		Animal: Animal{name: "peacock", origin: "India"},
		fly:    true,
		eat:    "rice",
	}

	car := cars {
		name: "",
		engine: ""
	}

	fmt.Println(b)
	fmt.Println('car', car)

}
