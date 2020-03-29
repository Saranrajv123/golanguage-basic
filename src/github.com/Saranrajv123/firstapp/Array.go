package main

import (
	"fmt"
)

func main() {
	studentsArr := [...]int{34, 45, 56, 576, 7868, 67867, 976, 56, 45, 53, 23}

	// // won't affetc the original Array
	// copyOfStudentsArr := studentsArr
	// copyOfStudentsArr[1] = 54

	// // affect original Array if use "&" symbol
	// secondCopyOfStd := &studentsArr
	// secondCopyOfStd[1] = 0

	// var employeeArr [10]string
	// employeeArr[2] = "saranraj"
	// employeeArr[4] = "saran"
	// employeeArr[6] = "raj"
	// employeeArr[8] = "go "
	// employeeArr[0] = "js"

	// var sampleArr [10]int
	// fmt.Println(sampleArr, "sampleArr")

	// // need to learn more
	// // do some research array  matrix
	// var matrixArr [3][3]int

	// matrixArr[0] = [3]int{1, 0, 0}

	// // slice Array
	// sliceArr := []string{"go", "programming"}
	// copySliceArr := sliceArr
	// copySliceArr[0] = "golang"

	// // make key word
	// usingmakeinbuilffun := make([]int, 5, 1000)
	// usingmakeinbuilffun[0] = 10

	// fmt.Println("usingmakeinbuilffun", usingmakeinbuilffun)
	// fmt.Println("length of the usingmakeinbuilffun", len(usingmakeinbuilffun))
	// fmt.Println("capacity of the usingmakeinbuilffun", cap(usingmakeinbuilffun))

	// fmt.Println("Working array")
	// fmt.Printf("%v %T\n", studentsArr, studentsArr)
	// fmt.Printf("%v %T\n", employeeArr, employeeArr)
	// fmt.Printf("find the length of the emaployee array %v\n", len(employeeArr))
	// fmt.Println(copyOfStudentsArr, "copyOfStudentsArr")
	// fmt.Println(secondCopyOfStd, "copyOfStudentsArr")

	// fmt.Println("matrix")
	// fmt.Println(matrixArr)

	// fmt.Printf("%v %T\n", sliceArr, sliceArr)
	// fmt.Printf("%v %T\n", copySliceArr, copySliceArr)

	// appending this work in array as well as slice
	newArr := []int{}
	newArr = append(newArr, 1, 3, 4, 5)

	// studentsArr = studentsArr[1:5]
	copyStdArr := studentsArr[:len(studentsArr)-1]

	fmt.Println("newArr", newArr)
	fmt.Println("capacity ", cap(newArr))

	fmt.Println("stud arr", studentsArr)
	fmt.Println("copy stud arr", copyStdArr)
	fmt.Println("stud arr", studentsArr)

}
