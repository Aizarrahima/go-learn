package main

import "fmt"

func random() interface{} {
	return true
}

func main() {
	var result interface{} = random()
	// var resultString string = result.(string) // conversion interface to string; apabila salah melakukan conversion maka akan terjadi panic
	// fmt.Println(resultString)

	// Type Assertions Switch
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}
}