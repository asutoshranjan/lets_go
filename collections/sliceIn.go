package collections

import (
	"fmt"
)

func SliceInGo() {
	// Slices are reference type by default
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := a[:]
	e := a[3:]
	f := a[:6]
	g := a[3:6]
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Printf("Length: %v \n", len(g))
	fmt.Printf("Capacity: %v \n", cap(g))
	// Initialize Slice with make func
	d := make([]int, 3, 10) // type, length, capacity
	fmt.Println(d)
	d = append(d, 1)
	h := []int{}
	fmt.Println(h)
	h = append(h, 23)
	fmt.Println(h)
	fmt.Printf("Length: %v \n", len(h))
	fmt.Printf("Capacity: %v \n", cap(h))
	fmt.Println(d)
	fmt.Printf("Length: %v \n", len(d))
	fmt.Printf("Capacity: %v \n", cap(d))
	fmt.Println()
}

func StackinSlice() {
	a := []int{}
	a = append(a, 1)
	fmt.Printf("%v, length: %v capacity: %v \n", a, len(a), cap(a))
	a = append(a, 2, 3, 4, 5)
	fmt.Printf("%v, length: %v capacity: %v \n", a, len(a), cap(a))
	b := []int{6, 7, 8, 9}
	a = append(a, b...)
	fmt.Printf("%v, length: %v capacity: %v \n", a, len(a), cap(a))
	fmt.Println()
}

func StackPushPop() {
	a1 := [...]int{1, 2, 3, 4, 5}
	a := a1
	b := a[:len(a)-1] //pop element
	fmt.Println(b)
	b = append(a[:2], a[3:]...) // remove only middle
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(a1)
	fmt.Println()
}
