package main

import (
	"fmt"
	"strings"
)

func main() {
	text := fmt.Sprintf("I love %s very very much", "cats")
	newText := strings.Replace(text, "very ", "", 1)
	fmt.Println(newText)
}
