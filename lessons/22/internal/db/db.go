package db

import (
	"database/sql"
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
		err := rows.Scan(&task.Id, &task.ClientId, &task.Text, &task.Created, &task.Answer, &task.Done)
		if err != nil {
			log.Println(err)
			continue
		}
		tasks = append(tasks, task)
	}
	return tasks
}
