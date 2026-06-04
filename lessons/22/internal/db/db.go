package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "modernc.org/sqlite"
)

/*
SelectAllClients() []Client
SelectAllTasks() []Tasks
*/

func SelectAllClients() []Client {
	db, err := sql.Open("sqlite", "lesson22")
	if err != nil {
		log.Println(err)
		return []Client{}
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("ping db:", err)
	}

	rows, err := db.Query("SELECT * FROM clients")
	if err != nil {
		log.Println(err)
		return []Client{}
	}

	defer rows.Close()
	clients := []Client{}

	for rows.Next() {
		ph := Client{}
		err := rows.Scan(&ph.Id, &ph.Name, &ph.Created, &ph.WasOnline)
		if err != nil {
			log.Println(err)
			continue
		}
		clients = append(clients, ph)
	}
	return clients
}

func SelectAllTasks() []Task {
	db, err := sql.Open("sqlite", "lesson22")
	if err != nil {
		log.Println(err)
		return []Task{}
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("ping db:", err)
	}

	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Println(err)
		return []Task{}
	}

	defer rows.Close()
	tasks := []Task{}

	for rows.Next() {
		ph := Task{}
		var TaskAnswer sql.NullString
		err := rows.Scan(&ph.Id, &ph.ClientId, &ph.Text, &ph.Created, &TaskAnswer, &ph.Done)
		if err != nil {
			log.Println(err)
			continue
		}
		if TaskAnswer.Valid {
			ph.Answer = TaskAnswer.String
		} else {
			ph.Answer = ""
		}
		tasks = append(tasks, ph)
	}
	return tasks
}

func CreateNewTask(newTask Task) error {
	db, err := sql.Open("sqlite", "lesson22")
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO tasks (clientId, text, created, done) VALUES ($1, $2, $3, $4)", newTask.ClientId, newTask.Text, newTask.Created, newTask.Done)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func UpdateTask(clientAnswer TaskAnswer) error {
	db, err := sql.Open("sqlite", "lesson22")
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	var statusCode int
	switch clientAnswer.Status {
	case "ok":
		statusCode = 1
	case "none":
		statusCode = 2
	}

	_, err = db.Exec("UPDATE tasks SET answer=$1, done=$2 WHERE id=$3", clientAnswer.Answer, statusCode, clientAnswer.Id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func InsertNewClient(clientId int) error {
	db, err := sql.Open("sqlite", "lesson22")
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	_, err = db.Exec(`INSERT INTO clients (id, name, created, online) VALUES ($1, $2, $3, $4)`,
		clientId, fmt.Sprintf("Client-%d", clientId), time.Now().Format("02 Jan 2006"), time.Now().Format("02 Jan 2006"))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
