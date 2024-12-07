package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	mulRegex   = regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	mulEnabled = true
)

func findMultiplicationsSum(s string) int {
	matches := mulRegex.FindAllStringSubmatch(s, -1)
	if matches == nil {
		return 0
	}

	sum := 0
	for _, match := range matches {
		if match[0] == "do()" {
			mulEnabled = true
		} else if match[0] == "don't()" {
			mulEnabled = false
		} else if mulEnabled {
			a, err := strconv.Atoi(match[1])
			if err != nil {
				log.Fatal(err)
			}
			b, err := strconv.Atoi(match[2])
			if err != nil {
				log.Fatal(err)
			}

			sum += a * b
		}
	}

	return sum
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	answer := 0
	for scanner.Scan() {
		line := scanner.Text()
		answer += findMultiplicationsSum(line)
	}

	log.Printf("Answer: %d", answer)
}
