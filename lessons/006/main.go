package main

import (
	"fmt"
	"math"
)

func main() {
	//var a int = 10 // int->float64
	//var b float32 = float32(a)

	var c float64 = 12.999999
	var d int = int(math.Round(c))

	//fmt.Println(b)
	fmt.Println(d)
}
