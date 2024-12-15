package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func incrOps(ops []int) []int {
	for i := 0; i < len(ops); i++ {
		ops[i] = (ops[i] + 1) % 3
		if ops[i] != 0 {
			break
		}
	}

	return ops
}

func intPow(x, y int) int {
	acc := 1
	for i := 0; i < y; i++ {
		acc *= x
	}
	return acc
}

func test(result int, args []int) bool {
	ops := make([]int, len(args)-1)

	for i := 0; i < intPow(3, len(ops)); i++ {
		acc := args[0]
		for i := 1; i < len(args); i++ {
			op := ops[i-1]
			switch op {
			case 0:
				acc += args[i]
			case 1:
				acc *= args[i]
			case 2:
				str1 := strconv.Itoa(acc)
				str2 := strconv.Itoa(args[i])
				acc, _ = strconv.Atoi(str1 + str2)
			}
		}

		if acc == result {
			return true
		}

		ops = incrOps(ops)
	}

	return false
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

		if test(calibrationResult, args) {
			log.Printf("%d: %v is correct!", calibrationResult, args)
			answer += calibrationResult
		} else {
			log.Printf("%d: %v is incorrect!", calibrationResult, args)
		}
	}

	log.Println(answer)
}
