package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File name:", file.Name())
	fmt.Println("File size:", stat.Size())
	fmt.Println("File dir:", stat.IsDir())
	fmt.Println("File mode:", stat.Mode())
	fmt.Println("Mod time", stat.ModTime().Format("02.01.2006 15:04"))
	fmt.Println("Mod time", stat.Sys())
}
