package pointers

import (
	"fmt"
)

func PointersInGo() {
	a := 67
	b := &a
	var f *int = &a
	fmt.Println(a, *b, *f)
	a = 23
	fmt.Println(a, *b, *f)
	*b = 47
	fmt.Println(a, *b)
	c := [3]int{1, 2, 3}
	d := &c[0]
	e := &c[1]
	*d = 78 + *e
	fmt.Println(c, *d, *e)
}
