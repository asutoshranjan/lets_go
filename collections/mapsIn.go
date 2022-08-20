package collections

import (
	"fmt"
)

func MapsInGo() {
	// Map = key Value pair
	statePopulations := map[string]int{
		"California": 89590489,
		"Texas":      8949005,
		"Florida":    93893003,
		"New York":   9048404085,
		"Ohio":       94994087,
	}
	statePopulations["Georgia"] = 29200005
	fmt.Println(statePopulations)
	fmt.Println(len(statePopulations))
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Texas"])
	odisha, ok := statePopulations["Odisha"] // , ok syntax of map
	fmt.Println(odisha, ok)
	sp := statePopulations
	sp["Odisha"] = 95899899994
	fmt.Println(sp)
	fmt.Println(statePopulations)
	fmt.Println(len(statePopulations))
	delete(sp, "Odisha")
	fmt.Println(statePopulations)
	fmt.Println(len(statePopulations))
	fmt.Println()
}
