package main

import (
	"fmt"
	"strconv"
)

func main() {
	//10 -> "10"
	//strconv
	//Atoi ascii to integer| Itoa integer to ascii

	var a string = "10"
	b, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v %T\n", b, b)

}
