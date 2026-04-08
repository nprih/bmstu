package calculateLen

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

const (
	titleOne = "Выполняется задача №1"
	reqOne   = "Введите строку: "
	resOne   = "Количество символов в введенной строке:"
	futter   = "Задача №1 решена"
)

var text string
var err error

func Run() {
	printHeader()
	inputText()
	fmt.Print(resOne, " ", utf8.RuneCountInString(strings.TrimSpace(text)), "\n\n", futter, "\n\n")
}

func printHeader() {
	fmt.Printf("%s\n\n", titleOne)
	fmt.Print(reqOne)
}

func inputText() {
	reader := bufio.NewReader(os.Stdin)
	text, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
}
