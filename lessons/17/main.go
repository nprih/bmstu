package main

import (
	"net/http"
	_ "net/http/pprof"
)

const sliceSize = 10_000_000

func foo() {
	for {
		var s []int
		for i := 0; i < sliceSize; i++ {
			s = append(s, i)
		}
	}
}

func main() {
	go foo()
	http.ListenAndServe(":8080", nil)
}
