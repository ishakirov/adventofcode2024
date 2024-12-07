package main

import (
	"bufio"
	"log"
	"os"
)

func findXDashMas(letters []string) int {
	count := 0

	for i := 1; i < len(letters)-1; i++ {
		for j := 1; j < len(letters[i])-1; j++ {
			if letters[i][j] == 'A' {
				if (letters[i+1][j+1] == 'M' && letters[i-1][j-1] == 'S' || letters[i+1][j+1] == 'S' && letters[i-1][j-1] == 'M') &&
					(letters[i+1][j-1] == 'M' && letters[i-1][j+1] == 'S' || letters[i+1][j-1] == 'S' && letters[i-1][j+1] == 'M') {
					count++
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

	log.Println("X-MAS count:", findXDashMas(letters))
}
