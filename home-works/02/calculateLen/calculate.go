package calculateLen

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tasks/enum"
	"unicode/utf8"
)

var text string
var err error
var stop string

func Run() {
	stop = ""
	for stop != enum.Quit {
		printHeader()
		inputText()
		fmt.Print(enum.ResOne, " ", utf8.RuneCountInString(strings.TrimSpace(text)), "\n\n", enum.FutterOne, "\n\n")
		inputNext()
	}
}

func printHeader() {
	fmt.Printf("%s\n\n", enum.TitleOne)
	fmt.Print(enum.ReqOne)
}

func inputText() {
	reader := bufio.NewReader(os.Stdin)
	text, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
}

func inputNext() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(enum.QuestionOne)
	text, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	if strings.ToLower(strings.TrimSpace(text)) == "n" {
		stop = enum.Quit
	}
}
