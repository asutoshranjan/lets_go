package flowcontrol

import (
	"fmt"
)

var Flag bool = true

func IfElseInGo() {
	if Flag {
		fmt.Println("This test is True")
		fmt.Println()
	}
}

func NumberGuessWithIF() {
	num := 50
	guess := -7
	if guess < 1 || guess > 100 {
		fmt.Println("Invalid Number Should be in 1 to 100 range")
	} else {
		if guess < num {
			fmt.Println("Too Low")
		}
		if guess > num {
			fmt.Println("Too High")
		}
		if guess == num {
			fmt.Println("You Got it!!")
		}
	}

	if nums := 20; nums < 40 {
		fmt.Printf("yeah the nums: %v is safe! \n", nums)
	}
	fmt.Println()
}
