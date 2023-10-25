package main

import "fmt"

func main() {
	var a, b, c string
	fmt.Scanln(&a, &b, &c)
	plavaet := a == "Да"
	letaet := b == "Да"
	dlina := c == "Да"
	if plavaet {
		if letaet {
			if dlina {
				fmt.Println("Утка")
			} else {
				fmt.Println("Пингвин")
			}
		} else {
			if dlina {
				fmt.Println("Плезиозавры")
			} else {
				fmt.Println("Дельфин")
			}
		}
	} else {

		if letaet {
			if dlina {
				fmt.Println("Страус")
			} else {
				fmt.Println("Курица")
			}
		} else {

			if dlina {
				fmt.Println("Жираф")
			} else {
				fmt.Println("Кот")
			}
		}
	}

}