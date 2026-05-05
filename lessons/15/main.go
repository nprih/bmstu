package main

import (
	"errors"
	"fmt"
)

func main() {
	err := fmt.Errorf("Some %s", "Error")
	err = errors.New(fmt.Sprintf("Another %s", "Error"))
}
