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
	} else {
		ans := solvePart2("./input_smpl.txt")
		fmt.Println("Answer (sample):", ans)

		// ans = solvePart2("./input.txt")
		// fmt.Println("Answer:", ans)
	}
}

func solvePart1(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	locks, keys := parseSchematics(file)

	matches := make(map[int][]int)
	for i, lock := range locks {
		for j, key := range keys {
			if keyFitsLock(key, lock) && !slices.Contains(matches[i], j) {
				matches[i] = append(matches[i], j)
			}
		}
	}

	// log.Println(matches)

	sum := 0
	for _, keys := range matches {
		sum += len(keys)
	}
	return sum
}

type schematic [][]string

func keyFitsLock(key schematic, lock schematic) bool {
	keyHeights := make([]int, len(key[0]))
	lockHeights := make([]int, len(lock[0]))

	for _, row := range key[:len(key)-1] {
		for j, col := range row {
			if col == "#" {
				keyHeights[j]++
			}
		}
	}

	for _, row := range lock[1:] {
		for j, col := range row {
			if col == "#" {
				lockHeights[j]++
			}
		}
	}

	totHeight := len(key) - 2
	for i := 0; i < len(key[0]); i++ {
		if keyHeights[i]+lockHeights[i] > totHeight {
			return false
		}
	}
	return true
}

func parseSchematics(file *os.File) ([]schematic, []schematic) {
	scanner := bufio.NewScanner(file)
	scanner.Split(splitSchematics)

	var locks []schematic
	var keys []schematic

	for scanner.Scan() {
		var schem schematic
		lines := strings.Split(strings.TrimSpace(scanner.Text()), "\n")
		for _, line := range lines {
			row := strings.Split(line, "")
			schem = append(schem, row)
		}

		isLock := true
		for _, c := range schem[0] {
			if c != "#" {
				isLock = false
				break
			}
		}

		if isLock {
			locks = append(locks, schem)
		} else {
			keys = append(keys, schem)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return locks, keys
}

func printSchematic(s schematic) {
	for _, row := range s {
		fmt.Println(row)
	}
}

func splitSchematics(data []byte, atEOF bool) (advance int, token []byte, err error) {
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
func solvePart2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return 0
}
