package main

import "fmt"

func main() {

	name := "Aizar"

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
	
}