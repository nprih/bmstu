package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func SelectAllClients() []Client {
	db, err := sql.Open("sqlite", "lesson22")
	if err != nil {
		log.Println(err)
		return []Client{}
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("ping", err)
	}

	rows, err := db.Query("SELECT * FROM clients")
	if err != nil {
		log.Println(err)
		return []Client{}
	}
	defer rows.Close()

	clients := []Client{}
	for rows.Next() {
		client := Client{}
		err := rows.Scan(&client.Id, &client.Name, &client.Created, &client.WasOnline)
		if err != nil {
			log.Println(err)
			continue
		}
		clients = append(clients, client)
	}
	return clients
}

func SelectAllTask() []Task {
	db, err := sql.Open("sqlite", "lesson22")
	if err != nil {
		log.Println(err)
		return []Task{}
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("ping", err)
	}

	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Println(err)
		return []Task{}
	}
	defer rows.Close()

	tasks := []Task{}
	for rows.Next() {
		task := Task{}
		var TaskAnswer sql.NullString
		err := rows.Scan(&task.Id, &task.ClientId, &task.Text, &task.Created, &TaskAnswer, &task.Done)
		if err != nil {
			log.Println(err)
			continue
		}
		if TaskAnswer.Valid {
			task.Answer = TaskAnswer.String
		} else {
			task.Answer = ""
		}
		tasks = append(tasks, task)
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

	_, err = db.Exec("INSERT INTO tasks (clientId, text, created, done) VALUES ($1, $2, $3, $4)",
		newTask.ClientId, newTask.Text, newTask.Created, newTask.Done)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func UpdateTask(clientTaskAnswer TaskAnswer) error {
	db, err := sql.Open("sqlite", "lesson22")
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	var statusCode int
	switch clientTaskAnswer.Status {
	case "ok":
		statusCode = 1
	case "none":
		statusCode = 2
	}

	_, err = db.Exec("UPDATE tasks SET answer=$1, done=$2 WHERE id=$3", clientTaskAnswer.Answer, statusCode, clientTaskAnswer.Id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
