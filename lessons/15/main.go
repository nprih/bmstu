package main

import (
	"fmt"
	"os"
)

func WriteToFile(name string) error {
	f, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = fmt.Fprintln(f, "Lesson 15 golang")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return err
	}
	//some logic
	_, err = fmt.Fprintln(f, "Finished Writing")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return err
	}
	f.Close()
	return nil
}

func main() {
	err := WriteToFile("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("All is OK!")
}
