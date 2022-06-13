package main

import (
	"fmt"
)

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)

	// function slice
	fmt.Println("Function Slice")
	fmt.Println(len(slice1)) // len => mendapatkan panjang
	fmt.Println(cap(slice1)) // cap => mendapatkan capacity

	// Apabila kita mengubah slice, maka data array juga ikut terubah
	// months[5] = "Diubah"
	// fmt.Println(slice1)

	// slice1[0] = "Mei Lagi"
	// fmt.Println(months)

	// Append Slice => menambah data baru pada posisi akhir ketika kapasitas masih muat, apabila kapasitas sudah tidak muat akan membuat array baru
	var slice2 = months[11:]
	fmt.Println("slice2 = ", slice2)

	var slice3 = append(slice2, "Aizar")
	fmt.Println("slice3 = ", slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println("slice3 = ", slice3)

	fmt.Println("slice2 = ", slice2)
	fmt.Println("months = ", months)

	// Make Slice => membuat slice baru dari awal
	newSlice := make([]string, 2, 5) // 2 => length; 5 => capacity
	newSlice[0] = "Aizar"
	newSlice[1] = "Adinda"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))   // len => panjang data
	fmt.Println(cap((newSlice))) // kapasitas

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
