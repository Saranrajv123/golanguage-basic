package main

import (
	"fmt"
)

const myConst int = 100

const (
	// errorSpecialist = iota
	// catSpecialist  + 6
	// dogSpecialist
	// ratSpecialist

	// _  = iota
	// KB = 1 << (10 * iota)
	// MB
	// GB
	// TB
	// PB
	// EB
	// ZB
	// YB

	isAdmin = 1 << iota
	canSeeHeadquaters
	canSeeFinancials

	canSeeEurope
	canSeeAsia
	// canSeeEurope
	canSeeAmerica
)

func main() {

	const myConst = 300
	var addedValue int64 = 400
	fmt.Println("hello world")
	fmt.Printf("%v, %T\n", myConst+addedValue, myConst+addedValue)
	// fmt.Printf("%v %T\n", myCon, myCon)

	fmt.Println("testing itoa in const")
	// var newSpecialist int
	// fmt.Printf("%v", catSpecialist)

	// fmt.Println("working with bitshift...")
	// fmt.Printf("%v, %T", KB, KB)
	// fmt.Println()

	// fileSize := 4000000000
	// fmt.Printf("%.2fGB", fileSize/GB)

	fmt.Println("next example ")
	fmt.Printf("%v, %T\n", isAdmin, isAdmin)
	fmt.Printf("%v, %T\n", canSeeHeadquaters, canSeeHeadquaters)

	fmt.Printf("%b\n", isAdmin)
	fmt.Printf("%b\n", canSeeHeadquaters)
	fmt.Printf("%b\n", canSeeFinancials)

	var roles byte = isAdmin | canSeeFinancials
	fmt.Printf("%b\n", roles)

	fmt.Printf("is Admin %v", isAdmin&roles)
}
