package main

import (
	"fmt"
)


type myStruct struct {
	foo int
}



func main() {
	//In order to use pointer we have use longer syntax not shorter syntax
	//* key before the data type in actually referencing memory location of a
	//*b - when a value change b also get the same value of what a assigned
	//if we assigning new value to *b, a also get the b value
	// slice and map are pointer reference of array not only value


	//var a int = 22
	//var b *int = &a /* this is will referencing memory of a*/
	//fmt.Println(a, *b) /* *b is actually getting value from a memory, */
	//a = 100
	//fmt.Println(a, *b)
	//*b = 150
	//fmt.Println(a, *b)

	var ms *myStruct
	fmt.Println(ms)
	ms = new(myStruct)
	ms.foo = 20
	//(*ms).foo = 10
	fmt.Print(ms.foo)




}