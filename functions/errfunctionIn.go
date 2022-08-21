package functions

import (
	"fmt"
)

func ErrorHandelInGo() {
	a, b := 2.0, 0.0
	res, err := divide(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
