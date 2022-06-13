package main

import "fmt"

func main() {

	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke -", counter)
		counter++
	}

	
	
	// for dengan statement => code lebih ringkas
	for counter := 1;  counter <= 10; counter++ { // counter := 1 => init statement; counter <= adalah kondisi; counter++ => post statement
		fmt.Println("Perulangan ke - ", counter)
	}

	slice := []string{"Aizar", "Rahima", "Suprayitno", "Adinda", "Alyshia"} // iterasi data
	for  i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	
	
	// For Rage
	// names := []string{"Rara", "Vania", "Virda"}
	// for index, name := range names{ // index => key; name => value; names => bisa dimasukkan slice, array, maupun maps (nama variable).
	// 	fmt.Println("index", index, "=", name)
	// }

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "Eko"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

}