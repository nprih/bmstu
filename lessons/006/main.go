package main

import (
	"fmt"
	"strconv"
)

func main() {
	//10 -> "10"
	//var pi rune = 1009 //rune/int32 -> string
	//fmt.Println(string(pi))
	//strconv
	//Atoi ascii to integer| Itoa integer to ascii
	var a int = 10
	var b string = strconv.Itoa(a)

	fmt.Printf("%v %T\n", b, b)

}
