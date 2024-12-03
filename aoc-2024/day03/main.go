package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
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
		ans := solvePart2("./input_smpl_2.txt")
		fmt.Println("Answer (sample):", ans)

		ans = solvePart2("./input.txt")
		fmt.Println("Answer:", ans)
	}
}

func solvePart1(input string) int {
	mem, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	memStr := string(mem)
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	parsedMem := re.FindAllString(memStr, -1)

	res := make([]int, len(parsedMem))
	for i, op := range parsedMem {
		res[i] = mul(op)
	}
	return sum(res)
}

func mul(expr string) int {
	re := regexp.MustCompile(`\d+,\d+`)
	pair := re.FindString(expr)
	nums := strings.Split(pair, ",")

	x, err := strconv.Atoi(nums[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(nums[1])
	if err != nil {
		log.Fatal(err)
	}
	return x * y
}

func sum(s []int) int {
	sum := 0
	for _, x := range s {
		sum += x
	}
	return sum
}

func solvePart2(input string) int {
	mem, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	memStr := string(mem)
	re := regexp.MustCompile(`(mul\(\d+,\d+\))|(do\(\))|(don't\(\))`)
	parsedMem := re.FindAllString(memStr, -1)

	var dont bool
	doSwitch := "do()"
	dontSwitch := "don't()"

	var validOps []string
	for _, op := range parsedMem {
		if op == dontSwitch {
			dont = true
		} else if op == doSwitch {
			dont = false
		} else {
			if !dont {
				validOps = append(validOps, op)
			}
		}
	}

	res := make([]int, len(validOps))
	for i, op := range validOps {
		res[i] = mul(op)
	}
	return sum(res)
}
