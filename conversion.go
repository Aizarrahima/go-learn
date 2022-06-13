package main

import "fmt"

func main() {
	// konversi angka
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32) // 100000
	fmt.Println(nilai64) // 100000
	fmt.Println(nilai8) // -96 (karena terjadi integer overflow, saat mencapai batas maka balik ke nilai yang paling bawah lalu terus menambah hingga mencapai batas maksimum dan balik lagi ke nilai yang paling bawah)


	// konversi string
	var name = "Aizar"
	var e byte = name[0] // memanggil karakter dan mengkonversi ke byte
	var eString string = string(e) // konversi byte ke string

	fmt.Println(name)
	fmt.Println(eString)
}
