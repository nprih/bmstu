package main

import (
	"fmt"
	"os"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Enter your name: ")
	//input, err := reader.ReadString('\n')
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(input)

	fmt.Fprintln(os.Stdout, "Hello", "World")
}
