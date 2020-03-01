package main

import (
	"fmt"
)

var I int = 50

const myConst int16 = 10

func main() {
	// fmt.Println("global", I)
	// // i = 40
	// // fmt.Println(i)
	// // var i int = 60
	// // fmt.Println(i)

	// var i int = 60
	// // j := 70

	// var j float32 = float32(i)
	// j = 75.80

	// var k int = int(j)
	// // var l = string(j) through error

	// var l = strconv.Itoa(i)

	// fmt.Println(i)
	// fmt.Printf("%T\n", j)
	// fmt.Println(k)
	// fmt.Println(l)

	//Constant

	// var a = strconv.Itoa(10)
	const a = 10

	const (
		i = iota
		j
		k
		l
	)

	const (
		x = iota
		y
		z
	)
	var addinValue int = 20
	const b = "string value"

	fmt.Printf("%v, %T\n", a+addinValue, a+addinValue)
	fmt.Printf("%v, %T\n", b, b)

	fmt.Printf("%v, %T\n", myConst, myConst)

	fmt.Printf("%v\n", i)
	fmt.Println("value", j)
	fmt.Println("value", k)
	fmt.Println("value", l)

	fmt.Println("value", x)
	fmt.Println("value", y)
	fmt.Println("value", z)

}
