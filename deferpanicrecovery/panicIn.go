package deferpanicrecovery

import (
	"fmt"
)

func PanicInGo() {
	a, b := 1, 1

	fmt.Println(b)
	if b == 0 {
		fmt.Println("This is a Panic")
		panic("Some error happened!!")
	} else {
		fmt.Println(a / b)
	}
}
