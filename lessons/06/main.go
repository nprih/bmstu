package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func run() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Новая строка: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		i, errRes := strconv.Atoi(strings.TrimSpace(text))
		printRes(i, errRes)
	}
}

func printRes(i int, err error) {
	res := "Вы ввели"
	if err != nil {
		fmt.Println(res, "не число")
		return
	}
	if i%2 == 0 {
		res += " четное"
	} else {
		res += " не четное"
	}
	fmt.Println(res, "число")
}

func main() {
	run()
}
