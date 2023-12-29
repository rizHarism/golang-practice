package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan ke", i, "guys")
	}
	fmt.Println("Done !")

	//  Looping slice manual
	names := []string{"Rizqi", "Harisma", "Uchrowi"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i], " ")
	}

	// looping slice with for range
	for _, name := range names {
		fmt.Println(name)
	}

	// for loop break
	for i := 0; i < 10; i++ {
		if i == 7 {
			break
		}
		fmt.Println("Perulangan ke", i)
	}

	// for loop continue
	fmt.Println("for loop continue")
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue // baris di bawah continue tidak dilanjutkan, dia langsung melakukan looping selanjutnya
		}
		fmt.Println("Perulangan ke", i) // akan dieksekusi jika perkondisian diatas bernilai false
	}
}
