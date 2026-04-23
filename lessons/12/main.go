package main

import (
	"fmt"
	"io"
	"os"
)

type User struct {
	Email    string
	Password string
	Age      int
}

func main() {
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//data, err := io.ReadAll(file)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(data))
	data := make([]byte, 64)
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		fmt.Println(string(data[:n]))
	}
}
