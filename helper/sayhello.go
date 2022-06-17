package helper

import "fmt"

var version = 1
var Application = "Belajar Golang"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

// func yang tidak bisa diakses diluar package
func sayGoodBye(name string) {
	fmt.Println("Goodbye", name)
}