package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	question := "¿Cómo estás?"
	engQestion := "How are you?"

	//fmt.Println(len(question))
	//fmt.Println(len(engQestion))
	fmt.Println(utf8.RuneCountInString(question))
	fmt.Println(utf8.RuneCountInString(engQestion))
}
