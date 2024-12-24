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

	comps := parseComputers(file)
	lans := findLAN(comps)
	return len(lans)
}

func findLAN(comps map[string][]string) [][]string {
	var lans [][]string
	for pc1, nbrs1 := range comps {
		for _, pc2 := range nbrs1 {
			for _, pc3 := range comps[pc2] {
				if slices.Contains(comps[pc3], pc1) && (strings.HasPrefix(pc1, "t") || strings.HasPrefix(pc2, "t") || strings.HasPrefix(pc3, "t")) {
					lan := []string{pc1, pc2, pc3}
					if !lanExists(lans, lan) {
						lans = append(lans, lan)
					}
				}
			}
		}
	}
	return lans
}

func lanExists(lans [][]string, lan []string) bool {
	for _, l := range lans {
		if lanEqual(l, lan) {
			return true
		}
	}
	return false
}

func lanEqual(x []string, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	diff := make(map[string]int, len(x))
	for _, _x := range x {
		diff[_x]++
	}
	for _, _y := range y {
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y]--
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	return len(diff) == 0
}

func parseComputers(file *os.File) map[string][]string {
	adj := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		comps := strings.Split(scanner.Text(), "-")
		adj[comps[0]] = append(adj[comps[0]], comps[1])
		adj[comps[1]] = append(adj[comps[1]], comps[0])
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return adj
}

func solvePart2(input string) string {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	comps := parseComputers(file)
	lans := make(map[string]bool)
	for pc := range comps {
		findLANRec(pc, []string{pc}, comps, lans)
	}

	var largestLan string
	for lan := range lans {
		if len(lan) > len(largestLan) {
			largestLan = lan
		}
	}
	return largestLan
}

func findLANRec(pc string, curLan []string, adj map[string][]string, lans map[string]bool) {
	slices.Sort(curLan)
	concat := strings.Join(curLan, ",")
	if lans[concat] {
		return
	}
	lans[concat] = true

	nbrs := adj[pc]
	for _, nbr := range nbrs {
		if slices.Contains(curLan, nbr) {
			continue
		}

		// if connToCurLan false, then found a neighbor that isn't connected
		// to each member of curr set -> disregard because it can't make our LAN larger
		connToCurLan := true
		for _, r := range curLan {
			if !slices.Contains(adj[r], nbr) {
				connToCurLan = false
			}
		}
		if connToCurLan {
			nReq := append(curLan, nbr)
			findLANRec(nbr, nReq, adj, lans)
		}
	}
}
