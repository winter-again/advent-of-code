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

	var numSafe int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r := strings.Fields(scanner.Text())
		report := parseReport(r)
		if reportSafe(report) {
			numSafe += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return numSafe
}

func reportSafe(report []int) bool {
	i := 0
	j := 1
	var inc bool
	if report[j] > report[i] {
		inc = true
	} else if report[j] < report[i] {
		inc = false
	} else {
		return false
	}

	if (absDiff(report[i], report[j]) < 1) || (absDiff(report[i], report[j]) > 3) {
		return false
	}

	safe := true
	for {
		i += 1
		j += 1
		if j > len(report)-1 {
			break
		}

		var incN bool
		if report[j] > report[i] {
			incN = true
		} else if report[j] < report[i] {
			incN = false
		} else {
			safe = false
			break
		}
		if incN != inc {
			safe = false
			break
		}

		if (absDiff(report[i], report[j]) < 1) || (absDiff(report[i], report[j]) > 3) {
			safe = false
			break
		}
	}
	return safe
}

func solvePart2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var numSafe int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r := strings.Fields(scanner.Text())
		report := parseReport(r)

		if reportSafe(report) {
			numSafe += 1
		} else {
			for i := 0; i < len(report); i++ {
				if reportSafe(remove(report, i)) {
					numSafe += 1
					break
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return numSafe
}

func remove(r []int, i int) []int {
	ret := make([]int, 0)
	ret = append(ret, r[:i]...)
	return append(ret, r[i+1:]...)
}

func parseReport(r []string) []int {
	report := make([]int, len(r))
	for i, v := range r {
		val, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		report[i] = val
	}
	return report
}

func absDiff(x int, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
