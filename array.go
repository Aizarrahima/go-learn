package main 

import "fmt"

func main() {
	
	// Array secara manual
	var names [3]string // 3 adalah jumlah batas maksimum data array

	// apabila melebihi batas maksimum data array maka akan terjadi error
	names[0] = "Aizar"
	names[1] = "Rahima"
	names[2] = "Suprayitno"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])


	// Array secara Langsung
	fmt.Println("Array secara langsung")
	var values = [3]int {
		90,
		80,
		95,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])


	// Function Array
	// len = untuk memanggil panjang array bukan jumlah datanya
	fmt.Println(len(names))
	fmt.Println(len(values))

	var coba [10]string // 10 adalah panjang array
	fmt.Println(len(coba))
}