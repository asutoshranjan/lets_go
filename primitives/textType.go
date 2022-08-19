package primitives

import (
	"fmt"
)

func TextType() {
	s := "this is a String"
	s2 := "Hi I am A file"
	b := []byte(s2)
	fmt.Printf("%v, %T \n", s, s)
	fmt.Printf("%v, %T \n", s[2], s[2])
	fmt.Println(string(s[2])) //also acts like an arry
	fmt.Println(s + " this is great news as java")
	fmt.Printf("%v, %T \n", b, b)
	// converting Strings later Files to byte slices
}

func RuneGo() {
	n := 258
	var char string = string(rune(n))
	r := 'A'
	fmt.Printf("%v, %T \n", r, r)
	fmt.Printf("%v, %T \n", char, char)
}
