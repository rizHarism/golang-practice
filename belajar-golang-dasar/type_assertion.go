package main

import (
	"fmt"
)

func Random() any {
	return 123
}

func main() {
	result := Random()

	switch val := result.(type) {
	case string:
		fmt.Println("String", val)
	case int:
		fmt.Println("Integer", val)
	default:
		fmt.Println("Unknown", val)
	}

}
