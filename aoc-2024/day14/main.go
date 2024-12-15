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

		ans = solvePart1Modulo("./input_smpl.txt", 7, 11)
		fmt.Println("Answer (sample w/ modulo):", ans)

		ans = solvePart1("./input.txt", 103, 101)
		fmt.Println("Answer:", ans)

		ans = solvePart1Modulo("./input.txt", 103, 101)
		fmt.Println("Answer (w/ modulo):", ans)
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

func solvePart1Modulo(input string, nRows int, nCols int) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	robots := parseRobots(scanner)
	t := 100
	for _, robot := range robots {
		robot.pos.x = mod(robot.pos.x+robot.vel.x*t, nCols)
		robot.pos.y = mod(robot.pos.y+robot.vel.y*t, nRows)
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

// Go modulo behavior is diff from expected
func mod(a int, b int) int {
	return (a%b + b) % b
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

	initGrid := make([][]string, nRows)
	for i := range initGrid {
		initGrid[i] = make([]string, nCols)
	}
	for i, row := range initGrid {
		for j := range row {
			initGrid[i][j] = "."
		}
	}
	for _, r := range robots {
		initGrid[r.pos.y][r.pos.x] = "#"
	}

	grid := make([][]string, nRows)
	for i := range grid {
		grid[i] = make([]string, nCols)
	}
	for i, row := range grid {
		for j := range row {
			grid[i][j] = "."
		}
	}

	maxMeanNbrs := float64(0)
	bestIter := 0
	// NOTE: turns out that after nRows*nCols time steps, grid returns to original state
	iters := nRows * nCols
	for t := 1; t <= iters; t++ {
		for _, robot := range robots {
			robot.move(nRows, nCols)
		}

		for i, row := range grid {
			for j := range row {
				grid[i][j] = "."
			}
		}
		for _, r := range robots {
			grid[r.pos.y][r.pos.x] = "#"
		}

		meanNbrs := meanNeighbors(grid)
		if meanNbrs > maxMeanNbrs {
			maxMeanNbrs = meanNbrs
			bestIter = t
		}
	}

	// if matches(grid, initGrid) {
	// 	printGrid(initGrid)
	// 	fmt.Println("---------------------------------------------------")
	// 	printGrid(grid)
	// }
	return bestIter
}

func matches(cand [][]string, target [][]string) bool {
	for i, row := range cand {
		for j := range row {
			if cand[i][j] != target[i][j] {
				return false
			}
		}
	}
	return true
}

func meanNeighbors(grid [][]string) float64 {
	dirs := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	var nbrList []int
	for i, row := range grid {
		for j := range row {
			if grid[i][j] == "#" {
				nbrs := 0
				for _, dir := range dirs {
					if (i+dir[0] >= 0 && i+dir[0] < len(grid) && j+dir[1] >= 0 && j+dir[1] < len(grid[0])) && (grid[i+dir[0]][j+dir[1]] == "#") {
						nbrs++
					}
				}
				nbrList = append(nbrList, nbrs)
			}
		}
	}

	sum := 0
	for _, num := range nbrList {
		sum += num
	}
	return float64(sum) / float64(len(nbrList))
}
