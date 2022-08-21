package looping

import (
	"fmt"
)

func ForWhileLoopGo() {
	fmt.Println()
	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}
	fmt.Println()
	//continue statement
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
