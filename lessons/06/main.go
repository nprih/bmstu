package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calc(text string) {
	i, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		fmt.Println("Вы ввели не число")
		return
	}

	res := i % 2
	response := ""
	if res == 0 {
		response = "четное"
	} else {
		response = "не четное"
	}

	fmt.Println("Ваше число:", strings.TrimSpace(response))

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("New line: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		calc(text)
	}
}
