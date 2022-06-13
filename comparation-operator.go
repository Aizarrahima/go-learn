package main

import "fmt"

func main() {

	var fruit1 = "Apple"
	var fruit2 = "Apple"
	var fruit3 = "Mango"

	var result1 bool = fruit1 == fruit2
	fmt.Println(result1)

	var result2 bool = fruit1 == fruit3
	fmt.Println(result2)
	

	fmt.Println("Result Comparation Integer")
	var value1 = 100
	var value2 = 200
	fmt.Println(value1 > value2)	
	fmt.Println(value1 < value2)	
	fmt.Println(value1 == value2)	
	fmt.Println(value1 != value2)	
}