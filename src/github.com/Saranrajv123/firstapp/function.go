package main

import (
	"fmt"
)

func main() {
	greeting := "hello"
	name := "saranraj"

	sayGreeting(&greeting, &name)
	fmt.Println(name, "inside main function")

}

func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "go lang"
	fmt.Println(*name, "name")
}
