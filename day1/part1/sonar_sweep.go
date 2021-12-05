package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open("../input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()

}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	measurements_larger := 0
	var previous int
	for i, line := range lines {
		number, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("printing loop: %d -> %s", i, err)
		}
		if i == 0 {
			previous = number
		} else if previous < number {
			measurements_larger += 1
			previous = number
		} else {
			previous = number
		}
	}
	fmt.Println(measurements_larger)

}
