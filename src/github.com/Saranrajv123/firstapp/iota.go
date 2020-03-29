package main

import (
	"fmt"
)

// const (
// 	_  = iota
// 	kb = 1 << (10 * iota)
// 	mb
// 	gb
// 	tb
// 	pb
// )

const (
	isAdmin = 1 << iota
	canSeeFinancials
	canSeeEurope
)

func main() {
	// fileSize := 4000000000
	// fmt.Println("itoa ", kb)
	// fmt.Printf("%v, %T\n", fileSize, fileSize)
	// fmt.Printf("%.2fGB", fileSize/gb)

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope

	fmt.Println("value", isAdmin)

	fmt.Printf("Is Admin %v\n", isAdmin)
	fmt.Printf("Is Admin %v\n", roles)
	fmt.Printf("Is Admin %v\n", isAdmin&roles)

}
