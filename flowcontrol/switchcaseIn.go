package flowcontrol

import (
	"fmt"
)

func SwitchCaseInGo() {
	switch i := 2 + 3; i { //Syntax for initializing and checking for values
	case 1, 5, 10:
		fmt.Println("One or Five or Ten")
	case 2, 4, 6:
		fmt.Println("Two or Four or Six")
	default:
		fmt.Println("Some Other Num")
	}

	j := 15

	switch {
	case j <= 10:
		fmt.Println("less than or equal to Ten")
	case j <= 20:
		fmt.Println("less than or equal to Twenty")
	default:
		fmt.Println("greater than twenty")
	}
}
