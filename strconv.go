package main

import (
	"fmt"
	"strconv"
)

func main() {
	// boolean
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}


	// int
	number, err := strconv.ParseInt("1000000", 10, 64) // di dalam parameter terdiri dari string, base int, dan bitSize int
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(1000000, 10) // di dalam parameter terdiri dari int, base int(decimal, biner, hexadecimal)
	fmt.Println(value)


	// string to int
	valueInt, err := strconv.Atoi("2000000")
	fmt.Println(valueInt)
}