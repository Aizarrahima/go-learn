package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKTPku NoKTP = "1230019121271" // tidak memerlukan string karena pada NoKTP sudah alias pada string
	var marriedStatus Married = false

	fmt.Println(noKTPku)
	fmt.Println(marriedStatus)
}