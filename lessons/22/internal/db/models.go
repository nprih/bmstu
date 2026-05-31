package db

type Client struct {
	Id        int
	Name      string
	Created   string
	WasOnline string
}

type Task struct {
	Id       int
	ClientId int
	Text     string
	Created  string
	Answer   string
	Done     int
}
