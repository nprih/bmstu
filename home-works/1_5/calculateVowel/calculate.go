package calculateVowel

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tasks/enum"
)

var text string
var sum int
var err error
var stop string

func Run() {
	stop = ""
	for stop != enum.Quit {
		printHeader()
		inputText()
		calcVowels()
		printRes()
		inputNext()
	}
}

func printHeader() {
	fmt.Printf("%s\n%s №2\n\n%s", enum.Line, enum.TitleTask, enum.ReqText)
}

func inputText() {
	reader := bufio.NewReader(os.Stdin)
	text, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
}

func calcVowels() {
	sum = 0
	for _, char := range strings.TrimSpace(strings.ToLower(text)) {
		for _, vowel := range enum.Vowels {
			if char == vowel {
				sum++
			}
		}
	}
}

func printRes() {
	fmt.Print(enum.ResTwo, " ", sum, "\n\n", getFooter(), "\n\n")
}

func getFooter() string {
	futter := strings.Split(strings.TrimSpace(enum.Futter), " ")
	return fmt.Sprintf("%s №2 %s", futter[0], futter[1])
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
