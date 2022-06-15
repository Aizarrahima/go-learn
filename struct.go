package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var aizar Customer
	aizar.Name = "Aizar Rahima"
	aizar.Address = "Indonesia"
	aizar.Age = 18

	fmt.Println(aizar.Name)
	fmt.Println(aizar.Address)
	fmt.Println(aizar.Age)

	adinda := Customer {
		Name: "Adinda",
		Address: "Malang",
		Age: 18,
	} 
	fmt.Println(adinda)
}