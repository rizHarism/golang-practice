package main

import "fmt"

func endApp() {
	fmt.Println("End Application")
	message := recover()
	fmt.Println("Terjadi Error", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi Bermasalah")
	}
}

func main() {
	runApp(false)
	fmt.Println("Defer, Panic, Recover")
}
