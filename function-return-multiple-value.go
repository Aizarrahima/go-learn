package main

import "fmt"

func getFullName() (string, string, string) {
	return "Aizar", "Rahima", "Suprayitno"
}

func main() {
	firstName, _, lastName := getFullName() // _ => untuk menghiraukan multiple value
	fmt.Println(firstName)
	// fmt.Println(middleName)
	fmt.Println(lastName)
}
