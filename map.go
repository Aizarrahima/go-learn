package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Aizar",
		"address": "Blitar",
	}

	person["title"] = "Programmer"

	fmt.Println(person)

	// contoh map[key]
	fmt.Println(person["name"]) 
	fmt.Println(person["address"])
	fmt.Println(person["title"])


	// function
	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Eko"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}