package functions

import (
	"fmt"
)

// pass by reference
func FunctionsWithPointers() {
	fmt.Println()
	greeting := "Hello"
	name := "Sanya"
	sayGreeting(&greeting, &name)
	fmt.Println(name)
}

func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}
