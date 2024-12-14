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

func fixUpdate(update []int, rules map[int][]int) []int {
	numberSet := map[int]struct{}{}
	for _, number := range update {
		numberSet[number] = struct{}{}
	}

	filteredRules := make(map[int][]int)
	for from, toList := range rules {
		if _, ok := numberSet[from]; !ok {
			continue
		}

		for _, to := range toList {
			if _, ok := numberSet[to]; ok {
				if _, ok := filteredRules[from]; !ok {
					filteredRules[from] = []int{}
				}
				filteredRules[from] = append(filteredRules[from], to)
			}
		}
	}

	for _, start := range update {
		topSorted := topSort(start, filteredRules)
		if len(topSorted) == len(update) {
			return topSorted
		}
	}

	return nil
}

func topSort(start int, g map[int][]int) []int {
	visited := map[int]struct{}{}
	result := []int{}

	var dfs func(int)
	dfs = func(v int) {
		if _, ok := visited[v]; ok {
			return
		}

		visited[v] = struct{}{}
		for _, to := range g[v] {
			dfs(to)
		}

		result = append(result, v)
	}

	dfs(start)

	reverse := make([]int, 0, len(result))
	for i := len(result) - 1; i >= 0; i-- {
		reverse = append(reverse, result[i])
	}

	return reverse
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

			if !isUpdateCorrect(update, rulesMap) {
				newUpdate := fixUpdate(update, rulesMap)
				if newUpdate == nil {
					log.Fatalf("couldnt fix update %v", update)
				} else {
					log.Printf("old: %v", update)
					log.Printf("new: %v\n", newUpdate)
					answer += newUpdate[len(newUpdate)/2]
				}
			}
		}
	}

	log.Printf("Answer: %d", answer)
}
