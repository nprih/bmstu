package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"
)

var id int = 3

type Task struct {
	Id   int
	Text string
}

type ExecutionResult struct {
	Id     int
	Status string
	Answer string
}

func executeTask(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("ошибка выполнения команды: %s", err)
	}
	return string(output), nil
}

func askForTasks(id int) (Task, error) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8090/ask-tasks?id=%d", id))
	if err != nil {
		return Task{}, err
	}

	var myTask Task
	err = json.NewDecoder(resp.Body).Decode(&myTask)
	if err != nil {
		return Task{}, err
	}
	return myTask, nil
}

func SendAnswerToServer(result ExecutionResult) error {
	//resp, err := http.Post("http://localhost:8090/answer", "application/json", strings.NewReader(result.Answer))
	client := &http.Client{
		Timeout: time.Second * 3,
	}
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return err
	}
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8090/ask-task", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Println("Answer wosn't sent")
		return fmt.Errorf("Answer wasn't sent")
	} else {
		fmt.Printf("Answer was sent")
		return nil
	}
}

func main() {
	task, err := askForTasks(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	if task.Id == 0 {
		fmt.Println(err)
		return
	}

	var result ExecutionResult
	answer, err := executeTask(task.Text)
	result.Id = task.Id
	if err != nil {

		result.Status = "none"
		result.Answer = err.Error()
		if err = SendAnswerToServer(result); err != nil {
			return
		}
	}
	result.Status = "ok"
	result.Answer = answer
	err = SendAnswerToServer(result)
	if err != nil {
		log.Println(err)
		return
	}
}
