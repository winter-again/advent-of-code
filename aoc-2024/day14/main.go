package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	part := flag.Int("part", 1, "problem part")
	flag.Parse()
	fmt.Println("Solving part", *part)

	if *part == 1 {
		ans := solvePart1("./input_smpl.txt", 7, 11)
		fmt.Println("Answer (sample):", ans)

		ans = solvePart1("./input.txt", 103, 101)
		fmt.Println("Answer:", ans)
	} else {
		ans := solvePart2("./input_smpl.txt", 7, 11)
		fmt.Println("Answer (sample):", ans)

		ans = solvePart2("./input.txt", 103, 101)
		fmt.Println("Answer:", ans)
	}
}

func solvePart1(input string, nRows int, nCols int) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	robots := parseRobots(scanner)

	for t := 0; t < 100; t++ {
		for _, robot := range robots {
			robot.move(nRows, nCols)
		}
	}

	xMid := (nCols - 1) / 2
	yMid := (nRows - 1) / 2
	quads := make([]int, 4)
	for _, r := range robots {
		if r.pos.x < xMid && r.pos.y < yMid {
			quads[0]++
		} else if r.pos.x > xMid && r.pos.y < yMid {
			quads[1]++
		} else if r.pos.x < xMid && r.pos.y > yMid {
			quads[2]++
		} else if r.pos.x > xMid && r.pos.y > yMid {
			quads[3]++
		}
	}

	safetyFactor := quads[0] * quads[1] * quads[2] * quads[3]
	return safetyFactor
}

func printGrid[T int | string](grid [][]T) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func (r *robot) move(nRows int, nCols int) {
	maxRow := nRows - 1
	maxCol := nCols - 1
	nx := r.pos.x + r.vel.x
	ny := r.pos.y + r.vel.y

	if nx < 0 {
		toEdge := r.pos.x
		diff := r.vel.x + (toEdge + 1) // r.vel.x < 0; +1 to bring onto other side
		nx = maxCol + diff
	} else if nx > maxCol {
		toEdge := maxCol - r.pos.x
		diff := r.vel.x - (toEdge + 1) // r.vel.x > 0
		nx = diff
	}

	if ny < 0 {
		toEdge := r.pos.y
		diff := r.vel.y + (toEdge + 1) // r.ve.ly < 0
		ny = maxRow + diff
	} else if ny > maxRow {
		toEdge := maxRow - r.pos.y
		diff := r.vel.y - (toEdge + 1) // r.vel.y > 0
		ny = diff
	}

	r.pos.x = nx
	r.pos.y = ny
}

type robot struct {
	pos *coord
	vel coord
}

type coord struct {
	x int
	y int
}

func (r *robot) String() string {
	return fmt.Sprintf("robot{pos:%v, v:%v}", r.pos, r.vel)
}

func (c *coord) String() string {
	return fmt.Sprintf("pos{x:%d, y:%d}", c.x, c.y)
}

func parseRobots(s *bufio.Scanner) []robot {
	var robots []robot
	for s.Scan() {
		fields := strings.FieldsFunc(s.Text(), robotFields)
		px, _ := strconv.Atoi(fields[0])
		py, _ := strconv.Atoi(fields[1])
		vx, _ := strconv.Atoi(fields[2])
		vy, _ := strconv.Atoi(fields[3])

		robot := robot{
			pos: &coord{px, py},
			vel: coord{vx, vy},
		}
		robots = append(robots, robot)
	}

	if err := s.Err(); err != nil {
		log.Println(err)
	}
	return robots
}

func robotFields(c rune) bool {
	return (!unicode.IsNumber(c) && c != '-')
}

func solvePart2(input string, nRows int, nCols int) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	robots := parseRobots(scanner)

	xMid := (nCols - 1) / 2
	yMid := (nRows - 1) / 2

	quads := make([]int, 4)
	bestFactor := 9999999999
	bestIter := 0
	for t := 1; t <= 100_000; t++ {
		for _, robot := range robots {
			robot.move(nRows, nCols)
		}

		for _, r := range robots {
			if r.pos.x < xMid && r.pos.y < yMid {
				quads[0]++
			} else if r.pos.x > xMid && r.pos.y < yMid {
				quads[1]++
			} else if r.pos.x < xMid && r.pos.y > yMid {
				quads[2]++
			} else if r.pos.x > xMid && r.pos.y > yMid {
				quads[3]++
			}
		}

		safetyFactor := quads[0] * quads[1] * quads[2] * quads[3]
		if safetyFactor < bestFactor {
			grid := make([][]string, nRows)
			for i := range grid {
				grid[i] = make([]string, nCols)
			}

			for i, row := range grid {
				for j := range row {
					grid[i][j] = "."
				}
			}
			for _, r := range robots {
				grid[r.pos.y][r.pos.x] = "#"
			}
			// fmt.Println("t=", t, "--------------------------------")
			// printGrid(grid)

			bestFactor = safetyFactor
			bestIter = t
		}

		for n := range quads {
			quads[n] = 0
		}
	}
	return bestIter
}
