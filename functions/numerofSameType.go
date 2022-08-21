package functions

import (
	"fmt"
)

func SumOfNumbers() {
	i := [5]int{1, 2, 3, 4, 5}
	sum(i)
	fmt.Println(i)
	sumDir(1, 2, 3, 4, 5)
}

func sum(values [5]int) {
	fmt.Println(values)
	result := 0
	for i, v := range values {
		result += v
		values[i] = 0
	}
	fmt.Println(result)
	fmt.Println(values)
}

func sumDir(values ...int) {
	fmt.Println(values)
	result := 0
	for i, v := range values {
		result += v
		values[i] = 0
	}
	fmt.Println(result)
}
