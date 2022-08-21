package looping

import (
	"fmt"
)

func SimpleLoopGo() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	for i, j := 0, 0; i <= 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	a := 0
	for ; a < 5; a++ {
		fmt.Println(a)
	}
	fmt.Println(a)
}
