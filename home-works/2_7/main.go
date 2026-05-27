package main

import (
	"bufio"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"golang.org/x/crypto/argon2"
	_ "modernc.org/sqlite"
)

type Argon2Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

func HashPasswordArgon2(password string) (string, error) {
	params := &Argon2Params{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	salt := make([]byte, params.SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, params.Iterations, params.Memory, params.Parallelism, params.KeyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		params.Memory, params.Iterations, params.Parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

type User struct {
	Id       int64
	Email    string
	Password string
}

func getDefaultUsers() []User {
	return []User{
		{
			Email:    "first@test.ru",
			Password: "123321",
		},
		{
			Email:    "second@test.ru",
			Password: "123qwe",
		},
		{
			Email:    "third@test.ru",
			Password: "ASD123",
		},
	}
}

func usersHashingPass() []User {
	users := getDefaultUsers()
	for i, user := range users {
		users[i].Password, _ = HashPasswordArgon2(user.Password)
	}
	return users
}

func dbConnection() (error, *sql.DB) {
	db, err := sql.Open("sqlite", "users")
	if err != nil {
		return err, db
	}
	return nil, db
}

func insertDefaultUsers() {
	err, db := dbConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	clearDb(db)

	var placeholders []string
	var args []interface{}
	for _, user := range usersHashingPass() {
		placeholders = append(placeholders, "(?, ?)")
		args = append(args, user.Email, user.Password)
	}

	query := `INSERT INTO users (email, password) VALUES ` + strings.Join(placeholders, ", ")

	_, err = db.Exec(query, args...)
	if err != nil {
		log.Fatal(err)
	}
}

func clearDb(db *sql.DB) {
	_, err := db.Exec("DELETE FROM users")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("DELETE FROM sqlite_sequence WHERE name = 'users'")
	if err != nil {
		log.Fatal(err)
	}
}

func toLowerCase(text string) string {
	text = strings.TrimSpace(text)
	char := []rune(text)
	for i := range char {
		char[i] = unicode.ToLower(char[i])
	}
	return string(char)
}

func main() {
	insertDefaultUsers()

	fmt.Print("Введите логин (всё что до символа: '@' в указанном вами email): ")
	reader := bufio.NewReader(os.Stdin)
	login, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(toLowerCase(login))
}
