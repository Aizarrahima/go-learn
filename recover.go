package main

import "fmt"

func endApp() {
	message := recover() // recover => menangkap data pada panic
	if message != nil {
		fmt.Println("Error dengan message", message)
	}
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(true)
	// untuk mengecek apakah aplikasi berhenti jika terdapat error
	fmt.Println("Aizar")
}
