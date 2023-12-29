package main

import "fmt"

type Man struct {
	Name, Status string
}

// method Married with pointer
func (man *Man) Married() {
	if man.Status == "Married" {
		man.Name = "Mr. " + man.Name
	} else {
		man.Name = man.Name + " Jr."
	}
}

func main() {
	eko := Man{"Eko", "Married"}
	eko.Married()

	rizqi := Man{"Rizqi", "Single"}
	rizqi.Married()

	fmt.Println(eko.Name)
	fmt.Println(rizqi.Name)

	// memanggil seluruh data struct
	fmt.Println(eko)
	fmt.Println(rizqi)

}
