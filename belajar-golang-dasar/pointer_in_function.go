package main

import "fmt"

type Alamat struct {
	Kota, Provinsi, Negara string
}

func ChangeCountry(value *Alamat) {
	value.Negara = "Indonesia"
}

func main() {
	alamat := Alamat{"Malang", "Jawa Timur", "Singapore"}
	// membuat value yang dikirim ke func ChangeCountry menjadi pointer
	ChangeCountry(&alamat)
	fmt.Println(alamat)
}
