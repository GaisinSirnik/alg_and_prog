package main

import "fmt"

var a = [5]int{5000, 1000, 500, 200, 100}

func main() {
	count := [5]int{0, 0, 0, 0, 0}
	var b int
	fmt.Scanln(&b)
	for i, c := range a {
		count[i] = b / c
		b %= c
	}
	fmt.Println(count[0], count[1], count[2], count[3], count[4])
}