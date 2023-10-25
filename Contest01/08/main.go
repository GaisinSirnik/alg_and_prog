package main

import "fmt"

func main() {

	var a int
	var b int
	var c int
	fmt.Scan(&a, &b, &c)
	if a > b && a > c {
		fmt.Println(a)
	}
	if b > a && b > c {
		fmt.Println(b)
	}
	if c > a && c > b {
		fmt.Println(c)
	}

}