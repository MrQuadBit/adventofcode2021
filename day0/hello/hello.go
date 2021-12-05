package main

import (
	"fmt"
	"log"

	"Desktop/adventofcode/adventofcode2021/day1/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
