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
		ans := solvePart1("./input_smpl.txt")
		fmt.Println("Answer (sample):", ans)

		ans = solvePart1("./input.txt")
		fmt.Println("Answer:", ans)

		ans = solvePart1NoBlacklist("./input_smpl.txt")
		fmt.Println("Answer (sample):", ans)

		ans = solvePart1NoBlacklist("./input.txt")
		fmt.Println("Answer:", ans)
	} else {
		ans := solvePart2("./input_smpl.txt")
		fmt.Println("Answer (sample):", ans)

		ans = solvePart2("./input.txt")
		fmt.Println("Answer:", ans)
	}
}

type coord struct {
	i int
	j int
}

func solvePart1(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ws := parseWordSearch(scanner)
	nRows := len(ws)
	nCols := len(ws[0])
	count := make(map[string]int, 8)
	blacklist := make(map[string][]coord)
	wMap := map[string][]string{
		"X": {"M", "A", "S"},
		"S": {"A", "M", "X"},
	}
	for i, row := range ws {
		for j, col := range row {
			if col == "X" || col == "S" {
				letters := wMap[col]
				if i > 2 {
					// up
					if !slices.Contains(blacklist["up"], coord{i, j}) {
						// log.Printf("Exploring UP from (%d, %d)", i, j)
						for p := 0; p < 3; p++ {
							if ws[i-1-p][j] != letters[p] {
								break
							}
							if p == 2 {
								count["up"]++
							}
						}
					}
				}
				if i > 2 && j < nCols-3 {
					// up-right
					if !slices.Contains(blacklist["up-right"], coord{i, j}) {
						// log.Printf("Exploring UP-RIGHT from (%d, %d)", i, j)
						for p := 0; p < 3; p++ {
							if ws[i-1-p][j+1+p] != letters[p] {
								break
							}
							if p == 2 {
								// log.Printf("UP-RIGHT from (%d, %d)", i, j)
								count["up-right"]++
							}
						}
					}
				}
				if j < nCols-3 {
					// right
					// log.Printf("Exploring RIGHT from (%d, %d)", i, j)
					for p := 0; p < 3; p++ {
						if ws[i][j+1+p] != letters[p] {
							break
						}
						if p == 2 {
							// log.Printf("RIGHT from (%d, %d)", i, j)
							blacklist["left"] = append(blacklist["left"], coord{i, j + 1 + p})
							count["right"]++
						}
					}
				}
				// note: only have to blacklist for right (left), down-right (up-left),
				// down (up), down-left (up-right)?
				if i < nRows-3 && j < nCols-3 {
					// down-right
					for p := 0; p < 3; p++ {
						if ws[i+1+p][j+1+p] != letters[p] {
							break
						}
						if p == 2 {
							// log.Printf("DOWN-RIGHT from (%d, %d)", i, j)
							blacklist["up-left"] = append(blacklist["up-left"], coord{i + 1 + p, j + 1 + p})
							count["down-right"]++
						}
					}
				}
				if i < nRows-3 {
					// down
					for p := 0; p < 3; p++ {
						if ws[i+1+p][j] != letters[p] {
							break
						}
						if p == 2 {
							// log.Printf("DOWN from (%d, %d)", i, j)
							blacklist["up"] = append(blacklist["up"], coord{i + 1 + p, j})
							count["down"]++
						}
					}
				}
				if i < nRows-3 && j > 3 {
					// down-left
					for p := 0; p < 3; p++ {
						if ws[i+1+p][j-1-p] != letters[p] {
							break
						}
						if p == 2 {
							// log.Printf("DOWN-LEFT from (%d, %d)", i, j)
							blacklist["up-right"] = append(blacklist["up-right"], coord{i + 1 + p, j - 1 - p})
							count["down-left"]++
						}
					}
				}
				if j > 2 {
					// left
					if !slices.Contains(blacklist["left"], coord{i, j}) {
						// log.Printf("Exploring LEFT from (%d, %d)", i, j)
						for p := 0; p < 3; p++ {
							if ws[i][j-1-p] != letters[p] {
								break
							}
							if p == 2 {
								count["left"]++
							}
						}
					}
				}
				if i > 3 && j > 3 {
					// up-left
					if !slices.Contains(blacklist["up-left"], coord{i, j}) {
						// log.Printf("Exploring UP-LEFT from (%d, %d)", i, j)
						for p := 0; p < 3; p++ {
							if ws[i-1-p][j-1-p] != letters[p] {
								break
							}
							if p == 2 {
								count["up-left"]++
							}
						}
					}
				}
			}
		}
	}
	return sum(count)
}

func solvePart1NoBlacklist(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ws := parseWordSearch(scanner)
	nRows := len(ws)
	nCols := len(ws[0])
	dirs := []coord{
		{0, -1},  // left
		{0, 1},   // right
		{-1, 0},  // up
		{1, 0},   // down
		{-1, -1}, // up-left
		{1, 1},   // down-right
		{-1, 1},  // up-left
		{1, -1},  // down-left
	}
	count := 0
	for i, row := range ws {
		for j, col := range row {
			if col == "X" {
				for _, dir := range dirs {
					if (i+3*dir.i >= 0 && i+3*dir.i < nRows) && (j+3*dir.j >= 0 && j+3*dir.j < nCols) {
						if ws[i+dir.i][j+dir.j] == "M" && ws[i+2*dir.i][j+2*dir.j] == "A" && ws[i+3*dir.i][j+3*dir.j] == "S" {
							count += 1
						}
					}
				}
			}
		}
	}
	return count
}

func parseWordSearch(s *bufio.Scanner) [][]string {
	var ws [][]string
	for s.Scan() {
		ws = append(ws, strings.Split(s.Text(), ""))
	}
	if err := s.Err(); err != nil {
		log.Println(err)
	}
	return ws
}

func sum(s map[string]int) int {
	sum := 0
	for _, c := range s {
		sum += c
	}
	return sum
}

func solvePart2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ws := parseWordSearch(scanner)
	nRows := len(ws)
	nCols := len(ws[0])
	count := 0
	for i := 1; i < nRows-1; i++ {
		for j := 1; j < nCols-1; j++ {
			if ws[i][j] == "A" {
				if isXmas(ws, i, j) {
					count++
				}
			}
		}
	}
	return count
}

func isXmas(ws [][]string, i int, j int) bool {
	var back bool
	var fwd bool
	if (ws[i-1][j-1] == "M" && ws[i+1][j+1] == "S") || (ws[i-1][j-1] == "S" && ws[i+1][j+1] == "M") {
		back = true
	}
	if (ws[i-1][j+1] == "M" && ws[i+1][j-1] == "S") || (ws[i-1][j+1] == "S" && ws[i+1][j-1] == "M") {
		fwd = true
	}
	if back && fwd {
		return true
	}
	return false
}
