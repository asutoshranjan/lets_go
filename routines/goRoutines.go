package routines

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func GoRoutines() {
	var msg string = "Hello"
	wg.Add(2)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	msg = "GoodBye"
	go sayHello()
	wg.Done()
	wg.Wait()
}

func sayHello() {
	fmt.Println("Hello Asu!! Go Routine")
}
