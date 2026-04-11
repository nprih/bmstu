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
		printRes()
		inputNext()
	}
}

func printHeader() {
	fmt.Printf("%s\n%s\n\n%s", enum.Line, enum.TitleOne, enum.ReqOne)
}

func printRes() {
	fmt.Print(enum.ResOne, " ", utf8.RuneCountInString(strings.TrimSpace(text)), "\n\n", enum.FutterOne, "\n\n")
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
	fmt.Print(enum.QuestionNext)
	text, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	if strings.ToLower(strings.TrimSpace(text)) == "y" {
		stop = enum.Quit
		fmt.Println(enum.Line)
	}
}
