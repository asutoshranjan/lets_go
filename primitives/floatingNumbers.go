package primitives

import (
	"fmt"
)

func FloatingNums() {
	var a float32 = 78.69850
	n := 3.14
	fmt.Printf("%v, %T \n", n, n)
	n = 2.1e14 + 1.0
	fmt.Printf("%v, %T \n", a, a)
	fmt.Printf("%v, %T \n", n, n)
	fmt.Println(float64(a) + n)
}

func FloatingArithmatic() {
	var (
		a float32
		b float32
	)
	a = 10.2 //1010
	b = 3.7  //0011
	fmt.Println(a + b)
	fmt.Println(a - 3)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(int(a) &^ int(b))
}
