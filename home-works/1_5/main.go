package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tasks/calculateLen"
	"tasks/calculateVowel"
	"tasks/capitalizeWords"
	"tasks/checkFormula"
	"tasks/enum"
)

var numTask int
var stop string

func main() {
	for stop != enum.Quit {
		numTask = 0
		taskSelection()
		taskSolution()
	}
}

func taskSolution() {
	switch numTask {
	case 1:
		calculateLen.Run()
	case 2:
		calculateVowel.Run()
	case 3:
		capitalizeWords.Run()
	case 4:
		checkFormula.Run()
	}
}

func taskSelection() {
	printTasks(getTasks())
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	checkNumber(text)
}

func getTasks() []string {
	return []string{enum.One, enum.Two, enum.Three, enum.Four, enum.Five}
}

func printTasks(tasks []string) {
	listTasks := ""
	for i := 0; i < len(tasks); i++ {
		listTasks += fmt.Sprintln(i+1, "-", tasks[i])
		if i == len(tasks)-2 {
			listTasks += fmt.Sprint("\n")
		}
	}
	listTasks += fmt.Sprintf("\n%s ", enum.Title)
	fmt.Print(listTasks)
}

func checkNumber(text string) {
	numberTask, err := strconv.Atoi(strings.TrimSpace(text))
	switch true {
	case err != nil:
		fmt.Print(strings.TrimSpace(text), enum.NotCorrect)
		return
	case numberTask <= 0:
		fmt.Print(strings.TrimSpace(text), enum.NotCorrect)
		return
	case numberTask == 5:
		fmt.Print(enum.Exit, "\n")
		stop = enum.Quit
		return
	case numberTask > 5:
		fmt.Print("У нас не так много задач\n")
		return
	default:
		numTask = numberTask
		fmt.Print(enum.Selected, numTask, "\n")
	}
}
