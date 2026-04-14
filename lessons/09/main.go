package main

import "fmt"

type location struct {
	lat, long float64
}

func main() {
	var solaris location = location{long: 12.12, lat: 10.11}

	fmt.Printf("%+v\n", solaris)
}
