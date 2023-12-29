package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("Nama " + name + " telah di blacklist")
	} else {
		fmt.Println("Selamat datang " + name)
	}
}

func main() {

	blackListName := func(name string) bool {
		return name == "Anjing"
	}

	registerUser("Anjing", blackListName)
	registerUser("Suryo", blackListName)

	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})

}
