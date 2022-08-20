package flowcontrol

import (
	"fmt"
)

func SwitchCaseInGo() {
	switch 2 {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Not one or two")
	}
}
