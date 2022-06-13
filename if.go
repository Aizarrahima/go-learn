package main

import "fmt"

func main() {

	var name = "Nana"

	if name == "Aizar" { // cek kondisi, apabila kondisi benar maka akan jalan perintah selanjutnya. Jika kondisi salah maka perintah berikutnya tidak jalan
		fmt.Println("Hello Aizar")
	} else if name == "Adinda" {
		fmt.Println("Hello Adinda")
	} else if name == "Nana" {
		fmt.Println("Hello Nana")
	} else {
		fmt.Println("Yuk kenalan!")
	}


	// if short statements
	// var length = len(name)
	if length := len(name) ; length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}