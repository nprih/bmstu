package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile(file *os.File) {
	scanner := bufio.NewScanner(file)
	errorCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(strings.ToLower(line), "error") {
			errorCount++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Ошибка чтения файла: %v", err)
	}
	fmt.Printf("Количество строк, содержащих 'error': %d\n", errorCount)
}

func main() {
	file, err := os.Open("server.log")
	if err != nil {
		log.Fatalf("Ошибка открытия файла: %v", err)
	}
	defer file.Close()
	readFile(file)
}
