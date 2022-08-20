package collections

import (
	"fmt"
)

func GradesArray() {
	grades := []int{97, 85, 93}
	var list [5]int
	list[3] = 67
	fmt.Printf("Grades :%v \n", grades)
	fmt.Printf("List :%v \n", list)
	fmt.Printf("Number of grades :%v \n", len(grades))
}

func IdentityMatrix() {
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)
	fmt.Println()
}

func CopyOfArray() {
	a := [...]int{1, 2, 3, 4}
	b := a
	c := &a //and pointer points at address of a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	c[1] = 5
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println()
}
