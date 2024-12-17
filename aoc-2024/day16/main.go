package main

import (
	"bufio"
	"container/heap"
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
	maze := parseMaze(scanner)

	var start pos
	for i, row := range maze {
		for j := range row {
			if maze[i][j] == "S" {
				// init facing E
				start = pos{i: i, j: j, di: 0, dj: 1, index: 0, cost: 0}
				break
			}
		}
	}

	search := make(minHeap, 1)
	search[0] = &start
	heap.Init(&search)
	// NOTE: turns out using []coord and a slices.Contains() check for seen is
	// much slower than this string-boolean map
	seen := make(map[coord]bool, (len(maze)-2)*(len(maze[0])-2))
	costs := make(map[coord]int, (len(maze)-2)*(len(maze[0])-2))

	var cost int
	for len(search) > 0 {
		cur := heap.Pop(&search).(*pos)
		seen[coord{i: cur.i, j: cur.j, di: cur.di, dj: cur.dj}] = true
		if maze[cur.i][cur.j] == "E" {
			cost = cur.cost
			break
		}
		// (i, j) -> (j, -i) for CW and (-j, i) for CCW
		dirs := []pos{
			{i: cur.i + cur.di, j: cur.j + cur.dj, di: cur.di, dj: cur.dj, cost: cur.cost + 1},
			{i: cur.i + cur.dj, j: cur.j + (-cur.di), di: cur.dj, dj: -cur.di, cost: cur.cost + 1000 + 1},
			{i: cur.i + (-cur.dj), j: cur.j + cur.di, di: -cur.dj, dj: cur.di, cost: cur.cost + 1000 + 1},
		}
		for _, dir := range dirs {
			if maze[dir.i][dir.j] == "#" {
				continue
			}
			if seen[coord{i: dir.i, j: dir.j, di: dir.di, dj: dir.dj}] {
				continue
			}
			if dir.cost < getMapDefault(costs, coord{i: dir.i, j: dir.j, di: dir.di, dj: dir.dj}, 999999999999) {
				if dir.i == 7 && dir.j == 5 {
					log.Printf("LESS: cur=%v dir.cost=%d map hit=%d\n", cur, dir.cost, getMapDefault(costs, coord{i: dir.i, j: dir.j, di: dir.di, dj: dir.dj}, 999999999999))
				}
				costs[coord{i: dir.i, j: dir.j, di: dir.di, dj: dir.dj}] = dir.cost
				heap.Push(&search, &dir)
			}
		}
	}
	return cost
}

type coord struct {
	i  int
	j  int
	di int
	dj int
}

type pos struct {
	i     int
	j     int
	di    int
	dj    int
	index int
	cost  int
}

type minHeap []*pos

func (mh *minHeap) Len() int {
	return len(*mh)
}

func (mh *minHeap) Less(i int, j int) bool {
	h := *mh
	return h[i].cost < h[j].cost
}

func (mh *minHeap) Swap(i int, j int) {
	if len(*mh) > 0 {
		h := *mh
		h[i], h[j] = h[j], h[i]
		h[i].index = i
		h[j].index = j
	}
}

func (mh *minHeap) Push(p any) {
	n := len(*mh)
	item := p.(*pos)
	item.index = n
	*mh = append(*mh, item)
}

// NOTE: should check if heap is empty
func (mh *minHeap) Pop() any {
	old := *mh
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*mh = old[0 : n-1]
	return item
}

func parseMaze(scanner *bufio.Scanner) [][]string {
	var maze [][]string
	for scanner.Scan() {
		maze = append(maze, strings.Split(scanner.Text(), ""))
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return maze
}

func printMaze(m [][]string) {
	for _, row := range m {
		fmt.Println(row)
	}
}

func getMapDefault(m map[coord]int, key coord, defaultValue int) int {
	if value, exists := m[key]; exists {
		return value
	}
	return defaultValue
}

func solvePart2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	maze := parseMaze(scanner)

	var start pos
	for i, row := range maze {
		for j := range row {
			if maze[i][j] == "S" {
				start = pos{i: i, j: j, di: 0, dj: 1, index: 0, cost: 0}
				break
			}
		}
	}

	search := make(minHeap, 1)
	search[0] = &start
	heap.Init(&search)
	costs := make(map[coord]int, (len(maze)-2)*(len(maze[0])-2))
	trail := make(map[coord][]coord, (len(maze)-2)*(len(maze[0])-2))
	var end coord

	for len(search) > 0 {
		cur := heap.Pop(&search).(*pos)
		if maze[cur.i][cur.j] == "E" {
			end = coord{i: cur.i, j: cur.j, di: cur.di, dj: cur.dj}
			break
		}
		dirs := []pos{
			{i: cur.i + cur.di, j: cur.j + cur.dj, di: cur.di, dj: cur.dj, cost: cur.cost + 1},
			{i: cur.i + cur.dj, j: cur.j + (-cur.di), di: cur.dj, dj: -cur.di, cost: cur.cost + 1000 + 1},
			{i: cur.i + (-cur.dj), j: cur.j + cur.di, di: -cur.dj, dj: cur.di, cost: cur.cost + 1000 + 1},
		}
		for _, dir := range dirs {
			if maze[dir.i][dir.j] == "#" {
				continue
			}

			if dir.cost < getMapDefault(costs, coord{i: dir.i, j: dir.j, di: dir.di, dj: dir.dj}, 999999999999) {
				costs[coord{i: dir.i, j: dir.j, di: dir.di, dj: dir.dj}] = dir.cost
				trail[coord{i: dir.i, j: dir.j, di: dir.di, dj: dir.dj}] = []coord{{i: cur.i, j: cur.j, di: cur.di, dj: cur.dj}}
				heap.Push(&search, &dir)
			} else if dir.cost == getMapDefault(costs, coord{i: dir.i, j: dir.j, di: dir.di, dj: dir.dj}, 999999999999) {
				exist := trail[coord{i: dir.i, j: dir.j, di: dir.di, dj: dir.dj}]
				if !slices.Contains(exist, coord{i: cur.i, j: cur.j, di: cur.di, dj: cur.dj}) {
					trail[coord{i: dir.i, j: dir.j, di: dir.di, dj: dir.dj}] = append(trail[coord{i: dir.i, j: dir.j, di: dir.di, dj: dir.dj}], coord{i: cur.i, j: cur.j, di: cur.di, dj: cur.dj})
				}
			}
		}
	}

	stack := []coord{end}
	good := []coord{end}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		origins := trail[cur]

		maze[cur.i][cur.j] = "O"

		for _, back := range origins {
			if !slices.Contains(good, back) {
				good = append(good, back)
				stack = append(stack, back)
			}
		}
	}

	ans := 0
	for _, row := range maze {
		for _, col := range row {
			if col == "O" {
				ans++
			}
		}
	}
	return ans
}
