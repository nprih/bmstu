package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
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
		return "", fmt.Errorf("ошибка выполнения команды: %w", err)
	}
	return string(output), nil
}

func askForTasks(id int) (Task, error) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/ask-task?id=%d", id))
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

func sendAnswerToServer(result ExecutionResult) error {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return err
	}
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/ask-task", bytes.NewBuffer(data))
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
		fmt.Println("Answer wasn't sent")
		return fmt.Errorf("Answer wasn't sent")
	} else {
		fmt.Println("Answer was sent")
		return nil
	}

}

func main() {
	for {
		task, err := askForTasks(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		if task.Id == 0 {
			return
		}

		var result ExecutionResult
		tsk := strings.Split(task.Text, " ")
		answer, err := executeTask(tsk[0], tsk[1:]...)
		result.Id = task.Id

		if err != nil {
			result.Status = "none"
			result.Answer = err.Error()
			if err = sendAnswerToServer(result); err != nil {
				fmt.Println(err)
				return
			}
			return
		}
		result.Status = "ok"
		result.Answer = answer
		err = sendAnswerToServer(result)
		if err != nil {
			log.Println(err)
			return
		}
		time.Sleep(3 * time.Second)
	}
}
