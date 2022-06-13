package main

import "fmt"

func sayHello() { // menggunakan camel case
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
}

func main() {
	for i := 0; i < 10; i++ {
		sayHello()
	}
}
