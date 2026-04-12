package capitalizeWords

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tasks/enum"
)

var text string
var err error
var stop string

func Run() {
	stop = ""
	for stop != enum.Quit {
		printHeader()
		inputNext()
	}
}

func printHeader() {
	fmt.Printf("%s\n%s\n\n%s", enum.Line, enum.TitleThree, enum.ReqText)
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
