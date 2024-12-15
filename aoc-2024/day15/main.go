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
		fmt.Println("Answer (sample 2):", ans)

		ans = solvePart1("./input.txt")
		fmt.Println("Answer:", ans)
	} else {
		ans := solvePart2("./input_smpl.txt")
		fmt.Println("Answer (sample):", ans)

		// ans = solvePart2("./input_manual.txt")
		// fmt.Println("Answer (manual):", ans)

		ans = solvePart2("./input.txt")
		fmt.Println("Answer:", ans)
	}
}

type grid [][]string

type pos struct {
	i int
	j int
}

var dirs = map[string]pos{
	"^": {-1, 0},
	">": {0, 1},
	"v": {1, 0},
	"<": {0, -1},
}

func solvePart1(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid, moves := parseDoc(file, 1)
	robot := grid.robotStart()

	// printGrid(grid)

	cur := robot
	for _, move := range moves {
		dir := dirs[move]
		// log.Println("Dir:", move, dir)
		next := pos{cur.i + dir.i, cur.j + dir.j}
		// TODO: make this move a method on grid? have to use pointers?
		if grid[next.i][next.j] == "#" {
			// printGrid(grid)
			continue
		}

		if grid[next.i][next.j] == "." {
			grid[next.i][next.j] = "@"
			grid[cur.i][cur.j] = "."
			cur = next
		} else if grid[next.i][next.j] == "O" {
			grpLen := 0
			for grid[next.i+grpLen*dir.i][next.j+grpLen*dir.j] == "O" {
				grpLen++
			}
			// log.Println("grpLen:", grpLen)

			// firstBox := next
			lastBox := pos{next.i + (grpLen-1)*dir.i, next.j + (grpLen-1)*dir.j}
			// log.Println(firstBox, lastBox)

			if grid[lastBox.i+dir.i][lastBox.j+dir.j] != "#" {
				grid[lastBox.i+dir.i][lastBox.j+dir.j] = "O"
				grid[cur.i+dir.i][cur.j+dir.j] = "@"
				grid[cur.i][cur.j] = "."
				cur = next
			}
		}
		// printGrid(grid)
	}

	sum := 0
	for i, row := range grid {
		for j := range row {
			if grid[i][j] == "O" {
				sum += 100*i + j
			}
		}
	}
	return sum
}

func (grid grid) robotStart() pos {
	for i, row := range grid {
		for j, col := range row {
			if col == "@" {
				return pos{i, j}
			}
		}
	}
	return pos{}
}

func parseDoc(file *os.File, part int) (grid, []string) {
	scanner := bufio.NewScanner(file)
	scanner.Split(splitDoc)

	scanner.Scan()
	lines := strings.Split(strings.TrimSpace(scanner.Text()), "\n")
	var mp [][]string

	for _, line := range lines {
		if part == 1 {
			mp = append(mp, strings.Split(line, ""))
		} else {
			var dblLine []string
			tiles := strings.Split(line, "")
			for _, tile := range tiles {
				if tile == "O" {
					dblLine = append(dblLine, "[", "]")
				} else if tile == "@" {
					dblLine = append(dblLine, "@", ".")
				} else {
					dblLine = append(dblLine, tile, tile)
				}
			}
			mp = append(mp, dblLine)
		}
	}

	scanner.Scan()
	moves := strings.Split(strings.ReplaceAll(strings.TrimSpace(scanner.Text()), "\n", ""), "")

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return mp, moves
}

func splitDoc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := strings.Index(string(data), "\n\n"); i >= 0 {
		return i + 1, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}
	return
}

func printGrid(mp [][]string) {
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

	grid, moves := parseDoc(file, 2)
	robot := grid.robotStart()

	// printGrid(grid)

	cur := robot
	for _, move := range moves {
		dir := dirs[move]
		next := pos{cur.i + dir.i, cur.j + dir.j}

		// log.Printf("t=%d move=%s, dir=%d\n", t, move, dir)
		if grid[next.i][next.j] == "#" {
			continue
		}
		if grid[next.i][next.j] == "." {
			grid[next.i][next.j] = "@"
			grid[cur.i][cur.j] = "."
			cur = next
		} else if grid[next.i][next.j] == "[" || grid[next.i][next.j] == "]" {
			boxChar := grid[next.i][next.j]
			if move == "<" || move == ">" {
				grpLen := 0
				for grid[next.i+grpLen*dir.i][next.j+grpLen*dir.j] == boxChar {
					grpLen += 2
				}
				lastBox := pos{cur.i + grpLen*dir.i, cur.j + grpLen*dir.j}

				if boxChar == "[" && grid[lastBox.i+dir.i][lastBox.j+dir.j] != "#" {
					// log.Println("MOVE: RIGHT")
					for s := 0; s < grpLen; s++ {
						if grid[next.i+s*dir.i][next.j+s*dir.j] == "[" {
							grid[next.i+s*dir.i][next.j+s*dir.j] = "]"
						} else {
							grid[next.i+s*dir.i][next.j+s*dir.j] = "["
						}
					}
					grid[lastBox.i+dir.i][lastBox.j+dir.j] = "]"
					grid[cur.i+dir.i][cur.j+dir.j] = "@"
					grid[cur.i][cur.j] = "."
					cur = next
				} else if boxChar == "]" && grid[lastBox.i+dir.i][lastBox.j+dir.j] != "#" {
					// log.Println("MOVE: LEFT")
					for s := 0; s < grpLen; s++ {
						if grid[next.i+s*dir.i][next.j+s*dir.j] == "[" {
							grid[next.i+s*dir.i][next.j+s*dir.j] = "]"
						} else {
							grid[next.i+s*dir.i][next.j+s*dir.j] = "["
						}
					}
					grid[lastBox.i+dir.i][lastBox.j+dir.j] = "["
					grid[next.i][next.j] = "@"
					grid[cur.i][cur.j] = "."
					cur = next
				}
			} else {
				boxHalves := findBoxHalves(next, dir, grid)
				canMove := true
				for halfPos := range boxHalves {
					if grid[halfPos.i+dir.i][halfPos.j+dir.j] == "#" {
						canMove = false
						break
					}
				}

				if canMove {
					for halfPos := range boxHalves {
						grid[halfPos.i][halfPos.j] = "."
					}
					for halfPos, halfChar := range boxHalves {
						grid[halfPos.i+dir.i][halfPos.j+dir.j] = halfChar
					}

					grid[next.i][next.j] = "@"
					grid[cur.i][cur.j] = "."
					cur = next
				}
			}
		}
		// printGrid(grid)
	}

	sum := 0
	for i, row := range grid {
		for j := range row {
			if grid[i][j] == "[" {
				sum += 100*i + j
			}
		}
	}
	return sum

}

func findBoxHalves(start pos, move pos, grid [][]string) map[pos]string {
	boxes := make(map[pos]string, 100)
	var seen []pos
	q := []pos{start}
	allowDirs := []pos{move, {0, -1}, {0, 1}}
	for {
		if len(q) == 0 {
			break
		}
		// pop from queue
		cur := q[0]
		q = q[1:]
		for _, dir := range allowDirs {
			boxChar := grid[cur.i][cur.j]
			n := pos{cur.i + dir.i, cur.j + dir.j}
			if grid[n.i][n.j] == "#" || grid[n.i][n.j] == "." {
				continue
			}
			if slices.Contains(seen, n) {
				continue
			}
			if dir.j == -1 && boxChar == "[" && grid[n.i][n.j] == "]" {
				continue
			}
			if dir.j == 1 && boxChar == "]" && grid[n.i][n.j] == "[" {
				continue
			}

			boxes[n] = grid[n.i][n.j]
			seen = append(seen, n)
			q = append(q, n)
		}
	}
	return boxes
}
