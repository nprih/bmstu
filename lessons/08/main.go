package main

import "fmt"

func AnimalTalk(who string, how func(string) string) {
	fmt.Println(how(who))
}

func main() {
	AnimalTalk("dog", func(animal string) string { return "Wuf!" })
}
