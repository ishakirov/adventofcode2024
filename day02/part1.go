package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isReportSafe(report []int) bool {
	if len(report) < 2 {
		return true
	}

	increasing := report[0] < report[1]
	for i := 0; i < len(report)-1; i++ {
		diff := abs(report[i] - report[i+1])
		if diff < 1 || diff > 3 {
			return false
		}

		if increasing && report[i] >= report[i+1] {
			return false
		} else if !increasing && report[i] <= report[i+1] {
			return false
		}
	}

	return true
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	in := make([]int, 0)

	safeReportsCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		for _, v := range parts {
			num, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}

			in = append(in, num)
		}

		if isReportSafe(in) {
			safeReportsCount++
		}

		in = in[:0]
	}

	log.Printf("Safe reports count: %d", safeReportsCount)
}
