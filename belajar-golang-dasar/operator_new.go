package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// cara lain penulisan variable | dalam kasus ini variable pointer
	var address1 *Address = new(Address) //new() sama dengan &Address{} atau type Address pointer kosong
	var address2 *Address = address1

	// merubah isi dari struct Address dengan 2 variable pointer
	address2.City = "Blitar"
	address2.Province = "Jawa Timur"
	address1.Country = "Indonesia"

	// Hasilnya akan sama, karena address2 itu reference ke address1 dan sama2 pointer
	fmt.Println(address1)
	fmt.Println(address2)
}
