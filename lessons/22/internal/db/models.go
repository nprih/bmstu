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

type TaskAnswer struct {
	Id     int
	Status string
	Answer string
}
