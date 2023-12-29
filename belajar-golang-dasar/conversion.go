package main

import "fmt"

func main() {
	var (
		nilai32 int32 = 32768
		nilai64 int64 = int64(nilai32)
		nilai16 int16 = int16(nilai32)
		name          = "Bonsky"
		e             = name[0]
		eString       = string(e)
	)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	fmt.Println(" ")

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)

	a := 5
	fmt.Println(a)

	a = 10
	fmt.Println(a)

}
