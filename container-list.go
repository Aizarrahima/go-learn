package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	// PushBack => untuk menambahkan ke data yang paling belakang
	data.PushBack("Aizar")
	data.PushBack("Rahima")
	data.PushBack("Suprayitno")

	// PushFront => untuuk menambahkan ke data yang paling awal/depan
	data.PushFront("Adinda")

	fmt.Println(data.Front().Value) // mengambil data yang paling depan
	fmt.Println(data.Back().Value)  // mengambil data yang paling belakang

	fmt.Println(data.Front().Prev()) // Prev => sebelumnya
	fmt.Println(data.Back().Next())  // Next => setelahnya
	

	// untuk meng-iterate semua data yang ada
	// data dari depan ke belakang
	for element := data.Front(); element != nil; element.Next() {
		fmt.Println(element.Value)
	}

	// meng-iterate data dari belakang ke depan
	for element := data.Back(); element != nil; element.Prev() {
		fmt.Println(element.Value)
	}
}

