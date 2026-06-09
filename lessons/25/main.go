package main

import (
	"crypto/aes"
	"fmt"
	"log"
)

var key = []byte("examplekey123456")

func main() {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
		return
	}
	myMessage := []byte("Hello my friends!")
	ciphertext := make([]byte, len(myMessage))
	block.Encrypt(ciphertext, myMessage)
	fmt.Printf("%x\n", ciphertext)

	clearMessage := make([]byte, len(ciphertext))
	block.Decrypt(clearMessage, ciphertext)
	fmt.Printf("%s\n", clearMessage)
}
