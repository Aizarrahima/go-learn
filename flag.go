package main

import (
	"flag"
	"fmt"
)

func main() {
	// hasil dari variable berupa pointer
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")
	number := flag.Int("number", 100, "Put your number")

	flag.Parse()

	// fmt.Println(*host, *username, *password)
	fmt.Println("Host : ", *host )
	fmt.Println("Username : ", *username)
	fmt.Println("Password : ", *password)
	fmt.Println("Number : ", *number)

	// note: apabila kita memasukkan value tidak sesuai dengan tipe datanya, maka akan muncul helper nya sesuai dengan default.
}