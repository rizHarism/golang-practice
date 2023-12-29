package main

import (
	"fmt"
)

func sumSlice(text string, numbers ...int) (string, int) {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return text, total
}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
func main() {

	slice1 := []int{20, 20, 20, 20, 20}

	hasil := sumAll(10, 10, 10, 10, 10)
	fmt.Println(hasil)
	fmt.Println(sumSlice("Total adalah :", slice1...))
}
