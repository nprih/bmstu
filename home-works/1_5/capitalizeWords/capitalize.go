package capitalizeWords

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tasks/enum"
	"unicode"
)

var text string
var err error
var stop string
var firstCharUpperWords []string

func Run() {
	stop = ""
	for stop != enum.Quit {
		printHeader()
		inputText()
		capitalizeWords()
		printRes()
		inputNext()
	}
}

func printHeader() {
	fmt.Printf("%s\n%s №3\n\n%s", enum.Line, enum.TitleTask, enum.ReqText)
}

func inputText() {
	reader := bufio.NewReader(os.Stdin)
	text, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
}

func capitalizeWords() {
	words := strings.Split(strings.TrimSpace(text), " ")
	firstCharUpperWords = firstCharUpperWords[:0]
	for _, word := range words {
		char := []rune(word)
		for i := range char {
			if i == 0 {
				char[i] = unicode.ToUpper(char[i])
			} else {
				char[i] = unicode.ToLower(char[i])
			}
		}
		firstCharUpperWords = append(firstCharUpperWords, string(char))
	}
}

func printRes() {
	fmt.Print(enum.ResThree, " ", strings.Join(firstCharUpperWords, " "),
		"\n\n", getFooter(), "\n\n",
	)
}

func getFooter() string {
	futter := strings.Split(strings.TrimSpace(enum.Futter), " ")
	return fmt.Sprintf("%s №3 %s", futter[0], futter[1])
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
