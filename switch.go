package main

import "fmt"

func main() {

	name := "Aizara"

	switch name {
	case "Aizar":
		fmt.Println("Hello Aizar")
		fmt.Println("Hello Aizar")
	case "Adinda":
		fmt.Println("Hello Adinda")
		fmt.Println("Hello Adinda")
	default:
		fmt.Println("Yuk Kenalan!")
		fmt.Println("Yuk Kenalan!")
	}

	// switch short statements
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// switch tanpa kondisi
	length := len(name)
	switch {
	case length > 18:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}