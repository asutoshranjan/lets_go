package looping

import (
	"fmt"
)

func LabledLoop() {
	fmt.Println()
	//This Loop Block is Labled as Loop
Loop:
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}
}
