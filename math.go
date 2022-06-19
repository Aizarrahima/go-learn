package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(1.7)) // 2 => karena dibulatkan paling depan ke atas
	fmt.Println(math.Round(1.3)) // 1 => karena dibulatkan paling depan ke bawah
	fmt.Println(math.Floor(1.7)) // 1 => meskipun dibulatkan ke atas, dengan menggunakan floor akan dipaksa dibulatkan ke bawah
	fmt.Println(math.Ceil(1.3)) // 2 => meskipun dibulatkan ke bawah, dengan menggunakan ceil akan dipaksa dibulatkan ke atas

	fmt.Println(math.Max(10, 20)) // 20; Max => untuk mengecek yang terbesar
	fmt.Println(math.Min(10, 20)) // 10; Min => untuk mengecek yang terkecil
}