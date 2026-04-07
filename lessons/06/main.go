package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("New line: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Ur text: ", strings.TrimSpace(text))
}
