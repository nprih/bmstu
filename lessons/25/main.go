package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func GenerateHMAC(key, message []byte) string {
	h := hmac.New(sha256.New, key)
	h.Write(message)
	return hex.EncodeToString(h.Sum(nil))
}
func VerifyHMAC(key, message []byte, signature string) bool {
	message = []byte("Hello with HMAC")
	h := hmac.New(sha256.New, key)
	h.Write(message)
	expectedSignature := hex.EncodeToString(h.Sum(nil))
	expectedBytes, expErr := hex.DecodeString(expectedSignature)
	receivedBytes, recErr := hex.DecodeString(signature)
	if expErr != nil || recErr != nil {
		return false
	}
	return hmac.Equal(expectedBytes, receivedBytes)
}
func main() {
	key := []byte("secret-key-123")
	message := []byte("Hello with HMAC!")
	signature := GenerateHMAC(key, message)
	fmt.Printf("Message: %s\n", string(message))
	fmt.Printf("HMAC signature: %s\n\n", signature)
	isValid := VerifyHMAC(key, message, signature)
	fmt.Printf("IsValid: %v\n", isValid)
}
