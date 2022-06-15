package main

import "fmt"

type HasName interface {
	GetName() string // kontrak
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string { // kontrak
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var aizar Person
	aizar.Name = "Aizar"

	sayHello(aizar)

	cat := Animal{
		Name: "Push",
	}
	sayHello(cat)
}