package main

import (
	"fmt"
	"os"
)

func main() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Can't read directory =(\n", err)
		return
		//os.Exit(1)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
