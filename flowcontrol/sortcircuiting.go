package flowcontrol

import (
	"fmt"
)

func TestForShortcircuiting() {
	if false || returnTrue() {
		fmt.Println("Hey this is Inner Test")
	}
}

func returnTrue() bool {
	fmt.Println("Returning True")
	return true
}
