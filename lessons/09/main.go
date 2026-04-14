package main

import "fmt"

type location struct {
	lat, long float64
}

func main() {
	var solaris location = location{10.11, 12.12}

	fmt.Printf("%+v\n", solaris)
}
