package main

import (
	"fmt"
	"os"
)

func main() {
	//file, err := os.Create("test.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	_, err := os.Open("test.txt")
	if err != nil {
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	//file2, err := os.OpenFile("test.txt") создает с флагами и правами, если его нет

}
