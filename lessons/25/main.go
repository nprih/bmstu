package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	data := []byte("Hello my friends!")
	hash := sha256.Sum256(data)
	fmt.Printf("%x\n", hash)
}
