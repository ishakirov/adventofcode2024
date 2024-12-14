package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	ruleRegex = regexp.MustCompile(`^(\d+)\|(\d+)`)
)

func parseRule(line string) (int, int, error) {
	matches := ruleRegex.FindStringSubmatch(line)
	if matches == nil {
		return 0, 0, fmt.Errorf("string [%s] doesnt contain a rule", line)
	}

	a, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, 0, err
	}
	b, err := strconv.Atoi(matches[2])
	if err != nil {
		return 0, 0, err
	}

	return a, b, nil
}

func isUpdateCorrect(update []int, rules map[int][]int) bool {
	indexesMap := make(map[int]int)
	for i, number := range update {
		indexesMap[number] = i
	}

	for i, number := range update {
		if _, ok := rules[number]; ok {
			for _, higher := range rules[number] {
				if higherIndex, ok := indexesMap[higher]; ok {
					if i > higherIndex {
						return false
					}
				}
			}
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

	rulesMap := make(map[int][]int)
	answer := 0

	scanRules := true
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			scanRules = false
			continue
		}

		if scanRules {
			a, b, err := parseRule(line)
			if err != nil {
				log.Fatal(err)
			}

			if _, ok := rulesMap[a]; !ok {
				rulesMap[a] = []int{}
			}
			rulesMap[a] = append(rulesMap[a], b)
		} else {
			update := []int{}
			for _, number := range strings.Split(line, ",") {
				n, err := strconv.Atoi(number)
				if err != nil {
					log.Fatal(err)
				}

				update = append(update, n)
			}

			if isUpdateCorrect(update, rulesMap) {
				log.Printf("Update %v is correct, adding %d to answer", update, update[len(update)/2])
				answer += update[len(update)/2]
			}
		}
	}

	log.Printf("Answer: %d", answer)
}
