package main

import "fmt"

type counter int

func (c *counter) incr() {
	*c++
}

func main() {
	var c counter = 0
	c.incr()
	fmt.Println(c)
}
