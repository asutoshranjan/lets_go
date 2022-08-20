package main

import (
	"fmt"
	"go_lang/collections"
	"go_lang/consts"
	"go_lang/consts/projects"
	"go_lang/flowcontrol"
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
	fmt.Println()
	consts.GoConstants()
	fmt.Println()
	projects.FileSizeConverter()
	fmt.Println()
	projects.AdminAccess()
	fmt.Println()
	collections.GradesArray()
	fmt.Println()
	collections.IdentityMatrix()
	collections.CopyOfArray()
	collections.SliceInGo()
	collections.StackinSlice()
	collections.StackPushPop()
	collections.MapsInGo()
	collections.DoctorStruct()
	collections.StructInheritance()
	flowcontrol.IfElseInGo()
	flowcontrol.NumberGuessWithIF()
	flowcontrol.TestForShortcircuiting()
	flowcontrol.SwitchCaseInGo()
	flowcontrol.TypeSwitchInGo()
}
