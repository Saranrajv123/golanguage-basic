package main

import (
	"fmt"
)

func main() {
	studentsArr := [...]int{34, 45, 56, 576, 7868, 67867, 976, 56, 45, 53, 23}

	// won't affetc the original Array
	copyOfStudentsArr := studentsArr
	copyOfStudentsArr[1] = 54

	// affect original Array if use "&" symbol
	secondCopyOfStd := &studentsArr
	secondCopyOfStd[1] = 0

	var employeeArr [10]string
	employeeArr[2] = "saranraj"
	employeeArr[4] = "saran"
	employeeArr[6] = "raj"
	employeeArr[8] = "go "
	employeeArr[0] = "js"

	var sampleArr [10]int
	fmt.Println(sampleArr, "sampleArr")

	// need to learn more
	// do some research array  matrix
	var matrixArr [3][3]int

	matrixArr[0] = [3]int{1, 0, 0}

	// slice Array
	sliceArr := []string{"go", "programming"}
	copySliceArr := sliceArr
	copySliceArr[0] = "golang"

	fmt.Println("Working array")
	fmt.Printf("%v %T\n", studentsArr, studentsArr)
	fmt.Printf("%v %T\n", employeeArr, employeeArr)
	fmt.Printf("find the length of the emaployee array %v\n", len(employeeArr))
	fmt.Println(copyOfStudentsArr, "copyOfStudentsArr")
	fmt.Println(secondCopyOfStd, "copyOfStudentsArr")

	fmt.Println("matrix")
	fmt.Println(matrixArr)

	fmt.Printf("%v %T\n", sliceArr, sliceArr)
	fmt.Printf("%v %T\n", copySliceArr, copySliceArr)

}
