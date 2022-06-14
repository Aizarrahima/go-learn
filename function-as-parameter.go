package main

import "fmt"

// type declaration 
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spanFilter(name string) string {
	if name == "Lala" {
		return "...."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Aizar", spanFilter) // Hello Aizar
	sayHelloWithFilter("Lala", spanFilter) // Hello ....
}