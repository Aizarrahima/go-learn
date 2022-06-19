package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	data := ring.New(5)
	for i := 0; i<data.Len(); i++ { // len => jumlah ring
		data.Value = "Data" + strconv.FormatInt(int64(i), 10)
		data = data.Next() // supaya mengarah pada ring berikutnya
	}

	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}