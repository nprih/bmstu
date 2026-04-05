package main

import (
	"fmt"
)

func decrypt(code string) {
	var c byte
	deCode := ""
	for i := 0; i < len(code); i++ {
		c = code[i]
		if c >= 'a' && c <= 'z' {
			c = c - 3
			if c < 'a' {
				c = c + 26
			}
		}
		deCode += string(c)
	}
	fmt.Println(deCode)
}

func main() {
	peace := "execute, my code"
	var c byte
	code := ""

	for i := 0; i < len(peace); i++ {
		c = peace[i]
		if c >= 'a' && c <= 'z' {
			c = c + 3
			if c > 'z' {
				c = c - 26
			}
		}
		code += string(c)
	}

	fmt.Println("\nТеперь дешифруем:")

	decrypt(code)
}
