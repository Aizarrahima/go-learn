package main

import "fmt"

func main() {
	// scope atas
	name := "Aizar"
	counter := 0

	// scope bawah
	increment := func() {
		name := "Adinda"
		fmt.Println("Increment")
		counter++ // closure => mengakses data yang berada diluar scope
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}

// note: scope yang berada dibawah tidak bisa diakses diluar scope tersebut. sedangkan scope di atas dapat diakses didalam function/scope bawah.