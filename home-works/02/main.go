package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tasks/calculateLen"
	"tasks/calculateVowel"
)

const (
	title      = "Укажите номер задачи, которую вы хотите выполнить:"
	one        = "Напишите программу, которая запрашивает у пользователя ввод строки, а затем выводит число - количество символов в строке"
	two        = "Напишите программу, которая подсчитывает количество гласных букв (а, е, ё, и, о, у, ы, э, ю, я) в введённой пользователем строке"
	three      = "Создайте функцию capitalizeWords(s string) string, которая преобразует каждое слово в строке так, чтобы первая буква была заглавной, а остальные — строчными"
	four       = "Напишите программу, которая запрашивает у пользователя ввод строки-формулы, а выводит сообщение о правильности написания круглых скобок"
	five       = "Выход из приложения"
	selected   = "Вы выбрали: "
	notCorrect = " - не корректное число\n"
	line       = "\n------------------------------------------------------------------------------------------------------------------------------------------------------------\n"
	quit       = "quit"
	exit       = "Выход"
)

var numTask int
var stop string

func main() {
	for stop != quit {
		numTask = 0
		taskSelection()
		fmt.Println(line)
		taskSolution()
	}
}

func taskSolution() {
	switch numTask {
	case 1:
		calculateLen.Run()
	case 2:
		calculateVowel.Run()
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
	return []string{one, two, three, four, five}
}

func printTasks(tasks []string) {
	listTasks := ""
	for i := 0; i < len(tasks); i++ {
		listTasks += fmt.Sprintln(i+1, "-", tasks[i])
		if i == len(tasks)-2 {
			listTasks += fmt.Sprint("\n")
		}
	}
	listTasks += fmt.Sprintf("\n%s ", title)
	fmt.Print(listTasks)
}

func checkNumber(text string) {
	numberTask, err := strconv.Atoi(strings.TrimSpace(text))
	switch true {
	case err != nil:
		fmt.Print(strings.TrimSpace(text), notCorrect)
		return
	case numberTask <= 0:
		fmt.Print(strings.TrimSpace(text), notCorrect)
		return
	case numberTask == 5:
		fmt.Print(exit, "\n")
		stop = quit
		return
	case numberTask > 5:
		fmt.Print("У нас не так много задач\n")
		return
	default:
		numTask = numberTask
		fmt.Print(selected, numTask, "\n")
	}
}
