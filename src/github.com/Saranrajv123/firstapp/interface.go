package main

import (
	"fmt"
)

type vehicle interface {
	speed()
	//wheels()
	//bodyType()

}

type bike struct {
	engine string
	brand  string
}

func (b bike) speed() {
	fmt.Println("engine: ", b.engine)
	fmt.Println("Brand: ", b.brand)
}

func vehicleFunc(v vehicle) {
	fmt.Println(v)
}

func main() {
	bk := bike{
		engine: "BS6",
		brand:  "Yamaha",
	}

	fmt.Println("bk", bk)
	vehicleFunc(bk)

}
