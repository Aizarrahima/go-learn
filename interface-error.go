package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (int, error) { // error => interface
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil // terdapat "nil" karena terdapat interface error
	}
}

func main() {
	// var contohError error = errors.New("Ups Error")
	// fmt.Println(contohError.Error())

	hasil, err := Pembagi(100, 0)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}

}