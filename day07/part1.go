package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func evaluate(result int, args []int, bitmask uint) bool {
	acc := args[0]
	for i := 1; i < len(args); i++ {
		if bitmask&(1<<uint(i-1)) != 0 {
			acc += args[i]
		} else {
			acc *= args[i]
		}
	}

	return acc == result
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	answer := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		calibrationResultString := parts[0][:len(parts[0])-1]
		calibrationResult, err := strconv.Atoi(calibrationResultString)
		if err != nil {
			log.Fatal(err)
		}

		args := make([]int, 0, len(parts)-1)
		for _, part := range parts[1:] {
			arg, err := strconv.Atoi(part)
			if err != nil {
				log.Fatal(err)
			}
			args = append(args, arg)
		}

		for i := 0; i < 1<<uint(len(args)-1); i++ {
			if evaluate(calibrationResult, args, uint(i)) {
				answer += calibrationResult
				break
			}
		}

	}

	log.Println(answer)
}
