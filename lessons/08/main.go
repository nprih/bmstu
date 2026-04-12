package main

import "fmt"

func Say(animal string) string {
	if animal == "cat" {
		return "Meuw!"
	} else if animal == "dog" {
		return "Wuf!"
	}
	return "I dont know!"
}

func AnimalTalk(who string, how func(string) string) {
	fmt.Println(how(who))
}

func main() {
	var voice func(string) string
	voice = Say
	AnimalTalk("cat", voice)
}
