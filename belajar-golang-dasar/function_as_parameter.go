package main

import "fmt"

func sayHelloFilter(value string, filter func(string) string) {
	filteredName := filter(value)
	fmt.Println("Hello " + filteredName)
}

func spamCheck(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {

	sayHelloFilter("Anjing", spamCheck)

	withVariable := sayHelloFilter

	name2 := "ibnu"
	withVariable("soni", spamCheck)
	withVariable(name2, spamCheck)

}
