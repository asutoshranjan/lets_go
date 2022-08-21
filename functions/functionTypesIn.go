package functions

import (
	"fmt"
)

func Function1() {
	fmt.Println()
	fmt.Println("Function1 Starts")
	s := sayMessage("Hello Riyali")
	fmt.Println(s)
}

func sayMessage(msg string) string {
	msg = "Hello Nana"
	fmt.Println(msg)
	return "I am a string"
}
