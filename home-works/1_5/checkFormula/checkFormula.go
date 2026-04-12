package checkFormula

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tasks/enum"
)

var text string
var err error
var openerCount, closingCount int
var stop string

func Run() {
	stop = ""
	for stop != enum.Quit {
		printHeader()
		inputText()
		calcBrackets()
		printRes()
		inputNext()
	}
}

func printHeader() {
	fmt.Printf("%s\n%s №4\n\n%s", enum.Line, enum.TitleTask, enum.ReqFour)
}

func inputText() {
	reader := bufio.NewReader(os.Stdin)
	text, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
}

func calcBrackets() {
	openerCount, closingCount = 0, 0
	for _, char := range strings.TrimSpace(text) {
		if char == '(' {
			openerCount++
		}
		if char == ')' {
			closingCount++
		}
	}
}

func printRes() {
	res := ""
	if openerCount != closingCount {
		res += " неправильно, "
	} else {
		res += " верно, "
	}
	fmt.Print(enum.ResFour, res,
		openerCount, " открывающиеся, ",
		closingCount, " закрывающиеся",
		"\n\n", getFooter(), "\n\n",
	)
}

func getFooter() string {
	futter := strings.Split(strings.TrimSpace(enum.Futter), " ")
	return fmt.Sprintf("%s №4 %s", futter[0], futter[1])
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
