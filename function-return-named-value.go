package main

import "fmt"

func getFullNamae() (firstName, middleName, lastName string) { // apabila tipe datanya sama, dapat dituliskan di akhir saja.
	firstName = "Aizar"
	middleName = "Rahima"
	lastName = "Suprayitno"

	return
}

func main() {
	firstName, middleName, lastName := getFullNamae() // untuk nama variabel disini tidak harus sama dengan nama variabel pada parameter di function atas.
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
