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

		ans = solvePart1("./input_smpl_2.txt")
		fmt.Println("Answer:", ans)

		ans = solvePart1("./input.txt")
		fmt.Println("Answer:", ans)
	} else {
		ans := solvePart2("./input_smpl.txt")
		fmt.Println("Answer (sample):", ans)

		ans = solvePart2("./input_smpl_2.txt")
		fmt.Println("Answer (sample 2):", ans)

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
	garden := parseGarden(scanner)

	regions := findRegions(garden)

	price := 0
	for _, region := range regions {
		perimCost := 0
		plant := garden[region[0].i][region[0].j]
		for _, p := range region {
			for _, dir := range dirs {
				n := pos{p.i + dir.i, p.j + dir.j}
				if n.i < 0 || n.i >= len(garden) || n.j < 0 || n.j >= len(garden[0]) {
					perimCost += 1
					continue
				}
				if garden[n.i][n.j] != plant {
					perimCost += 1
				} else {
					continue
				}
			}
		}
		price += len(region) * perimCost
	}
	return price
}

type pos struct {
	i int
	j int
}

var dirs = []pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// flood fill algo with queue
// recursive gave stack overflow
func findRegions(garden [][]string) [][]pos {
	var regions [][]pos
	var seen []pos
	for i, row := range garden {
		for j := range row {
			if slices.Contains(seen, pos{i, j}) {
				continue
			}
			seen = append(seen, pos{i, j})

			var region []pos
			q := []pos{{i, j}}
			for {
				if len(q) == 0 {
					break
				}
				// pop from queue
				cur := q[0]
				q = q[1:]
				for _, dir := range dirs {
					n := pos{cur.i + dir.i, cur.j + dir.j}
					if n.i < 0 || n.i >= len(garden) || n.j < 0 || n.j >= len(garden[0]) {
						continue
					}
					if garden[n.i][n.j] != garden[cur.i][cur.j] {
						if !slices.Contains(region, cur) {
							region = append(region, cur)
						}
						continue
					}
					if slices.Contains(region, n) {
						continue
					}
					region = append(region, n)
					seen = append(seen, n)
					q = append(q, n)
				}
			}
			regions = append(regions, region)
		}
	}
	return regions
}

func parseGarden(s *bufio.Scanner) [][]string {
	var garden [][]string
	for s.Scan() {
		garden = append(garden, strings.Split(s.Text(), ""))
	}

	if err := s.Err(); err != nil {
		log.Println(err)
	}
	return garden
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
	garden := parseGarden(scanner)
	regions := findRegions(garden)

	price := 0
	for _, region := range regions {
		sides := 0
		for _, p := range region {
			vertices := numVertices(p, garden)
			sides += vertices
		}
		price += len(region) * sides
	}
	return price
}

func numVertices(p pos, garden [][]string) int {
	vertices := 0

	up := pos{p.i + dirs[0].i, p.j + dirs[0].j}
	down := pos{p.i + dirs[1].i, p.j + dirs[1].j}
	left := pos{p.i + dirs[2].i, p.j + dirs[2].j}
	right := pos{p.i + dirs[3].i, p.j + dirs[3].j}

	upLeft := pos{p.i - 1, p.j - 1}
	upRight := pos{p.i - 1, p.j + 1}
	downLeft := pos{p.i + 1, p.j - 1}
	downRight := pos{p.i + 1, p.j + 1}

	if invalid(p, up, garden) {
		if invalid(p, left, garden) {
			vertices++
		}
		if invalid(p, right, garden) {
			vertices++
		}
	}
	if invalid(p, down, garden) {
		if invalid(p, left, garden) {
			vertices++
		}
		if invalid(p, right, garden) {
			vertices++
		}
	}

	if !invalid(p, up, garden) && !invalid(p, left, garden) && invalid(p, upLeft, garden) {
		vertices++
	}
	if !invalid(p, up, garden) && !invalid(p, right, garden) && invalid(p, upRight, garden) {
		vertices++
	}
	if !invalid(p, down, garden) && !invalid(p, left, garden) && invalid(p, downLeft, garden) {
		vertices++
	}
	if !invalid(p, down, garden) && !invalid(p, right, garden) && invalid(p, downRight, garden) {
		vertices++
	}
	return vertices
}

func invalid(p pos, dir pos, garden [][]string) bool {
	if dir.i < 0 || dir.i >= len(garden) || dir.j < 0 || dir.j >= len(garden[0]) {
		return true
	}
	if garden[dir.i][dir.j] != garden[p.i][p.j] {
		return true
	}
	return false
}
