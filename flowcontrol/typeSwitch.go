package flowcontrol

import (
	"fmt"
)

func TypeSwitchInGo() {
	var i interface{} = 1.07
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is an float64")
	case string:
		fmt.Println("i is string")
	default:
		fmt.Println("i is another type")
	}
}
