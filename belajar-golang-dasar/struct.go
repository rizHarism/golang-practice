package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// struct method (function in struct)
func (customer Customer) sayHello(name string, address string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
	fmt.Println("My address is", customer.Address)
	fmt.Println("is your address " + address + "?")
}

func main() {
	var rizqi Customer
	fmt.Println(rizqi)

	rizqi.Name = "Rizqi"
	rizqi.Address = "Blitar"
	rizqi.Age = 21
	fmt.Println(rizqi.Name)
	fmt.Println(rizqi.Address)
	fmt.Println(rizqi.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Kanigoro",
		Age:     32,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Jakarta", 43}
	fmt.Println(budi)

	rizqi.sayHello(budi.Name, budi.Address)
}
