package primitives

import (
	"fmt"
)

func BoolFunction() {
	var n bool
	flag := 1 != 2
	fmt.Printf("%v, %T \n", n, n)
	if flag {
		fmt.Println("Hello I am Bool")
	}
}
