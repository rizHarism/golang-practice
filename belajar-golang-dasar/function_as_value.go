package main

import "fmt"

func goodBye(name string) string {
	return "selamat jalan " + name
}

func main() {
	variable1 := goodBye
	variable2 := goodBye

	fmt.Println(variable1("dobleh"))
	fmt.Println(variable2("jono"))
}
