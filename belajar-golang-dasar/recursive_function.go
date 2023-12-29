package main

import "fmt"

func factorialLoop(number int) int {
	result := 1

	for i := number; i > 0; i-- {
		result *= i
	}

	return result
}

func factorialRecursive(number int) int {
	if number == 1 {
		return number
	} else {
		return number * factorialRecursive(number-1)
	}
}

func main() {
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursive(10))
}
