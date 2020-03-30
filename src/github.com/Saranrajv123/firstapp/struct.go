package main

import (
	"fmt"
)

type Doctor struct {
	number     int
	doctorName string
	companions []string
	expisode   []string
}

func main() {

	cpyDoctor := Doctor{
		// expisode:   []string{},
		number:     3,
		doctorName: "go",
		companions: []string{
			"go",
			"programming",
			"language",
		},
	}

	fmt.Println("cpyDoctor", cpyDoctor)

}
