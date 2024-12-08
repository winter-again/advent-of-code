package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
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

type pos struct {
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
	mp := parseMap(scanner)

	antennas := make(map[string][]pos)
	re := regexp.MustCompile(`^[A-Za-z0-9]$`)
	for i, row := range mp {
		for j, col := range row {
			if re.MatchString(col) {
				antennas[col] = append(antennas[col], pos{i, j})
			}
		}
	}

	antinodes := make(map[pos]bool)
	for _, locs := range antennas {
		for i, antA := range locs[:len(locs)-1] {
			for _, antB := range locs[i+1:] {
				// di, dj is how to get to antB from antA
				di := antB.i - antA.i
				dj := antB.j - antA.j

				antinodeA := pos{antA.i - di, antA.j - dj}
				if (antinodeA.i >= 0 && antinodeA.i < len(mp)) && (antinodeA.j >= 0 && antinodeA.j < len(mp[0])) {
					if !antinodes[antinodeA] {
						antinodes[antinodeA] = true
					}
				}

				antinodeB := pos{antB.i + di, antB.j + dj}
				if (antinodeB.i >= 0 && antinodeB.i < len(mp)) && (antinodeB.j >= 0 && antinodeB.j < len(mp[0])) {
					if !antinodes[antinodeB] {
						antinodes[antinodeB] = true
					}
				}
			}
		}
	}
	return len(antinodes)
}

func parseMap(s *bufio.Scanner) [][]string {
	var mp [][]string
	for s.Scan() {
		line := strings.Split(s.Text(), "")
		mp = append(mp, line)
	}

	if err := s.Err(); err != nil {
		log.Println(err)
	}
	return mp
}

func printMap(mp [][]string) {
	for _, line := range mp {
		fmt.Println(line)
	}
}

func solvePart2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	mp := parseMap(scanner)

	antennas := make(map[string][]pos)
	re := regexp.MustCompile(`^[A-Za-z0-9]$`)
	for i, row := range mp {
		for j, col := range row {
			if re.MatchString(col) {
				antennas[col] = append(antennas[col], pos{i, j})
			}
		}
	}

	antinodes := make(map[pos]bool)
	for _, locs := range antennas {
		for i, antA := range locs[:len(locs)-1] {
			for _, antB := range locs[i+1:] {
				if !antinodes[antA] {
					antinodes[antA] = true
				}

				if !antinodes[antB] {
					antinodes[antB] = true
				}

				di := antB.i - antA.i
				dj := antB.j - antA.j

				curr := pos{antA.i - di, antA.j - dj}
				for {
					if (curr.i >= 0 && curr.i < len(mp)) && (curr.j >= 0 && curr.j < len(mp[0])) {
						if !antinodes[curr] {
							antinodes[curr] = true
							mp[curr.i][curr.j] = "#"
						}
						curr = pos{curr.i - di, curr.j - dj}
					} else {
						break
					}
				}

				curr = pos{antB.i + di, antB.j + dj}
				for {
					if (curr.i >= 0 && curr.i < len(mp)) && (curr.j >= 0 && curr.j < len(mp[0])) {
						if !antinodes[curr] {
							antinodes[curr] = true
							mp[curr.i][curr.j] = "#"
						}
						curr = pos{curr.i + di, curr.j + dj}
					} else {
						break
					}
				}
			}
		}
	}
	return len(antinodes)
}
