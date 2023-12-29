package main

import "fmt"

func main() {
	hari := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	slice1 := hari[5:]
	fmt.Println(hari)
	fmt.Println(slice1)
	// slice1[0] = "Saturday"
	// slice1[1] = "Sunday"
	fmt.Println(hari)
	fmt.Println(slice1)

	slice2 := append(hari[:], "Monday")
	fmt.Println(hari)
	fmt.Println(slice1)
	fmt.Println(slice2)

}
