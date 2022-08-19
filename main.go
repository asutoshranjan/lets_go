package main

import (
	"fmt"
	"go_lang/lib"
	"go_lang/primitives"
	"go_lang/space"
)

func main() {
	fmt.Println("This is Main Function being executed")
	space.SayHello()
	lib.Model1()
	space.Model2()
	primitives.BoolFunction()
	primitives.PlayWithNumbers()
	primitives.Calculation()
	fmt.Println()
	primitives.BitWise()
	fmt.Println()
	primitives.BitShift()
	fmt.Println()
	primitives.FloatingNums()
	fmt.Println()
	primitives.FloatingArithmatic()
	fmt.Println()
	primitives.ComplexNumbers()
	fmt.Println()
	primitives.TextType()
	fmt.Println()
	primitives.RuneGo()
}
