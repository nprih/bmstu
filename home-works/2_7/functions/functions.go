package functions

import (
	"2_7/argon2"
	"2_7/models"
	"unicode"
)

func ToLowerCase(text string) string {
	char := []rune(text)
	for i := range char {
		char[i] = unicode.ToLower(char[i])
	}
	return string(char)
}

func UsersHashingPass() []models.User {
	users := models.DefaultUsers
	for i, user := range users {
		users[i].Password, _ = argon2.HashPasswordArgon2(user.Password)
	}
	return users
}
