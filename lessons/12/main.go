package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const CRAW = os.O_CREATE | os.O_RDWR | os.O_APPEND

var dirName string
var dir []os.DirEntry
var err error
var fileText string

func main() {
	setDirName()
	setDir()
	setFileString()
	writeFile()
}

func setDirName() {
	fmt.Print("Введите название папки: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	dirName = strings.TrimSpace(text)
}

func setDir() {
	dir, err = os.ReadDir(dirName)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func setFileString() {
	var fileType string
	for _, file := range dir {
		if file.IsDir() {
			fileType = "папка"
		} else {
			fileType = "файл"
		}
		fileText += fmt.Sprintf("%s - это %s\n", file.Name(), fileType)
	}
	fileText += "\n"
}

func writeFile() {
	file, err := os.OpenFile("result.txt", CRAW, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(fileText)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Отчет представлен в файле ./result.txt")
}
