package main

import (
	"bufio"
	"log"
	"os"
)

var directions = [][2]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

func inBounds(g [][]byte, i, j int) bool {
	return i >= 0 && i < len(g) && j >= 0 && j < len(g[i])
}

func computeScore(g [][]byte, i, j int) int {
	queue := make([][2]int, 0)
	visited := make(map[[2]int]struct{})

	queue = append(queue, [2]int{i, j})
	visited[[2]int{i, j}] = struct{}{}

	fullTrails := 0

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		if g[v[0]][v[1]] == '9' {
			fullTrails++
		}

		for _, dir := range directions {
			x, y := v[0]+dir[0], v[1]+dir[1]
			if !inBounds(g, x, y) {
				continue
			}

			if _, ok := visited[[2]int{x, y}]; !ok && g[x][y] == g[v[0]][v[1]]+1 {
				queue = append(queue, [2]int{x, y})
				visited[[2]int{x, y}] = struct{}{}
			}
		}
	}

	return fullTrails
}

func computeRating(g [][]byte, i, j int) int {
	memo := make(map[[2]int]int)

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if v, ok := memo[[2]int{i, j}]; ok {
			return v
		}

		if g[i][j] == '9' {
			return 1
		}

		rating := 0
		for _, dir := range directions {
			x, y := i+dir[0], j+dir[1]
			if !inBounds(g, x, y) {
				continue
			}

			if g[x][y] == g[i][j]+1 {
				rating += dfs(x, y)
			}
		}

		memo[[2]int{i, j}] = rating
		return rating
	}

	return dfs(i, j)
}

func part1(g [][]byte) int {
	answer := 0
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			if g[i][j] == '0' {
				score := computeScore(g, i, j)
				// log.Println(score)
				answer += score
			}
		}
	}
	return answer
}

func part2(g [][]byte) int {
	answer := 0
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			if g[i][j] == '0' {
				rating := computeRating(g, i, j)
				log.Println(rating)
				answer += rating
			}
		}
	}
	return answer
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

	log.Println(part1(g))
	log.Print(part2(g))
}
