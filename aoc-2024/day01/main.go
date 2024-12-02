package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
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

	var a []int
	var b []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Fields(scanner.Text())
		a_val, _ := strconv.Atoi(values[0])
		b_val, _ := strconv.Atoi(values[1])
		a = append(a, a_val)
		b = append(b, b_val)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	slices.Sort(a)
	slices.Sort(b)

	dist := 0
	for i := 0; i < len(a); i++ {
		dist = dist + absDiff(a[i], b[i])
	}
	return dist
}

func absDiff(x int, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func solvePart2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var a []int
	var b []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Fields(scanner.Text())
		a_val, _ := strconv.Atoi(values[0])
		b_val, _ := strconv.Atoi(values[1])
		a = append(a, a_val)
		b = append(b, b_val)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	// O(n^2); hash map would make it O(n)?
	counts := make([]int, len(a))
	for i, a_val := range a {
		for _, b_val := range b {
			if b_val == a_val {
				counts[i]++
			}
		}
	}

	sim := 0
	for i, val := range a {
		sim = sim + (val * counts[i])
	}
	return sim
}
