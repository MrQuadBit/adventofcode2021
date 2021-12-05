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
		log.Fatalf("error: %s", e)
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
	check(err)

	measurements_larger := 0
	var previous int
	for i, _ := range lines {

		if i == 0 {
			window_1, err_1 := strconv.Atoi(lines[i])
			check(err_1)
			window_2, err_2 := strconv.Atoi(lines[i+1])
			check(err_2)
			window_3, err_3 := strconv.Atoi(lines[i+2])
			check(err_3)

			previous = window_1 + window_2 + window_3
		} else if i > len(lines)-3 {
			fmt.Printf("lines[%d] = %s\n", i, lines[i])
		} else {
			window_1, err_1 := strconv.Atoi(lines[i])
			check(err_1)
			window_2, err_2 := strconv.Atoi(lines[i+1])
			check(err_2)
			window_3, err_3 := strconv.Atoi(lines[i+2])
			check(err_3)

			window := window_1 + window_2 + window_3
			fmt.Printf("window[%d] = %d, %d, %d = %d\n", i, window_1, window_2, window_3, window)

			if window > previous {
				fmt.Printf("%d is larger than %d\n", window, previous)
				measurements_larger += 1
				previous = window
			} else {
				previous = window
			}
		}
	}
	fmt.Println(measurements_larger)
}
