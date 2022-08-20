package consts

import (
	"fmt"
	//"math"
)

const myConst int16 = 68

// const block

const (
	a1 = iota + 1
	b1
	c1
)

const (
	//write only constant
	_  = iota
	a2 = iota
)

func GoConstants() {
	fmt.Println("Constants ...")
	fmt.Printf("%v, %T \n", myConst, myConst)
	// const myConst float64 = math.Sin(1.57) ..error
	const myConst float32 = 3.14
	fmt.Printf("%v, %T \n", myConst, myConst)
	// implicit conversion
	const a = 32
	var b uint8 = 87
	fmt.Printf("%v, %T \n", a+b, a+b)
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
	fmt.Println(a2)
	fmt.Println()
}
