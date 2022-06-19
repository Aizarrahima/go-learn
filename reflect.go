package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type ContohLagi struct {
	Name string `required:"true"`
	Description string `required:"true"`
}

// validation library
func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data) // logic untuk mengambil semua data reflection
	// lalu melakukan iterasi semua field-nya
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		// kemudian melakukan pengecekan apakah field "required" memiliki value "true"
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface() // melakukan validasi
			if value == "" {
				return false
			}
		}
	}
	return true
}


func main() {
	sample := Sample{"Aizar"}
	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0) // Name

	fmt.Println(sampleType.NumField()) // 1
	fmt.Println(structField.Name) // Name
	fmt.Println(sampleType.Field(0).Tag.Get("required")) // ture
	fmt.Println(sampleType.Field(0).Tag.Get("max")) // 10
	fmt.Println(sampleType.Field(0).Tag.Get("min")) // string kosong => karena field-nya kosong

	// melakukan validasi
	sample.Name = ""
	fmt.Println(IsValid(sample))

	// contoh lain
	contoh := ContohLagi{"Aizar", "Oke"}
	fmt.Println(IsValid(contoh)) // true => meskipun tidak ada value apapun tetap menghasilkan true ketikan di compile, karena tidak terdapat "required"
}