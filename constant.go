package main

import "fmt"

func main() {
	const firstName string = "Aizar"
	const lastName = "Rahima"
	const value = 1000

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	// Multiple Constant
	const (
		namaAwal string = "Aizar"
		namaAkhir		= "Rahima"
		number			= 1000
	)
	fmt.Println(namaAwal)
	fmt.Println(namaAkhir)
	fmt.Println(number)
}
