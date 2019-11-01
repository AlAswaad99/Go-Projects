package main

import "fmt"

func main() {
	fmt.Print("Enter a number to evaluate its Fibonacci Sequence \t")
	input := 0
	fmt.Scan(&input)

	Fibonacci(input)

}

func Fibonacci(num int) {

	counter, num1, num2 := 0, 0, 1

	fmt.Println(num1)

	for num2 <= num {
		fmt.Println(num2)

		num2 = num1 + num2
		num1 = num2 - num1

		counter++
	}
}
