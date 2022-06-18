package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Aizar Rahima", "Aizar")) // true
	fmt.Println(strings.Contains("Aizar Rahima", "Lisa")) // false

	fmt.Println(strings.Split("Aizar Rahima", " ")) // [Aizar Rahima]
	fmt.Println(strings.ToLower("Aizar Rahima")) // aizar rahima
	fmt.Println(strings.ToUpper("Aizar Rahima")) // AIZAR RAHIMA
	fmt.Println(strings.ToTitle("Aizar Rahima")) // AIZAR RAHIMA
	fmt.Println(strings.Trim("         Aizar Rahima         ", " ")) // Aizar Rahima (trim => menghapus data yang ada di kanan dan kirinya,selana menemukan data yang telah ditentukan maka akan terhapus )
	fmt.Println(strings.ReplaceAll("Aizar Aizar Aizar Aizar", "Aizar", "Budi")) // Budi Budi Budi Budi

}