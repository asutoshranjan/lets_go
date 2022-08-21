package looping

import (
	"fmt"
)

func ForEachLoopGo() {
	fmt.Println()
	s := []int{1, 2, 3, 4, 5}
	s = append(s, 6)
	//key Value
	for k, v := range s {
		fmt.Println(k, v)
	}
	fmt.Println()
	m := map[string]int{
		"Ashu":    8473,
		"Shoriya": 9083,
		"Raju":    9484,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println()
	a := "Hello Go!"
	for k, v := range a {
		fmt.Println(k, string(v))
	}
}
