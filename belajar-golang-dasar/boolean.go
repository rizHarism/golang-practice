package main

import "fmt"

func main() {
	uas := 90
	absensi := 85

	var lulusUas bool = uas > 80
	lulusAbsensi := absensi > 80

	lulus := lulusUas && lulusAbsensi

	fmt.Println(lulus)
}
