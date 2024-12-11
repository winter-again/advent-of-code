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
	mp, positions := parseMap(scanner)

	score := 0
	for _, head := range positions {
		ends := make(map[pos]bool)
		walkTrail(mp, head, ends)
		score += len(ends)
	}
	return score
}

var dirs = []pos{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func walkTrail(mp [][]string, cur pos, ends map[pos]bool) {
	if mp[cur.i][cur.j] == "9" {
		ends[pos{cur.i, cur.j}] = true
		return
	}

	// NOTE: shouldn't need to drop last used dir since the to-from == 1 will exclude it
	for _, dir := range dirs {
		next := pos{cur.i + dir.i, cur.j + dir.j}
		if (next.i >= 0 && next.i < len(mp)) && (next.j >= 0 && next.j < len(mp[0])) {
			from, _ := strconv.Atoi(mp[cur.i][cur.j])
			to, _ := strconv.Atoi(mp[next.i][next.j])
			if to-from == 1 {
				walkTrail(mp, next, ends)
			}
		}
	}
}

type pos struct {
	i int
	j int
}

func parseMap(s *bufio.Scanner) ([][]string, []pos) {
	var mp [][]string
	for s.Scan() {
		mp = append(mp, strings.Split(s.Text(), ""))
	}

	if err := s.Err(); err != nil {
		log.Println(err)
	}

	var positions []pos
	for i, row := range mp {
		for j, val := range row {
			if val == "0" {
				positions = append(positions, pos{i, j})
			}
		}
	}
	return mp, positions
}

func printMap(mp [][]string) {
	for _, row := range mp {
		fmt.Println(row)
	}
}

func solvePart2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	mp, positions := parseMap(scanner)

	ratings := 0
	for _, head := range positions {
		trails := 0
		walkTrailPart2(mp, head, &trails)
		ratings += trails
	}
	return ratings
}
func walkTrailPart2(mp [][]string, cur pos, trails *int) {
	if mp[cur.i][cur.j] == "9" {
		*trails++
		return
	}

	// NOTE: again, shouldn't need to drop last used dir since the to-from == 1 will exclude it
	for _, dir := range dirs {
		next := pos{cur.i + dir.i, cur.j + dir.j}
		if (next.i >= 0 && next.i < len(mp)) && (next.j >= 0 && next.j < len(mp[0])) {
			from, _ := strconv.Atoi(mp[cur.i][cur.j])
			to, _ := strconv.Atoi(mp[next.i][next.j])
			if to-from == 1 {
				walkTrailPart2(mp, next, trails)
			}
		}
	}
	return
}
