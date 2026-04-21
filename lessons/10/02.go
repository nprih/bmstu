package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Title    string
	Done     bool
	Executor string
}

var tasks []Task

func showMenu() {
	fmt.Println("~~~ To-Do Лист ~~~")
	fmt.Println("1. Показать все задачи")
	fmt.Println("2. Добавить задачу")
	fmt.Println(`3. Пометить как "Выполнена"`)
	fmt.Println("4. Удалить задачу")
	fmt.Println("5. Выйти")
	fmt.Println("6. Сменить исполнителя")
}

func addTask(reader *bufio.Reader) {
	fmt.Print("Введите название задачи: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Введите исполнителя: ")
	executor, _ := reader.ReadString('\n')
	executor = strings.TrimSpace(executor)

	tasks = append(tasks, Task{Title: title, Done: false, Executor: executor})
	fmt.Println("Задача добавлена!")
}

func showTasks() {
	if len(tasks) == 0 {
		fmt.Println("Список задач пуст")
		return
	}
	fmt.Println("~~~ Список Задач ~~~")
	for i, task := range tasks {
		status := "[ ]"
		if task.Done {
			status = "[✓]"
		}
		fmt.Printf("%d. %s %s (%s)\n", i+1, status, task.Title, task.Executor)
	}
}

func completeTask(reader *bufio.Reader) {
	showTasks()
	fmt.Print("Введите номер задачи: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	index, _ := strconv.Atoi(input)
	tasks[index-1].Done = true
	fmt.Println("Задача выполнена!")
}

func deleteTask(reader *bufio.Reader) {
	showTasks()
	fmt.Print("Введите номер задачи: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	index, _ := strconv.Atoi(input)
	tasks = append(tasks[:index-1], tasks[index:]...)
	fmt.Println("Задача удалена!")
}

func changeExecutor(reader *bufio.Reader) {
	showTasks()
	fmt.Print("Введите номер задачи: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	index, _ := strconv.Atoi(input)

	fmt.Print("Введите нового исполнителя: ")
	executor, _ := reader.ReadString('\n')
	executor = strings.TrimSpace(executor)

	tasks[index-1].Executor = executor
	fmt.Println("Исполнитель изменён!")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		showMenu()
		fmt.Print("Ваш выбор: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			showTasks()
		case "2":
			addTask(reader)
		case "3":
			completeTask(reader)
		case "4":
			deleteTask(reader)
		case "5":
			fmt.Println("До свидания!")
			return
		case "6":
			changeExecutor(reader)
		default:
			fmt.Println("Неверный выбор, попробуй снова!")
		}
	}
}
