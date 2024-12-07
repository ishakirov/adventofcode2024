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

func isReportSafeWithProblemDampener(report []int) bool {
	if isReportSafe(report) {
		return true
	}

	for i := 0; i < len(report); i++ {
		modifiedReport := make([]int, 0,for len(report)-1)
		modifiedReport = append(modifiedReport, report[:i]...)
		modifiedReport = append(modifiedReport, report[i+1:]...)
		if isReportSafe(modifiedReport) {
			return true
		}
	}

	return false
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

		log.Println(in)
		if isReportSafeWithProblemDampener(in) {
			log.Println(in, "is safe")
			safeReportsCount++
		} else {
			log.Println(in, "is not safe")
		}

		in = in[:0]
	}

	log.Printf("Safe reports count: %d", safeReportsCount)
}
