package main

import "fmt"

func main() {

	i, j := 1, 1
	for j <= 12 {
		for i <= 12 {
			fmt.Print(i*j, "\t")
			i++
		}
		fmt.Print("\n\n")
		j++
		i = 1
	}

	fmt.Scanln("\n\nPress Enter key to exit")

}
