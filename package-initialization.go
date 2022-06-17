package main

import (
	"fmt"
	"go-learn/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}