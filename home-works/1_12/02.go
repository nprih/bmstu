package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var textSummary string

func getFileNames(files []os.DirEntry) (fileNames []string) {
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".txt" {
			fileNames = append(fileNames, file.Name())
		}
	}
	return fileNames
}

func readFiles(fileNames []string) {
	for _, fileName := range fileNames {
		text, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
			return
		}
		textSummary += string(text) + "\n\n"
	}
}

func writeFile() {
	file, err := os.Create("combined.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, textSummary)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	files, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
		return
	}

	readFiles(getFileNames(files))
	writeFile()
}
