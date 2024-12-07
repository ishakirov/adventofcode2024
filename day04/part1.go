package main

import (
	"bufio"
	"log"
	"os"
)

func findXmas(letters []string) int {
	searchWord := "XMAS"
	directions := [][2]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
		{1, 1},
		{-1, -1},
		{1, -1},
		{-1, 1},
	}

	count := 0
	for i := 0; i < len(letters); i++ {
		for j := 0; j < len(letters[i]); j++ {
			if letters[i][j] == searchWord[0] {
				for _, dir := range directions {
					searchWordIndex := 0
					searchI := i
					searchJ := j
					for searchI >= 0 && searchI < len(letters) && searchJ >= 0 && searchJ < len(letters[searchI]) && searchWordIndex < len(searchWord) {
						if letters[searchI][searchJ] == searchWord[searchWordIndex] {
							searchWordIndex++
						} else {
							break
						}

						searchI += dir[0]
						searchJ += dir[1]
					}

					if searchWordIndex == len(searchWord) {
						count++
					}
				}
			}
		}
	}

	return count
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	letters := make([]string, 0)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		letters = append(letters, line)
	}

	log.Println("XMAS count:", findXmas(letters))
}
