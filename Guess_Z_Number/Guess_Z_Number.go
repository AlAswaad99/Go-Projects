package main

import (
	"fmt"
	"math/rand"
)

func main() {
	tries := 5
	answer := rand.Intn(20) + 1
	var numList []int

	fmt.Print("I am thinking of a random number between 1 and 20. Do think u can find it?\nHere. Try: ")

	for tries > 0 {

		var input int
		fmt.Scanln(&input)
		flag := AlreadyGuessed(numList, input)

		if flag {
			fmt.Println("You have already guesed the number")
		} else {
			if input < answer {
				fmt.Println("Try guessing a larger number. You have ", tries-1, " tries remaining")
			} else if input > answer {
				fmt.Println("Try guessing a smaller number. You have ", tries-1, " tries remaining")
			} else if input == answer {
				fmt.Println("You are correct!! You finished with ", tries-1, " tries remaining")
				break
			}
			tries--
		}
		numList = append(numList, input)
	}
	fmt.Println("Game Over Loser!!!!The answer was", answer)
}

func AlreadyGuessed(numGuessed []int, input int) bool {
	for _, item := range numGuessed {
		if item == input {
			return true
		}
	}
	return false
}
