package functions

import (
	"fmt"
)

type Greeter struct {
	Name  string
	Greet string
}

func FuncAsVariable() {
	var f func() = func() {
		fmt.Println("Hello Go !")
	}
	f()
	g := func() {
		fmt.Println("This is B")
	}
	g()

	g1 := Greeter{
		Name:  "Asutosh",
		Greet: "Hello",
	}
	fmt.Println(g1.Greet, g1.Name)
}
