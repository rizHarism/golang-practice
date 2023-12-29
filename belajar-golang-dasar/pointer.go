package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	address1 := Address{"Blitar", "Jawa Timur", "Indonesia"}

	// default in golang variable is passing by value, not by reference
	address2 := address1 // Copy value | pass value
	address2.City = "Kediri"

	// variable address1 tidak berubah ketika ada perubahan pada address2
	fmt.Println(address1)
	fmt.Println(address2)

	address3 := Address{"Surabaya", "Jawa Timur", "Indonesia"}
	address4 := &address3 // dengan pointer (&) menjadikan address3 reference dengan address1
	address4.City = "Malang"
	// variable address3 berubah ketika ada perubahan pada address4
	fmt.Println(address3)
	fmt.Println(address4)

	// addres4 membuat Address baru dengan pointer, address3 tidak ikut berubah
	address4 = &Address{"Boyolali", "Jawa Tengah", "Indonesia"}
	fmt.Println(address3)
	fmt.Println(address4)

	// Asterik operator * pada awal variable
	// Merubah semua data srtructh yang bersinggungan dengan address4 | dalam kasus ini adalah address3
	address5 := Address{"Nganjuk", "Jawa Timur", "Indonesia"}
	address6 := &address5
	address6.City = "Pacitan"
	fmt.Println(address5)
	fmt.Println(address6)
	*address6 = Address{"Boyolali", "Jawa Tengah", "Indonesia"}
	fmt.Println(address5)
	fmt.Println(address6)

}
