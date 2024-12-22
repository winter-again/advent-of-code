package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
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

	initNums := parseInitNums(file)

	sum := 0
	for _, init := range initNums {
		in := init
		for j := 0; j < 2000; j++ {
			in = evolve(in)
		}
		sum += in
	}
	return sum
}

func evolve(secret int) int {
	s1 := (secret ^ (secret * 64)) % 16777216
	s2 := (s1 ^ (s1 / 32)) % 16777216
	s3 := (s2 ^ (s2 * 2048)) % 16777216
	return s3
}

func parseInitNums(file *os.File) []int {
	scanner := bufio.NewScanner(file)
	var nums []int
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return nums
}

// TODO: for each buyer:
// calc prices based on the secret num
// calc price diffs
// (there's prob some pattern?)
// then have to find a len 4 seq common to ALL buyers
// such that the first occurrence of the seq ends on high price so that the sum is maximized
func solvePart2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	initNums := parseInitNums(file)

	prices := make(map[int][]int, len(initNums))
	numSecrets := 2000
	for _, init := range initNums {
		prices[init] = append(prices[init], init%10)
		secret := init
		for n := 0; n < numSecrets-1; n++ {
			s := evolve(secret)
			prices[init] = append(prices[init], s%10)
			secret = s
		}
	}

	diffs := make(map[int][]int, len(initNums))
	for k, v := range prices {
		for i := 1; i < len(v); i++ {
			diffs[k] = append(diffs[k], v[i]-v[i-1])
		}
	}

	// NOTE: this works but is very slow
	candDiff := diffs[initNums[0]]
	maxBananas := 0
	for i := 0; i < len(candDiff)-3; i++ {
		targSeq := candDiff[i : i+4]
		bananas := prices[initNums[0]][i+4]
		for init, diff := range diffs {
			if init == initNums[0] {
				continue
			}
			loc, err := findSeq(diff, targSeq)
			if err != nil {
				continue
			}
			price := prices[init][loc+1]
			bananas += price
		}
		if bananas > maxBananas {
			maxBananas = bananas
		}
	}
	return maxBananas
}

func findSeq(changes []int, seq []int) (int, error) {
	var loc int
	for i := 0; i < len(changes)-3; i++ {
		if slices.Equal(changes[i:i+4], seq) {
			loc = i + 3
			return loc, nil
		}
	}
	return loc, fmt.Errorf("Not found: %v", seq)
}

func printMap(m map[int][]int) {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}
