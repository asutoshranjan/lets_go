package primitives

import (
	"fmt"
)

func PlayWithNumbers() {
	var n int = 99594
	m := 3989892729709088
	var num uint16 = 89
	var i uint8 = 2
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)
	fmt.Printf("%v, %T\n", num, num)
	fmt.Printf("%v, %T\n", i, i)
}

func Calculation() {
	a := 45
	b := 10
	var (
		c int  = 67
		d int8 = 8
	)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	fmt.Println(c + int(d))
}
