package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("a([a-z])r")

	fmt.Println(regex.MatchString("aizar")) // false => lebih dari 3 karakter
	fmt.Println(regex.MatchString("air")) // true => sesuai
	fmt.Println(regex.MatchString("aIr")) // false => I huruf kapital

	search := regex.FindAllString("air ayr aer aur azr aizar", -1) // -1 => tidak terbatas
	fmt.Println(search) // [air ayr aer aur azr] => "aizar" tidak dapat tampil karena tidak sesuai dengan ketentuan yang telah dibuat
}