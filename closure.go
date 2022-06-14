package main

import "fmt"

func main() {

	name := "Aizar"
	counter := 0

	increment := func() {
		name = "Adinda"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}