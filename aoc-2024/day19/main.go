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

	towels, patterns := parseTowels(file)
	tm := make(map[rune][]string)
	for _, towel := range towels {
		c := rune(towel[0])
		tm[c] = append(tm[c], towel)
	}

	res := 0
	cache := make(map[string]bool)
	for _, patt := range patterns {
		if tryTowel(patt, tm, cache) {
			res++
		}
	}
	return res
}

func tryTowel(design string, tm map[rune][]string, cache map[string]bool) bool {
	if len(design) == 0 {
		return true
	}
	_, ok := cache[design]
	if ok {
		return cache[design]
	}
	first := rune(design[0])
	patterns := tm[first]
	for _, patt := range patterns {
		if len(patt) > len(design) {
			continue
		}
		if design[:len(patt)] == patt && tryTowel(design[len(patt):], tm, cache) {
			cache[design] = true
			return true
		}
	}
	cache[design] = false
	return false
}

func parseTowels(file *os.File) ([]string, []string) {
	scanner := bufio.NewScanner(file)
	scanner.Split(splitTowels)

	scanner.Scan()
	towels := strings.Split(scanner.Text(), ", ")

	scanner.Scan()
	designs := strings.Split(strings.TrimSpace(scanner.Text()), "\n")

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return towels, designs
}

func splitTowels(data []byte, atEOF bool) (advance int, token []byte, err error) {
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

	towels, patterns := parseTowels(file)
	tm := make(map[rune][]string)
	for _, towel := range towels {
		c := rune(towel[0])
		tm[c] = append(tm[c], towel)
	}

	res := 0
	cache := make(map[string]int)
	for _, patt := range patterns {
		res += countTowels(patt, tm, cache)
	}
	return res
}

func countTowels(design string, tm map[rune][]string, cache map[string]int) int {
	if len(design) == 0 {
		return 1
	}
	ct, ok := cache[design]
	if ok {
		return ct
	}

	add := 0
	first := rune(design[0])
	patterns := tm[first]
	for _, patt := range patterns {
		if len(patt) > len(design) {
			continue
		}
		if design[:len(patt)] == patt {
			more := countTowels(design[len(patt):], tm, cache)
			cache[design] += more
			add += more
		}
	}
	return add
}
