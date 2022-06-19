package main

import (
	"fmt"
	"time"
)

func main() {
	// sekarang
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	
	// membuat waktu manual
	utc := time.Date(2022, time.June, 20, 20, 0, 0,0, time.UTC)
	fmt.Println(utc)

	
	// melakukan parsing (dari string)
	// cara 1
	parse, _ := time.Parse(time.RFC3339, "2022-10-20T15:04:05Z") // time.RFC3339 => format waktu
	fmt.Println(parse)
	
	// cara 2
	layout := "2006-01-02"
	parse1, _ := time.Parse(layout, "2020-10-01")
	fmt.Println(parse1)
}