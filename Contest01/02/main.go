package main

import  "fmt"
	    

func main() {
	men := 0.5 * 365

	var dub = int(men/20 + 1)
	var topol = int(men/32 + 1)

	fmt.Println(men, topol, dub)

}
	