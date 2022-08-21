package functions

import "fmt"

func PointerCallingFun() {
	s := sumReturnPointer(1, 2, 3, 4, 5)
	fmt.Println(*s)
	fmt.Println()
}

func sumReturnPointer(values ...int) *int {
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}
