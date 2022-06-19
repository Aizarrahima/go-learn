package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User // Alias pada user slice

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []User {
		{"Aizar", 18},
		{"Vania", 30},
		{"Lisa", 27},
		{"Dinda", 29},
	}

	fmt.Println(users)
	// melakukan pengurutan
	sort.Sort(UserSlice(users)) // dilakukan konversi berdasarkan kontrak-nya
	fmt.Println(users)
}