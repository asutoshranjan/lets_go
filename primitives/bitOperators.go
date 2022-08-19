package primitives

import (
	"fmt"
)

func BitWise() {
	a := 10 // 1010
	b := 4  // 0100
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)  //1110
	fmt.Println(a &^ b) //1010
}

func BitShift() {
	a := 8 //1000
	fmt.Println(a >> 3)
	fmt.Println(a << 3)
}
