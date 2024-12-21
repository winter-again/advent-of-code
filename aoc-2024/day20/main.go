package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	part := flag.Int("part", 1, "problem part")
	flag.Parse()
	fmt.Println("Solving part", *part)

	if *part == 1 {
		ans := solvePart1("./input_smpl.txt", true)
		fmt.Println("Answer (sample):", ans)

		ans = solvePart1("./input.txt", false)
		fmt.Println("Answer:", ans)
	} else {
		ans := solvePart2("./input_smpl.txt", true)
		fmt.Println("Answer (sample):", ans)

		ans = solvePart2("./input.txt", false)
		fmt.Println("Answer:", ans)
	}
}

func solvePart1(input string, sample bool) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := parseGrid(scanner)

	var start pos
	var end pos
	for i, row := range grid {
		for j := range row {
			if grid[i][j] == "S" {
				start = pos{i: i, j: j}
			} else if grid[i][j] == "E" {
				end = pos{i: i, j: j}
			}
		}
	}

	locs := findFastest(grid, start, end)

	var thresh int
	if sample {
		// NOTE: not explicitly given as sample cond and result; just summed up the
		// sample results given
		thresh = 2
	} else {
		thresh = 100
	}

	ct := 0
	for s := 0; s < len(locs); s++ {
		savedCheat := tryCheat(grid, s, locs, thresh)
		ct += savedCheat
	}
	return ct
}

type pos struct {
	i int
	j int
}

var dirs = []pos{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func tryCheat(grid [][]string, n int, locs []pos, thresh int) int {
	res := 0
	visited := make(map[pos]bool)
	for i := 0; i < n; i++ {
		visited[locs[i]] = true
	}
	cheatStart := locs[n]

	for _, dir := range dirs {
		wall := pos{i: cheatStart.i + dir.i, j: cheatStart.j + dir.j}
		cheatEnd := pos{i: cheatStart.i + 2*dir.i, j: cheatStart.j + 2*dir.j}
		if (wall.i >= 0 && wall.i < len(grid) && wall.j >= 0 && wall.j < len(grid[0])) && (cheatEnd.i >= 0 && cheatEnd.i < len(grid) && cheatEnd.j >= 0 && cheatEnd.j < len(grid[0])) && (grid[wall.i][wall.j] == "#" && grid[cheatEnd.i][cheatEnd.j] != "#" && !visited[cheatEnd]) {
			cheat := slices.Index(locs, cheatEnd)
			diff := cheat - n - 2
			if diff >= thresh {
				res++
			}
			visited[cheatEnd] = true
		}
	}
	return res
}

func findFastest(grid [][]string, start pos, end pos) []pos {
	visited := make(map[pos]bool)
	var locs []pos
	cur := start
	for {
		if cur == end {
			locs = append(locs, cur)
			break
		}
		for _, dir := range dirs {
			next := pos{i: cur.i + dir.i, j: cur.j + dir.j}
			if grid[next.i][next.j] != "#" && !visited[pos{i: next.i, j: next.j}] {
				visited[cur] = true
				locs = append(locs, cur)
				cur = next
			}
		}
	}
	return locs
}

func parseGrid(scanner *bufio.Scanner) [][]string {
	var track [][]string

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		track = append(track, line)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return track
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func solvePart2(input string, sample bool) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := parseGrid(scanner)

	var start pos
	var end pos
	for i, row := range grid {
		for j := range row {
			if grid[i][j] == "S" {
				start = pos{i: i, j: j}
			} else if grid[i][j] == "E" {
				end = pos{i: i, j: j}
			}
		}
	}

	locs := findFastest(grid, start, end)

	var thresh int
	if sample {
		// NOTE: not explicitly given as sample cond and result; just summed up the
		// sample results given
		thresh = 50
	} else {
		thresh = 100
	}

	ct := 0
	for s := 0; s < len(locs); s++ {
		ct += exploreCheat(grid, s, locs, thresh)
	}
	return ct
}

func exploreCheat(grid [][]string, s int, locs []pos, thresh int) int {
	ct := 0
	visitedCheat := make(map[pos]bool)
	for r := 2; r <= 20; r++ {
		for di := 0; di <= r; di++ {
			dj := r - di
			cur := locs[s]
			ends := []pos{{i: cur.i + di, j: cur.j + dj}, {i: cur.i + di, j: cur.j - dj}, {i: cur.i - di, j: cur.j + dj}, {i: cur.i - di, j: cur.j - dj}}
			for _, end := range ends {
				if (end.i >= 0 && end.i < len(grid) && end.j >= 0 && end.j < len(grid[0])) && grid[end.i][end.j] != "#" && !visitedCheat[end] {
					visitedCheat[end] = true
					cheatEnd := slices.Index(locs, end)
					diff := cheatEnd - s - r
					if diff >= thresh {
						ct++
					}
				}
			}
		}
	}
	return ct
}
