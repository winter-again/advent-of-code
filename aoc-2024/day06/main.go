package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	part := flag.Int("part", 1, "problem part")
	flag.Parse()
	fmt.Println("Solving part", *part)

	if *part == 1 {
		ans := solvePart1("./input_smpl.txt")
		fmt.Println("Answer (sample):", ans)

		ans = solvePart1("./input.txt")
		fmt.Println("Answer:", ans)
	} else {
		ans := solvePart2("./input_smpl.txt")
		fmt.Println("Answer (sample):", ans)

		ans = solvePart2("./input.txt")
		fmt.Println("Answer:", ans)
	}
}

func solvePart1(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	mp := parseLabMap(scanner)

	var initI int
	var initJ int
	var found bool
	for i := 0; i < len(mp); i++ {
		if found {
			break
		}
		for j := 0; j < len(mp[i]); j++ {
			if mp[i][j] == "^" {
				found = true
				initI = i
				initJ = j
				break
			}
		}
	}
	path, _ := findPath(mp, initI, initJ)
	return len(path)
}

func parseLabMap(s *bufio.Scanner) [][]string {
	var m [][]string
	for s.Scan() {
		m = append(m, strings.Split(s.Text(), ""))
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return m
}

// note: order is important, with up first then 90 deg cw turns
var dirs = [][]int{
	{-1, 0}, {0, 1}, {1, 0}, {0, -1},
}

func findPath(m [][]string, i int, j int) ([][]int, [][]string) {
	mp := copyMap(m)

	dir := 0
	var path [][]int
	for {
		ni := i + dirs[dir][0]
		nj := j + dirs[dir][1]

		if (ni < 0 || ni > len(mp)-1) || (nj < 0 || nj > len(mp[0])-1) {
			mp[i][j] = "X"
			path = append(path, []int{i, j})
			break
		}

		if mp[ni][nj] == "#" {
			dir = (dir + 1) % 4
		} else {
			if mp[i][j] != "X" {
				mp[i][j] = "X"
				path = append(path, []int{i, j})
			}
			i = ni
			j = nj
		}
	}
	return path, mp
}

func copyMap(mp [][]string) [][]string {
	mpc := make([][]string, len(mp))
	for i, row := range mp {
		mpc[i] = make([]string, len(row))
		copy(mpc[i], row)
	}
	return mpc
}

func solvePart2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	mp := parseLabMap(scanner)

	var initI int
	var initJ int
	var found bool
	for i := 0; i < len(mp); i++ {
		for j := 0; j < len(mp[i]); j++ {
			if mp[i][j] == "^" {
				found = true
				initI = i
				initJ = j
				break
			}
		}
		if found {
			break
		}
	}

	path, _ := findPath(mp, initI, initJ)

	// NOTE: max answer is 5,444
	pos := 0
	for _, p := range path[1:] {
		isCycle, _ := isCycle(mp, initI, initJ, p)
		if isCycle {
			pos++
		}
	}
	return pos
}

type direction struct {
	x int
	y int
}

func isCycle(m [][]string, i int, j int, p []int) (bool, [][]string) {
	mp := copyMap(m)
	dir := 0
	mp[p[0]][p[1]] = "O"

	visited := make(map[[3]int]bool, len(m))

	for {
		if visited[[3]int{i, j, dir}] {
			return true, mp
		}
		visited[[3]int{i, j, dir}] = true
		ni := i + dirs[dir][0]
		nj := j + dirs[dir][1]

		if (ni < 0 || ni > len(mp)-1) || (nj < 0 || nj > len(mp[0])-1) {
			return false, mp
		}

		if mp[ni][nj] == "O" || mp[ni][nj] == "#" {
			dir = (dir + 1) % 4
		} else {
			i = ni
			j = nj
		}
	}
}

func printMap(mp [][]string) {
	for _, r := range mp {
		log.Println(r)
	}
}
