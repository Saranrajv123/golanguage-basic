package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var m = sync.RWMutex{}
var counter = 0

func sayHello() {
	fmt.Println("Hello", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}

func main() {
	//name := "Hello"
	//go func(name string) {
	//	fmt.Println("name", name)
	//}(name)
	//fmt.Println("name 2", name)

	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}
