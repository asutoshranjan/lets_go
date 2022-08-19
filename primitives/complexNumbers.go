package primitives

// complexnumber in go is terated as a first class citizen

import (
	"fmt"
)

func ComplexNumbers() {
	var n complex64 = 2i
	var m complex64 = 2 + 6i
	var o complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", o, o)
	fmt.Printf("%v, %T\n", imag(m), imag(m))
	fmt.Println(m + n)
	fmt.Println(m - n)
	fmt.Println(m * n)
	fmt.Println(m / n)
}
