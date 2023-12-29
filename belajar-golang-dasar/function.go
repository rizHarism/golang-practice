package main

import "fmt"

// function
func sayHello(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

// function return value
func greetings(name string) string {
	return "Hello " + name
}

// function multiple return value
func multipleGreetings() (string, string) {
	return "Halo", "Semua" // pengembalian 2 retun bertipe data string
}
func multipleTipeData(name string, text string, number int32) (string, string, int32) {
	return name, text, number
}

func main() {

	sayHello("rizqi", "ngenob")

	response := greetings("riz")
	println(response)

	firstText, secondText := multipleGreetings()
	fmt.Println(firstText, secondText)

	name, text, number := multipleTipeData("Joko", "juara", 1)
	fmt.Println(name, text, number)
	fmt.Println(multipleTipeData("Rano", "kelas", 10))
}
