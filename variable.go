package main

import "fmt"

func main(){
	// menggunakan kata kunci "var"
	var name string

	name = "Aizar Rahima"
	fmt.Println(name)

	name = "Alyshia"
	fmt.Println(name)


	// tidak menggunakan jenis type data
	var friendName = "Vania"
	fmt.Println(friendName)

	var age = 18
	fmt.Println(age)


	// tidak menggunakan kata kunci "var" dan menggantinya menggunakan ":="
	country := "Indonesia" // awal deklarasi
	fmt.Println(country)

	country = "Amerika" // variable sama, beda pada value
	fmt.Println(country)


	// Deklarasi Multiple Variable
	var (
		firstName = "Aizar"
		lastName = "Rahima"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}