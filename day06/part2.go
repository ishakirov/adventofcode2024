package main

import (
	"bufio"
	"log"
	"os"
)

var (
	directions = [][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	directionLabels = map[byte]int{
		'^': 0,
		'>': 1,
		'v': 2,
		'<': 3,
	}
)

func inBounds(g [][]byte, i, j int) bool {
	return i >= 0 && i < len(g) && j >= 0 && j < len(g[i])
}

func detectLoop(g [][]byte) bool {
	guardI, guardJ := -1, -1
	guardDirectionIndex := -1

	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			if directionIndex, ok := directionLabels[g[i][j]]; ok {
				guardI, guardJ = i, j
				guardDirectionIndex = directionIndex
				break
			}
		}
	}

	visited := make(map[[2]int]int)
	visited[[2]int{guardI, guardJ}] = guardDirectionIndex
	for {
		newI, newJ := guardI+directions[guardDirectionIndex][0], guardJ+directions[guardDirectionIndex][1]
		if !inBounds(g, newI, newJ) {
			break
		}
		if g[newI][newJ] == '#' {
			guardDirectionIndex = (guardDirectionIndex + 1) % len(directions)
			continue
		}

		if directionIndex, ok := visited[[2]int{newI, newJ}]; ok {
			if directionIndex == guardDirectionIndex {
				return true
			}
		} else {
			visited[[2]int{newI, newJ}] = guardDirectionIndex
		}

		guardI, guardJ = newI, newJ
	}

	return false
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	g := make([][]byte, 0)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		g = append(g, []byte(line))
	}

	answer := 0
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			if g[i][j] == '#' {
				continue
			}
			if _, ok := directionLabels[g[i][j]]; ok {
				continue
			}
			g[i][j] = '#'
			if detectLoop(g) {
				answer++
			}
			g[i][j] = '.'
		}
	}

	log.Println(answer)
}
