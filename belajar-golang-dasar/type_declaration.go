package main

import "fmt"

func main() {
	type ktp string
	var ktp1 ktp = "4789127821384723489"
	// var ktp2 = ktp("7823487435874")

	fmt.Println(ktp1)
	fmt.Println(ktp("7897345457754"))
}
