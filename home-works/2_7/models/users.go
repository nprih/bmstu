package models

type User struct {
	Id       int64
	Email    string
	Password string
}

var DefaultUsers = []User{
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
