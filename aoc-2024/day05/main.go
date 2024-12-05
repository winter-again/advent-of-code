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

	scanner := bufio.NewScanner(file)
	var rawRules []string
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			break
		}
		rawRules = append(rawRules, scanner.Text())
	}

	var rawUpdates []string
	for scanner.Scan() {
		rawUpdates = append(rawUpdates, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	rules := parseRules(rawRules, false)
	updates := parseUpdates(rawUpdates)

	sum := 0
	for _, updt := range updates {
		ok := true
		for i, u := range updt {
			rule := rules[u]
			cands := updt[i+1:]
			for _, cand := range cands {
				if !slices.Contains(rule, cand) {
					ok = false
					break
				}
			}
		}
		if ok {
			sum += midpt(updt)
		}
	}
	return sum
}

func parseRules(r []string, rev bool) map[int][]int {
	rules := make(map[int][]int)
	for _, rul := range r {
		rr := strings.Split(rul, "|")
		x, _ := strconv.Atoi(rr[0])
		y, _ := strconv.Atoi(rr[1])
		if rev {
			rules[y] = append(rules[y], x)
		} else {
			rules[x] = append(rules[x], y)
		}
	}
	return rules
}

func parseUpdates(u []string) [][]int {
	updates := make([][]int, len(u))
	for i, upd := range u {
		uu := strings.Split(upd, ",")
		var uInt []int
		for _, uint := range uu {
			val, _ := strconv.Atoi(uint)
			uInt = append(uInt, val)
		}
		updates[i] = uInt
	}
	return updates
}

func midpt(s []int) int {
	x := (len(s) - 1) / 2
	return s[x]
}

func solvePart2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rawRules []string
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			break
		}
		rawRules = append(rawRules, scanner.Text())
	}

	var rawUpdates []string
	for scanner.Scan() {
		rawUpdates = append(rawUpdates, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	rules := parseRules(rawRules, false)
	updates := parseUpdates(rawUpdates)

	var bads [][]int
	for _, updt := range updates {
		bad := false
		for i, u := range updt {
			rule := rules[u]
			cands := updt[i+1:]
			for _, cand := range cands {
				if !slices.Contains(rule, cand) {
					bad = true
					break
				}
			}
		}
		if bad {
			bads = append(bads, updt)
		}
	}

	rules = parseRules(rawRules, true)

	for _, bad := range bads {
		// log.Println("FIXING >", bad)
		n := len(bad)
		var sorted bool
		for {
			if sorted {
				break
			}

			sorted = true
			for i := 0; i < n-1; i++ {
				rule := rules[bad[i]]
				// log.Println("COMP >", bad[i], bad[i+1])
				if slices.Contains(rule, bad[i+1]) {
					sorted = false
					tmp := bad[i]
					bad[i] = bad[i+1]
					bad[i+1] = tmp
					// log.Println("SWAPPED >", bad)
				}
			}
			n = n - 1
		}
	}

	sum := 0
	for _, bad := range bads {
		sum += midpt(bad)
	}
	return sum
}
