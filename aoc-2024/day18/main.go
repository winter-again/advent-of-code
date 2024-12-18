package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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
	bts := parseBytes(scanner)

	var nRows int
	var nCols int
	var n int
	if sample {
		nRows = 7
		nCols = 7
		n = 12
	} else {
		nRows = 71
		nCols = 71
		n = 1024
	}
	grid := mapBytes(nRows, nCols, bts, n)
	steps := walkGrid(grid)
	return steps
}

type pos struct {
	i int
	j int
}

var dirs = []pos{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

func walkGrid(grid [][]string) int {
	queue := []pos{{0, 0}}
	seen := make(map[pos]bool)
	seen[pos{0, 0}] = true
	parents := make(map[pos]pos)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.i == len(grid[0])-1 && cur.j == len(grid)-1 {
			break
		}
		for _, dir := range dirs {
			next := pos{i: cur.i + dir.i, j: cur.j + dir.j}
			if seen[next] || next.i < 0 || next.i >= len(grid[0]) || next.j < 0 || next.j >= len(grid) || grid[next.i][next.j] == "#" {
				continue
			}
			seen[next] = true
			parents[next] = cur
			queue = append(queue, next)
		}
	}

	steps := 0
	stack := []pos{{len(grid[0]) - 1, len(grid) - 1}}
	origin := pos{i: 0, j: 0}
	for len(stack) > 0 {
		sq := stack[len(stack)-1]
		if sq == origin {
			break
		}
		stack = stack[:len(stack)-1]
		stack = append(stack, parents[sq])
		steps++
	}
	return steps
}

func mapBytes(nRows int, nCols int, bts []pos, n int) [][]string {
	grid := make([][]string, nRows)
	for i := range grid {
		grid[i] = make([]string, nCols)
	}

	for i, row := range grid {
		for j := range row {
			grid[i][j] = "."
		}
	}

	s := 0
	for s < n {
		w := bts[s]
		grid[w.i][w.j] = "#"
		s++
	}
	return grid
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func parseBytes(scanner *bufio.Scanner) []pos {
	var bts []pos
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		bts = append(bts, pos{i: y, j: x})
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return bts
}

func solvePart2(input string, sample bool) string {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	bts := parseBytes(scanner)

	var nRows int
	var nCols int
	if sample {
		nRows = 7
		nCols = 7
	} else {
		nRows = 71
		nCols = 71
	}

	badByte := 0
	for n := 1; n <= len(bts); n++ {
		grid := mapBytes(nRows, nCols, bts, n)

		canExit := findNoExit(grid)
		if !canExit {
			badByte = n
			break
		}
	}

	flipped := []int{bts[badByte-1].j, bts[badByte-1].i}
	flippedStr := make([]string, len(flipped))
	for i, s := range flipped {
		flippedStr[i] = strconv.Itoa(s)
	}
	ans := strings.Join(flippedStr, ",")
	return ans
}

func findNoExit(grid [][]string) bool {
	queue := []pos{{0, 0}}
	seen := make(map[pos]bool)
	seen[pos{0, 0}] = true
	parents := make(map[pos]pos)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.i == len(grid[0])-1 && cur.j == len(grid)-1 {
			break
		}
		for _, dir := range dirs {
			next := pos{i: cur.i + dir.i, j: cur.j + dir.j}
			if seen[next] || next.i < 0 || next.i >= len(grid[0]) || next.j < 0 || next.j >= len(grid) || grid[next.i][next.j] == "#" {
				continue
			}
			seen[next] = true
			parents[next] = cur
			queue = append(queue, next)
		}
	}

	dest := pos{len(grid[0]) - 1, len(grid) - 1}
	_, ok := parents[dest]
	if !ok {
		return false
	}
	return true
}
