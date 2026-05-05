package main

import (
	"errors"
	"fmt"
	"os"
)

func readConfig() (string, error) {
	_, err := os.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("File not in directory: %w", err)
	}
	return "", nil
}

func main() {
	_, err := readConfig()
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("OMG! Not exists")
	}
}
