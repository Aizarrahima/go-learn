package main

import (
	"fmt"
	"go-learn/helper"
)

func main() {
	helper.SayHello("Aizar")

	// tidak bisa diakses karena awalan nama function menggunakan huruf kecil
	// helper.sayGoodBye("Aizar") // error

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error, karena menggunakan awalan nama variable menggunakan huruf kecil
}