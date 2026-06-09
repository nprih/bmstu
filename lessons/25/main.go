package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"log"
)

func main() {
	//Generate keys
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Println(err)
		return
	}
	publicKey := &privateKey.PublicKey
	message := []byte("Hello my friends!")
	//create chipher
	ciphertext, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		publicKey,
		message,
		nil,
	)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("Message: %s\n", message)
	fmt.Printf("Cipher: %x\n", ciphertext)
	//clear text
	clearText, err := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		privateKey,
		ciphertext,
		nil,
	)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("Message: %s\n", clearText)
}
