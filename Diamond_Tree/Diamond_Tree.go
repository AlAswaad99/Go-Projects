package main

import "fmt"

func main() {
	drawDiamondUp()
	drawDiamondDown()
}

func drawDiamondUp() {
	i, j, h := 1, 1, 10
	k := h - 1

	for i = 1; i <= h; i++ {
		for k = h; k >= i; k-- {
			fmt.Print(" ")
		}
		for j = 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func drawDiamondDown() {
	i, j, h := 1, 1, 10
	k := h - 1

	for i = h - 1; i >= 1; i-- {
		for k = h - 1; k >= i; k-- {
			fmt.Print(" ")

		}
		for j = i; j >= 1; j-- {
			fmt.Print(" *")
		}
		fmt.Println()
	}
}
