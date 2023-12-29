package main

import "fmt"

// nil hanya bisa sebagai digunakan pada interface{}, slice, map, function()m pointer dan channel

func NewSlice(name string) []string {
	if name == "" {
		return nil
	} else {
		return []string{name}
	}
}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{"name": name}
	}
}

func main() {
	dataSlice := NewSlice("eko")
	dataMap := NewMap("andre")

	if dataSlice == nil {
		fmt.Println("data slice kosong")
	} else {
		fmt.Println(dataSlice[0])
	}

	if dataMap == nil {
		fmt.Println("data map kosong")
	} else {
		fmt.Println(dataMap["name"])
	}
}
