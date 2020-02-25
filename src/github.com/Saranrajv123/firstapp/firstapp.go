package main

import (
	"fmt"
	"strconv"
)

var I int = 50

func main() {
	fmt.Println("global", I)
	// i = 40
	// fmt.Println(i)
	// var i int = 60
	// fmt.Println(i)

	var i int = 60
	// j := 70

	var j float32 = float32(i)
	j = 75.80

	var k int = int(j)
	// var l = string(j) through error

	var l = strconv.Itoa(i)

	fmt.Println(i)
	fmt.Printf("%T\n", j)
	fmt.Println(k)
	fmt.Println(l)
}
