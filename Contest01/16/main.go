package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scanln(&a)
	b := make([]float64, a)
	for i := range b {
		fmt.Scan(&b[i])
	}
	buff := make([]any, a)
	for i, v := range b {
		if i == 0 || i == a-1 {
			buff[i] = v
			continue
		}
		buff[i] = (b[i-1] + v + b[i+1]) / 3
	}
	fmt.Println(buff...)
}