package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
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

	scanner := bufio.NewScanner(file)
	eqs := parseEquations(scanner)

	var res int
	for _, eq := range eqs {
		if ok := tryEquation(&eq, eq.nums[0], 1, len(eq.nums)); ok {
			res += eq.target
			continue
		}
	}
	return res
}

func tryEquation(eq *equation, res int, i int, n int) bool {
	if i == n {
		if res == eq.target {
			return true
		}
		return false
	}
	plus := tryEquation(eq, res+eq.nums[i], i+1, n)
	multi := tryEquation(eq, res*eq.nums[i], i+1, n)

	if plus || multi {
		return true
	}
	return false
}

type equation struct {
	target int
	nums   []int
}

func parseEquations(s *bufio.Scanner) []equation {
	var eqs []equation
	for s.Scan() {
		var eq equation
		ln := strings.Split(s.Text(), ":")
		eq.target, _ = strconv.Atoi(ln[0])

		numStr := strings.Fields(ln[1])
		nums := make([]int, 0, len(numStr))
		for _, num := range numStr {
			n, _ := strconv.Atoi(num)
			nums = append(nums, n)
		}
		eq.nums = nums
		eqs = append(eqs, eq)
	}

	if err := s.Err(); err != nil {
		log.Println(err)
	}
	return eqs
}

func solvePart2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	eqs := parseEquations(scanner)

	var res int
	for _, eq := range eqs {
		if ok := tryEquationConcat(&eq, eq.nums[0], 1, len(eq.nums)); ok {
			res += eq.target
			continue
		}
	}
	return res
}

func tryEquationConcat(eq *equation, res int, i int, n int) bool {
	if i == n {
		if res == eq.target {
			return true
		}
		return false
	}

	plus := tryEquationConcat(eq, res+eq.nums[i], i+1, n)
	multi := tryEquationConcat(eq, res*eq.nums[i], i+1, n)

	// concat via string builder
	// var sb strings.Builder
	// sb.WriteString(strconv.Itoa(res))
	// sb.WriteString(strconv.Itoa(eq.nums[i]))
	// interm, _ := strconv.Atoi(sb.String())

	interm := res*(int(math.Pow10(len(strconv.Itoa(eq.nums[i]))))) + eq.nums[i]
	concat := tryEquationConcat(eq, interm, i+1, n)
	if plus || multi || concat {
		return true
	}
	return false
}
