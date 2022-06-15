package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// func sayHi(customer Customer, name string){
// 	fmt.Println("Hello", name, "My Name is", customer.Name)
// }

func (customer Customer) sayHi(name string){
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayHello(){
	fmt.Println("Hello from", a.Name)
}

func main() {
	var aizar Customer
	aizar.Name = "Aizar Rahima"
	aizar.Address = "Indonesia"
	aizar.Age = 18

	// sayHi(aizar, "Lisa")

	aizar.sayHi("Lisa")
	aizar.sayHello()

	// fmt.Println(aizar.Name)
	// fmt.Println(aizar.Address)
	// fmt.Println(aizar.Age)

	// adinda := Customer {
	// 	Name: "Adinda",
	// 	Address: "Malang",
	// 	Age: 18,
	// } 
	// fmt.Println(adinda)
}