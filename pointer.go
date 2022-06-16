package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// Pointer di Function
func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	// operator &
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1 

	address2.City = "Bandung" // City => field; address2 => pointer

	// operator *
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)


	// function new
	var address4 *Address = new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)


	// pointer di function
	// data asli
	var alamat = Address { 
		City: "Subang",
		Province: "Jawa Barat",
		Country: "",
	}
	var alamatPointer *Address = &alamat
	ChangeCountryToIndonesia(alamatPointer) 
	fmt.Println(alamat)
}