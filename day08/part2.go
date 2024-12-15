package main

import (
	"bufio"
	"log"
	"os"
)

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

	antennas := make(map[byte][][2]int)
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			if g[i][j] != '.' {
				if _, ok := antennas[g[i][j]]; !ok {
					antennas[g[i][j]] = make([][2]int, 0)
				}
				antennas[g[i][j]] = append(antennas[g[i][j]], [2]int{i, j})
			}
		}
	}

	antinodes := make(map[[2]int]struct{})
	for _, points := range antennas {
		for i := 0; i < len(points)-1; i++ {
			for j := i + 1; j < len(points); j++ {
				antinode1 := [2]int{
					points[i][0],
					points[i][1],
				}
				for antinode1[0] >= 0 && antinode1[0] < len(g) && antinode1[1] >= 0 && antinode1[1] < len(g[antinode1[0]]) {
					antinodes[antinode1] = struct{}{}

					antinode1[0] += points[i][0] - points[j][0]
					antinode1[1] += points[i][1] - points[j][1]
				}

				antinode2 := [2]int{
					points[j][0],
					points[j][1],
				}
				for antinode2[0] >= 0 && antinode2[0] < len(g) && antinode2[1] >= 0 && antinode2[1] < len(g[antinode2[0]]) {
					antinodes[antinode2] = struct{}{}

					antinode2[0] += points[j][0] - points[i][0]
					antinode2[1] += points[j][1] - points[i][1]
				}
			}
		}
	}

	log.Println(len(antinodes))
}
