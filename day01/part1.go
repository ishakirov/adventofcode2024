package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	list1 := make([]int, 0)
	list2 := make([]int, 0)

	scanner := bufio.NewScanner(input)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		first, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		second, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}

		list1 = append(list1, first)
		list2 = append(list2, second)

		i++
	}

	sort.Ints(list1)
	sort.Ints(list2)

	distanceSum := 0
	for i := 0; i < len(list1); i++ {
		distanceSum += abs(list1[i] - list2[i])
	}

	log.Printf("Distance sum: %d", distanceSum)
}
